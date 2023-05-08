import { Book, mockBook } from "../types/book.ts";
import BookForm from "./BookForm.vue";
import { DOMWrapper, mount, VueWrapper } from "@vue/test-utils";
import { beforeEach, expect, Mock } from "vitest";
import { VueQueryPlugin } from "@tanstack/vue-query";

export class FormObj {
  constructor(private app: VueWrapper) {}

  private input(label: string): DOMWrapper<HTMLInputElement> {
    const input = this.app.find(`[aria-label="${label}"]`);
    expect(
      input.exists(),
      `Input element with ref ${label} not found`
    ).toBeTruthy();
    return input as DOMWrapper<HTMLInputElement>;
  }

  private button(text: RegExp) {
    const button = this.app.findAll("button").find((x) => x.text().match(text));
    expect(button?.exists(), `Expected button with text ${text}`).toBeTruthy();
    return button!;
  }

  private get fesLibrarySelect() {
    return this.app.find(
      '[aria-label="FES library"]'
    ) as DOMWrapper<HTMLSelectElement>;
  }

  private get readStatusSelect() {
    return this.app.find('[aria-label="Read status"]');
  }

  get authorInputs() {
    return this.app.findAll('[data-test="author"]');
  }

  get priceInputs() {
    return this.app.findAll('[data-test="price"]');
  }

  async setInputValue(label: string, value: string) {
    const input = this.input(label);
    return input.setValue(value);
  }

  async setTitleInput(value: string) {
    return this.setInputValue("Book title", value);
  }

  expectInputValue(label: string, value: string) {
    const input = this.input(label);
    expect(input.element.value).toEqual(value);
  }

  async clickAddAuthor() {
    const addAuthorButton = this.button(/Add author/);
    await addAuthorButton.trigger("click");
  }

  expectPrice(location: string, price: string) {
    const priceDiv = this.app.find(`[aria-label="price for ${location}"]`);
    expect(priceDiv.exists(), `Expected price for ${location}`).toBeTruthy();
    expect(priceDiv.text()).toContain(location);
    expect(priceDiv.find("input").element.value).toEqual(price);
  }

  async clickAddPrice() {
    const addPriceBtn = this.button(/Add location/);
    await addPriceBtn.trigger("click");
  }

  async selectFESLibraryValue(value: string) {
    const option = this.fesLibrarySelect
      .findAll("option")
      .find((x) => x.text() === value);
    expect(option?.exists()).toBeTruthy();
    await option!.setValue();
  }

  async selectReadStatusValue(value: string) {
    const option = this.readStatusSelect
      .findAll("option")
      .find((x) => x.text() === value);
    expect(option?.exists()).toBeTruthy();
    await option!.setValue();
  }

  expectFESLibraryValue(value: string) {
    expect(this.fesLibrarySelect.element.value).toEqual(value);
  }

  async clickSubmit() {
    const addBookBtn = this.button(/Add book/);
    await addBookBtn.trigger("click");
  }
}

describe("Book Form", () => {
  let fetch: Mock<any, any>;

  const wrapper = (book: Book, locations: string[] = []) =>
    mount(BookForm, {
      props: { book, locations },
      global: {
        plugins: [VueQueryPlugin],
      },
    });

  beforeEach(() => {
    fetch = vi.fn();
    global.fetch = fetch;
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

  it("can edit the title", async () => {
    const app = wrapper({ id: "", title: "", authors: [] });
    const form = new FormObj(app);
    await form.setTitleInput("new title here");
    form.expectInputValue("Book title", "new title here");
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

  describe("prices", () => {
    let app: VueWrapper;
    let form: FormObj;

    beforeEach(() => {
      app = wrapper(
        mockBook({
          prices: {
            Logos: 1234,
            Kindle: 850,
          },
        }),
        ["SKS", "BD"]
      );

      form = new FormObj(app);
    });

    it("displays existing prices", () => {
      form.expectPrice("Logos", "12.34");
      form.expectPrice("Kindle", "8.50");
    });

    it("can edit existing prices", async () => {
      await form.setInputValue("Price at Logos", "1.00");
      form.expectPrice("Logos", "1.00");
    });

    it("displays existing categories", () => {
      form.expectPrice("SKS", "");
      form.expectPrice("BD", "");
    });

    describe("adding a price for a new location", () => {
      it("can add a new location", async () => {
        await form.setInputValue("New price location", "My New Location");
        await form.clickAddPrice();
        expect(form.priceInputs).toHaveLength(5);
        form.expectInputValue("New price location", "");
        form.expectPrice("My New Location", "");
      });

      it("does not add a new price with an empty location", async () => {
        expect(form.priceInputs).toHaveLength(4);
        await form.clickAddPrice();
        expect(form.priceInputs).toHaveLength(4);
      });

      it("does not add an existing location", async () => {
        expect(form.priceInputs).toHaveLength(4);
        await form.setInputValue("New price location", "Logos");
        await form.clickAddPrice();
        expect(form.priceInputs).toHaveLength(4);
        form.expectPrice("Logos", "12.34");
      });
    });
  });

  describe("library selection", () => {
    it.each([
      { data: undefined, value: "Not checked" },
      { data: true, value: "true" },
      { data: false, value: "false" },
    ])(
      "correctly maps data to select option when initialising component %s",
      ({ data, value }) => {
        const app = wrapper({
          id: "",
          title: "",
          authors: [],
          fesLibrary: data,
        });
        const form = new FormObj(app);
        form.expectFESLibraryValue(value);
      }
    );
  });

  describe("submit", () => {
    it("can create new book", async () => {
      const app = wrapper({ id: "", title: "", authors: [""] });
      const form = new FormObj(app);
      await form.setTitleInput("Preaching");
      await form.setInputValue("Author 1", "Tim Keller");

      await form.setInputValue("New price location", "SKS");
      await form.clickAddPrice();
      await form.setInputValue("Price at SKS", "24.30");

      await form.selectFESLibraryValue("Present");
      await form.selectReadStatusValue("Completed");

      await form.clickSubmit();

      expect(fetch).toHaveBeenCalledWith("http://localhost:5000/book/", {
        method: "POST",
        body: '{"id":"","title":"Preaching","authors":["Tim Keller"],"fesLibrary":true,"readStatus":"Completed","prices":{"SKS":2430}}',
      });
    });
  });
});
