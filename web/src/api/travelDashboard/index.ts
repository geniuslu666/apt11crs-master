import { http } from '@/utils/http/axios';

// 获取一日游概况
export function dashboard(params) {
  return http.request({ url: '/travel/dashboard', method: 'GET', params });
}
