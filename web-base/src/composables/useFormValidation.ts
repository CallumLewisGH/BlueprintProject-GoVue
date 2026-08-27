import { ref, reactive, computed } from 'vue'

interface ValidationRule {
  validate: (value: any) => boolean | string
  message?: string
}

export function useFormValidation() {
  const errors = reactive<Record<string, string>>({})
  const touched = reactive<Record<string, boolean>>({})

  const rules: Record<string, ValidationRule[]> = {}

  function registerField(fieldName: string, fieldRules: ValidationRule[]) {
    rules[fieldName] = fieldRules
    if (!errors[fieldName]) {
      errors[fieldName] = ''
    }
    if (!touched[fieldName]) {
      touched[fieldName] = false
    }
  }

  function validateField(fieldName: string, value: any): boolean {
    const fieldRules = rules[fieldName]
    if (!fieldRules || fieldRules.length === 0) return true

    for (const rule of fieldRules) {
      const result = rule.validate(value)
      if (result !== true) {
        errors[fieldName] = typeof result === 'string' ? result : (rule.message || 'Invalid value')
        return false
      }
    }
    errors[fieldName] = ''
    return true
  }

  function validateAll(formData: Record<string, any>): boolean {
    let isValid = true

    for (const fieldName of Object.keys(rules)) {
      if (!validateField(fieldName, formData[fieldName])) {
        isValid = false
      }
    }

    return isValid
  }

  function markTouched(fieldName: string) {
    touched[fieldName] = true
  }

  function resetValidation() {
    Object.keys(errors).forEach(key => {
      errors[key] = ''
    })
    Object.keys(touched).forEach(key => {
      touched[key] = false
    })
  }

  const hasErrors = computed(() => Object.values(errors).some(e => e !== ''))

  return {
    errors,
    touched,
    registerField,
    validateField,
    validateAll,
    markTouched,
    resetValidation,
    hasErrors
  }
}
