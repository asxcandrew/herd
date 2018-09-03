import { BaseResource } from './resource';
import { axios } from '../utils';

class MediaResource extends BaseResource {
  constructor(base) {
    super()
    this.base = base;
    this.api = axios.public;    
  };

  createMedia(contentType){
    return this.post({content_type: contentType});
  };

  uploadFile(url, file){
    return this.api.put(
      url,
      file,
      {
        headers: {
          'Content-Type': file.type,
        }
      }
    )
  };
}

export default new MediaResource('media');
