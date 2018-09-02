<template>
  <el-row
    v-if="isStoryReady"
    justify="space-around"
  >
    <el-col :span="20" class="author-card__wrapper">
      <div class="author-card__avatar">
        <router-link
          :to="{ name: 'public-profile', params: { username: story.user.username }}"
          class="unstyled-link">
          <avatar
            :username="story.user.name"
            :size="60"
          />
        </router-link>
      </div>
      <div class="author-card__body">
        {{ story.user.name }}
        <el-button size="mini" plain>Follow</el-button><br>
        <ul class="ul-unstyled">
          <li class="substring">{{ formatDate(story.publised_at) }}</li>
          <li class="substring">{{ countWords(story) }}</li>
        </ul>
      </div>
    </el-col>
    <el-col :span="20" class="public-story-container">
      <h2>{{ story.title }}</h2>
      <div v-html="story.html_body"></div>
    </el-col>
  </el-row>
</template>

<script>
import {
  GET_FEED_STORY,
  GET_FEED_BODY,
} from '@/store/actions.type';
import storyMixin from '@/mixins/story-utils';
import Avatar from 'vue-avatar';

export default {
  name: 'show-story',
  mixins: [storyMixin],
  components: {
    Avatar
  },
  data() {
    return {
      isStoryReady: false
    }
  },
  computed: {
    story(){
      return this.$store.getters.getFeedStoryByUid(this.uidFromUri(this.$route.params.uri))
    },
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$store.dispatch(GET_FEED_STORY, vm.uidFromUri(vm.$route.params.uri)).then( story => {
        vm.$store.dispatch(GET_FEED_BODY, story.uid).then(_ => {
          vm.isStoryReady = true;
        }).then(_ => {
          vm.title = 'sdsdsd'
          vm.$emit('updateHead')
        })
      })
    })
  },
};
</script>
<style lang="scss">
.public-story-container{
  font-size: 15pt;
  line-height: 1.3;
}
.author-card__wrapper{
  display: flex;
  .author-card__avatar{
    margin-right: 10px;
  }
}

</style>
