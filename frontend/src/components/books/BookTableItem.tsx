import React from 'react'

import { Book } from '@/types/book'
import { Locale } from '@/i18n.config'
import CustomLink from '../global/custom-link'
import { ExternalLink } from 'lucide-react'

interface BookTableItemProps {
  book: Book
  locale: Locale
}

export default function BookTableItem({ book, locale }: BookTableItemProps) {
  if (!book) return null
  return (
    <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
      <th scope="row" className="px-6 py-4 text-left  font-medium text-gray-900 dark:text-white whitespace-nowrap">
        {book.title}
      </th>
      <td className="px-6 py-4 text-left">
        {book.author_name}
      </td>
      <td className="px-6 py-4 text-left">
        {book.category}
      </td>
      <td className="px-6 py-4 text-left">
        {book.rating}
      </td>
      <td className="px-6 py-4 text-left">
        {book.publication_year}
      </td>
      <td className="px-6 py-4 text-left pl-7">
        <CustomLink
          target="_blank"
          locale={locale}
          href={`/books/${book.id}`}
          className="text-primary font-bold"
        >
          <ExternalLink className='h-4 w-4' />
        </CustomLink>
      </td>
    </tr>
  )
}
