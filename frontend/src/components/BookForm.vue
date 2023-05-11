<script setup lang="ts">
import { Book } from "../types/book.ts";
import { computed, ref } from "vue";
import { addBookMutation } from "../api/api.ts";
import { useQueryClient } from "@tanstack/vue-query";

const props = defineProps<{
  book: Book;
  locations: string[];
}>();

const generatePrices = () => {
  let pricesObj: { [location: string]: string } = Object.fromEntries(
    props.locations.map((location) => [location, ""])
  );
  if (props.book.prices) {
    for (const [loc, price] of Object.entries(props.book.prices)) {
      pricesObj[loc] = (price / 100).toFixed(2);
    }
  }
  return pricesObj;
};

const prices = ref(generatePrices());
const newPriceLocation = ref("");

const pricesInCents = computed(() => {
  let output: { [l: string]: number } = {};
  for (const [loc, priceStr] of Object.entries(prices.value)) {
    let priceFloat = parseFloat(priceStr);
    if (priceFloat > 0) {
      output[loc] = Math.round(priceFloat * 100);
    }
  }
  return output;
});

const book = ref(props.book);

const addNewPriceLocation = () => {
  if (newPriceLocation.value === "") return;
  if (newPriceLocation.value in prices.value) return;

  prices.value[newPriceLocation.value] = "";
  newPriceLocation.value = "";
};

const addAuthor = () => {
  if (!book.value.authors) {
    book.value.authors = [""];
  } else {
    book.value.authors.push("");
  }
};

const queryClient = useQueryClient();
const { mutate } = addBookMutation(queryClient);

const submitForm = () => {
  if (book.value.title === "") return;
  if (!book.value.authors) return;

  book.value.authors = book.value.authors.filter((x) => x.length > 0);
  if (book.value.authors.length === 0) return;

  book.value.prices = pricesInCents.value;
  mutate(book.value);
};
</script>

<template>
  <h1 v-if="book.id === ''">New book</h1>
  <h1 v-else>Edit book</h1>

  <form @submit.prevent="submitForm">
    <h2>Book details</h2>
    <div class="form-row">
      <div class="w-2/12">Title:</div>
      <input
        type="text"
        v-model="book.title"
        aria-label="Book title"
        class="flex-1"
      />
    </div>
    <form @submit.prevent="addAuthor">
      <div class="form-row">
        <div class="w-2/12">Author(s):</div>
        <div class="flex-1 flex flex-col">
          <input
            type="text"
            v-if="book.authors"
            v-for="index in book.authors.length"
            :key="index"
            v-model="book.authors[index - 1]"
            :aria-label="`Author ${index}`"
            data-test="author"
          />
          <button @click.prevent="addAuthor" class="secondary-btn mt-2">
            Add author
          </button>
        </div>
      </div>
    </form>
    <h2>Availability</h2>
    <div
      v-for="(_, location) in prices"
      class="form-row"
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
    <form @submit.prevent="addNewPriceLocation">
      <div class="form-row items-center">
        <div class="w-2/12">
          <input
            type="text"
            aria-label="New price location"
            v-model="newPriceLocation"
          />
        </div>
        <button @click.prevent="addNewPriceLocation" class="secondary-btn">
          Add location
        </button>
      </div>
    </form>
    <div class="form-row">
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
    <div class="form-row">
      <div class="w-2/12">Read Status:</div>
      <div>
        <select aria-label="Read status" v-model="book.readStatus">
          <option>Not started</option>
          <option>In progress</option>
          <option>Completed</option>
        </select>
      </div>
    </div>
    <button @click.prevent="submitForm" class="primary-btn">Add book</button>
  </form>
</template>

<style scoped>
.form-row {
  @apply flex flex-row mb-2;
}
</style>
