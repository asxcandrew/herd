<template>
  <div class="navbar-plugin story-plugin">
    <el-button
      v-if="story"
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
  UPDATE_STORY,
} from '@/store/actions.type';

export default {
  name: 'story-navbar-plugin',
  computed: {
    story(){
      return this.$store.getters.getStoryById(this.$route.params.id)
    },
  },
  methods: {
    onClickPublish() {
      this.$store.dispatch(
        UPDATE_STORY,
        {
          id: this.story.id,
          params: { active: true, published_at: Date.now() },
        },
      );
    }
  },
}
</script>
