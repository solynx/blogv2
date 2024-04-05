interface PostOverview {
  title: string;
  slug: string;
  description: string | null;
  category: object | null;
  author: object | null;
  createdAt: string | null;
}

interface PostDetail {
  title: string;
  content: string | null;
  category: object | null;
  author: object | null;
  createdAt: string | null;
}

export const POST_PAGE_LIMIT = 2;
export type { PostOverview };
