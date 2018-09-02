import { PublicService } from '@/services';
import {
  FETCH_FEED,
  GET_FEED_BODY,
  GET_FEED_STORY,
} from '../actions.type'
import {
  SET_FEED,
  ADD_FEED_STORY,
  ADD_FEED_BODY,
} from '../mutations.type'

const initState = {
  feed: []
};

const state = Object.assign({}, initState)

const mutations = {
  [SET_FEED] (state, stories) {
    state.feed = stories
  },
  [ADD_FEED_STORY] (state, story) {
    state.feed.push(story)
  },
  [ADD_FEED_BODY] (state, data) {
    let story = state.feed.find(item => item.uid == data.uid)
    story.html_body = data.html_body
  }
};

const getters = {
  feed: (state) => {
    return state.feed
  },
  getFeedStoryByUid: (state) => (uid) => {
    return state.feed.find(story => story.uid == uid)
  }
};

const actions = {
  [GET_FEED_STORY] ({ commit }, uid) {
    return PublicService.feedStory(uid)
      .then(({ data }) => {
        commit(ADD_FEED_STORY, data.data)
        return data.data
      })
  },
  [FETCH_FEED] ({ commit }, params) {
    return PublicService.feed()
      .then(({ data }) => {
        commit(SET_FEED, data.data)
        return data
      })
  },
  [GET_FEED_BODY] ({ commit }, uid) {
    return PublicService.feedStoryBody(uid)
      .then(({ data }) => {
        commit(ADD_FEED_BODY, data.data)
        return data
      })
  },
}

export default { state, actions, getters, mutations };
