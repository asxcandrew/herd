<template>
  <div>
    <h2>{{ $t('views.editor.newStoryTitle') }}</h2>
    <editor content="content">
  </div>
</template>

<script>
import editor from '@/components/editor';

export default {
  name: 'new-story',
  components: {
    editor,
  },
  computed: {
    content(){
      return 'edit your story';
    }
  }
};
</script>

<style>

</style>
