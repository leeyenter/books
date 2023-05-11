import { QueryClient, useMutation, useQuery } from "@tanstack/vue-query";
import { Book } from "../types/book.ts";
import { router } from "../main.ts";
import {
  AuthenticationResponseJSON,
  RegistrationResponseJSON,
} from "@simplewebauthn/typescript-types";

const BACKEND_URL = "http://localhost:5000";

const getAuthJWT = (): string => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; auth=`);
  if (parts.length === 1) return "";
  const cookie = parts.pop()?.split(";").shift();
  console.log(cookie);
  return cookie || "";
};

export const API = {
  auth: {
    checkLogin: async () => {
      const resp = await fetch(`${BACKEND_URL}/auth/`, {
        headers: { auth: getAuthJWT() },
      });
      return resp.ok;
    },
    checkRegistration: async (): Promise<number> => {
      const resp = await fetch(`${BACKEND_URL}/auth/registration/`);
      const jsonObj = await resp.json();
      return jsonObj["num"];
    },
    beginRegistration: async () => {
      const resp = await fetch(`${BACKEND_URL}/auth/registration/begin`, {
        method: "POST",
      });
      return await resp.json();
    },
    finishRegistration: async (
      attResp: RegistrationResponseJSON
    ): Promise<string> => {
      const resp = await fetch(`${BACKEND_URL}/auth/registration/finish`, {
        method: "POST",
        body: JSON.stringify(attResp),
      });
      const jsonObj = await resp.json();
      return jsonObj["token"];
    },
    beginLogin: async () => {
      const resp = await fetch(`${BACKEND_URL}/auth/login/begin`, {
        method: "POST",
      });
      return await resp.json();
    },
    finishLogin: async (
      attResp: AuthenticationResponseJSON
    ): Promise<string> => {
      const resp = await fetch(`${BACKEND_URL}/auth/login/finish`, {
        method: "POST",
        body: JSON.stringify(attResp),
      });
      const jsonObj = await resp.json();
      return jsonObj["token"];
    },
  },
  book: {
    get: async (): Promise<Book[]> => {
      console.log("GET /book/");
      const resp = await fetch(`${BACKEND_URL}/book/`);
      return await resp.json();
    },
    post: async (book: Book) => {
      console.log("POST /book/");
      return await fetch(`${BACKEND_URL}/book/${book.id}`, {
        method: "POST",
        headers: { auth: getAuthJWT() },
        body: JSON.stringify(book),
      });
    },
  },
};

export const checkLoginQuery = () =>
  useQuery({
    queryKey: ["login"],
    queryFn: API.auth.checkLogin,
  });

export const getBooksQuery = () =>
  useQuery({
    queryKey: ["books"],
    queryFn: API.book.get,
    staleTime: 60_000,
  });

export const addBookMutation = (queryClient: QueryClient) =>
  useMutation({
    mutationFn: API.book.post,
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["books"] });
      await router.push("/");
    },
  });
