import { http } from '@/utils/http/axios';

// 获取优惠券列表
export function List(params) {
  return http.request({
    url: '/pmsCouponType/list',
    method: 'get',
    params,
  });
}

// 获取所有物业列表
export function All(params) {
  return http.request({
    url: '/pmsCouponType/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除优惠券
export function Delete(params) {
  return http.request({
    url: '/pmsCouponType/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑优惠券
export function Edit(params) {
  return http.request({
    url: '/pmsCouponType/edit',
    method: 'POST',
    params,
  });
}

// 修改优惠券状态
export function Status(params) {
  return http.request({
    url: '/pmsCouponType/status',
    method: 'POST',
    params,
  });
}

// 获取优惠券指定详情
export function View(params) {
  return http.request({
    url: '/pmsCouponType/view',
    method: 'GET',
    params,
  });
}

// 获取优惠券最大排序
export function MaxSort() {
  return http.request({
    url: '/pmsCouponType/maxSort',
    method: 'GET',
  });
}

// 发放会员优惠券
export function SendMemberCoupon(params) {
  return http.request({
    url: '/pmsCouponType/sendCoupon',
    method: 'POST',
    params,
  });
}


