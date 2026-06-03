import { http } from '@/utils/http/axios';

// 获取会员等级列表
export function List(params) {
  return http.request({
    url: '/pmsMemberLevel/list',
    method: 'get',
    params,
  });
}

// 获取会员等级全部列表
export function All(params) {
  return http.request({
    url: '/pmsMemberLevel/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除会员等级
export function Delete(params) {
  return http.request({
    url: '/pmsMemberLevel/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑会员等级
export function Edit(params) {
  return http.request({
    url: '/pmsMemberLevel/edit',
    method: 'POST',
    params,
  });
}

// 获取会员等级指定详情
export function View(params) {
  return http.request({
    url: '/pmsMemberLevel/view',
    method: 'GET',
    params,
  });
}


// 修改渠道管理状态
export function Status(params) {
  return http.request({
    url: '/pmsMemberLevel/status',
    method: 'POST',
    params,
  });
}
