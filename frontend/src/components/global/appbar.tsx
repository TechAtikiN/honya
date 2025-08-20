import MobileSidebar from "./mobile-sidebar";

export default function Appbar() {
  return (
    <div className="flex items-center justify-end md:justify-end p-2">
      <MobileSidebar />
    </div>
  );
}
