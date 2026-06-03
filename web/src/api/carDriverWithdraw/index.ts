import { http } from '@/utils/http/axios';

// 获取司机提现申请表列表
export function WithdrawList(params) {
  return http.request({
    url: '/carDriverWithdraw/list',
    method: 'get',
    params,
  });
}

// 获取提现申请表指定详情
export function View(params) {
  return http.request({
    url: '/carDriverWithdraw/view',
    method: 'GET',
    params,
  });
}

// 同意司机提现
export function Agree(params) {
  return http.request({
    url: '/carDriverWithdraw/agree',
    method: 'POST',
    params,
  });
}

// 拒绝司机提现
export function Disagree(params) {
  return http.request({
    url: '/carDriverWithdraw/disagree',
    method: 'POST',
    params,
  });
}

// 司机提现转账
export function Transfer(params) {
  return http.request({
    url: 'carDriverWithdraw/transfer',
    method: 'POST',
    params,
  });
}

