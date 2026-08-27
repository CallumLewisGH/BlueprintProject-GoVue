import type { Config } from 'tailwindcss'

export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'selector',
  theme: {
    extend: {
      colors: {
        brand: {
          primary: '#7c3aed',
          dark: '#1e1b2e',
          light: '#f3f1f9',
        }
      }
    },
  },
  plugins: [],
} satisfies Config