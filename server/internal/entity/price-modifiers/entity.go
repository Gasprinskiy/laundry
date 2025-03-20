package pricemodifiers

import "laundry/tools/sqlnull"

type PriceModifier struct {
	ID          int                `db:"id"`
	Percent     float64            `db:"percent"`
	ModifierID  int                `db:"modifier_id"`
	Modifier    int                `db:"modifier"`
	Description sqlnull.NullString `db:"description"`
}

type UnitPriceModifier struct {
	ID           int                `db:"id"`
	UnitID       int                `db:"unit_id"`
	UnitQuantity float64            `db:"unit_quantity"`
	Percent      float64            `db:"percent"`
	ModifierID   int                `db:"modifier_id"`
	Modifier     int                `db:"modifier"`
	Description  sqlnull.NullString `db:"description"`
}

type PriceModifierCommonData struct {
	Percent     float64 `json:"percent"`
	Description string  `json:"description"`
	Modifier    int     `json:"modifier"`
}
