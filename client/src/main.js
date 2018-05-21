import Vue from 'vue';
import Element from 'element-ui';
import VeeValidate from 'vee-validate';
import { sync } from 'vuex-router-sync';
import App from './App';
import router from './router';
import store from './store';

import './assets/styles/main.scss';

sync(store, router, { moduleName: 'route' });

Vue.use(Element);
Vue.use(VeeValidate);

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>',
});
