import { getLocale } from "@/i18n.config";
import { getDictionary } from "@/lib/locales";
import Header from "@/components/books/Header";

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
    <div>
      {/* Search input and Add button */}
      <Header />

      {/* Filters and Sort */}

      {/* Book list */}
    </div>
  );
}
