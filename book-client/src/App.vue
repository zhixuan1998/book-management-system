<template>
  <div class="app-container">
    <custom-header />
    <div class="container">
      <router-view />
    </div>
    <div class="modal-overlay" v-show="showModal">
      <div class="modal-container">
        <div class="modal-title">{{ modal.title }}</div>
        <div class="modal-body">{{ modal.message }}</div>
        <div class="modal-button">
          <custom-button @click="hideModal" v-if="!!modal.cancelButtonText">{{
            modal.cancelButtonText
          }}</custom-button>
          <custom-button @click="hideModalOnConfirm()">{{ modal.confirmButtonText }}</custom-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router';
import { ref, reactive, provide, inject } from 'vue';

const $messages = inject('messages');

const showModal = ref(false);

const noop = () => {};

const modal = reactive({
  title: '',
  message: '',
  cancelButtonText: '',
  confirmButtonText: '',
  onConfirm: noop
});

function openModal(obj) {
  modal.title = obj?.title ?? $messages.error.oops();
  modal.message = obj?.message ?? $messages.error.general();
  modal.cancelButtonText = obj?.cancelButtonText ?? '';
  modal.confirmButtonText = obj?.confirmButtonText ?? $messages.button.ok();
  modal.onConfirm = obj?.onConfirm ?? noop;

  showModal.value = true;
}

function hideModal() {
  showModal.value = false;
  resetModal();
}

async function hideModalOnConfirm() {
  await modal.onConfirm();
  hideModal();
}

function resetModal() {
  modal.title = modal.message = modal.cancelButtonText = modal.confirmButtonText = '';
  modal.onConfirm = noop;
}

provide('modal', { open: openModal, hide: hideModal });
</script>

<style scoped>
.app-container {
  display: grid;
  width: 100%;
  min-height: 100vh;
  grid-template-rows: 100px 1fr;
  position: relative;

  .modal-overlay {
    position: fixed;
    display: flex;
    top: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgb(0, 0, 0, 0.5);
    z-index: 20;

    .modal-container {
      margin: auto;
      display: flex;
      flex-direction: column;
      width: min(100%, 500px);
      background-color: #ffffff;
      box-shadow: 0 3px 10px 0 rgba(0, 0, 0, 0.14);
      padding: 30px;
      row-gap: 20px;

      & > * {
        width: 100%;
      }

      .modal-title,
      .modal-body,
      .modal-button {
        text-align: center;
      }

      .modal-title {
        font-size: 20px;
        font-weight: 550;
      }

      .modal-body {
        font-size: 15px;
      }

      .modal-button {
        display: flex;
        gap: 20px;
        justify-content: center;
      }
    }
  }
}
</style>
