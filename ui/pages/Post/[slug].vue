<template>
  <div>
    <div  v-if="!foundPost">
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
              Bài viết không tồn tại
            </p>
          </div>
        </div>
      </div>
      <div class="flex justify-center mt-3">
        <Button :href="'/'"> Trang chủ</Button>
      </div>
    </div>
    <Content :header="header"   v-if="foundPost">
      <div class="grid md:gap-4">
        <div
          class="flex items-center flex-col justify-center my-3 relative"
          v-if="showLoading"
        >
          <Skeleton class="block" />
          <Spin class="absolute"></Spin>
        </div>
        <PostDetail :data="postDetail" v-if="!showLoading" />
      </div>
      <TextArea class="mt-3" @handle-create-contribute="handleCreateContribute"/>
      <div v-if="!showLoading">
        <h2 class="text-xl md:text-2xl font-segoe text-bold inline-block mt-3">
          Bài viết liên quan
        </h2>
        <div class="grid gap-4 grid-cols-1 md:grid-cols-4 pt-3">
          <PostSmallCart :data="relatedPostsData" />
        </div>
      </div>

    </Content>
  </div>
</template>

<script setup lang="ts">
import { GET, POST } from "~/types/method";
import { type PostDetail } from "~/types/post";
const header = {
  title: "Bài viết",
};
const postDetail = ref<PostDetail>({
  title: "",
  content: "",
  createdAt: "",
  author: {},
  category: {},
});
const foundPost = ref(false)
const relatedPostsData = ref(null);
const route = useRoute();
const showLoading = ref(false);
const contributeContent = ref("")
const getPostDetail = async () => {
  showLoading.value = true;
  let post = await useRestApi(GET, `/post.json?slug=${route.params?.slug}`);
  foundPost.value = post.status
  if (post.status) {
    let data = post.data;
    postDetail.value = {
      id: data.id,
      title: data.title || "",
      content: data.content || "",
      author: data.user,
      category: data.category,
      createdAt: useFormatDate(data.created_at),
    };
    await getRelatedPost();
  }
  setTimeout(() => {
    showLoading.value = false;
  }, 600);
};

async function getRelatedPost() {
    let relatedPosts = await useRestApi(POST, `/related-post.json`, {
    user_id: postDetail.value.author.id,
    category_id: postDetail.value.category.id,
    slug: route.params?.slug,
    });
    if (relatedPosts.status) {
      relatedPostsData.value = relatedPosts.data;
    }
}

async function handleCreateContribute(content: object, status: object, showMessage: object) {
  let result = await useRestApi(POST, `/contribute.json`, {
    content: content.value,
    post_id: postDetail.value.id,
    type: "post"
  });
  if(result.status) {
    status.value = true
    content.value = ""
    showMessage.value = true
  }
  else{
    status.value = false
  }
}
getPostDetail();

// getPostsRelated();
</script>

<style lang="scss" scoped></style>
