import { http } from '@/utils/http/axios';

// 获取车辆列表
export function List(params) {
  return http.request({
    url: '/carCar/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除车辆
export function Delete(params) {
  return http.request({
    url: '/carCar/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑车辆
export function Edit(params) {
  return http.request({
    url: '/carCar/edit',
    method: 'POST',
    params,
  });
}

// 修改车辆状态
export function Status(params) {
  return http.request({
    url: '/carCar/status',
    method: 'POST',
    params,
  });
}

// 获取车辆指定详情
export function View(params) {
  return http.request({
    url: '/carCar/view',
    method: 'GET',
    params,
  });
}

// 获取车辆最大排序
export function MaxSort() {
  return http.request({
    url: '/carCar/maxSort',
    method: 'GET',
  });
}

// 修改司机管理状态
export function WorkStatus(params) {
  return http.request({
    url: '/carCar/workStatus',
    method: 'POST',
    params,
  });
}

