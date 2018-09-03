<template>
  <el-row>
    <el-col
      :lg="12"
      :sm="24"
      v-for="story in feed"
      :key="story.id"
    >
      <el-card shadow="disabled">
        <img src="" class="image">
        <div>
          <h3>
            <router-link
              :to="{ name: 'show-story', params: { uri: storyUri(story) }}"
              class="unstyled-link"
            >{{ story.title }}</router-link>
          </h3>
          <div class="bottom clearfix">
            <author-card :author="story.user" :story="story" :small="true" />
            <!-- <time class="substring">{{ formatDate(story.publised_at) }}</time> -->
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import { FETCH_FEED } from '@/store/actions.type';
import { mapGetters } from 'vuex';
import storyMixin from '@/mixins/story-utils';
import authorCard from '@/components/author-card';

export default {
  name: 'feed',
  mixins: [ storyMixin ],
  components: {
    authorCard
  },
  computed: {
    ...mapGetters(['feed']),
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$store.dispatch(FETCH_FEED);
    })
  },
};
</script>
<style>
.el-card {
  border: 0;
}
</style>
