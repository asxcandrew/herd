class BaseResource {
  query(params) {
    let url = `/${this.base}`;
    return this.api.get(url, params);
  }
  get(id, options) {
    let url = `/${this.base}`;
    if (id !== undefined) {
      url += `/${id}`;
    }
    return this.api.get(url, options);
  }
  delete(id, options) {
    const url = `/${this.base}/${id}`;
    return this.api.delete(url, options);
  }
  post(options) {
    const url = `/${this.base}`;
    return this.api.post(url, options);
  }
  put(id, options) {
    const url = `/${this.base}/${id}`;
    return this.api.put(url, options);
  }
  patch(id, options) {
    const url = `/${this.base}/${id}`;
    return this.api.patch(url, options);
  }
}

export { BaseResource };
