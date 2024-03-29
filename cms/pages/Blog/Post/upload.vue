<template>
  <div>
    
    <NCard title="Write Post" hoverable>
      <n-space vertical class="mb-3">
        <NCard title="Tiêu đề" :bordered="false">
            <n-input v-model:value="value" type="text" placeholder="Title" />
        </NCard>
        <NCard title="Mô tả ngắn" :bordered="false">
            <n-input v-model:value="value" type="text" placeholder="Description" />
        </NCard>

      </n-space>
      <NCard title="Nội dung " :bordered="false">
        <Tiptap />
      </NCard>
      <template #footer>
       <NButton>Upload</NButton>
    </template>
    </NCard>
  </div>
</template>

<script setup>
const value = ref('')
</script>

<style lang="scss" scoped></style>
