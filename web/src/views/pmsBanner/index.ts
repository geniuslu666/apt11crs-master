import { http, jumpExport } from '@/utils/http/axios';

// 获取banner 横幅列表
export function List(params) {
  return http.request({
    url: '/pmsBanner/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除banner 横幅
export function Delete(params) {
  return http.request({
    url: '/pmsBanner/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑banner 横幅
export function Edit(params) {
  return http.request({
    url: '/pmsBanner/edit',
    method: 'POST',
    params,
  });
}

// 获取banner 横幅指定详情
export function View(params) {
  return http.request({
    url: '/pmsBanner/view',
    method: 'GET',
    params,
  });
}

// 导出banner 横幅
export function Export(params) {
  jumpExport('/pmsBanner/export', params);
}


