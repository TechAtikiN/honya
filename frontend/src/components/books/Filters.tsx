import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from '../ui/dropdown-menu'
import { Button } from '../ui/button'
import RangeFilter from './RangeFilter'
import { ChevronDown } from 'lucide-react'

export default function Filters() {
  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger asChild
        >
          <Button variant="secondary" className="flex items-center justify-between min-w-[160px]">
            <span>
              Filters
            </span>
            <ChevronDown className="h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent side='bottom' align='start' sideOffset={10}>
          <DropdownMenuItem className="flex flex-col gap-4">
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
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
