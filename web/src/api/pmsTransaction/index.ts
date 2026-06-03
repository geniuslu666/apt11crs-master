import { http, jumpExport } from '@/utils/http/axios';

// 获取支付流水列表
export function List(params) {
  return http.request({
    url: '/pmsTransaction/list',
    method: 'get',
    params,
  });
}

// 获取支付流水指定详情
export function View(params) {
  return http.request({
    url: '/pmsTransaction/view',
    method: 'GET',
    params,
  });
}

// 导出支付流水
export function Export(params) {
  jumpExport('/pmsTransaction/export', params);
}


