import { axios } from '../utils';

function PublicResource(base, actions) {
  this.base = base;
  this.api = axios.public;
  Object.assign(this, actions);
}

function AuthorizedResource(base, actions) {
  this.base = base;
  this.api = axios.authorized;
  Object.assign(this, actions);
}

const proto = {
  get(id, options) {
    let url = `/${this.base}`;
    if (id !== undefined) {
      url += `/${id}`;
    }
    return this.api.get(url, options);
  },
  delete(id, options) {
    const url = `/${this.base}/${id}`;
    return this.api.delete(url, options);
  },
  post(options) {
    const url = `/${this.base}`;
    return this.api.post(url, options);
  },
  put(id, options) {
    const url = `/${this.base}/${id}`;
    return this.api.put(url, options);
  },
  patch(id, options) {
    const url = `/${this.base}/${id}`;
    return this.api.patch(url, options);
  },
};

AuthorizedResource.prototype = Object.create(proto, {});
PublicResource.prototype = Object.create(proto, {});

export { PublicResource, AuthorizedResource };
