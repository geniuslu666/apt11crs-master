import { http, jumpExport } from '@/utils/http/axios';

// 获取客户档案列表
export function List(params) {
  return http.request({
    url: '/pmsGuestProfile/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除客户档案
export function Delete(params) {
  return http.request({
    url: '/pmsGuestProfile/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑客户档案
export function Edit(params) {
  return http.request({
    url: '/pmsGuestProfile/edit',
    method: 'POST',
    params,
  });
}

// 获取客户档案指定详情
export function View(params) {
  return http.request({
    url: '/pmsGuestProfile/view',
    method: 'GET',
    params,
  });
}

// 导出客户档案
export function Export(params) {
  jumpExport('/pmsGuestProfile/export', params);
}