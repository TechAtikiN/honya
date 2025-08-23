import { getBookDetails, getBooks } from "@/actions/book.actions";

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
export interface Filters {
  publication_year?: string
  category?: string
  query?: string
  sort?: string
  pages?: string
  rating?: string
}

export type Book = Awaited<ReturnType<typeof getBookDetails>>;
export interface BooksResponse {
  data: Book[]
  meta: {
    total_count: number
    limit: number
    offset: number
  }
}

export type BooksDetails = Awaited<ReturnType<typeof getBooks>>;