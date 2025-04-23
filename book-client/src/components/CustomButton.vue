<template>
  <button class="button" :class="type" :disabled="disabled"><slot></slot></button>
</template>

<script setup>
import { inject } from 'vue';
const $messages = inject('messages');
const props = defineProps({
  type: {
    type: String,
    default: 'primary'
  },
  disabled: Boolean
});
</script>

<style lang="scss" scoped>
@use '../assets/styles/functions.scss' as func;

.button {
  display: flex;
  gap: 10px;
  padding: 10px 20px;
  background-color: func.theme-color(xs);
  border: 1px solid func.theme-color(s);
  border-radius: 5px;
  cursor: pointer;
  transition: 0.15s;

  //   &.primary {
  //     color: func.theme-color(xs);
  //     background-color: func.theme-color(xs);
  //     border: 1px solid func.theme-color(s);

  //     &:hover {
  //       background-color: func.theme-color(s, 0.5);
  //     }
  //   }

  //   &.secondary {
  //     background-color: func.theme-color(xs);
  //     border: 1px solid func.theme-color(s);

  //     &:hover {
  //       background-color: func.theme-color(s, 0.5);
  //     }
  //   }

  &:hover {
    background-color: func.theme-color(s, 0.5);
  }

  &:active {
    transform: scale(0.95);
  }
}
</style>
