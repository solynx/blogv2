
import type { DataTableColumns } from "naive-ui";
interface CategoryDataTable {
    no: number;
    title: string;
    author: string;
    createdAt: string;
    value: string;
};

interface CategoryDetail {
    id: string,
    name: string,
    author: string,
    slug: string,
    index: number,
}

interface CategoryOption {
    label: string,
    value: string
}
export const CATEGORY_ENDPOINT = "/v1/system/category.json";
export const CATEGORY_DETAIL_ENDPOINT = "/v1/system/category/detail.json";

export type {CategoryDataTable, CategoryDetail, CategoryOption}