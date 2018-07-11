import Vue from 'vue';
import Router from 'vue-router';
import { auth } from '../utils';
import routes from './routes';

Vue.use(Router);

const router = new Router({
  base: '/',
  mode: 'history',
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!auth.loggedIn()) {
      next({
        path: '/',
      });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
