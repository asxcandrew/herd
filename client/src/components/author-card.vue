<template>
  <div class="author-card__wrapper">
    <div class="author-card__avatar">
      <router-link
        :to="{ name: 'public-profile', params: { username: author.username }}"
        class="unstyled-link">
        <avatar
          v-if="!small"
          :username="author.name"
          :size="60"
        />
      </router-link>
    </div>
    <div class="author-card__body">
      {{ author.name }}
      <el-button size="mini" plain v-if="!small">Follow</el-button><br>
      <ul class="ul-unstyled">
        <li class="substring">{{ formatDate(story.publised_at) }}</li>
        <li class="substring" v-if="story.html_body">
          {{ countWords(story)}} words
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import storyMixin from '@/mixins/story-utils';
import Avatar from 'vue-avatar';

export default {
  name: 'author-card',
  mixins: [storyMixin],
  components: {
    Avatar
  },
  props: {
    author: {
      required: true,
    },
    story: {
      required: true,
    },
    small: {
      default: false
    }
  },
};
</script>
<style lang="scss">
.author-card__wrapper{
  display: flex;
  .author-card__avatar{
    margin-right: 10px;
  }
}
</style>
