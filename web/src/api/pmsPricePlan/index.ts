import { http } from '@/utils/http/axios';

// 获取房型列表
export function List(params) {
  return http.request({
    url: '/pricePlan/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除房型
export function Delete(params) {
  return http.request({
    url: '/pricePlan/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑房型
export function Edit(params) {
  return http.request({
    url: '/pricePlan/edit',
    method: 'POST',
    params,
  });
}

// 获取房型指定详情
export function View(params) {
  return http.request({
    url: '/pricePlan/view',
    method: 'GET',
    params,
  });
}

// 修改状态
export function Status(params) {
  return http.request({
    url: '/pricePlan/status',
    method: 'POST',
    params,
  });
}

// 价格计划-恢复
export function Recycle(params) {
  return http.request({
    url: '/pricePlan/recycle',
    method: 'POST',
    params,
  });
}
