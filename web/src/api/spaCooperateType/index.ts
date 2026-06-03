import { http, jumpExport } from '@/utils/http/axios';

// 获取按摩营业类型列表
export function List(params) {
  return http.request({
    url: '/spaCooperateType/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除按摩营业类型
export function Delete(params) {
  return http.request({
    url: '/spaCooperateType/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑按摩营业类型
export function Edit(params) {
  return http.request({
    url: '/spaCooperateType/edit',
    method: 'POST',
    params,
  });
}

// 修改按摩营业类型状态
export function Status(params) {
  return http.request({
    url: '/spaCooperateType/status',
    method: 'POST',
    params,
  });
}

// 获取按摩营业类型指定详情
export function View(params) {
  return http.request({
    url: '/spaCooperateType/view',
    method: 'GET',
    params,
  });
}

// 导出按摩营业类型
export function Export(params) {
  jumpExport('/spaCooperateType/export', params);
}


