import { http } from '@/utils/http/axios';

//获取概况
export function dashboard(params) {
  return http.request({
    url: '/car/dashboard',
    method: 'post',
    params,
  });
}
