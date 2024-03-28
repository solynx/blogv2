<template>
    <Content :header="header">
      <div class="my-6">
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
    title : "Danh Mục Bài Viết"
  }
  </script>