<template>
  <div class="space-y-2">
    <label class="block text-xs font-semibold" :style="{ color: 'var(--muted)' }">
      {{ label }}
      <span v-if="required" class="text-red-500">*</span>
    </label>
    <div class="relative">
      <input
        :type="showPassword ? 'text' : 'password'"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        @input="updateValue"
        @blur="handleBlur"
        @focus="handleFocus"
        class="w-full px-3 py-2 pr-10 border rounded-lg outline-none focus:ring-2 text-sm transition-colors"
        :style="{
          backgroundColor: 'var(--surface)',
          borderColor: isFocused ? 'var(--accent)' : 'var(--border)',
          color: 'var(--text-color)',
          '--tw-ring-color': 'var(--accent)',
          opacity: disabled ? 0.6 : 1
        }"
      />
      <button
        @click="showPassword = !showPassword"
        type="button"
        class="absolute right-3 top-1/2 transform -translate-y-1/2 transition-opacity hover:opacity-70"
        :style="{ color: 'var(--muted)' }"
        :aria-label="showPassword ? 'Hide password' : 'Show password'"
      >
        <EyeSlashIcon v-if="showPassword" class="w-5 h-5" />
        <EyeIcon v-else class="w-5 h-5" />
      </button>
    </div>
    <p v-if="error" class="text-xs text-red-500 mt-1">{{ error }}</p>
    <p v-if="hint" class="text-xs mt-1" :style="{ color: 'var(--muted)' }">{{ hint }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { EyeIcon, EyeSlashIcon } from '@heroicons/vue/24/outline'

defineProps<{
  modelValue: string
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
const showPassword = ref(false)

function updateValue(event: Event) {
  const target = event.target as HTMLInputElement
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
