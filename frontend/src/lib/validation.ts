import { z } from "zod";
import { bookCategory } from "@/types/book";

export const bookFormSchema = z.object({
  title: z
    .string()
    .trim()
    .min(1, "Title is required")
    .min(2, "Title must be at least 2 characters long")
    .max(200, "Title cannot exceed 200 characters"),

  description: z
    .string()
    .trim()
    .min(1, "Description is required")
    .min(10, "Description must be at least 10 characters long")
    .max(1000, "Description cannot exceed 1000 characters"),

  category: z.nativeEnum(bookCategory, {
    message: "Please select a valid book category",
  }),

  image: z
    .string()
    .url("Image must be a valid URL")
    .optional()
    .or(z.literal("")), // Allow empty string for optional image

  rating: z
    .number()
    .min(0, "Rating must be between 0 and 5")
    .max(5, "Rating must be between 0 and 5"),

  publicationYear: z
    .union([
      z.number()
        .int({ message: "Publication year must be a whole number" })
        .min(1950, { message: "Publication year must be 1950 or later" })
        .max(new Date().getFullYear(), {
          message: `Publication year cannot be in the future`,
        }),
      z.nan()
    ])
    .refine((val) => !isNaN(val), {
      message: "Publication year must be a valid number",
    }),

  pages: z
    .union([
      z.number()
        .int({ message: "Number of pages must be a whole number" })
        .min(1, { message: "Book must have at least 1 page" }),
      z.nan()
    ])
    .refine((val) => !isNaN(val), {
      message: "Number of pages must be a valid number",
    }),


  isbn: z
    .string()
    .trim()
    .min(1, "ISBN is required")
    .max(20, "ISBN cannot exceed 20 characters"),

  authorName: z
    .string()
    .trim()
    .min(1, "Author name is required")
    .min(2, "Author name must be at least 2 characters long")
    .max(100, "Author name too long")
    .regex(/^[a-zA-Z\s\.\-']+$/, "Author name can only contain letters, spaces, periods, hyphens, and apostrophes"),
});

export const reviewFormSchema = z.object({
  name: z
    .string()
    .trim()
    .min(1, "Name is required")
    .min(2, "Name must be at least 2 characters long")
    .max(100, "Name cannot exceed 100 characters"),
  email: z
    .string()
    .trim()
    .min(1, "Email is required")
    .email("Email must be a valid email address")
    .max(100, "Email cannot exceed 100 characters"),
  content: z
    .string()
    .trim()
    .min(1, "Review content is required")
    .max(1000, "Review content cannot exceed 1000 characters"),
});