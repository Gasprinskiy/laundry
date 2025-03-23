import { UnitType } from '../types/unit-type';

export const UnitTypeTitle: { [key in UnitType]: string } = {
  [UnitType.UnitKG]: 'Вес',
  [UnitType.UnitPCS]: 'Количество',
};

export const UnitTypePointer: { [key in UnitType]: string } = {
  [UnitType.UnitKG]: 'кг.',
  [UnitType.UnitPCS]: 'шт.',
};
