import { http, jumpExport } from '@/utils/http/axios';

// 获取会员意向表列表
export function List(params) {
  return http.request({
    url: '/pmsMemberIntention/list',
    method: 'get',
    params,
  });
}

// 获取会员意向表指定详情
export function View(params) {
  return http.request({
    url: '/pmsMemberIntention/view',
    method: 'GET',
    params,
  });
}

// 导出会员意向表
export function Export(params) {
  jumpExport('/pmsMemberIntention/export', params);
}


