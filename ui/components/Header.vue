<template>
  <header
    class="shadow-md py-4 px-4 sm:px-10 bg-white font-[sans-serif] min-h-[70px]"
  >
    <div class="flex flex-wrap items-center justify-between gap-5 relative">
      <NuxtLink to="/" class="w-8"> <img
                src="https://api.dicebear.com/8.x/adventurer/svg?seed=Boo"
                alt="avatar"
              /> </NuxtLink>
      <div class="flex lg:order-1 max-sm:ml-auto">
        <button
          class="px-4 py-2 text-sm rounded-full font-bold text-white border-2 border-[#007bff] bg-[#007bff] transition-all ease-in-out duration-300 hover:bg-transparent hover:text-[#007bff]"
        >
          Login
        </button>
        <button
          class="px-4 py-2 text-sm rounded-full font-bold text-white border-2 border-[#007bff] bg-[#007bff] transition-all ease-in-out duration-300 hover:bg-transparent hover:text-[#007bff] ml-3"
        >
          Sign up
        </button>
        <button id="toggle" class="lg:hidden ml-7">
          <svg
            class="w-7 h-7"
            fill="#000"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
              clip-rule="evenodd"
            ></path>
          </svg>
        </button>
      </div>
      <ul
        id="collapseMenu"
        class="lg:!flex lg:space-x-5 max-lg:space-y-2 max-lg:hidden max-lg:py-4 max-lg:w-full"
      >
      <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink to="/"  class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]">Home</NuxtLink>
        </li>
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink to="/project"  class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]">Projects</NuxtLink>
        </li>
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink to="/category"  class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]">Categories</NuxtLink>
        </li>
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink to="/about"  class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]">About</NuxtLink>
        </li>
        <li class="max-lg:border-b max-lg:py-2 px-3 max-lg:rounded">
          <NuxtLink to="/contact"  class="lg:hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]">Contact</NuxtLink>
        </li>
      </ul>
    </div>
  </header>
</template>

<script>
export default {
  setup() {
    return {};
  },
};
</script>

<style lang="scss" scoped></style>
