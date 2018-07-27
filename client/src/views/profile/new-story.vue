<template>
  <div>
    <h2>{{ $t('views.editor.newStoryTitle') }}</h2>
    <quill-editor ref="myTextEditor"
              v-model="content"
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

export default {
  name: 'new-story',
  components: {
    quillEditor,
  },
  methods: {
    onEditorChange(event) {
      // this.content = html
    }
  },
  data() {
    return {
      content: '',
      editorOption: {
        modules: {
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'],
            ['blockquote', 'code-block'],
            [{ 'header': 1 }, { 'header': 2 }],
            [{ 'list': 'ordered' }, { 'list': 'bullet' }],
            [{ 'indent': '-1' }, { 'indent': '+1' }],
            [{ 'size': ['small', false, 'large', 'huge'] }],
            [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
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

</style>
