import { http,jumpExport} from '@/utils/http/axios';

// 财务概况
export function Stat(params) {
  return http.request({
    url: '/finance/stat',
    method: 'get',
    params,
  });
}

export function List(params) {
  return http.request({
    url: '/finance/list',
    method: 'get',
    params,
  });
}

// 导出报表
export function Export(params) {
  jumpExport('/finance/export', params);
}


