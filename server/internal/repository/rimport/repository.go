package rimport

import "laundry/internal/repository"

type RepositoryImport struct {
	Services         repository.Services
	FulfillmentTypes repository.FulfillmentTypes
	PriceModifiers   repository.PriceModifiers
	Items            repository.Items
}
