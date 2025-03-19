package services

import (
	"laundry/internal/entity/items"
)

type SubService struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Service struct {
	ID            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	UnitID        int    `json:"unit_id" db:"unit_id"`
	HasSubService bool   `json:"has_sub_service" db:"has_sub_service"`
}

type ServiceItems struct {
	ID        int     `json:"id" db:"id"`
	ItemID    int     `json:"item_id" db:"item_id"`
	ItemName  string  `json:"item_name" db:"item_name"`
	Price     float64 `json:"-" db:"price"`
	ServiceID int     `json:"-" db:"service_id"`
}

type ServicesCommonResponse struct {
	Services  []Service         `json:"services"`
	ItemTypes []items.ItemTypes `json:"item_types"`
}
