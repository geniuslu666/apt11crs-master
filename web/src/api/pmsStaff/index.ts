import { http } from '@/utils/http/axios';

// 获取员工管理列表
export function List(params) {
  return http.request({
    url: '/pmsStaff/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除员工管理
export function Delete(params) {
  return http.request({
    url: '/pmsStaff/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑员工管理
export function Edit(params) {
  return http.request({
    url: '/pmsStaff/edit',
    method: 'POST',
    params,
  });
}

// 修改员工管理状态
export function Status(params) {
  return http.request({
    url: '/pmsStaff/status',
    method: 'POST',
    params,
  });
}

// 获取员工管理指定详情
export function View(params) {
  return http.request({
    url: '/pmsStaff/view',
    method: 'GET',
    params,
  });
}

// 员工绑定会员
export function Bind(params) {
  return http.request({
    url: '/pmsStaff/bind',
    method: 'POST',
    params,
  });
}

// 解绑员工管理
export function Unbind(params) {
  return http.request({
    url: '/pmsStaff/unbind',
    method: 'POST',
    params,
  });
}
