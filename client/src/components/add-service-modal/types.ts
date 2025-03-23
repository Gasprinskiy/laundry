import type { ItemType } from '@/packages/api/types/items';
import type { Service, ServiceItem, ServiceSubService } from '@/packages/api/types/services';

export interface AddServiceModalProps {
  services: Service[];
  itemType: ItemType[];
}

export interface AddedItem {
  item: ServiceItem;
  quantity: number;
}

export interface AddServiceParams {
  chosenService: Service;
  chosenSubService?: ServiceSubService;
  addedItems: AddedItem[];
  itemType: ItemType;
}

export type AddServiceModalEmits = {
  onAdd: [item: AddServiceParams];
};
