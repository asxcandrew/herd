import { UserService } from '@/services';
import {
  CHANGE_SESSION,
  ADD_USER,
} from '../mutations.type'
import {
  GET_CURRENT_USER,
  GET_USER,
} from '../actions.type';

const initState = {
  users: []
};

const state = Object.assign({}, initState)

const getters = {
  users: (state) => {
    return state.feed
  },
  getUserByUsername: (state) => (username) => {
    return state.users.find(user => users.username == username)
  }
};

const mutations = {
  [ADD_USER] (state, user) {
    state.users.push(user)
  },
};

const actions = {
  [GET_CURRENT_USER]({ commit }) {
    return UserService.get('me')
      .then((res) => {
        commit(CHANGE_SESSION, { user: res.data.data });
      });
  },
  [GET_USER]({ commit }, username) {
    return UserService.get(username)
      .then((res) => {
        commit(CHANGE_SESSION, { user: res.data.data });
      });
  },
};

// Export module
export default { state, getters, mutations, actions };
