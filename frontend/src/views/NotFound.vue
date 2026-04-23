<template>
  <div class="relative min-h-screen flex flex-col items-center justify-center px-6 overflow-hidden">

    <div class="grain-overlay" />

    <div class="relative z-10 max-w-md text-center space-y-6">

      <p class="text-8xl font-bold text-white/15 select-none">404</p>

      <div class="space-y-3 -mt-4">
        <h1 class="text-2xl font-semibold tracking-tight">
          This page wasn't in the patch notes.
        </h1>
        <p class="text-white/35 text-sm leading-relaxed">
          Either it was never built, or it was quietly removed
          in a hotfix nobody was told about. Classic.
        </p>
      </div>

      <router-link to="/"
        class="inline-flex items-center gap-2 px-6 py-3 bg-white text-black font-semibold
               rounded-xl hover:bg-white/90 transition-all duration-200 text-sm group">
        <ArrowLeft :size="15" class="group-hover:-translate-x-0.5 transition-transform" />
        Back to lobby
      </router-link>

    </div>
  </div>
</template>

<script setup>
import { ArrowLeft } from 'lucide-vue-next'
</script>

<style scoped>
.grain-overlay {
  position: fixed;
  inset: 0;
  pointer-events: none;
  z-index: 0;
  opacity: 0.03;
  background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 256 256' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noise'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noise)'/%3E%3C/svg%3E");
  background-repeat: repeat;
  background-size: 128px 128px;
}
</style>