<template>
  <div class="w-full">
    <label class="block text-xs font-semibold mb-2" :style="{ color: 'var(--muted)' }">
      {{ label }}
      <span v-if="required" class="text-red-500">*</span>
    </label>
    <div class="relative">
      <input
        :value="searchQuery"
        type="text"
        :placeholder="placeholder"
        @input="handleInput"
        @focus="isOpen = true"
        @blur="handleBlur"
        class="w-full px-3 py-2 border rounded-lg outline-none focus:ring-2 text-sm"
        :style="{
          backgroundColor: 'var(--surface)',
          borderColor: 'var(--border)',
          color: 'var(--text-color)',
          '--tw-ring-color': 'var(--accent)'
        }"
      />
      
      <div v-if="isOpen && filteredOptions.length > 0"
           class="absolute top-full left-0 right-0 mt-1 rounded-lg shadow-xl border z-50 max-h-60 overflow-y-auto"
           :style="{ backgroundColor: 'var(--card-bg)', borderColor: 'var(--border)' }">
        <button
          v-for="(option, index) in filteredOptions"
          :key="option.value"
          @click="selectOption(option)"
          @mouseover="highlightedIndex = index"
          class="w-full text-left px-3 py-2 text-sm transition-colors"
          :style="{
            backgroundColor: highlightedIndex === index ? 'var(--border)' : 'transparent',
            color: 'var(--text-color)'
          }"
        >
          {{ option.label }}
        </button>
      </div>

      <p v-if="error" class="text-xs text-red-500 mt-1">{{ error }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface Option {
  value: string | number
  label: string
}

const props = defineProps<{
  modelValue: string | number
  options: Option[]
  label?: string
  placeholder?: string
  disabled?: boolean
  required?: boolean
  error?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
  'blur': []
}>()

const searchQuery = ref('')
const isOpen = ref(false)
const highlightedIndex = ref(0)

const filteredOptions = computed(() => {
  if (!searchQuery.value) return props.options
  
  const query = searchQuery.value.toLowerCase()
  return props.options.filter(opt => 
    opt.label.toLowerCase().includes(query)
  )
})

function handleInput(event: Event) {
  const target = event.target as HTMLInputElement
  searchQuery.value = target.value
  highlightedIndex.value = 0
}

function selectOption(option: Option) {
  emit('update:modelValue', option.value)
  searchQuery.value = option.label
  isOpen.value = false
}

function handleBlur() {
  setTimeout(() => {
    isOpen.value = false
  }, 150)
  emit('blur')
}
</script>
