package orders

import (
	fulfillmenttypes "laundry/internal/entity/fulfillment-types"
	"laundry/tools/sqlnull"
	"time"
)

type ServiceCommonItem struct {
	ID                 int               `json:"id"`
	ItemID             int               `json:"item_id"`
	ItemTypeID         int               `json:"item_type_id"`
	ItemTypeModifierID sqlnull.NullInt64 `json:"modifier_id"`
	Quantity           float64           `json:"quantity"`
}

type ServiceCommonResponseItem struct {
	ID         int     `json:"id"`
	ItemID     int     `json:"item_id"`
	ItemName   string  `json:"item_name"`
	ItemTypeID int     `json:"item_type_id"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
}

type CreateOrderServiceParam[T any] struct {
	ServiceID    int               `json:"service_id"`
	ServiceName  string            `json:"service_name"`
	SubServiceID sqlnull.NullInt64 `json:"subservice_id"`
	Items        []T               `json:"services"`
	UnitID       int               `json:"unit_id"`
	Total        float64           `json:"total"`
}

type CreateOrderParam struct {
	UserName        string                                       `json:"user_name"`
	UserPhoneNumber string                                       `json:"user_phone_number"`
	Fulfillment     fulfillmenttypes.FulfillmentType             `json:"fulfillment"`
	Services        []CreateOrderServiceParam[ServiceCommonItem] `json:"services"`
}

type OrderPriceModificators struct {
	Title   string `json:"title"`
	Percent string `json:"percent"`
}

type CreateOrderResponse struct {
	OrderID       int                                                  `json:"order_id"`
	OrderDate     time.Time                                            `json:"order_date"`
	OrderServices []CreateOrderServiceParam[ServiceCommonResponseItem] `json:"order_services"`
	Fulfillment   fulfillmenttypes.FulfillmentType                     `json:"fulfillment"`
	Discounts     OrderPriceModificators                               `json:"discounts"`
	Markups       OrderPriceModificators                               `json:"markups"`
	Total         float64                                              `json:"total"`
}
