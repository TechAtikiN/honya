'use client';

import { Locale } from '@/i18n.config';
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

  const isBookHomePage = pathname === `/`
  const isBookDetailPage = pathname?.startsWith(`/books/`)
  const isActive = link.href === pathname || (link.href === '/' && (isBookHomePage || isBookDetailPage));

  return (
    <CustomLink
      locale={locale}
      href={link.href}
      className={`flex items-center justify-start space-x-3 p-2 rounded-sm
        ${isActive ? 'bg-secondary border font-medium border-primary/60' : 'font-normal'}
        `}
    >
      <div>{link.icon}</div>
      <p
        className={`${collapse ? 'opacity-0 w-0' : 'opacity-100 w-auto'
          }`}
      >
        {link.name}
      </p>
    </CustomLink>
  );
}
