import { http, jumpExport } from '@/utils/http/axios';


// APP预定_获取最新订单 https://crsdev.yeebok.cn/admin/pmsAppReservation/lastOrder
export function lastOrder(params) {
  return http.request({
    url: '/pmsAppReservation/lastOrder',
    method: 'get',
    params,
  });
}

// 获取预约订单列表
export function List(params) {
  return http.request({
    url: '/pmsRoomReservation/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除预约订单
export function Delete(params) {
  return http.request({
    url: '/pmsRoomReservation/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑预约订单
export function Edit(params) {
  return http.request({
    url: '/pmsRoomReservation/edit',
    method: 'POST',
    params,
  });
}

// 修改预约订单状态
export function Status(params) {
  return http.request({
    url: '/pmsRoomReservation/status',
    method: 'POST',
    params,
  });
}

// 获取预约订单指定详情
export function pmsRoomReservationView(params) {
  return http.request({
    url: '/pmsRoomReservation/view',
    method: 'GET',
    params,
  });
}

// 导出预约订单
export function Export(params) {
  jumpExport('/pmsRoomReservation/export', params);
}

// 订单取消预约
export function Cancel(params) {
  return http.request({
    url: '/pmsAppReservation/cancel',
    method: 'POST',
    params,
  });
}


// 获取预约订单简易详情（所属物业、下单时间）
export function singleView(params) {
  return http.request({
    url: '/pmsAppReservation/singleView',
    method: 'GET',
    params,
  });
}
