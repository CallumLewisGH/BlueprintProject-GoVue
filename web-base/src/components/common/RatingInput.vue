<template>
  <div class="space-y-2">
    <div class="flex items-center justify-between">
      <label class="text-sm font-semibold" :style="{ color: 'var(--text-color)' }">
        {{ label }}
        <span v-if="required" class="text-red-500">*</span>
      </label>
      <span v-if="showRating" class="text-sm flex items-center gap-1">{{ modelValue }}/{{ maxRating }} <StarIcon class="w-4 h-4" style="color: var(--accent);" /></span>
    </div>
    <div class="flex gap-1">
      <button
        v-for="i in maxRating"
        :key="i"
        @click="updateRating(i)"
        @hover="hoverRating = i"
        class="transition-transform hover:scale-110 cursor-pointer"
        :style="{ opacity: i <= (hoverRating || modelValue) ? 1 : 0.3, color: 'var(--accent)' }"
      >
        <StarIcon class="w-6 h-6" />
      </button>
    </div>
    <p v-if="hint" class="text-xs" :style="{ color: 'var(--muted)' }">{{ hint }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { StarIcon } from '@heroicons/vue/24/solid'

defineProps<{
  modelValue: number
  label?: string
  maxRating?: number
  required?: boolean
  showRating?: boolean
  hint?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: number]
}>()

const hoverRating = ref(0)

function updateRating(value: number) {
  emit('update:modelValue', value)
  hoverRating.value = 0
}
</script>
