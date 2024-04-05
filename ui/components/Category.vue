<template>
  <div class="mb-3 md:m-0 p-6 w-full h-full bg-slate-100 text-center" v-for="item in props.data">
    <NuxtLink :to="'/post/category/' + item.slug">
      {{ item.name ?? '' }}
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
const props = defineProps(['data'])
</script>

<style lang="scss" scoped></style>
