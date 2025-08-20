
export default function BrandLogo({ collapse = false }: { collapse?: boolean }) {
    return (
        <p
            className={`text-primary text-3xl font-bold ${collapse ? "hidden" : ""
                }`}
        >
            Honya
        </p>
    )
}
