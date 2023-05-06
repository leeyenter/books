import { DOMWrapper, mount, VueWrapper } from "@vue/test-utils";
import { Book, mockBook } from "../types/book.ts";
import BooksTable from "./BooksTable.vue";
import { expect } from "vitest";
import { faker } from "@faker-js/faker";

class BookPage {
  constructor(private wrapper: VueWrapper) {}

  row(i: number) {
    return new BookRow(this.wrapper.find("tbody").findAll("tr")[i]);
  }

  expectHeaders(colNames: string[]) {
    const headers = this.wrapper.find("thead").findAll("th");
    for (let i = 0; i < colNames.length; i++) {
      expect(headers[i].text()).toEqual(colNames[i]);
    }
    return this;
  }
}

class BookRow {
  constructor(private wrapper: DOMWrapper<HTMLElementTagNameMap["tr"]>) {}

  expectTitle(title: string) {
    expect(this.wrapper.find("[aria-label=title]").text()).toEqual(title);
    return this;
  }

  expectAuthors(authors: string) {
    expect(this.wrapper.find("[aria-label=authors]").text()).toEqual(authors);
    return this;
  }

  expectPrices(prices: string[]) {
    const rowPrices = this.wrapper.findAll('[data-test="price"]');
    expect(rowPrices.length).toEqual(prices.length);
    for (let i = 0; i < prices.length; i++) {
      expect(rowPrices[i].text()).toEqual(prices[i]);
    }
    return this;
  }

  expectFESLibrary(text: string) {
    expect(this.wrapper.find('[aria-label="fes library"]').text()).toEqual(
      text
    );
    return this;
  }

  expectOwned(owned: string) {
    expect(this.wrapper.find("[aria-label=owned]").text()).toEqual(owned);
    return this;
  }

  expectReadStatus(status: string) {
    expect(this.wrapper.find("[aria-label=status]").text()).toEqual(status);
    return this;
  }
}

describe("Books Table component", async () => {
  const wrapper = (books: Book[]) =>
    mount(BooksTable, {
      props: { books },
    });

  it("has an empty state", () => {
    const app = wrapper([]);
    expect(app.text()).toContain("No books found");
  });

  it("renders books correctly", () => {
    const app = wrapper([
      mockBook({
        title: "A long obedience",
        authors: ["Eugene Peterson"],
        boughtDate: faker.date.past().toDateString(),
        boughtType: "Kindle",
        readStatus: "Not started",
        prices: {
          Logos: 20_00,
          Kindle: 18_50,
        },
        fesLibrary: true,
      }),
      mockBook({
        title: "Another book",
        authors: ["Author 1", "Author 2"],
        readStatus: "In progress",
        prices: {
          SKS: 20_10,
          Kindle: 15_00,
        },
        fesLibrary: false,
      }),
      mockBook(),
    ]);

    const pageObj = new BookPage(app);

    pageObj.expectHeaders([
      "Title",
      "Author(s)",
      "Kindle",
      "Logos",
      "SKS",
      "FES Library",
      "Owned",
      "Status",
    ]);

    pageObj
      .row(0)
      .expectTitle("A long obedience")
      .expectAuthors("Eugene Peterson")
      .expectPrices(["$18.50", "$20.00", "-"])
      .expectFESLibrary("Yes")
      .expectOwned("Kindle")
      .expectReadStatus("Not started");

    pageObj
      .row(1)
      .expectTitle("Another book")
      .expectAuthors("Author 1, Author 2")
      .expectPrices(["$15.00", "-", "$20.10"])
      .expectFESLibrary("No")
      .expectOwned("-")
      .expectReadStatus("In progress");

    pageObj.row(2).expectFESLibrary("-");
  });
});
