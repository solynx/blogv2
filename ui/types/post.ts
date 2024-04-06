interface PostOverview {
  title: string;
  slug: string;
  description: string | null;
  category: object | null;
  author: object | null;
  createdAt: string | null;
}

interface PostDetail {
  id: string | null,
  title: string;
  content: string | null;
  category: object | null;
  author: object | null;
  createdAt: string | null;
}

export const POST_PAGE_LIMIT = 6;
export type { PostOverview };
