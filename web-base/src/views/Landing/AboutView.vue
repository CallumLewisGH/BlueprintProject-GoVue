<template>
  <!-- Hero Section -->
  <header class="relative pt-32 pb-20 px-4 overflow-hidden"
          :style="{ backgroundColor: 'var(--header-bg)' }">
    <div class="absolute inset-0 pointer-events-none">
      <div class="absolute -top-[10%] left-1/4 w-[40%] h-[40%] rounded-full blur-[120px]"
            :style="{ backgroundColor: 'var(--glow-orange)' }"></div>
      <div class="absolute -bottom-[10%] right-1/4 w-[40%] h-[40%] rounded-full blur-[120px]"
            :style="{ backgroundColor: 'var(--glow-amber)' }"></div>
    </div>

    <div class="relative z-10 text-center max-w-4xl mx-auto">
      <span class="inline-block py-1 px-3 rounded-full text-sm font-bold mb-6 tracking-wide uppercase border"
            :style="{
              backgroundColor: 'var(--badge-bg)',
              borderColor: 'var(--badge-border)',
              color: 'var(--accent)'
            }">
        About
      </span>
      <h1 class="text-5xl md:text-6xl font-extrabold mb-6 leading-tight"
          :style="{ color: 'var(--header-text)' }">
        Built by <span :style="{ color: 'var(--accent)' }">You</span>
      </h1>
      <p class="text-lg md:text-xl mb-10 max-w-2xl mx-auto"
          :style="{ color: 'var(--header-muted)' }">
        Replace this with the real story of why this product exists.
      </p>
    </div>
  </header>

  <!-- Main Content -->
  <main class="py-20 px-4 sm:px-6 lg:px-8"
        :style="{ backgroundColor: 'var(--surface)' }">
    <div class="max-w-5xl mx-auto">

      <!-- Mission Section -->
      <section class="mb-20">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
          <div>
            <h2 class="text-3xl md:text-4xl font-bold mb-6"
                :style="{ color: 'var(--text-color)' }">
              Why We Exist
            </h2>
            <div class="h-1 w-20 rounded mb-8"
                  :style="{ backgroundColor: 'var(--accent)' }"></div>
            <p class="text-lg mb-6 leading-relaxed"
                :style="{ color: 'var(--muted)' }">
              Replace this paragraph with your actual mission statement.
            </p>
          </div>
          <div class="relative">
            <div class="aspect-square rounded-2xl overflow-hidden border-2 shadow-2xl"
                  :style="{
                    borderColor: 'var(--border)',
                    boxShadow: isDark ? '0 25px 50px -12px rgba(0, 0, 0, 0.5)' : '0 25px 50px -12px rgba(0, 0, 0, 0.15)'
                  }">
              <div class="w-full h-full flex items-center justify-center"
                    :style="{
                      backgroundColor: isDark ? 'var(--card-bg)' : 'var(--border)'
                    }">
                <SparklesIcon class="w-14 h-14" :style="{ color: 'var(--accent)' }" />
              </div>
            </div>
            <div class="absolute -bottom-4 -right-4 w-2/3 h-2/3 rounded-2xl border-2 -z-10"
                  :style="{
                    borderColor: 'var(--accent)',
                    backgroundColor: 'var(--badge-bg)'
                  }"></div>
          </div>
        </div>
      </section>

      <!-- Values Section -->
      <section class="mb-20">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold mb-4"
              :style="{ color: 'var(--text-color)' }">
            Our Values
          </h2>
          <div class="h-1 w-20 mx-auto rounded"
                :style="{ backgroundColor: 'var(--accent)' }"></div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div
            v-for="(value, index) in values"
            :key="index"
            class="group p-8 rounded-2xl hover:-translate-y-2 transition-all duration-300 border"
            :style="{
              backgroundColor: 'var(--card-bg)',
              borderColor: 'var(--border)'
            }"
          >
            <div class="mb-6 w-20 h-20 flex items-center justify-center rounded-full group-hover:scale-110 transition-transform"
                  :style="{
                    backgroundColor: isDark ? 'var(--border)' : 'var(--surface)'
                  }">
              <component :is="value.icon" class="w-10 h-10" :style="{ color: 'var(--accent)' }" />
            </div>
            <h3 class="text-xl font-bold mb-3"
                :style="{ color: 'var(--text-color)' }">
              {{ value.title }}
            </h3>
            <p class="leading-relaxed"
                :style="{ color: 'var(--muted)' }">
              {{ value.desc }}
            </p>
          </div>
        </div>
      </section>

      <!-- Call to Action -->
      <section class="text-center py-12 px-6 rounded-3xl border"
                :style="{
                  backgroundColor: 'var(--card-bg)',
                  borderColor: 'var(--border)'
                }">
        <h2 class="text-3xl md:text-4xl font-bold mb-6"
            :style="{ color: 'var(--text-color)' }">
          Ready to Get Started?
        </h2>
        <p class="text-lg mb-8 max-w-2xl mx-auto"
            :style="{ color: 'var(--muted)' }">
          Replace this call-to-action copy with the real pitch.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="handleAuth"
            class="px-8 py-4 rounded-lg text-lg font-bold transition-all transform hover:scale-105 shadow-lg text-white"
            :style="{ backgroundColor: 'var(--accent)' }"
          >
            Get Started
          </button>
        </div>
      </section>
    </div>
  </main>
</template>

<script setup lang="ts">
import router from '@/router'

import { useDarkMode } from '@/composables/useDarkMode'
import { UsersIcon, ShieldCheckIcon, SparklesIcon } from '@heroicons/vue/24/outline'

const { isDark } = useDarkMode()

const handleAuth = () => {
  router.push('/login')
}

// Placeholder values - swap for the real ones.
const values = [
  {
    title: 'Value One',
    desc: 'Describe the first thing you stand for.',
    icon: ShieldCheckIcon
  },
  {
    title: 'Value Two',
    desc: 'Describe the second thing you stand for.',
    icon: UsersIcon
  },
  {
    title: 'Value Three',
    desc: 'Describe the third thing you stand for.',
    icon: SparklesIcon
  }
]
</script>
