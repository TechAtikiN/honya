import { Locale } from '@/i18n.config';
import MobileSidebar from './mobile-sidebar';

interface AppbarProps {
  locale: Locale;
  sidebarLinks: { name: string; href: string; icon: React.ReactNode }[];
}

export default function Appbar({ locale, sidebarLinks }: AppbarProps) {
  return <MobileSidebar locale={locale} sidebarLinks={sidebarLinks} />;
}
