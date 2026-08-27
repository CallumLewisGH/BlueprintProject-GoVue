<template>
  <div class="w-full">
    <label v-if="label" class="block text-xs font-semibold mb-1" :style="{ color: 'var(--muted)' }">
      {{ label }}
      <span v-if="required" class="text-red-500">*</span>
    </label>
    <select
      :value="modelValue"
      :disabled="disabled"
      @change="updateValue"
      @blur="handleBlur"
      @focus="handleFocus"
      class="w-full px-3 py-2 border rounded-lg outline-none focus:ring-2 text-sm transition-colors"
      :style="{
        backgroundColor: 'var(--surface)',
        borderColor: isFocused ? 'var(--accent)' : 'var(--border)',
        color: 'var(--text-color)',
        '--tw-ring-color': 'var(--accent)',
        opacity: disabled ? 0.6 : 1
      }"
    >
      <option value="" disabled>{{ placeholder }}</option>
      <option v-for="option in options" :key="option.value" :value="option.value">
        {{ option.label }}
      </option>
    </select>
    <p v-if="error" class="text-xs text-red-500 mt-1">{{ error }}</p>
    <p v-if="hint" class="text-xs mt-1" :style="{ color: 'var(--muted)' }">{{ hint }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Option {
  value: string | number
  label: string
}

defineProps<{
  modelValue: string | number
  options: Option[]
  label?: string
  placeholder?: string
  disabled?: boolean
  required?: boolean
  error?: string
  hint?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string]
  'blur': []
  'focus': []
}>()

const isFocused = ref(false)

function updateValue(event: Event) {
  const target = event.target as HTMLSelectElement
  emit('update:modelValue', target.value)
}

function handleBlur() {
  isFocused.value = false
  emit('blur')
}

function handleFocus() {
  isFocused.value = true
  emit('focus')
}
</script>
