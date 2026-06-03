import { http, jumpExport } from '@/utils/http/axios';

// 获取房型列表
export function List(params) {
  return http.request({
    url: '/pmsRoomType/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除房型
export function Delete(params) {
  return http.request({
    url: '/pmsRoomType/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑房型
export function Edit(params) {
  return http.request({
    url: '/pmsRoomType/edit',
    method: 'POST',
    params,
  });
}

// 获取房型指定详情
export function View(params) {
  return http.request({
    url: '/pmsRoomType/view',
    method: 'GET',
    params,
  });
}

// 更新房型状态
export function Status(params) {
  return http.request({
    url: '/pmsRoomType/status',
    method: 'POST',
    params,
  });
}
