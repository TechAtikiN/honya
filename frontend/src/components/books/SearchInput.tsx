'use client'
import { useState, useCallback, ChangeEvent } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { Search } from 'lucide-react';
import debounce from 'lodash.debounce';

import { LocaleDict } from '@/lib/locales';
import { Locale } from '@/i18n.config';

interface SearchInputProps {
  translations: LocaleDict
  locale: Locale
}

export default function SearchInput({
  translations,
  locale
}: SearchInputProps) {
  const [query, setQuery] = useState<string>('');
  const router = useRouter();
  const searchParams = useSearchParams();

  const handleSearch = useCallback(
    debounce((newQuery: string) => {
      const params = new URLSearchParams(searchParams.toString());

      params.delete('page');

      if (newQuery) {
        params.set('query', newQuery);
      } else {
        params.delete('query');
      }

      router.push(`/?${params.toString()}`);
    }, 350),
    [searchParams, router]
  );

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setQuery(value);
    handleSearch(value);
  };

  return (
    <div className="flex flex-1 items-center md:w-1/2 -space-x-7">
      <input
        type="text"
        value={query}
        onChange={handleInputChange}
        placeholder={translations.page.home.filters?.searchBooksPlaceholder || "Search books..."}
        className="w-full form-input"
      />
      <Search className="text-muted-foreground h-4 w-4" />
    </div>
  );
} 
