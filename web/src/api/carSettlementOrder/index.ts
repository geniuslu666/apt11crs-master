import { http } from '@/utils/http/axios';

// 获取结算单列表
export function List(params) {
  return http.request({
    url: '/carSettlementOrder/list',
    method: 'get',
    params,
  });
}

// 结算概况
export function Stat(params) {
  return http.request({
    url: '/carSettlementOrder/stat',
    method: 'get',
    params,
  });
}

// 结算详情
export function View(params) {
  return http.request({
    url: '/carSettlementOrder/view',
    method: 'GET',
    params,
  });
}

// 结算单-核账
export function Verify(params) {
  return http.request({
    url: '/carSettlementOrder/verify',
    method: 'POST',
    params,
  });
}
