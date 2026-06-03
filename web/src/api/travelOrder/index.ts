import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({ url: '/travel/order/list', method: 'GET', params });
}

export function View(params) {
  return http.request({ url: '/travel/order/view', method: 'GET', params });
}

export function Refund(params) {
  return http.request({ url: '/travel/order/refund', method: 'POST', params });
}
