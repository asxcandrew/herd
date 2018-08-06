import { PublicService } from '@/services';
import defaults from './defaults';
import {
  UPDATE_STATUS_BAR,
  CLEAR_STATUS_BAR,
  SIGN_IN,
  SIGN_UP,
  SIGN_OUT,
  SHOW_MODAL,
  CLOSE_MODAL,
} from './actions.type';

export default {
  [SIGN_IN] ({ commit, dispatch }, form) {
    return PublicService.signin({
      username: form.email.trim(),
      password: form.password.trim(),
    }).then(res => {
      commit('CHANGE_SESSION', { token: res.data.token, token_expiration: Date.parse(res.data.expire) });
      return dispatch('getCurrentUser');
    });
  },
  [SIGN_UP] ({ commit }, form ) {
    return PublicService.signup(form).then((res) => {
      commit('CHANGE_SESSION', { token: res.data.token, token_expiration: res.data.expire });
    }).catch(() => {
      resolve(false);
    });
  },
  [SIGN_OUT] ({ commit }) {
    return (async () => {
      commit('DESTROY_SESSION');
    })();
  },
  [UPDATE_STATUS_BAR] ({ commit }, params) {
    commit('CHANGE_STATUS_BAR', params);
  },
  [CLEAR_STATUS_BAR] ({ commit }) {
    commit('CHANGE_STATUS_BAR', defaults().statusBar);
  },
  [SHOW_MODAL] ({ commit }, name) {
    commit('CHANGE_MODALS', Object.assign(defaults().modals, { [name]: true }));
  },
  [CLOSE_MODAL] ({ commit }, name) {
    commit('CHANGE_MODALS', { [name]: false });
  },
};
