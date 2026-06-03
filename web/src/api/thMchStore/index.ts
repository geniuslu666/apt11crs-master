import { http } from '@/utils/http/axios';

// 获取商户门店列表
export function List(params) {
  return http.request({
    url: '/thMchStore/list',
    method: 'get',
    params,
  });
}

// 获取商户门店全部列表
export function All(params) {
  return http.request({
    url: '/thMchStore/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除商户门店
export function Delete(params) {
  return http.request({
    url: '/thMchStore/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑商户门店
export function Edit(params) {
  return http.request({
    url: '/thMchStore/edit',
    method: 'POST',
    params,
  });
}

// 获取商户门店指定详情
export function View(params) {
  return http.request({
    url: '/thMchStore/view',
    method: 'GET',
    params,
  });
}

// 状态开启关闭
export function Switch(params) {
  return http.request({
    url: '/thMchStore/switch',
    method: 'POST',
    params,
  });
}


