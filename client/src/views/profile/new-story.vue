<template>
  <div>
    <editor :content="content" />
  </div>
</template>

<script>
import editor from '@/components/editor';

export default {
  name: 'new-story',
  components: {
    editor,
  },
  data() {
    return {
      content: 'Write your story...',
    };
  }
};
</script>

<style>
</style>
