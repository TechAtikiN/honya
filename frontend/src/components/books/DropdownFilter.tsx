'use client'

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Button } from '../ui/button'
import { ChevronDown } from 'lucide-react'
import { useSearchParams, useRouter, usePathname } from 'next/navigation'
import { useMemo } from 'react'

interface DropdownFilterProps {
  label: string
  searchParamKey: string
  defaultValue: string
  list: { value: string; label: string; icon: React.ReactNode }[]
}

export default function DropdownFilter({
  label,
  searchParamKey,
  defaultValue,
  list,
}: DropdownFilterProps) {
  const searchParams = useSearchParams()
  const router = useRouter()
  const pathname = usePathname()

  const currentValue = searchParams.get(searchParamKey) || defaultValue

  const selectedItem = useMemo(() => {
    return list.find((item) => item.value === currentValue)
  }, [list, currentValue])

  const handleSelect = (value: string) => {
    const params = new URLSearchParams(searchParams.toString())
    params.set(searchParamKey, value)
    router.push(`${pathname}?${params.toString()}`)
  }

  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="secondary" className="flex items-center justify-between min-w-[160px]">
            <span>{selectedItem?.label || label}</span>
            <ChevronDown className="h-4 w-4 ml-2" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-56 mr-5" align="center" side="bottom" sideOffset={4}>
          {list.map((category) => (
            <DropdownMenuItem
              key={category.value}
              onSelect={() => handleSelect(category.value)}
              className={`
                cursor-pointer flex items-center justify-between
                ${currentValue === category.value ? 'bg-secondary text-primary' : 'text-muted-foreground'}
                `}
            >
              <p className="font-medium text-primary">{category.label}</p>
              <div>{category.icon}</div>
            </DropdownMenuItem>
          ))}
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
