import Vue from 'vue';
import Element from 'element-ui';
import { sync } from 'vuex-router-sync';
import App from './App';
import store from './store';
import router from './router';
import i18n from './lang';

import './assets/styles/main.scss';

import defaultLayout from './views/layout';

sync(store, router, { moduleName: 'route' });

Vue.use(Element);
Vue.component('default-layout', defaultLayout);

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  i18n,
  render: h => h(App)
});
