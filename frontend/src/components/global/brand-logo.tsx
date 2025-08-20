export default function BrandLogo({
  collapse = false,
}: {
  collapse?: boolean;
}) {
  return (
    <p
      className={`text-primary text-3xl font-bold transition-all duration-200 ease-in-out whitespace-nowrap overflow-hidden ${
        collapse ? 'opacity-0 w-0' : 'opacity-100 w-auto'
      }`}
    >
      Honya
    </p>
  );
}
