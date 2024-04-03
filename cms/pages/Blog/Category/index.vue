<template>
  <n-card title="Quản trị" hoverable>
    <ClientOnly>
      <NRow gutter="12">
      <n-col :span="12">
        <n-form-item label="Thêm danh mục">
          <n-input
            v-model:value="categoryName"
            show-count
            :maxlength="50"
            placeholder="Nhập tên danh mục..."
          />
          <NButton type="success" class="mx-3" @click="handleCreateCategory"
            >Thêm mới</NButton
          >
        </n-form-item>
      </n-col>
    </NRow>
      <n-data-table
        :columns="columns"
        :data="categoryTableData"
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
                <n-input type="text" placeholder="Waiting..." v-model:value="categoryDetail.name" />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Slug">
                <n-input type="text" placeholder="Waiting..." v-model:value="categoryDetail.slug" disabled />
              </NFormItem>
            </NGi>
            <NGi>
              <NFormItem label="Vị trí">
                  <n-input-number
                  v-model:value="categoryDetail.index"
                  placeholder="Waiting..."
                  :min="1"
                  :max="99"
                  disable
                />
              </NFormItem>
          
            </NGi>
            <NGi>
              <NFormItem label="Người tạo">
                <n-input type="text" placeholder="Waiting..." disabled v-model:value="categoryDetail.author"/>
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
import { type CategoryDetail, CATEGORY_ENDPOINT, CATEGORY_DETAIL_ENDPOINT, type CategoryDataTable } from "~/types/category";
import { Method } from "~/types/requestMethod";
import { formatDate } from "~/helpers/date";
import {
  EyeSharp as EyeIcon,
  PencilSharp as PencilIcon,
  TrashOutline as TrashIcon,
} from "@vicons/ionicons5";

const showDetailCategory = ref(false);
const showDeleteConfirm = ref(false);
const categoryName = ref<String>("");
const pagination = {
  pageSize: 12,
};
const categoryTableData = ref<CategoryDataTable[]>([]);
const message = useMessage();
const loadingBar = useLoadingBar()
const categoryDetail = ref<CategoryDetail>({id: "", name: "", slug: "", index: 1, author: ""});
const createColumns = ({
  view,
  edit,
  remove,
}: {
  view: (row: CategoryDataTable) => void;
  edit: (row: CategoryDataTable) => void;
  remove: (row: CategoryDataTable) => void;
}): DataTableColumns<CategoryDataTable> => {
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
  view(row: CategoryDataTable) {
    handleGetDetailCategory(row.value)
    showDetailCategory.value = !showDetailCategory.value;
  },
  edit(row: CategoryDataTable) {
    handleGetDetailCategory(row.value)
    showDetailCategory.value = !showDetailCategory.value;
  },
  remove(row: CategoryDataTable) {
    showDeleteConfirm.value = !showDeleteConfirm.value
  },
});

const convertDataForTable = async () => {
  const listCategory = await useRestApi(Method.GET, CATEGORY_ENDPOINT, {});
  categoryTableData.value = [];
  listCategory.data.forEach((item: object, index: number) => {
    categoryTableData.value.push({
      no: index + 1,
      title: item?.name,
      author: item.user?.full_name ?? "",
      createdAt: formatDate(item.create_at),
      value: item.id
    });
  });
};
const handleGetDetailCategory = async (id : string) => {
  let result = await useRestApi(Method.POST, CATEGORY_DETAIL_ENDPOINT, {id});
  if (result.data) {
    categoryDetail.value = {
      id: id,
      name: result.data.name,
      slug: result.data.slug,
      index: result.data.index,
      author: result.data.user.full_name
    }
  }
}
const handleUpdateCategory = async () => {
  loadingBar.start();
  let result = await useRestApi(Method.PATCH, CATEGORY_ENDPOINT, categoryDetail.value);
  showDetailCategory.value = !showDetailCategory.value
  if(result.status) {
    loadingBar.finish();
    message.success("Cập nhật thành công!", {duration: 1000});
    return convertDataForTable();
  }
  else{
    loadingBar.error();
    return message.error("Thêm thất bại")
  }
}
const onConfirmDeleteClick = async (id : string) => {
  let result = await useRestApi(Method.DELETE, CATEGORY_ENDPOINT, {id});
  if (result.status) {
    message.success("Xóa thành công!", {duration: 1000});
    return convertDataForTable();
  }
  return message.error("Xóa thất bại")
}
const handleCreateCategory = async () => {
  let result = await useRestApi(Method.POST, CATEGORY_ENDPOINT, {name: categoryName.value});
  categoryName.value = "";
  if (result.status) {
    message.success("Thêm thành công!", {duration: 1000});
    return convertDataForTable();
  }
  return message.error("Xóa thất bại")
}
convertDataForTable();
</script>

<style lang="scss" scoped></style>
