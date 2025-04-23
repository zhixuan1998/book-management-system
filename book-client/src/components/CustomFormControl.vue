<template>
  <div class="detail-row">
    <label v-if="label">{{ label }}</label>
    <input v-model="formattedModelValue" class="input" :type="validatedType" :readonly="readonly" />
  </div>
  <custom-separator />
</template>

<script setup>
import { computed } from 'vue';
import { convertDateToDatePickerFormat } from '@/utils/dateFormatter';

const defaultType = 'text';
const validTypes = [defaultType, 'number', 'date'];

const modelValue = defineModel();

const props = defineProps({
  label: {
    type: String,
    default: ''
  },
  type: {
    type: String,
    default: defaultType
  },
  readonly: {
    type: Boolean,
    default: true
  }
});

const validatedType = computed(() => (validTypes.includes(props.type) ? props.type : defaultType));

const formattedModelValue = computed({
  get() {
    return props.type === 'date'
      ? convertDateToDatePickerFormat(modelValue.value)
      : modelValue.value;
  },
  set(newValue) {
    modelValue.value = newValue;
  }
});
</script>

<style lang="scss">
@use '../assets/styles/functions.scss' as func;

.detail-row {
  --spacing: 20px;

  display: flex;
  flex-direction: column;
  padding: var(--spacing);
  gap: var(--spacing);

  label {
    font-weight: 600;
  }

  .input {
    border: none;
    padding: var(--spacing);
    transition: 0.15s;
    outline: 1px solid rgba(0, 0, 0, 0.15);
    position: relative;

    &:focus {
      outline-color: rgba(0, 0, 0, 0.5);
    }

    &:read-only {
      cursor: auto;
      outline-color: transparent;
    }

    &::-webkit-calendar-picker-indicator {
      opacity: 0;
      position: absolute;
      inset: 0;
      width: 100%;
      height: 100%;
    }
  }
}
</style>
