import AuthorizedResource from './authorized';

class StoryService extends AuthorizedResource {
  getBody(id){
    return this.get(id, undefined, 'body');
  };
}

export default new StoryService('stories');
