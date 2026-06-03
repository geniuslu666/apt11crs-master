import { http } from '@/utils/http/axios';

// 获取套餐管理列表
export function List(params) {
    return http.request({
        url: '/foodGoods/list',
        method: 'get',
        params,
    });
}

// 删除/批量删除套餐管理
export function Delete(params) {
    return http.request({
        url: '/foodGoods/delete',
        method: 'POST',
        params,
    });
}

// 添加/编辑套餐管理
export function Edit(params) {
    return http.request({
        url: '/foodGoods/edit',
        method: 'POST',
        params,
    });
}

// 修改套餐状态
export function Status(params) {
    return http.request({
        url: '/foodGoods/status',
        method: 'POST',
        params,
    });
}

// 获取套餐管理指定详情
export function View(params) {
    return http.request({
        url: '/foodGoods/view',
        method: 'GET',
        params,
    });
}


