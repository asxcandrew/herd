import { storage } from '../utils';

const state = {
  session: storage.get('wedn_net_session_info') || {},
};

// if (state.session && state.session.token) {
//   // init axios headers
//   axios.defaults.headers.Authorization = `Bearer ${state.session.token}`
// }

export default state;
