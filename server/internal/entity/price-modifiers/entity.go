package pricemodifiers

type PriceModifier struct {
	ID       int     `db:"id"`
	Percent  float64 `db:"percent"`
	Modifier string  `db:"modifier"`
}
