package postgres

import (
	fulfillmenttypes "laundry/internal/entity/fulfillment-types"
	"laundry/internal/repository"

	"github.com/jmoiron/sqlx"
)

type fulfillmentTypesRepository struct{}

func NewFulfillmentTypesRepository() repository.FulfillmentTypes {
	return &fulfillmentTypesRepository{}
}

func (r *fulfillmentTypesRepository) FindAllFulfillmentTypes(tx *sqlx.Tx) (data []fulfillmenttypes.FulfillmentType, err error) {
	sqlQuery := `
	SELECT 
		ft.id,
		ft.name,
		ft.modifier_id
	FROM public.fulfillment_types ft
	`

	err = tx.Select(&data, sqlQuery)

	return
}
