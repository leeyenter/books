<script setup lang="ts">
import NavBar from "./components/NavBar.vue";
import { getBooksQuery } from "./api/api.ts";
const { isLoading, isError, error } = getBooksQuery();
</script>

<template>
  <template v-if="isLoading">Loading...</template>
  <template v-if="isError">Error: {{ error.message }}</template>
  <template v-else>
    <NavBar />
    <router-view />
  </template>
</template>
