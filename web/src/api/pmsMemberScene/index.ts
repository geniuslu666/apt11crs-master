import { http } from '@/utils/http/axios';

// 获取会员等级场景列表
export function List(params) {
  return http.request({
    url: '/pmsMemberScene/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除会员等级场景
export function Delete(params) {
  return http.request({
    url: '/pmsMemberScene/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑会员等级场景
export function Edit(params) {
  return http.request({
    url: '/pmsMemberScene/edit',
    method: 'POST',
    params,
  });
}

// 操作会员等级场景开关
export function Switch(params) {
  return http.request({
    url: '/pmsMemberScene/switch',
    method: 'POST',
    params,
  });
}

// 获取会员等级场景指定详情
export function View(params) {
  return http.request({
    url: '/pmsMemberScene/view',
    method: 'GET',
    params,
  });
}


