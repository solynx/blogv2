<template>
  <div
    class="pt-3 cursor-pointer rounded overflow-hidden group"
    v-for="item in props.data"
  >
    <span class="text-sm block text-gray-400 mb-3 flex justify-between">
      <p>{{ useFormatDate(item.created_at) }}</p>
    </span>
    <NuxtLink
      :to="'/post/' + item.slug"
      class="text-md group-hover:text-blue-500 transition-all"
    >
      {{ item.title }}
    </NuxtLink>
    <div class="flex flex-wrap items-center gap-3 mt-3">
      <p class="text-xs text-gray-400">BY {{ item.user.full_name }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps(["data"]);
</script>

<style scoped></style>
