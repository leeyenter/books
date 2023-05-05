import { useQuery } from "@tanstack/vue-query";
import { Book } from "../types/book.ts";

const BACKEND_URL = "http://localhost:5000";

const getBooks = async (): Promise<Book[]> => {
  const resp = await fetch(`${BACKEND_URL}/book/`);
  return await resp.json();
};

export const getBooksQuery = () =>
  useQuery({
    queryKey: ["books"],
    queryFn: getBooks,
  });
