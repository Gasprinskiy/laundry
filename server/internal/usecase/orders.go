package usecase

import (
	"laundry/internal/entity/orders"
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

func (u *OrdersUsecase) ProcessOrder(param orders.CreateOrderParam) (result orders.CreateOrderResponse, err error) {

	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) (orders.CreateOrderResponse, error) {
			serviceItems, err := u.repo.Services.FindAllServiceItems(tx)
			if err != nil {
				return result, err
			}

			subServiceItems, err := u.repo.Services.FindAllSubServiceItems(tx)
			if err != nil {
				return result, err
			}

			sItems := append(serviceItems, subServiceItems...)

			serviceMap := make(map[int][]services.ServiceItems)

			for _, srv := range param.Services {
				serviceMap[srv.ServiceID] = slice.Filter(
					sItems,
					func(item services.ServiceItems, index int) bool {
						return srv.ServiceID == item.ServiceID
					},
				)
			}

			return result, nil
		},
		"Не удалось обработать заказ",
	)

}
