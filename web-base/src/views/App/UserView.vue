<template>
  <div v-if="userProfile" class="max-w-2xl mx-auto p-6 min-h-screen" style="color: var(--text-color);">
    <div class="flex items-center gap-6 mb-8">
      <div class="w-24 h-24 rounded-full overflow-hidden border-2 flex-shrink-0" style="border-color: var(--border);">
        <img v-if="userProfile.profilePicture" :src="userProfile.profilePicture" class="w-full h-full object-cover" />
        <div v-else class="w-full h-full flex items-center justify-center" style="background-color: var(--surface); color: var(--muted);"><UserIcon class="w-10 h-10" /></div>
      </div>
      <div class="flex-1 min-w-0">
        <h1 class="text-3xl font-bold truncate">{{ userProfile.username }}</h1>
        <p v-if="userProfile.bio" class="text-sm mt-1 italic opacity-70">"{{ userProfile.bio }}"</p>
        <p class="text-[10px] uppercase font-bold tracking-widest mt-2 opacity-40">Joined {{ new Date(userProfile.createdAt).getFullYear() }}</p>
      </div>
    </div>
  </div>

  <div v-else-if="loading" class="flex justify-center items-center min-h-[400px]">
    <div class="animate-spin rounded-full h-12 w-12 border-b-2" style="border-color: var(--accent);"></div>
  </div>

  <div v-else class="flex justify-center items-center min-h-[400px]" style="color: var(--muted);">
    User not found.
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import type { UserPublicProfile } from '@/api'
import { QueryBuilder } from '@/helpers/queryBuilder'
import { UserQueryFilters } from '@/queryFilters/userQueryFilters'
import { UserIcon } from '@heroicons/vue/24/outline'
import { UsersService } from '@/services/UsersService'

const route = useRoute()
const userProfile = ref<UserPublicProfile | null>(null)
const loading = ref(true)

async function fetchUser() {
  const username = route.params.name as string
  if (!username) return (loading.value = false)
  loading.value = true
  try {
    const query = new QueryBuilder().addParameter(UserQueryFilters.WithUsernames([username])).build()
    const result = await UsersService.getUsers(query)
    if (result.data?.[0]) {
      userProfile.value = result.data[0]
    }
  } finally {
    loading.value = false
  }
}

watch(() => route.params.name, fetchUser)
onMounted(fetchUser)
</script>
