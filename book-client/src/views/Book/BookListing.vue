<template>
  <custom-separator />
  <div class="pagination-wrapper">
    <div class="pagination-indicator">
      <custom-button @click="previousPage" :disabled="currentPage == 1">
        <font-awesome-icon :icon="faAngleLeft" />
      </custom-button>
      {{ pagination?.itemFrom }} - {{ pagination?.itemTo }}
      <custom-button @click="nextPage" :disabled="pagination?.isLastPage">
        <font-awesome-icon :icon="faAngleRight" />
      </custom-button>
      <custom-button @click="goToAddPage">{{ $messages.button.addNew() }}</custom-button>
      <custom-button @click="resetConfirmation"
        ><font-awesome-icon :icon="faTriangleExclamation" class="warning-icon" />{{
          $messages.button.reset()
        }}</custom-button
      >
    </div>
    <div class="pagination-dropdown">
      <span class="fw-bold">Items per page</span>
      <custom-dropdown v-model="perPage" :options="perPageOptions" />
    </div>
  </div>
  <custom-separator />
  <div class="main-content">
    <table>
      <tbody>
        <tr>
          <th>{{ $messages.book.bookId() }}</th>
          <th>{{ $messages.book.title() }}</th>
          <th>{{ $messages.book.author() }}</th>
          <th>{{ $messages.book.genre() }}</th>
          <th>{{ $messages.book.publisher() }}</th>
          <th>{{ $messages.book.publishedAt() }}</th>
          <th>{{ $messages.general.actions() }}</th>
        </tr>
        <book-listing-row v-for="item of books" :key="item.bookId" :item="item" />
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, inject } from 'vue';
import { useRouter } from 'vue-router';
import BookListingRow from './components/BookListingRow.vue';
import {
  faAngleLeft,
  faAngleRight,
  faTriangleExclamation
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const $modal = inject('modal');
const $messages = inject('messages');
const $repositories = inject('repositories');

const perPageOptions = ref([]);
const books = ref([]);
const pagination = ref(null);
let currentPage = 1;
const perPage = ref(10);

onMounted(async () => {
  const perPageOptionsResult = $repositories.lookupRepository.getRecordPerPageOptions();
  perPageOptions.value = perPageOptionsResult;

  getBooks();
});

async function getBooks() {
  await $repositories.bookRepository
    .getAll({ limit: perPage.value, page: currentPage })
    .then(({ data }) => {
      books.value = data?.data?.results ?? [];
      pagination.value = data?.data?.pagination;
    });
}

watch(perPage, async () => {
  currentPage = 1;
  await getBooks();
});

function resetConfirmation() {
  $modal.open({
    title: $messages.error.beCareful(),
    message: $messages.book.resetWarningMessage(),
    cancelButtonText: $messages.button.cancel(),
    onConfirm: resetBookList
  });
}

async function resetBookList() {
  await $repositories.bookRepository
    .reset()
    .then(async () => {
      await getBooks();
    })
    .catch(() => $modal.open());
}

function goToAddPage() {
  router.push({ name: 'BookAdd' });
}

async function previousPage() {
  currentPage -= 1;
  await getBooks();
}

async function nextPage() {
  currentPage += 1;
  await getBooks();
}
</script>

<style lang="scss">
@use '../../assets/styles/functions.scss' as func;

.pagination-wrapper {
  padding: 10px 0;
  justify-content: space-between;

  &,
  .pagination-dropdown,
  .pagination-indicator {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .pagination-dropdown,
  .pagination-indicator {
    gap: 10px;
    justify-content: flex-end;
  }

  .custom-dropdown {
    height: 35px;

    .dropdown-container {
      border: 1px solid func.theme-color(l);
      background-color: #fff;
    }
  }
}

.warning-icon {
  color: #cca333;
}
</style>
