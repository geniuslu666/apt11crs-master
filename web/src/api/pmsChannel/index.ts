import { http } from '@/utils/http/axios';

// 获取渠道管理列表
export function List(params) {
  return http.request({
    url: '/pmsChannel/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除渠道管理
export function Delete(params) {
  return http.request({
    url: '/pmsChannel/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑渠道管理
export function Edit(params) {
  return http.request({
    url: '/pmsChannel/edit',
    method: 'POST',
    params,
  });
}

// 修改渠道管理状态
export function Status(params) {
  return http.request({
    url: '/pmsChannel/status',
    method: 'POST',
    params,
  });
}

// 获取渠道管理指定详情
export function View(params) {
  return http.request({
    url: '/pmsChannel/view',
    method: 'GET',
    params,
  });
}

// 渠道绑定会员
export function Bind(params) {
  return http.request({
    url: '/pmsChannel/bind',
    method: 'POST',
    params,
  });
}

// 解绑渠道管理
export function Unbind(params) {
  return http.request({
    url: '/pmsChannel/unbind',
    method: 'POST',
    params,
  });
}


