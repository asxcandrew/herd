
// import mainRoutes from './main-routes'

export default [
  // ## login page
  {
    name: 'login',
    path: '/login',
    meta: { requiresAuth: false },
    component: () => import('../views/login'),
  },
  // ## main page
  {
    path: '/',
    meta: { requiresAuth: true },
    component: () => import(/* webpackChunkName: 'common' */ '../views/layout'),
    // children: mainRoutes
  },
  // ## not found page
  // {
  //   name: 'not-found',
  //   path: '*',
  //   meta: { requiresAuth: false },
  //   component: () => import(/* webpackChunkName: 'common' */ '../views/error')
  // }
];
