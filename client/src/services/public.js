import { PublicResource } from './resource';

const usernameAvailability = (username) => {
  const url = `/username_availability/${username}`;
  return this.a.api.get(url, {});
};

const signup = (data) => {
  const url = '/sign_up';
  return this.a.api.post(url, data);
};

const signin = (data) => {
  const url = '/sign_in';
  return this.a.api.post(url, data);
};

const res = new PublicResource('', { usernameAvailability, signup, signin });

export default res;
