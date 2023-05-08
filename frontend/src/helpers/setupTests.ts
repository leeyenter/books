import { setupServer } from "msw/node";
import { rest } from "msw";
import { mockBook } from "../types/book.ts";

const restHandlers = [
  rest.get("http://localhost:5000/book/", (_, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json([
        mockBook({ title: "A Long Obedience" }),
        mockBook(),
        mockBook(),
      ])
    );
  }),
];

const server = setupServer(...restHandlers);

beforeAll(() => server.listen({ onUnhandledRequest: "error" }));
afterEach(() => server.resetHandlers());
afterAll(() => server.close());
