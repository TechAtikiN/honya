import Appbar from "@/components/global/appbar";
import Sidebar from "@/components/global/sidebar";

export default function MainLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="w-full flex bg-accent">
      {/* Sidebar */}
      <div className="h-[calc(100vh)]">
        <Sidebar />
      </div>

      <div className="w-full bg-white md:m-3 md:ml-0 rounded-sm h-[calc(100vh-10px)] overflow-auto invisible-scrollbar">
        {/* Appbar */}
        <Appbar />

        {/* Content */}
        <main className="max-w-7xl mx-auto w-full px-2 md:px-6">
          {children}
        </main>
      </div>
    </div>
  );
}
