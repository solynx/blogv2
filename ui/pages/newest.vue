<template>
  <Content :header="header">
    <hr class="md:my-6 my-3" />
    <div class="grid md:gap-4">
      <PostItem :data="listNewPosts" />
    </div>

    <div class="w-full text-center mt-3">
      <Pagination :data="paginationData" />
    </div>
  </Content>
</template>

<script setup lang="ts">
import { GET, POST } from "~/types/method";
import { type PostOverview, POST_PAGE_LIMIT } from "~/types/post";
const route = useRoute();
const paginationData = ref({
  total: 0,
  page: parseInt(route.query?.page) > 0 ? parseInt(route.query?.page) : 1,
  path: "/newest",
  limit: POST_PAGE_LIMIT,
});
const header = {
  title: "Mới nhất",
  description: "Lastest Post",
};
const listNewPosts = ref<PostOverview[]>([]);
const getNewPosts = async () => {
  const postsData = await useRestApi(POST, "/posts.json", {
    page: 1,
    limit: POST_PAGE_LIMIT,
  });
  if (postsData.status) {
    postsData.data.forEach((item, index) => {
      let post: PostOverview = {
        title: item.title || "",
        description: item.description || "",
        slug: item.slug || item.id,
        createdAt: useFormatDate(item.created_at),
        author: item.user,
        category: item.category,
      };
      listNewPosts.value.push(post);
    });
    paginationData.value.total = postsData.metadata?.total || 1;
  }
};
await getNewPosts();
</script>

<style lang="scss" scoped></style>
