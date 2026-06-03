import { http } from '@/utils/http/axios';

// 获取接送机预订单日志列表
export function carOrderLogList(params) {
  return http.request({
    url: '/carOrderLog/list',
    method: 'get',
    params,
  });
}


// 获取餐厅预订单日志列表
export function foodOrderLogList(params) {
  return http.request({
    url: '/foodOrderLog/list',
    method: 'get',
    params,
  });
}

// 获取按摩预订单日志列表
export function spaOrderLogList(params) {
  return http.request({
    url: '/spaOrderLog/list',
    method: 'get',
    params,
  });
}

// 获取酒店预订单日志列表
export function hotelOrderLogList(params) {
  return http.request({
    url: '/hotelOrderLog/list',
    method: 'get',
    params,
  });
}
