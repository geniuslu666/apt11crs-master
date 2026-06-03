import { http, jumpExport } from '@/utils/http/axios';

// 获取房间列表
export function List(params) {
  return http.request({
    url: '/pmsRoomUnit/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除房间
export function Delete(params) {
  return http.request({
    url: '/pmsRoomUnit/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑房间
export function Edit(params) {
  return http.request({
    url: '/pmsRoomUnit/edit',
    method: 'POST',
    params,
  });
}

// 获取房间指定详情
export function View(params) {
  return http.request({
    url: '/pmsRoomUnit/view',
    method: 'GET',
    params,
  });
}

// 导出房间
export function Export(params) {
  jumpExport('/pmsRoomUnit/export', params);
}