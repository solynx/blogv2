<template>
  <Content :header="header">
    <hr class="md:my-6 my-3" />

    <div class="grid md:gap-4 relative">
      <PostItem :data="listNewPosts" v-if="!showLoading" />
    </div>
    <div
      class="flex items-center flex-col justify-center my-3 relative"
      v-if="showLoading"
    >
      <Skeleton class="block" />
      <Skeleton class="block" />
      <Skeleton class="block" />
      <Skeleton class="block" />
      <Skeleton class="block" />
      <Skeleton class="block" />
      <Spin class="absolute"></Spin>
    </div>
    <div class="w-full text-center mt-3">
      <Pagination @handle-paginate="handlePaginate" :data="paginationData" />
    </div>
  </Content>
</template>

<script setup lang="ts">
definePageMeta({
  title: "Tất cả bài viết",
});
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
  description: "Tất cả bài viết",
};
const listNewPosts = ref<PostOverview[]>([]);
const showLoading = ref(false);
const getNewPosts = async () => {
  showLoading.value = true;
  listNewPosts.value = [];
  const postsData = await useRestApi(POST, "/posts.json", {
    page: paginationData.value.page,
    limit: POST_PAGE_LIMIT,
  });
  if (postsData.status) {
    let listPosts: PostOverView[] = [];
    postsData.data.forEach((item, index) => {
      let post: PostOverview = {
        title: item.title || "",
        description: item.description || "",
        slug: item.slug || item.id,
        createdAt: useFormatDate(item.created_at),
        author: item.user,
        category: item.category,
      };
      listPosts.push(post);
    });
    listNewPosts.value = listPosts;
    paginationData.value.total = postsData.metadata?.total || 1;
  }
  setTimeout(() => {
    showLoading.value = false;
  }, 800);
};
const handlePaginate = async (page: string) => {
  let currentPage = parseInt(page);
  if (!currentPage) {
    return;
  }
  paginationData.value.page = currentPage;
  await getNewPosts();
};
await getNewPosts();
</script>

<style lang="scss" scoped></style>
