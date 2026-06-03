import { http, jumpExport } from '@/utils/http/axios';

// 获取退款流水列表
export function List(params) {
  return http.request({
    url: '/pmsTransactionRefund/list',
    method: 'get',
    params,
  });
}

// 获取退款流水指定详情
export function View(params) {
  return http.request({
    url: '/pmsTransactionRefund/view',
    method: 'GET',
    params,
  });
}

// 导出退款流水
export function Export(params) {
  jumpExport('/pmsTransactionRefund/export', params);
}


