import { type ClassValue, clsx } from "clsx";
import dayjs from "dayjs";
import { twMerge } from "tailwind-merge";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(relativeTime);

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const getMethodColor = (method: string) => {
  switch (method) {
    case "GET":
      return "bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200";
    case "POST":
      return "bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200";
    case "PUT":
      return "bg-orange-100 text-orange-800 dark:bg-orange-900 dark:text-orange-200";
    case "DELETE":
      return "bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200";
    case "PATCH":
      return "bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-200";
    default:
      return "bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200";
  }
};


export const formattedTime = (time: string) => {
  return dayjs(time).fromNow();
};
