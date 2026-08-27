<template>
  <div class="flex items-center gap-2">
    <input
      :id="id"
      :checked="modelValue === value"
      type="radio"
      :name="name"
      :value="value"
      @change="updateValue"
      :disabled="disabled"
      class="w-4 h-4 cursor-pointer transition-colors"
      :style="{
        accentColor: 'var(--accent)',
        opacity: disabled ? 0.6 : 1
      }"
    />
    <label v-if="label" :for="id" class="text-sm cursor-pointer" :style="{ color: 'var(--text-color)' }">
      {{ label }}
    </label>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  modelValue: string | number
  value: string | number
  label?: string
  name: string
  disabled?: boolean
  id?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
}>()

const computedId = computed(() => props.id || `radio-${Math.random().toString(36).substr(2, 9)}`)

function updateValue() {
  emit('update:modelValue', props.value)
}
</script>
