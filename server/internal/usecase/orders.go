package usecase

import (
	"fmt"
	"laundry/internal/entity/orders"
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/entity/services"
	"laundry/internal/entity/units"
	"laundry/internal/repository/rimport"
	"laundry/tools/appdate"
	"laundry/tools/appmath"
	"laundry/tools/slice"
	"laundry/tools/sqlnull"
	transactiongeneric "laundry/tools/transaction-generic"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OrdersUsecase struct {
	repo *rimport.RepositoryImport
	db   *sqlx.DB
}

func NewOrdersUsecase(
	repo *rimport.RepositoryImport,
	db *sqlx.DB,
) *OrdersUsecase {
	return &OrdersUsecase{
		repo,
		db,
	}
}

// CalculateOrder расчитывает цену заказа с учетом цен за вещи в каждой услуге,
// применяет скидки если условия скидок были выполнены
func (u *OrdersUsecase) CalculateOrder(param orders.CalculateOrderParam) (orders.CalculateOrderResponse, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) (res orders.CalculateOrderResponse, err error) {
			// получение всех доступных вещей с их ценами
			serviceItems, err := u.repo.Services.FindAllServiceItems(tx)
			if err != nil {
				return
			}
			// получение модификаторов цен единиц измерения (шт, кг)
			unitModifiers, err := u.repo.PriceModifiers.FindAllUnitModifiers(tx)
			if err != nil {
				return
			}
			// получение модификаторов цен типа вещей (врослые, детские)
			itemTypeModifiers, err := u.repo.PriceModifiers.FindAllItemTypeModifiers(tx)
			if err != nil {
				return
			}

			// превращение данных в map для последующей работы с ними
			ableItems := slice.Reduce(
				serviceItems,
				func(
					acc map[string]services.ServiceItems,
					item services.ServiceItems,
					index int,
				) map[string]services.ServiceItems {
					serviceId := item.ServiceID.GetInt()
					if item.SubServiceID.Valid {
						serviceId = item.SubServiceID.GetInt()
					}
					key := fmt.Sprintf("%d:%d", serviceId, item.ID)
					acc[key] = item
					return acc
				},
				make(map[string]services.ServiceItems, len(serviceItems)),
			)

			ableUnitModifiers := slice.Reduce(
				unitModifiers,
				func(
					acc map[int]pricemodifiers.UnitPriceModifier,
					unitM pricemodifiers.UnitPriceModifier,
					index int,
				) map[int]pricemodifiers.UnitPriceModifier {
					acc[unitM.UnitID] = unitM
					return acc
				},
				make(map[int]pricemodifiers.UnitPriceModifier, len(unitModifiers)),
			)

			ableItemTypeModifiers := slice.Reduce(
				itemTypeModifiers,
				func(
					acc map[int]pricemodifiers.PriceModifier,
					itemTypeM pricemodifiers.PriceModifier,
					index int,
				) map[int]pricemodifiers.PriceModifier {
					acc[itemTypeM.ID] = itemTypeM
					return acc
				},
				make(map[int]pricemodifiers.PriceModifier, len(itemTypeModifiers)),
			)

			var calculatedServices []orders.CalculateOrderResponseService

			// обработка каждой услуги заказа
			for _, service := range param.Services {
				processed := u.calculateSingleService(orders.CalculateSingleServiceParam{
					OrderedServices:       service,
					AbleItems:             ableItems,
					AbleUnitModifiers:     ableUnitModifiers,
					AbleItemTypeModifiers: ableItemTypeModifiers,
				})

				calculatedServices = append(calculatedServices, processed)
			}

			// расчет общей суммы всех услуг со скидкой и без
			reduce := slice.Reduce(
				calculatedServices,
				func(
					acc orders.CalculateOrderReduceResult,
					item orders.CalculateOrderResponseService,
					index int,
				) orders.CalculateOrderReduceResult {
					acc.Final += item.Final
					acc.Total += item.Total
					return acc
				},
				orders.CalculateOrderReduceResult{
					Total: 0,
					Final: 0,
				},
			)

			final := reduce.Final

			// применения общих скидок к заказу
			var fulfillmentPriceModifier pricemodifiers.PriceModifier
			var discounts []pricemodifiers.PriceModifierCommonData
			var markups []pricemodifiers.PriceModifierCommonData

			if param.Fulfillment.ModifierID.Valid {
				fulfillmentPriceModifier, err = u.repo.PriceModifiers.FindFulfillmentModifierByID(tx, param.Fulfillment.ModifierID.GetInt())
				if err != nil {
					return
				}

				modifier := pricemodifiers.PriceModifierCommonData{
					Percent:     fulfillmentPriceModifier.Percent,
					Description: fulfillmentPriceModifier.Description.String,
					Modifier:    fulfillmentPriceModifier.Modifier,
					ModifierID:  fulfillmentPriceModifier.ModifierID,
				}

				finalResult, isDiscount := u.countMarkupsAndDiscounts(
					modifier.Modifier,
					modifier.Percent,
					final,
				)

				final = finalResult

				if isDiscount {
					discounts = append(discounts, modifier)
				} else {
					markups = append(markups, modifier)
				}
			}

			res = orders.CalculateOrderResponse{
				TemporaryID:   uuid.New().String(),
				OrderServices: calculatedServices,
				Fulfillment:   param.Fulfillment,
				Discounts:     discounts,
				Markups:       markups,
				Total:         reduce.Total,
				Final:         appmath.RoundToDecimals(final, 1),
			}

			return res, nil
		},
		"Не удалось обработать заказ",
	)
}

// CreateOrder создает заказ
func (u *OrdersUsecase) CreateOrder(param orders.CreateOrderParamWithPreCalculatedData) (int, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) (orderID int, err error) {
			// создание заказа в таблице заказов
			orderID, err = u.repo.Orders.CreateOrder(tx, param.UserParam)
			if err != nil {
				return
			}

			// запись услуг заказа
			for _, service := range param.PreCalculatedData.OrderServices {
				err = u.processCreateOrderService(tx, orderID, service)
				if err != nil {
					return
				}
			}

			// запись примененных к заказу модификаторов цены
			for _, modifier := range append(param.PreCalculatedData.Discounts, param.PreCalculatedData.Markups...) {
				err = u.repo.Orders.CreateOrderPriceModifiersRecord(tx, orders.CreateOrderPriceModifiersRecord{
					Modifier:    modifier.Modifier,
					Description: modifier.Description,
					Percent:     modifier.Percent,
					OrderID:     sqlnull.NewInt64(orderID),
				})

				if err != nil {
					return
				}
			}

			return
		},
		"Не удалось создать заказ",
	)
}

// FindTodayOrders полученение заказов за сегодня
func (u *OrdersUsecase) FindTodayOrders() ([]orders.Order, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) ([]orders.Order, error) {
			param := orders.GetOrderByDateRangeParam{
				StartDate: appdate.GetStartOfDay(),
				EndDate:   appdate.GetEndOfDay(),
			}

			return u.repo.Orders.FindOrdersByDateRange(tx, param)
		},
		"Не удалось найти заказы за сегодня",
	)
}

// calculateSingleService обработка услуги и применение скидок
func (u *OrdersUsecase) calculateSingleService(param orders.CalculateSingleServiceParam) orders.CalculateOrderResponseService {
	chosenItems := []orders.ServiceCommonResponseItem{}

	// если подуслуга то используется id подуслуги
	serviceID := param.OrderedServices.ServiceID
	if param.OrderedServices.SubServiceID.Valid {
		serviceID = param.OrderedServices.SubServiceID.GetInt()
	}

	// получение выбранных вещей из доступных
	for _, chosenItem := range param.OrderedServices.Items {
		key := fmt.Sprintf("%d:%d", serviceID, chosenItem.ID)
		ableItem := param.AbleItems[key]
		chosenItems = append(chosenItems, orders.ServiceCommonResponseItem{
			ID:          ableItem.ID,
			ItemID:      ableItem.ItemID,
			ItemName:    ableItem.ItemName,
			PriceForOne: ableItem.Price,
			PriceForAll: appmath.RoundToDecimals(ableItem.Price*chosenItem.Quantity, 1),
			Quantity:    chosenItem.Quantity,
		})
	}

	var commonModifiers []pricemodifiers.PriceModifierCommonData

	// получение модификаторов цен типа вещей (врослые, детские)
	priceModifier, exists := param.AbleItemTypeModifiers[param.OrderedServices.ItemsTypeID]

	if exists {
		commonModifiers = append(commonModifiers, pricemodifiers.PriceModifierCommonData{
			Percent:     priceModifier.Percent,
			Description: priceModifier.Description.String,
			Modifier:    priceModifier.Modifier,
			ModifierID:  priceModifier.ModifierID,
		})
	}

	// получение модификаторов цен единиц измерения (шт, кг)
	var unitPriceModifierID int
	unitPriceModifier, exists := param.AbleUnitModifiers[param.OrderedServices.UnitID]

	reduced := slice.Reduce(
		chosenItems,
		func(
			acc orders.CalculateSingleServiceItemReduceResult,
			value orders.ServiceCommonResponseItem,
			index int,
		) orders.CalculateSingleServiceItemReduceResult {
			acc.TotalSum += value.PriceForAll
			acc.TotalUnitQuantity += value.Quantity
			return acc
		},
		orders.CalculateSingleServiceItemReduceResult{
			TotalSum:          0,
			TotalUnitQuantity: 0,
		},
	)

	if reduced.TotalUnitQuantity > unitPriceModifier.UnitQuantity && exists {
		commonModifiers = append(commonModifiers, pricemodifiers.PriceModifierCommonData{
			Percent:     unitPriceModifier.Percent,
			Description: unitPriceModifier.Description.String,
			Modifier:    unitPriceModifier.Modifier,
			ModifierID:  unitPriceModifier.ModifierID,
		})
		unitPriceModifierID = unitPriceModifier.ID
	}

	var discounts []pricemodifiers.PriceModifierCommonData
	var markups []pricemodifiers.PriceModifierCommonData

	// применение модификаторов цен
	final := reduced.TotalSum
	for _, modifer := range commonModifiers {
		result, isDiscount := u.countMarkupsAndDiscounts(
			modifer.Modifier,
			modifer.Percent,
			final,
		)

		final = result

		if isDiscount {
			discounts = append(discounts, modifer)
		} else {
			markups = append(markups, modifer)
		}
	}

	return orders.CalculateOrderResponseService{
		ServiceID:      param.OrderedServices.ServiceID,
		SubServiceID:   param.OrderedServices.SubServiceID,
		ServiceName:    param.OrderedServices.ServiceName,
		SubServiceName: param.OrderedServices.SubServiceName,
		Total:          appmath.RoundToDecimals(reduced.TotalSum, 1),
		Final:          appmath.RoundToDecimals(final, 1),
		Items:          chosenItems,
		Discounts:      discounts,
		Markups:        markups,
		UnitID:         param.OrderedServices.UnitID,
		UnitTitle:      units.UnitTitle[param.OrderedServices.UnitID],
		UnitModifierID: sqlnull.NewInt64(unitPriceModifierID),
		ItemsTypeID:    param.OrderedServices.ItemsTypeID,
	}
}

// processCreateOrderService записывает услуги и вещи у созданного заказа
func (u *OrdersUsecase) processCreateOrderService(
	tx *sqlx.Tx,
	orderID int,
	param orders.CalculateOrderResponseService,
) error {
	// запись услуги
	id, err := u.repo.Orders.CreateOrderServiceRecord(tx, orderID, param.ServiceID)
	if err != nil {
		return err
	}

	// запись вещей
	for _, item := range param.Items {
		err = u.repo.Orders.CreateOrderServiceItemRecord(tx, orders.CreateOrderServiceItemRecord{
			ServiceItemID:  item.ID,
			Quantity:       item.Quantity,
			Price:          item.PriceForOne,
			OrderServiceId: id,
		})

		if err != nil {
			return err
		}
	}

	// запись модификаторов цены у услуги
	for _, modifier := range append(param.Discounts, param.Markups...) {
		err = u.repo.Orders.CreateOrderPriceModifiersRecord(tx, orders.CreateOrderPriceModifiersRecord{
			Modifier:    modifier.Modifier,
			Description: modifier.Description,
			Percent:     modifier.Percent,
			ServiceID:   sqlnull.NewInt64(id),
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// countMarkupsAndDiscounts получает модификтор и процент, высчитывает или добавляет процент к сумме
func (u *OrdersUsecase) countMarkupsAndDiscounts(
	modifier int,
	percent float64,
	sum float64,
) (result float64, isDiscount bool) {

	switch modifier {
	case pricemodifiers.ModifierDiscount:
		isDiscount = true
		result = sum - appmath.CaclPercentFromSum(sum, percent)

	case pricemodifiers.ModifierMarkup:
		result = sum + appmath.CaclPercentFromSum(sum, percent)
	}

	return
}
