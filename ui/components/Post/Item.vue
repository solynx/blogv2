<template>
    <div class="mt-1 cursor-pointer rounded overflow-hidden group" v-for="item in props.data">
      <div>
        <span class="text-sm block text-gray-400 mb-3 flex justify-between">
          <p class="rounded-full bg-slate-100 px-3 py-1">{{ item.category.name}}</p>
          <p>{{ item.createdAt }}</p>
        </span>
        <NuxtLink
          :to="'/post/' + item.slug" class="text-xl font-bold text-[#333] group-hover:text-blue-500 transition-all "
        >
          {{ item.title }}
        </NuxtLink>
        <div class="mt-3" v-if="item.description">
          <p class="text-gray-400 text-sm">
            {{ item.description }}
          </p>
        </div>
      </div>
      <div class="flex flex-wrap items-center gap-3 mt-3 mb-3">
        <p class="text-xs text-gray-400">@ {{ item.author.full_name }}</p>
      </div>
      <hr />
    </div>

  </template>
  
<script setup lang="ts">
  const props = defineProps(["data"])

</script>
  
  <style lang="scss" scoped></style>
  