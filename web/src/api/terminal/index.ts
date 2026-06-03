import { http } from '@/utils/http/axios';

// 获取终端列表
export function List(params) {
  return http.request({
    url: '/terminal/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除终端
export function Delete(params) {
  return http.request({
    url: '/terminal/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑终端
export function Edit(params) {
  return http.request({
    url: '/terminal/edit',
    method: 'POST',
    params,
  });
}

// 获取终端指定详情
export function View(params) {
  return http.request({
    url: '/terminal/view',
    method: 'GET',
    params,
  });
}

// 获取品牌型号列表
export function BrandList(params) {
  return http.request({
    url: '/brandModel/list',
    method: 'get',
    params,
  });
}


// 添加/编辑终端
export function BrandEdit(params) {
  return http.request({
    url: '/brandModel/edit',
    method: 'POST',
    params,
  });
}

// 获取终端指定详情
export function BrandView(params) {
  return http.request({
    url: '/brandModel/view',
    method: 'GET',
    params,
  });
}
