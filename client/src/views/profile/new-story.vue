<template>
  <div>
    <editor :story="story" v-if="story"/>
  </div>
</template>

<script>
import editor from '@/components/editor';
import { CREATE_STORY } from '@/store/actions.type';

export default {
  name: 'new-story',
  components: {
    editor,
  },
  created() {
    this.$store.dispatch(CREATE_STORY, this.params).then(data => {
      this.id = data.data.id
    })
  },
  computed: {
    story(){
      return this.$store.getters.getStoryById(this.id)
    },
  },
  data() {
    return {
      id: '',
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
