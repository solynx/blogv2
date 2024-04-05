<template>
  <div>
    <ul class="flex md:space-x-4 justify-center">
      <li
        class="flex items-center justify-center shrink-0 hover:bg-gray-50 cursor-pointer w-10 h-10 rounded-lg"
        v-if="paginationListButton.length > 1"
      >
        <NuxtLink
          :to="
            paginationListButton.length > 1
              ? (path || '') + '?page=' + paginationListButton[0]
              : null
          "
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-4 fill-gray-400"
            viewBox="0 0 55.753 55.753"
          >
            <path
              d="M12.745 23.915c.283-.282.59-.52.913-.727L35.266 1.581a5.4 5.4 0 0 1 7.637 7.638L24.294 27.828l18.705 18.706a5.4 5.4 0 0 1-7.636 7.637L13.658 32.464a5.367 5.367 0 0 1-.913-.727 5.367 5.367 0 0 1-1.572-3.911 5.369 5.369 0 0 1 1.572-3.911z"
              data-original="#000000"
            />
          </svg>
        </NuxtLink>
      </li>

      <li
        class="flex items-center justify-center shrink-0 cursor-pointer text-base font-bold w-10 h-10 rounded-lg"
        v-for="page in paginationListButton"
        :class="page == currentPage ? 'bg-blue-500 text-white' : null"
      >
        <NuxtLink :to="(path || '') + '?page=' + page">
          {{ page }}
        </NuxtLink>
      </li>

      <li
        class="flex items-center justify-center shrink-0 hover:bg-gray-50 cursor-pointer w-10 h-10 rounded-lg"
        v-if="paginationListButton.length > 1"
      >
        <NuxtLink
          :to="
            paginationListButton.length > 1
              ? (path || '') +
                '?page=' +
                paginationListButton[paginationListButton.length - 1]
              : null
          "
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-4 fill-gray-400 rotate-180"
            viewBox="0 0 55.753 55.753"
          >
            <path
              d="M12.745 23.915c.283-.282.59-.52.913-.727L35.266 1.581a5.4 5.4 0 0 1 7.637 7.638L24.294 27.828l18.705 18.706a5.4 5.4 0 0 1-7.636 7.637L13.658 32.464a5.367 5.367 0 0 1-.913-.727 5.367 5.367 0 0 1-1.572-3.911 5.369 5.369 0 0 1 1.572-3.911z"
              data-original="#000000"
            />
          </svg>
        </NuxtLink>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
const props = defineProps(["data"]);
const currentPage = props.data.page;
const path = props.data.path;
const totalPages = await Math.ceil(props.data.total / props.data.limit);
const paginationListButton = ref([]);

async function createPaginationListButton() {
  if (currentPage <= totalPages) {
    if (currentPage > 1) {
      paginationListButton.value.push(currentPage - 1);
    }
    for (let i = currentPage; i <= totalPages; i++) {
      if (i > currentPage + 3 && i < totalPages) {
        paginationListButton.value.push("...");
        paginationListButton.value.push(i);
        break;
      }
      paginationListButton.value.push(i);
    }
  }
}
await createPaginationListButton();
</script>

<style lang="scss" scoped></style>
