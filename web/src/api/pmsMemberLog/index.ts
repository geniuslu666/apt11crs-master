import { http, jumpExport } from '@/utils/http/axios';

// 获取会员登录日志列表
export function List(params) {
  return http.request({
    url: '/pmsMemberLog/list',
    method: 'get',
    params,
  });
}

// 获取会员登录日志指定详情
export function View(params) {
  return http.request({
    url: '/pmsMemberLog/view',
    method: 'GET',
    params,
  });
}

// 导出会员登录日志
export function Export(params) {
  jumpExport('/pmsMemberLog/export', params);
}