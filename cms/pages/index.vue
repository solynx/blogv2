<template>
  <ClientOnly>
    <n-space vertical>
      <n-layout has-sider>
        <n-layout-sider
          bordered
          collapse-mode="width"
          :collapsed-width="64"
          :width="240"
          :collapsed="collapsed"
          show-trigger
          @collapse="collapsed = true"
          @expand="collapsed = false"
        >
          <n-menu
            v-model:value="activeKey"
            :collapsed="collapsed"
            :collapsed-width="64"
            :collapsed-icon-size="22"
            :options="menuOptions"
          />
        </n-layout-sider>
        <n-layout>
          <span>Content</span>
        </n-layout>
      </n-layout>
    </n-space>
  </ClientOnly>
</template>

<script lang="ts" setup>
import { NIcon } from "naive-ui";
import type { MenuOption } from "naive-ui";
import {
  BookOutline as BookIcon,
  PersonOutline as PersonIcon,
  WineOutline as WineIcon,
  PieChartOutline as PieChartIcon,
  NewspaperOutline as PaperIcon,
  ListOutline as ListIcon,
  AddOutline as AddIcon,
  Add,
} from "@vicons/ionicons5";

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions: MenuOption[] = [
  {
    label: "Overview",
    key: "overview",
    icon: renderIcon(PieChartIcon),
  },
  {
    label: "Quản trị tin tức",
    key: "blog-management",
    icon: renderIcon(PaperIcon),
    children: [
      {
        type: "group",
        label: "Bài viết",
        key: "post",
        children: [
          {
            label: "Quản trị bài viết",
            key: "post-all",
            icon: renderIcon(ListIcon),
          },
          {
            label: "Đăng bài",
            key: "upload-post",
            icon: renderIcon(AddIcon),
          },
        ],
      },
      {
        type: "group",
        label: "Tác giả",
        key: "author",
        children: [
          {
            label: "Danh sách tác giả",
            key: "author-all",
            icon: renderIcon(PersonIcon),
          },
          {
            label: "Thêm tác giả",
            key: "new-author",
            icon: renderIcon(AddIcon),
          },
        ],
      },
      {
        type: "group",
        label: "Danh mục",
        key: "category",
        children: [
          {
            label: "Tất cả danh mục",
            key: "category-all",
            icon: renderIcon(BookIcon),
          },
          {
            label: "Thêm danh mục",
            key: "new-category",
            icon: renderIcon(AddIcon),
          },
        ],
      },
    ],
  },
];

const activeKey = ref<string | null>(null);
const collapsed = ref(true);
</script>
