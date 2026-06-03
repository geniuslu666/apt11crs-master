import { http } from '@/utils/http/axios';

// 获取门店管理列表
export function List(params) {
  return http.request({
    url: '/spaStore/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除门店管理
export function Delete(params) {
  return http.request({
    url: '/spaStore/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑门店管理
export function Edit(params) {
  return http.request({
    url: '/spaStore/edit',
    method: 'POST',
    params,
  });
}

// 获取门店管理指定详情
export function View(params) {
  return http.request({
    url: '/spaStore/view',
    method: 'GET',
    params,
  });
}

// 获取门店管理指定详情
export function Latest() {
  return http.request({
    url: '/spaStore/latest',
    method: 'GET',
  });
}

