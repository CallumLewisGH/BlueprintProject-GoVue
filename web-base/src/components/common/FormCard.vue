<template>
  <div class="p-6 rounded-xl border" 
       :style="{ backgroundColor: 'var(--card-bg)', borderColor: 'var(--border)' }">
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <slot />
      <div class="flex gap-3 pt-4">
        <Button
          :label="submitLabel"
          type="submit"
          variant="primary"
          :disabled="isLoading"
          :loading-text="loadingLabel"
        />
        <Button
          v-if="showCancel"
          :label="cancelLabel"
          type="button"
          variant="outline"
          @click="$emit('cancel')"
        />
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import Button from '@/components/ui/Button.vue'

interface Props {
  isLoading?: boolean
  submitLabel?: string
  cancelLabel?: string
  loadingLabel?: string
  showCancel?: boolean
}

withDefaults(defineProps<Props>(), {
  isLoading: false,
  submitLabel: 'Submit',
  cancelLabel: 'Cancel',
  loadingLabel: 'Submitting...',
  showCancel: true
})

const emit = defineEmits<{
  'submit': []
  'cancel': []
}>()

function handleSubmit() {
  emit('submit')
}
</script>
