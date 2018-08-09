import { StoryService } from '@/services';
import {
  FETCH_STORIES,
  GET_STORY,
  CREATE_STORY,
} from '../actions.type'
import {
  SET_STORIES,
  ADD_STORY
} from '../mutations.type'

const initState = {
  stories: []
};

const state = Object.assign({}, initState)

const getters = {
  stories (state) {
    return state.stories
  },
};

const mutations = {
  [SET_STORIES] (state, stories) {
    state.stories = stories
  },
  [ADD_STORY] (state, story) {
    state.stories.push(story)
  },
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
    // TODO: avoid extronuous network call if story exists
    return StoryService.get(id)
      .then(({ data }) => {
        commit(ADD_STORY, data.data)
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
};

export default { state, getters, mutations, actions };
