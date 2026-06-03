import { http } from '@/utils/http/axios';

// 获取首页文章列表
export function List(params) {
  return http.request({
    url: '/article/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除首页文章
export function Delete(params) {
  return http.request({
    url: '/article/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑首页文章
export function Edit(params) {
  return http.request({
    url: '/article/edit',
    method: 'POST',
    params,
  });
}

// 获取首页文章指定详情
export function View(params) {
  return http.request({
    url: '/article/view',
    method: 'GET',
    params,
  });
}

// 修改首页活动状态
export function Status(params) {
  return http.request({
    url: '/article/status',
    method: 'POST',
    params,
  });
}
