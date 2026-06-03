import { http, jumpExport } from '@/utils/http/axios';

// 获取地区列表
export function List(params) {
  return http.request({
    url: '/propertyRegion/list',
    method: 'get',
    params,
  });
}


