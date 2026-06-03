import { http, jumpExport } from '@/utils/http/axios';

// 获取附件分类列表
export function List(params) {
  return http.request({
    url: '/sysAttachmentKind/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除附件分类
export function Delete(params) {
  return http.request({
    url: '/sysAttachmentKind/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑附件分类
export function Edit(params) {
  return http.request({
    url: '/sysAttachmentKind/edit',
    method: 'POST',
    params,
  });
}

// 获取附件分类指定详情
export function View(params) {
  return http.request({
    url: '/sysAttachmentKind/view',
    method: 'GET',
    params,
  });
}