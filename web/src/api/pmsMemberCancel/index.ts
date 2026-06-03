import { http } from '@/utils/http/axios';

// 获取会员注销列表
export function List(params) {
  return http.request({
    url: '/pmsMemberCancel/list',
    method: 'get',
    params,
  });
}

// 同意注销
export function Agree(params) {
  return http.request({
    url: '/pmsMemberCancel/agree',
    method: 'POST',
    params,
  });
}

// 拒绝注销
export function Disagree(params) {
  return http.request({
    url: '/pmsMemberCancel/disagree',
    method: 'POST',
    params,
  });
}
