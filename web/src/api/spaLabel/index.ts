import { http } from '@/utils/http/axios';

// 获取餐厅标签列表
export function List(params) {
  return http.request({
    url: '/spaLabel/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除餐厅标签
export function Delete(params) {
  return http.request({
    url: '/spaLabel/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑餐厅标签
export function Edit(params) {
  return http.request({
    url: '/spaLabel/edit',
    method: 'POST',
    params,
  });
}

// 修改餐厅标签状态
export function Status(params) {
  return http.request({
    url: '/spaLabel/status',
    method: 'POST',
    params,
  });
}

// 获取餐厅标签指定详情
export function View(params) {
  return http.request({
    url: '/spaLabel/view',
    method: 'GET',
    params,
  });
}

// 获取餐厅标签最大排序
export function MaxSort() {
  return http.request({
    url: '/spaLabel/maxSort',
    method: 'GET',
  });
}


