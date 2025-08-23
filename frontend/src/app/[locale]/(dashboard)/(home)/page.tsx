import { getLocale } from "@/i18n.config";
import { getBooks } from "@/actions/book.actions";
import Header from "@/components/books/Header";
import BookList from "@/components/books/BookList";
import FilterAndSortSection from "@/components/books/FilterAndSortSection";
import { getFilters, getPagination } from "@/lib/utils";
import CustomLink from "@/components/global/custom-link";
import { ChevronLeft, ChevronRight } from "lucide-react";
import { Book, Filters } from "@/types/book";
import { getDictionary } from "@/lib/locales";

interface HomePageProps {
  params: Promise<{ locale: string }>;
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}

export default async function Home({
  params,
  searchParams
}: HomePageProps) {
  const locale = await params;
  const lang = getLocale(locale.locale);
  const translations = await getDictionary(lang)

  const filters = await searchParams;
  const formattedFilters: Filters = getFilters(filters);
  const pagination = getPagination(filters.page);

  const response = await getBooks(formattedFilters, pagination);

  const { data, meta } = response || { data: [], meta: { total_count: 0 } }

  return (
    <div className="flex flex-col space-y-6 h-[calc(100vh-30px)] overflow-auto invisible-scrollbar pb-5">
      {/* Search input and Add button */}
      <Header
        translations={translations}
        locale={lang}
      />

      {/* Filters and Sort */}
      <FilterAndSortSection
        filters={formattedFilters}
        translations={translations}
        locale={lang}
      />

      {!data || data.length === 0 ? (
        <div className="flex flex-col items-center justify-center h-full space-y-3">
          <p className="text-primary text-lg">No books found.</p>
        </div>
      ) : (
        <div className="w-full flex flex-col justify-between h-full gap-y-3">
          <BookList books={data as Book[]} />

          {/* Pagination */}
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
              Page {pagination.currentPage} of {Math.ceil(meta.total_count / pagination.limit)}
            </p>
            {(pagination.currentPage * pagination.limit) < meta.total_count ? (
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
        </div>
      )}
    </div>
  );
}
