export interface Book {
  id: string;
  title: string;
  authors: string[];
  prices?: { [source: string]: number };
  boughtDate?: string;
  boughtType?: string;
  readStatus?: string;
}
