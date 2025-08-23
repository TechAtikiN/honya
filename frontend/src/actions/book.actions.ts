'use server'

import { BooksResponse, Filters } from "@/types/book"

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://localhost:8080/api'

export async function getBooks(filters: Filters, page: number = 1, limit: number = 10): Promise<BooksResponse> {
  const offset = (page - 1) * limit;

  const URL = `${BACKEND_API_URL}/books?limit=${limit}&offset=${offset}&` + new URLSearchParams(filters as Record<string, string>);

  console.log('Fetching books from URL:', URL);

  const res = await fetch(URL, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    cache: 'no-store',
  });

  const data = await res.json();
  return data;
}
