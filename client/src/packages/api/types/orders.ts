export interface CalculateOrderItem {
  id: number;
  quantity: number;
}

export interface ServiceCommonResponseItem {
  id: number;
  item_id: number;
  item_name: string;
  quantity: number;
  price_for_one: number;
  price_for_all: number;
}

export interface CalculateOrderService {
  service_id: number;
  service_name: string;
  subservice_id: number | null;
  subservice_name: string | null;
  items: CalculateOrderItem[];
  unit_id: number;
  item_type_id: number;
}
