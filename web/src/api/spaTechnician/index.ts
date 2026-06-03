import { http } from '@/utils/http/axios';

// 获取技师管理列表
export function List(params) {
    return http.request({
        url: '/spaTechnician/list',
        method: 'get',
        params,
    });
}

// 获取技师列表
export function All(params) {
  return http.request({
    url: '/spaTechnician/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除技师管理
export function Delete(params) {
    return http.request({
        url: '/spaTechnician/delete',
        method: 'POST',
        params,
    });
}

// 添加/编辑技师管理
export function Edit(params) {
    return http.request({
        url: '/spaTechnician/edit',
        method: 'POST',
        params,
    });
}

// 修改技师管理状态
export function WorkStatus(params) {
  return http.request({
    url: '/spaTechnician/workStatus',
    method: 'POST',
    params,
  });
}

// 修改技师管理状态
export function Status(params) {
    return http.request({
        url: '/spaTechnician/status',
        method: 'POST',
        params,
    });
}

// 获取技师管理指定详情
export function View(params) {
    return http.request({
        url: '/spaTechnician/view',
        method: 'GET',
        params,
    });
}

// 绑定用户
export function Bind(params) {
  return http.request({
    url: '/spaTechnician/bind',
    method: 'POST',
    params,
  });
}

// 解绑用户
export function Unbind(params) {
  return http.request({
    url: '/spaTechnician/unbind',
    method: 'POST',
    params,
  });
}


