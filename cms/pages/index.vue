<template>
  <n-page-header @back="handleBack">
    <n-grid :cols="5">
      <n-gi>
        <n-statistic label="Tổng bài viết" value="125" />
      </n-gi>
      <n-gi>
        <n-statistic label="Tổng bài viết hiển thị" value="22" />
      </n-gi>
      <n-gi>
        <n-statistic label="Bài viết trong tháng" value="36" />
      </n-gi>
      <n-gi>
        <n-statistic label="Danh mục" value="83" />
      </n-gi>
      <n-gi>
        <n-statistic label="Số lượt xem" value="2,346" />
      </n-gi>
    </n-grid>
    <template #title> Thống kê </template>
    <template #extra>
      <n-space>
        <n-button>Làm mới</n-button>
      </n-space>
    </template>
  </n-page-header>
</template>

<script setup>
import { useMessage } from "naive-ui";
const message = useMessage();
const route = useRoute();
const handleBack = () => {
  message.info("[onBack]");
};
const options = [
  {
    label: "More episodes",
    key: "1",
  },
  {
    label: "More episodes",
    key: "2",
  },
  {
    label: "More episodes",
    key: "3",
  },
];
</script>

<style lang="scss" scoped></style>
