<script setup lang="ts">
const skills = [
  { name: 'Go', icon: '🐹', level: 90 },
  { name: 'Vue.js', icon: '💚', level: 85 },
  { name: 'TypeScript', icon: '📘', level: 80 },
  { name: 'Docker', icon: '🐳', level: 75 },
  { name: 'PostgreSQL', icon: '🐘', level: 70 },
  { name: 'Linux', icon: '🐧', level: 80 }
]

const timeline = [
  {
    year: '2025',
    title: 'Started Byte Cabinet',
    description: 'Created this blog to document my learning journey and share technical insights.'
  },
  {
    year: '2024',
    title: 'Diving into Go',
    description: 'Started exploring Go for building high-performance backend services.'
  },
  {
    year: '2023',
    title: 'Full Stack Development',
    description: 'Expanded skills to include frontend development with Vue.js and React.'
  }
]
</script>

<template>
  <div class="min-h-screen py-12">
    <div class="max-w-4xl mx-auto px-4">
      <!-- Header Section -->
      <div class="text-center mb-16 animate-fade-in">
        <div
          class="w-32 h-32 mx-auto mb-6 rounded-full bg-gradient-to-br from-primary-400 to-primary-600 flex items-center justify-center text-5xl shadow-2xl shadow-primary-500/30">
          👨‍💻
        </div>
        <h1 class="text-4xl md:text-5xl font-bold mb-4">
          About <span class="text-gradient">Me</span>
        </h1>
        <p class="text-xl text-[var(--color-text-secondary)] max-w-2xl mx-auto">
          Software developer passionate about building elegant solutions to complex problems.
        </p>
      </div>

      <!-- Bio Section -->
      <section class="bento-card mb-8 animate-slide-up">
        <h2 class="text-2xl font-bold mb-4 flex items-center gap-2">
          <span>📖</span> Bio
        </h2>
        <div class="prose-content space-y-4">
          <p>
            Hello! I'm a software developer who loves exploring new technologies and building things that matter.
            This blog, <strong>Byte Cabinet</strong>, is my personal space for documenting technical notes,
            sharing what I learn, and organizing my thoughts on various programming topics.
          </p>
          <p>
            I believe in continuous learning and the power of sharing knowledge. Whether it's a quick tip,
            an in-depth tutorial, or lessons learned from a challenging project, I hope you find something
            useful here.
          </p>
          <p>
            When I'm not coding, you'll find me reading tech blogs, experimenting with new frameworks,
            or contributing to open-source projects.
          </p>
        </div>
      </section>

      <!-- Skills Section -->
      <section class="bento-card mb-8">
        <h2 class="text-2xl font-bold mb-6 flex items-center gap-2">
          <span>🛠️</span> Skills
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div v-for="skill in skills" :key="skill.name" class="flex items-center gap-4">
            <span class="text-2xl w-10 text-center">{{ skill.icon }}</span>
            <div class="flex-grow">
              <div class="flex justify-between mb-1">
                <span class="font-medium">{{ skill.name }}</span>
                <span class="text-sm text-[var(--color-text-secondary)]">{{ skill.level }}%</span>
              </div>
              <div class="h-2 bg-[var(--color-bg)] rounded-full overflow-hidden">
                <div
                  class="h-full bg-gradient-to-r from-primary-500 to-primary-600 rounded-full transition-all duration-500"
                  :style="{ width: skill.level + '%' }">
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Timeline Section -->
      <section class="bento-card mb-8">
        <h2 class="text-2xl font-bold mb-6 flex items-center gap-2">
          <span>🗓️</span> Journey
        </h2>
        <div class="relative">
          <!-- Timeline line -->
          <div class="absolute left-4 top-0 bottom-0 w-0.5 bg-[var(--color-border)]"></div>

          <div v-for="(item, index) in timeline" :key="index" class="relative pl-12 pb-8 last:pb-0">
            <!-- Timeline dot -->
            <div
              class="absolute left-2 w-5 h-5 rounded-full bg-primary-500 border-4 border-[var(--color-bg-secondary)]">
            </div>

            <div class="bg-[var(--color-bg)] rounded-xl p-4 border border-[var(--color-border)]">
              <span class="text-sm text-primary-500 font-medium">{{ item.year }}</span>
              <h3 class="text-lg font-semibold mt-1">{{ item.title }}</h3>
              <p class="text-[var(--color-text-secondary)] mt-2 text-sm">{{ item.description }}</p>
            </div>
          </div>
        </div>
      </section>

      <!-- Contact Section -->
      <section class="bento-card-highlight text-center">
        <h2 class="text-2xl font-bold mb-4">Let's Connect!</h2>
        <p class="text-white/80 mb-6">
          Feel free to reach out if you want to collaborate or just say hi.
        </p>
        <div class="flex justify-center gap-4">
          <a href="https://github.com/aliaxy" target="_blank" rel="noopener noreferrer"
            class="p-3 rounded-xl bg-white/10 hover:bg-white/20 transition-colors" aria-label="GitHub">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
              <path
                d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
            </svg>
          </a>
          <a href="mailto:hello@example.com" class="p-3 rounded-xl bg-white/10 hover:bg-white/20 transition-colors"
            aria-label="Email">
            <svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
          </a>
        </div>
      </section>

      <!-- Blog Info -->
      <section class="mt-8 text-center text-sm text-[var(--color-text-secondary)]">
        <p>
          This blog is built with
          <a href="https://gofiber.io" target="_blank" class="text-primary-500 hover:underline">Go Fiber</a>
          and
          <a href="https://vuejs.org" target="_blank" class="text-primary-500 hover:underline">Vue.js</a>.
          The source code is available on
          <a href="https://github.com/aliaxy/byte-cabinet" target="_blank"
            class="text-primary-500 hover:underline">GitHub</a>.
        </p>
      </section>
    </div>
  </div>
</template>
