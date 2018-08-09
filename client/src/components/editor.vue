<template>
  <div>
    <h3>
      <input 
        class="frameless editor-title"
        v-model="editorTitle"
        @change="onEditorChange"
        :placeholder="this.$t('views.editor.newStoryTitle')">
    </h3>
    <quill-editor
              v-model="editorContent"
              :options="editorOption"
              @change="onEditorChange($event)">
    </quill-editor>
  </div>
</template>

<script>
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

import { quillEditor } from 'vue-quill-editor';
import { UPDATE_STATUS_BAR } from '@/store/actions.type';

export default {
  name: 'editor',
  components: {
    quillEditor,
  },
  methods: {
    onEditorChange(event) {
      this.contentHTML = event.html
      // this.$store.dispatch(UPDATE_STATUS_BAR, { isLoading: true, notice: 'saving' });
    },
    updatePost() {
      console.log('updete post')
    }
  },
  props: ['content', 'title'],
  data() {
    return {
      editorContent: this.content,
      editorTitle: this.title,
      contentHTML: '',
      editorOption: {
        modules: {
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'],
            ['blockquote', 'code-block'],
            [{ 'header': 1 }, { 'header': 2 }],
            [{ 'list': 'bullet' }],
            [{ 'font': [] }],
            [{ 'align': [] }],
            ['link', 'image', 'video']
          ],
          syntax: {
            highlight: text => hljs.highlightAuto(text).value
          }
        }
      },
    };
  }
};
</script>
<style>
.editor-title {
  width: 100%;
  text-align: center;
}
.frameless {
  border: 0;
  outline:0px solid transparent;
}
</style>
