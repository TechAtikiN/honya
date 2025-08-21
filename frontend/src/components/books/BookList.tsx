import { MOCK_BOOKS_DATA } from '@/constants/books'
import BookListItem from './BookListItem'

export default function BookList() {
  return (
    <div className="grid max-[420px]:grid-cols-1 grid-cols-2 sm:grid-cols-3 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-8">
      {MOCK_BOOKS_DATA.map((book) => (
        <BookListItem
          key={book.id}
          book={book}
        />
      ))}
    </div>
  )
}
