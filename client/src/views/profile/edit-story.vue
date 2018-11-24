<template>
  <el-row v-if="storyIsReady">
    <el-col :span="24">
      <author-card :author="story.user" :story="story" :small="true"/>
    </el-col>
    <el-col :span="24">
      <text-editor :story="story"></text-editor>
    </el-col>
  </el-row>
</template>

<script>
import textEditor from '@/components/text-editor';
import authorCard from '@/components/author-card';

import { GET_STORY, GET_STORY_BODY } from '@/store/actions.type';

export default {
  name: 'edit-story',
  props: ['uid'],
  components: {
    textEditor,
    authorCard,
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
