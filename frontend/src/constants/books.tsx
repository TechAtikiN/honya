import { bookCategory } from "@/types/book";
import { Aperture, Atom, AudioWaveform, BadgeQuestionMark, BellDot, BellRing, BookHeart, BookPlus, BookType, CalendarClock, Component, CookingPot, HeartHandshake, History, Laugh, Layers, MapPlus, SquareLibrary, Star, UserRoundPen, Zap } from "lucide-react";

export const BOOK_CATEGORIES = [
  {
    value: 'all',
    label: "All Categories",
    icon: <SquareLibrary className="h-4 w-4" />
  },
  {
    value: 'fiction',
    label: "Fiction",
    icon: <Aperture className="h-4 w-4" />
  },
  {
    value: 'non-fiction',
    label: "Non-Fiction",
    icon: <AudioWaveform className="h-4 w-4" />
  },
  {
    value: 'science',
    label: "Science",
    icon: <Atom className="h-4 w-4" />
  },
  {
    value: 'history',
    label: "History",
    icon: <History className="h-4 w-4" />
  },
  {
    value: 'fantasy',
    label: "Fantasy",
    icon: <Component className='h-4 w-4' />
  },
  {
    value: 'mystery',
    label: "Mystery",
    icon: <BadgeQuestionMark className="h-4 w-4" />
  },
  {
    value: 'biography',
    label: "Biography",
    icon: <UserRoundPen className="h-4 w-4" />
  },
  {
    value: 'romance',
    label: "Romance",
    icon: <BookHeart className="h-4 w-4" />
  },
  {
    value: 'thriller',
    label: "Thriller",
    icon: <BellRing className="h-4 w-4" />
  },
  {
    value: 'self-help',
    label: "Self Help",
    icon: <HeartHandshake className="h-4 w-4" />
  },
  {
    value: 'cooking',
    label: "Cooking",
    icon: <CookingPot className="h-4 w-4" />
  },
  {
    value: 'travel',
    label: "Travel",
    icon: <MapPlus className="h-4 w-4" />
  },
  {
    value: 'classics',
    label: "Classics",
    icon: <Zap className="h-4 w-4" />
  },
  {
    value: 'comics',
    label: "Comics",
    icon: <Laugh className="h-4 w-4" />
  }
];

export const BOOK_SORT_OPTIONS = [
  {
    value: 'title',
    label: "Title",
    icon: <BookType className="h-4 w-4" />
  },
  {
    value: 'rating',
    label: "Rating",
    icon: <Star className="h-4 w-4" />
  },
  {
    value: 'recently_added',
    label: "Recently Added",
    icon: <BookPlus className="h-4 w-4" />
  },
  {
    value: 'recently_updated',
    label: "Recently Updated",
    icon: <BellDot className="h-4 w-4" />
  },
  {
    value: 'publication_year',
    label: "Publication Year",
    icon: <CalendarClock className="h-4 w-4" />
  },
  {
    value: "pages",
    label: "Number of Pages",
    icon: <Layers className="h-4 w-4" />
  }

]

export const MOCK_BOOKS_DATA = [
  {
    id: "1",
    title: "The Great Gatsby",
    authorName: "F. Scott Fitzgerald",
    description:
      "A classic American novel set in the Jazz Age, exploring themes of wealth, love, and the American Dream.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 1925,
    rating: 4.2,
    pages: 180,
    isbn: "978-0-7432-7356-5",
    createdAt: new Date("2024-01-15"),
    updatedAt: new Date("2024-01-15"),
  },
  {
    id: "2",
    title: "To Kill a Mockingbird",
    authorName: "Harper Lee",
    description: "A gripping tale of racial injustice and childhood innocence in the American South.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 1960,
    rating: 4.5,
    pages: 324,
    isbn: "978-0-06-112008-4",
    createdAt: new Date("2024-01-10"),
    updatedAt: new Date("2024-01-10"),
  },
  {
    id: "3",
    title: "Sapiens: A Brief History of Humankind",
    authorName: "Yuval Noah Harari",
    description:
      "An exploration of how Homo sapiens came to dominate the world through cognitive, agricultural, and scientific revolutions.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 2011,
    rating: 4.4,
    pages: 443,
    isbn: "978-0-06-231609-7",
    createdAt: new Date("2024-01-20"),
    updatedAt: new Date("2024-01-20"),
  },
  {
    id: "4",
    title: "The Pragmatic Programmer",
    authorName: "David Thomas, Andrew Hunt",
    description: "A guide to becoming a better programmer through practical advice and timeless principles.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 1999,
    rating: 4.3,
    pages: 352,
    isbn: "978-0-201-61622-4",
    createdAt: new Date("2024-01-25"),
    updatedAt: new Date("2024-01-25"),
  },
  {
    id: "5",
    title: "Dune",
    authorName: "Frank Herbert",
    description:
      "An epic science fiction novel set on the desert planet Arrakis, following Paul Atreides and his journey.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 1965,
    rating: 4.6,
    pages: 688,
    isbn: "978-0-441-17271-9",
    createdAt: new Date("2024-01-12"),
    updatedAt: new Date("2024-01-12"),
  },
  {
    id: "6",
    title: "The Midnight Library",
    authorName: "Matt Haig",
    description:
      "A philosophical novel about life, regret, and the infinite possibilities that exist between life and death.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 2020,
    rating: 4.1,
    pages: 288,
    isbn: "978-0-525-55948-1",
    createdAt: new Date("2024-02-01"),
    updatedAt: new Date("2024-02-01"),
  },
  {
    id: "7",
    title: "Educated",
    authorName: "Tara Westover",
    description: "A powerful memoir about education, family, and the struggle between loyalty and independence.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 2018,
    rating: 4.7,
    pages: 334,
    isbn: "978-0-399-59050-4",
    createdAt: new Date("2024-01-30"),
    updatedAt: new Date("2024-01-30"),
  },
  {
    id: "8",
    title: "The Seven Husbands of Evelyn Hugo",
    authorName: "Taylor Jenkins Reid",
    description: "A captivating novel about a reclusive Hollywood icon who finally decides to tell her story.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 2017,
    rating: 4.8,
    pages: 400,
    isbn: "978-1-5011-3981-2",
    createdAt: new Date("2024-02-05"),
    updatedAt: new Date("2024-02-05"),
  },
  {
    id: "9",
    title: "The Great Gatsby",
    authorName: "F. Scott Fitzgerald",
    description:
      "A classic American novel set in the Jazz Age, exploring themes of wealth, love, and the American Dream.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 1925,
    rating: 4.2,
    pages: 180,
    isbn: "978-0-7432-7356-5",
    createdAt: new Date("2024-01-15"),
    updatedAt: new Date("2024-01-15"),
  },
  {
    id: "10",
    title: "To Kill a Mockingbird",
    authorName: "Harper Lee",
    description: "A gripping tale of racial injustice and childhood innocence in the American South.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 1960,
    rating: 4.5,
    pages: 324,
    isbn: "978-0-06-112008-4",
    createdAt: new Date("2024-01-10"),
    updatedAt: new Date("2024-01-10"),
  },
  {
    id: "11",
    title: "Sapiens: A Brief History of Humankind",
    authorName: "Yuval Noah Harari",
    description:
      "An exploration of how Homo sapiens came to dominate the world through cognitive, agricultural, and scientific revolutions.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 2011,
    rating: 4.4,
    pages: 443,
    isbn: "978-0-06-231609-7",
    createdAt: new Date("2024-01-20"),
    updatedAt: new Date("2024-01-20"),
  },
  {
    id: "12",
    title: "The Pragmatic Programmer",
    authorName: "David Thomas, Andrew Hunt",
    description: "A guide to becoming a better programmer through practical advice and timeless principles.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 1999,
    rating: 4.3,
    pages: 352,
    isbn: "978-0-201-61622-4",
    createdAt: new Date("2024-01-25"),
    updatedAt: new Date("2024-01-25"),
  },
  {
    id: "13",
    title: "Dune",
    authorName: "Frank Herbert",
    description:
      "An epic science fiction novel set on the desert planet Arrakis, following Paul Atreides and his journey.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 1965,
    rating: 4.6,
    pages: 688,
    isbn: "978-0-441-17271-9",
    createdAt: new Date("2024-01-12"),
    updatedAt: new Date("2024-01-12"),
  },
  {
    id: "14",
    title: "The Midnight Library",
    authorName: "Matt Haig",
    description:
      "A philosophical novel about life, regret, and the infinite possibilities that exist between life and death.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 2020,
    rating: 4.1,
    pages: 288,
    isbn: "978-0-525-55948-1",
    createdAt: new Date("2024-02-01"),
    updatedAt: new Date("2024-02-01"),
  },
  {
    id: "15",
    title: "Educated",
    authorName: "Tara Westover",
    description: "A powerful memoir about education, family, and the struggle between loyalty and independence.",
    category: bookCategory.FICTION,
    image: "/assets/books/harry-potter.png",
    publicationYear: 2018,
    rating: 4.7,
    pages: 334,
    isbn: "978-0-399-59050-4",
    createdAt: new Date("2024-01-30"),
    updatedAt: new Date("2024-01-30"),
  },
  {
    id: "16",
    title: "The Seven Husbands of Evelyn Hugo",
    authorName: "Taylor Jenkins Reid",
    description: "A captivating novel about a reclusive Hollywood icon who finally decides to tell her story.",
    category: bookCategory.FICTION,
    image: "/assets/books/great-gatsby.png",
    publicationYear: 2017,
    rating: 4.8,
    pages: 400,
    isbn: "978-1-5011-3981-2",
    createdAt: new Date("2024-02-05"),
    updatedAt: new Date("2024-02-05"),
  },
]