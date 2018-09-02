import { StoryService, PublicService } from '@/services';
import {
  FETCH_STORIES,
  GET_STORY,
  CREATE_STORY,
  UPDATE_STORY,
  PUBLISH_STORY,
  HIDE_STORY,
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
  getStoryByUid: (state) => (uid) => {
    return state.stories.find(story => story.uid == uid)
  }
};

const mutations = {
  [SET_STORIES] (state, stories) {
    state.stories = stories
  },
  [ADD_STORY] (state, story) {
    state.stories.push(story)
  },
  [REMOVE_STORY] (state, uid) {
    var index = state.stories.findIndex(item => item.uid === uid);
    state.stories.splice(index, 1);
  },
  [ADD_BODY_STORY] (state, data) {
    let story = state.stories.find(item => item.uid == data.uid)
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
  [GET_STORY] ({ commit, getters }, uid) {
    if (getters.getStoryByUid(uid)) {
      return getters.getStoryByUid(uid);
    }
    return StoryService.get(uid)
      .then(({ data }) => {
        commit(ADD_STORY, data.data)
        return data.data
      })
  },
  [GET_STORY_BODY] ({ commit }, uid) {
    return StoryService.getBody(uid)
      .then(({ data }) => {
        commit(ADD_BODY_STORY, data.data)
        return data
      })
  },
  [CREATE_STORY] ({ commit }, params) {
    return StoryService.post(params)
      .then(({ data }) => {
        commit(ADD_STORY, data.data)
        return data.data
      })
  },
  [UPDATE_STORY] ({ commit }, {uid, params}) {
    return StoryService.put(uid, params)
      .then(({ data }) => {
        commit(REMOVE_STORY, uid)
        commit(ADD_STORY, data.data)
        return data
      })
  },
  [PUBLISH_STORY] ({ dispatch }, uid) {
    let date = new Date();
    return dispatch(
      UPDATE_STORY,
      {
        uid: uid,
        params: { active: 'true', published_at: date.toISOString() },
      },
    );
  },
  [HIDE_STORY] ({ dispatch }, uid) {
    return dispatch(
      UPDATE_STORY,
      {
        uid: uid,
        params: { active: 'false' },
      },
    );
  },
  [DELETE_STORY] ({ commit }, uid) {
    return StoryService.delete(uid)
      .then(({ status, data }) => {
        if (status == 200) {
          commit(REMOVE_STORY, uid)
        } else {
          //show error
        }
        return data
      })
  },
};

export default { state, getters, mutations, actions };
