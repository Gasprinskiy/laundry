package usecase

import (
	"fmt"
	"laundry/internal/entity/orders"
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/entity/services"
	"laundry/internal/repository/rimport"
	"laundry/tools/slice"
	transactiongeneric "laundry/tools/transaction-generic"

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

type Shit struct {
	Total     float64
	Final     float64
	Discounts []pricemodifiers.PriceModifierCommonData
	Markups   []pricemodifiers.PriceModifierCommonData
}

func (u *OrdersUsecase) ProcessOrder(param orders.CreateOrderParam) (result []Shit, err error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) ([]Shit, error) {
			serviceItems, err := u.repo.Services.FindAllServiceItems(tx)
			if err != nil {
				return result, err
			}

			subServiceItems, err := u.repo.Services.FindAllSubServiceItems(tx)
			if err != nil {
				return result, err
			}

			unitModifiers, err := u.repo.PriceModifiers.FindAllUnitModifiers(tx)
			if err != nil {
				return result, err
			}

			priceModifiers, err := u.repo.PriceModifiers.FindAllItemTypeModifiers(tx)
			if err != nil {
				return result, err
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
				ablePriceModifiers[priceM.ModifierID] = priceM
			}

			for _, service := range param.Services {
				processed := u.processSingleService(orders.ProcessSingleServiceParam{
					OrderedServices:    service,
					AbleItems:          ableItems,
					AbleUnitModifiers:  ableUnitModifiers,
					AblePriceModifiers: ablePriceModifiers,
				})

				result = append(result, processed)
			}

			return result, nil
		},
		"Не удалось обработать заказ",
	)

}

func (u *OrdersUsecase) processSingleService(param orders.ProcessSingleServiceParam) Shit {
	var serviceID int

	if param.OrderedServices.SubServiceID.Valid {
		serviceID = param.OrderedServices.SubServiceID.GetInt()
	} else {
		serviceID = param.OrderedServices.ServiceID
	}

	chosenServices := []orders.ProcessSingleServiceItemReduce{}

	for _, chosenItem := range param.OrderedServices.Items {
		key := fmt.Sprintf("%d:%d", serviceID, chosenItem.ID)
		chosenServices = append(chosenServices, orders.ProcessSingleServiceItemReduce{
			Item:     param.AbleItems[key],
			Quantity: chosenItem.Quantity,
		})
	}

	var commonModifiers []pricemodifiers.PriceModifierCommonData

	if param.OrderedServices.ItemsTypeModifierID.Valid {
		modifierId := param.OrderedServices.ItemsTypeModifierID.GetInt()
		priceModifier := param.AblePriceModifiers[modifierId]

		commonModifiers = append(commonModifiers, pricemodifiers.PriceModifierCommonData{
			Percent:     priceModifier.Percent,
			Description: priceModifier.Description.String,
			Modifier:    priceModifier.Modifier,
		})
	}

	unitPriceModifier := param.AbleUnitModifiers[param.OrderedServices.UnitID]

	reduced := slice.Reduce(
		chosenServices,
		func(
			acc orders.ProcessSingleServiceItemReduceResult,
			value orders.ProcessSingleServiceItemReduce,
			index int,
		) orders.ProcessSingleServiceItemReduceResult {
			acc.TotalSub = acc.TotalSub + (value.Item.Price * value.Quantity)
			acc.TotalUnitQuantity = acc.TotalUnitQuantity + value.Quantity
			return acc
		},
		orders.ProcessSingleServiceItemReduceResult{
			TotalSub:          0,
			TotalUnitQuantity: 0,
		},
	)

	if reduced.TotalUnitQuantity > unitPriceModifier.UnitQuantity {
		commonModifiers = append(commonModifiers, pricemodifiers.PriceModifierCommonData{
			Percent:     unitPriceModifier.Percent,
			Description: unitPriceModifier.Description.String,
			Modifier:    unitPriceModifier.Modifier,
		})
	}

	var discounts []pricemodifiers.PriceModifierCommonData
	var markups []pricemodifiers.PriceModifierCommonData

	final := reduced.TotalSub

	fmt.Println("commonModifiers: ", commonModifiers)

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

	return Shit{
		Total:     reduced.TotalSub,
		Final:     final,
		Discounts: discounts,
		Markups:   markups,
	}
}

func (u *OrdersUsecase) countMarkupsAndDiscounts(
	modifierID int,
	percent float64,
	sum float64,
) (result float64, isDiscount bool) {

	if modifierID == pricemodifiers.ModifierDiscount {
		isDiscount = true
	}

	switch modifierID {
	case pricemodifiers.ModifierDiscount:
		isDiscount = true
		result = sum - (sum * percent / 100)

	case pricemodifiers.ModifierMarkup:
		result = sum + (sum * percent / 100)
	}

	return
}
