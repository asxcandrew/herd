import Vue from 'vue';
import Element from 'element-ui';
import VueHead from 'vue-head';
import { sync } from 'vuex-router-sync';
import App from '@/App';
import store from '@/store';
import router from '@/router';
import i18n from '@/lang';

sync(store, router, { moduleName: 'route' });

Vue.use(Element);
Vue.use(VueHead);

import defaultLayout from '@/views/layout';
import storyNavbarPlugin from '@/components/story-navbar-plugin';
import searchNavbarPlugin from '@/components/search-navbar-plugin';

Vue.component('default-layout', defaultLayout);
Vue.component('story-navbar-plugin', storyNavbarPlugin);
Vue.component('search-navbar-plugin', searchNavbarPlugin);

import "normalize.css";
import "element-ui/lib/theme-chalk/index.css";
import "element-ui/lib/theme-chalk/display.css";

import '@/assets/styles/main.scss';

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  i18n,
  render: h => h(App)
});
