const CHANGE_STATUS_BAR = (state, params) => {
  Object.assign(state.statusBar, params);
};

const CHANGE_MODALS = (state, modals) => {
  Object.assign(state.modals, modals);
};

export { CHANGE_MODALS, CHANGE_STATUS_BAR };
