import { http } from '@/utils/http/axios';

// 获取车辆车型列表
export function List(params) {
  return http.request({
    url: '/printer/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除车辆车型
export function Delete(params) {
  return http.request({
    url: '/printer/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑车辆车型
export function Edit(params) {
  return http.request({
    url: '/printer/edit',
    method: 'POST',
    params,
  });
}

// 修改车辆车型状态
export function Status(params) {
  return http.request({
    url: '/printer/status',
    method: 'POST',
    params,
  });
}

// 获取车辆车型指定详情
export function View(params) {
  return http.request({
    url: '/printer/view',
    method: 'GET',
    params,
  });
}

// 获取车辆车型最大排序
export function MaxSort() {
  return http.request({
    url: '/printer/maxSort',
    method: 'GET',
  });
}

// 打印接送机订单
export function PrinterCarOrder(params) {
  return http.request({
    url: '/printer/printCarOrder',
    method: 'POST',
    params,
  });
}

// 打印餐厅订单
export function PrinterFoodOrder(params) {
  return http.request({
    url: '/printer/printFoodOrder',
    method: 'POST',
    params,
  });
}

// 打印按摩订单
export function PrinterSpaOrder(params) {
  return http.request({
    url: '/printer/printSpaOrder',
    method: 'POST',
    params,
  });
}

