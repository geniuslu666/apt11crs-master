import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({ url: '/travel/productSku/list', method: 'GET', params });
}

export function View(params) {
  return http.request({ url: '/travel/productSku/view', method: 'GET', params });
}

export function Edit(params) {
  return http.request({ url: '/travel/productSku/edit', method: 'POST', params });
}

export function Delete(params) {
  return http.request({ url: '/travel/productSku/delete', method: 'POST', params });
}

export function Status(params) {
  return http.request({ url: '/travel/productSku/status', method: 'POST', params });
}
