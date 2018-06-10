
import profileRoutes from './profile-routes';

export default [
  {
    path: '/',
    meta: { requiresAuth: false },
    component: () => import(/* webpackChunkName: 'common' */ '../views/layout'),
    children: profileRoutes,
  },
  // ## not found page
  // {
  //   name: 'not-found',
  //   path: '*',
  //   meta: { requiresAuth: false },
  //   component: () => import(/* webpackChunkName: 'common' */ '../views/error')
  // }
];
