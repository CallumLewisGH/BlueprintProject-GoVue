<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center" role="presentation">
        <!-- Backdrop -->
        <div
          class="absolute inset-0 transition-opacity"
          :style="{ backgroundColor: 'rgba(0, 0, 0, 0.5)' }"
          @click="closeModal"
        />

        <!-- Modal Content -->
        <div
          ref="modalContent"
          class="relative z-10 w-full max-w-md mx-4 rounded-xl shadow-xl border max-h-[90vh] overflow-y-auto"
          :style="{
            backgroundColor: 'var(--card-bg)',
            borderColor: 'var(--border)'
          }"
          @click.stop
          role="dialog"
          :aria-modal="true"
          :aria-labelledby="title ? 'modal-title' : undefined"
        >
          <!-- Header -->
          <div v-if="title" class="flex items-center justify-between p-6 border-b" :style="{ borderColor: 'var(--border)' }">
            <h2 id="modal-title" class="text-lg font-bold" :style="{ color: 'var(--text-color)' }">{{ title }}</h2>
            <button
              @click="closeModal"
              class="p-1 rounded hover:opacity-70 transition-opacity"
              :style="{ color: 'var(--muted)' }"
              aria-label="Close modal"
            >
              <XMarkIcon class="w-5 h-5" />
            </button>
          </div>

          <!-- Body -->
          <div class="p-6">
            <slot />
          </div>

          <!-- Footer -->
          <div v-if="$slots.footer" class="border-t p-6 flex gap-3 justify-end" :style="{ borderColor: 'var(--border)' }">
            <slot name="footer" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, nextTick } from 'vue'
import { XMarkIcon } from '@heroicons/vue/24/outline'

interface Props {
  isOpen: boolean
  title?: string
}

const props = withDefaults(defineProps<Props>(), {
  title: ''
})

const emit = defineEmits<{
  'close': []
}>()

const modalContent = ref<HTMLElement>()
const previousActiveElement = ref<HTMLElement | null>(null)

function closeModal() {
  emit('close')
}

function handleKeydown(event: KeyboardEvent) {
  if (event.key === 'Escape') {
    closeModal()
  }
  
  // Focus trap: keep focus within modal
  if (event.key === 'Tab' && modalContent.value) {
    const focusableElements = modalContent.value.querySelectorAll(
      'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
    )
    const firstElement = focusableElements[0] as HTMLElement
    const lastElement = focusableElements[focusableElements.length - 1] as HTMLElement

    if (!firstElement || !lastElement) return

    if (event.shiftKey) {
      // Shift + Tab
      if (document.activeElement === firstElement) {
        lastElement.focus()
        event.preventDefault()
      }
    } else {
      // Tab
      if (document.activeElement === lastElement) {
        firstElement.focus()
        event.preventDefault()
      }
    }
  }
}

async function setupFocusTrap() {
  if (!modalContent.value) return
  
  await nextTick()
  
  // Store the previously focused element
  previousActiveElement.value = document.activeElement as HTMLElement
  
  // Find the first focusable element
  const focusableElements = modalContent.value.querySelectorAll(
    'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
  )
  
  if (focusableElements.length > 0) {
    (focusableElements[0] as HTMLElement).focus()
  } else {
    modalContent.value.focus()
  }
}

function restoreFocus() {
  if (previousActiveElement.value) {
    previousActiveElement.value.focus()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
  document.body.style.overflow = 'auto'
  restoreFocus()
})

watch(() => props.isOpen, async (newVal) => {
  if (newVal) {
    document.body.style.overflow = 'hidden'
    await setupFocusTrap()
  } else {
    document.body.style.overflow = 'auto'
    restoreFocus()
  }
})
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
