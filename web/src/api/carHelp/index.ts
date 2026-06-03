import { http } from '@/utils/http/axios';

// 获取Help列表
export function List(params) {
  return http.request({
    url: '/carHelp/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除Help
export function Delete(params) {
  return http.request({
    url: '/carHelp/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑Help
export function Edit(params) {
  return http.request({
    url: '/carHelp/edit',
    method: 'POST',
    params,
  });
}

// 获取Help指定详情
export function View(params) {
  return http.request({
    url: '/carHelp/view',
    method: 'GET',
    params,
  });
}

// 获取HelpLanguage列表
export function LanguageList(params) {
  return http.request({
    url: '/carHelp/languageList',
    method: 'get',
    params,
  });
}


