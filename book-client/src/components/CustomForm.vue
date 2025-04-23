<template>
  <div class="main-content">
    <div class="content-header" v-if="headerText">{{ headerText }}</div>
    <slot></slot>
  </div>
</template>

<script setup>
const props = defineProps({
  headerText: {
    type: String,
    default: ''
  }
});
</script>

<style lang="scss">
@use '../assets/styles/functions.scss' as func;

.main-content {
  .content-header {
    font-size: 1.5rem;
    font-weight: 500;
    padding: var(--padding);
    background-color: func.theme-color(xs);
    border-radius: 10px 10px 0 0;
  }
}
</style>
