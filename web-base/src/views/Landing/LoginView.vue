<template>
  <div class="min-h-screen transition-colors duration-300 font-sans flex items-center justify-center" 
       :style="{ backgroundColor: 'var(--bg-color)', color: 'var(--text-color)' }">
    
    <!-- Background decorative elements matching landing page -->
    <div class="absolute inset-0 pointer-events-none overflow-hidden">
      <div class="absolute -top-[10%] -left-[10%] w-[40%] h-[40%] rounded-full blur-[120px]"
           :style="{ backgroundColor: 'var(--glow-orange)' }"></div>
      <div class="absolute -bottom-[10%] -right-[10%] w-[40%] h-[40%] rounded-full blur-[120px]"
           :style="{ backgroundColor: 'var(--glow-amber)' }"></div>
    </div>

    <!-- Simple back button to home -->
    <router-link to="/"
                 class="absolute top-8 left-8 p-2 rounded-full transition text-white hover:text-violet-200 flex items-center gap-1"
                 :style="{ backgroundColor: isDark ? 'rgba(0, 0, 0, 0.2)' : 'rgba(0, 0, 0, 0.1)' }">
      <ArrowLeftIcon class="w-4 h-4" /> Back
    </router-link>

    <!-- Dark mode toggle -->
    <button
      @click="toggleDarkMode"
      class="absolute top-8 right-8 p-2 rounded-full transition text-white hover:text-violet-200"
      :style="{ backgroundColor: isDark ? 'rgba(0, 0, 0, 0.2)' : 'rgba(0, 0, 0, 0.1)' }"
      :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
    >
      <SunIcon v-if="isDark" class="w-5 h-5" />
      <MoonIcon v-else class="w-5 h-5" />
    </button>

    <!-- Main login box -->
    <div class="relative z-10 w-full max-w-sm px-4">
      <div class="p-8 rounded-2xl shadow-xl border text-center"
           :style="{ 
             backgroundColor: 'var(--card-bg)',
             borderColor: 'var(--border)'
           }">
        
        <!-- Logo -->
        <div class="mb-6">
          <span class="font-bold text-2xl tracking-wider uppercase"
                :style="{ color: 'var(--text-color)' }">
            App
          </span>
        </div>

        <!-- Icon -->
        <div class="w-16 h-16 mx-auto mb-6 flex items-center justify-center rounded-full text-3xl shadow-sm"
             :style="{
               backgroundColor: isDark ? 'var(--border)' : 'var(--surface)'
             }">
          <UserCircleIcon class="w-11 h-11" :style="{ color: 'var(--accent)' }" />
        </div>

        <!-- Headline -->
        <h1 class="text-2xl font-bold mb-2"
            :style="{ color: 'var(--text-color)' }">
          Welcome <span :style="{ color: 'var(--accent)' }">Back</span>
        </h1>
        <p class="text-sm mb-8"
           :style="{ color: 'var(--muted)' }">
          Sign in to continue
        </p>

        <!-- Google Sign In Button -->
        <button
          @click="signInWithGoogle"
          class="w-full py-4 rounded-lg font-semibold transition-all transform hover:scale-[1.02] shadow-lg flex items-center justify-center gap-3 mb-6"
          :style="{
            backgroundColor: 'var(--surface)',
            border: '1px solid var(--border)',
            color: 'var(--text-color)'
          }"
        >
          <img src="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/google.svg" alt="Google" class="w-5 h-5"/>
          <span>Sign in with Google</span>
        </button>

        <p class="text-xs mb-2" :style="{ color: 'var(--muted)' }">
          By continuing you agree to our
          <router-link to="/privacy" class="underline hover:no-underline" :style="{ color: 'var(--accent)' }">Privacy Policy</router-link>.
        </p>

        <!-- Beta badge -->
        <div class="mt-6 pt-6 border-t"
             :style="{ borderColor: 'var(--border)' }">
          <span class="inline-block py-1 px-3 rounded-full text-xs font-bold tracking-wide uppercase border"
                :style="{ 
                  backgroundColor: 'var(--badge-bg)',
                  borderColor: 'var(--badge-border)',
                  color: 'var(--accent)'
                }">
            Beta Access
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import router from '@/router';
import { useUserStore } from '@/stores/userStore';
import { ArrowLeftIcon, SunIcon, MoonIcon, UserCircleIcon } from '@heroicons/vue/24/outline';

const userStore = useUserStore();

const isDark = ref(localStorage.getItem('darkMode') !== 'false');

const updateDOM = () => {
  if (isDark.value) {
    document.documentElement.classList.add('dark');
  } else {
    document.documentElement.classList.remove('dark');
  }
  localStorage.setItem('darkMode', isDark.value.toString());
};

const toggleDarkMode = () => {
  isDark.value = !isDark.value;
  updateDOM();
};

onMounted(() => {
  if (!localStorage.getItem('darkMode')) {
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
    isDark.value = prefersDark;
  }
  updateDOM();
});

const signInWithGoogle = () => {
  userStore.login('google');
};
</script>