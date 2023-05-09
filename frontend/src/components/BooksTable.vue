<script setup lang="ts">
import { Book } from "../types/book.ts";
import {
  formatOptionalBoolean,
  getBookPriceLocations,
} from "../helpers/books.ts";
import BookPriceCell from "../subcomponents/BookPriceCell.vue";

const props = defineProps<{
  books: Book[];
}>();

const locations = getBookPriceLocations(props.books);
</script>

<template>
  <div v-if="props.books.length === 0">No books found</div>
  <table v-else class="w-full">
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
        <td aria-label="authors">{{ book.authors?.join(", ") }}</td>
        <template v-for="loc in locations" data-test="price">
          <BookPriceCell :prices="book.prices" :loc="loc" />
        </template>
        <td aria-label="fes library" class="text-center">
          {{ formatOptionalBoolean(book.fesLibrary) }}
        </td>
        <td aria-label="owned">{{ book.boughtType ?? "-" }}</td>
        <td aria-label="status">{{ book.readStatus }}</td>
        <td>
          <router-link
            :to="'/edit/' + book.id"
            aria-label="Edit book"
            class="secondary-btn"
          >
            Edit
          </router-link>
          <button class="secondary-btn">Delete</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<style scoped>
td,
th {
  @apply py-1 px-3;
}

th {
  @apply border-b-2;
}

td {
  @apply border-b;
}
</style>
