import { StoryService } from '@/services';
import {
  FETCH_STORIES,
  GET_STORY,
  CREATE_STORY,
  UPDATE_STORY,
  DELETE_STORY,
  GET_STORY_BODY,
} from '../actions.type'
import {
  SET_STORIES,
  ADD_STORY,
  REMOVE_STORY,
  ADD_BODY_STORY,
} from '../mutations.type'

const initState = {
  stories: []
};

const state = Object.assign({}, initState)

const getters = {
  stories: (state) => {
    return state.stories
  },
  getStoryById: (state) => (id) => {
    return state.stories.find(story => story.id == id)
  }
};

const mutations = {
  [SET_STORIES] (state, stories) {
    state.stories = stories
  },
  [ADD_STORY] (state, story) {
    state.stories.push(story)
  },
  [REMOVE_STORY] (state, id) {
    var index = state.stories.findIndex(item => item.id === id);
    state.stories.splice(index, 1);
  },
  [ADD_BODY_STORY] (state, data) {
    let story = state.stories.find(item => item.id == data.id)
    story.html_body = data.html_body
  }
};

const actions = {
  [FETCH_STORIES] ({ commit }, params) {
    return StoryService.query(params.filter)
      .then(({ data }) => {
        commit(SET_STORIES, data.data)
        return data
      })
  },
  [GET_STORY] ({ commit }, id) {
    if (state.stories.findIndex(story => story.id === id) != -1) {
      return (async () => {})();
    }
    return StoryService.get(id)
      .then(({ data }) => {
        commit(ADD_STORY, data.data)
        return data
      })
  },
  [GET_STORY_BODY] ({ commit }, id) {
    return StoryService.getBody(id)
      .then(({ data }) => {
        commit(ADD_BODY_STORY, data.data)
        return data
      })
  },
  [CREATE_STORY] ({ commit }, params) {
    return StoryService.post(params)
      .then(({ data }) => {
        commit(ADD_STORY, data.data)
        return data
      })
  },
  [UPDATE_STORY] ({ commit }, {id, params}) {
    return StoryService.put(id, params)
      .then(({ data }) => {
        commit(REMOVE_STORY, id)
        commit(ADD_STORY, data.data)
        return data
      })
  },
  [DELETE_STORY] ({ commit }, id) {
    return StoryService.delete(id)
      .then(({ status, data }) => {
        if (status == 200) {
          commit(REMOVE_STORY, id)
        } else {
          //show error
        }
        return data
      })
  },
};

export default { state, getters, mutations, actions };
