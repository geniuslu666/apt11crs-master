import { http } from '@/utils/http/axios';

// 获取区域管理列表
export function List(params) {
  return http.request({
    url: '/foodArea/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除区域管理
export function Delete(params) {
  return http.request({
    url: '/foodArea/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑区域管理
export function Edit(params) {
  return http.request({
    url: '/foodArea/edit',
    method: 'POST',
    params,
  });
}

// 获取区域管理指定详情
export function View(params) {
  return http.request({
    url: '/foodArea/view',
    method: 'GET',
    params,
  });
}

// 修改区域管理状态
export function Status(params) {
  return http.request({
    url: '/foodArea/status',
    method: 'POST',
    params,
  });
}


