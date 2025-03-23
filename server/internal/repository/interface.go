package repository

import (
	fulfillmenttypes "laundry/internal/entity/fulfillment-types"
	"laundry/internal/entity/items"
	"laundry/internal/entity/orders"
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/entity/services"

	"github.com/jmoiron/sqlx"
)

type Services interface {
	FindAllServices(tx *sqlx.Tx) (data []services.Service, err error)
	FindServiceItemsByID(tx *sqlx.Tx, id int) (data []services.ServiceItems, err error)
	FindServiceSubServiceItemsByID(tx *sqlx.Tx, id int) (data []services.ServiceItems, err error)
	FindServiceSubServiceById(tx *sqlx.Tx, id int) (data []services.SubService, err error)
	FindAllServiceItems(tx *sqlx.Tx) (data []services.ServiceItems, err error)
}

type FulfillmentTypes interface {
	FindAllFulfillmentTypes(tx *sqlx.Tx) (data []fulfillmenttypes.FulfillmentType, err error)
}

type PriceModifiers interface {
	FindFulfillmentModifierByID(tx *sqlx.Tx, id int) (data pricemodifiers.PriceModifier, err error)
	FindAllItemTypeModifiers(tx *sqlx.Tx) (data []pricemodifiers.PriceModifier, err error)
	FindAllUnitModifiers(tx *sqlx.Tx) (data []pricemodifiers.UnitPriceModifier, err error)
}

type Items interface {
	FindAllItemTypes(tx *sqlx.Tx) (data []items.ItemTypes, err error)
}

type Orders interface {
	CreateOrder(tx *sqlx.Tx, param orders.CreateOrderDbParam) (id int, err error)
	CreateOrderPriceModifiersRecord(tx *sqlx.Tx, param orders.CreateOrderPriceModifiersRecord) error
	CreateOrderServiceRecord(tx *sqlx.Tx, orderID int, serviceID int) (id int, err error)
	CreateOrderServiceItemRecord(tx *sqlx.Tx, param orders.CreateOrderServiceItemRecord) error
}
