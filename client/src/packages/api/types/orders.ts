import type { FulfillmentType } from './fulfilments';
import type { ItemTypeId } from './items';
import type { PriceModifierCommonData } from './price-modifiers';
import type { UnitType } from './unit-type';

export interface Order {
  id: number;
  user_name: string;
  phone_number: string;
  total: number;
  final: number;
}

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
  subservice_id?: number;
  subservice_name?: string;
  items: CalculateOrderItem[];
  unit_id: number;
  item_type_id: number;
}

export interface CalculateOrderParam {
  fulfillment: FulfillmentType;
  services: CalculateOrderService[];
}

export interface ServiceCommonResponseItem {
  id: number;
  item_id: number;
  item_name: string;
  quantity: number;
  price_for_one: number;
  price_for_all: number;
}

export interface CalculateOrderResponseService {
  service_id: number;
  service_name: string;
  subservice_id: number | null;
  subservice_name: string | null;
  total: number;
  final: number;
  items: ServiceCommonResponseItem[];
  discounts: PriceModifierCommonData[];
  markups: PriceModifierCommonData[];
  unit_id: UnitType;
  unit_title: string;
  unit_modifier_id: number | null;
  items_type_id: ItemTypeId;
}

export interface CalculateOrderResponse {
  temporary_id: string;
  order_services: CalculateOrderResponseService[];
  fulfillment: FulfillmentType;
  discounts: PriceModifierCommonData[];
  markups: PriceModifierCommonData[];
  total: number;
  final: number;
}

export interface CreateOrderParam {
  user_name: string;
  phone_number: string;
}

export interface CreateResponse {
  id: number;
}
