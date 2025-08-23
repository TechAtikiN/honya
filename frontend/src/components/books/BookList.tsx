import BookListItem from './BookListItem'
import { Book } from '@/types/book'

interface BookListProps {
  books: Book[]
}

export default function BookList({ books }: BookListProps) {
  return (
    <div className="grid max-[420px]:grid-cols-1 grid-cols-2 sm:grid-cols-3 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-5">
      {books.map((book) => (
        <BookListItem
          key={book.id}
          book={book}
        />
      ))}
    </div>
  )
}
