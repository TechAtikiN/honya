import { BOOK_CATEGORIES, BOOK_SORT_OPTIONS } from "@/constants/books"
import DropdownFilter from "./DropdownFilter"
import RangeFilter from "./RangeFilter"

export default function FilterAndSortSection() {
  return (
    <div className="flex items-center justify-between">
      <div className="flex items-center justify-start space-x-4">
        <RangeFilter
          searchParamKey="publication_year"
          label="Publication Year"
          defaultValue={new Date().getFullYear()}
          max={new Date().getFullYear()}
          step={10}
          fromLabel={1950}
        />
        <RangeFilter
          searchParamKey="rating"
          label="Minimum Rating"
          defaultValue={0}
          max={5}
          step={0.5}
          fromLabel={0}
        />
        <RangeFilter
          searchParamKey="pages"
          label="Number of Pages"
          defaultValue={10000}
          max={10000}
          step={1000}
          fromLabel={1}
        />
      </div>

      <div className="flex items-center justify-start space-x-4">
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
