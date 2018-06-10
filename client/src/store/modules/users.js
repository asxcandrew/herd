import { UserService } from '../../services';

const state = {

};

/**
 * Getters
 * @type {Object}
 */
const getters = {

};

/**
 * Mutations
 * @type {Object}
 */
const mutations = {

};

/**
 * Actions
 * @type {Object}
 */
const actions = {
  getCurrentUser({ commit }) {
    return UserService.get('me')
      .then((res) => {
        commit('CHANGE_SESSION', { user: res.data.data });
      });
  },
};

// Export module
export default { state, getters, mutations, actions };
