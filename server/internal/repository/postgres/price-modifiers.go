package postgres

import (
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/repository"

	"github.com/jmoiron/sqlx"
)

type priceModifiersRepository struct{}

func NewPriceModifiersRepository() repository.PriceModifiers {
	return &priceModifiersRepository{}
}

func (r *priceModifiersRepository) FindPriceModifierByID(tx *sqlx.Tx, id int) (data pricemodifiers.PriceModifier, err error) {
	sqlQuery := `
	SELECT 
		pm.id,
		pm.percent,
		pmt.id as modifier
	FROM public.price_modifiers pm
		JOIN public.price_modifiers_types pmt ON pmt.id = pm.modifier_type_id
	WHERE pm.id = $1`

	err = tx.Get(&data, sqlQuery, id)

	return
}

func (r *priceModifiersRepository) FindAllPriceModifiers(tx *sqlx.Tx) (data []pricemodifiers.PriceModifier, err error) {
	sqlQuery := `
	SELECT 
		pm.id,
		pm.percent,
		pmt.id as modifier,
	FROM public.price_modifiers pm
		JOIN public.price_modifiers_types pmt ON pmt.id = pm.modifier_type_id
	`

	err = tx.Select(&data, sqlQuery)

	return
}

func (r *priceModifiersRepository) FindAllItemTypeModifiers(tx *sqlx.Tx) (data []pricemodifiers.PriceModifier, err error) {
	sqlQuery := `
	SELECT 
		it.id,
		pm.id as modifier_id,
		pmt.id as modifier,
		pm.percent,
		it.description
	FROM public.item_types it
		JOIN public.price_modifiers pm ON pm.id = it.modifier_id
		JOIN public.price_modifiers_types pmt ON pmt.id = pm.modifier_type_id
	`

	err = tx.Select(&data, sqlQuery)

	return
}

func (r *priceModifiersRepository) FindAllUnitModifiers(tx *sqlx.Tx) (data []pricemodifiers.UnitPriceModifier, err error) {
	sqlQuery := `
	SELECT 
		um.id,
		um.unit_id,
		um.unit_quantity,
		pm.percent,
		pm.id as modifier_id,
		pmt.id as modifier,
		um.description
	FROM public.unit_modifiers um
		JOIN public.price_modifiers pm ON pm.id = um.modifier_id
		JOIN public.price_modifiers_types pmt ON pmt.id = pm.modifier_type_id`

	err = tx.Select(&data, sqlQuery)

	return
}
