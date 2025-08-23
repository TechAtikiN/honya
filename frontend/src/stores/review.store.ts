import { create } from 'zustand';
import { persist } from 'zustand/middleware';

interface ReviewState {
  submittedReviews: Record<string, boolean>;
  markReviewSubmitted: (bookId: string) => void;
  hasSubmittedReview: (bookId: string) => boolean;
}

export const useReviewStore = create<ReviewState>()(
  persist(
    (set, get) => ({
      submittedReviews: {},
      markReviewSubmitted: (bookId: string) =>
        set((state) => ({
          submittedReviews: {
            ...state.submittedReviews,
            [bookId]: true,
          },
        })),
      hasSubmittedReview: (bookId: string) => {
        return !!get().submittedReviews[bookId];
      },
    }),
    {
      name: 'review-storage',
    }
  )
);
