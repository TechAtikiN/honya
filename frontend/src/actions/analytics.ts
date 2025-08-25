'use server'

const BACKEND_API_URL = process.env.BACKEND_API_URL || 'http://localhost:8080/api'

export async function getReviewsData(): Promise<{ name: string; count: number }[] | null> {
  try {
    const res = await fetch(`${BACKEND_API_URL}/dashboard/reviews-data`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });

    if (res.ok) {
      const { data } = await res.json();
      return data;
    } else {
      console.error('Error fetching reviews data', res.statusText);
      return null;
    }
  } catch (error) {
    console.error('Error fetching reviews data:', error);
    throw error;
  }
}

export async function getBookData(filterBy: string): Promise<{ name: string; count: number }[] | null> {
  try {
    const res = await fetch(`${BACKEND_API_URL}/dashboard/books-data?filter_by=${filterBy}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });

    if (res.ok) {
      const { data } = await res.json();
      return data;
    } else {
      console.error('Error fetching books data', res.statusText);
      return null;
    }
  } catch (error) {
    console.error('Error fetching books data:', error);
    throw error;
  }
}