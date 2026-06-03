import { http, jumpExport } from '@/utils/http/axios';

// 获取物业列表列表
export function List(params) {
  return http.request({
    url: '/pmsProperty/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除物业列表
export function Delete(params) {
  return http.request({
    url: '/ten/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑物业列表
export function Edit(params) {
  return http.request({
    url: '/ten/edit',
    method: 'POST',
    params,
  });
}

// 修改物业列表状态
export function Status(params) {
  return http.request({
    url: '/ten/status',
    method: 'POST',
    params,
  });
}

// 获取物业列表指定详情
export function View(params) {
  return http.request({
    url: '/ten/view',
    method: 'GET',
    params,
  });
}

// 导出物业列表
export function Export(params) {
  jumpExport('/ten/export', params);
}