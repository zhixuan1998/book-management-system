<template>
  <header class="header-container">
    <h1>{{ $messages.book.header() }}</h1>
  </header>
</template>

<script setup>
import { inject } from 'vue'
const $messages = inject('messages')
</script>

<style lang="scss">
@use '../assets/styles/functions.scss' as func;

.header-container {
  display: flex;
  width: 100%;
  justify-content: center;
  align-items: center;
  background-color: #{func.theme-color(m)};
}
</style>
