import Vue from 'vue';
import Router from 'vue-router';
import routes from './routes';
import store from '@/store';

import {
  CLEAR_STATUS_BAR,
} from '@/store/actions.type';

Vue.use(Router);

const router = new Router({
  base: '/',
  mode: 'history',
  routes,
});

router.beforeEach((to, from, next) => {
  store.dispatch(CLEAR_STATUS_BAR);

  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.getters.loggedIn) {
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
