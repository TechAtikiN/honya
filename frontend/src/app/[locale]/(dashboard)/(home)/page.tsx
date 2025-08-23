import { getLocale } from "@/i18n.config";
import { getBooks } from "@/actions/book.actions";
import Header from "@/components/books/Header";
import BookList from "@/components/books/BookList";
import FilterAndSortSection from "@/components/books/FilterAndSortSection";
import { getFilters } from "@/lib/utils";
import CustomLink from "@/components/global/custom-link";
import { ChevronLeft, ChevronRight } from "lucide-react";

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

  const filters = await searchParams;
  const formattedFilters = getFilters(filters);

  const currentPage = formattedFilters.page ? parseInt(formattedFilters.page[0] as string, 10) : 1;

  const limit = 10;
  const { data, meta } = await getBooks(formattedFilters, currentPage, limit);

  return (
    <div className="flex flex-col space-y-6 h-[calc(100vh-30px)] overflow-auto invisible-scrollbar pb-5">
      {/* Search input and Add button */}
      <Header />

      {/* Filters and Sort */}
      <FilterAndSortSection />

      {!data || data.length === 0 ? (
        <div className="flex flex-col items-center justify-center h-full space-y-3">
          <p className="text-primary text-lg">No books found.</p>

        </div>
      ) : (
        <div className="w-full flex flex-col justify-between h-full gap-y-3">
          <BookList books={data} />

          {/* Pagination */}
          <div className="flex items-center justify-center space-x-5 -ml-28">
            {currentPage > 1 ? (
              <CustomLink
                locale={lang}
                href={`?page=${currentPage - 1}`}
                className="flex items-center justify-center space-x-1 min-w-24"
              >
                <ChevronLeft className="h-5 w-5 text-primary" />
                <span className="font-medium">Previous</span>
              </CustomLink>
            ) : (
              <div className="min-w-24"></div>
            )}
            <p className="text-primary font-normal text-sm">
              Page {currentPage} of {Math.ceil(meta.total_count / limit)}
            </p>
            {(currentPage * limit) < meta.total_count ? (
              <CustomLink
                locale={lang}
                href={`?page=${currentPage + 1}`}
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
