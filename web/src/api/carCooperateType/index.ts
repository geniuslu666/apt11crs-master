import { http, jumpExport } from '@/utils/http/axios';

// 获取接送机营业类型列表
export function List(params) {
  return http.request({
    url: '/carCooperateType/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除接送机营业类型
export function Delete(params) {
  return http.request({
    url: '/carCooperateType/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑接送机营业类型
export function Edit(params) {
  return http.request({
    url: '/carCooperateType/edit',
    method: 'POST',
    params,
  });
}

// 修改接送机营业类型状态
export function Status(params) {
  return http.request({
    url: '/carCooperateType/status',
    method: 'POST',
    params,
  });
}

// 获取接送机营业类型指定详情
export function View(params) {
  return http.request({
    url: '/carCooperateType/view',
    method: 'GET',
    params,
  });
}

// 导出接送机营业类型
export function Export(params) {
  jumpExport('/carCooperateType/export', params);
}


