import { Book } from '@/types/book';
import BookTableItem from './BookTableItem';
import { Locale } from '@/i18n.config';
import { LocaleDict } from '@/lib/locales';

interface BookListProps {
  books: Book[] | null;
  locale: Locale;
  translations: LocaleDict;
}

export default function BookTable({
  books,
  locale,
  translations,
}: BookListProps) {
  if (!books) return null;
  return (
    <div className='min-h-[300px]'>
      <div className='relative overflow-x-auto shadow-sm sm:rounded-lg'>
        {books.length > 0 ? (
          <table className='w-full text-sm'>
            <thead className='text-sm bg-secondary rounded-md font-medium text-primary uppercase'>
              <tr>
                <th scope='col' className='table-header'>
                  {translations.page.analytics.title}
                </th>
                <th scope='col' className='table-header'>
                  {translations.page.analytics.author}
                </th>
                <th scope='col' className='table-header'>
                  {translations.page.analytics.category}
                </th>
                <th scope='col' className='table-header'>
                  {translations.page.analytics.rating}
                </th>
                <th scope='col' className='table-header'>
                  {translations.page.analytics.publicationYear}
                </th>
                <th scope='col' className='table-header'>
                  {translations.page.analytics.view}
                </th>
              </tr>
            </thead>
            <tbody className=''>
              {books.length > 0 &&
                books.map((book) => (
                  <BookTableItem key={book?.id} book={book} locale={locale} />
                ))}
            </tbody>
          </table>
        ) : (
          <div className='p-4 text-center text-gray-500'>
            {translations.page.analytics.noDataFound}
          </div>
        )}
      </div>
    </div>
  );
}
