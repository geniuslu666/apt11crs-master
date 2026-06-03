import { http } from '@/utils/http/axios';

// 获取Banner列表
export function List(params) {
  return http.request({
    url: '/carBanner/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除Banner
export function Delete(params) {
  return http.request({
    url: '/carBanner/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑Banner
export function Edit(params) {
  return http.request({
    url: '/carBanner/edit',
    method: 'POST',
    params,
  });
}

// 修改Banner状态
export function Status(params) {
  return http.request({
    url: '/carBanner/status',
    method: 'POST',
    params,
  });
}

// 获取Banner指定详情
export function View(params) {
  return http.request({
    url: '/carBanner/view',
    method: 'GET',
    params,
  });
}

// 获取Banner最大排序
export function MaxSort() {
  return http.request({
    url: '/carBanner/maxSort',
    method: 'GET',
  });
}


