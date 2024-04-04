interface PostUploadData {
  title: string;
  description: string;
  content: string;
  category_id: string;
  published: boolean;
  seo_title: string;
  seo_description: string | null;
  seo_keywords: string | null;
}
interface PostDataTable {
  no: number;
  title: string;
  category: string,
  author: string;
  createdAt: string;
  id: string;
};

interface PostDataDetail {
  id: string, 
  title: string, 
  slug: string,
  description: string,
  author: string,
  content: string,
  published: boolean,
  category_id: string | null,
  seo_title: string | null,
  seo_keywords: string | null,
  seo_description: string | null,
}
export const POST_ENDPOINT = "/v1/system/post.json";
export const POST_DETAIL_ENDPOINT = "/v1/system/post/detail.json";
export type { PostUploadData, PostDataTable, PostDataDetail };
