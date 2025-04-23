<template>
  <tr>
    <td>{{ item.bookId || '-' }}</td>
    <td>{{ item.title || '-' }}</td>
    <td>{{ item.author || '-' }}</td>
    <td>{{ item.genre || '-' }}</td>
    <td>{{ item.publisher || '-' }}</td>
    <td>{{ convertDateToFriendlyDisplay(item.publishedAt) || '-' }}</td>
    <td>
      <custom-button @click="goToDetail">{{ $messages.button.view() }}</custom-button>
    </td>
  </tr>
</template>

<script setup>
import { inject } from 'vue';
import { useRouter } from 'vue-router';
import { convertDateToFriendlyDisplay } from '@/utils/dateFormatter';

const props = defineProps({
  item: Object
});

const router = useRouter();
const $messages = inject('messages');

function goToDetail() {
  router.push({ name: 'BookDetail', params: { bookId: props.item.bookId } });
}
</script>
