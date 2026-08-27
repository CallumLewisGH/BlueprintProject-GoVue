<template>
  <div class="border-b" :style="{ borderColor: 'var(--border)' }">
    <div class="flex space-x-4">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        @click="$emit('select', tab.id)"
        :class="[
          'px-4 py-2 font-medium transition-all border-b-2',
          isActive(tab.id) ? 'border-b-accent' : 'border-b-transparent'
        ]"
        :style="{
          borderBottomColor: isActive(tab.id) ? 'var(--accent)' : 'transparent',
          color: isActive(tab.id) ? 'var(--accent)' : 'var(--muted)'
        }"
      >
        {{ tab.label }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Tab {
  id: string | number
  label: string
}

const props = defineProps<{
  tabs: Tab[]
  activeTab: string | number
}>()

defineEmits<{
  'select': [tabId: string | number]
}>()

function isActive(tabId: string | number): boolean {
  return tabId === props.activeTab
}
</script>
