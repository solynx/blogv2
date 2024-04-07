<template>
  <header
    class="shadow-md py-4 px-4 sm:px-10 bg-white font-[sans-serif] min-h-[70px]"
  >
    <div class="flex flex-wrap items-center justify-between gap-5 relative">
      <NuxtLink to="/" class="w-12 hidden lg:block">
        <img src="/public/favicon.ico" alt="avatar" class="rounded" />
      </NuxtLink>

      <ul
        id="collapseMenu"
        class="lg:!flex lg:space-x-5 flex max-lg:py-4 max-lg:w-full"
      >
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink
            to="/"
            class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]"
            >Home</NuxtLink
          >
        </li>
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink
            to="/project"
            class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]"
            >Projects</NuxtLink
          >
        </li>
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink
            to="/about"
            class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]"
            >About</NuxtLink
          >
        </li>
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink
            to="/contact"
            class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]"
            >Contact</NuxtLink
          >
        </li>
      </ul>
    </div>
  </header>
</template>

<script setup></script>

<style lang="scss" scoped></style>
