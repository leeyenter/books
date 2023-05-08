<script setup lang="ts">
import { Book } from "../types/book.ts";
import {
  formatMoney,
  formatOptionalBoolean,
  getBookPriceLocations,
} from "../helpers/books.ts";

const props = defineProps<{
  books: Book[];
}>();

const locations = getBookPriceLocations(props.books);
</script>

<template>
  <div v-if="props.books.length === 0">No books found</div>
  <table v-else>
    <thead>
      <tr>
        <th>Title</th>
        <th>Author(s)</th>
        <th v-for="loc in locations">{{ loc }}</th>
        <th>FES Library</th>
        <th>Owned</th>
        <th>Status</th>
        <th></th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="book in props.books" :key="book.id">
        <td aria-label="title">{{ book.title }}</td>
        <td aria-label="authors">{{ book.authors.join(", ") }}</td>
        <td v-for="loc in locations" data-test="price">
          <template v-if="book.prices">
            {{ formatMoney(book.prices[loc]) }}
          </template>
          <template v-else>-</template>
        </td>
        <td aria-label="fes library">
          {{ formatOptionalBoolean(book.fesLibrary) }}
        </td>
        <td aria-label="owned">{{ book.boughtType ?? "-" }}</td>
        <td aria-label="status">{{ book.readStatus }}</td>
        <td>
          <router-link :to="'/edit/' + book.id" aria-label="Edit book">
            Edit
          </router-link>
          <button>Delete</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>
