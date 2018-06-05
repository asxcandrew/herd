import { storage, axios } from '../utils';

const CHANGE_SESSION = (state, session) => {
  if (session && session.token) {
    // change axios authorization header
    axios.authorized.defaults.headers.Authorization = `Bearer ${session.token}`;
  }
  // TODO: new session mixin
  Object.assign(state.session, session);
  storage.set('wedn_net_session_info', state.session);
};

const DESTROY_SESSION = (state) => {
  axios.authorized.defaults.headers.Authorization = '';
  state.session = {};
  storage.set('wedn_net_session_info', state.session);
};

const CHANGE_MODALS = (state, modals) => {
  Object.assign(state.modals, modals);
};

export { CHANGE_SESSION, CHANGE_MODALS, DESTROY_SESSION };
