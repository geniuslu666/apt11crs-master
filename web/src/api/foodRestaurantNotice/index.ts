import { http } from '@/utils/http/axios';

// 获取商家通知列表
export function List(params) {
  return http.request({
    url: '/foodRestaurantNotice/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除商家通知
export function Delete(params) {
  return http.request({
    url: '/foodRestaurantNotice/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑商家通知
export function Edit(params) {
  return http.request({
    url: '/foodRestaurantNotice/edit',
    method: 'POST',
    params,
  });
}

// 修改商家通知状态
export function Status(params) {
  return http.request({
    url: '/foodRestaurantNotice/status',
    method: 'POST',
    params,
  });
}

// 获取商家通知指定详情
export function View(params) {
  return http.request({
    url: '/foodRestaurantNotice/view',
    method: 'GET',
    params,
  });
}

// 获取商家通知最大排序
export function MaxSort() {
  return http.request({
    url: '/foodRestaurantNotice/maxSort',
    method: 'GET',
  });
}


