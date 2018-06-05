import { PublicService } from '../services';
import defaults from './defaults';

const signIn = ({ commit, dispatch }, form) => {
  return PublicService.signin({
    username: form.email.trim(),
    password: form.password.trim(),
  }).then(res => {
    commit('CHANGE_SESSION', { token: res.data.token, token_expiration: res.data.expire });
    return dispatch('getCurrentUser');
  });
};

const signUp = ({ commit }, form ) => {
  return PublicService.signup(form).then((res) => {
    commit('CHANGE_SESSION', { token: res.data.token, token_expiration: res.data.expire });
  }).catch(() => {
    resolve(false);
  });
};

const signOut = ({ commit }) => {
  return new Promise((resolve, reject) => {
    commit('DESTROY_SESSION');
  }).catch(() => {
    resolve(false);
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

const showModal = ({ commit }, name) => {
  commit('CHANGE_MODALS', Object.assign(defaults().modals, { [name]: true }));
};

const closeModal = ({ commit }, name) => {
  commit('CHANGE_MODALS', { [name]: false });
};

export { signOut, showModal, closeModal, signUp, signIn };
