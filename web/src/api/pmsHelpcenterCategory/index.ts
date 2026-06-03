import { http, jumpExport } from '@/utils/http/axios';

// 获取帮助中心分类列表
export function List(params) {
  return http.request({
    url: '/pmsHelpcenterCategory/list',
    method: 'get',
    params,
  });
}

// 获取帮助中心分类全部列表
export function All(params) {
  return http.request({
    url: '/pmsHelpcenterCategory/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除帮助中心分类
export function Delete(params) {
  return http.request({
    url: '/pmsHelpcenterCategory/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑帮助中心分类
export function Edit(params) {
  return http.request({
    url: '/pmsHelpcenterCategory/edit',
    method: 'POST',
    params,
  });
}

// 获取帮助中心分类指定详情
export function View(params) {
  return http.request({
    url: '/pmsHelpcenterCategory/view',
    method: 'GET',
    params,
  });
}


