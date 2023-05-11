<script setup lang="ts">
import NavBar from "./components/NavBar.vue";
import { checkLoginQuery, getBooksQuery } from "./api/api.ts";

const {
  isLoading: booksIsLoading,
  isError: booksIsError,
  error: booksError,
} = getBooksQuery();

const {
  data: isLoggedIn,
  isLoading: authIsLoading,
  isError: authIsError,
  error: authError,
} = checkLoginQuery();
</script>

<template>
  <template v-if="booksIsLoading || authIsLoading">Loading...</template>
  <template v-if="booksIsError">Error: {{ booksError.message }}</template>
  <template v-else-if="authIsError">Error: {{ authError.message }}</template>
  <template v-else>
    <NavBar :isLoggedIn="isLoggedIn" />
    <div id="body">
      <router-view />
    </div>
  </template>
</template>

<style scoped>
#body {
  @apply w-full max-w-screen-xl mx-auto px-4 pb-8;
}
</style>
