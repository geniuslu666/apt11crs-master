import { http } from '@/utils/http/axios';


// 修改/新增取消政策 https://crsdev.yeebok.cn/admin/pmsCancelRate/edit
export function pmsCancelRateedit(params) {
  return http.request({
    url: '/pmsCancelRate/edit',
    method: 'post',
    params,
  });
}
// 获取取消政策列表 https://crsdev.yeebok.cn/admin/pmsCancelRate/list
export function pmsCancelRatelist(params) {
  return http.request({
    url: '/pmsCancelRate/list',
    method: 'GET',
    params,
  });
}

//获取房态
export function roomStatus(params) {
  return http.request({
    url: '/pms/roomStatus',
    method: 'post',
    params,
  });
}
//PMS价格日历_列表 https://crsdev.yeebok.cn/admin/pms/priceCalendar
export function priceCalendar(params) {
  return http.request({
    url: '/pms/priceCalendar',
    method: 'post',
    params,
  });
}

export function dashboard(params) {
  return http.request({
    url: '/pms/dashboard',
    method: 'post',
    params,
  });
}

export function dashboardAll(params) {
  return http.request({
    url: '/pms/dashboard/all',
    method: 'post',
    params,
  });
}

// pms路由http://192.168.10.91:8000/admin/pms/menu
export function pmsmenu(params) {
  return http.request({
    url: '/pms/menu',
    method: 'post',
    params,
  });
}
