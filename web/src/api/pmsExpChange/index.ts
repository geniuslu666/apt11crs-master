import { http } from '@/utils/http/axios';

// 获取成长值列表
export function List(params) {
  return http.request({
    url: '/pmsExpChange/list',
    method: 'get',
    params,
  });
}


