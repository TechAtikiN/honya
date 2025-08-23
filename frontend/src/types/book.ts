import { getBooks } from "@/actions/book.actions";

export enum bookCategory {
  FICTION = "fiction",
  NON_FICTION = "non_fiction",
  SCIENCE = "science",
  HISTORY = "history",
  FANTASY = "fantasy",
  MYSTERY = "mystery",
  THRILLER = "thriller",
  COOKING = "cooking",
  TRAVEL = "travel",
  CLASSICS = "classics",
}

export type Book = {
  id: string;
  title: string;
  description: string;
  author_name: string;
  category: bookCategory;
  image: string;
  publication_year: number;
  rating: number;
  pages: number;
  isbn: string;
  created_at: Date;
  updated_at: Date;
}

export interface Filters {
  publication_year?: string
  category?: string
  query?: string
  sort?: string
  pages?: string
  rating?: string
}


export interface BooksResponse {
  data: Book[]
  meta: {
    total_count: number
    limit: number
    offset: number
  }
}

export type BooksDetails = Awaited<ReturnType<typeof getBooks>>;