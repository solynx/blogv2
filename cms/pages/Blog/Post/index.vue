<template>
  <n-card title="Quản trị" hoverable>
    <ClientOnly>
      <n-data-table
        :columns="columns"
        :data="postTableData"
        :bordered="false"
        :pagination="pagination"
      />

      <n-modal v-model:show="showDetailPost" transform-origin="center">
        <n-card
          style="position: fixed; width: 80%; top: 10%; left: 10%"
          title="Sửa danh mục"
          :bordered="false"
          size="huge"
          role="dialog"
          aria-modal="true"
        >
          <NGrid x-gap="12" :cols="4">
            <NGi span="2">
              <NFormItem label="Tiêu đề bài viết">
                <n-input
                  type="text"
                  placeholder="Waiting..."
                  v-model:value="postDataDetail.title"
                />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Slug">
                <n-input
                  type="text"
                  placeholder="Waiting..."
                  v-model:value="postDataDetail.slug"
                  disabled
                />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Người tạo">
                <n-input
                  type="text"
                  placeholder="Waiting..."
                  disabled
                  v-model:value="postDataDetail.author"
                />
              </NFormItem>
            </NGi>
            <NGi span="2">
              <NFormItem label="Mô tả ngắn">
                <n-input
                  type="text"
                  placeholder="Waiting..."
                  v-model:value="postDataDetail.description"
                />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Danh mục">
                <n-select
                  v-model:value="postDataDetail.category_id"
                  filterable
                  placeholder="Chọn danh mục"
                  :options="categoryOptions"
                />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Trạng thái">
                <n-switch v-model:value="postDataDetail.published" />
              </NFormItem>
            </NGi>
            <NGi span="4">
              <NFormItem label="Content">
              <TiptapEditor :editor="editor" class="w-full" />
            </NFormItem>
            </NGi>
            <NGi span="2">
              <NFormItem label="Tiêu đề SEO">
                <n-input
                  type="text"
                  placeholder="Nhập tiêu đề..."
                  v-model:value="postDataDetail.seo_title"
                />
              </NFormItem>
            </NGi>
            <NGi span="2">
              <NFormItem label="Từ khóa SEO">
                <n-input
                  type="text"
                  placeholder="Nhập từ khóa..."
                  v-model:value="postDataDetail.seo_keywords"
                />
              </NFormItem>
            </NGi>
            <NGi span="4">
              <NFormItem label="Mô tả SEO">
                <n-input
                  type="textarea"
                  placeholder="Nhập mô tả..."
                  v-model:value="postDataDetail.seo_description"
                />
              </NFormItem>
            </NGi>

          </NGrid>
          <template #footer>
              <NButton type="info" @click="handleUpdatePost">Lưu thông tin</NButton>
          </template>
        </n-card>
      </n-modal>
      <n-modal
        v-model:show="showDeleteConfirm"
        :mask-closable="false"
        preset="dialog"
        title="Dialog"
        content="Are you sure?"
        positive-text="Confirm"
        negative-text="Cancel"
        @positive-click="onConfirmDeleteClick"
      />
    </ClientOnly>
  </n-card>
</template>

<script setup lang="ts">
import {
  NButton,
  NIcon,
  useMessage,
  NPopconfirm,
  useLoadingBar,
} from "naive-ui";
import type { DataTableColumns, NInputGroup } from "naive-ui";
import {
  POST_DETAIL_ENDPOINT,
  POST_ENDPOINT,
  type PostDataTable,
  type PostDataDetail,
} from "~/types/post";
import { type CategoryOption, CATEGORY_ENDPOINT } from "~/types/category";
import { Method } from "~/types/requestMethod";
import { formatDate } from "~/helpers/date";
import {
  EyeSharp as EyeIcon,
  PencilSharp as PencilIcon,
  TrashOutline as TrashIcon,
} from "@vicons/ionicons5";

const showDetailPost = ref(false);
const showDeleteConfirm = ref(false);
const pagination = {
  pageSize: 12,
};
const postTableData = ref<PostDataTable[]>([]);
const categoryOptions = ref<CategoryOption[]>([]);
const message = useMessage();
const loadingBar = useLoadingBar();

const postDataDetail = ref<PostDataDetail>({
  id: "",
  title: "",
  slug: "",
  description: "",
  content: "",
  author: "",
  category_id: null,
  seo_title: null,
  seo_keywords: null,
  seo_description: null,
  published: false
});
const createColumns = ({
  view,
  edit,
  remove,
}: {
  view: (row: PostDataTable) => void;
  edit: (row: PostDataTable) => void;
  remove: (row: PostDataTable) => void;
}): DataTableColumns<PostDataTable> => {
  return [
    {
      title: "No",
      key: "no",
    },
    {
      title: "Tiêu đề",
      key: "title",
    },
    {
      title: "Danh muc",
      key: "category",
    },
    {
      title: "Tác giả",
      key: "author",
    },
    {
      title: "Ngày tạo",
      key: "createdAt",
    },
    {
      title: "Action",
      key: "actions",
      render(row: any) {
        return [
          h(
            NButton,
            {
              info: true,
              type: "warning",
              size: "medium",
              class: "mx-1",
              title: "Edit",
              onClick: () => edit(row),
            },
            { default: () => h(NIcon, null, { default: () => h(PencilIcon) }) }
          ),
          h(
            NPopconfirm,
            {
              onPositiveClick: () => onConfirmDeleteClick(row.id),
              negativeText: "Cancel",
              positiveText: "Delete",
            },
            {
              trigger: () =>
                h(
                  NButton,
                  { size: "medium", type: "error" },
                  {
                    default: () =>
                      h(NIcon, null, { default: () => h(TrashIcon) }),
                  }
                ),
              default: () => h("span", {}, { default: () => "Xác nhận xóa" }),
            }
          ),
        ];
      },
    },
  ];
};

const columns = createColumns({
  view(row: PostDataTable) {
    // handleGetDetailCategory(row.value)
    showDetailPost.value = !showDetailPost.value;
  },
  edit(row: PostDataTable) {
    handleGetDetailPost(row.id);
    showDetailPost.value = !showDetailPost.value;
  },
  remove(row: PostDataTable) {
    showDeleteConfirm.value = !showDeleteConfirm.value;
  },
});

const editor = useEditor({
  content: postDataDetail.value.content,
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
    postDataDetail.value.content = value.editor.getHTML();
  },
});

const convertDataForTable = async () => {
  const listPost = await useRestApi(Method.GET, POST_ENDPOINT, {});
  postTableData.value = [];
  listPost.data.forEach((item: object, index: number) => {
    postTableData.value.push({
      no: index + 1,
      title: item?.title,
      author: item.user?.full_name ?? "",
      createdAt: formatDate(item.created_at),
      id: item.id,
      category_id: item.category.id,
      category: item.category.name,
      seo_title: item.seo_title,
      seo_keywords: item.seo_keywords,
      seo_description: item.seo_description,
      published: item.published ?? false,
    });
  });
};
const handleGetDetailPost = async (id: string) => {
  let result = await useRestApi(
    Method.GET,
    POST_DETAIL_ENDPOINT + `?id=${id}`,
    {}
  );

  if (result.status) {
    let data = result.data;
    postDataDetail.value = {
      id: id,
      title: data.title,
      slug: data.slug,
      description: data.description,
      content: data.content,
      author: data.user.full_name,
      published: data.published ?? false,
      seo_description: data.seo_description,
      seo_title: data.seo_title,
      seo_keywords: data.seo_keywords,
      category_id: data.category.id,
    };
    editor.value.commands.setContent(postDataDetail.value.content)
  }
};

const handleGetListCategory = async () => {
  const listCategory = await useRestApi(Method.GET, CATEGORY_ENDPOINT, {});
  if (listCategory.status) {
    listCategory.data.forEach((item: object, index: number) => {
      categoryOptions.value.push({
        label: item.name,
        value: item.id,
      });
    });
    return;
  }
  return message.error("Failed to get data");
};

const handleUpdatePost = async () => {
  loadingBar.start();
  let result = await useRestApi(Method.PATCH, POST_ENDPOINT, postDataDetail.value);
  showDetailPost.value = !showDetailPost.value
  if(result.status) {
    loadingBar.finish();
    message.success("Cập nhật thành công!", {duration: 1000});
    return convertDataForTable();
  }
  else{
    loadingBar.error();
    return message.error("Cập nhật không thành công!")
  }
}
const onConfirmDeleteClick = async (id : string) => {
  let result = await useRestApi(Method.DELETE, POST_ENDPOINT, {id});
  if (result.status) {
    message.success("Xóa thành công!", {duration: 1000});
    return convertDataForTable();
  }
  return message.error("Xóa thất bại")
}
convertDataForTable();
handleGetListCategory();
</script>

<style lang="scss" scoped></style>
