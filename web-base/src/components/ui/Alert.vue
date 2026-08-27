<template>
  <div class="p-4 rounded-lg border" :style="{ backgroundColor: `var(--${type}-bg, var(--card-bg))`, borderColor: 'var(--border)' }">
    <div class="flex items-start gap-3">
      <component :is="icon" class="w-5 h-5 flex-shrink-0" :style="{ color: typeColor }" />
      <div class="flex-1 min-w-0">
        <h3 v-if="title" class="font-semibold text-sm mb-1" :style="{ color: 'var(--text-color)' }">
          {{ title }}
        </h3>
        <p class="text-sm" :style="{ color: typeColor }">
          <slot>{{ message }}</slot>
        </p>
      </div>
      <button
        v-if="closable"
        @click="isVisible = false"
        class="flex-shrink-0 p-1 rounded hover:opacity-70 transition-opacity"
        :style="{ color: 'var(--muted)' }"
        aria-label="Close alert"
      >
        <XMarkIcon class="w-4 h-4" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import {
  CheckCircleIcon,
  XCircleIcon,
  ExclamationTriangleIcon,
  InformationCircleIcon,
  XMarkIcon
} from '@heroicons/vue/24/outline'

const props = defineProps<{
  type?: 'success' | 'error' | 'warning' | 'info'
  title?: string
  message?: string
  closable?: boolean
}>()

const isVisible = ref(true)

const iconMap = {
  success: CheckCircleIcon,
  error: XCircleIcon,
  warning: ExclamationTriangleIcon,
  info: InformationCircleIcon
}

const colorMap = {
  success: 'rgb(34, 197, 94)',
  error: 'rgb(239, 68, 68)',
  warning: 'rgb(234, 179, 8)',
  info: 'rgb(59, 130, 246)'
}

const type = computed(() => props.type || 'info')
const icon = computed(() => iconMap[type.value])
const typeColor = computed(() => colorMap[type.value])

defineExpose({
  isVisible
})
</script>
