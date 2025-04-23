<template>
  <div class="book button-container justify-space-between">
    <div class="book button-container">
      <custom-button @click="goToListingPage">
        <font-awesome-icon :icon="faAngleLeft" />
        {{ $messages.button.back() }}
      </custom-button>
    </div>
    <slot name="top-button"></slot>
  </div>

  <custom-form :headerText="bookInitialTitle">
    <custom-form-control v-if="!isAddPage" v-model="book.bookId" :label="$messages.book.bookId()" />
    <custom-form-control
      v-model="book.title"
      :label="$messages.book.title()"
      :readonly="!isEditMode"
    />
    <custom-form-control
      v-model="book.author"
      :label="$messages.book.author()"
      :readonly="!isEditMode"
    />
    <custom-form-control
      v-model="book.genre"
      :label="$messages.book.genre()"
      :readonly="!isEditMode"
    />
    <custom-form-control
      v-model="book.description"
      :label="$messages.book.description()"
      :readonly="!isEditMode"
    />
    <custom-form-control
      v-model="book.isbn"
      :label="$messages.book.isbn()"
      :readonly="!isEditMode"
    />
    <custom-form-control
      v-model="book.publisher"
      :label="$messages.book.publisher()"
      :readonly="!isEditMode"
    />
    <custom-form-control
      v-model="book.publishedAt"
      :label="$messages.book.publishedAt()"
      type="date"
      :readonly="!isEditMode"
    />
  </custom-form>
  <slot name="bottom-button"></slot>
</template>

<script setup>
import { ref, inject, computed, watch } from 'vue';
import { useRouter } from 'vue-router';
import { faAngleLeft } from '@fortawesome/free-solid-svg-icons';

const book = defineModel();

const props = defineProps({
  isEditMode: Boolean
});

const router = useRouter();
const $messages = inject('messages');

const isAddPage = computed(() => !book.value.bookId);
const bookInitialTitle = ref(null);

const unwatchTitle = watch(
  () => book?.value.title,
  (newTitle) => {
    if (isAddPage) {
      unwatchTitle();
    } else if (newTitle) {
      bookInitialTitle.value = newTitle;
      unwatchTitle();
    }
  }
);

function goToListingPage() {
  router.push({ name: 'BookListing' });
}
</script>

<style>
.book.button-container {
  display: flex;
  padding: 0;
  gap: 20px;
}
.justify-flex-end {
  justify-content: flex-end;
}
.justify-space-between {
  justify-content: space-between;
}
</style>
