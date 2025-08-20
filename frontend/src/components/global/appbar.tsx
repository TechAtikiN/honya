import { Locale } from "@/i18n.config";
import { LocaleDict } from "@/lib/locales";
import MobileSidebar from "./mobile-sidebar";

interface AppbarProps {
  locale: Locale;
  sidebarLinks: { name: string; href: string; icon: React.ReactNode }[];
}

export default function Appbar({
  locale,
  sidebarLinks,
}: AppbarProps) {
  return (
    <div className="flex items-center justify-end p-2">
      <MobileSidebar
        locale={locale}
        sidebarLinks={sidebarLinks}
      />
    </div>
  );
}
