<template>
  <div class="flex items-center">
    <input
      :id="id"
      :checked="modelValue"
      type="checkbox"
      @change="updateValue"
      :disabled="disabled"
      class="w-4 h-4 rounded border cursor-pointer transition-colors"
      :style="{
        backgroundColor: modelValue ? 'var(--accent)' : 'var(--surface)',
        borderColor: 'var(--border)',
        accentColor: 'var(--accent)',
        opacity: disabled ? 0.6 : 1
      }"
    />
    <label v-if="label" :for="id" class="ml-2 text-sm cursor-pointer" :style="{ color: 'var(--text-color)' }">
      {{ label }}
    </label>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  modelValue: boolean
  label?: string
  disabled?: boolean
  id?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()

const computedId = computed(() => props.id || `checkbox-${Math.random().toString(36).substr(2, 9)}`)

function updateValue(event: Event) {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.checked)
}
</script>
