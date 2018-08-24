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
  props: ['uid'],
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
      return this.$store.getters.getStoryByUid(this.uid)
    },
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$store.dispatch(GET_STORY, vm.uid).then( _ => {
        vm.$store.dispatch(GET_STORY_BODY, vm.uid).then( _ => { vm.storyIsReady = true })
      })
    })
  },
};
</script>

<style>

</style>
