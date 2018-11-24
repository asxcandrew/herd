<template>
  <div class="container">
    <editor-floating-menu :editor="editor">
      <div
        slot-scope="{ commands, isActive, menu }"
        class="editor__floating-menu"
        :class="{ 'is-active': menu.isActive }"
        :style="`top: ${menu.top}px`"
      >
        <font-awesome-icon
          icon="plus"
          class="floating-menu__button"
          @click="showFloatingMenu()"
        />

      </div>
    </editor-floating-menu>
    <div class="editor">
      <div class="editor__floating-menu" v-if="floatingMenuIsActive">
        <button
          class="menubar__button"
          :class="{ 'is-active': isActive.heading({ level: 1 }) }"
          @click="commands.heading({ level: 1 })"
        >
          H1
        </button>
      </div>
      <editor-menu-bubble :editor="editor">
        <div
          class="editor__menububble"
          slot-scope="{ commands, isActive, menu, getMarkAttrs }"
          :class="{ 'is-active': menu.isActive }"
          :style="`left: ${menu.left}px; bottom: ${menu.bottom + 10}px;`"
        >
            <form class="menububble__form" v-if="linkMenuIsActive" @submit.prevent="setLinkUrl(commands.link, linkUrl)">
              <input class="menububble__input" type="text" v-model="linkUrl" placeholder="https://" ref="linkInput" @keydown.esc="hideLinkMenu"/>
              <button class="menububble__button" @click="setLinkUrl(commands.link, null)" type="button">
                <font-awesome-icon icon="trash" />
              </button>
            </form>
              <button
                class="menububble__button"
                :class="{ 'is-active': isActive.bold() }"
                @click="commands.bold"
              >
                <font-awesome-icon icon="bold" />
              </button>
              <button
                class="menububble__button"
                :class="{ 'is-active': isActive.italic() }"
                @click="commands.italic"
              >
                <font-awesome-icon icon="italic" />
              </button>
              <button
                class="menububble__button"
                :class="{ 'is-active': isActive.code() }"
                @click="commands.code"
              >
                <font-awesome-icon icon="code" />
              </button>
              <button
                  class="menububble__button"
                  @click="showLinkMenu(getMarkAttrs('link'))"
                  :class="{ 'is-active': isActive.link() }"
                >
                  <span>Add Link</span>
                  <font-awesome-icon icon="link" />
              </button>
        </div>
      </editor-menu-bubble>
      <editor-content class="editor__content" :editor="titleEditor" />
      <editor-content class="editor__content" :editor="editor" />
    </div>
  </div>
</template>

<script>

import { Editor, EditorContent, EditorFloatingMenu, EditorMenuBubble } from 'tiptap';
import {
  Blockquote,
  BulletList,
  CodeBlock,
  HardBreak,
  Heading,
  ListItem,
  OrderedList,
  TodoItem,
  TodoList,
  Bold,
  Code,
  Italic,
  Link,
  History,
  Placeholder,
} from 'tiptap-extensions';

import {
  UPDATE_STATUS_BAR,
  UPDATE_STORY,
} from '@/store/actions.type';

export default {
  name: 'text-editor',
  components: {
    EditorFloatingMenu,
    EditorContent,
    EditorMenuBubble,
  },
  methods: {
    showLinkMenu(type) {
			this.linkUrl = type.attrs.href
			this.linkMenuIsActive = true
			this.$nextTick(() => {
				this.$refs.linkInput.focus()
			})
    },
    hideLinkMenu() {
			this.linkUrl = null
			this.linkMenuIsActive = false
    },
    showFloatingMenu() {
      this.floatingMenuIsActive = true
    },
		setLinkUrl(url, type, focus) {
			type.command({ href: url })
			this.hideLinkMenu()
			focus()
		},
    updatePost() {
      // this.$store.dispatch(UPDATE_STATUS_BAR, { isLoading: true, notice: 'saving...' });
      this.$store.dispatch(
        UPDATE_STORY,
        {
          uid: this.story.uid,
          params: { html_body: this.editorContent, title: this.editorTitle },
        },
      ).then(() => {
        this.$store.dispatch(UPDATE_STATUS_BAR, { notice: 'saved' });
      });
    }
  },
  props: ['story'],
  data() {
    return {
      // editorContent: this.story.html_body,
      // editorTitle: this.story.title,
      linkUrl: null,
      linkMenuIsActive: false,
      floatingMenuIsActive: false,
      titleEditor: new Editor({
        extensions: [
          new Placeholder({
            emptyClass: 'is-empty',
          }),
        ],
        onUpdate({ getJSON, getHTML }) {
          console.log(getJSON())
          // console.log(getHTML())
          // this.updatePost(getHTML());
        },
      }),
      editor: new Editor({
        extensions: [
          new Blockquote(),
          new BulletList(),
          new CodeBlock(),
          new HardBreak(),
          new ListItem(),
          new Heading({ maxLevel: 3 }),
          new Bold(),
          new Code(),
          new OrderedList(),
          new Italic(),
          new TodoItem(),
          new TodoList(),
          new Link(),
          new History(),
          new Placeholder({
            emptyNodeClass: 'is-empty',
          }),
        ],
        onUpdate({ state }) {
          console.log(state)
          // console.log(getHTML())
          // this.updatePost(getHTML());
        },
      }),
    };
  },
  beforeDestroy() {
    this.editor.destroy();
  },
};
</script>
<style lang="scss">
.editor-title {
  width: 100%;
  text-align: center;
}
.frameless {
  border: 0;
  outline:0px solid transparent;
}
</style>
