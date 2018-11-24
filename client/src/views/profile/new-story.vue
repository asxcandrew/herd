<template>
  <div>
    <text-editor :story="story" v-if="story"/>
  </div>
</template>

<script>
import textEditor from '@/components/text-editor';
import { CREATE_STORY } from '@/store/actions.type';

export default {
  name: 'new-story',
  components: {
    textEditor,
  },
  created() {
    this.$store.dispatch(CREATE_STORY, this.params).then(story => {
      this.$router.push({ name: 'edit-story', params: { uid: story.uid }});
    });
  },
  computed: {
    story(){
      return this.$store.getters.getStoryByUid(this.uid)
    },
  },
  data() {
    return {
      uid: '',
      params: {
        html_body: 'Write your story...',
        title: 'Title',
      }
    };
  },
};
</script>

<style>
</style>
