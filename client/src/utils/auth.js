import store from '@/store';

export default {
  authHeader() {
    return `Bearer ${store.getters.session.token}`;
  },
};
