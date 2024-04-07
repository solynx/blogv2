interface PostOverview {
  title: string;
  slug: string;
  description: string | null;
  category: object | null;
  author: object | null;
  createdAt: string | null;
}

interface PostDetail {
  id: string | null;
  title: string;
  content: string | null;
  category: object | null;
  author: object | null;
  createdAt: string | null;
}

export const POST_PAGE_LIMIT = 8,
  SEO_TITLE_DEFAULT =
    "Retherer Blog là trang web học tập thiên hướng về lập trình và giới thiệu sách, tích lũy kinh nghiệm.",
  SEO_DESCRIPTION_DEFAULT = `Retherer Blog được phát triển bởi một nhóm sinh viên trẻ và giàu nhiệt huyết đam mê công nghệ và yêu sáng tạo.
                                    Qua blog này các bạn có thể tìm hiểu được các kiến thức cơ bản của lập trình, các ngôn ngữ lập trình như Go, Rust, Javascript và framework tạo
                                    ra một trang web nhanh chóng. Tập trung vào khả năng sáng tạo nên tất cả mọi thứ đều không có giới hạn!`,
  SEO_KEYWORDS_DEFAULT = `Retherer, Retherer blog, blog, dev, Program language, Go, Rust, Blockchain, Web3, Solana, OOP, C, C++`;

export type { PostOverview };
