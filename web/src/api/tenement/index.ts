import { http, jumpExport } from '@/utils/http/axios';

// 获取物业列表
export function List(params) {
  return http.request({
    url: '/tenement/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除物业
export function Delete(params) {
  return http.request({
    url: '/tenement/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑物业
export function Edit(params) {
  return http.request({
    url: '/tenement/edit',
    method: 'POST',
    params,
  });
}

// 修改物业状态
export function Status(params) {
  return http.request({
    url: '/tenement/status',
    method: 'POST',
    params,
  });
}

// 获取物业指定详情
export function View(params) {
  return http.request({
    url: '/tenement/view',
    method: 'GET',
    params,
  });
}

// 导出物业
export function Export(params) {
  jumpExport('/tenement/export', params);
}