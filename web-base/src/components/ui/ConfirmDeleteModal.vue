<template>
  <Modal :is-open="isOpen" :title="title" @close="handleCancel">
    <p class="text-sm" :style="{ color: 'var(--text-color)' }">
      {{ message }}
    </p>
    
    <template #footer>
      <Button
        label="Cancel"
        variant="outline"
        @click="handleCancel"
      />
      <Button
        label="Delete"
        variant="danger"
        :disabled="isDeleting"
        :loading-text="'Deleting...'"
        @click="handleConfirm"
      />
    </template>
  </Modal>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Modal from './Modal.vue'
import Button from './Button.vue'

interface Props {
  isOpen: boolean
  title?: string
  message?: string
  isDeleting?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Confirm Deletion',
  message: 'Are you sure you want to delete this? This action cannot be undone.',
  isDeleting: false
})

const emit = defineEmits<{
  'confirm': []
  'cancel': []
  'close': []
}>()

const isDeleting = ref(false)

function handleConfirm() {
  emit('confirm')
}

function handleCancel() {
  isDeleting.value = false
  emit('cancel')
  emit('close')
}
</script>
