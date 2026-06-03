import { http } from '@/utils/http/axios';

// 获取通知模版列表
export function List(params) {
  return http.request({
    url: '/smsTemplate/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除通知模版
export function Delete(params) {
  return http.request({
    url: '/smsTemplate/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑通知模版
export function Edit(params) {
  return http.request({
    url: '/smsTemplate/edit',
    method: 'POST',
    params,
  });
}

// 获取通知模版指定详情
export function View(params) {
  return http.request({
    url: '/smsTemplate/view',
    method: 'GET',
    params,
  });
}

// 获取发送记录列表
export function Log(params) {
  return http.request({
    url: '/smsTemplate/log',
    method: 'get',
    params,
  });
}

