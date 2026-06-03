import { http } from '@/utils/http/axios';

// 获取订餐-座位表列表
export function List(params) {
  return http.request({
    url: '/foodSeat/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除订餐-座位表
export function Delete(params) {
  return http.request({
    url: '/foodSeat/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑订餐-座位表
export function Edit(params) {
  return http.request({
    url: '/foodSeat/edit',
    method: 'POST',
    params,
  });
}

// 修改订餐-座位表状态
export function Status(params) {
  return http.request({
    url: '/foodSeat/status',
    method: 'POST',
    params,
  });
}

// 获取订餐-座位表指定详情
export function View(params) {
  return http.request({
    url: '/foodSeat/view',
    method: 'GET',
    params,
  });
}


