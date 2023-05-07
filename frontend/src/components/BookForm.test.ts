import { Book, mockBook } from "../types/book.ts";
import BookForm from "./BookForm.vue";
import { mount, VueWrapper } from "@vue/test-utils";

class FormObj {
  constructor(private app: VueWrapper) {}

  setInputValue(label: string, value: string) {
    this.app.find({ ref: label }).setValue(value);
  }

  get authorInputs() {
    return this.app.findAll('[data-test="author"]');
  }

  async clickAddAuthor() {
    const addAuthorButton = this.app
      .findAll("button")
      .find((x) => x.text().match(/Add author/));
    expect(addAuthorButton).toBeTruthy();
    await addAuthorButton!.trigger("click");
  }
}

describe("Book Form", () => {
  const wrapper = (book: Book) =>
    mount(BookForm, {
      props: { book },
    });

  it('displays "New book" when passing in with empty props', () => {
    const app = wrapper({ id: "", title: "", authors: [] });
    expect(app.text()).toContain("New book");
    expect(app.text()).toContain("Title");
  });

  it('displays "Edit book" when passing in an existing book', () => {
    const app = wrapper(mockBook());
    expect(app.text()).toContain("Edit book");
    expect(app.text()).toContain("Title");
  });

  it("can edit the title", () => {
    const app = wrapper({ id: "", title: "", authors: [] });
    const form = new FormObj(app);
    form.setInputValue("title", "new title here");
  });

  describe("authors", () => {
    it("displays no inputs when no authors are set", () => {
      const app = wrapper({ id: "", title: "", authors: [] });
      const form = new FormObj(app);
      expect(form.authorInputs).toHaveLength(0);
    });

    it("displays two inputs when two authors are set", () => {
      const app = wrapper({
        id: "",
        title: "",
        authors: ["First author", "Second author"],
      });
      const form = new FormObj(app);
      expect(form.authorInputs).toHaveLength(2);
    });

    it("displays an extra input when add author button is clicked", async () => {
      const app = wrapper({ id: "", title: "", authors: [] });
      const form = new FormObj(app);
      await form.clickAddAuthor();
      expect(form.authorInputs).toHaveLength(1);
    });
  });
});
