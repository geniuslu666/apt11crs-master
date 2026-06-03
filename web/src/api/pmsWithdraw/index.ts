import { http } from '@/utils/http/axios';

// 获取员工提现申请表列表
export function StaffList(params) {
  return http.request({
    url: '/pmsWithdraw/staffList',
    method: 'get',
    params,
  });
}

// 获取渠道提现申请表列表
export function ChannelList(params) {
  return http.request({
    url: '/pmsWithdraw/channelList',
    method: 'get',
    params,
  });
}

// 获取提现申请表指定详情
export function View(params) {
  return http.request({
    url: '/pmsWithdraw/view',
    method: 'GET',
    params,
  });
}

// 同意员工提现
export function AgreeStaff(params) {
  return http.request({
    url: '/pmsWithdraw/agreeStaff',
    method: 'POST',
    params,
  });
}

// 拒绝员工提现
export function DisagreeStaff(params) {
  return http.request({
    url: '/pmsWithdraw/disagreeStaff',
    method: 'POST',
    params,
  });
}

// 同意渠道提现
export function AgreeChannel(params) {
  return http.request({
    url: '/pmsWithdraw/agreeChannel',
    method: 'POST',
    params,
  });
}

// 拒绝员工提现
export function DisagreeChannel(params) {
  return http.request({
    url: '/pmsWithdraw/disagreeChannel',
    method: 'POST',
    params,
  });
}

// 员工提现转账
export function TransferStaff(params) {
  return http.request({
    url: 'pmsWithdraw/transferStaff',
    method: 'POST',
    params,
  });
}

// 渠道提现转账
export function TransferChannel(params) {
  return http.request({
    url: '/pmsWithdraw/transferChannel',
    method: 'POST',
    params,
  });
}
