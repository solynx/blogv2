<template>
  <div>
    <ClientOnly>
      <NCard title="Write Post" hoverable>
      <n-space vertical class="mb-3">
        <n-grid :x-gap="12" :y-gap="12" :cols="4" layout-shift-disabled>
          <n-gi :span="2">
            <NFormItem label="Title">
              <n-input
                v-model:value="post.title"
                type="text"
                placeholder="Title"
              />
            </NFormItem>
          </n-gi>

          <n-gi>
            <NFormItem label="Category">
              <n-select
                v-model:value="post.category_id"
                filterable
                placeholder="Chọn danh mục"
                :options="categoryOptions"
              />
            </NFormItem>
          </n-gi>
        </n-grid>
        <NFormItem label="Description">
          <n-input
            v-model:value="post.description"
            type="text"
            placeholder="Description"
          />
        </NFormItem>
      </n-space>
      <NFormItem label="Content">
        <TiptapEditor :editor="editor" class="w-full" />
      </NFormItem>
      <NButton type="info" @click="handlePostUpload">Upload</NButton>
    </NCard>
    </ClientOnly>
  </div>
</template>

<script setup lang="ts">
import { useMessage, useLoadingBar } from "naive-ui";
import  { type PostUploadData, POST_ENDPOINT} from "~/types/post";
import {  CATEGORY_ENDPOINT, type CategoryOption } from "~/types/category";
import { Method } from "~/types/requestMethod";
const value = ref("");
const message = useMessage();
const loadingBar = useLoadingBar()
const selectedValue = ref(null);
const options = [
  {
    label: "Drive My Car",
    value: "song1",
  },
  {
    label: "Norwegian Wood",
    value: "song2",
  },
  {
    label: "You Won't See",
    value: "song3",
  },
  {
    label: "Nowhere Man",
    value: "song4",
  },
  {
    label: "Think For Yourself",
    value: "song5",
  },
  {
    label: "The Word",
    value: "song6",
  },
  {
    label: "Michelle",
    value: "song7",
  },
  {
    label: "What goes on",
    value: "song8",
  },
  {
    label: "Girl",
    value: "song9",
  },
  {
    label: "I'm looking through you",
    value: "song10",
  },
  {
    label: "In My Life",
    value: "song11",
  },
  {
    label: "Wait",
    value: "song12",
  },
];

const post = ref<PostUploadData>({
  title: "",
  description: "",
  content: "",
  category_id: "",
});
const editor = useEditor({
  content: post.value.content,
  extensions: [
    TiptapStarterKit,
    TiptapPlaceholder.configure({
      placeholder: "Write something …",
    }),
  ],
  editorProps: {
    attributes: {
      class: "border-[1px] border-solid border-inherit px-3",
    },
  },
  onUpdate: (value) => {
    post.value.content = value.editor.getHTML();
  },
});

const categoryOptions = ref<CategoryOption[]>([])
const handleGetLi = async () => {
  const listCategory = await useRestApi(Method.GET, CATEGORY_ENDPOINT, {});
  if(listCategory.status) {
    listCategory.data.forEach((item: object, index: number) => {
    categoryOptions.value.push({
      label: item.name,
      value: item.id
      });
    });
    return;
  }
  return message.error('Failed to get data')
}
const handlePostUpload = async () => {
  loadingBar.start();
  if(!post.value.category_id) {
    return message.error("Vui lòng chọn danh mục")
  }
  const result = await useRestApi(Method.POST, POST_ENDPOINT, post.value)
  if(result.status) {
    loadingBar.finish();
    post.value = {
        title: "",
        description: "",
        content: "",
        category_id: "",
    }
    return message.success("Thêm bài viết thành công!")
  }
  loadingBar.error();
  return message.error("Lỗi thêm bài viết!")
};
handleGetLi();
</script>
