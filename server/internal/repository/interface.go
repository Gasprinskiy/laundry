package repository

import (
	fulfillmenttypes "laundry/internal/entity/fulfillment-types"
	"laundry/internal/entity/items"
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/entity/services"

	"github.com/jmoiron/sqlx"
)

type Services interface {
	FindAllServices(tx *sqlx.Tx) (data []services.Service, err error)
	FindServiceItemsByID(tx *sqlx.Tx, id int) (data []services.ServiceItems, err error)
	FindServiceSubServiceById(tx *sqlx.Tx, id int) (data []services.SubService, err error)
	FindServiceSubServiceItemsById(tx *sqlx.Tx, id int) (data []services.ServiceItems, err error)
	FindAllServiceItems(tx *sqlx.Tx) (data []services.ServiceItems, err error)
	FindAllSubServiceItems(tx *sqlx.Tx) (data []services.ServiceItems, err error)
}

type FulfillmentTypes interface {
	FindAllFulfillmentTypes(tx *sqlx.Tx) (data []fulfillmenttypes.FulfillmentType, err error)
}

type PriceModifiers interface {
	FindPriceModifierByID(tx *sqlx.Tx, id int) (data pricemodifiers.PriceModifier, err error)
}

type Items interface {
	FindAllItemTypes(tx *sqlx.Tx) (data []items.ItemTypes, err error)
}
