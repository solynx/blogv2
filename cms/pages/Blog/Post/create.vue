<template>
  <div>
    <ClientOnly>
      <NCard title="Write Post" hoverable>
      <n-space vertical class="mb-3">
        <n-grid :x-gap="12" :y-gap="12" :cols="4" layout-shift-disabled>
          <n-gi :span="3">
            <NFormItem label="Tiêu đề">
              <n-input
                v-model:value="post.title"
                type="text"
                placeholder="Nhập tiêu đề..."
              />
            </NFormItem>
          </n-gi>

          <n-gi>
            <NFormItem label="Danh mục">
              <n-select
                v-model:value="post.category_id"
                filterable
                placeholder="Chọn danh mục"
                :options="categoryOptions"
              />
            </NFormItem>
          </n-gi>

          <NGi span="3">
            <NFormItem label="Mô tả ngắn">
              <n-input
                v-model:value="post.description"
                type="text"
                placeholder="Nhập mô tả..."
              />
            </NFormItem>
          </NGi>
          <n-gi >
            <NFormItem label="Trạng thái" class="justify-center">
              <n-switch
                v-model:value="post.published"
              />
            </NFormItem>
          </n-gi>
          <NGi span="4">
            <NFormItem label="Nội dung chính">
              <TiptapEditor :editor="editor" class="w-full" />
            </NFormItem>
          </NGi>
          <NGi span="2">
            <NFormItem label="Tiêu đề SEO">
              <n-input
                v-model:value="post.seo_title"
                type="text"
                placeholder="Nhập tiêu đề..."
              />
            </NFormItem>
          </NGi>
          <NGi span="2">
            <NFormItem label="Từ khóa SEO">
              <n-input
                v-model:value="post.seo_keywords"
                type="text"
                placeholder="Nhập từ khóa..."
              />
            </NFormItem>
          </NGi>
          <NGi span="4">
            <NFormItem label="Mô tả SEO">
              <n-input
                v-model:value="post.seo_description"
                type="textarea"
                placeholder="Nhập mô tả..."
              />
            </NFormItem>
          </NGi>
        </n-grid>
        
      </n-space>

      <NButton type="info" @click="handlePostUpload">Thêm bài viết</NButton>
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


const post = ref<PostUploadData>({
  title: "",
  description: "",
  content: "",
  category_id: null,
  published: false,
  seo_title: null,
  seo_description: null,
  seo_keywords: null
});

const editor = useEditor({
  content: post.value.content,
  extensions: [
    TiptapStarterKit,
    TiptapPlaceholder.configure({
      placeholder: "Write something …",
    }),
    TiptapImage.configure({
      inline: true,
    }),
    TiptapTextAlign.configure({
    types: ['heading', 'paragraph'],}),
    TiptapYoutube.configure({
          controls: false,
        })
  ],
  editorProps: {
    attributes: {
      class: "border-[1px] border-solid border-inherit px-3",
    },
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
  if(!post.value.category_id) {
    return message.error("Vui lòng chọn danh mục")
  }
  post.value.content = editor.value.getHTML()
  loadingBar.start();
  const result = await useRestApi(Method.POST, POST_ENDPOINT, post.value)
  if(result.status) {
    loadingBar.finish();
    post.value = {
        title: "",
        description: "",
        content: "",
        category_id: null,
        published: false,
        seo_description: null,
        seo_title: null,
        seo_keywords: null
    }
    editor.value.commands.clearContent()
    return message.success("Thêm bài viết thành công!")
  }
  loadingBar.error();
  return message.error("Lỗi thêm bài viết!")
};
handleGetLi();
</script>
