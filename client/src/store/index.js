import Vue from 'vue';
import Vuex from 'vuex';

import state from './state';
import * as getters from './getters';
import * as mutations from './mutations';
import actions from './actions';
import modules from './modules';

Vue.use(Vuex);

const strict = process.env.NODE_ENV !== 'production';

const plugins = [];

const store = new Vuex.Store({ state, getters, mutations, actions, modules, strict, plugins });

export default store;
