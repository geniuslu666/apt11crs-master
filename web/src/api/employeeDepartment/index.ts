import { http } from '@/utils/http/axios';

/**
 * @description: 获取员工部门列表
 */
export function List(params?: any) {
  return http.request({
    url: '/employeeDepartment/list',
    method: 'get',
    params,
  });
}

/**
 * @description: 获取员工部门详情
 */
export function View(params: { id: number }) {
  return http.request({
    url: '/employeeDepartment/view',
    method: 'get',
    params,
  });
}

/**
 * @description: 新增/编辑员工部门
 */
export function Edit(params: any) {
  return http.request({
    url: '/employeeDepartment/edit',
    method: 'post',
    data: params,
  });
}

/**
 * @description: 删除员工部门
 */
export function Delete(params: { id: number }) {
  return http.request({
    url: '/employeeDepartment/delete',
    method: 'post',
    data: params,
  });
}

/**
 * @description: 批量删除员工部门
 */
export function BatchDelete(params: { ids: number[] }) {
  return http.request({
    url: '/employeeDepartment/batchDelete',
    method: 'post',
    data: params,
  });
}

/**
 * @description: 切换员工部门状态
 */
export function Switch(params: { id: number; status: number }) {
  return http.request({
    url: '/employeeDepartment/switch',
    method: 'post',
    data: params,
  });
}

/**
 * @description: 获取部门树结构
 */
export function GetDepartmentTree(params?: any) {
  return http.request({
    url: '/employeeDepartment/tree',
    method: 'get',
    params,
  });
}
