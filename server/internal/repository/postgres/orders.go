package postgres

import (
	"laundry/internal/entity/orders"
	"laundry/internal/repository"

	"github.com/jmoiron/sqlx"
)

type ordersRepository struct{}

func NewOrdersRepository() repository.Orders {
	return &ordersRepository{}
}

func (r *ordersRepository) CreateOrder(tx *sqlx.Tx, param orders.CreateOrderDbParam) (id int, err error) {
	sqlQuery := `
	INSERT INTO 
	public.orders (user_name, phone_number, creation_date, total, final)
	VALUES (:user_name, :phone_number, :creation_date, :total, :final)
	RETURNING id`

	stmt, err := tx.PrepareNamed(sqlQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(param).Scan(&id)
	if err != nil {
		return
	}

	return
}

func (r *ordersRepository) CreateOrderPriceModifiersRecord(tx *sqlx.Tx, param orders.CreateOrderPriceModifiersRecord) error {
	sqlQuery := `
	INSERT INTO
	public.order_price_modifiers (modifier_type_id, description, percent, order_id, service_id)
	VALUES (:modifier_type_id, :description, :percent, :order_id, :service_id)`

	_, err := tx.NamedExec(sqlQuery, param)

	return err
}

func (r *ordersRepository) CreateOrderServiceRecord(tx *sqlx.Tx, orderID int, serviceID int) (id int, err error) {
	sqlQuery := `
	INSERT INTO 
	public.orders_services (order_id, service_id)
	VALUES ($1, $2)
	RETURNING id`

	err = tx.QueryRow(sqlQuery, orderID, serviceID).Scan(&id)
	if err != nil {
		return
	}

	return
}

func (r *ordersRepository) CreateOrderServiceItemRecord(tx *sqlx.Tx, param orders.CreateOrderServiceItemRecord) error {
	sqlQuery := `
	INSERT INTO
	public.orders_service_items (service_item_id, quantity, price, order_service_id)
	VALUES (:service_item_id, :quantity, :price, :order_service_id)`

	_, err := tx.NamedExec(sqlQuery, param)

	return err
}
