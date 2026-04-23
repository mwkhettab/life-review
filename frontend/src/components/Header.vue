<template>
  <header class="fixed top-0 left-0 right-0 z-40">
    <div class="max-w-5xl mx-auto px-6 h-16 flex items-center justify-between">
      <!-- Logo -->
      <router-link to="/" class="flex items-center gap-2.5 group">
        <span
          class="text-sm font-semibold tracking-tight text-white/50 group-hover:text-white/80 transition-colors"
        >
          life review
        </span>
      </router-link>

      <!-- GitHub -->
      <a
        href="https://github.com/mwkhettab/life-review"
        target="_blank"
        rel="noopener noreferrer"
        class="w-8 h-8 flex items-center justify-center rounded-lg text-white/25 hover:text-white/60 transition-colors duration-200"
      >
        <Github />
      </a>
    </div>
  </header>
</template>

<script setup>
import { Sparkles, Github } from "lucide-vue-next";
</script>
