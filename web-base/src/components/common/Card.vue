<template>
  <div class="p-4 rounded-xl border cursor-pointer hover:scale-[1.01] transition-all overflow-hidden"
       :style="{ backgroundColor: 'var(--card-bg)', borderColor: 'var(--border)' }"
       @click="$emit('click')">
    <div v-if="$slots.header" class="mb-3">
      <slot name="header" />
    </div>

    <div class="flex items-center gap-4">
      <div v-if="$slots.icon" class="flex-shrink-0">
        <slot name="icon" />
      </div>

      <div class="flex-1 min-w-0">
        <h3 class="font-bold truncate" :style="{ color: 'var(--text-color)' }">
          <slot name="title" />
        </h3>
        <p v-if="$slots.description" class="text-sm line-clamp-2" :style="{ color: 'var(--muted)' }">
          <slot name="description" />
        </p>
      </div>

      <div v-if="$slots.action" class="flex-shrink-0">
        <slot name="action" />
      </div>
    </div>

    <div v-if="$slots.footer" class="mt-3 pt-3 border-t" :style="{ borderColor: 'var(--border)' }">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup lang="ts">
defineEmits<{
  'click': []
}>()
</script>
