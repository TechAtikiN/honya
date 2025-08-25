import { Locale } from '@/i18n.config'
import { ChevronLeft, ChevronRight } from 'lucide-react'
import CustomLink from '../global/custom-link'

interface BooksPaginationProps {
  totalCount: number
  locale: Locale
  pagination: {
    currentPage: number
    limit: number
  }
}

export default function BooksPagination({
  totalCount,
  locale: lang,
  pagination
}: BooksPaginationProps) {
  return (
    <div className="flex items-center justify-center space-x-5 -ml-28">
      {pagination.currentPage > 1 ? (
        <CustomLink
          locale={lang}
          href={`?page=${pagination.currentPage - 1}`}
          className="flex items-center justify-center space-x-1 min-w-24"
        >
          <ChevronLeft className="h-5 w-5 text-primary" />
          <span className="font-medium">Previous</span>
        </CustomLink>
      ) : (
        <div className="min-w-24"></div>
      )}

      <p className="text-primary font-normal text-sm">
        ({totalCount} Results)  Page {pagination.currentPage} of {Math.ceil((totalCount || 0) / pagination.limit)}
      </p>

      {(pagination.currentPage * pagination.limit) < totalCount ? (
        <CustomLink
          locale={lang}
          href={`?page=${pagination.currentPage + 1}`}
          className="flex items-center justify-center space-x-1 min-w-24"
        >
          <span className="font-medium">Next</span>
          <ChevronRight className="h-5 w-5 text-primary" />
        </CustomLink>
      ) : (
        <div className="min-w-24"></div>
      )}
    </div>
  )
}
