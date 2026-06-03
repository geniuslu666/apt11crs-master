import { http, jumpExport } from '@/utils/http/axios';

// 获取接送机地点管理列表
export function List(params) {
  return http.request({
    url: '/CarAddress/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除接送机地点管理
export function Delete(params) {
  return http.request({
    url: '/CarAddress/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑接送机地点管理
export function Edit(params) {
  return http.request({
    url: '/CarAddress/edit',
    method: 'POST',
    params,
  });
}

// 修改接送机地点管理状态
export function Status(params) {
  return http.request({
    url: '/CarAddress/status',
    method: 'POST',
    params,
  });
}

// 获取接送机地点管理指定详情
export function View(params) {
  return http.request({
    url: '/CarAddress/view',
    method: 'GET',
    params,
  });
}

// 导出接送机地点管理
export function Export(params) {
  jumpExport('/CarAddress/export', params);
}


