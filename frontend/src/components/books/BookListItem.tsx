import { Book } from "@/types/book"
import Image from "next/image"
import Link from "next/link"
import HintLabel from "../global/hint-label"

interface BookListItemProps {
  book: Book
}

export default function BookListItem({
  book
}: BookListItemProps) {
  return (
    <Link
      href={`/books/${book.id}`}
      className="flex flex-col items-center justify-center space-y-3"
    >
      <div className="mx-auto">
        <div className="relative w-[160px] h-[220px] rounded-md overflow-hidden drop-shadow-md shadow-lg">
          <Image
            src={book.image || "/placeholder.png"}
            alt={book.title}
            fill
            className="object-fill"
          />
        </div>
      </div>
      <div>
        {book.title.length > 16 ? (
          <HintLabel
            label={book.title}
            side="bottom"
          >
            <p className="font-semibold text-primary">
              {book.title.slice(0, 16)}...
            </p>
          </HintLabel>
        ) : (
          <p className="font-semibold text-primary">
            {book.title}
          </p>
        )}
        <p className="font-medium text-sm text-primary/50">{book.authorName}</p>
      </div>
    </Link>
  )
}
