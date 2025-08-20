import { Locale } from "@/i18n.config";
import MobileSidebar from "./mobile-sidebar";
import { LocaleDict } from "@/lib/locales";

interface AppbarProps {
  locale: Locale;
  translations: LocaleDict;
}

export default function Appbar({
  locale,
  translations,
}: AppbarProps) {
  return (
    <div className="flex items-center justify-end md:justify-end p-2">
      <MobileSidebar
        locale={locale}
        translations={translations}
      />
    </div>
  );
}
