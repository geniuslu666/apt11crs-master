import { http } from '@/utils/http/axios';

// 获取服务商管理列表
export function List(params) {
    return http.request({
        url: '/spaIsp/list',
        method: 'get',
        params,
    });
}

// 获取服务商列表
export function All(params) {
  return http.request({
    url: '/spaIsp/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除服务商管理
export function Delete(params) {
    return http.request({
        url: '/spaIsp/delete',
        method: 'POST',
        params,
    });
}

// 添加/编辑服务商管理
export function Edit(params) {
    return http.request({
        url: '/spaIsp/edit',
        method: 'POST',
        params,
    });
}

// 修改服务商管理状态
export function WorkStatus(params) {
  return http.request({
    url: '/spaIsp/workStatus',
    method: 'POST',
    params,
  });
}

// 修改服务商管理状态
export function Status(params) {
    return http.request({
        url: '/spaIsp/status',
        method: 'POST',
        params,
    });
}

// 获取服务商管理指定详情
export function View(params) {
    return http.request({
        url: '/spaIsp/view',
        method: 'GET',
        params,
    });
}

// 绑定用户
export function Bind(params) {
  return http.request({
    url: '/spaIsp/bind',
    method: 'POST',
    params,
  });
}

// 解绑用户
export function Unbind(params) {
  return http.request({
    url: '/spaIsp/unbind',
    method: 'POST',
    params,
  });
}


