<template>
  <ClientOnly>
    <div>
      <div v-if="props.editor">
        <button
          @click="props.editor.chain().focus().toggleBold().run()"
          :disabled="!editor.can().chain().focus().toggleBold().run()"
          :class="{ 'is-active': editor.isActive('bold') }"
        >
          bold
        </button>
        <button
          @click="props.editor.chain().focus().toggleItalic().run()"
          :disabled="!editor.can().chain().focus().toggleItalic().run()"
          :class="{ 'is-active': editor.isActive('italic') }"
        >
          italic
        </button>
        <button
          @click="props.editor.chain().focus().toggleStrike().run()"
          :disabled="!editor.can().chain().focus().toggleStrike().run()"
          :class="{ 'is-active': editor.isActive('strike') }"
        >
          strike
        </button>
        <button
          @click="props.editor.chain().focus().toggleCode().run()"
          :disabled="!editor.can().chain().focus().toggleCode().run()"
          :class="{ 'is-active': editor.isActive('code') }"
        >
          code
        </button>
        <button @click="props.editor.chain().focus().unsetAllMarks().run()">
          clear marks
        </button>
        <button @click="props.editor.chain().focus().clearNodes().run()">
          clear nodes
        </button>
        <button
          @click="props.editor.chain().focus().setParagraph().run()"
          :class="{ 'is-active': editor.isActive('paragraph') }"
        >
          paragraph
        </button>
        <button
          @click="
            props.editor.chain().focus().toggleHeading({ level: 1 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }"
        >
          h1
        </button>
        <button
          @click="
            props.editor.chain().focus().toggleHeading({ level: 2 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }"
        >
          h2
        </button>
        <button
          @click="
            props.editor.chain().focus().toggleHeading({ level: 3 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }"
        >
          h3
        </button>
        <button
          @click="
            props.editor.chain().focus().toggleHeading({ level: 4 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 4 }) }"
        >
          h4
        </button>
        <button
          @click="
            props.editor.chain().focus().toggleHeading({ level: 5 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 5 }) }"
        >
          h5
        </button>
        <button
          @click="
            props.editor.chain().focus().toggleHeading({ level: 6 }).run()
          "
          :class="{ 'is-active': editor.isActive('heading', { level: 6 }) }"
        >
          h6
        </button>
        <button
          @click="props.editor.chain().focus().toggleBulletList().run()"
          :class="{ 'is-active': editor.isActive('bulletList') }"
        >
          bullet list
        </button>
        <button
          @click="props.editor.chain().focus().toggleOrderedList().run()"
          :class="{ 'is-active': editor.isActive('orderedList') }"
        >
          ordered list
        </button>
        <button
          @click="props.editor.chain().focus().toggleCodeBlock().run()"
          :class="{ 'is-active': editor.isActive('codeBlock') }"
        >
          code block
        </button>
        <button
          @click="props.editor.chain().focus().toggleBlockquote().run()"
          :class="{ 'is-active': editor.isActive('blockquote') }"
        >
          blockquote
        </button>
        <button @click="props.editor.chain().focus().setHorizontalRule().run()">
          horizontal rule
        </button>
        <button @click="props.editor.chain().focus().setHardBreak().run()">
          hard break
        </button>
        <button
          @click="props.editor.chain().focus().undo().run()"
          :disabled="!editor.can().chain().focus().undo().run()"
        >
          undo
        </button>
        <button
          @click="props.editor.chain().focus().redo().run()"
          :disabled="!editor.can().chain().focus().redo().run()"
        >
          redo
        </button>
      </div>
      <TiptapEditorContent
        class="mt-3"
        v-model="props.content"
        :editor="props.editor"
      />
    </div>
  </ClientOnly>
</template>

<script setup>
const props = defineProps(["editor"]);

onMounted(() => {
  if (!!unref(props.editor)) {
    unref(props.editor).commands.setContent(props.content);
  }
});

onBeforeUnmount(() => {
  unref(props.editor).destroy();
});
</script>
<style scoped>
button {
  padding: 3px;
  margin: 0 5px;
  border: 1px solid;
}
</style>
