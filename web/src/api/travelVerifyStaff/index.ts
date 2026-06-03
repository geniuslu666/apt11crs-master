import { http } from '@/utils/http/axios';

export function List(params) {
 return http.request({ url: '/travel/verifyStaff/list', method: 'GET', params });
}

export function View(params) {
 return http.request({ url: '/travel/verifyStaff/view', method: 'GET', params });
}

export function Edit(params) {
 return http.request({ url: '/travel/verifyStaff/edit', method: 'POST', params });
}

export function Delete(params) {
 return http.request({ url: '/travel/verifyStaff/delete', method: 'POST', params });
}

export function ScopeOptions() {
 return http.request({ url: '/travel/verifyStaff/scopeOptions', method: 'GET' });
}

export function Status(params) {
 return http.request({ url: '/travel/verifyStaff/status', method: 'POST', params });
}
