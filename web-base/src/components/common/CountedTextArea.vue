<template>
  <div class="w-full">
    <div class="flex items-center justify-between mb-2">
      <label v-if="label" class="text-sm font-semibold" :style="{ color: 'var(--text-color)' }">
        {{ label }}
      </label>
      <span v-if="showCount" class="text-xs" :style="{ color: 'var(--muted)' }">
        {{ modelValue.length }}/{{ maxLength }}
      </span>
    </div>
    <TextArea
      :model-value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :rows="rows"
      :error="error"
      @update:model-value="updateValue"
      @blur="$emit('blur')"
      @focus="$emit('focus')"
    />
  </div>
</template>

<script setup lang="ts">
import TextArea from '@/components/ui/TextArea.vue'

const props = defineProps<{
  modelValue: string
  label?: string
  placeholder?: string
  disabled?: boolean
  maxLength?: number
  rows?: number
  error?: string
  showCount?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string]
  'blur': []
  'focus': []
}>()

function updateValue(value: string) {
  if (props.maxLength && value.length > props.maxLength) {
    return
  }
  emit('update:modelValue', value)
}
</script>
