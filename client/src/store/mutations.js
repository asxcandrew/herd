import { storage } from '../utils';

const CHANGE_SESSION = (state, session) => {
  // TODO: new session mixin
  Object.assign(state.session, session);
  storage.set('session_info', state.session);
};

const CHANGE_STATUS_BAR = (state, params) => {
  Object.assign(state.statusBar, params);
};

const DESTROY_SESSION = (state) => {
  state.session = {};
  storage.set('session_info', state.session);
};

const CHANGE_MODALS = (state, modals) => {
  Object.assign(state.modals, modals);
};

export { CHANGE_SESSION, CHANGE_MODALS, DESTROY_SESSION, CHANGE_STATUS_BAR };
