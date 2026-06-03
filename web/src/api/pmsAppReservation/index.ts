import { http, jumpExport } from '@/utils/http/axios';

//https://crsdev.yeebok.cn/admin/pmsAppReservation/syncAllOrder
//同步所有订单
export function syncAllOrder(params) {
  return http.request({
    url: '/pmsAppReservation/syncAllOrder',
    method: 'POST',
    params,
  });
}
//同步订单
//https://crsdev.yeebok.cn/admin/pmsAppReservation/syncOrder
export function syncOrder(params) {
  return http.request({
    url: '/pmsAppReservation/syncOrder',
    method: 'POST',
    params,
  });
}
// 获取系统入住订单列表
export function List(params) {
  return http.request({
    url: '/pmsAppReservation/list',
    method: 'get',
    params,
  });
}

// 获取系统入住订单列表
export function RoomList(params) {
  return http.request({
    url: '/pmsAppReservation/roomList',
    method: 'get',
    params,
  });
}

// 获取系统入住订单列表
export function rebateList(params) {
  return http.request({
    url: '/pmsAppReservation/rebateList',
    method: 'get',
    params,
  });
}

// 删除/批量删除系统入住订单
export function Delete(params) {
  return http.request({
    url: '/pmsAppReservation/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑系统入住订单
export function Edit(params) {
  return http.request({
    url: '/pmsAppReservation/edit',
    method: 'POST',
    params,
  });
}

// 修改系统入住订单状态
export function Status(params) {
  return http.request({
    url: '/pmsAppReservation/status',
    method: 'POST',
    params,
  });
}

// 获取系统入住订单指定详情
export function View(params) {
  return http.request({
    url: '/pmsAppReservation/view',
    method: 'GET',
    params,
  });
}

// 导出系统入住订单
export function Export(params) {
  jumpExport('/pmsAppReservation/export', params);
}
// 支付云退款接口 https://crsdev.yeebok.cn/admin/payRefund/paycloud
export function paycloud(params) {
  return http.request({
    url: '/payRefund/paycloud',
    method: 'POST',
    params,
  });
}
// 发送短信  https://crsdev.yeebok.cn/admin/pmsMember/sendSms
export function sendSms(params) {
  return http.request({
    url: '/pmsMember/sendSms',
    method: 'POST',
    params,
  });
}
// 获取系统入住分销订单列表
export function ReferrerList(params) {
  return http.request({
    url: '/pmsAppReservation/referrerList',
    method: 'get',
    params,
  });
}

