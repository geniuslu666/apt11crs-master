import { http, jumpExport } from '@/utils/http/axios';

// 获取礼品券列表
export function List(params) {
  return http.request({
    url: '/thMemberCoupon/list',
    method: 'get',
    params,
  });
}

// 导出支付流水
export function Export(params) {
  jumpExport('/thMemberCoupon/export', params);
}

// 获取会员礼品券指详情
export function View(params) {
  return http.request({
    url: '/thMemberCoupon/view',
    method: 'GET',
    params,
  });
}

// 回收礼品券
export function Recycle(params) {
  return http.request({
    url: '/thMemberCoupon/recycle',
    method: 'POST',
    params,
  });
}

// 手动核销
export function ManualVerify(params) {
  return http.request({
    url: '/thMemberCoupon/manualVerify',
    method: 'POST',
    params,
  });
}
