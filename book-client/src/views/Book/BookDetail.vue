<template>
  <book-detail-form v-model="book" :isEditMode="isEditMode">
    <template #top-button>
      <div class="book button-container justify-flex-end" v-if="!isEditMode">
        <custom-button @click="goToAddPage">{{ $messages.button.addNew() }}</custom-button>
        <custom-button @click="enableEditMode">{{ $messages.button.edit() }}</custom-button>
      </div>
    </template>

    <template #bottom-button>
      <div class="book button-container" v-if="!isEditMode">
        <custom-button @click="deleteConfirmation">{{ $messages.button.delete() }}</custom-button>
      </div>

      <div class="book button-container justify-flex-end" v-else>
        <custom-button @click="disableEditMode">{{ $messages.button.cancel() }}</custom-button>
        <custom-button @click="updateBook">{{ $messages.button.save() }}</custom-button>
      </div>
    </template>
  </book-detail-form>
</template>

<script setup>
import { ref, inject, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import BookDetailForm from './components/BookDetailForm.vue';
import { generateRequestData } from './book';

const route = useRoute();
const router = useRouter();

const $modal = inject('modal');
const $messages = inject('messages');
const $repositories = inject('repositories');

const book = ref({});
const isEditMode = ref(false);

onMounted(async () => {
  await getBook(route.params.bookId);
});

async function getBook(bookId) {
  await $repositories.bookRepository
    .get(bookId)
    .then(({ data }) => {
      book.value = data?.data ?? {};
    })
    .catch(() => $modal.open());
}

async function updateBook() {
  const bookData = generateRequestData(book.value);
  await $repositories.bookRepository
    .update(bookData)
    .then(() => {
      $modal.open({
        title: $messages.general.successful(),
        message: $messages.general.updateSuccessful()
      });
    })
    .catch(() => {
      $modal.open();
    })
    .finally(() => {
      disableEditMode();
    });
}

function deleteConfirmation() {
  $modal.open({
    title: $messages.general.deleteConfirmation(),
    message: $messages.book.deleteMessage(bookId.value),
    cancelButtonText: $messages.button.cancel(),
    onConfirm: deleteBook
  });
}

async function deleteBook() {
  await $repositories.bookRepository
    .delete(book.value.bookId)
    .then(() => router.push({ name: 'BookListing' }))
    .catch(() => $modal.open());
}

function enableEditMode() {
  isEditMode.value = true;
}

function disableEditMode() {
  isEditMode.value = false;
  getBook(book.value.bookId);
}

function goToAddPage() {
  router.push({ name: 'BookAdd' });
}
</script>
