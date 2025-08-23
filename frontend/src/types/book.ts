export enum bookCategory {
    FICTION = "fiction",
    NON_FICTION = "non_fiction",
    SCIENCE = "science",
    HISTORY = "history",
    FANTASY = "fantasy",
    MYSTERY = "mystery",
    THRILLER = "thriller",
    COOKING = "cooking",
    TRAVEL = "travel",
    CLASSICS = "classics",
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