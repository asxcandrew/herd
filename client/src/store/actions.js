import defaults from './defaults';
import {
  UPDATE_STATUS_BAR,
  CLEAR_STATUS_BAR,
  SHOW_MODAL,
  CLOSE_MODAL,
} from './actions.type';

export default {
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
