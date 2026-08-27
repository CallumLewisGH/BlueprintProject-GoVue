<template>
  <div class="w-full">
    <div class="flex items-center justify-between mb-2">
      <label v-if="label" class="text-sm font-semibold" :style="{ color: 'var(--text-color)' }">
        {{ label }}
      </label>
      <span class="text-xs" :style="{ color: 'var(--muted)' }">{{ modelValue.length }}/{{ maxImages }}</span>
    </div>

    <p v-if="error || uploadError" class="text-xs mb-2" style="color: rgb(239, 68, 68);">
      {{ error || uploadError }}
    </p>

    <div class="flex flex-wrap gap-3">
      <div
        v-for="(url, index) in modelValue"
        :key="url"
        class="relative w-20 h-20 rounded-lg overflow-hidden border group"
        :style="{ borderColor: 'var(--border)' }"
      >
        <img :src="url" class="w-full h-full object-cover" alt="Uploaded photo" />
        <button
          type="button"
          @click="removeAt(index)"
          class="absolute top-1 right-1 p-0.5 rounded-full bg-black/60 text-white opacity-0 group-hover:opacity-100 transition-opacity"
          aria-label="Remove photo"
        >
          <XMarkIcon class="w-3.5 h-3.5" />
        </button>
      </div>

      <button
        v-if="modelValue.length < maxImages"
        type="button"
        @click="fileInput?.click()"
        :disabled="isUploading"
        class="w-20 h-20 rounded-lg border-2 border-dashed flex items-center justify-center transition-opacity disabled:opacity-60"
        :style="{ borderColor: 'var(--border)', color: 'var(--muted)' }"
      >
        <ArrowPathIcon v-if="isUploading" class="w-5 h-5 animate-spin" />
        <PlusIcon v-else class="w-6 h-6" />
      </button>
    </div>

    <input
      ref="fileInput"
      type="file"
      accept="image/jpeg,image/png,image/webp"
      class="hidden"
      @change="handleFileChange"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { XMarkIcon, PlusIcon, ArrowPathIcon } from '@heroicons/vue/24/outline'
import { useImageUpload } from '@/composables/useImageUpload'

const props = withDefaults(defineProps<{
  modelValue: string[]
  label?: string
  maxImages?: number
  error?: string
  purpose?: string
}>(), {
  maxImages: 10,
  purpose: 'profile'
})

const emit = defineEmits<{
  'update:modelValue': [value: string[]]
}>()

const fileInput = ref<HTMLInputElement | null>(null)
const { isUploading, error: uploadError, uploadImage } = useImageUpload()

async function handleFileChange(event: Event) {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  target.value = ''
  if (!file) return

  const publicUrl = await uploadImage(file, props.purpose)
  if (!publicUrl) return

  emit('update:modelValue', [...props.modelValue, publicUrl])
}

function removeAt(index: number) {
  emit('update:modelValue', props.modelValue.filter((_, i) => i !== index))
}
</script>
