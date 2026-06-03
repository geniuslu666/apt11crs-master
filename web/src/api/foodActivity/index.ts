import { http, jumpExport } from '@/utils/http/axios';

// 获取订餐活动表列表
export function List(params) {
  return http.request({
    url: '/foodActivity/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除订餐活动表
export function Delete(params) {
  return http.request({
    url: '/foodActivity/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑订餐活动表
export function Edit(params) {
  return http.request({
    url: '/foodActivity/edit',
    method: 'POST',
    params,
  });
}

// 修改订餐活动表状态
export function Status(params) {
  return http.request({
    url: '/foodActivity/status',
    method: 'POST',
    params,
  });
}

// 获取订餐活动表指定详情
export function View(params) {
  return http.request({
    url: '/foodActivity/view',
    method: 'GET',
    params,
  });
}

// 导出订餐活动表
export function Export(params) {
  jumpExport('/foodActivity/export', params);
}


