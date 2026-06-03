import { http } from '@/utils/http/axios';

// 获取渠道管理列表
export function List(params) {
  return http.request({
    url: '/pmsBalanceChange/list',
    method: 'get',
    params,
  });
}

// 积分概况
export function Stat(params) {
  return http.request({
    url: '/pmsBalanceChange/stat',
    method: 'get',
    params,
  });
}


