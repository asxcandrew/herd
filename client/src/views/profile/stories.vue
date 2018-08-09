<template>
  <div>
    <h2>{{ $t('views.profile.stories.title') }}</h2>
    <el-tabs v-model="activeTab">
      <el-tab-pane
        :label="this.$t('views.profile.stories.draftsTabTitle', { mun: draftStories.length})"
        name="drafts">
        <p v-for="story in draftStories">
          {{ story.title }}
        </p>
      </el-tab-pane>
      <el-tab-pane
        :label="this.$t('views.profile.stories.publishedTabTitle', { mun: publishedStories.length})"
        name="published">
        <p v-for="story in publishedStories">
          {{ story.title }}
        </p></el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { FETCH_STORIES } from '@/store/actions.type';
import { mapGetters } from 'vuex';

export default {
  name: 'stories',
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
  mounted() {
    this.$store.dispatch(FETCH_STORIES, {filter: { user_id: this.session.user.id }});
  }
};
</script>

<style>

</style>
