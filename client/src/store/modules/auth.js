import { storage } from '@/utils';
import { PublicService } from '@/services';
import {
  SIGN_IN,
  SIGN_UP,
  SIGN_OUT,
  GET_CURRENT_USER,
} from '../actions.type';
import {
  DESTROY_SESSION,
  CHANGE_SESSION,
} from '../mutations.type'

const state = {
  session: storage.get('session_info') || {},
};

const getters = {
  session: (state) => {
    return state.session;
  },
  loggedIn: (state) => {
    return state.session && state.session.user && state.session.token_expiration > Date.now();
  }
};

const mutations = {
  [DESTROY_SESSION] (state) {
    state.session = {};
    storage.set('session_info', state.session);
  },
  [CHANGE_SESSION] (state, data) {
    state.session = Object.assign({}, state.session, data);
    storage.set('session_info', state.session);
  },
};

const actions = {
  [SIGN_IN] ({ commit, dispatch }, form) {
    return PublicService.signin({
      email: form.email.trim(),
      password: form.password.trim(),
    }).then(res => {
        commit(CHANGE_SESSION, { token: res.data.token, token_expiration: Date.parse(res.data.expire) });
    }).then(_ => {
      return dispatch(GET_CURRENT_USER);
    })
  },
  [SIGN_UP] ({ commit }, form ) {
    return PublicService.signup(form).then((res) => {
      commit(CHANGE_SESSION, { token: res.data.token, token_expiration: res.data.expire });
    }).catch(() => {
      resolve(false);
    });
  },
  [SIGN_OUT] ({ commit }) {
    return (async () => {
      commit(DESTROY_SESSION);
    })();
  },
};

export default { state, getters, mutations, actions };
