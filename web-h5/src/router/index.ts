import { createRouter, createWebHashHistory } from 'vue-router'
import { useTravelStaffStore } from '../store/travelStaff'

const router = createRouter({
  history: createWebHashHistory('/daytrip/'),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/Login.vue'),
      meta: { public: true },
    },
    {
      path: '/',
      name: 'Index',
      component: () => import('../views/Index.vue'),
      meta: { requireAuth: true },
    },
    {
      path: '/verify',
      name: 'ScanCode',
      component: () => import('../views/Index.vue'),
      meta: { requireAuth: true },
    },
    {
      path: '/verify/confirm',
      name: 'VerifyConfirm',
      component: () => import('../views/VerifyConfirm.vue'),
      meta: { requireAuth: true },
    },
    {
      path: '/verify/result',
      name: 'VerifyResult',
      component: () => import('../views/VerifyResult.vue'),
      meta: { requireAuth: true },
    },
    {
      path: '/verify/list',
      name: 'VerifyList',
      component: () => import('../views/VerifyList.vue'),
      meta: { requireAuth: true },
    },
    {
      path: '/verify/detail/:id',
      name: 'VerifyDetail',
      component: () => import('../views/VerifyDetail.vue'),
      meta: { requireAuth: true },
    },
  ],
})

router.beforeEach((to) => {
  const store = useTravelStaffStore()
  if (to.meta.requireAuth && !store.isLoggedIn) {
    return { path: '/login', query: { redirect: to.fullPath } }
  }
  if (to.path === '/login' && store.isLoggedIn) {
    return { path: '/' }
  }
})

export default router
