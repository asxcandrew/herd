import axios from 'axios';
import auth from './auth';

const _rootApi = process.env.API_BASE;

const createAxios = (url)=> {
  return axios.create({
    baseURL: _rootApi + url,
    responseType: 'json',
    timeout: 5000,
    headers: {
      // 'X-Custom-Header': 'foobar',
      // // true: need, false: dont need
      // 'Authorization': true,
      // 'X-Requested-With': 'XMLHttpRequest'
    },
  });
};

const instance = {
  public: createAxios('/v1/p'),
  authorized: createAxios('/v1/a'),
};

instance.authorized.interceptors.request.use((config) => {
  config.headers.Authorization = auth.authHeader();

  return config;
});

export default instance;
