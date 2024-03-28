<template>
  <div>
    <Content :header="header">
      <div class="grid md:gap-4">
        <PostDetail />
      </div>
      <div>
        <h2 class="text-3xl font-segoe text-bold inline-block mt-3">
          Bài viết liên quan
        </h2>
        <div class="grid gap-4 grid-cols-1 md:grid-cols-4 pt-3">
          <PostSmallCart/>
          <PostSmallCart/>
          <PostSmallCart/>
          <PostSmallCart/>
        </div>
      </div>
    </Content>
  </div>
</template>

<script setup>
const header = {
  title: "Tìm hiểu công cụ Github",
};
</script>

<style lang="scss" scoped></style>
