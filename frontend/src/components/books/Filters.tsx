import { useState } from 'react'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger
} from '../ui/dropdown-menu'
import { Button } from '../ui/button'
import RangeFilter from './RangeFilter'
import { ListFilterPlus } from 'lucide-react'

interface SelectedFilter {
  key: string
  label: string
  value: number
}

export default function Filters() {
  const [selectedFilters, setSelectedFilters] = useState<SelectedFilter[]>([])

  const handleFilterChange = (key: string, label: string, value: number) => {
    setSelectedFilters(prev => {
      const existing = prev.find(f => f.key === key)
      if (existing) {
        return prev.map(f => f.key === key ? { key, label, value } : f)
      } else {
        return [...prev, { key, label, value }]
      }
    })
  }

  const getButtonLabel = () => {
    if (selectedFilters.length === 0) return 'Filters'
    if (selectedFilters.length === 1) {
      const { label, value } = selectedFilters[0]
      return `${label}: ${value}`
    }
    return `${selectedFilters.length} selected`
  }

  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="secondary" className="flex items-center justify-between min-w-[180px] text-primary">
            <span>{getButtonLabel()}</span>
            <ListFilterPlus className="h-4 w-4" />
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
              onFilterChange={handleFilterChange}
            />
            <RangeFilter
              searchParamKey="rating"
              label="Minimum Rating"
              defaultValue={0}
              max={5}
              step={0.5}
              fromLabel={0}
              onFilterChange={handleFilterChange}
            />
            <RangeFilter
              searchParamKey="pages"
              label="Number of Pages"
              defaultValue={10000}
              max={10000}
              step={1000}
              fromLabel={1}
              onFilterChange={handleFilterChange}
            />
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
