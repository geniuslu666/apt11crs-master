import { http, jumpExport } from '@/utils/http/axios';

// 设置物业相册封面
// http://192.168.10.91:8000/admin/pmsProperty/setGalleryCover
export function setGalleryCover(params) {
  return http.request({
    url: '/pmsProperty/setGalleryCover',
    method: 'post',
    params,
  });
}
// 获取相册以及封面图
// http://192.168.10.91:8000/admin/pmsProperty/getGallery
export function getGallery(params) {
  return http.request({
    url: '/pmsProperty/getGallery',
    method: 'post',
    params,
  });
}
//dru2mp.natappfree.cc/admin/pms/editPropertyLanguage
// 编辑物业多语言属性信息
export function editPropertyLanguage(params) {
  return http.request({
    url: '/pms/editPropertyLanguage',
    method: 'post',
    params,
  });
}
// 获取物业列表
export function List(params) {
  return http.request({
    url: '/pmsProperty/list',
    method: 'get',
    params,
  });
}

// 获取所有物业列表
export function All(params) {
  return http.request({
    url: '/pmsProperty/all',
    method: 'get',
    params,
  });
}

// 删除/批量删除物业
export function Delete(params) {
  return http.request({
    url: '/pmsProperty/delete',
    method: 'POST',
    params,
  });
}

// 恢复物业
export function Recycle(params) {
  return http.request({
    url: '/pmsProperty/recycle',
    method: 'POST',
    params,
  });
}

// 添加/编辑物业
export function Edit(params) {
  return http.request({
    url: '/pmsProperty/edit',
    method: 'POST',
    params,
  });
}

// 物业绑定会员分组
export function EditGroup(params) {
  return http.request({
    url: '/pmsProperty/editGroup',
    method: 'POST',
    params,
  });
}

// 获取物业指定详情
export function View(params) {
  return http.request({
    url: '/pmsProperty/view',
    method: 'GET',
    params,
  });
}

// 导出物业
export function Export(params) {
  jumpExport('/pmsProperty/export', params);
}



// 导出物业 http://10.8.1.58:8000/admin/kefu/oss
export function kefuoss(params) {
  return http.request({
    url: '/kefu/oss',
    method: 'GET',
    params,
  });
}

// 物业排序
export function Sort(params) {
  return http.request({
    url: '/pmsProperty/sortUpdate',
    method: 'POST',
    params,
  });
}

// 按摩-预定状态开启关闭
export function SwitchSpaCanOrder(params) {
  return http.request({
    url: '/pmsProperty/switchSpaCanOrder',
    method: 'POST',
    params,
  });
}

// 按摩-门禁密码
export function UpdateAccessPass(params) {
  return http.request({
    url: '/pmsProperty/updateAccessPass',
    method: 'POST',
    params,
  });
}
