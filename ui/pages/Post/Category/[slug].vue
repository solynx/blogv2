<template>
    <div>
      <Content :header="header">
        <div class="grid md:gap-4">
            <!-- <PostItem :data="listNewPosts"/> -->
        </div>
      </Content>
    </div>
</template>
  
<script setup lang="ts">

const route = useRoute()
const header = {
title: `#${route.params.slug}`,
};
// const listNewPosts = [
//     {title: "as"}
// ]
// const getNewPosts = async () => {
//   const postsData = await useRestApi(POST, "/posts.json", { page: 1, limit: 6})
//   if (postsData.status) {
//     postsData.data.forEach((item, index) => {
//       let post : PostOverview = {
//         title: item.title || "",
//         description: item.description || "",
//         slug: item.slug || item.id,
//         createdAt: useFormatDate(item.created_at),
//         author: item.user,
//         category: item.category
//       }
//       listNewPosts.value.push(post)
//     });
//   }
// }
</script>
  
<style lang="scss" scoped></style>
  