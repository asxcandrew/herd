import { BaseResource } from './resource';
import { axios } from '../utils';

class AuthorizedResource extends BaseResource {
  constructor (base) {
    super()
    this.api = axios.authorized;
    this.base = base;
  }
};

export default AuthorizedResource;
