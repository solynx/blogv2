<template>
  <Content :header="header">
    <hr class="md:my-6 my-3" />
      <div class="grid md:gap-4">
        <PostItem />
        <PostItem />
      </div>
      <Pagination class="mt-5" />
    <div class="mx-3 my-6">
      <div class="my-3">
        <h2 class="text-3xl font-extrabold text-[#333] inline-block">Category</h2>
      </div>
      <div class="block md:grid md:grid-cols-3 md:gap-3">
        <Category />
        <Category />
        <Category />
        <Category />
        <Category />
        <Category />
      </div>
    </div>
  </Content>
</template>
<script setup>
const header = {
  title : "News",
  description: 'Lastest Post'
}
</script>