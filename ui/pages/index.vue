<template>
  <Content :header="header">
    <hr class="md:my-6 my-3" />
      <div class="grid md:gap-4">
        <PostItem :data="listNewPosts"/>
      </div>
      <Pagination class="mt-5" />
    <div class="mx-3 my-6">
      <div class="my-3">
        <h2 class="text-3xl font-extrabold text-[#333] inline-block">Category</h2>
      </div>
      <div class="block md:grid md:grid-cols-3 md:gap-3">
        <Category :data="listNewCategory"/>
      </div>
    </div>
  </Content>
</template>
<script setup lang="ts">
import { GET, POST } from "~/types/method"
import { type PostOverview } from "~/types/post"
const header = {
  title : "Mới nhất",
  description: 'Lastest Post'
}
const listNewPosts = ref<PostOverview[]>([]);
const listNewCategory = ref([]);
const getNewPosts = async () => {
  const postsData = await useRestApi(POST, "/posts.json", { page: 1, limit: 6})
  if (postsData.status) {
    postsData.data.forEach((item, index) => {
      let post : PostOverview = {
        title: item.title || "",
        description: item.description || "",
        slug: item.slug || item.id,
        createdAt: useFormatDate(item.created_at),
        author: item.user,
        category: item.category
      }
      listNewPosts.value.push(post)
    });
  }
}

const getListNewCategory = async () => {
  const categoriesData = await useRestApi(GET, "/new-categories.json", {})
  if (categoriesData.status) {
    listNewCategory.value = categoriesData.data
  }
}
getNewPosts();
getListNewCategory();
</script>