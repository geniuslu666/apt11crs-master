import { http } from '@/utils/http/axios';

// 获取礼品券分类列表
export function List(params) {
  return http.request({
    url: '/thCouponCategory/list',
    method: 'get',
    params,
  });
}

// 获取礼品券分类全部列表
export function All(params) {
  return http.request({
    url: '/thCouponCategory/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除礼品券分类
export function Delete(params) {
  return http.request({
    url: '/thCouponCategory/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑礼品券分类
export function Edit(params) {
  return http.request({
    url: '/thCouponCategory/edit',
    method: 'POST',
    params,
  });
}

// 获取礼品券分类指定详情
export function View(params) {
  return http.request({
    url: '/thCouponCategory/view',
    method: 'GET',
    params,
  });
}

// 状态开启关闭
export function Switch(params) {
  return http.request({
    url: '/thCouponCategory/switch',
    method: 'POST',
    params,
  });
}


