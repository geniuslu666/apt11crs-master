import { http, jumpExport } from '@/utils/http/axios';

// 获取帮助中心列表
export function List(params) {
  return http.request({
    url: '/pmsHelpcenter/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除帮助中心
export function Delete(params) {
  return http.request({
    url: '/pmsHelpcenter/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑帮助中心
export function Edit(params) {
  return http.request({
    url: '/pmsHelpcenter/edit',
    method: 'POST',
    params,
  });
}

// 获取帮助中心指定详情
export function View(params) {
  return http.request({
    url: '/pmsHelpcenter/view',
    method: 'GET',
    params,
  });
}


