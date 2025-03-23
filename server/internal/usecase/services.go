package usecase

import (
	"laundry/internal/entity/services"
	"laundry/internal/repository/rimport"
	transactiongeneric "laundry/tools/transaction-generic"

	"github.com/jmoiron/sqlx"
)

type ServicesUsecase struct {
	repo *rimport.RepositoryImport
	db   *sqlx.DB
}

func NewServicesUsecase(
	repo *rimport.RepositoryImport,
	db *sqlx.DB,
) *ServicesUsecase {
	return &ServicesUsecase{
		repo,
		db,
	}
}

// FindAllServices собирает все доступные услуги и типы одежды
func (u *ServicesUsecase) FindAllServices() (services.ServicesCommonResponse, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) (data services.ServicesCommonResponse, err error) {

			servicesList, err := u.repo.Services.FindAllServices(tx)
			if err != nil {
				return
			}

			itemTypes, err := u.repo.Items.FindAllItemTypes(tx)
			if err != nil {
				return
			}

			fulfillmentTypes, err := u.repo.FulfillmentTypes.FindAllFulfillmentTypes(tx)
			if err != nil {
				return
			}

			data = services.ServicesCommonResponse{
				Services:         servicesList,
				ItemTypes:        itemTypes,
				FulfillmentTypes: fulfillmentTypes,
			}

			return
		},
		"Ошибка при получении услуг",
	)
}

// FindServiceItemsByID поиск списка доступной одежды в услуге
func (u *ServicesUsecase) FindServiceItemsByID(id int, isSub bool) ([]services.ServiceItems, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) ([]services.ServiceItems, error) {
			if isSub {
				return u.repo.Services.FindServiceSubServiceItemsByID(tx, id)
			}
			return u.repo.Services.FindServiceItemsByID(tx, id)
		},
		"Ошибка при получении списка доступной одежды в услуге",
	)
}

// FindServiceSubServiceById поиск списка доп. услуг у основной услуги
func (u *ServicesUsecase) FindServiceSubServiceById(id int) ([]services.SubService, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *sqlx.Tx) ([]services.SubService, error) {
			return u.repo.Services.FindServiceSubServiceById(tx, id)
		},
		"Ошибка при получении доп. услуг у основной услуги",
	)
}
