'use client'
import { Button } from '../ui/button'
import { Trash2 } from 'lucide-react'
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '../ui/dialog'
import { useState, useTransition } from 'react'
import { deleteBook } from '@/actions/book.actions'
import { toast } from 'sonner'
import { useRouter } from 'next/navigation'

interface DeleteBookProps {
  bookId: string
}

export default function DeleteBook({ bookId }: DeleteBookProps) {
  const [isOpen, setIsOpen] = useState(false)
  const [isPending, startTransition] = useTransition()
  const router = useRouter()

  const handleDeleteBook = () => {
    try {
      startTransition(async () => {
        const response = await deleteBook(bookId);
        if (response.success) {
          toast.success(response.message || 'Book deleted successfully');
          setIsOpen(false);
          window.location.href = '/';
        } else {
          console.error('Error updating book:', response?.message);
          toast.error(response?.message || 'Failed to delete book');
        }
      });
    } catch (error) {
      console.error('Error deleting book:', error);
      toast.error('An unexpected error occurred while deleting the book');
    }
  }

  return (
    <div>
      <div className="flex flex-col space-y-3">
        <div className="flex items-center justify-between">
          <p className="text-2xl font-bold text-primary">Danger Zone</p>
          <Dialog open={isOpen} onOpenChange={setIsOpen}>
            <DialogTrigger asChild>
              <Button variant="destructive" className="flex items-center">
                <Trash2 className="h-4 w-4" />
                <span>Delete Book</span>
              </Button>
            </DialogTrigger>
            <DialogContent className="w-full md:min-w-[350px]">
              <DialogHeader>
                <DialogTitle>Confirm Deletion</DialogTitle>
                <DialogDescription>
                  Are you sure you want to delete this book? This action cannot be undone.
                </DialogDescription>
              </DialogHeader>

              <DialogFooter className="sm:justify-end">
                <DialogClose asChild>
                  <Button type="button" variant="secondary">
                    Cancel
                  </Button>
                </DialogClose>
                <Button
                  onClick={handleDeleteBook}
                  disabled={isPending}
                  type="button" variant="destructive"
                >
                  {isPending ? 'Deleting...' : 'Confirm Delete'}
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>
    </div>
  )
}
