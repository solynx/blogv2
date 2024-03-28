<template>
  <div class="pt-3 cursor-pointer rounded overflow-hidden group">
    <div>
      <span class="text-sm block text-gray-400 mb-3 flex justify-between">
        <p class="rounded-full bg-slate-100 px-2 py-1">Category</p>
        <p>10/6/2024</p>
      </span>
      <div class="mt-3">
        <p class="text-gray-400 text-sm px-2">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis
          accumsan, nunc et tempus blandit, metus mi consectetur felis turpis
          vitae ligula. Tất cả
        </p>
      </div>
    </div>
    <div class="flex flex-wrap items-center gap-3 mt-8">
      <img
        :src="'https://api.dicebear.com/8.x/adventurer/svg?seed=' + Math.random(0,100)"
        alt="avatar"
        class="w-10 h-10 rounded-full"
      />
      <p class="text-xs text-gray-400">BY SIMON KONECKI</p>
    </div>
  </div>
</template>

<script>
export default {
  setup() {
    return {};
  },
};
</script>

<style lang="scss" scoped></style>
