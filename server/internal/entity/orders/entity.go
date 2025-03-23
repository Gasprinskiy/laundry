package orders

import (
	fulfillmenttypes "laundry/internal/entity/fulfillment-types"
	pricemodifiers "laundry/internal/entity/price-modifiers"
	"laundry/internal/entity/services"
	"laundry/tools/sqlnull"
	"time"
)

type Order struct {
	ID          int     `json:"id" db:"id"`
	UserName    string  `json:"user_name" db:"user_name"`
	PhoneNumber string  `json:"phone_number" db:"phone_number"`
	Total       float64 `json:"total" db:"total"`
	Final       float64 `json:"final" db:"final"`
}

type GetOrderByDateRangeParam struct {
	StartDate time.Time
	EndDate   time.Time
}

type CalculateOrderItem struct {
	ID       int     `json:"id"`
	Quantity float64 `json:"quantity"`
}

type ServiceCommonResponseItem struct {
	ID          int     `json:"id"`
	ItemID      int     `json:"item_id"`
	ItemName    string  `json:"item_name"`
	Quantity    float64 `json:"quantity"`
	PriceForOne float64 `json:"price_for_one"`
	PriceForAll float64 `json:"price_for_all"`
}

type CalculateOrderService struct {
	ServiceID      int                  `json:"service_id"`
	ServiceName    string               `json:"service_name"`
	SubServiceID   sqlnull.NullInt64    `json:"subservice_id"`
	SubServiceName sqlnull.NullString   `json:"subservice_name"`
	Items          []CalculateOrderItem `json:"items"`
	UnitID         int                  `json:"unit_id"`
	ItemsTypeID    int                  `json:"item_type_id"`
}

type CalculateOrderResponseService struct {
	ServiceID      int                                      `json:"service_id"`
	ServiceName    string                                   `json:"service_name"`
	SubServiceID   sqlnull.NullInt64                        `json:"subservice_id"`
	SubServiceName sqlnull.NullString                       `json:"subservice_name"`
	Total          float64                                  `json:"total"`
	Final          float64                                  `json:"final"`
	Items          []ServiceCommonResponseItem              `json:"items"`
	Discounts      []pricemodifiers.PriceModifierCommonData `json:"discounts"`
	Markups        []pricemodifiers.PriceModifierCommonData `json:"markups"`
	UnitID         int                                      `json:"unit_id"`
	UnitTitle      string                                   `json:"unit_title"`
	UnitModifierID sqlnull.NullInt64                        `json:"unit_modifier_id"`
	ItemsTypeID    int                                      `json:"items_type_id"`
}

type CalculateOrderParam struct {
	Fulfillment fulfillmenttypes.FulfillmentType `json:"fulfillment"`
	Services    []CalculateOrderService          `json:"services"`
}

type CalculateOrderResponse struct {
	TemporaryID   string                                   `json:"temporary_id"`
	OrderServices []CalculateOrderResponseService          `json:"order_services"`
	Fulfillment   fulfillmenttypes.FulfillmentType         `json:"fulfillment"`
	Discounts     []pricemodifiers.PriceModifierCommonData `json:"discounts"`
	Markups       []pricemodifiers.PriceModifierCommonData `json:"markups"`
	Total         float64                                  `json:"total"`
	Final         float64                                  `json:"final"`
}

type CreateOrderParam struct {
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
}

type CreateOrderDbParam struct {
	UserName     string    `db:"user_name"`
	PhoneNumber  string    `db:"phone_number"`
	CreationDate time.Time `db:"creation_date"`
	Total        float64   `db:"total"`
	Final        float64   `db:"final"`
}

type CreateOrderServiceItemRecord struct {
	ServiceItemID  int     `db:"service_item_id"`
	Quantity       float64 `db:"quantity"`
	Price          float64 `db:"price"`
	OrderServiceId int     `db:"order_service_id"`
}

type CreateOrderPriceModifiersRecord struct {
	Modifier    int               `db:"modifier_type_id"`
	Description string            `db:"description"`
	Percent     float64           `db:"percent"`
	OrderID     sqlnull.NullInt64 `db:"order_id"`
	ServiceID   sqlnull.NullInt64 `db:"service_id"`
}

type CreateOrderParamWithPreCalculatedData struct {
	UserParam         CreateOrderDbParam
	PreCalculatedData CalculateOrderResponse
}

type CalculateSingleServiceParam struct {
	OrderedServices       CalculateOrderService
	AbleItems             map[string]services.ServiceItems
	AbleUnitModifiers     map[int]pricemodifiers.UnitPriceModifier
	AbleItemTypeModifiers map[int]pricemodifiers.PriceModifier
}

type CalculateSingleServiceItemReduceResult struct {
	TotalSum          float64
	TotalUnitQuantity float64
}

type CalculateOrderReduceResult struct {
	Total float64
	Final float64
}
