import { UserService } from '@/services';
import {
  CHANGE_SESSION,
} from '../mutations.type'
import {
  GET_CURRENT_USER,
} from '../actions.type';

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
  [GET_CURRENT_USER]({ commit }) {
    return UserService.get('me')
      .then((res) => {
        commit(CHANGE_SESSION, { user: res.data.data });
      });
  },
};

// Export module
export default { state, getters, mutations, actions };
