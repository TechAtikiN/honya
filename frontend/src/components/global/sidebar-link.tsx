"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

interface SidebarLinkProps {
  link: {
    name: string;
    href: string;
    icon: React.ReactNode;
  };
  collapse?: boolean;
}

export default function SidebarLink({ link, collapse }: SidebarLinkProps) {
  const pathname = usePathname();

  const isHome = link.href === "/";
  const isActive = isHome
    ? pathname === "/"
    : pathname === link.href || pathname.startsWith(link.href + "/");

  return (
    <Link
      href={link.href}
      className={`flex items-center justify-start space-x-3 p-2 rounded-md
        ${isActive ? "bg-secondary font-medium" : ""}
        `}
    >
      <div>{link.icon}</div>
      <p className={`${collapse ? "hidden" : ""}`}>{link.name}</p>
    </Link>
  );
}
