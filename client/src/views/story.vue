<template>
  <el-row
    v-if="isStoryReady"
    justify="center"
  >
    <el-col :span="24">
      <author-card :author="story.user" :story="story" />
    </el-col>
    <el-col :span="24" class="public-story-container">
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
import authorCard from '@/components/author-card';
import storyMixin from '@/mixins/story-utils';

export default {
  name: 'show-story',
  mixins: [storyMixin],
  components: {
    authorCard,
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
</style>
