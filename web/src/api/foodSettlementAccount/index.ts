import { http } from '@/utils/http/axios';

// 获取订餐结算账户表列表
export function List(params) {
  return http.request({
    url: '/foodSettlementAccount/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除订餐结算账户表
export function Delete(params) {
  return http.request({
    url: '/foodSettlementAccount/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑订餐结算账户表
export function Edit(params) {
  return http.request({
    url: '/foodSettlementAccount/edit',
    method: 'POST',
    params,
  });
}

// 修改订餐结算账户表状态
export function Status(params) {
  return http.request({
    url: '/foodSettlementAccount/status',
    method: 'POST',
    params,
  });
}

// 获取订餐结算账户表指定详情
export function View(params) {
  return http.request({
    url: '/foodSettlementAccount/view',
    method: 'GET',
    params,
  });
}


