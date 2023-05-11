<script setup lang="ts">
const props = defineProps<{
  isLoggedIn?: boolean;
}>();
</script>

<template>
  <nav class="p-4 flex flex-row">
    <router-link to="/" class="font-extrabold text-lg">Books</router-link>
    <div class="flex-1" />
    <router-link v-if="isLoggedIn" to="/new" class="primary-btn">
      New
    </router-link>
    <router-link v-else to="/login" class="primary-btn">Login</router-link>
  </nav>
</template>
