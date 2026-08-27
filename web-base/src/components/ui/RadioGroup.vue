<template>
  <div class="space-y-3">
    <div v-if="label" class="text-sm font-semibold" :style="{ color: 'var(--text-color)' }">
      {{ label }}
      <span v-if="required" class="text-red-500">*</span>
    </div>
    <div class="space-y-2">
      <Radio
        v-for="option in options"
        :key="option.value"
        :model-value="modelValue"
        :value="option.value"
        :label="option.label"
        :name="name"
        :disabled="disabled"
        @update:model-value="updateValue"
      />
    </div>
    <p v-if="error" class="text-xs text-red-500">{{ error }}</p>
    <p v-if="hint" class="text-xs" :style="{ color: 'var(--muted)' }">{{ hint }}</p>
  </div>
</template>

<script setup lang="ts">
import Radio from './Radio.vue'

interface Option {
  value: string | number
  label: string
}

defineProps<{
  modelValue: string | number
  options: Option[]
  name: string
  label?: string
  disabled?: boolean
  required?: boolean
  error?: string
  hint?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
}>()

function updateValue(value: string | number) {
  emit('update:modelValue', value)
}
</script>
