import { Search } from 'lucide-react'

export default function Header() {
  return (
    <div className="flex items-center justify-between space-x-5">
      <div className="flex items-center w-full md:w-1/2 -space-x-7">
        <input
          type="text"
          placeholder={'Search books by title, description, author, category'}
          className='w-full form-input'
        />
        <Search className="text-muted-foreground h-4 w-4" />

      </div>

    </div>
  )
}
