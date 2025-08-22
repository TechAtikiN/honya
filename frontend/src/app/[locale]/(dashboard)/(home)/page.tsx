import { getLocale } from "@/i18n.config";
import { getDictionary } from "@/lib/locales";
import Header from "@/components/books/Header";
import BookList from "@/components/books/BookList";
import FilterAndSortSection from "@/components/books/FilterAndSortSection";

interface HomePageProps {
  params: Promise<{ locale: string }>;
}

export default async function Home({
  params
}: HomePageProps) {
  const locale = await params;
  const lang = getLocale(locale.locale);
  const { page } = await getDictionary(lang);

  return (
    <div className="flex flex-col space-y-6 h-[calc(100vh-30px)] overflow-auto invisible-scrollbar">
      <p>{page.home.title}</p>
      {/* Search input and Add button */}
      <Header />

      {/* Filters and Sort */}
      <FilterAndSortSection />

      {/* Book list */}
      <div className=" h-[calc(100vh-30px)] overflow-auto invisible-scrollbar">
        <BookList />
      </div>
    </div>
  );
}
