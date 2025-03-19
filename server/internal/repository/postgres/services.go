package postgres

import (
	"laundry/internal/entity/services"
	"laundry/internal/repository"

	"github.com/jmoiron/sqlx"
)

type servicesRepository struct{}

func NewServicesRepository() repository.Services {
	return &servicesRepository{}
}

func (r *servicesRepository) FindAllServices(tx *sqlx.Tx) (data []services.Service, err error) {
	sqlQuery := `
	SELECT DISTINCT ON(s.id)
		s.id, 
		s.name, 
		s.unit_id,
		ss.id IS NOT NULL AS has_sub_service
	FROM public.services s
		LEFT OUTER JOIN public.sub_services ss ON ss.service_id = s.id`

	err = tx.Select(&data, sqlQuery)

	return
}

func (r *servicesRepository) FindServiceItemsByID(tx *sqlx.Tx, id int) (data []services.ServiceItems, err error) {
	sqlQuery := `
	SELECT DISTINCT ON(it.id)
		si.id,
		it.id as item_id,
		it.name as item_name,
		si.price
	FROM public.service_items si
		JOIN public.services s ON s.id = si.service_id
		JOIN public.items it ON it.id = si.item_id
	WHERE si.service_id = $1`

	err = tx.Select(&data, sqlQuery, id)

	return
}

func (r *servicesRepository) FindServiceSubServiceById(tx *sqlx.Tx, id int) (data []services.SubService, err error) {
	sqlQuery := `
	SELECT
		ss.id,
		ss.name
	FROM public.sub_services ss 
		WHERE ss.service_id = $1`

	err = tx.Select(&data, sqlQuery, id)

	return
}

func (r *servicesRepository) FindServiceSubServiceItemsById(tx *sqlx.Tx, id int) (data []services.ServiceItems, err error) {
	sqlQuery := `
	SELECT DISTINCT ON(it.id)
		si.id,
		it.id as item_id,
		it.name as item_name,
		si.price
	FROM public.sub_service_items si
		JOIN public.sub_services ss ON ss.id = si.service_id
		JOIN public.items it ON it.id = si.item_id
	WHERE si.service_id = $1`

	err = tx.Select(&data, sqlQuery, id)

	return
}

func (r *servicesRepository) FindAllServiceItems(tx *sqlx.Tx) (data []services.ServiceItems, err error) {
	sqlQuery := `
	SELECT 
		si.id,
		it.id as item_id,
		it.name as item_name,
		si.price,
		si.service_id
	FROM public.service_items si
		JOIN public.services s ON s.id = si.service_id
		JOIN public.items it ON it.id = si.item_id`

	err = tx.Select(&data, sqlQuery)

	return
}

func (r *servicesRepository) FindAllSubServiceItems(tx *sqlx.Tx) (data []services.ServiceItems, err error) {
	sqlQuery := `
	SELECT 
		si.id,
		it.id as item_id,
		it.name as item_name,
		si.price,
		si.service_id
	FROM public.sub_service_items si
		JOIN public.sub_services ss ON ss.id = si.service_id
		JOIN public.items it ON it.id = si.item_id`

	err = tx.Select(&data, sqlQuery)

	return
}
