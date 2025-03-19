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
		pmt.modifier
	FROM public.price_modifiers pm
		JOIN public.price_modifiers_types pmt ON pmt.id = pm.modifier_type_id
	WHERE pm.id = $1`

	err = tx.Get(&data, sqlQuery, id)

	return
}
