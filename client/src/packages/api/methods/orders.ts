import type { CalculateOrderParam, CalculateOrderResponse, CreateOrderParam, CreateResponse } from '../types/orders';
import $api from '../worker';

export function calculateOrder(param: CalculateOrderParam): Promise<CalculateOrderResponse> {
  return $api('/orders/calculate', {
    method: 'POST',
    body: param,
  });
}

export function createOrder(id: string, param: CreateOrderParam): Promise<CreateResponse> {
  return $api(`/orders/create/${id}`, {
    method: 'POST',
    body: param,
  });
};
