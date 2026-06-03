import { http } from '@/utils/http/axios';

// 获取预订单列表
export function List(params) {
  return http.request({
    url: '/spaOrder/list',
    method: 'get',
    params,
  });
}

// 获取预订单指定详情
export function View(params) {
  return http.request({
    url: '/spaOrder/view',
    method: 'GET',
    params,
  });
}

// 订单确认
export function ConfirmAgree(params) {
  return http.request({
    url: '/spaOrder/confirmAgree',
    method: 'POST',
    params,
  });
}

// 拒绝员工提现
export function ConfirmDisagree(params) {
  return http.request({
    url: '/spaOrder/confirmDisagree',
    method: 'POST',
    params,
  });
}

// 获取订单可以选择的技师
export function OrderTechnician(params) {
  return http.request({
    url: '/spaOrder/technicianList',
    method: 'GET',
    params,
  });
}

// 订单调度
export function Dispatch(params) {
  return http.request({
    url: '/spaOrder/dispatch',
    method: 'POST',
    params,
  });
}


// 获取技师结算订单列表
export function TechnicianOrderList(params) {
  return http.request({
    url: '/spaOrder/technicianOrderList',
    method: 'get',
    params,
  });
}

// 出发
export function GoOut(params) {
  return http.request({
    url: '/spaOrder/goOut',
    method: 'POST',
    params,
  });
}

// 到店
export function Arrive(params) {
  return http.request({
    url: '/spaOrder/arrive',
    method: 'POST',
    params,
  });
}

// 开始服务
export function StartService(params) {
  return http.request({
    url: '/spaOrder/startService',
    method: 'POST',
    params,
  });
}

// 结束服务
export function EndService(params) {
  return http.request({
    url: '/spaOrder/endService',
    method: 'POST',
    params,
  });
}

// 取消
export function CancelPay(params) {
  return http.request({
    url: '/spaOrder/cancelPay',
    method: 'POST',
    params,
  });
}

// 导出
export function ExportOrder(params) {
  return http.request({
    url: '/spaOrder/export',
    method: 'POST',
    params,
  });
}

// 获取导出列表
export function ExportList(params) {
  return http.request({
    url: '/spaOrder/exportList',
    method: 'get',
    params,
  });
}

// 异常处理
export function Abnormal(params) {
  return http.request({
    url: '/spaOrder/abnormal',
    method: 'POST',
    params,
  });
}
