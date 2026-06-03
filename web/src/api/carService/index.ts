import { http } from '@/utils/http/axios';

// 获取服务列表
export function List(params) {
  return http.request({
    url: '/carService/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除服务
export function Delete(params) {
  return http.request({
    url: '/carService/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑服务
export function Edit(params) {
  return http.request({
    url: '/carService/edit',
    method: 'POST',
    params,
  });
}

// 修改服务状态
export function Status(params) {
  return http.request({
    url: '/carService/status',
    method: 'POST',
    params,
  });
}

// 获取服务指定详情
export function View(params) {
  return http.request({
    url: '/carService/view',
    method: 'GET',
    params,
  });
}

// 获取服务最大排序
export function MaxSort() {
  return http.request({
    url: '/carService/maxSort',
    method: 'GET',
  });
}

// 服务排序
export function Sort(params) {
  return http.request({
    url: '/carService/sortUpdate',
    method: 'POST',
    params,
  });
}

