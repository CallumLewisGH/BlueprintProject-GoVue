import { ref } from 'vue'

export function useAsync<T, E = Error>(asyncFunction: () => Promise<T>) {
  const data = ref<T | null>(null)
  const error = ref<E | null>(null)
  const isLoading = ref(false)

  async function execute() {
    isLoading.value = true
    error.value = null
    
    const result = await asyncFunction()
    
    if (result instanceof Error) {
      error.value = result as E
      return
    }
    
    data.value = result
  }

  function reset() {
    data.value = null
    error.value = null
    isLoading.value = false
  }

  return {
    data,
    error,
    isLoading,
    execute,
    reset
  }
}
