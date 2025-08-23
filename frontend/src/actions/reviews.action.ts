'use server'

import { ReviewFormData } from "@/components/book-details/AddReview";
import { ReviewsResponse } from "@/types/book";
import { revalidatePath } from "next/cache";

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://localhost:8080/api'

export async function getBookReviews(bookId: string): Promise<ReviewsResponse | null> {
  try {
    const res = await fetch(`${BACKEND_API_URL}/reviews/book/${bookId}`, {
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
      console.error('Error fetching reviews:', res.statusText);
      return null;
    }
  } catch (error) {
    console.error('Error fetching reviews:', error);
    throw error;
  }
}

export async function addReview(data: ReviewFormData, bookId: string) {
  try {
    const res = await fetch(`${BACKEND_API_URL}/reviews`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ ...data, book_id: bookId }),
      cache: 'no-store',
    });

    if (res.ok) {
      revalidatePath(`/books/${bookId}`);
      return { success: true, message: 'Review added successfully' };

    } else {
      if (res.status === 400) {
        return {
          success: false,
          message: 'Invalid data provided. Please try again.',
        }
      } else if (res.status === 500) {
        return {
          success: false,
          message: "Server error, Please try again later.",
        }
      } else if (res.status === 409) {
        return {
          success: false,
          message: "Conflict error, Please try again later.",
        }
      } else if (res.status === 404) {
        return {
          success: false,
          message: "Not found, Please try again later.",
        }
      } else {
        console.error('Unexpected error:', res.statusText);
        return {
          success: false,
          message: 'Unexpected error occurred',
        }
      }
    }
  } catch (error) {
    console.error('Error adding review:', error);
    throw error;
  }
}