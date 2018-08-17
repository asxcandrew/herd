<template>
  <div>
    <h2>{{ $t('views.profile.stories.title') }}</h2>
    <el-tabs v-model="activeTab">
      <el-tab-pane
        :label="$t('views.profile.stories.draftsTabTitle', { num: draftStories.length})"
        name="drafts">
        <el-row>
          <story-list-item
            v-for="story in draftStories"
            :key="story.id"
            :story=story
          />
        </el-row>
      </el-tab-pane>
      <el-tab-pane
        :label="$t('views.profile.stories.publishedTabTitle', { num: publishedStories.length})"
        name="published">
        <el-row>
          <story-list-item
            v-for="story in publishedStories"
            :key="story.id"
            :story=story
          />
        </el-row>
        </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { FETCH_STORIES } from '@/store/actions.type';
import { mapGetters } from 'vuex';
import storysListItem from '@/components/story-list-item';

export default {
  name: 'stories',
  components: { 'story-list-item': storysListItem },
  computed: {
    ...mapGetters(['session', 'stories']),
    draftStories: function() {
      return this.stories.filter((s) => {
        return !s.active
      });
    },
    publishedStories: function() {
      return this.stories.filter((s) => {
        return s.active
      });
    },
  },
  data(){
    return {
      activeTab: 'drafts',
    }
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$store.dispatch(FETCH_STORIES, {filter: { user_id: vm.session.user.id }});
    })
  },
};
</script>

<style>

</style>
