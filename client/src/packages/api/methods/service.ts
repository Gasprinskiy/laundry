import type { ServiceItem, ServicesCommonResponse, ServiceSubService } from '../types/services';
import $api from '../worker';

export function fetchAllAbleServices(): Promise<ServicesCommonResponse> {
  return $api('/services', {
    method: 'GET',
  });
}

export function fetchServiceItems(id: number): Promise<ServiceItem[]> {
  return $api(`/services/${id}/items`, {
    method: 'GET',
  });
}

export function fetchServiceSubService(id: number): Promise<ServiceSubService[]> {
  return $api(`/services/${id}/sub`, {
    method: 'GET',
  });
}

export function fetchSubServiceItems(id: number): Promise<ServiceItem[]> {
  return $api(`/services/sub/${id}/items`, {
    method: 'GET',
  });
}
