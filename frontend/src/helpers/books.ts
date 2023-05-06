import { Book } from "../types/book.ts";

export const getBookPriceLocations = (books: Book[]) => {
  let locationsMap: { [location: string]: boolean } = {};
  books.forEach((book) =>
    Object.keys(book.prices ?? []).forEach(
      (location) => (locationsMap[location] = true)
    )
  );

  return Object.keys(locationsMap).sort();
};

const moneyFormatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
});

export const formatMoney = (amount?: number) => {
  if (!amount) return "-";
  return moneyFormatter.format(amount / 100);
};

export const formatOptionalBoolean = (value?: boolean) => {
  if (value == null) return "-";
  return value ? "Yes" : "No";
};
