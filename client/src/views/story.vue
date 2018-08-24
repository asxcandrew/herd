<template>
    <p>story</p>
</template>

<script>
import {
  GET_STORY,
  GET_STORY_BODY,
} from '@/store/actions.type';
import storyMixin from '@/mixins/story-utils';

export default {
  name: 'show-story',
  mixins: [storyMixin],
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$store.dispatch(GET_STORY, vm.uidFromUri(vm.$route.params.uri)).then( story => {
        vm.$store.dispatch(GET_STORY_BODY, story.uid).then( _ => {  })
      })
    })
  },
};
</script>
