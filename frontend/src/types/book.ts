export enum bookCategory {
    FICTION = "fiction",
    NON_FICTION = "non_fiction",
    SCIENCE = "science",
    HISTORY = "history",
    FANTASY = "fantasy",
    MYSTERY = "mystery",
    BIOGRAPHY = "biography",
    ROMANCE = "romance",
    THRILLER = "thriller",
    SELF_HELP = "self_help",
    COOKING = "cooking",
    TRAVEL = "travel",
    CLASSICS = "classics",
    COMICS = "comics",
}

export type Book = {
    id: string;
    title: string;
    description: string;
    authorName: string;
    category: bookCategory;
    image: string;
    publicationYear: number;
    rating: number;
    pages: number;
    isbn: string;
    createdAt: Date;
    updatedAt: Date;
}