<script setup lang="ts">
import { tryLogin } from "../helpers/auth.ts";
import { useRouter } from "vue-router";
import { useQueryClient } from "@tanstack/vue-query";

tryLogin(useRouter(), useQueryClient());
</script>

<template></template>

<style scoped></style>
