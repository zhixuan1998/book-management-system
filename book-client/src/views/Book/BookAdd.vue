<template>
  <book-detail-form v-model="book" :isEditMode="true">
    <template #bottom-button>
      <div class="book button-container justify-flex-end">
        <custom-button @click="handleAddBook">{{ $messages.button.confirm() }}</custom-button>
      </div>
    </template>
  </book-detail-form>
</template>

<script setup>
import { ref, inject } from 'vue';
import { useRouter } from 'vue-router';
import BookDetailForm from './components/BookDetailForm.vue';
import { getDefaultData, generateRequestData } from './book';

const router = useRouter();
const $modal = inject('modal');
const $messages = inject('messages');
const $repositories = inject('repositories');

const book = ref(getDefaultData());

async function handleAddBook() {
  if (!isDataValid()) {
    return $modal.open({
      message: $messages.error.allFieldRequired()
    })
  }

  await addBook();
}

function isDataValid() {
  if (
    !book.value.title ||
    !book.value.author ||
    !book.value.genre ||
    !book.value.description ||
    !book.value.isbn ||
    !book.value.publisher ||
    isNaN(new Date(book.value.publishedAt).getTime())
  ) {
    return false;
  }

  return true;
}

async function addBook() {
  const bookData = generateRequestData(book.value);

  await $repositories.bookRepository.add(bookData).then(({ data }) => {
    router.push({ name: 'BookDetail', params: { bookId: data.data.bookId } });
  });
}
</script>
