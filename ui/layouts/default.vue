<template>
  <div class="wrapper w-full">
    <Header />
    <section class="mx-auto block md:flex md:flex-row w-full">
      <div class="basis-0/5 md:basis-1/5 wrapper-sidebar__category">
        <Sidebar :list-category="listNewCategory" />
      </div>
      <div class="basis-5/5 md:basis-4/5">
        <main>
          <slot />
        </main>
        <Footer />
      </div>
    </section>
  </div>
</template>

<script setup>
const route = useRoute();
useHead({
  meta: [
    { property: "og:title", content: `Retherer Blog - ${route.meta.title}` },
  ],
  title: "Retherer Blog",
});
import { GET } from "~/types/method";
const listNewCategory = ref({});
const getListNewCategory = async () => {
  const categoriesData = await useRestApi(GET, "/new-categories.json", {});
  if (categoriesData.status) {
    listNewCategory.value = categoriesData.data;
  }
};
getListNewCategory();
</script>

<style lang="scss" scoped></style>
