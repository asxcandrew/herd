import Vue from 'vue';
import Element from 'element-ui';
import { sync } from 'vuex-router-sync';
import App from './App';
import store from './store';
import router from './router';
import i18n from './lang';

import './assets/styles/main.scss';

sync(store, router, { moduleName: 'route' });

Vue.use(Element);

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  i18n,
  components: { App },
  template: '<App/>',
});
