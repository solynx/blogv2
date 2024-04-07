<template>
  <div>
    <div v-if="!foundCategory">
      <div
        class="bg-red-50 border-s-4 border-red-500 p-4 dark:bg-red-800/30"
        role="alert"
      >
        <div class="flex">
          <div class="flex-shrink-0">
            <!-- Icon -->
            <span
              class="inline-flex justify-center items-center size-8 rounded-full border-4 border-red-100 bg-red-200 text-red-800 dark:border-red-900 dark:bg-red-800 dark:text-red-400"
            >
              <svg
                class="flex-shrink-0 size-4"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <path d="M18 6 6 18"></path>
                <path d="m6 6 12 12"></path>
              </svg>
            </span>
            <!-- End Icon -->
          </div>
          <div class="ms-3">
            <h3 class="text-gray-800 font-semibold dark:text-white">Lỗi</h3>
            <p class="text-sm text-gray-700 dark:text-gray-400">
              Không tìm thấy danh mục
            </p>
          </div>
        </div>
      </div>
      <div class="flex justify-center mt-3">
        <Button :href="'/'"> Trang chủ</Button>
      </div>
    </div>
    <Content :header="header" v-else>
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
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  title: "Danh mục",
});
import { GET, POST } from "~/types/method";
import { type PostOverview, POST_PAGE_LIMIT } from "~/types/post";
const route = useRoute();

const showLoading = ref(false);
const foundCategory = ref(true);
const header = {
  title: `#${route.params.slug}`,
};
const categoryFounded = ref(null);
const paginationData = ref({
  total: 0,
  page: parseInt(route.query?.page) > 0 ? parseInt(route.query?.page) : 1,
  path: "/post/category/" + route.params.slug,
  limit: POST_PAGE_LIMIT,
});
const listNewPosts = ref<PostOverview[]>([]);
async function getCategoryBySlug() {
  showLoading.value = true;
  const category = await useRestApi(
    GET,
    "/category.json?slug=" + route.params.slug,
    {}
  );
  if (category?.status) {
    categoryFounded.value = category.data;
    setTimeout(() => {
      showLoading.value = false;
    }, 800);
  }
  foundCategory.value = category?.status;
}

async function getRelatedPosts() {
  if (categoryFounded.value?.id) {
    const postsData = await useRestApi(POST, "/category-posts.json", {
      category_id: categoryFounded.value?.id,
      page: paginationData.value.page,
      limit: POST_PAGE_LIMIT,
    });

    if (postsData.status) {
      paginationData.value.total = postsData.metadata?.total || 1;
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
    }
  }
}
const handlePaginate = async (page: string) => {
  let currentPage = parseInt(page);
  if (!currentPage) {
    return;
  }
  paginationData.value.page = currentPage;
  await getRelatedPosts();
};
await getCategoryBySlug();
await getRelatedPosts();
</script>

<style lang="scss" scoped></style>
