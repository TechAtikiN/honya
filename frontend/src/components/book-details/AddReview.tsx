'use client'

import { useState, useTransition } from 'react';
import { Button } from '@/components/ui/button';
import { UserStar } from 'lucide-react';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from '../ui/dialog';
import { z } from "zod";
import { reviewFormSchema } from '@/lib/validation';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { toast } from 'sonner';
import { addReview } from '@/actions/reviews.action';
import { useReviewStore } from '@/stores/review.store';

interface AddReviewProps {
  bookId: string;
}

export type ReviewFormData = z.infer<typeof reviewFormSchema>;

export default function AddReview({ bookId }: AddReviewProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [isPending, startTransition] = useTransition();

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors, isSubmitting },
  } = useForm<ReviewFormData>({
    resolver: zodResolver(reviewFormSchema),
    mode: 'onChange',
  });

  const { hasSubmittedReview, markReviewSubmitted } = useReviewStore();

  const onSubmit = async (data: ReviewFormData) => {
    try {
      startTransition(async () => {
        const response = await addReview(data, bookId);

        if (response?.success) {
          toast.success(response.message || 'Review added successfully');
          markReviewSubmitted(bookId);
          reset();
          setIsOpen(false);
        } else {
          console.error('Failed to create review:', response?.message);
          toast.error(response?.message || 'Failed to add review. Please try again.');
        }
      });
    } catch (error) {
      console.error('Error submitting form:', error);
      toast.error('An error occurred while submitting the form. Please try again.');
    }
  };

  const alreadySubmitted = hasSubmittedReview(bookId);

  return (
    <div>
      <Dialog open={isOpen} onOpenChange={setIsOpen}>
        <DialogTrigger asChild>
          <Button
            variant="secondary"
            className="flex items-center"
            onClick={() => setIsOpen(true)}
            disabled={alreadySubmitted}
          >
            <UserStar className="h-4 w-4" />
            <span>{alreadySubmitted ? 'Review Submitted' : 'Add Review'}</span>
          </Button>
        </DialogTrigger>
        <DialogContent className="w-full md:min-w-[350px]">
          <DialogHeader>
            <DialogTitle className="flex items-center space-x-2 justify-start">
              <UserStar className="h-5 w-5" />
              <p>Add Review</p>
            </DialogTitle>
            <DialogDescription className="sr-only">
              Add your review for this book.
            </DialogDescription>
          </DialogHeader>
          <hr />
          <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col space-y-4">
            <div className="grid grid-cols-2 gap-4">
              <div className="flex flex-col space-y-1 col-span-2 md:col-span-1">
                <label htmlFor="name" className="form-label">Name</label>
                <input
                  type="text"
                  required
                  {...register("name")}
                  className="form-input"
                  placeholder="Daisy Jones"
                />
                {errors.name && (
                  <p className="text-destructive text-xs">{errors.name.message}</p>
                )}
              </div>
              <div className="flex flex-col space-y-1 col-span-2 md:col-span-1">
                <label htmlFor="email" className="form-label">Email</label>
                <input
                  type="email"
                  required
                  {...register("email")}
                  className="form-input"
                  placeholder="jones@gmail.com"
                />
                {errors.email && (
                  <p className="text-destructive text-xs">{errors.email.message}</p>
                )}
              </div>
              <div className="flex flex-col space-y-1 col-span-2">
                <label htmlFor="content" className="form-label">Review Content</label>
                <textarea
                  {...register("content")}
                  className="form-input"
                  placeholder="Your review..."
                  rows={4}
                  style={{ resize: 'none' }}
                ></textarea>
                {errors.content && (
                  <p className="text-destructive text-xs">{errors.content.message}</p>
                )}
              </div>
            </div>
            <DialogFooter className="flex items-center justify-end">
              <DialogClose asChild>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => {
                    reset();
                    setIsOpen(false);
                  }}
                >
                  Cancel
                </Button>
              </DialogClose>
              <Button type="submit" disabled={isSubmitting || isPending}>
                {isSubmitting || isPending ? 'Submitting...' : 'Submit Review'}
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>
    </div>
  );
}
