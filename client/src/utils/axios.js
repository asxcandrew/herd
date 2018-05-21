import axios from 'axios';
// import storage from './storage'

const instance = axios.create({
  baseURL: process.env.BASE_URL,
  timeout: 5000,
  headers: {
    // 'X-Custom-Header': 'foobar',
    // // true: need, false: dont need
    // 'Authorization': true,
    // 'X-Requested-With': 'XMLHttpRequest'
  },
});

// instance.interceptors.request.use(config => {
//   // Add authorization in the header
//   // TODO: token in store
//   const token = storage.get('wedn_net_access_token')
//   if (token && config.headers.Authorization) {
//     config.headers.Authorization = `Bearer ${token}`
//   }
//   return config
// })

export default instance;
