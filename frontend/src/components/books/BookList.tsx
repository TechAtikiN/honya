'use client'

import { useSidebarStore } from '@/stores/sidebar.store'
import BookListItem from './BookListItem'
import { Book } from '@/types/book'

interface BookListProps {
  books: Book[] | null
}

export default function BookList({ books }: BookListProps) {
  if (!books) return null
  const { collapse } = useSidebarStore();

  return (
    <div className={`grid max-[420px]:grid-cols-1 grid-cols-2 sm:grid-cols-2 
   ${collapse ? 'md:grid-cols-3' : 'md:grid-cols-2'}
    ${collapse ? 'lg:grid-cols-4' : 'lg:grid-cols-4'}
     xl:grid-cols-5 gap-5`}>
      {books && books.map((book) => (
        <BookListItem
          key={book?.id}
          book={book}
        />
      ))}
    </div>
  )
}
