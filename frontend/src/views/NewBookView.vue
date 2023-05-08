<script setup lang="ts">
import BookForm from "../components/BookForm.vue";
</script>

<template>
  <BookForm :book="{ id: '', title: '', authors: [''] }" :locations="[]" />
</template>

<style scoped></style>
