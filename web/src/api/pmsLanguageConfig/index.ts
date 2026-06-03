import { http, jumpExport } from '@/utils/http/axios';

// 获取多语言项列表
export function List(params) {
  return http.request({
    url: '/pmsLanguageConfig/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除多语言项
export function Delete(params) {
  return http.request({
    url: '/pmsLanguageConfig/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑多语言项
export function Edit(params) {
  return http.request({
    url: '/pmsLanguageConfig/edit',
    method: 'POST',
    params,
  });
}

// 获取多语言项指定详情
export function View(params) {
  return http.request({
    url: '/pmsLanguageConfig/view',
    method: 'GET',
    params,
  });
}