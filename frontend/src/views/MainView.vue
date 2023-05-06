<script setup lang="ts">
import { getBooksQuery } from "../api/api.ts";
import BooksTable from "../components/BooksTable.vue";

const { isLoading, isError, data, error } = getBooksQuery();
</script>

<template>
  <div v-if="isLoading">Loading...</div>
  <div v-else-if="isError">Error: {{ error.message }}</div>
  <div v-else>
    <BooksTable :books="data" />
  </div>
</template>

<style scoped></style>
