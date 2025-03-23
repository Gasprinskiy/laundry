export enum ItemTypeId {
  ItemTypeChildren = 1,
  ItemTypeAdult = 2,
}

export interface ItemType {
  id: ItemTypeId;
  name: string;
  modifier_id: number | null;
}
