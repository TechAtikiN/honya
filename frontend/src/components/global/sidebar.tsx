"use client";

import { useSidebarStore } from "@/stores/sidebar.store";
import { Blocks, PanelLeftOpen, PanelRightOpen, ScrollText } from "lucide-react";

import { Button } from "../ui/button";
import SidebarLink from "./sidebar-link";
import BrandLogo from "./brand-logo";
import { Locale } from "@/i18n.config";
import LanguageSelector from "./language-selector";
import { LocaleDict } from "@/lib/locales";


interface SidebarProps {
  locale: Locale;
  translations: LocaleDict
}

export default function Sidebar({ locale, translations }: SidebarProps) {
  const { collapse, toggleCollapse } = useSidebarStore();

  const SIDEBAR_LINKS = [
    {
      name: translations.sidebar.navigation.books,
      href: "/",
      icon: <ScrollText className="w-5 h-5" />,
    },
    {
      name: translations.sidebar.navigation.analytics,
      href: "/analytics",
      icon: <Blocks className="w-5 h-5" />,
    },
  ];

  return (
    <div
      className={`h-full hidden md:flex flex-col items-start justify-between p-3 md:py-6 transition-all duration-200 ease-in-out
       ${collapse ? "w-16" : "w-[270px]"}
   `}
    >
      <div className="space-y-8 w-full">
        {/* logo  */}
        <div className="flex items-center justify-between">
          <BrandLogo collapse={collapse} />
          <Button
            className="hidden md:block"
            onClick={toggleCollapse}
            variant={"ghost"}
            size={"sm"}
          >
            {collapse ? (
              <PanelLeftOpen className="h-5 w-5" />
            ) : (
              <PanelRightOpen className="h-5 w-5" />
            )}
          </Button>
        </div>

        {/* links */}
        <div className="space-y-2">
          {SIDEBAR_LINKS.map((link) => (
            <SidebarLink
              key={link.name} link={link} collapse={collapse} locale={locale}
            />
          ))}
        </div>
      </div>

      {/* footer */}
      <LanguageSelector
        collapse={collapse}
        locale={locale}
      />
    </div>
  );
}
