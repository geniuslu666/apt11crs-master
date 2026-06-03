import { http } from '@/utils/http/axios';

// 获取接送机地点类型列表
export function List(params) {
  return http.request({
    url: '/carAddressType/list',
    method: 'get',
    params,
  });
}

