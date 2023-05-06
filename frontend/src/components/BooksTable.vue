<script setup lang="ts">
import { Book } from "../types/book.ts";

const props = defineProps<{
  books: Book[];
}>();
</script>

<template>
  <ul>
    <li v-for="book in props.books" :key="book.id">{{ book.title }}</li>
  </ul>
</template>
