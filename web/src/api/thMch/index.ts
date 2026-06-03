import { http } from '@/utils/http/axios';

// 获取商户列表
export function List(params) {
  return http.request({
    url: '/thMch/list',
    method: 'get',
    params,
  });
}

// 获取商户全部列表
export function All(params) {
  return http.request({
    url: '/thMch/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除商户
export function Delete(params) {
  return http.request({
    url: '/thMch/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑商户
export function Edit(params) {
  return http.request({
    url: '/thMch/edit',
    method: 'POST',
    params,
  });
}

// 获取商户指定详情
export function View(params) {
  return http.request({
    url: '/thMch/view',
    method: 'GET',
    params,
  });
}

// 状态开启关闭
export function Switch(params) {
  return http.request({
    url: '/thMch/switch',
    method: 'POST',
    params,
  });
}


