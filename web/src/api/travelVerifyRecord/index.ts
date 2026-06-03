import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({ url: '/travel/verifyRecord/list', method: 'GET', params });
}
