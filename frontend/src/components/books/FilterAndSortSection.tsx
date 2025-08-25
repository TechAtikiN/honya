'use client'
import { BOOK_CATEGORIES, BOOK_SORT_OPTIONS } from "@/constants/books"
import DropdownFilter from "./DropdownFilter"
import Filters from "./Filters"
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { Button } from "../ui/button";
import { Trash2 } from "lucide-react";
import HintLabel from "../global/hint-label";
import { Filters as TFilters } from "@/types/book";
import { LocaleDict } from "@/lib/locales";
import { Locale } from "@/i18n.config";

interface FiltersAndSortSectionProps {
  filters: TFilters
  translations: LocaleDict
  locale: Locale
}

export default function FilterAndSortSection({
  filters,
  translations,
  locale
}: FiltersAndSortSectionProps) {
  const searchParams = useSearchParams();
  const pathname = usePathname()

  return (
    <div className="flex flex-wrap items-center justify-between gap-2">
      <div className="flex flex-wrap items-center justify-start gap-2">
        <Filters
          translations={translations}
          locale={locale}
        />
      </div>

      <div className="flex items-center justify-start gap-2">
        {filters && Object.keys(filters).length > 0 && (
          <HintLabel
            label={translations.page.home.filters.clearAllFilters}
            side="bottom"
          >
            <Button
              type="button"
              variant={'destructive'}
              onClick={() => {
                const params = new URLSearchParams(searchParams.toString());
                params.delete('category');
                params.delete('publication_year');
                params.delete('sort');
                params.delete('page');
                params.delete('filter_by');
                params.delete("sort")
                const queryString = params.toString();
                const newPath = queryString ? `${pathname}?${queryString}` : pathname;
                window.location.href = newPath;
              }}
            >
              <Trash2 className="h-4 w-4" />
            </Button>
          </HintLabel>
        )}

        <DropdownFilter
          label={translations.page.home.filters.category}
          searchParamKey="category"
          defaultValue="all"
          list={BOOK_CATEGORIES}
          locale={locale}
        />
        <DropdownFilter
          label="Sort"
          searchParamKey="sort"
          defaultValue="recently_added"
          list={BOOK_SORT_OPTIONS}
          locale={locale}
        />
      </div>
    </div>
  )
}
