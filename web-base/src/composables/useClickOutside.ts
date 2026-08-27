import { ref, onMounted, onUnmounted } from 'vue'

export function useClickOutside(callback: () => void) {
  const elementRef = ref<HTMLElement | null>(null)

  function handleClickOutside(event: MouseEvent) {
    if (elementRef.value && !elementRef.value.contains(event.target as Node)) {
      callback()
    }
  }

  onMounted(() => {
    document.addEventListener('click', handleClickOutside)
  })

  onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside)
  })

  return {
    elementRef
  }
}
