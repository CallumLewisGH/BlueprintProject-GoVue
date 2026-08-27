<template>
  <div class="space-y-4">
    <div class="p-4 rounded-xl border text-center"
         :style="{ backgroundColor: 'var(--card-bg)', borderColor: 'var(--border)' }">
      <component :is="icon" v-if="icon" class="w-9 h-9 mx-auto mb-3" :style="{ color: 'var(--muted)' }" />
      <h3 class="font-bold text-lg mb-1" :style="{ color: 'var(--text-color)' }">{{ title }}</h3>
      <p :style="{ color: 'var(--muted)' }">{{ message }}</p>
    </div>
    <slot />
  </div>
</template>

<script setup lang="ts">
import type { Component } from 'vue'

defineProps<{
  icon?: Component
  title: string
  message: string
}>()
</script>
