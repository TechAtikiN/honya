import { Github } from 'lucide-react'
import AddNewBook from './AddNewBook'
import Link from 'next/link'
import SearchInput from './SearchInput'
import { LocaleDict } from '@/lib/locales'
import { Locale } from '@/i18n.config'

interface HeaderProps {
  translations: LocaleDict
  locale: Locale
}

export default function Header({
  translations,
  locale
}: HeaderProps) {
  return (
    <div className="flex items-center justify-between space-x-3 md:space-x-0">
      <div className="flex w-full md:w-2/3 items-center gap-x-2 md:gap-x-4">
        <SearchInput
          translations={translations}
        />
        <AddNewBook
          translations={translations}
          locale={locale}
        />
      </div>

      <Link
        className='bg-primary rounded-full p-2 hover:opacity-80 transition'
        href="https://github.com/TechAtikiN" target='_blank'>
        <Github className="h-5 w-5 ml-[2px] text-white fill-white" />
      </Link>
    </div>
  )
}
