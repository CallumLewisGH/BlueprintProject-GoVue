<template>
  <!-- Hero Section -->
  <header class="relative h-[80vh] flex items-center justify-center overflow-hidden"
          :style="{ backgroundColor: 'var(--header-bg)' }">
    <div class="absolute inset-0 pointer-events-none">
      <div class="absolute -top-[10%] -left-[10%] w-[40%] h-[40%] rounded-full blur-[120px]"
            :style="{ backgroundColor: 'var(--glow-orange)' }"></div>
      <div class="absolute -bottom-[10%] -right-[10%] w-[40%] h-[40%] rounded-full blur-[120px]"
            :style="{ backgroundColor: 'var(--glow-amber)' }"></div>
    </div>

    <div class="relative z-10 text-center px-4 max-w-4xl mx-auto">
      <span class="inline-block py-1 px-3 rounded-full text-sm font-bold mb-6 tracking-wide uppercase border"
            :style="{
              backgroundColor: 'var(--badge-bg)',
              borderColor: 'var(--badge-border)',
              color: 'var(--accent)'
            }">
        Beta
      </span>
      <h1 class="text-5xl md:text-7xl font-extrabold mb-6 leading-tight"
          :style="{ color: 'var(--header-text)' }">
        Your Product <span :style="{ color: 'var(--accent)' }">Name</span>
      </h1>
      <p class="text-lg md:text-xl mb-10 max-w-2xl mx-auto"
          :style="{ color: 'var(--header-muted)' }">
        Replace this hero copy with your actual pitch.
      </p>

      <div class="flex flex-col sm:flex-row gap-4 justify-center">
        <router-link to="/login" class="px-8 py-4 rounded-lg text-lg font-bold transition shadow-lg text-white"
                :style="{ backgroundColor: 'var(--accent)' }">
          Get Started
        </router-link>
      </div>
    </div>
  </header>

  <!-- Features Section -->
  <section class="py-20 px-4 sm:px-6 lg:px-8"
            :style="{ backgroundColor: 'var(--surface)' }">
    <div class="max-w-7xl mx-auto">
      <div class="text-center mb-16">
        <h2 class="text-3xl md:text-4xl font-bold mb-4"
            :style="{ color: 'var(--text-color)' }">
          Feature Highlights
        </h2>
        <div class="h-1 w-20 mx-auto rounded"
              :style="{ backgroundColor: 'var(--accent)' }"></div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div
          v-for="(feature, index) in features"
          :key="index"
          class="group p-8 rounded-2xl shadow-xl hover:-translate-y-2 transition-transform duration-300 border"
          :style="{
            backgroundColor: 'var(--card-bg)',
            borderColor: 'var(--border)'
          }"
        >
          <div class="mb-6 w-16 h-16 flex items-center justify-center rounded-full group-hover:scale-110 transition-transform shadow-sm"
                :style="{
                  backgroundColor: isDark ? 'var(--border)' : 'var(--surface)'
                }">
            <component :is="feature.icon" class="w-8 h-8" :style="{ color: 'var(--accent)' }" />
          </div>
          <h3 class="text-xl font-bold mb-3"
              :style="{ color: 'var(--text-color)' }">
            {{ feature.title }}
          </h3>
          <p class="leading-relaxed"
              :style="{ color: 'var(--muted)' }">
            {{ feature.desc }}
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { useDarkMode } from '@/composables/useDarkMode'
import { SparklesIcon, ShieldCheckIcon, BoltIcon } from '@heroicons/vue/24/outline'

const { isDark } = useDarkMode()

// Placeholder feature list - swap for the real ones.
const features = [
  {
    title: 'Feature One',
    desc: 'Describe the first thing that makes this product worth using.',
    icon: SparklesIcon
  },
  {
    title: 'Feature Two',
    desc: 'Describe the second thing that makes this product worth using.',
    icon: ShieldCheckIcon
  },
  {
    title: 'Feature Three',
    desc: 'Describe the third thing that makes this product worth using.',
    icon: BoltIcon
  }
]
</script>
