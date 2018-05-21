import { TokenService, UserService } from '../services';

const createToken = ({ commit }, { username, password }) => {
  return TokenService.post({
    username: username.trim(),
    password: password.trim(),
  }).then(res => {
    commit('CHANGE_SESSION', { token: res.data.token });
    return res.data.token;
  });
};

const checkToken = ({ commit, getters }) => {
  return new Promise((resolve, reject) => {
    // validate local store
    if (!getters.session.token) {
      return resolve(false);
    }
    // remote
    TokenService.get()
      .then(res => resolve(true))
      .catch(err => {
        commit('CHANGE_SESSION', { token: null });
        resolve(false);
      });
  });
};

const getCurrentUser = ({ commit }) => {
  return UserService.get('me')
    .then(res => {
      commit('CHANGE_SESSION', { user: res.data });
      return res.data;
    });
};

export { createToken, checkToken, getCurrentUser };
