import { ref, onMounted } from 'vue'

export function useDarkMode() {
  const isDark = ref(localStorage.getItem('darkMode') !== 'false')

  const updateDOM = () => {
    if (isDark.value) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
    localStorage.setItem('darkMode', isDark.value.toString())
  }

  const toggleDarkMode = () => {
    isDark.value = !isDark.value
    updateDOM()
  }

  onMounted(() => {
    if (!localStorage.getItem('darkMode')) {
      const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
      isDark.value = prefersDark
    }
    updateDOM()
  })

  return {
    isDark,
    toggleDarkMode
  }
}