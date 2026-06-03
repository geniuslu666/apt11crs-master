import { http } from '@/utils/http/axios';

// 获取餐厅套餐预订单列表
export function List(params) {
  return http.request({
    url: '/foodOrder/list',
    method: 'get',
    params,
  });
}

// 获取餐厅套餐预订单指定详情
export function View(params) {
  return http.request({
    url: '/foodOrder/view',
    method: 'GET',
    params,
  });
}

// 订单确认
export function ConfirmAgree(params) {
  return http.request({
    url: '/foodOrder/confirmAgree',
    method: 'POST',
    params,
  });
}

// 拒绝员工提现
export function ConfirmDisagree(params) {
  return http.request({
    url: '/foodOrder/confirmDisagree',
    method: 'POST',
    params,
  });
}

// 获取餐厅结算订单列表
export function SettleOrderList(params) {
  return http.request({
    url: '/foodOrder/settleOrderList',
    method: 'get',
    params,
  });
}

// 取消
export function CancelPay(params) {
  return http.request({
    url: '/foodOrder/cancelPay',
    method: 'POST',
    params,
  });
}

// 导出
export function ExportOrder(params) {
  return http.request({
    url: '/foodOrder/export',
    method: 'POST',
    params,
  });
}

// 获取导出列表
export function ExportList(params) {
  return http.request({
    url: '/foodOrder/exportList',
    method: 'get',
    params,
  });
}
