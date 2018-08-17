<template>
  <div class="center">
    <h3>
      <input 
        class="frameless editor-title"
        v-model="editorTitle"
        @change="onChangeTitle"
        :placeholder="this.$t('views.editor.newStoryTitle')">
    </h3>
    <quill-editor
              v-model="editorContent"
              :options="editorOption"
              @change="onChange($event)"
              @blur="onBlur($event)">
    </quill-editor>
  </div>
</template>

<script>
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

import { quillEditor } from 'vue-quill-editor';
import {
  UPDATE_STATUS_BAR,
  CLEAR_STATUS_BAR,
  UPDATE_STORY,
} from '@/store/actions.type';

export default {
  name: 'editor',
  components: {
    quillEditor,
  },
  methods: {
    onChange(event) {
      this.updatePost();
    },
    onChangeTitle() {
      this.updatePost();
    },
    onBlur(event) {
      this.updatePost();
    },
    updatePost() {
      this.$store.dispatch(UPDATE_STATUS_BAR, { isLoading: true, notice: 'saving' });
      this.$store.dispatch(
        UPDATE_STORY,
        {
          id: this.story.id,
          params: { html_body: this.editorContent, title: this.editorTitle },
        },
      ).then(() => {
        this.$store.dispatch(CLEAR_STATUS_BAR)
      });
    }
  },
  props: ['story'],
  data() {
    return {
      editorContent: this.story.html_body,
      editorTitle: this.story.title,
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
