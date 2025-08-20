import { getLocale } from "@/i18n.config";
import { getDictionary } from "@/lib/locales";

interface HomePageProps {
    params: Promise<{ locale: string }>;
}

export default async function Analytics({
    params
}: HomePageProps) {
    const locale = await params;
    const lang = getLocale(locale.locale);
    const { page } = await getDictionary(lang);

    return (
        <div>
            {page.analytics.title}
        </div>
    );
}
