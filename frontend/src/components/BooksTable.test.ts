import { mount } from "@vue/test-utils";
import { Book } from "../types/book.ts";
import BooksTable from "./BooksTable.vue";

test("Books Table component", async () => {
  const wrapper = (books: Book[]) =>
    mount(BooksTable, {
      props: { books },
    });

  it.todo("has an empty state");
  it.todo("renders books correctly");
});
