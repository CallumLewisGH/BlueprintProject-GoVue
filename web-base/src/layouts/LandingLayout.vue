<template>
  <div class="min-h-screen flex flex-col transition-colors duration-300 font-sans" 
       :style="{ backgroundColor: 'var(--bg-color)', color: 'var(--text-color)' }">
    
    <LandingHeader :active-page="activePage" class="flex-shrink-0" />

    <main class="flex-1 pt-20">
      <router-view />
    </main>

    <footer class="py-12 border-t flex-shrink-0"
            :style="{ 
              backgroundColor: 'var(--bg-color)',
              borderColor: 'var(--border)'
            }">
      <div class="max-w-7xl mx-auto px-4 text-center space-y-3">
        <p class="text-sm"
           :style="{ color: 'var(--muted)' }">
          © 2026 App. All rights reserved.
        </p>
        <router-link to="/privacy" class="text-sm underline hover:no-underline"
                     :style="{ color: 'var(--muted)' }">
          Privacy Policy
        </router-link>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import LandingHeader from '@/components/Landing/LandingHeader.vue'

defineProps({
  activePage: {
    type: String,
    default: ''
  }
})
</script>