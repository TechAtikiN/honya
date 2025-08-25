import BookListItem from './BookListItem'
import { Book } from '@/types/book'
import BookTableItem from './BookTableItem'
import { Locale } from '@/i18n.config'

interface BookListProps {
  books: Book[] | null
  locale: Locale
}

export default function BookTable({ books, locale }: BookListProps) {
  if (!books) return null
  return (
    <div className="">
      <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
        {books.length > 0 ? (
          <table className="w-full text-sm">
            <thead className="text-sm bg-secondary rounded-md font-medium text-primary uppercase">
              <tr>
                <th scope="col" className="px-6 py-3 text-left">
                  Title
                </th>
                <th scope="col" className="px-6 py-3 text-left">
                  Author
                </th>
                <th scope="col" className="px-6 py-3 text-left">
                  Category
                </th>
                <th scope="col" className="px-6 py-3 text-left">
                  Rating
                </th>
                <th scope="col" className="px-6 py-3 text-left">
                  Publication year
                </th>
                <th scope="col" className="px-6 py-3 text-left">
                  Action
                </th>

              </tr>
            </thead>
            <tbody className=''>
              {books.length > 0 && books.map((book) => (
                <BookTableItem
                  key={book?.id}
                  book={book}
                  locale={locale}
                />
              ))}
            </tbody>
          </table>
        ) : (
          <div className='p-4 text-center text-gray-500'>No books found.</div>
        )}
      </div>

    </div>
  )
}
