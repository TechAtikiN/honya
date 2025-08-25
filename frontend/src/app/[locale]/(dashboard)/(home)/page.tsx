import { getLocale } from "@/i18n.config";
import { getBooks } from "@/actions/book.actions";
import Header from "@/components/books/Header";
import BookList from "@/components/books/BookList";
import FilterAndSortSection from "@/components/books/FilterAndSortSection";
import { getFilters, getPagination } from "@/lib/utils";
import { Book, Filters } from "@/types/book";
import { getDictionary } from "@/lib/locales";
import BooksPagination from "@/components/books/BooksPagination";

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
    <div className="flex flex-col justify-between gap-y-5 md:gap-y-10 h-[calc(100vh-30px)] overflow-auto invisible-scrollbar pb-5">
      {/* Search input and Add button */}
      <div className="flex flex-col space-y-4">
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
      </div>

      {!data || data.length === 0 ? (
        <div className="flex flex-col items-center justify-center h-full space-y-3">
          <p className="text-primary text-lg">No books found.</p>
        </div>
      ) : (
        <div className="w-full flex flex-col justify-between h-full gap-y-6">
          <BookList books={data as Book[]} />

          {meta && meta.total_count > 0 && (
            <BooksPagination
              totalCount={meta?.total_count || 0}
              locale={lang}
              pagination={pagination}
            />
          )}
        </div>
      )}
    </div>
  );
}
