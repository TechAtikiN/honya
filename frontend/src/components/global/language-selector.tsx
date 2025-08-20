'use client';

import { Globe, Languages } from 'lucide-react';
import { Button } from '../ui/button';
import { Popover, PopoverContent, PopoverTrigger } from '../ui/popover';
import { i18n, Locale } from '@/i18n.config';
import { usePathname } from 'next/navigation';
import Link from 'next/link';

interface LanguageSelectorProps {
  collapse: boolean;
  locale: Locale;
}

export default function LanguageSelector({
  collapse,
  locale,
}: LanguageSelectorProps) {
  const pathname = usePathname();

  const getRedirectPathName = (loc: string) => {
    if (!pathname) return '/';

    const pathnameIsMissingLocale = i18n.locales.every(
      (locale) =>
        !pathname.startsWith(`/${locale}/`) && pathname !== `/${locale}`
    );

    if (pathnameIsMissingLocale) {
      if (loc === i18n.defaultLocale) return pathname;
      return `/${loc}${pathname}`;
    }
    if (loc === i18n.defaultLocale) {
      const segments = pathname.split('/');
      const isHome = segments.length === 2;
      if (isHome) return '/';

      segments.splice(1, 1);
      return segments.join('/');
    }

    const segments = pathname.split('/');
    segments[1] = loc;
    return segments.join('/');
  };

  const setLocaleAndReload = (loc: string) => {
    // Set the locale in a cookie to persist the user's choice
    document.cookie = `NEXT_LOCALE=${loc}; path=/; max-age=31536000`; // 1 year cookie

    // Redirect to the new locale and reload the page to apply changes
    const newPath = getRedirectPathName(loc);
    window.location.href = newPath;
  };

  return (
    <Popover>
      <PopoverTrigger asChild className='p-4'>
        <Button
          variant={'default'}
          className='flex items-center justify-between w-full hover:cursor-pointer p-5'
        >
          <p
            className={`text-white transition-all duration-200 ease-in-out whitespace-nowrap overflow-hidden ${
              collapse ? 'opacity-0 w-0' : 'opacity-100 w-auto'
            }`}
          >
            Select Language
          </p>
          <Globe className='h-5 w-5' />
        </Button>
      </PopoverTrigger>

      <PopoverContent
        className='w-full min-w-40'
        align='end'
        side='right'
        sideOffset={10}
      >
        <div className='flex flex-col items-start space-y-2 w-full'>
          {i18n.locales.map((lang) => (
            <Button
              key={lang}
              variant={'link'}
              className={`justify-start w-full hover:no-underline ${
                lang === locale ? 'bg-neutral-200' : ''
              }`}
              onClick={() => setLocaleAndReload(lang)}
            >
              {lang === Locale.EN ? 'English' : '日本語'}
            </Button>
          ))}
        </div>
      </PopoverContent>
    </Popover>
  );
}
