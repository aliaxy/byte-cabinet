<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppHeader from './components/layout/AppHeader.vue'
import AppFooter from './components/layout/AppFooter.vue'
</script>

<template>
  <div class="min-h-screen flex flex-col custom-scrollbar">
    <AppHeader />

    <main class="flex-grow">
      <RouterView />
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
/* Additional app-level styles if needed */
</style>
