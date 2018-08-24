import { BaseResource } from './resource';
import { axios } from '../utils';

class PublicResource extends BaseResource {
  constructor(base) {
    super()
    this.base = base;
    this.api = axios.public;    
  };

  usernameAvailability(username){
    const url = `/username_availability/${username}`;
    return this.api.get(url, {});
  };

  feed(){
    const url = `/feed`;
    return this.api.get(url, {});
  };
  
  signup(data){
    const url = '/sign_up';
    return this.api.post(url, data);
  };

  signin(data){
    const url = '/sign_in';
    return this.api.post(url, data);
  };

}

const res = new PublicResource('');

export default res;
