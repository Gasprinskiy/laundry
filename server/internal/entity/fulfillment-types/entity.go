package fulfillmenttypes

import "laundry/tools/sqlnull"

type FulfillmentType struct {
	ID         int               `json:"id" db:"id"`
	Name       string            `json:"name" db:"name"`
	ModifierID sqlnull.NullInt64 `json:"modifier_id" db:"modifier_id"`
}
