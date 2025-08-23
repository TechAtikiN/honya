'use client'
import { BOOK_CATEGORIES, BOOK_SORT_OPTIONS } from "@/constants/books"
import DropdownFilter from "./DropdownFilter"
import Filters from "./Filters"
import { useRouter } from "next/navigation";
import { Button } from "../ui/button";
import { Trash2 } from "lucide-react";
import HintLabel from "../global/hint-label";
import { Filters as TFilters } from "@/types/book";

interface FiltersAndSortSectionProps {
  filters: TFilters
}

export default function FilterAndSortSection({
  filters
}: FiltersAndSortSectionProps) {
  const router = useRouter()
  return (
    <div className="flex flex-wrap items-center justify-between gap-2">
      <div className="flex flex-wrap items-center justify-start gap-2">
        <Filters />
      </div>

      <div className="flex items-center justify-start gap-2">
        {filters && Object.keys(filters).length > 0 && (
          <HintLabel
            label="Clear all filters"
            side="bottom"
          >
            <Button
              type="button"
              variant={'destructive'}
              onClick={() => {
                router.push('/')
              }}
            >
              <Trash2 className="h-4 w-4" />
            </Button>
          </HintLabel>
        )}

        <DropdownFilter
          label="Category"
          searchParamKey="category"
          defaultValue="all"
          list={BOOK_CATEGORIES}
        />
        <DropdownFilter
          label="Sort"
          searchParamKey="sort"
          defaultValue="recently_added"
          list={BOOK_SORT_OPTIONS}
        />
      </div>
    </div>
  )
}
