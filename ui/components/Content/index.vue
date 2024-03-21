<template>
  <div class="bg-white sm:px-6 px-4 py-6 font-[sans-serif] overflow-hidden">
    <div class="w-full">
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
      <hr class="my-10" />
      <div class="grid md:gap-16">
        <div class="cursor-pointer rounded overflow-hidden group">
          <div>
            <span class="text-sm block text-gray-400 mb-4">10 FEB 2023</span>
            <h3
              class="text-xl font-bold text-[#333] group-hover:text-blue-500 transition-all"
            >
              A Guide to Igniting Your Imagination
            </h3>
            <div class="mt-6">
              <p class="text-gray-400 text-sm">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis
                accumsan, nunc et tempus blandit, metus mi consectetur felis
                turpis vitae ligula.
              </p>
            </div>
          </div>
          <div class="flex flex-wrap items-center gap-3 mt-8">
            <img
              src="https://readymadeui.com/team-1.webp"
              class="w-10 h-10 rounded-full"
            />
            <p class="text-xs text-gray-400">BY JOHN DOE</p>
          </div>
        </div>
        <div class="cursor-pointer rounded overflow-hidden group">
          <div>
            <span class="text-sm block text-gray-400 mb-4">7 JUN 2023</span>
            <h3
              class="text-xl font-bold text-[#333] group-hover:text-blue-500 transition-all"
            >
              Hacks to Supercharge Your Day
            </h3>
            <div class="mt-6">
              <p class="text-gray-400 text-sm">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis
                accumsan, nunc et tempus blandit, metus mi consectetur felis
                turpis vitae ligula.
              </p>
            </div>
          </div>
          <div class="flex flex-wrap items-center gap-3 mt-8">
            <img
              src="https://readymadeui.com/team-2.webp"
              class="w-10 h-10 rounded-full"
            />
            <p class="text-xs text-gray-400">BY MARK ADAIR</p>
          </div>
        </div>
        <div class="cursor-pointer rounded overflow-hidden group">
          <div>
            <span class="text-sm block text-gray-400 mb-4">5 OCT 2023</span>
            <h3
              class="text-xl font-bold text-[#333] group-hover:text-blue-500 transition-all"
            >
              Trends and Predictions
            </h3>
            <div class="mt-6">
              <p class="text-gray-400 text-sm">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis
                accumsan, nunc et tempus blandit, metus mi consectetur felis
                turpis vitae ligula.
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
      </div>
      <Pagination />
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
