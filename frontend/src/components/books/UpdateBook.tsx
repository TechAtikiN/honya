'use client'
import { useState } from 'react';
import BookForm from './BookForm';

export default function UpdateBook() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div>
      <BookForm setIsOpen={setIsOpen} isOpen={isOpen} isEdit={true} />
    </div>
  )
}