import { ref, reactive } from 'vue'

export interface FormState {
  isLoading: boolean
  submitError: string | null
}

export function useForm<T extends Record<string, any>>(initialValues: T) {
  const formData = reactive<T>({ ...initialValues })
  const state = reactive<FormState>({
    isLoading: false,
    submitError: null
  })

  function updateField(fieldName: string, value: any) {
    formData[fieldName] = value
  }

  function setError(error: string) {
    state.submitError = error
  }

  function clearError() {
    state.submitError = null
  }

  function setLoading(loading: boolean) {
    state.isLoading = loading
  }

  function resetForm() {
    Object.keys(initialValues).forEach(key => {
      formData[key] = initialValues[key]
    })
    clearError()
    setLoading(false)
  }

  return {
    formData,
    state,
    updateField,
    setError,
    clearError,
    setLoading,
    resetForm
  }
}
