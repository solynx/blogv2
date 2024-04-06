<template>
<NuxtLink
  class="inline-flex items-center gap-2 rounded border border-indigo-600 bg-indigo-600 px-8 py-3 text-white hover:opacity-90 focus:outline-none focus:ring active:text-indigo-500"
  :to="props.href"
>
  <span class="text-sm font-medium"> <slot/> </span>

  <svg
    class="size-5 rtl:rotate-180"
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    stroke="currentColor"
  >
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      d="M17 8l4 4m0 0l-4 4m4-4H3"
    />
  </svg>
</NuxtLink>
</template>

<script setup>
const props = defineProps(["href"])
</script>

<style lang="scss" scoped>

</style>