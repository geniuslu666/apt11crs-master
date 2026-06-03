import { http } from '@/utils/http/axios';

// 获取礼品券列表
export function List(params) {
  return http.request({
    url: '/thCoupon/list',
    method: 'get',
    params,
  });
}

// 获取所有物业列表
export function All(params) {
  return http.request({
    url: '/thCoupon/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除礼品券
export function Delete(params) {
  return http.request({
    url: '/thCoupon/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑礼品券
export function Edit(params) {
  return http.request({
    url: '/thCoupon/edit',
    method: 'POST',
    params,
  });
}

// 修改礼品券状态
export function Status(params) {
  return http.request({
    url: '/thCoupon/status',
    method: 'POST',
    params,
  });
}

// 修改礼品券使用状态
export function UseStatus(params) {
  return http.request({
    url: '/thCoupon/useStatus',
    method: 'POST',
    params,
  });
}

// 获取礼品券指定详情
export function View(params) {
  return http.request({
    url: '/thCoupon/view',
    method: 'GET',
    params,
  });
}

// 获取礼品券最大排序
export function MaxSort() {
  return http.request({
    url: '/thCoupon/maxSort',
    method: 'GET',
  });
}

// 发放会员优惠券
export function SendMemberCoupon(params) {
  return http.request({
    url: '/thCoupon/sendCoupon',
    method: 'POST',
    params,
  });
}

