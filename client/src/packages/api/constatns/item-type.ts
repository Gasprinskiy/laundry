import { ItemTypeId } from '../types/items';

export const ItemTypeTitle: { [key in ItemTypeId]: string } = {
  [ItemTypeId.ItemTypeAdult]: 'Взрослая',
  [ItemTypeId.ItemTypeChildren]: 'Детская',
} as const;
