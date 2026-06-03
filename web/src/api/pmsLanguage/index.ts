import { http, jumpExport } from '@/utils/http/axios';

// 获取语言字典列表
export function List(params) {
  return http.request({
    url: '/pmsLanguage/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除语言字典
export function Delete(params) {
  return http.request({
    url: '/pmsLanguage/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑语言字典
export function Edit(params) {
  return http.request({
    url: '/pmsLanguage/edit',
    method: 'POST',
    params,
  });
}

// 获取语言字典指定详情
export function View(params) {
  return http.request({
    url: '/pmsLanguage/view',
    method: 'GET',
    params,
  });
}

// 导出语言字典
export function Export(params) {
  jumpExport('/pmsLanguage/export', params);
}