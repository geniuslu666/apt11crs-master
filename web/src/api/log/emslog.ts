import { http } from '@/utils/http/axios';

export function getLogList(params) {
  return http.request({
    url: '/emsLog/list',
    method: 'get',
    params,
  });
}
export function Delete(params) {
  return http.request({
    url: '/emsLog/delete',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/emsLog/view',
    method: 'GET',
    params,
  });
}
