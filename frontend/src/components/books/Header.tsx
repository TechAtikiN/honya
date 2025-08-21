import { Search } from 'lucide-react'
import AddNewBook from './AddNewBook'

export default function Header() {
  return (
    <div className="flex items-center justify-between space-x-5">
      <div className="flex items-center w-full -space-x-7">
        <Search className="text-muted-foreground h-4 w-4" />
        <input
          type="text"
          placeholder={'Search books by title, description, author, category'}
          className='pl-9 w-full form-input'
        />
      </div>

      <AddNewBook />
    </div>
  )
}
