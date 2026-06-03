import { http, jumpExport } from '@/utils/http/axios';

// 获取预约订单列表
export function List(params) {
  return http.request({
    url: '/hgexample/pmsRoomReservation/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除预约订单
export function Delete(params) {
  return http.request({
    url: '/hgexample/pmsRoomReservation/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑预约订单
export function Edit(params) {
  return http.request({
    url: '/hgexample/pmsRoomReservation/edit',
    method: 'POST',
    params,
  });
}

// 修改预约订单状态
export function Status(params) {
  return http.request({
    url: '/hgexample/pmsRoomReservation/status',
    method: 'POST',
    params,
  });
}

// 获取预约订单指定详情
export function View(params) {
  return http.request({
    url: '/hgexample/pmsRoomReservation/view',
    method: 'GET',
    params,
  });
}

// 导出预约订单
export function Export(params) {
  jumpExport('/hgexample/pmsRoomReservation/export', params);
}