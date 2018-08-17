<template>
  <div v-if="storyIsReady">
    <editor :story="story"></editor>
  </div>
</template>

<script>
import editor from '@/components/editor';
import { GET_STORY, GET_STORY_BODY } from '@/store/actions.type';

export default {
  name: 'edit-story',
  props: ['id'],
  components: {
    editor,
  },
  data() {
    return {
      storyIsReady: false,
    }
  },
  computed: {
    story(){
      return this.$store.getters.getStoryById(this.id)
    },
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$store.dispatch(GET_STORY, vm.id).then( _ => {
        vm.$store.dispatch(GET_STORY_BODY, vm.id).then( _ => { vm.storyIsReady = true })
      })
    })
  },
};
</script>

<style>

</style>
