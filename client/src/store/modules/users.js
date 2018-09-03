import { UsersService } from '@/services';
import {
  ADD_USER,
} from '../mutations.type'
import {
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
  [GET_USER]({ commit }, username) {
    return UsersService.get(username)
      .then((res) => {
        commit(ADD_USER, res.data.data);
      });
  },
};

// Export module
export default { state, getters, mutations, actions };
