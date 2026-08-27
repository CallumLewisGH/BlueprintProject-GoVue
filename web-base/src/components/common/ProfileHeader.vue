<template>
  <div v-if="profile" class="flex items-center gap-4">
    <div class="flex-shrink-0">
      <div v-if="profile.profilePicture" 
           class="w-20 h-20 rounded-full overflow-hidden border-2" 
           :style="{ borderColor: 'var(--border)' }">
        <img :src="profile.profilePicture" alt="Profile" class="w-full h-full object-cover" />
      </div>
      <div v-else
           class="w-20 h-20 rounded-full flex items-center justify-center border-2"
           :style="{ backgroundColor: 'var(--surface)', borderColor: 'var(--border)', color: 'var(--muted)' }">
        <UserIcon class="w-10 h-10" />
      </div>
    </div>

    <div class="flex-1 min-w-0">
      <h2 class="text-2xl font-bold truncate" :style="{ color: 'var(--text-color)' }">
        {{ profile.username }}
      </h2>
      <p v-if="profile.email" class="text-sm" :style="{ color: 'var(--muted)' }">
        {{ profile.email }}
      </p>
      <p v-if="profile.bio" class="text-sm italic mt-1" :style="{ color: 'var(--muted)' }">
        "{{ profile.bio }}"
      </p>
    </div>

    <div v-if="showEditButton" class="flex-shrink-0">
      <button 
        @click="$emit('edit')"
        class="p-3 rounded-full hover:opacity-80 transition-opacity"
        :style="{ backgroundColor: 'var(--accent)', color: 'var(--header-text)' }"
        aria-label="Edit profile"
      >
        <PencilSquareIcon class="w-5 h-5" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { UserIcon, PencilSquareIcon } from '@heroicons/vue/24/outline'

interface Profile {
  username: string
  email?: string
  bio?: string
  profilePicture?: string
}

defineProps<{
  profile: Profile | null
  showEditButton?: boolean
}>()

defineEmits<{
  'edit': []
}>()
</script>
