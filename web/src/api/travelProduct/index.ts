import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({ url: '/travel/product/list', method: 'GET', params });
}

export function View(params) {
  return http.request({ url: '/travel/product/view', method: 'GET', params });
}

export function Edit(params) {
  return http.request({ url: '/travel/product/edit', method: 'POST', params });
}

export function Delete(params) {
  return http.request({ url: '/travel/product/delete', method: 'POST', params });
}

export function Status(params) {
  return http.request({ url: '/travel/product/status', method: 'POST', params });
}

export function RecycleList(params) {
  return http.request({ url: '/travel/product/recycle', method: 'GET', params });
}

export function Restore(params) {
  return http.request({ url: '/travel/product/restore', method: 'POST', params });
}
