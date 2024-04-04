<template>
  <ClientOnly>
    <div>
      <div v-if="editor">
        <button
          @click="editor.chain().focus().toggleBold().run()"
          :disabled="!editor.can().chain().focus().toggleBold().run()"
          :class="{ 'is-active': editor.isActive('bold') }"
        >
          bold
        </button>
        <button
          @click="editor.chain().focus().toggleItalic().run()"
          :disabled="!editor.can().chain().focus().toggleItalic().run()"
          :class="{ 'is-active': editor.isActive('italic') }"
        >
          italic
        </button>
        <button
          @click="editor.chain().focus().toggleStrike().run()"
          :disabled="!editor.can().chain().focus().toggleStrike().run()"
          :class="{ 'is-active': editor.isActive('strike') }"
        >
          strike
        </button>
        <button
          @click="editor.chain().focus().toggleCode().run()"
          :disabled="!editor.can().chain().focus().toggleCode().run()"
          :class="{ 'is-active': editor.isActive('code') }"
        >
          code
        </button>
        <button @click="editor.chain().focus().unsetAllMarks().run()">
          clear marks
        </button>
        <button @click="editor.chain().focus().clearNodes().run()">
          clear nodes
        </button>
        <button
          @click="editor.chain().focus().setParagraph().run()"
          :class="{ 'is-active': editor.isActive('paragraph') }"
        >
          paragraph
        </button>
        <button
          @click="
           editor.chain().focus().toggleHeading({ level: 1 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }"
        >
          h1
        </button>
        <button
          @click="
            editor.chain().focus().toggleHeading({ level: 2 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }"
        >
          h2
        </button>
        <button
          @click="
            editor.chain().focus().toggleHeading({ level: 3 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }"
        >
          h3
        </button>
        <button
          @click="
           editor.chain().focus().toggleHeading({ level: 4 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 4 }) }"
        >
          h4
        </button>
        <button
          @click="
            editor.chain().focus().toggleHeading({ level: 5 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 5 }) }"
        >
          h5
        </button>
        <button
          @click="
            editor.chain().focus().toggleHeading({ level: 6 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 6 }) }"
        >
          h6
        </button>
        <button
          @click="editor.chain().focus().toggleBulletList().run()"
          :class="{ 'is-active': editor.isActive('bulletList') }"
        >
          bullet list
        </button>
        <button
          @click="editor.chain().focus().toggleOrderedList().run()"
          :class="{ 'is-active': editor.isActive('orderedList') }"
        >
          ordered list
        </button>
        <button
          @click="editor.chain().focus().toggleCodeBlock().run()"
          :class="{ 'is-active': editor.isActive('codeBlock') }"
        >
          code block
        </button>
        <button
          @click="editor.chain().focus().toggleBlockquote().run()"
          :class="{ 'is-active': editor.isActive('blockquote') }"
        >
          blockquote
        </button>
        <button @click="editor.chain().focus().setHorizontalRule().run()">
          horizontal rule
        </button>
        <button @click="editor.chain().focus().setHardBreak().run()">
          hard break
        </button>
        <button
          @click="editor.chain().focus().undo().run()"
          :disabled="!editor.can().chain().focus().undo().run()"
        >
          undo
        </button>
        <button
          @click="editor.chain().focus().redo().run()"
          :disabled="!editor.can().chain().focus().redo().run()"
        >
          redo
        </button>
        <button @click="editor.chain().focus().setTextAlign('left').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'left' }) }">
        left
        </button>
        <button @click="editor.chain().focus().setTextAlign('center').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'center' }) }">
          center
        </button>
        <button @click="editor.chain().focus().setTextAlign('right').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'right' }) }">
          right
        </button>
        <button @click="editor.chain().focus().setTextAlign('justify').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'justify' }) }">
          justify
        </button>
        <button @click="editor.chain().focus().unsetTextAlign().run()">
          unsetTextAlign
        </button>
        <button @click="addImage">
          setImage
        </button>
        <button id="add" @click="addVideo">
        Add YouTube video
      </button>
      </div>
      <TiptapEditorContent
        class="mt-3"
        :editor="editor"
      />
    </div>
  </ClientOnly>
</template>

<script setup lang="ts">
const props = defineProps({
  editor : {
    type: Object,
    required:true
  }
});


const editor = ref(props.editor)

const addImage = () => {
  const url = window.prompt('URL')
  if (url) {
    editor.value.chain().focus().setImage({ src: url }).run()
  }
}

const addVideo = () => {
  const url = prompt('Enter YouTube URL')

  editor.value.commands.setYoutubeVideo({
    src: url,
    width:640,
    height:  480,
  })
}
// onMounted(() => {
//   if (!!unref(editor)) {
//     unref(editor).commands.setContent(props.data);
//   }
// });

// onBeforeUnmount(() => {
//   unref(editor).destroy();
// });
</script>
<style scoped>
button {
  padding: 3px;
  margin: 5px 5px;
  border: 1px solid;
}
</style>
