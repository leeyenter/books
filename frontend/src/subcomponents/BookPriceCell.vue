<script setup lang="ts">
import { Book } from "../types/book.ts";
import { formatMoney } from "../helpers/books.ts";
import { computed } from "vue";

const props = defineProps<{
  prices?: { [p: string]: number };
  loc: string;
}>();

const className = computed(() => {
  if (money.value === "-") return "text-center text-gray-400";
  return "text-right";
});

const money = computed(() => {
  if (!props.prices) return "-";
  return formatMoney(props.prices[props.loc]);
});
</script>

<template>
  <td :class="className">
    {{ money }}
  </td>
</template>

<style scoped></style>
