<template>
  <div class="bg-white sm:px-6 px-4 py-6 font-[sans-serif] overflow-hidden">
    <div class="w-full h-full">
      <div>
        <h2 class="text-3xl font-extrabold text-[#333] inline-block">
          LATEST BLOGS
        </h2>
        <p class="text-gray-400 text-sm mt-4">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis
          accumsan, nunc et tempus blandit, metus mi consectetur felis turpis
          vitae ligula.
        </p>
      </div>
      <hr class="md:my-6 my-3" />
      <div class="grid md:gap-4">
        <Post />
        <Post />
      </div>
      <Pagination class="mt-5" />
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
