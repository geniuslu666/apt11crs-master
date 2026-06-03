import { http, jumpExport } from '@/utils/http/axios';

// 获取订餐合作类型列表
export function List(params) {
  return http.request({
    url: '/foodCooperateType/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除订餐合作类型
export function Delete(params) {
  return http.request({
    url: '/foodCooperateType/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑订餐合作类型
export function Edit(params) {
  return http.request({
    url: '/foodCooperateType/edit',
    method: 'POST',
    params,
  });
}

// 修改订餐合作类型状态
export function Status(params) {
  return http.request({
    url: '/foodCooperateType/status',
    method: 'POST',
    params,
  });
}

// 获取订餐合作类型指定详情
export function View(params) {
  return http.request({
    url: '/foodCooperateType/view',
    method: 'GET',
    params,
  });
}

// 导出订餐合作类型
export function Export(params) {
  jumpExport('/foodCooperateType/export', params);
}


