<template>
  <NModalProvider>
    <NMessageProvider>
      <div class="relative h-screen">
        <n-layout position="absolute">
          <n-layout-header
            style="height: 64px; padding: 24px; display: flex"
            bordered
            class="bg-gradient-to-r from-[#000428] to-[#004e92]"
          >
            <div style="width: 240px">
              <img
                src="https://api.dicebear.com/8.x/identicon/svg?seed=Coco"
                alt="avatar"
                class="w-6 h-6"
              />
            </div>
            <div class="text-white">HEADER</div>
          </n-layout-header>
          <n-layout
            has-sider
            position="absolute"
            style="top: 64px; bottom: 64px"
          >
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
                </n-layout>
              </n-space>
            </ClientOnly>
            <n-layout embedded content-style="padding: 24px;">
              <slot />
            </n-layout>
          </n-layout>
          <n-layout-footer
            bordered
            position="absolute"
            style="height: 64px; padding: 24px"
            class="bg-gradient-to-l from-[#000428] to-[#004e92]"
          >
            Footer
          </n-layout-footer>
        </n-layout>
      </div>
    </NMessageProvider>
  </NModalProvider>
</template>

<script lang="ts" setup>
import { NIcon, NMessageProvider, NModalProvider } from "naive-ui";
import type { MenuOption } from "naive-ui";
import MyNuxtLink from "~/components/MyNuxtLink.ts";
import {
  BookOutline as BookIcon,
  PersonOutline as PersonIcon,
  WineOutline as WineIcon,
  PieChartOutline as PieChartIcon,
  NewspaperOutline as PaperIcon,
  ListOutline as ListIcon,
  AddOutline as AddIcon,
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
        label: "Bài viết & Danh mục",
        key: "post",
        children: [
          {
            label: () =>
              h(
                MyNuxtLink,
                {
                  to: "/blog/post",
                },
                "Quản trị bài viết"
              ),
            key: "post-all",
            icon: renderIcon(ListIcon),
          },
          {
            label: () =>
              h(
                MyNuxtLink,
                {
                  to: "/blog/post/create",
                },
                "Đăng bài"
              ),
            key: "create-post",
            icon: renderIcon(AddIcon),
          },
          {
            label: () =>
              h(
                MyNuxtLink,
                {
                  to: "/blog/category",
                },
                "Danh mục"
              ),
            key: "category-all",
            icon: renderIcon(BookIcon),
          },
        ],
      },
      {
        type: "group",
        label: "Tác giả",
        key: "author",
        children: [
          {
            label: () =>
              h(
                MyNuxtLink,
                {
                  to: "/blog/author",
                },
                "Tất cả"
              ),
            key: "author-all",
            icon: renderIcon(PersonIcon),
          },
          {
            label: () =>
              h(
                MyNuxtLink,
                {
                  to: "/blog/author/create",
                },
                "Thêm mới"
              ),
            key: "new-author",
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
