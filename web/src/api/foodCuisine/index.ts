import { http } from '@/utils/http/axios';

// 获取餐厅菜系列表
export function List(params) {
  return http.request({
    url: '/foodCuisine/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除餐厅菜系
export function Delete(params) {
  return http.request({
    url: '/foodCuisine/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑餐厅菜系
export function Edit(params) {
  return http.request({
    url: '/foodCuisine/edit',
    method: 'POST',
    params,
  });
}

// 修改餐厅菜系状态
export function Status(params) {
  return http.request({
    url: '/foodCuisine/status',
    method: 'POST',
    params,
  });
}

// 获取餐厅菜系指定详情
export function View(params) {
  return http.request({
    url: '/foodCuisine/view',
    method: 'GET',
    params,
  });
}

// 获取菜系最大排序
export function MaxSort() {
  return http.request({
    url: '/foodCuisine/maxSort',
    method: 'GET',
  });
}


