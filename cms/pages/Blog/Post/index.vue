<template>
  <n-card title="Quản trị" hoverable>
    <ClientOnly>
      <n-data-table
        :columns="columns"
        :data="postTableData"
        :bordered="false"
        :pagination="pagination"
      />

      <n-modal v-model:show="showDetailCategory" transform-origin="center">
        <n-card
          style="position: fixed; width: 60%; top: 30%; left: 20%"
          title="Sửa danh mục"
          :bordered="false"
          size="huge"
          role="dialog"
          aria-modal="true"
        >
          <NGrid x-gap="12" :cols="2">
            <NGi>
              <NFormItem label="Tên danh mục">
                <n-input type="text" placeholder="Waiting..." v-model:value="PostDataDetail.name" />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Slug">
                <n-input type="text" placeholder="Waiting..." v-model:value="PostDataDetail.slug" disabled />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Vị trí">
                  <n-input-number
                  v-model:value="PostDataDetail.index"
                  placeholder="Waiting..."
                  :min="1"
                  :max="99"
                  disable
                />
              </NFormItem>
          
            </NGi>
            <NGi>
              <NFormItem label="Người tạo">
                <n-input type="text" placeholder="Waiting..." disabled v-model:value="PostDataDetail.author"/>
              </NFormItem>
            </NGi>  
          </NGrid>
          <template #footer>
            <NButton type="info" @click="handleUpdateCategory">Sửa</NButton>
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
import { NButton, NIcon, useMessage, NPopconfirm, useLoadingBar } from "naive-ui";
import type { DataTableColumns, NInputGroup } from "naive-ui";
import { POST_ENDPOINT, type PostDataTable, type PostDataDetail } from "~/types/post";
import { Method } from "~/types/requestMethod";
import { formatDate } from "~/helpers/date";
import {
  EyeSharp as EyeIcon,
  PencilSharp as PencilIcon,
  TrashOutline as TrashIcon,
} from "@vicons/ionicons5";

const showDetailCategory = ref(false);
const showDeleteConfirm = ref(false);
const pagination = {
  pageSize: 12,
};
const postTableData = ref<PostDataTable[]>([]);
const message = useMessage();
const loadingBar = useLoadingBar()
const postDataDetail = ref<PostDataDetail>({id: "", name: "", slug: "", index: 1, author: ""});
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
            { onPositiveClick: () => onConfirmDeleteClick(row.value), negativeText: 'Cancel', positiveText: 'Delete',},
            { trigger: () => h(NButton, 
                  {size: 'medium', type: "error"},
                  { default: () => h(NIcon, null, { default: () => h(TrashIcon)}) },),
                  default: () => h('span', {}, { default: () => 'Xác nhận xóa' }),},),
        ];
      },
    },
  ];
};


const columns = createColumns({
  view(row: PostDataTable) {
    // handleGetDetailCategory(row.value)
    showDetailCategory.value = !showDetailCategory.value;
  },
  edit(row: PostDataTable) {
    // handleGetDetailCategory(row.value)
    showDetailCategory.value = !showDetailCategory.value;
  },
  remove(row: PostDataTable) {
    showDeleteConfirm.value = !showDeleteConfirm.value
  },
});

const convertDataForTable = async () => {
  const listPost = await useRestApi(Method.GET, POST_ENDPOINT, {});
  postTableData.value = [];
  listPost.data.forEach((item: object, index: number) => {
    console.log(item)
    postTableData.value.push({
      no: index + 1,
      title: item?.title,
      author: item.user?.full_name ?? "",
      createdAt: formatDate(item.create_at),
      id: item.id,
      category: item.category.name,
      seoTitle: item.seo_title,
      seoDescription: item.seo_description
    });
  });
};
// const handleGetDetailCategory = async (id : string) => {
//   let result = await useRestApi(Method.POST, CATEGORY_DETAIL_ENDPOINT, {id});
//   if (result.data) {
//     PostDataDetail.value = {
//       id: id,
//       name: result.data.name,
//       slug: result.data.slug,
//       index: result.data.index,
//       author: result.data.user.full_name
//     }
//   }
// }
// const handleUpdateCategory = async () => {
//   loadingBar.start();
//   let result = await useRestApi(Method.PATCH, CATEGORY_ENDPOINT, PostDataDetail.value);
//   showDetailCategory.value = !showDetailCategory.value
//   if(result.status) {
//     loadingBar.finish();
//     message.success("Cập nhật thành công!", {duration: 1000});
//     return convertDataForTable();
//   }
//   else{
//     loadingBar.error();
//     return message.error("Thêm thất bại")
//   }
// }
// const onConfirmDeleteClick = async (id : string) => {
//   let result = await useRestApi(Method.DELETE, CATEGORY_ENDPOINT, {id});
//   if (result.status) {
//     message.success("Xóa thành công!", {duration: 1000});
//     return convertDataForTable();
//   }
//   return message.error("Xóa thất bại")
// }
convertDataForTable();
</script>

<style lang="scss" scoped></style>
