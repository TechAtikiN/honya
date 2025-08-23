import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function formatToAgo(date: Date | string | number) {
  const now = new Date();
  const pastDate = typeof date === 'number'
    ? new Date(date * 1000)
    : new Date(date);

  const seconds = Math.floor((now.getTime() - pastDate.getTime()) / 1000);
  if (seconds < 5) return 'just now';
  if (seconds < 60) return `${seconds} seconds ago`;

  const minutes = Math.floor(seconds / 60);
  if (minutes < 60) return `${minutes} minutes ago`;
  if (minutes < 120) return `1 hour ago`;

  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours} hours ago`;
  const days = Math.floor(hours / 24);
  return `${days} days ago`;
}

export function getPagination(page: string | string[] | undefined) {
  const currentPage = page ? parseInt(page as string, 10) : 1;
  const limit = 10;
  return { currentPage, limit };
}


export function getFilters(searchParams: { [key: string]: string | string[] | undefined }) {
  const filters: { [key: string]: string } = {};

  for (const key in searchParams) {
    const value = searchParams[key];
    if (typeof value === 'string' && value.trim() !== '') {
      filters[key] = value;
    }
  }

  return filters;
}

export function formatTitleCase(str: string) {
  return str.replace(/\w\S*/g, (txt) => {
    return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
  });
}