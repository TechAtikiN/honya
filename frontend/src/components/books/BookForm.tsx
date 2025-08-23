import { Button } from '../ui/button'
import {
  DialogFooter,
  DialogClose,
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription
} from '../ui/dialog'
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { bookFormSchema } from '@/lib/validation';
import { BOOK_CATEGORIES } from '@/constants/books';
import { Slider } from '../ui/slider';
import { useState, useTransition, useEffect } from 'react';
import { BookPlus, PencilLine } from 'lucide-react';
import UploadBookImage from './UploadBookImage';
import { addBook, updateBook } from '@/actions/book.actions';
import { toast } from 'sonner';
import { Book } from '@/types/book';

export type BookFormData = z.infer<typeof bookFormSchema>;

interface BookFormProps {
  bookDetails?: Book;
  isOpen?: boolean;
  setIsOpen: (open: boolean) => void;
  isEdit?: boolean;
}

export default function BookForm({ bookDetails, isOpen = false, setIsOpen, isEdit = false }: BookFormProps) {
  const [bookImagePreview, setbookImagePreview] = useState<string | ArrayBuffer | null>(null);
  const [bookImageInfo, setBookImageInfo] = useState<{ name: string; size: number } | null>(null);
  const [bookImageError, setbookImageError] = useState<string | null>(null);
  const [imageFile, setImageFile] = useState<File | null>(null);
  const [isPending, startTransition] = useTransition();

  const getDefaultValues = (book?: Book): Partial<BookFormData> => ({
    rating: book?.rating || 4,
    category: (book?.category || BOOK_CATEGORIES[0].value) as BookFormData['category'],
    publicationYear: book?.publication_year || new Date().getFullYear(),
    pages: book?.pages || 100,
    title: book?.title || '',
    isbn: book?.isbn || '',
    description: book?.description || '',
    authorName: book?.author_name || '',
  });

  const {
    register,
    handleSubmit,
    control,
    reset,
    formState: { errors, isSubmitting },
  } = useForm<BookFormData>({
    resolver: zodResolver(bookFormSchema),
    mode: 'onChange',
    defaultValues: getDefaultValues(bookDetails),
  });

  // Reset form when bookDetails changes (after successful update)
  useEffect(() => {
    if (bookDetails && isEdit) {
      reset(getDefaultValues(bookDetails));
      setImageFile(null);
      setbookImagePreview(null);
      setBookImageInfo(null);
      setbookImageError(null);
    }
  }, [bookDetails, isEdit, reset]);

  const onSubmit = async (data: BookFormData) => {
    try {
      if (isEdit) {
        startTransition(async () => {
          const response = await updateBook({ ...data, id: bookDetails?.id || '' }, imageFile || undefined);
          if (response?.success) {
            toast.success(response.message || 'Book updated successfully');
            // Don't reset here - let useEffect handle it when new data comes in
            setIsOpen(false);
          } else {
            console.error('Error updating book:', response?.message);
            toast.error(response?.message || 'Failed to update book');
          }
        });
      } else {
        startTransition(async () => {
          const response = await addBook(data, imageFile || undefined);
          if (response?.success) {
            toast.success(response.message || 'Book created successfully');
            reset();
            setImageFile(null);
            setbookImagePreview(null);
            setBookImageInfo(null);
            setbookImageError(null);
            setIsOpen(false);
          } else {
            console.error('Error creating book:', response?.message);
            toast.error(response?.message || 'Failed to create book');
          }
        });
      }
    } catch (error) {
      console.error('Error submitting form:', error);
      toast.error('An error occurred while submitting the form. Please try again.');
    }
  };

  const handleCancel = () => {
    // Reset to current book details when canceling
    if (isEdit && bookDetails) {
      reset(getDefaultValues(bookDetails));
    } else {
      reset();
    }
    setImageFile(null);
    setbookImagePreview(null);
    setBookImageInfo(null);
    setbookImageError(null);
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button className="flex items-center justify-center md:min-w-40">
          {isEdit ? <PencilLine className="h-4 w-4" /> : <BookPlus className="h-4 w-4" />}
          <span className='hidden md:block'>
            {isEdit ? 'Edit Book' : 'Add Book'}
          </span>
        </Button>
      </DialogTrigger>
      <DialogContent className="w-full md:min-w-[750px]">
        <DialogHeader>
          <DialogTitle className="flex items-center space-x-2 justify-start">
            <BookPlus className="h-5 w-5" />
            <p>
              {isEdit ? 'Edit Book' : 'Add New Book'}
            </p>
          </DialogTitle>
          <DialogDescription className="sr-only">
            {isEdit ? 'Edit book details' : 'Add a new book to your collection'}
          </DialogDescription>
        </DialogHeader>

        <hr />

        <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col space-y-4 full h-[600px] md:h-auto overflow-auto invisible-scrollbar'>
          <div className='grid grid-cols-1 md:grid-cols-2 gap-4'>
            <div className='flex flex-col space-y-4 border-r border-secondary pr-4'>
              <UploadBookImage
                bookImagePreview={bookImagePreview}
                setbookImagePreview={setbookImagePreview}
                bookImageInfo={bookImageInfo}
                setBookImageInfo={setBookImageInfo}
                bookImageError={bookImageError}
                setbookImageError={setbookImageError}
                setImageFile={setImageFile}
                imageURL={bookDetails?.image || null}
              />

              <div className='flex flex-col space-y-1'>
                <label
                  htmlFor='title'
                  className='form-label'>Title</label>
                <input
                  type="text"
                  required
                  {...register("title")}
                  className='form-input'
                  placeholder="Harry Potter and the Philosopher's Stone"
                />
                {errors.title && <p className='text-destructive text-xs'>{errors.title.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label
                  htmlFor='isbn'
                  className='form-label'>ISBN</label>
                <input
                  type="text"
                  required
                  {...register("isbn")}
                  className='form-input disabled:bg-secondary/50 disabled:cursor-not-allowed'
                  placeholder='978-3-16-148410-0'
                  disabled={isEdit}
                />
                {errors.isbn && <p className='text-destructive text-xs'>{errors.isbn.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label
                  htmlFor='description'
                  className='form-label'>Description</label>
                <textarea
                  required
                  {...register("description")}
                  className='form-input'
                  rows={3}
                  style={{ resize: 'none' }}
                  placeholder='A brief description of the book...'
                ></textarea>
                {errors.description && <p className='text-destructive text-xs'>{errors.description.message}</p>}
              </div>
            </div>

            <div className='flex flex-col space-y-4'>
              <div className='flex flex-col space-y-1'>
                <label
                  htmlFor='authorName'
                  className='form-label'>Author Name</label>
                <input
                  required
                  type="text"
                  {...register("authorName")}
                  className='form-input'
                  placeholder='JK Rowling'
                />
                {errors.authorName && <p className='text-destructive text-xs'>{errors.authorName.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label
                  htmlFor='category'
                  className='form-label'>Category</label>
                <select
                  required
                  {...register("category")}
                  className='form-input'
                >
                  {BOOK_CATEGORIES.map((category) => (
                    <option key={category.value} value={category.value}>
                      {category.label}
                    </option>
                  ))}
                </select>
                {errors.category && <p className='text-destructive text-xs'>{errors.category.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label
                  htmlFor='publicationYear'
                  className='form-label'>Publication Year</label>
                <input
                  required
                  type="number"
                  {...register("publicationYear", {
                    setValueAs: (value) => parseInt(value, 10),
                    valueAsNumber: true,
                  })}
                  className='form-input'
                  placeholder='2020'
                  min="1950"
                  max={new Date().getFullYear()}
                />
                {errors.publicationYear && <p className='text-destructive text-xs'>{errors.publicationYear.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label
                  htmlFor='pages'
                  className='form-label'>Number of pages</label>
                <input
                  required
                  type="number"
                  {...register("pages", {
                    setValueAs: (value) => parseInt(value, 10),
                    valueAsNumber: true,
                  })}
                  className='form-input'
                  placeholder='300'
                  min="1"
                />
                {errors.pages && <p className='text-destructive text-xs'>{errors.pages.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label htmlFor="rating" className='form-label'>
                  Rating
                </label>
                <Controller
                  name="rating"
                  control={control}
                  render={({ field }) => (
                    <Slider
                      onValueChange={(value) => {
                        field.onChange(value[0]);
                      }}
                      max={5}
                      min={0}
                      step={0.5}
                      defaultValue={[field.value]}
                      value={[field.value]}
                      className="w-full"
                    />
                  )}
                />
                {errors.rating && <p className='text-destructive text-xs'>{errors.rating.message}</p>}
              </div>
            </div>
          </div>

          <DialogFooter>
            <DialogClose asChild>
              <Button
                type="button"
                variant="outline"
                onClick={handleCancel}
              >
                Cancel
              </Button>
            </DialogClose>
            <Button
              type="submit"
              disabled={isSubmitting || isPending}
            >
              {isSubmitting || isPending ? (isEdit ? 'Updating...' : 'Adding...') : (isEdit ? 'Update Book' : 'Add Book')}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}