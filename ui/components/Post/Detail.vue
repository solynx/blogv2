<template>
  <div class="pt-3 cursor-pointer rounded overflow-hidden group">
    <div>
      <span class="text-sm block text-gray-400 mb-3 flex justify-between">
        <p class="rounded-full bg-slate-100 px-2 py-1">
          {{ props.data.category?.name ?? "" }}
        </p>
        <p>{{ props.data.createdAt ?? "" }}</p>
      </span>
      <h1 class="text-2xl font-bold">{{ props.data.title }}</h1>
      <div
        class="mt-3 !prose dark:prose-invert !prose-sm sm:prose-base max-w-screen-xl"
        v-html="props.data.content"
      ></div>
    </div>
    <div class="flex flex-wrap items-center gap-3 mt-8">
      <img
        :src="
          'https://api.dicebear.com/8.x/adventurer/svg?seed=' +
          props.data.user?.id
        "
        alt="avatar"
        class="w-10 h-10 rounded-full"
      />
      <p class="text-xs text-gray-400">
        @ {{ props.data.author?.full_name ?? "author" }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps(["data"]);
</script>

<style scoped>
.no-tailwindcss > * {
  all: unset;
}
</style>
