import { http } from '@/utils/http/axios';

// 获取预订单列表
export function List(params) {
  return http.request({
    url: '/cabinetOrder/list',
    method: 'get',
    params,
  });
}

// 获取预订单指定详情
export function View(params) {
  return http.request({
    url: '/cabinetOrder/view',
    method: 'GET',
    params,
  });
}

// 导出
export function ExportOrder(params) {
  return http.request({
    url: '/cabinetOrder/export',
    method: 'POST',
    params,
  });
}

// 获取导出列表
export function ExportList(params) {
  return http.request({
    url: '/cabinetOrder/exportList',
    method: 'get',
    params,
  });
}

// 订单完成
export function OrderComplete(params) {
  return http.request({
    url: '/cabinetOrder/complete',
    method: 'POST',
    params,
  });
}

// 订单退款
export function RefundOrder(params) {
  return http.request({
    url: '/cabinetOrder/refund',
    method: 'POST',
    params,
  });
}
