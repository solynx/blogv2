interface PostUploadData {
  title: string;
  description: string;
  content: string;
  category_id: string;
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
  content: string,
  seoTitle: string,
  seoKeyword: string, 
  seoDescription: string
}
export const POST_ENDPOINT = "/v1/system/post.json";
export type { PostUploadData, PostDataTable, PostDataDetail };
