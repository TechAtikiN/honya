import DeleteBook from "@/components/book-details/DeleteBook";
import StarRating from "@/components/book-details/StarRating";
import UpdateBook from "@/components/books/UpdateBook";
import CustomLink from "@/components/global/custom-link";
import { Button } from "@/components/ui/button";
import { MOCK_BOOK_DATA } from "@/constants/books";
import { getLocale } from "@/i18n.config";
import { formatToAgo } from "@/lib/utils";
import { ArrowLeft, UserStar } from "lucide-react";
import Image from "next/image";

interface BookDetailPageProps {
  params: Promise<{ locale: string; bookId: string }>;
}

export default async function BookDetailPage({ params }: BookDetailPageProps) {
  const locale = await params;
  const lang = getLocale(locale.locale);

  return (
    <div className="flex flex-col space-y-10 h-[calc(100vh-30px)] overflow-auto invisible-scrollbar px-2 pb-10">
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
        <UpdateBook />
      </div>

      <div className="flex flex-col md:flex-row items-start justify-center space-x-0 md:space-x-10 space-y-5 md:space-y-0">
        {/* Left section  */}
        <div className="relative flex items-center justify-center w-full md:w-1/2">
          <div className="hidden md:block md:absolute md:h-[520px] md:w-[520px] md:rounded-full md:bg-secondary/35 md:z-0 shadow-lg" />
          <Image
            src={MOCK_BOOK_DATA.image || "/placeholder.png"}
            alt={MOCK_BOOK_DATA.title}
            width={350}
            height={520}
            className="rounded-md object-cover z-10 shadow-xl"
          />
        </div>

        {/* Right section */}
        <div className="flex flex-col items-start justify-center space-y-4 w-full md:w-1/2">
          <p className="bg-secondary border border-primary rounded-xl font-bold px-3 text-primary">{MOCK_BOOK_DATA.category}</p>
          <p className="text-5xl font-extrabold text-primary">{MOCK_BOOK_DATA.title}</p>
          <p className="text-lg font-medium text-primary/50 mt-2">- by {MOCK_BOOK_DATA.authorName}</p>
          <p className="text-justify text-lg font-normal text-primary mt-2">{MOCK_BOOK_DATA.description}</p>
        </div>
      </div>


      <div className='flex items-center justify-start w-full space-x-5 font-semibold text-primary/50'>
        <p>Created {formatToAgo(Number(MOCK_BOOK_DATA.createdAt))}</p>
        <p>|</p>
        <p>Updated {formatToAgo(Number(MOCK_BOOK_DATA.updatedAt))}</p>
      </div>

      {/* bottom  */}
      <div className="flex flex-col space-y-10">
        <div className="flex flex-col md:flex-row items-start justify-between">
          <div className="flex flex-col items-center justify-between space-y-1">
            <StarRating rating={MOCK_BOOK_DATA.rating} />
            <p className="">Rating</p>
          </div>

          <div className="flex flex-col items-center justify-between space-y-1">
            <p className="text-lg text-primary font-semibold">{MOCK_BOOK_DATA.publicationYear}</p>
            <p className="font-normal text-primary">Publication Year</p>
          </div>

          <div className="flex flex-col items-center justify-between space-y-1">
            <p className="text-lg text-primary font-semibold">{MOCK_BOOK_DATA.pages}</p>
            <p className="font-normal text-primary">Pages</p>
          </div>

          <div className="flex flex-col items-center justify-between space-y-1">
            <p className="text-lg text-primary font-semibold">{MOCK_BOOK_DATA.isbn}</p>
            <p className="font-normal text-primary">ISBN</p>
          </div>
        </div>
        <hr />


        {/* reviews */}
        <div className="flex flex-col space-y-3">
          <div className="flex items-center justify-between">
            <p className="text-2xl font-bold text-primary">Reviews</p>
            <Button variant="secondary" className="flex items-center">
              <UserStar className="h-4 w-4" />
              <span>Add Review</span>
            </Button>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-5">
            {MOCK_BOOK_DATA.reviews.length > 0 &&
              MOCK_BOOK_DATA.reviews.map((review, index) => (
                <div key={index} className="flex flex-col items-start justify-start space-y-2 bg-white p-4 rounded-lg shadow-md">
                  <p className="text-lg font-semibold text-primary">{review.reviewerName}</p>
                  <p className="text-sm text-primary/70">&quot;{review.reviewText}&quot;</p>
                </div>
              ))
            }
          </div>
        </div>
      </div>

      <DeleteBook />
    </div>
  )
}
