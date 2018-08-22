<template>
  <div class="navbar-plugin story-plugin">
    <el-button
      v-if="isPublished"
      type="text"
      @click="onClickHide"
    >
      hide
    </el-button>
    <el-button
      v-else
      type="text"
      @click="onClickPublish"
    >
      publish
    </el-button>
  </div>
</template>

<script>
import {
  UPDATE_STATUS_BAR,
  PUBLISH_STORY,
  HIDE_STORY,
} from '@/store/actions.type';

export default {
  name: 'story-navbar-plugin',
  computed: {
    isPublished() {
      return this.story && this.story.active;
    },
    story(){
      return this.$store.getters.getStoryById(this.$route.params.id);
    },
  },
  methods: {
    onClickHide() {
      this.$store.dispatch(HIDE_STORY, this.story.id);
    },
    onClickPublish() {
      this.$store.dispatch(PUBLISH_STORY, this.story.id).then(() => {
        this.$store.dispatch(UPDATE_STATUS_BAR, { notice: 'published' });
      });
    }
  },
}
</script>
