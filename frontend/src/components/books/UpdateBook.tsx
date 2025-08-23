'use client'
import { useState } from 'react';
import BookForm from './BookForm';
import { Book } from '@/types/book';

interface UpdateBookProps {
  bookDetails: Book;
}

export default function UpdateBook({ bookDetails }: UpdateBookProps) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <BookForm
      bookDetails={bookDetails}
      setIsOpen={setIsOpen} isOpen={isOpen} isEdit={true}
    />
  )
}