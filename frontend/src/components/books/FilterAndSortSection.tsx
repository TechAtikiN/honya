import { BOOK_CATEGORIES, BOOK_SORT_OPTIONS } from "@/constants/books"
import DropdownFilter from "./DropdownFilter"
import Filters from "./Filters"

export default function FilterAndSortSection() {
  return (
    <div className="flex flex-wrap items-center justify-between gap-2">
      <div className="flex flex-wrap items-center justify-start gap-2">
        <Filters />
      </div>

      <div className="flex items-center justify-start gap-2">
        <DropdownFilter
          label="Category"
          searchParamKey="category"
          defaultValue="all"
          list={BOOK_CATEGORIES}
        />
        <DropdownFilter
          label="Sort"
          searchParamKey="sort_by"
          defaultValue="recently_added"
          list={BOOK_SORT_OPTIONS}
        />
      </div>
    </div>
  )
}
