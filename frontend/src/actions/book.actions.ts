'use server'

import { BookFormData } from "@/components/books/BookForm";
import { BookDetailsResponse, BooksResponse, Filters } from "@/types/book"
import { revalidatePath } from "next/cache";
import { getBookReviews } from "./reviews.action";

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://localhost:8080/api'

export async function getBooks(filters: Filters, pagination: { currentPage: number, limit: number }): Promise<BooksResponse | null> {
  try {
    const offset = (pagination.currentPage - 1) * pagination.limit;

    const URL = `${BACKEND_API_URL}/books?limit=${pagination.limit}&offset=${offset}&` + new URLSearchParams(filters as Record<string, string>);

    const res = await fetch(URL, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });

    if (res.ok) {
      const data = await res.json();
      return data;
    } else {
      console.error('Error fetching books', res.statusText);
      return null
    }

  } catch (error) {
    console.error('Error fetching books:', error);
    throw error;
  }
}

export async function getBookDetails(id: string): Promise<BookDetailsResponse | null> {
  try {
    const res = await fetch(`${BACKEND_API_URL}/books/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });

    const bookReviews = await getBookReviews(id);

    if (res.ok) {
      const bookDetails = await res.json();
      return { ...bookDetails, reviews: bookReviews };
    } else {
      console.error('Error fetching book:', res.statusText);
      return null;
    }
  } catch (error) {
    console.error('Error fetching book:', error);
    throw error;
  }
}

export async function addBook(data: BookFormData, imageFile?: File) {
  try {
    const formData = new FormData();

    formData.append('title', data.title);
    formData.append('isbn', data.isbn);
    formData.append('description', data.description);
    formData.append('author_name', data.authorName);
    formData.append('category', data.category);
    formData.append('publication_year', data.publicationYear.toString());
    formData.append('pages', data.pages.toString());
    formData.append('rating', data.rating.toString());

    if (imageFile) {
      formData.append('image', imageFile);
    }

    const res = await fetch(`${BACKEND_API_URL}/books`, {
      method: 'POST',
      body: formData,
    });

    if (res.ok) {
      revalidatePath('/');
      return {
        success: true,
        messageKey: 'actions.book.addSuccess',
      }
    } else {
      const errorData = await res.json();
      if (errorData.error.includes('A book with this ISBN already exists')) {
        return {
          success: false,
          messageKey: 'actions.book.isbnExists',
        }
      }
      if (res.status === 400) {
        return {
          success: false,
          messageKey: 'actions.book.invalidData',
        }
      } else if (res.status === 500) {
        return {
          success: false,
          messageKey: 'actions.book.serverError',
        }
      } else if (res.status === 409) {
        return {
          success: false,
          messageKey: 'actions.book.conflictError',
        }
      } else if (res.status === 404) {
        return {
          success: false,
          messageKey: 'actions.book.notFound',
        }
      } else {
        console.error('Unexpected error:', errorData);
        return {
          success: false,
          messageKey: 'actions.book.unexpectedError',
        }
      }
    }
  } catch (error) {
    console.error('Error creating book:', error);
    throw error;
  }
}

export async function updateBook(data: Partial<BookFormData> & { id: string }, imageFile?: File) {
  try {
    const formData = new FormData();

    if (data.title) formData.append('title', data.title);
    if (data.isbn) formData.append('isbn', data.isbn);
    if (data.description) formData.append('description', data.description);
    if (data.authorName) formData.append('author_name', data.authorName);
    if (data.category) formData.append('category', data.category);
    if (data.publicationYear) formData.append('publication_year', data.publicationYear.toString());
    if (data.pages) formData.append('pages', data.pages.toString());
    if (data.rating) formData.append('rating', data.rating.toString());

    if (imageFile) {
      formData.append('image', imageFile);
    }

    const res = await fetch(`${BACKEND_API_URL}/books/${data.id}`, {
      method: 'PATCH',
      body: formData,
    });

    if (res.ok) {
      revalidatePath(`/books/${data.id}`);
      return {
        success: true,
        messageKey: 'actions.book.updateSuccess',
      }
    } else {
      const errorData = await res.json();
      if (errorData.error.includes('A book with this ISBN already exists')) {
        return {
          success: false,
          messageKey: 'actions.book.isbnExists',
        }
      }
      if (res.status === 400) {
        console.error(errorData.error);
        return {
          success: false,
          messageKey: 'actions.book.invalidData',
        }
      } else if (res.status === 500) {
        console.error(errorData.error);
        return {
          success: false,
          messageKey: 'actions.book.serverError',
        }
      } else if (res.status === 409) {
        console.error(errorData.error);
        return {
          success: false,
          messageKey: 'actions.book.conflictError',
        }
      } else if (res.status === 404) {
        console.error(errorData.error);
        return {
          success: false,
          messageKey: 'actions.book.notFound',
        }
      } else {
        console.error('Unexpected error:', errorData);
        return {
          success: false,
          messageKey: 'actions.book.unexpectedError',
        }
      }
    }
  } catch (error) {
    console.error('Error updating book:', error);
    throw error;
  }
}

export async function deleteBook(id: string) {
  try {
    const res = await fetch(`${BACKEND_API_URL}/books/${id}`, {
      method: 'DELETE',
    });

    if (res.ok) {
      // revalidatePath('/');
      return {
        success: true,
        messageKey: 'actions.book.deleteSuccess',
      }
    } else {
      const errorData = await res.json();
      if (res.status === 400) {
        return {
          success: false,
          messageKey: 'actions.book.invalidData',
        }
      } else if (res.status === 500) {
        return {
          success: false,
          messageKey: 'actions.book.serverError',
        }
      } else if (res.status === 409) {
        return {
          success: false,
          messageKey: 'actions.book.conflictError',
        }
      } else if (res.status === 404) {
        return {
          success: false,
          messageKey: 'actions.book.notFound',
        }
      } else {
        console.error('Unexpected error:', errorData);
        return {
          success: false,
          messageKey: 'actions.book.unexpectedError',
        }
      }
    }
  } catch (error) {
    console.error('Error deleting book:', error);
    throw error;
  }
}