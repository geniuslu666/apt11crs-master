import { http } from '@/utils/http/axios';

// 获取结算模式列表
export function List(params) {
  return http.request({
    url: '/carSettlement/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除结算模式
export function Delete(params) {
  return http.request({
    url: '/carSettlement/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑结算模式
export function Edit(params) {
  return http.request({
    url: '/carSettlement/edit',
    method: 'POST',
    params,
  });
}

// 修改结算模式状态
export function Status(params) {
  return http.request({
    url: '/carSettlement/status',
    method: 'POST',
    params,
  });
}

// 获取结算模式指定详情
export function View(params) {
  return http.request({
    url: '/carSettlement/view',
    method: 'GET',
    params,
  });
}


