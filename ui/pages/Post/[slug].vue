<template>
  <div>
    <Content :header="header">
      <hr class="md:my-6 my-3" />
      <div class="grid md:gap-4">
        <PostDetail />
      </div>
      <Pagination class="mt-5" />
    </Content>
  </div>
</template>

<script setup>
const header = {
  title: "Post detail",
};
</script>

<style lang="scss" scoped></style>
