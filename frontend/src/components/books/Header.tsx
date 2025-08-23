import { Github, Search } from 'lucide-react'
import AddNewBook from './AddNewBook'
import Link from 'next/link'

export default function Header() {
  return (
    <div className="flex items-center justify-between gap-x-2">
      <div className="flex w-full md:w-2/3 items-center gap-x-2 md:gap-x-4">
        <div className="flex flex-1 items-center md:w-1/2 -space-x-7">
          <input
            type="text"
            placeholder={'Search books...'}
            className='w-full form-input'
          />
          <Search className="text-muted-foreground h-4 w-4" />
        </div>
        <AddNewBook />
      </div>

      <Link
        className='bg-primary rounded-full p-2 hover:opacity-80 transition'
        href="https://github.com/TechAtikiN" target='_blank'>
        <Github className="h-5 w-5 text-white" />
      </Link>
    </div>
  )
}
