"use client";

import { useSidebarStore } from "@/stores/sidebar.store";
import { Languages, PanelLeftOpen, PanelRightOpen } from "lucide-react";
import { SIDEBAR_LINKS } from "@/constants/sidebar";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { Button } from "../ui/button";
import SidebarLink from "./sidebar-link";
import BrandLogo from "./brand-logo";

const LANGUAGES = [
  {
    name: "English",
    code: "en"
  },
  {
    name: "日本語",
    code: "ja"
  }
]

export default function Sidebar() {
  const { collapse, toggleCollapse } = useSidebarStore();

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
            <SidebarLink key={link.name} link={link} collapse={collapse} />
          ))}
        </div>
      </div>

      {/* footer */}
      <div className="w-full">
        <Popover>
          <PopoverTrigger asChild className="p-4">
            <Button
              variant={"default"}
              className="flex items-center justify-center w-full hover:cursor-pointer p-5"
            >
              <Languages className="h-5 w-5" />
              <p className={`text-white ${collapse ? "hidden" : ""}`}>
                Select Language
              </p>
            </Button>
          </PopoverTrigger>

          <PopoverContent className="w-full">
            <div className="space-y-2 w-full">
              {LANGUAGES.map((language) => (
                <Button
                  key={language.code}
                  variant="ghost"
                  size={"sm"}
                  className="w-full justify-start"
                >
                  {language.name}
                </Button>
              ))}
            </div>
          </PopoverContent>
        </Popover>
      </div>
    </div>
  );
}
