package rimport

import (
	"laundry/internal/repository/postgres"
)

func NewRepositoryImports() *RepositoryImport {
	return &RepositoryImport{
		Services:         postgres.NewServicesRepository(),
		FulfillmentTypes: postgres.NewFulfillmentTypesRepository(),
		PriceModifiers:   postgres.NewPriceModifiersRepository(),
		Items:            postgres.NewItemsRepository(),
		Orders:           postgres.NewOrdersRepository(),
	}
}
