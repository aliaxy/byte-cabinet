<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()
const isDark = ref(false)
const isMobileMenuOpen = ref(false)

const navLinks = [
  { name: 'Home', path: '/' },
  { name: 'Archives', path: '/archives' },
  { name: 'Categories', path: '/categories' },
  { name: 'Tags', path: '/tags' },
  { name: 'About', path: '/about' },
]

const toggleTheme = () => {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

// Initialize theme from localStorage
if (typeof window !== 'undefined') {
  const savedTheme = localStorage.getItem('theme')
  const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  isDark.value = savedTheme === 'dark' || (!savedTheme && prefersDark)
  document.documentElement.classList.toggle('dark', isDark.value)
}
</script>

<template>
  <header class="sticky top-0 z-50 backdrop-blur-lg bg-[var(--color-bg)]/80 border-b border-[var(--color-border)]">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo -->
        <RouterLink to="/" class="flex items-center space-x-3 group">
          <div
            class="w-10 h-10 rounded-xl bg-gradient-to-br from-primary-500 to-primary-600 flex items-center justify-center shadow-lg shadow-primary-500/25 group-hover:shadow-primary-500/40 transition-shadow">
            <!-- Cabinet Icon -->
            <svg class="w-6 h-6" viewBox="0 0 24 24" fill="none">
              <rect x="4" y="4" width="16" height="16" rx="2" fill="#1e1b4b" />
              <rect x="6" y="6" width="12" height="3" rx="1" fill="#818cf8" />
              <rect x="9" y="7" width="6" height="1" rx="0.5" fill="#e0e7ff" />
              <rect x="6" y="10.5" width="12" height="3" rx="1" fill="#818cf8" />
              <rect x="9" y="11.5" width="6" height="1" rx="0.5" fill="#e0e7ff" />
              <rect x="6" y="15" width="12" height="3" rx="1" fill="#818cf8" />
              <rect x="9" y="16" width="6" height="1" rx="0.5" fill="#e0e7ff" />
            </svg>
          </div>
          <span class="text-xl font-bold hidden sm:block">
            Byte<span class="text-primary-500">Cabinet</span>
          </span>
        </RouterLink>

        <!-- Desktop Navigation -->
        <nav class="hidden md:flex items-center space-x-1">
          <RouterLink v-for="link in navLinks" :key="link.path" :to="link.path"
            class="px-4 py-2 rounded-lg text-sm font-medium transition-colors" :class="[
              route.path === link.path
                ? 'text-primary-600 dark:text-primary-400 bg-primary-500/10'
                : 'text-[var(--color-text-secondary)] hover:text-[var(--color-text)] hover:bg-[var(--color-bg-secondary)]'
            ]">
            {{ link.name }}
          </RouterLink>
        </nav>

        <!-- Right Actions -->
        <div class="flex items-center space-x-2">
          <!-- Theme Toggle -->
          <button @click="toggleTheme"
            class="p-2 rounded-lg text-[var(--color-text-secondary)] hover:text-[var(--color-text)] hover:bg-[var(--color-bg-secondary)] transition-colors"
            :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'">
            <!-- Sun Icon -->
            <svg v-if="isDark" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            <!-- Moon Icon -->
            <svg v-else class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
            </svg>
          </button>

          <!-- GitHub Link -->
          <a href="https://github.com/aliaxy/byte-cabinet" target="_blank" rel="noopener noreferrer"
            class="p-2 rounded-lg text-[var(--color-text-secondary)] hover:text-[var(--color-text)] hover:bg-[var(--color-bg-secondary)] transition-colors"
            aria-label="GitHub">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
              <path fill-rule="evenodd" clip-rule="evenodd"
                d="M12 2C6.477 2 2 6.477 2 12c0 4.42 2.865 8.17 6.839 9.49.5.092.682-.217.682-.482 0-.237-.008-.866-.013-1.7-2.782.604-3.369-1.34-3.369-1.34-.454-1.156-1.11-1.464-1.11-1.464-.908-.62.069-.608.069-.608 1.003.07 1.531 1.03 1.531 1.03.892 1.529 2.341 1.087 2.91.831.092-.646.35-1.086.636-1.336-2.22-.253-4.555-1.11-4.555-4.943 0-1.091.39-1.984 1.029-2.683-.103-.253-.446-1.27.098-2.647 0 0 .84-.269 2.75 1.025A9.578 9.578 0 0112 6.836c.85.004 1.705.115 2.504.337 1.909-1.294 2.747-1.025 2.747-1.025.546 1.377.203 2.394.1 2.647.64.699 1.028 1.592 1.028 2.683 0 3.842-2.339 4.687-4.566 4.935.359.309.678.919.678 1.852 0 1.336-.012 2.415-.012 2.743 0 .267.18.578.688.48C19.138 20.167 22 16.418 22 12c0-5.523-4.477-10-10-10z" />
            </svg>
          </a>

          <!-- Mobile Menu Button -->
          <button @click="toggleMobileMenu"
            class="md:hidden p-2 rounded-lg text-[var(--color-text-secondary)] hover:text-[var(--color-text)] hover:bg-[var(--color-bg-secondary)] transition-colors"
            aria-label="Toggle menu">
            <svg v-if="!isMobileMenuOpen" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
            <svg v-else class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Mobile Navigation -->
      <nav v-if="isMobileMenuOpen" class="md:hidden py-4 border-t border-[var(--color-border)]">
        <div class="flex flex-col space-y-1">
          <RouterLink v-for="link in navLinks" :key="link.path" :to="link.path" @click="isMobileMenuOpen = false"
            class="px-4 py-2 rounded-lg text-sm font-medium transition-colors" :class="[
              route.path === link.path
                ? 'text-primary-600 dark:text-primary-400 bg-primary-500/10'
                : 'text-[var(--color-text-secondary)] hover:text-[var(--color-text)] hover:bg-[var(--color-bg-secondary)]'
            ]">
            {{ link.name }}
          </RouterLink>
        </div>
      </nav>
    </div>
  </header>
</template>
