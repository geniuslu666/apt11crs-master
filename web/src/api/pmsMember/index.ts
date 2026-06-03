import { http, jumpExport } from '@/utils/http/axios';

// 获取会员信息列表
export function List(params) {
  return http.request({
    url: '/pmsMember/list',
    method: 'get',
    params,
  });
}

// 获取会员信息列表
export function SelectList(params) {
  return http.request({
    url: '/pmsMember/select',
    method: 'get',
    params,
  });
}

// 删除/批量删除会员信息
export function Delete(params) {
  return http.request({
    url: '/pmsMember/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑会员信息
export function Edit(params) {
  return http.request({
    url: '/pmsMember/edit',
    method: 'POST',
    params,
  });
}

// 添加/编辑会员信息
export function BaseEdit(params) {
  return http.request({
    url: '/pmsMember/baseEdit',
    method: 'POST',
    params,
  });
}

// 调整积分
export function BalanceEdit(params) {
  return http.request({
    url: '/pmsMember/balanceEdit',
    method: 'POST',
    params,
  });
}

// 调整成长值
export function ExpEdit(params) {
  return http.request({
    url: '/pmsMember/expEdit',
    method: 'POST',
    params,
  });
}

// 发送私信
export function SendMsg(params) {
  return http.request({
    url: '/pmsMember/sendMsg',
    method: 'POST',
    params,
  });
}

//https://crsdev.yeebok.cn/admin/pmsMember/sendNotify
export function sendNotify(params) {
  return http.request({
    url: '/pmsMember/sendNotify',
    method: 'POST',
    params,
  });
}

// 获取会员信息指定详情
export function View(params) {
  return http.request({
    url: '/pmsMember/view',
    method: 'GET',
    params,
  });
}

// 导出会员信息
export function Export(params) {
  jumpExport('/pmsMember/export', params);
}

// 获取会员统计数据
export function Stat(params) {
  return http.request({
    url: '/pmsMember/stat',
    method: 'get',
    params,
  });
}

// 修改会员状态
export function Status(params) {
  return http.request({
    url: '/pmsMember/status',
    method: 'POST',
    params,
  });
}

// 会员注销
export function Cancel(params) {
  return http.request({
    url: '/pmsMember/cancel',
    method: 'POST',
    params,
  });
}

// 发送注册登录验证码
export function SendRegisterCode(params) {
  return http.request({
    url: '/pmsMember/sendRegisterCode',
    method: 'POST',
    params,
  });
}
