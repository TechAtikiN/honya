'use client';

import { Locale } from '@/i18n.config';
import { LocaleDict } from '@/lib/locales';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import CustomLink from './custom-link';

interface SidebarLinkProps {
  link: {
    name: string;
    href: string;
    icon: React.ReactNode;
  };
  collapse?: boolean;
  locale: Locale;
}

export default function SidebarLink({
  link,
  collapse,
  locale,
}: SidebarLinkProps) {
  const pathname = usePathname();

  const isHome = link.href === '/';
  const isActive = isHome
    ? pathname === '/'
    : pathname === link.href || pathname.startsWith(link.href + '/');

  return (
    <CustomLink
      locale={locale}
      href={link.href}
      className={`flex items-center justify-start space-x-3 p-2 rounded-md
        ${isActive ? 'bg-secondary font-medium' : 'font-normal'}
        `}
    >
      <div>{link.icon}</div>
      <p
        className={`transition-all duration-200 ease-in-out whitespace-nowrap overflow-hidden ${
          collapse ? 'opacity-0 w-0' : 'opacity-100 w-auto'
        }`}
      >
        {link.name}
      </p>
    </CustomLink>
  );
}
