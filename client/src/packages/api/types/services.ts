import type { FulfillmentType } from './fulfilments';
import type { ItemType } from './items';
import type { UnitType } from './unit-type';

export interface Service {
  id: number;
  name: string;
  unit_id: UnitType;
  has_sub_service: boolean;
}

export type ServiceSubService = Omit<Service, 'unit_id' | 'has_sub_service'>;

export interface ServicesCommonResponse {
  services: Service[];
  item_types: ItemType[];
  fulfillment_types: FulfillmentType[];
}

export interface ServiceItem {
  id: number;
  item_id: number;
  item_name: string;
}
