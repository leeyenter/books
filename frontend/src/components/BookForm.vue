<script setup lang="ts">
import { Book } from "../types/book.ts";
import { ref } from "vue";

const props = defineProps<{
  book: Book;
  locations: string[];
}>();

const generatePrices = () => {
  let pricesObj: { [location: string]: string } = Object.fromEntries(
    props.locations.map((location) => [location, ""])
  );
  if (book.value.prices) {
    for (const [loc, price] of Object.entries(book.value.prices)) {
      pricesObj[loc] = (price / 100).toFixed(2);
    }
  }
  return pricesObj;
};

const book = ref(props.book);
const prices = ref(generatePrices());
const newPriceLocation = ref("");

const addNewPriceLocation = () => {
  if (newPriceLocation.value === "") return;
  if (newPriceLocation.value in prices.value) return;

  prices.value[newPriceLocation.value] = "";
  newPriceLocation.value = "";
};

const addAuthor = () => {
  book.value.authors.push("");
};

const submitForm = () => {};
</script>

<template>
  <h1 v-if="book.id === ''">New book</h1>
  <h1 v-else>Edit book</h1>

  <form>
    <h2>Book details</h2>
    <div class="flex flex-row">
      <div class="w-2/12">Title:</div>
      <input
        type="text"
        v-model="book.title"
        aria-label="Book title"
        class="flex-1"
      />
    </div>
    <div class="flex flex-row">
      <div class="w-2/12">Author(s):</div>
      <div class="flex-1 flex flex-col">
        <input
          type="text"
          v-for="index in book.authors.length"
          :key="index"
          v-model="book.authors[index]"
          :aria-label="`Author ${index}`"
          data-test="author"
        />
        <button @click.prevent="addAuthor">Add author</button>
      </div>
    </div>
    <h2>Availability</h2>
    <div
      v-for="(_, location) in prices"
      class="flex flex-row"
      :aria-label="`price for ${location}`"
    >
      <div class="w-2/12">{{ location }}</div>
      <div>
        <input
          type="text"
          v-model="prices[location]"
          :aria-label="`Price at ${location}`"
          data-test="price"
        />
      </div>
    </div>
    <div>
      <input
        type="text"
        aria-label="New price location"
        v-model="newPriceLocation"
      />
      <button @click.prevent="addNewPriceLocation">Add price</button>
    </div>
    <div class="flex flex-row">
      <div class="w-2/12">FES Library:</div>
      <div>
        <select aria-label="FES library" v-model="book.fesLibrary">
          <option :value="undefined">Not checked</option>
          <option :value="true">Present</option>
          <option :value="false">Absent</option>
        </select>
      </div>
    </div>
    <h2>Other Details</h2>
    <div class="flex flex-row">
      <div class="w-2/12">Read Status:</div>
      <div>
        <select aria-label="Read status" v-model="book.readStatus">
          <option>Not started</option>
          <option>In progress</option>
          <option>Completed</option>
        </select>
      </div>
    </div>
    <button @click.prevent="submitForm">Add book</button>
  </form>
</template>
