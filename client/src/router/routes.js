
import profileRoutes from './profile-routes';
import notFound from '../views/404';

const baseRoutes = [
  {
    path: '/',
    meta: { requiresAuth: false },
    component: require('@/views/feed').default,
  },
  {
    path: '/story/:uri',
    name: 'show-story',
    meta: { requiresAuth: false },
    component: require('@/views/story').default,
  },
  {
    path: '/@:username',
    name: 'public-profile',
    component: require('@/views/public-profile').default,
  },
  {
    name: 'not-found',
    path: '*',
    component: notFound,
  },
];

export default baseRoutes.concat(profileRoutes);
