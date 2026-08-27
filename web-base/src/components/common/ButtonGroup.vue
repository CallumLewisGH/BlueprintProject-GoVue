<template>
  <div class="space-y-3">
    <div v-if="label" class="text-sm font-semibold" :style="{ color: 'var(--text-color)' }">
      {{ label }}
      <span v-if="required" class="text-red-500">*</span>
    </div>
    <div class="flex gap-2 flex-wrap">
      <button
        v-for="option in options"
        :key="option.value"
        @click="toggleOption(option.value)"
        class="px-4 py-2 rounded-lg font-medium transition-all border"
        :style="{
          backgroundColor: isSelected(option.value) ? 'var(--accent)' : 'var(--surface)',
          borderColor: isSelected(option.value) ? 'var(--accent)' : 'var(--border)',
          color: isSelected(option.value) ? 'var(--header-text)' : 'var(--text-color)'
        }"
      >
        {{ option.label }}
      </button>
    </div>
    <p v-if="error" class="text-xs text-red-500">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
interface Option {
  value: string | number
  label: string
}

const props = defineProps<{
  modelValue: (string | number)[]
  options: Option[]
  label?: string
  required?: boolean
  error?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: (string | number)[]]
}>()

function isSelected(value: string | number): boolean {
  return props.modelValue.includes(value)
}

function toggleOption(value: string | number) {
  const newValue = isSelected(value)
    ? props.modelValue.filter(v => v !== value)
    : [...props.modelValue, value]
  emit('update:modelValue', newValue)
}
</script>
