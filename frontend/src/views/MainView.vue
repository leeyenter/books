<script setup lang="ts">
import BooksTable from "../components/BooksTable.vue";
import { getBooksQuery } from "../api/api.ts";
const { data } = getBooksQuery();
</script>

<template>
  <BooksTable v-if="data" :books="data" />
</template>

<style scoped></style>
