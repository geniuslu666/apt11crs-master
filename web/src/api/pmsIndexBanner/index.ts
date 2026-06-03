import { http } from '@/utils/http/axios';

// 获取列表
export function List(params) {
  return http.request({
    url: '/indexBanner/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除
export function Delete(params) {
  return http.request({
    url: '/indexBanner/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑
export function Edit(params) {
  return http.request({
    url: '/indexBanner/edit',
    method: 'POST',
    params,
  });
}

// 获取详情
export function View(params) {
  return http.request({
    url: '/indexBanner/view',
    method: 'GET',
    params,
  });
}

// 修改状态
export function Status(params) {
  return http.request({
    url: '/indexBanner/status',
    method: 'POST',
    params,
  });
}

// 修改状态
export function Switch(params) {
  return http.request({
    url: '/indexBanner/status',
    method: 'POST',
    params,
  });
}

// 排序
export function Sort(params) {
  return http.request({
    url: '/indexBanner/sortUpdate',
    method: 'POST',
    params,
  });
}
