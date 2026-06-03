import { http, jumpExport } from '@/utils/http/axios';

// 获取物业房型列表
export function List(params) {
  return http.request({
    url: '/room/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除物业房型
export function Delete(params) {
  return http.request({
    url: '/room/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑物业房型
export function Edit(params) {
  return http.request({
    url: '/room/edit',
    method: 'POST',
    params,
  });
}

// 获取物业房型指定详情
export function View(params) {
  return http.request({
    url: '/room/view',
    method: 'GET',
    params,
  });
}

// 导出物业房型
export function Export(params) {
  jumpExport('/room/export', params);
}