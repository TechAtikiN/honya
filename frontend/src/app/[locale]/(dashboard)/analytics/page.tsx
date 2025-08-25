import { getBookData, getReviewsData } from "@/actions/analytics.actions";
import { getBooks } from "@/actions/book.actions";
import BooksChart from "@/components/analytics/BooksChart";
import ReviewsChart from "@/components/analytics/ReviewsChart";
import AddNewBook from "@/components/books/AddNewBook";
import BooksPagination from "@/components/books/BooksPagination";
import BookTable from "@/components/books/BookTable";
import FilterAndSortSection from "@/components/books/FilterAndSortSection";
import { getLocale } from "@/i18n.config";
import { getDictionary } from "@/lib/locales";
import { getFilters, getPagination } from "@/lib/utils";
import { Book, Filters } from "@/types/book";

interface HomePageProps {
  params: Promise<{ locale: string }>;
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}

export default async function Analytics({ params, searchParams }: HomePageProps) {
  const locale = await params;
  const lang = getLocale(locale.locale);
  const translations = await getDictionary(lang)

  const _searchParams = await searchParams;
  const filterBy = (_searchParams.filter_by as string) || 'category';

  const filters = _searchParams;
  const formattedFilters: Filters = getFilters(filters);
  const pagination = getPagination(filters.page);

  const { reviewsData, booksData, data, meta } = await Promise.all([
    getReviewsData(),
    getBookData(filterBy),
    getBooks(formattedFilters, pagination)
  ]).then(([reviewsData, booksData, response]) => {
    return { reviewsData, booksData, ...response };
  }) || { reviewsData: null, booksData: null, data: [], meta: { total_count: 0 }, };

  return (
    <div className="flex flex-col space-y-6 h-[calc(100vh-30px)] overflow-auto invisible-scrollbar pb-5">
      {/* Charts */}
      <div className="grid grid-cols-1 min-[850px]:grid-cols-2 gap-5">
        <div className="p-3 rounded-md border border-primary/20 flex">
          {!booksData ? (
            <div className="">
              {translations.page.analytics.noBooks}
            </div>
          ) : (
            <BooksChart locale={lang} booksData={booksData} filterBy={filterBy} translations={translations} />
          )}
        </div>
        <div className="p-3 rounded-md border border-primary/20 flex">
          {!reviewsData ? (
            <div className="">
              {translations.page.analytics.noReviews}
            </div>
          ) : (
            <ReviewsChart locale={lang} reviewsData={reviewsData} translations={translations} />
          )}
        </div>
      </div>

      {/* Books Table */}
      <div className="w-full flex flex-col justify-between h-full gap-y-5 border border-primary/20 rounded-md p-3">
        <div className="flex flex-col gap-y-3">
          <div className="flex items-center justify-between">
            <p className="text-lg font-bold text-primary">
              {translations.page.analytics.booksTable}
            </p>
            <AddNewBook
              translations={translations}
              locale={lang}
            />
          </div>

          {/* Filters and Sort */}
          <FilterAndSortSection
            filters={formattedFilters}
            translations={translations}
            locale={lang}
          />

        </div>
        <div className="flex flex-col gap-y-5">
          {/* Books Table */}
          <BookTable books={data as Book[]} locale={lang} translations={translations} />

          {/* Pagination */}
          {meta && meta.total_count > 0 && (
            <BooksPagination
              totalCount={meta?.total_count || 0}
              locale={lang}
              pagination={pagination}
              translations={translations}
            />
          )}
        </div>
      </div>
    </div>
  );
}
