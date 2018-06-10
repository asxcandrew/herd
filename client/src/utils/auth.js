import store from '../store';

export default {
  // TODO: renew token here
  loggedIn() {
    if (store.getters.session.token) {
      if (Date.parse(store.getters.session.token_expiration) > Date.now()) {
        // Get new token
        return false;
      }
      return true;
    }
    return false;
  },
  authHeader() {
    return `Bearer ${store.getters.session.token}`;
  },
};
