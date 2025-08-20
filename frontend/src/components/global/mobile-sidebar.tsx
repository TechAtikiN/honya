import { SIDEBAR_LINKS } from "@/constants/sidebar";
import SidebarLink from "./sidebar-link";
import { Menu } from "lucide-react";
import {
    Sheet,
    SheetContent,
    SheetDescription,
    SheetHeader,
    SheetTitle,
    SheetTrigger,
} from "../ui/sheet";
import { Button } from "../ui/button";
import BrandLogo from "./brand-logo";

export default function MobileSidebar() {
    return (
        <div className="md:hidden flex items-center justify-between w-full">
            {/* Logo  */}
            <BrandLogo collapse={false} />

            {/* Mobile Sidebar Trigger and Content */}
            <Sheet>
                <SheetTrigger asChild>
                    <Button variant="ghost">
                        <Menu className="text-primary h-5 w-5" />
                    </Button>
                </SheetTrigger>
                <SheetContent side="left" className="w-11/12 h-full">
                    <SheetHeader className="hidden">
                        <SheetTitle className="sr-only">Sidebar</SheetTitle>
                        <SheetDescription className="sr-only">
                            Honya, your personal library.
                        </SheetDescription>
                    </SheetHeader>

                    <div
                        className={`h-full block md:hidden flex-col items-start justify-between p-3
   `}
                    >
                        <div className="space-y-5 w-full">
                            {/* logo  */}
                            <BrandLogo />
                            {/* links */}
                            <div className="space-y-2">
                                {SIDEBAR_LINKS.map((link) => (
                                    <SidebarLink key={link.name} link={link} />
                                ))}
                            </div>
                        </div>
                    </div>

                    {/* footer */}
                    <div>footer</div>
                </SheetContent>
            </Sheet>
        </div>
    );
}
