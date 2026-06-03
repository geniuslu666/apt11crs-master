import { http } from '@/utils/http/axios';

// 获取Maintenance列表
export function List(params) {
  return http.request({
    url: '/carMaintenance/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除Maintenance
export function Delete(params) {
  return http.request({
    url: '/carMaintenance/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑Maintenance
export function Edit(params) {
  return http.request({
    url: '/carMaintenance/edit',
    method: 'POST',
    params,
  });
}

// 获取Maintenance指定详情
export function View(params) {
  return http.request({
    url: '/carMaintenance/view',
    method: 'GET',
    params,
  });
}

// 获取MaintenanceLanguage列表
export function LanguageList(params) {
  return http.request({
    url: '/carMaintenance/languageList',
    method: 'get',
    params,
  });
}


