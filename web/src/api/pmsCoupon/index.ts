import { http } from '@/utils/http/axios';

// 获取优惠券列表
export function List(params) {
  return http.request({
    url: '/pmsCoupon/list',
    method: 'get',
    params,
  });
}


// 会员优惠券概况
export function Stat(params) {
  return http.request({
    url: '/pmsCoupon/stat',
    method: 'get',
    params,
  });
}

// 回收优惠券
export function Recycle(params) {
  return http.request({
    url: '/pmsCoupon/recycle',
    method: 'POST',
    params,
  });
}
