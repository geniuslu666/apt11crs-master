import { http } from '@/utils/http/axios';

// 获取商户分类列表
export function List(params) {
  return http.request({
    url: '/thMchCategory/list',
    method: 'get',
    params,
  });
}

// 获取商户分类全部列表
export function All(params) {
  return http.request({
    url: '/thMchCategory/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除商户分类
export function Delete(params) {
  return http.request({
    url: '/thMchCategory/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑商户分类
export function Edit(params) {
  return http.request({
    url: '/thMchCategory/edit',
    method: 'POST',
    params,
  });
}

// 获取商户分类指定详情
export function View(params) {
  return http.request({
    url: '/thMchCategory/view',
    method: 'GET',
    params,
  });
}

// 状态开启关闭
export function Switch(params) {
  return http.request({
    url: '/thMchCategory/switch',
    method: 'POST',
    params,
  });
}


