package postgres

import (
	"laundry/internal/entity/items"
	"laundry/internal/repository"

	"github.com/jmoiron/sqlx"
)

type itemsRepository struct{}

func NewItemsRepository() repository.Items {
	return &itemsRepository{}
}

func (r *itemsRepository) FindAllItemTypes(tx *sqlx.Tx) (data []items.ItemTypes, err error) {
	sqlQuery := `
	SELECT 
		it.id,
		it.name,
		it.modifier_id
	FROM public.item_types it`

	err = tx.Select(&data, sqlQuery)

	return
}
