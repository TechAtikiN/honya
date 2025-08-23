import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import { LocaleDict } from "./locales";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function formatToAgo(date: Date | string | number, translations?: LocaleDict) {
  const now = new Date();
  const pastDate = typeof date === 'number'
    ? new Date(date * 1000)
    : new Date(date);

  const seconds = Math.floor((now.getTime() - pastDate.getTime()) / 1000);

  // Get localized strings with fallbacks
  const timeAgo = translations?.page?.bookDetails?.timeAgo;
  const justNow = timeAgo?.justNow || 'just now';
  const secondsAgo = timeAgo?.secondsAgo || 'seconds ago';
  const minutesAgo = timeAgo?.minutesAgo || 'minutes ago';
  const oneHourAgo = timeAgo?.oneHourAgo || '1 hour ago';
  const hoursAgo = timeAgo?.hoursAgo || 'hours ago';
  const daysAgo = timeAgo?.daysAgo || 'days ago';

  if (seconds < 5) return justNow;
  if (seconds < 60) return `${seconds} ${secondsAgo}`;

  const minutes = Math.floor(seconds / 60);
  if (minutes < 60) return `${minutes} ${minutesAgo}`;
  if (minutes < 120) return oneHourAgo;

  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours} ${hoursAgo}`;
  const days = Math.floor(hours / 24);
  return `${days} ${daysAgo}`;
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

// Generic helper to resolve form validation error messages
export function getNestedTranslation(key: string | undefined, translations: LocaleDict, formType: 'book' | 'review' = 'book'): string {
  if (!key || !key.includes('form.error.')) return key || '';

  const errorKey = key.replace('form.error.', '');

  // Handle different form types based on the actual locale structure
  if (formType === 'book') {
    return (translations.page.home.bookForm.errorMessages as Record<string, string>)?.[errorKey] || key;
  } else if (formType === 'review') {
    return (translations.page.bookDetails.reviewForm as Record<string, string>)?.[errorKey] || key;
  }

  return key;
}

// Generic helper to resolve action message keys for both book and review actions
export function resolveActionMessage(messageKey: string, translations: LocaleDict): string {
  if (!messageKey.includes('actions.')) return messageKey;

  // Handle book actions
  if (messageKey.includes('actions.book.')) {
    const actionKey = messageKey.replace('actions.book.', '');
    const bookActions = translations?.actions?.book as Record<string, string>;
    return bookActions?.[actionKey] || messageKey;
  }

  // Handle review actions
  if (messageKey.includes('actions.review.')) {
    const actionKey = messageKey.replace('actions.review.', '');
    const reviewActions = translations?.actions?.review as Record<string, string>;
    return reviewActions?.[actionKey] || messageKey;
  }

  return messageKey;
};