import type { CreateOrderParam } from '@/packages/api/types/orders';

export type CreateOrderModalEmits = {
  onCreate: [param: CreateOrderParam];
};
