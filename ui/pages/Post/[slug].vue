<template>
  <div>
    <Content :header="header">
      <div class="grid md:gap-4">
        <PostDetail :data="postDetail"/>
      </div>
      <div >
        <h2 class="text-3xl font-segoe text-bold inline-block mt-3">
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
import { GET, POST } from "~/types/method"
import { type PostDetail } from "~/types/post"
const header = {
  title: "Bài viết",
};
const postDetail = ref<PostDetail>({title: "", content: "", createdAt: "", author: {}, category: {}})
const relatedPostsData = ref(null)
const route = useRoute()
const getPostDetail = async () => {
  let post = await useRestApi(GET, `/post.json?slug=${route.params?.slug}`)
  if (post.status) {
    let data = post.data
    postDetail.value = {
      title: data.title || '',
      content: data.content || '',
      author: data.user,
      category: data.category,
      createdAt: useFormatDate(data.created_at)
    }
  }
}

let relatedPosts = await useRestApi(POST, `/related-post.json`, { user_id: postDetail.value.author.id, category_id: postDetail.value.category.id, slug: route.params?.slug})
if (relatedPosts.status) {
  relatedPostsData.value = relatedPosts.data
}
getPostDetail();
// getPostsRelated();
</script>

<style lang="scss" scoped></style>
