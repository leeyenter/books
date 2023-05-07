<script setup lang="ts">
import { Book } from "../types/book.ts";
import { ref } from "vue";
const props = defineProps<{
  book: Book;
}>();

const book = ref(props.book);

const addAuthor = () => {
  book.value.authors.push("");
};
</script>

<template>
  <h1 v-if="book.id === ''">New book</h1>
  <h1 v-else>Edit book</h1>

  <form>
    <div class="flex flex-row">
      <label for="title" class="w-2/12">Title:</label>
      <input
        type="text"
        v-model="book.title"
        id="title"
        ref="title"
        class="flex-1"
      />
    </div>
    <div class="flex flex-row">
      <label class="w-2/12">Author(s):</label>
      <div class="flex-1 flex flex-col">
        <input
          type="text"
          v-for="(author, index) in book.authors"
          :key="index"
          data-test="author"
          v-model="book.authors[index]"
        />
        <button @click.prevent="addAuthor">Add author</button>
      </div>
    </div>
  </form>
</template>
