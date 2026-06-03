import { http } from '@/utils/http/axios';

// 获取预订单列表
export function List(params) {
  return http.request({
    url: '/carOrder/list',
    method: 'get',
    params,
  });
}

// 获取预订单指定详情
export function View(params) {
  return http.request({
    url: '/carOrder/view',
    method: 'GET',
    params,
  });
}

// 订单退款
export function RefundOrder(params) {
  return http.request({
    url: '/carOrder/refund',
    method: 'POST',
    params,
  });
}

// 订单确认
export function ConfirmAgree(params) {
  return http.request({
    url: '/carOrder/confirmAgree',
    method: 'POST',
    params,
  });
}

// 拒绝员工提现
export function ConfirmDisagree(params) {
  return http.request({
    url: '/carOrder/confirmDisagree',
    method: 'POST',
    params,
  });
}

// 获取订单可以选择的司机
export function OrderDriver(params) {
  return http.request({
    url: '/carOrder/driverList',
    method: 'GET',
    params,
  });
}

// 订单调度
export function Dispatch(params) {
  return http.request({
    url: '/carOrder/dispatch',
    method: 'POST',
    params,
  });
}

// 订单转单
export function Transfer(params) {
  return http.request({
    url: '/carOrder/transfer',
    method: 'POST',
    params,
  });
}

// 获取餐厅结算订单列表
export function SettleOrderList(params) {
  return http.request({
    url: '/carOrder/settleOrderList',
    method: 'get',
    params,
  });
}

// 司机出发
export function GoOut(params) {
  return http.request({
    url: '/carOrder/goOut',
    method: 'POST',
    params,
  });
}

// 开始服务
export function StartService(params) {
  return http.request({
    url: '/carOrder/startService',
    method: 'POST',
    params,
  });
}

// 结束服务
export function EndService(params) {
  return http.request({
    url: '/carOrder/endService',
    method: 'POST',
    params,
  });
}

// 异常处理
export function Abnormal(params) {
  return http.request({
    url: '/carOrder/abnormal',
    method: 'POST',
    params,
  });
}

// 导出
export function ExportOrder(params) {
  return http.request({
    url: '/carOrder/export',
    method: 'POST',
    params,
  });
}

// 获取导出列表
export function ExportList(params) {
  return http.request({
    url: '/carOrder/exportList',
    method: 'get',
    params,
  });
}
