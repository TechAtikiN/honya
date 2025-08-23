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
import { createBookFormSchema } from '@/lib/validation';
import { BOOK_CATEGORIES } from '@/constants/books';
import { Slider } from '../ui/slider';
import { useState } from 'react';
import { BookPlus, PencilLine } from 'lucide-react';
import UploadBookImage from './UploadBookImage';

type BookFormData = z.infer<typeof createBookFormSchema>;

interface BookFormProps {
  isOpen?: boolean;
  setIsOpen: (open: boolean) => void;
  isEdit?: boolean;
}

export default function BookForm({ isOpen = false, setIsOpen, isEdit = false }: BookFormProps) {
  const [rating, setRating] = useState([4]);

  const {
    register,
    handleSubmit,
    control,
    reset,
    formState: { errors, isSubmitting },
  } = useForm<BookFormData>({
    resolver: zodResolver(createBookFormSchema),
    mode: "onChange",
    defaultValues: {
      rating: 4,
      category: BOOK_CATEGORIES[0]?.value as BookFormData['category'],
      publicationYear: new Date().getFullYear(),
      pages: 1,
    }
  });

  const onSubmit = async (data: BookFormData) => {
    try {
      console.log(data);
      // Handle form submission logic here

      // Reset form and close dialog on success
      reset();
      setRating([4]);
      setIsOpen(false);
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  };

  const handleCancel = () => {
    reset();
    setRating([4]);
    setIsOpen(false);
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button className="flex items-center justify-center md:min-w-40">
          {isEdit ? <PencilLine className="h-4 w-4" /> : <BookPlus className="h-4 w-4" />}
          <span className='hidden md:block'>
            {isEdit ? 'Edit Book' : 'Add New Book'}
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
              <UploadBookImage />

              <div className='flex flex-col space-y-1'>
                <label className='form-label'>Title</label>
                <input
                  type="text"
                  {...register("title")}
                  className='form-input'
                  placeholder="Harry Potter and the Philosopher's Stone"
                />
                {errors.title && <p className='text-destructive text-xs'>{errors.title.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label className='form-label'>ISBN</label>
                <input
                  type="text"
                  {...register("isbn")}
                  className='form-input disabled:bg-secondary/50 disabled:cursor-not-allowed'
                  placeholder='978-3-16-148410-0'
                  disabled={isEdit}
                />
                {errors.isbn && <p className='text-destructive text-xs'>{errors.isbn.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label className='form-label'>Description</label>
                <textarea
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
                <label className='form-label'>Author Name</label>
                <input
                  type="text"
                  {...register("authorName")}
                  className='form-input'
                  placeholder='JK Rowling'
                />
                {errors.authorName && <p className='text-destructive text-xs'>{errors.authorName.message}</p>}
              </div>

              <div className='flex flex-col space-y-1'>
                <label className='form-label'>Category</label>
                <select
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
                <label className='form-label'>Publication Year</label>
                <input
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
                <label className='form-label'>Number of pages</label>
                <input
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
                  Rating: {rating[0]}
                </label>
                <Controller
                  name="rating"
                  control={control}
                  render={({ field }) => (
                    <Slider
                      value={rating}
                      onValueChange={(value) => {
                        setRating(value);
                        field.onChange(value[0]);
                      }}
                      max={5}
                      min={0}
                      step={0.5}
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
              disabled={isSubmitting}
            >
              {isSubmitting ? 'Adding...' : 'Add Book'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>

  )
}
