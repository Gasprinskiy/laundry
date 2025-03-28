import type { CalculateOrderParam, CalculateOrderResponse, CreateOrderParam, CreateResponse, Order } from '../types/orders';
import $api from '../client';

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

export function getTodayOrders(): Promise<Order[]> {
  return $api('/orders/today', {
    method: 'GET',
  });
};
