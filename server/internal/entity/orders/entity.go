package orders

import (
	fulfillmenttypes "laundry/internal/entity/fulfillment-types"
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/entity/services"
	"laundry/tools/sqlnull"
	"time"
)

type ServiceCommonItem struct {
	ID int `json:"id"`
	// ItemID   int     `json:"item_id"`
	Quantity float64 `json:"quantity"`
}

type ServiceCommonResponseItem struct {
	ID       int     `json:"id"`
	ItemID   int     `json:"item_id"`
	ItemName string  `json:"item_name"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
}

type CreateOrderServiceParam struct {
	ServiceID int `json:"service_id"`
	// ServiceName         string              `json:"service_name"`
	SubServiceID        sqlnull.NullInt64   `json:"subservice_id"`
	Items               []ServiceCommonItem `json:"items"`
	UnitID              int                 `json:"unit_id"`
	ItemsTypeModifierID sqlnull.NullInt64   `json:"modifier_id"`
}

type CreateOrderResponseServiceParam struct {
	ServiceID    int                         `json:"service_id"`
	ServiceName  string                      `json:"service_name"`
	SubServiceID sqlnull.NullInt64           `json:"subservice_id"`
	Items        []ServiceCommonResponseItem `json:"items"`
	UnitID       int                         `json:"unit_id"`
}

type CreateOrderParam struct {
	// UserName        string                           `json:"user_name"`
	// UserPhoneNumber string                           `json:"user_phone_number"`
	// Fulfillment     fulfillmenttypes.FulfillmentType `json:"fulfillment"`
	Services []CreateOrderServiceParam `json:"services"`
}

type OrderPriceModificators struct {
	Title   string `json:"title"`
	Percent string `json:"percent"`
}

type CreateOrderResponse struct {
	OrderID       int                               `json:"order_id"`
	OrderDate     time.Time                         `json:"order_date"`
	OrderServices []CreateOrderResponseServiceParam `json:"order_services"`
	Fulfillment   fulfillmenttypes.FulfillmentType  `json:"fulfillment"`
	Discounts     []OrderPriceModificators          `json:"discounts"`
	Markups       []OrderPriceModificators          `json:"markups"`
	Total         float64                           `json:"total"`
}

type ProcessSingleServiceParam struct {
	OrderedServices    CreateOrderServiceParam
	AbleItems          map[string]services.ServiceItems
	AbleUnitModifiers  map[int]pricemodifiers.UnitPriceModifier
	AblePriceModifiers map[int]pricemodifiers.PriceModifier
}

type ProcessSingleServiceItemReduce struct {
	Item     services.ServiceItems
	Quantity float64
}

type ProcessSingleServiceItemReduceResult struct {
	TotalSub          float64
	TotalUnitQuantity float64
}
