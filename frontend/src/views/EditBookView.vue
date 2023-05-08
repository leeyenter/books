<script setup lang="ts">
import BookForm from "../components/BookForm.vue";
import { getBooksQuery } from "../api/api.ts";
import { getBookPriceLocations } from "../helpers/books.ts";
const { data } = getBooksQuery();
</script>

<template>
  <BookForm
    v-if="data && data.length > 0"
    :book="data.find((x) => x.id === $route.params.id)"
    :locations="getBookPriceLocations(data)"
  />
</template>

<style scoped></style>
