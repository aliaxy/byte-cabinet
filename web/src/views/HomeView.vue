<script setup lang="ts">
import { ref } from 'vue'

// Mock data for demonstration - will be replaced with API calls
const featuredPost = ref({
  id: 1,
  title: 'Building a Modern Blog with Go Fiber and Vue.js',
  summary: 'Learn how to create a full-stack blog application using Go Fiber for the backend and Vue.js for the frontend. We will cover authentication, database design, and deployment strategies.',
  category: 'Tutorial',
  date: '2025-01-10',
  readTime: '12 min read',
  tags: ['Go', 'Vue.js', 'Tutorial']
})

const recentPosts = ref([
  {
    id: 2,
    title: 'Understanding JWT Authentication',
    summary: 'A deep dive into JSON Web Tokens and how to implement secure authentication.',
    category: 'Security',
    date: '2025-01-08',
    readTime: '8 min read'
  },
  {
    id: 3,
    title: 'SQLite vs PostgreSQL: Choosing the Right Database',
    summary: 'Comparing two popular databases for your next project.',
    category: 'Database',
    date: '2025-01-05',
    readTime: '6 min read'
  },
  {
    id: 4,
    title: 'CSS Grid Layout Mastery',
    summary: 'Master the art of creating complex layouts with CSS Grid.',
    category: 'Frontend',
    date: '2025-01-03',
    readTime: '10 min read'
  }
])

const categories = ref([
  { name: 'Tutorial', count: 12, icon: '📚' },
  { name: 'Backend', count: 8, icon: '⚙️' },
  { name: 'Frontend', count: 15, icon: '🎨' },
  { name: 'DevOps', count: 5, icon: '🚀' }
])

const popularTags = ref([
  'Go', 'Vue.js', 'TypeScript', 'Docker', 'PostgreSQL', 'REST API', 'JWT', 'Tailwind CSS'
])

const stats = ref({
  posts: 42,
  categories: 8,
  tags: 24
})
</script>

<template>
  <div class="min-h-screen">
    <!-- Hero Section -->
    <section class="py-12 md:py-20">
      <div class="max-w-6xl mx-auto px-4">
        <div class="text-center mb-12">
          <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold mb-6 animate-fade-in">
            Welcome to
            <span class="text-gradient">Byte Cabinet</span>
          </h1>
          <p class="text-lg md:text-xl text-[var(--color-text-secondary)] max-w-2xl mx-auto animate-slide-up">
            Personal technical notes and learning experiences.
            Exploring the world of software development, one byte at a time.
          </p>
        </div>

        <!-- Bento Grid -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6">

          <!-- Featured Post - Large Card -->
          <article class="bento-card-featured lg:col-span-2 lg:row-span-2 flex flex-col">
            <div class="flex items-center gap-2 mb-4">
              <span class="tag">{{ featuredPost.category }}</span>
              <span class="text-sm text-[var(--color-text-secondary)]">Featured</span>
            </div>
            <h2 class="text-2xl md:text-3xl font-bold mb-4 flex-grow">
              <a href="#" class="hover:text-primary-500 transition-colors">
                {{ featuredPost.title }}
              </a>
            </h2>
            <p class="text-[var(--color-text-secondary)] mb-6 line-clamp-3">
              {{ featuredPost.summary }}
            </p>
            <div class="flex items-center justify-between mt-auto">
              <div class="flex items-center gap-4 text-sm text-[var(--color-text-secondary)]">
                <span>{{ featuredPost.date }}</span>
                <span>{{ featuredPost.readTime }}</span>
              </div>
              <a href="#"
                class="text-primary-500 hover:text-primary-600 font-medium text-sm flex items-center gap-1 group">
                Read more
                <svg class="w-4 h-4 transform group-hover:translate-x-1 transition-transform" fill="none"
                  viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </a>
            </div>
          </article>

          <!-- Stats Card -->
          <div class="bento-card-highlight flex flex-col justify-center">
            <div class="text-center">
              <div class="text-4xl font-bold mb-2">{{ stats.posts }}</div>
              <div class="text-white/80 text-sm">Articles Published</div>
            </div>
            <div class="mt-6 pt-6 border-t border-white/20 grid grid-cols-2 gap-4 text-center">
              <div>
                <div class="text-xl font-bold">{{ stats.categories }}</div>
                <div class="text-white/60 text-xs">Categories</div>
              </div>
              <div>
                <div class="text-xl font-bold">{{ stats.tags }}</div>
                <div class="text-white/60 text-xs">Tags</div>
              </div>
            </div>
          </div>

          <!-- Categories Card -->
          <div class="bento-card">
            <h3 class="font-semibold mb-4 flex items-center gap-2">
              <span class="text-lg">📂</span>
              Categories
            </h3>
            <div class="space-y-3">
              <a v-for="cat in categories" :key="cat.name" href="#"
                class="flex items-center justify-between p-2 rounded-lg hover:bg-[var(--color-bg)] transition-colors group">
                <span class="flex items-center gap-2">
                  <span>{{ cat.icon }}</span>
                  <span class="group-hover:text-primary-500 transition-colors">{{ cat.name }}</span>
                </span>
                <span class="text-xs px-2 py-0.5 rounded-full bg-[var(--color-bg)] text-[var(--color-text-secondary)]">
                  {{ cat.count }}
                </span>
              </a>
            </div>
          </div>

          <!-- Recent Posts -->
          <article v-for="post in recentPosts.slice(0, 2)" :key="post.id" class="bento-card flex flex-col">
            <span class="tag mb-3 self-start">{{ post.category }}</span>
            <h3 class="font-semibold mb-2 line-clamp-2 flex-grow">
              <a href="#" class="hover:text-primary-500 transition-colors">
                {{ post.title }}
              </a>
            </h3>
            <div class="flex items-center gap-3 text-xs text-[var(--color-text-secondary)] mt-auto">
              <span>{{ post.date }}</span>
              <span>{{ post.readTime }}</span>
            </div>
          </article>

          <!-- Tags Cloud Card -->
          <div class="bento-card md:col-span-2">
            <h3 class="font-semibold mb-4 flex items-center gap-2">
              <span class="text-lg">🏷️</span>
              Popular Tags
            </h3>
            <div class="flex flex-wrap gap-2">
              <a v-for="tag in popularTags" :key="tag" href="#" class="tag">
                {{ tag }}
              </a>
            </div>
          </div>

          <!-- About Card -->
          <div class="bento-card flex flex-col items-center text-center">
            <div
              class="w-16 h-16 rounded-full bg-gradient-to-br from-primary-400 to-primary-600 flex items-center justify-center text-2xl mb-4 shadow-lg shadow-primary-500/25">
              👨‍💻
            </div>
            <h3 class="font-semibold mb-2">About Me</h3>
            <p class="text-sm text-[var(--color-text-secondary)] mb-4">
              Software developer passionate about building things.
            </p>
            <a href="/about" class="btn-secondary text-sm">
              Learn More
            </a>
          </div>

          <!-- Latest Post Card -->
          <article class="bento-card flex flex-col">
            <span class="tag mb-3 self-start">{{ recentPosts[2]?.category }}</span>
            <h3 class="font-semibold mb-2 line-clamp-2 flex-grow">
              <a href="#" class="hover:text-primary-500 transition-colors">
                {{ recentPosts[2]?.title }}
              </a>
            </h3>
            <p class="text-sm text-[var(--color-text-secondary)] mb-3 line-clamp-2">
              {{ recentPosts[2]?.summary }}
            </p>
            <div class="flex items-center gap-3 text-xs text-[var(--color-text-secondary)] mt-auto">
              <span>{{ recentPosts[2]?.date }}</span>
              <span>{{ recentPosts[2]?.readTime }}</span>
            </div>
          </article>

        </div>

        <!-- View All Posts Button -->
        <div class="text-center mt-12">
          <a href="/archives" class="btn-primary inline-flex items-center gap-2">
            View All Posts
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </a>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
