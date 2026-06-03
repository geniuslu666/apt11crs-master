import { http } from '@/utils/http/axios';

// 获取员工管理列表
export function Text(params) {
  return http.request({
    url: '/translate/text',
    method: 'get',
    params,
  });
}
