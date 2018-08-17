import defaults from './defaults';

const state = {
  modals: defaults().modals,
  statusBar: defaults().statusBar,
};

// if (state.session && state.session.token) {
//   // init axios headers
//   axios.defaults.headers.Authorization = `Bearer ${state.session.token}`
// }

export default state;
