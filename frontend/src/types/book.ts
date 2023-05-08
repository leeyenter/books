import { createPartialMock } from "./createPartialMock.ts";
import { faker } from "@faker-js/faker";

export interface Book {
  id: string;
  title: string;
  authors: string[];
  prices?: { [source: string]: number };
  fesLibrary?: boolean;
  boughtDate?: string;
  boughtType?: string;
  readStatus?: string;
}

export const mockBook = (overrides: Partial<Book> = {}) => {
  return createPartialMock(
    (): Book => ({
      id: faker.datatype.uuid(),
      title: faker.lorem.sentence(),
      authors: [faker.name.fullName(), faker.name.fullName()],
      prices: {
        Kindle: faker.datatype.number({ min: 1_00, max: 100_00 }),
        Logos: faker.datatype.number({ min: 1_00, max: 100_00 }),
      },
      readStatus: faker.helpers.arrayElement([
        "Not started",
        "In progress",
        "Completed",
      ]),
    }),
    overrides
  );
};
