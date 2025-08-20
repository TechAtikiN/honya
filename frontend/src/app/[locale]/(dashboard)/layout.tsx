import Appbar from '@/components/global/appbar';
import Sidebar from '@/components/global/sidebar';
import { getSidebarLinks } from '@/constants/sidebar';
import { getLocale } from '@/i18n.config';
import { getDictionary } from '@/lib/locales';

export default async function MainLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  const lang = getLocale(locale);
  const translations = await getDictionary(lang);
  const sidebarLinks = getSidebarLinks(translations)

  return (
    <div className='w-full flex bg-accent'>
      {/* Sidebar */}
      <div className='h-[calc(100vh)]'>
        <Sidebar locale={lang} sidebarLinks={sidebarLinks} />
      </div>

      <div className='w-full bg-white rounded-sm md:m-3 md:mr-0 md:mb-0 overflow-auto invisible-scrollbar'>
        {/* Appbar */}
        <Appbar locale={lang} sidebarLinks={sidebarLinks} />

        {/* Content */}
        <main className='max-w-7xl mx-auto w-full px-2 md:px-6'>
          {children}
        </main>
      </div>
    </div>
  );
}
