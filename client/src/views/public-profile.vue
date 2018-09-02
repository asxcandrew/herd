<template>
<div>public profile page in heavy development</div>
</template>

<script>
import {
  GET_USER,
} from '@/store/actions.type';
export default {
  name: 'public-profile',
  props: ['username'],
  computed: {
    user(){
      return this.$store.getters.getUserByUsername(this.username);
    },
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.$store.dispatch(GET_USER, vm.$route.params.username).then( user => {
      })
    })
  },
};
</script>

<style>

</style>
