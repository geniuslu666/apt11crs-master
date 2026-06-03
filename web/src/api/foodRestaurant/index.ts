import { http } from '@/utils/http/axios';

// 获取餐厅管理列表
export function List(params) {
    return http.request({
        url: '/foodRestaurant/list',
        method: 'get',
        params,
    });
}

// 获取餐厅列表
export function All(params) {
  return http.request({
    url: '/foodRestaurant/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除餐厅管理
export function Delete(params) {
    return http.request({
        url: '/foodRestaurant/delete',
        method: 'POST',
        params,
    });
}

// 添加/编辑餐厅管理
export function Edit(params) {
    return http.request({
        url: '/foodRestaurant/edit',
        method: 'POST',
        params,
    });
}

// 修改餐厅管理状态
export function Status(params) {
    return http.request({
        url: '/foodRestaurant/status',
        method: 'POST',
        params,
    });
}

// 获取餐厅管理指定详情
export function View(params) {
    return http.request({
        url: '/foodRestaurant/view',
        method: 'GET',
        params,
    });
}

// 重置核销码
export function ResetVerifyCode(params) {
  return http.request({
    url: '/foodRestaurant/resetVerifyCode',
    method: 'POST',
    params,
  });
}

// 预定状态开启关闭
export function Switch(params) {
  return http.request({
    url: '/foodRestaurant/switch',
    method: 'POST',
    params,
  });
}

// 餐厅排序
export function Sort(params) {
  return http.request({
    url: '/foodRestaurant/sortUpdate',
    method: 'POST',
    params,
  });
}


