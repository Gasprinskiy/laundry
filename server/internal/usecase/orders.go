package usecase

import (
	"fmt"
	"laundry/internal/entity/orders"
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/entity/services"
	"laundry/internal/entity/units"
	"laundry/internal/repository/rimport"
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

func (u *OrdersUsecase) CalculateOrder(param orders.CalculateOrderParam) (orders.CalculateOrderResponse, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) (res orders.CalculateOrderResponse, err error) {
			serviceItems, err := u.repo.Services.FindAllServiceItems(tx)
			if err != nil {
				return
			}

			subServiceItems, err := u.repo.Services.FindAllSubServiceItems(tx)
			if err != nil {
				return
			}

			unitModifiers, err := u.repo.PriceModifiers.FindAllUnitModifiers(tx)
			if err != nil {
				return
			}

			priceModifiers, err := u.repo.PriceModifiers.FindAllItemTypeModifiers(tx)
			if err != nil {
				return
			}

			ableItems := make(map[string]services.ServiceItems)
			ableUnitModifiers := make(map[int]pricemodifiers.UnitPriceModifier)
			ablePriceModifiers := make(map[int]pricemodifiers.PriceModifier)

			for _, item := range append(serviceItems, subServiceItems...) {
				key := fmt.Sprintf("%d:%d", item.ServiceID, item.ID)
				ableItems[key] = item
			}

			for _, unitM := range unitModifiers {
				ableUnitModifiers[unitM.UnitID] = unitM
			}

			for _, priceM := range priceModifiers {
				ablePriceModifiers[priceM.ID] = priceM
			}

			var calculatedServices []orders.CalculateOrderResponseService

			for _, service := range param.Services {
				processed := u.calculateSingleService(orders.CalculateSingleServiceParam{
					OrderedServices:    service,
					AbleItems:          ableItems,
					AbleUnitModifiers:  ableUnitModifiers,
					AblePriceModifiers: ablePriceModifiers,
				})

				calculatedServices = append(calculatedServices, processed)
			}

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

func (u *OrdersUsecase) calculateSingleService(param orders.CalculateSingleServiceParam) orders.CalculateOrderResponseService {
	var serviceID int

	if param.OrderedServices.SubServiceID.Valid {
		serviceID = param.OrderedServices.SubServiceID.GetInt()
	} else {
		serviceID = param.OrderedServices.ServiceID
	}

	chosenItems := []orders.ServiceCommonResponseItem{}

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

	priceModifier, exists := param.AblePriceModifiers[param.OrderedServices.ItemsTypeID]

	if exists {
		commonModifiers = append(commonModifiers, pricemodifiers.PriceModifierCommonData{
			Percent:     priceModifier.Percent,
			Description: priceModifier.Description.String,
			Modifier:    priceModifier.Modifier,
			ModifierID:  priceModifier.ModifierID,
		})
	}

	var unitPriceModifierID int
	unitPriceModifier := param.AbleUnitModifiers[param.OrderedServices.UnitID]

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

	if reduced.TotalUnitQuantity > unitPriceModifier.UnitQuantity {
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
