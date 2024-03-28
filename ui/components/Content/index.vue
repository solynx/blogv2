<template>
  <div class="bg-white sm:px-6 px-4 py-6 font-[sans-serif] overflow-hidden">
    <div class="w-full h-full">
      <div>
        <h2 class="text-3xl font-extrabold text-[#333] inline-block">
          {{ props.header?.title ?? 'Information' }}
        </h2>
        <p class="text-gray-400 text-sm mt-4">
          {{ props.header?.description ?? '' }}
        </p>
      </div>
      <slot/>
    </div>
  </div>
</template>

<script setup>
const props = defineProps(['header'])
</script>

<style lang="scss" scoped></style>
