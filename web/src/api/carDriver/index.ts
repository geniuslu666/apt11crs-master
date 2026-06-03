import { http } from '@/utils/http/axios';

// 获取司机管理列表
export function List(params) {
    return http.request({
        url: '/carDriver/list',
        method: 'get',
        params,
    });
}

// 获取司机列表
export function All(params) {
  return http.request({
    url: '/carDriver/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除司机管理
export function Delete(params) {
    return http.request({
        url: '/carDriver/delete',
        method: 'POST',
        params,
    });
}

// 添加/编辑司机管理
export function Edit(params) {
    return http.request({
        url: '/carDriver/edit',
        method: 'POST',
        params,
    });
}

// 修改司机管理状态
export function WorkStatus(params) {
  return http.request({
    url: '/carDriver/workStatus',
    method: 'POST',
    params,
  });
}

// 修改司机管理状态
export function Status(params) {
    return http.request({
        url: '/carDriver/status',
        method: 'POST',
        params,
    });
}

// 获取司机管理指定详情
export function View(params) {
    return http.request({
        url: '/carDriver/view',
        method: 'GET',
        params,
    });
}

// 绑定用户
export function Bind(params) {
  return http.request({
    url: '/carDriver/bind',
    method: 'POST',
    params,
  });
}

// 解绑用户
export function Unbind(params) {
  return http.request({
    url: '/carDriver/unbind',
    method: 'POST',
    params,
  });
}

// 绑定车辆
export function BindCar(params) {
  return http.request({
    url: '/carDriver/bindCar',
    method: 'POST',
    params,
  });
}
