<template>
  <n-card title="Quản trị" hoverable>
    <n-data-table
      :columns="columns"
      :data="data"
      :bordered="false"
      :pagination="pagination"
    />
  </n-card>
</template>

<script setup lang="ts">
import { NButton, NIcon, useMessage } from "naive-ui";
import type { DataTableColumns } from "naive-ui";
import {
  EyeSharp as EyeIcon,
  PencilSharp as PencilIcon,
  TrashOutline as TrashIcon,
} from "@vicons/ionicons5";
type Song = {
  no: number;
  title: string;
  author: string;
  createdAt: string;
};
const pagination = {
  pageSize: 10,
};
const createColumns = ({
  play,
}: {
  play: (row: Song) => void;
}): DataTableColumns<Song> => {
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
      title: "Ngày đăng",
      key: "createdAt",
    },
    {
      title: "Action",
      key: "actions",
      render(row) {
        return [
          h(
            NButton,
            {
              info: true,
              color: "#2080F0",
              size: "medium",
              class: "mr-2",
              title: "Detail",
              onClick: () => play(row),
            },
            { default: () => h(NIcon, null, { default: () => h(EyeIcon) }) }
          ),
          h(
            NButton,
            {
              info: true,
              color: "#F0A020",
              size: "medium",
              class: "mx-2",
              title: "Edit",
              onClick: () => play(row),
            },
            { default: () => h(NIcon, null, { default: () => h(PencilIcon) }) }
          ),
          h(
            NButton,
            {
              info: true,
              color: "#D03050",
              class: "mx-2",
              size: "medium",
              title: "Delete",
              onClick: () => play(row),
            },
            { default: () => h(NIcon, null, { default: () => h(TrashIcon) }) }
          ),
        ];
      },
    },
  ];
};

const data: Song[] = [
  { no: 3, title: "Wonderwall", author: "John Doe", createdAt: "12/12/2024" },
  {
    no: 4,
    title: "Don't Look Back in Anger",
    author: "Artisan",
    createdAt: "12/12/2024",
  },
  {
    no: 12,
    title: "Champagne Supernova",
    author: "Mkdir",
    createdAt: "12/12/2024",
  },
];
const message = useMessage();
const columns = createColumns({
  play(row: Song) {
    message.info(`Play ${row.title}`);
  },
});
</script>

<style lang="scss" scoped></style>
