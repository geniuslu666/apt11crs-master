import { http } from '@/utils/http/axios';

// 获取服务列表
export function List(params) {
  return http.request({
    url: '/spaService/list',
    method: 'get',
    params,
  });
}

// 获取服务列表-ALL
export function All(params) {
  return http.request({
    url: '/spaService/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除服务
export function Delete(params) {
  return http.request({
    url: '/spaService/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑服务
export function Edit(params) {
  return http.request({
    url: '/spaService/edit',
    method: 'POST',
    params,
  });
}

// 修改服务状态
export function Status(params) {
  return http.request({
    url: '/spaService/status',
    method: 'POST',
    params,
  });
}

// 获取服务指定详情
export function View(params) {
  return http.request({
    url: '/spaService/view',
    method: 'GET',
    params,
  });
}

// 获取服务最大排序
export function MaxSort() {
  return http.request({
    url: '/spaService/maxSort',
    method: 'GET',
  });
}

// 按摩服务-恢复
export function Recycle(params) {
  return http.request({
    url: '/spaService/recycle',
    method: 'POST',
    params,
  });
}

