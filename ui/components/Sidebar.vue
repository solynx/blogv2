<template>
  <nav
    class="bg-white shadow-lg h-screen sticky top-0 left-0 w-full md:min-w-[240px] py-6 px-4 font-[sans-serif] hidden md:block overflow-auto"
  >
    <div class="mt-4">
      <h6 class="text-blue-600 text-sm font-bold px-4">Danh mục</h6>
      <ul class="mt-2">
        <li v-for="item in props.listCategory">
          <NuxtLink
            :to="'/post/category/' + item.slug"
            class="text-black hover:text-blue-600 text-[15px] block hover:bg-blue-50 rounded px-4 py-2.5 transition-all"
          >
            {{ item.name }}
          </NuxtLink>
        </li>
      </ul>
    </div>
    <div class="mt-4">
      <h6 class="text-blue-600 text-sm font-bold px-4">Dự kiến tới</h6>
      <ul class="mt-2">
        <li>
          <a
            href="https://deno.com/"
            target="_blank"
            class="text-black hover:text-blue-600 text-[15px] block hover:bg-blue-50 rounded px-4 py-2.5 transition-all"
          >
            Deno
          </a>
        </li>
        <li>
          <a
            href="https://solana.com/vi/developers/guides/getstarted/setup-local-development"
            target="_blank"
            class="text-black hover:text-blue-600 text-[15px] block hover:bg-blue-50 rounded px-4 py-2.5 transition-all"
          >
            Solana
          </a>
        </li>
        <li>
          <a
            href="https://www.rust-lang.org/"
            target="_blank"
            class="text-black hover:text-blue-600 text-[15px] block hover:bg-blue-50 rounded px-4 py-2.5 transition-all"
          >
            Rust
          </a>
        </li>
        <li>
          <a
            href="https://docs.rs/axum/latest/axum/"
            target="_blank"
            class="text-black hover:text-blue-600 text-[15px] block hover:bg-blue-50 rounded px-4 py-2.5 transition-all"
          >
            Axum
          </a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script setup>
const props = defineProps(["listCategory"]);
</script>

<style lang="scss" scoped></style>
