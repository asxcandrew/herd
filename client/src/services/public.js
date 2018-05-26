import { PublicResource } from './resource';

const usernameAvailability = (username) => {
  const url = `/username_availability/${username}`;
  return this.a.api.get(url, {});
};

const res = new PublicResource('', { usernameAvailability });

export default res;
