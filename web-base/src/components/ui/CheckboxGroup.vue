<template>
  <div class="space-y-2">
    <div v-if="label" class="text-sm font-semibold" :style="{ color: 'var(--text-color)' }">
      {{ label }}
      <span v-if="required" class="text-red-500">*</span>
    </div>
    <div class="space-y-2">
      <Checkbox
        v-for="option in options"
        :key="option.value"
        :model-value="isChecked(option.value)"
        :label="option.label"
        :id="`checkbox-${option.value}`"
        :disabled="disabled"
        @update:model-value="toggleOption(option.value)"
      />
    </div>
    <p v-if="error" class="text-xs text-red-500">{{ error }}</p>
    <p v-if="hint" class="text-xs" :style="{ color: 'var(--muted)' }">{{ hint }}</p>
  </div>
</template>

<script setup lang="ts">
import Checkbox from './Checkbox.vue'

interface Option {
  value: string | number
  label: string
}

const props = defineProps<{
  modelValue: (string | number)[]
  options: Option[]
  label?: string
  disabled?: boolean
  required?: boolean
  error?: string
  hint?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: (string | number)[]]
}>()

function isChecked(value: string | number): boolean {
  return props.modelValue.includes(value)
}

function toggleOption(value: string | number) {
  const newValue = isChecked(value)
    ? props.modelValue.filter(v => v !== value)
    : [...props.modelValue, value]
  emit('update:modelValue', newValue)
}
</script>
