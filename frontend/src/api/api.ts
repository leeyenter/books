import { QueryClient, useMutation, useQuery } from "@tanstack/vue-query";
import { Book } from "../types/book.ts";

const BACKEND_URL = "http://localhost:5000";

const API = {
  book: {
    get: async (): Promise<Book[]> => {
      console.log("GET /book/");
      const resp = await fetch(`${BACKEND_URL}/book/`);
      return await resp.json();
    },
    post: async (book: Book) => {
      console.log("POST /book/");
      return await fetch(`${BACKEND_URL}/book/`, {
        method: "POST",
        body: JSON.stringify(book),
      });
    },
  },
};

export const getBooksQuery = () =>
  useQuery({
    queryKey: ["books"],
    queryFn: API.book.get,
    staleTime: 60_000,
  });

export const addBookMutation = (queryClient: QueryClient) =>
  useMutation({
    mutationFn: API.book.post,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["books"] }),
  });
