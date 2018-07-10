
import profileRoutes from './profile-routes';
import Layout from '../views/layout';

export default [
  {
    path: '/',
    meta: { requiresAuth: false },
    component: Layout,
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
