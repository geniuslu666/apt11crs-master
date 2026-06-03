import { http } from '@/utils/http/axios';

// 获取活动表列表
export function List(params) {
  return http.request({
    url: '/employeeActivity/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除活动表
export function Delete(params) {
  return http.request({
    url: '/employeeActivity/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑活动表
export function Edit(params) {
  return http.request({
    url: '/employeeActivity/edit',
    method: 'POST',
    params,
  });
}

// 修改活动表状态
export function Status(params) {
  return http.request({
    url: '/employeeActivity/status',
    method: 'POST',
    params,
  });
}

// 获取活动表指定详情
export function View(params) {
  return http.request({
    url: '/employeeActivity/view',
    method: 'GET',
    params,
  });
}

// 批量发放预约券
export function BatchIssueCoupons(params) {
  return http.request({
    url: '/employeeActivity/batchIssueCoupons',
    method: 'POST',
    params,
  });
}

// 获取券领取记录列表
export function CouponRecordList(params) {
  return http.request({
    url: '/employeeActivity/couponRecordList',
    method: 'GET',
    params,
  });
}
