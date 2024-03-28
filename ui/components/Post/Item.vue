<template>
    <div class="pt-3 cursor-pointer rounded overflow-hidden group">
      <div>
        <span class="text-sm block text-gray-400 mb-3 flex justify-between">
          <p class="rounded-full bg-slate-100 px-2 py-1">Category</p>
          <p>10/6/2024</p>
        </span>
        <NuxtLink
          to="/post/slug" class="text-xl font-bold text-[#333] group-hover:text-blue-500 transition-all"
        >
          Trends and Predictions
        </NuxtLink>
        <div class="mt-3">
          <p class="text-gray-400 text-sm">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis
            accumsan, nunc et tempus blandit, metus mi consectetur felis turpis
            vitae ligula.
          </p>
        </div>
      </div>
      <div class="flex flex-wrap items-center gap-3 mt-8">
        <img
          src="https://readymadeui.com/team-3.webp"
          class="w-10 h-10 rounded-full"
        />
        <p class="text-xs text-gray-400">BY SIMON KONECKI</p>
      </div>
    </div>
    <hr />
  </template>
  
  <script>
  export default {
    setup() {
      return {};
    },
  };
  </script>
  
  <style lang="scss" scoped></style>
  