import { http } from '@/utils/http/axios';

// 获取列表
export function List(params) {
  return http.request({
    url: '/pmsIndexNav/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除
export function Delete(params) {
  return http.request({
    url: '/pmsIndexNav/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑
export function Edit(params) {
  return http.request({
    url: '/pmsIndexNav/edit',
    method: 'POST',
    params,
  });
}

// 获取详情
export function View(params) {
  return http.request({
    url: '/pmsIndexNav/view',
    method: 'GET',
    params,
  });
}

// 修改状态
export function Status(params) {
  return http.request({
    url: '/pmsIndexNav/status',
    method: 'POST',
    params,
  });
}

// 修改状态
export function Switch(params) {
  return http.request({
    url: '/pmsIndexNav/switch',
    method: 'POST',
    params,
  });
}

// 排序
export function Sort(params) {
  return http.request({
    url: '/pmsIndexNav/sortUpdate',
    method: 'POST',
    params,
  });
}

// 修改小程序显示状态
export function MinappStatus(params) {
  return http.request({
    url: '/pmsIndexNav/minappStatus',
    method: 'POST',
    params,
  });
}
