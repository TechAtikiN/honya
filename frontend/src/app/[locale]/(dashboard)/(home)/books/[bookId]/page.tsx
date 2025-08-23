import { getBookDetails } from "@/actions/book.actions";
import { buttonVariants } from "@/components/ui/button";
import { getLocale } from "@/i18n.config";
import { cn, formatTitleCase, formatToAgo } from "@/lib/utils";
import { ArrowLeft, BookAlert } from "lucide-react";
import Image from "next/image";
import AddReview from "@/components/book-details/AddReview";
import DeleteBook from "@/components/book-details/DeleteBook";
import StarRating from "@/components/book-details/StarRating";
import UpdateBook from "@/components/books/UpdateBook";
import CustomLink from "@/components/global/custom-link";

interface BookDetailPageProps {
  params: Promise<{ locale: string; bookId: string }>;
}

export default async function BookDetailPage({ params }: BookDetailPageProps) {
  const { locale, bookId } = await params;
  const lang = getLocale(locale);

  const bookDetails = await getBookDetails(bookId)

  if (!bookDetails) {
    return (
      <div className="flex flex-col items-center justify-center gap-y-5 h-[calc(100vh-30px)]">
        <div className="flex flex-col items-center justify-center gap-2">
          <BookAlert className="h-16 w-16 text-primary mx-auto mb-4" />
          <p className="text-2xl font-bold text-primary">Book not found</p>
        </div>
        <CustomLink
          href={`/`}
          locale={lang}
          className={cn(buttonVariants({ variant: 'default' }), '')}
        >
          Go back to Library
        </CustomLink>
      </div>
    )
  }

  return (
    <div className="flex flex-col space-y-10 min-h-[calc(100vh-30px)] overflow-auto invisible-scrollbar px-2 pb-10">
      {/* back button */}
      <div className="flex items-center justify-between w-full">
        <CustomLink
          href={`/`}
          locale={lang}
          className="flex items-center space-x-2"
        >
          <ArrowLeft className="h-4 w-4" />
          <span className="font-medium text-primary hover:underline underline-offset-4">View all books</span>
        </CustomLink>
        <UpdateBook bookDetails={bookDetails} />
      </div>

      <div className="flex flex-col md:flex-row items-start justify-center space-x-0 md:space-x-10 space-y-5 md:space-y-0">
        {/* Left section  */}
        <div className="relative flex items-center justify-center w-full md:w-1/2">
          <div className="hidden md:block md:absolute md:h-[520px] md:w-[520px] md:rounded-full md:bg-secondary/35 md:z-0 shadow-lg" />
          <Image
            src={bookDetails.image || '/assets/books/book-placeholder.png'}
            alt={bookDetails.title || 'Book Image'}
            width={350}
            height={520}
            className="rounded-md object-cover z-10 shadow-xl"
          />
        </div>

        {/* Right section */}
        <div className="flex flex-col items-start justify-center space-y-4 w-full md:w-1/2">
          <p className="bg-secondary border border-primary rounded-xl font-bold px-3 text-primary">{formatTitleCase(bookDetails.category)}</p>
          <p className="text-5xl font-extrabold text-primary">{bookDetails.title}</p>
          <p className="text-lg font-medium text-primary/50 mt-2">- by {bookDetails.author_name}</p>
          <p className="text-justify text-lg font-normal text-primary mt-2">{bookDetails.description}</p>
        </div>
      </div>


      <div className='flex items-center justify-start w-full space-x-5 font-semibold text-primary/50'>
        <p>Created {formatToAgo(Number(bookDetails.created_at))}</p>
        <p>|</p>
        <p>Updated {formatToAgo(Number(bookDetails.updated_at))}</p>
      </div>

      {/* bottom  */}
      <div className="flex flex-col space-y-10">
        <div className="flex flex-col md:flex-row items-start justify-between">
          <div className="flex flex-col items-center justify-between space-y-1">
            <StarRating rating={bookDetails.rating} />
            <p className="">Rating</p>
          </div>

          <div className="flex flex-col items-center justify-between space-y-1">
            <p className="text-lg text-primary font-semibold">{bookDetails.publication_year}</p>
            <p className="font-normal text-primary">Publication Year</p>
          </div>

          <div className="flex flex-col items-center justify-between space-y-1">
            <p className="text-lg text-primary font-semibold">{bookDetails.pages}</p>
            <p className="font-normal text-primary">Pages</p>
          </div>

          <div className="flex flex-col items-center justify-between space-y-1">
            <p className="text-lg text-primary font-semibold">{bookDetails.isbn}</p>
            <p className="font-normal text-primary">ISBN</p>
          </div>
        </div>
        <hr />


        {/* reviews */}
        <div className="flex flex-col space-y-3">
          <div className="flex items-center justify-between">
            <p className="text-2xl font-bold text-primary">Reviews</p>
            <AddReview bookId={bookDetails.id} />
          </div>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-5">
            {bookDetails?.reviews && bookDetails.reviews.data.length > 0 ?
              bookDetails.reviews.data.map((review, index) => (
                <div key={index} className="flex flex-col items-start justify-start space-y-2 bg-accent p-4 rounded-md shadow-sm">
                  <div className="flex items-center gap-3">
                    <div className="flex items-center rounded-full bg-primary h-9 w-9 justify-center text-white font-bold">
                      {review.name.charAt(0).toUpperCase()}
                    </div>
                    <div>
                      <p className="text-lg font-semibold text-primary">
                        {review.name}
                      </p>
                      <p className="text-sm text-primary/50 font-normal">{formatToAgo(review.created_at)}</p>
                    </div>
                  </div>
                  <p className="text-sm text-primary/70">&quot;{review.content}&quot;</p>
                </div>
              )) : (
                <p className="text-primary/70">No reviews yet.</p>
              )
            }
          </div>
        </div>
      </div>

      <DeleteBook
        bookId={bookDetails.id}
      />
    </div>
  )
}
