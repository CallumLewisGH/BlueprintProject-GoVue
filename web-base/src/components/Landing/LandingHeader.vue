<template>
  <nav class="absolute w-full z-20 top-0 left-0 border-b backdrop-blur-md"
       :style="{ 
         borderColor: 'var(--border)',
         backgroundColor: 'var(--header-bg)'  
       }">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-20">
        <div class="flex-shrink-0 flex items-center gap-2 cursor-pointer" @click="goHome">
          <span class="font-bold text-2xl tracking-wider uppercase text-white">App</span>
        </div>

        <div class="hidden md:flex items-center space-x-8">
          <router-link
            to="/about" 
            :class="['text-white hover:text-violet-200 transition-colors', 
                     { 'font-semibold': activePage === 'about' }]"
          >
            About
          </router-link>
          <router-link 
            to="/contact" 
            :class="['text-white hover:text-violet-200 transition-colors', 
                     { 'font-semibold': activePage === 'contact' }]"
          >
            Contact
          </router-link>

          <button
            @click="toggleDarkMode"
            class="p-2 rounded-full transition text-white hover:text-violet-200"
            :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
          >
            <SunIcon v-if="isDark" class="w-5 h-5" />
            <MoonIcon v-else class="w-5 h-5" />
          </button>

          <button 
            @click="handleAuth"
            class="px-6 py-2 rounded-lg font-semibold transition-all transform hover:scale-105 shadow-lg text-white"
            :style="{ backgroundColor: 'var(--accent)' }"
          >
            Sign In / Join
          </button>
        </div>

        <div class="md:hidden flex items-center gap-4">
          <button
            @click="toggleDarkMode"
            class="p-2 text-white"
            :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
          >
            <SunIcon v-if="isDark" class="w-5 h-5" />
            <MoonIcon v-else class="w-5 h-5" />
          </button>
          <button @click="isMenuOpen = !isMenuOpen" class="focus:outline-none text-white" aria-label="Toggle menu">
            <Bars3Icon class="w-6 h-6" />
          </button>
        </div>
      </div>
    </div>
    
    <!-- Mobile menu -->
    <div v-if="isMenuOpen" class="md:hidden border-t"
         :style="{ 
           borderColor: 'var(--border)',
           backgroundColor: isDark ? 'rgba(0, 0, 0, 0.3)' : 'rgba(0, 0, 0, 0.2)'
         }">
      <div class="px-4 py-6 space-y-4">
        <router-link
          to="/about" 
          :class="['block text-white hover:text-violet-200 transition-colors py-2',
                   { 'font-semibold': activePage === 'about' }]"
          @click="isMenuOpen = false"
        >
          About
        </router-link>
        <router-link 
          to="/contact" 
          :class="['block text-white hover:text-violet-200 transition-colors py-2',
                   { 'font-semibold': activePage === 'contact' }]"
          @click="isMenuOpen = false"
        >
          Contact
        </router-link>
        <button 
          @click="handleAuth"
          class="w-full px-6 py-3 rounded-lg font-semibold transition-all shadow-lg text-white"
          :style="{ backgroundColor: 'var(--accent)' }"
        >
          Sign In / Join
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import router from '@/router'
import { ref } from 'vue'
import { useDarkMode } from '@/composables/useDarkMode'
import { SunIcon, MoonIcon, Bars3Icon } from '@heroicons/vue/24/outline'

const { isDark, toggleDarkMode } = useDarkMode()
const isMenuOpen = ref(false)

const props = defineProps({
  activePage: {
    type: String,
    default: ''
  }
})

const goHome = () => {
  router.push('/')
}

const handleAuth = () => {
  router.push('/login')
}

// Close mobile menu when route changes
import { watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
watch(() => route.path, () => {
  isMenuOpen.value = false
})
</script>