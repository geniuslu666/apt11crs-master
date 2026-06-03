import { http, jumpExport } from '@/utils/http/axios';

// 获取会员分组列表
export function List(params) {
  return http.request({
    url: '/pmsMemberGroup/list',
    method: 'get',
    params,
  });
}

// 获取会员分组全部列表
export function GroupAll(params) {
  return http.request({
    url: '/pmsMemberGroup/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除会员分组
export function Delete(params) {
  return http.request({
    url: '/pmsMemberGroup/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑会员分组
export function Edit(params) {
  return http.request({
    url: '/pmsMemberGroup/edit',
    method: 'POST',
    params,
  });
}

// 获取会员分组指定详情
export function View(params) {
  return http.request({
    url: '/pmsMemberGroup/view',
    method: 'GET',
    params,
  });
}

// 导出会员分组
export function Export(params) {
  jumpExport('/pmsMemberGroup/export', params);
}


