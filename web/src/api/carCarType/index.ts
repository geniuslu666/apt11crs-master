import { http } from '@/utils/http/axios';

// 获取车辆车型列表
export function List(params) {
  return http.request({
    url: '/carCarType/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除车辆车型
export function Delete(params) {
  return http.request({
    url: '/carCarType/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑车辆车型
export function Edit(params) {
  return http.request({
    url: '/carCarType/edit',
    method: 'POST',
    params,
  });
}

// 修改车辆车型状态
export function Status(params) {
  return http.request({
    url: '/carCarType/status',
    method: 'POST',
    params,
  });
}

// 获取车辆车型指定详情
export function View(params) {
  return http.request({
    url: '/carCarType/view',
    method: 'GET',
    params,
  });
}

// 获取车辆车型最大排序
export function MaxSort() {
  return http.request({
    url: '/carCarType/maxSort',
    method: 'GET',
  });
}


