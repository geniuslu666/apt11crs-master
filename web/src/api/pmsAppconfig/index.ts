import { http } from '@/utils/http/axios';

// 获取APP配置列表
export function List(params) {
  return http.request({
    url: '/pmsAppconfig/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除APP配置
export function Delete(params) {
  return http.request({
    url: '/pmsAppconfig/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑APP配置
export function Edit(params) {
  return http.request({
    url: '/pmsAppconfig/edit',
    method: 'POST',
    params,
  });
}

// 获取APP配置指定详情
export function View(params) {
  return http.request({
    url: '/pmsAppconfig/view',
    method: 'GET',
    params,
  });
}
//http://gyconf.nat300.top/admin/pmsMember/sendEmail
//发送邮件
export function sendEmll(params) {
  return http.request({
    url: '/pmsMember/sendEmail',
    method: 'POST',
    params,
  });
}