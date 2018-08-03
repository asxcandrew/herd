<template>
  <div id="app">
    <component :is="layout">
      <router-view />
    </component>
  </div>
</template>

<script>
const defaultLayout = 'default';

export default {
  name: 'App',
  computed: {
    layout(){
      return (this.$route.meta.layout || defaultLayout) + '-layout';
    },
    title(){
      return this.$t('meta.title.default');
    },
    titleTemplate(){
      return this.$t('meta.title.defaultTemplate');
    },
  },
  head: {
    title: function () {
      return {
        inner: this.title,
        separator: '|',
        complement: this.titleTemplate,
      }
    },
  },
};
</script>
