<template>
  <div
    class="relative min-h-screen flex flex-col items-center justify-center px-6 overflow-hidden"
  >
    <!-- Subtle radial glow -->
    <div
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-150 h-150 bg-green-500/3 rounded-full blur-[120px] pointer-events-none"
    />

    <div class="grain-overlay" />

    <!-- Content -->
    <div class="relative z-10 max-w-lg text-center space-y-8">
      <!-- Headline -->
      <div class="space-y-4">
        <h1 class="text-5xl sm:text-6xl font-bold tracking-tight leading-[1.1]">
          Rate the<br />
          <span class="text-green-400/90">human experience.</span>
        </h1>

        <p
          class="text-white/35 text-base sm:text-lg leading-relaxed max-w-md mx-auto"
        >
          A review platform for life itself. No sign-up required.
        </p>
      </div>

      <!-- CTA -->
      <div
        class="flex flex-col sm:flex-row items-center justify-center gap-3 pt-2"
      >
        <router-link
          to="/reviews"
          class="px-8 py-3.5 bg-white text-black font-semibold rounded-xl hover:bg-white/90 transition-all duration-200 text-sm inline-flex items-center gap-2 group"
        >
          Browse Reviews
          <ArrowRight
            :size="15"
            class="group-hover:translate-x-0.5 transition-transform"
          />
        </router-link>

        <router-link
          to="/reviews"
          @click.prevent="goToReviewsAndOpen"
          class="px-8 py-3.5 border border-white/12 text-white/60 font-medium rounded-xl hover:border-white/25 hover:text-white/80 transition-all duration-200 text-sm inline-flex items-center gap-2"
        >
          <PenLine :size="14" />
          Leave a Review
        </router-link>
      </div>

      <!-- Stats teaser -->
      <div class="flex items-center justify-center gap-8 pt-6">
        <div class="text-center">
          <p class="text-2xl font-bold text-white/80">
            {{ stats?.averageRating }}
          </p>
          <p class="text-[10px] uppercase tracking-widest text-white/25 mt-0.5">
            avg rating
          </p>
        </div>
        <div class="w-px h-8 bg-white/8" />
        <div class="text-center">
          <p class="text-2xl font-bold text-white/80">
            {{ stats?.totalReviews }}
          </p>
          <p class="text-[10px] uppercase tracking-widest text-white/25 mt-0.5">
            reviews
          </p>
        </div>
        <div class="w-px h-8 bg-white/8" />
        <div class="text-center">
          <p class="text-2xl font-bold text-white/80">{{ stats?.ratingCounts?.['5'] }}</p>
          <p class="text-[10px] uppercase tracking-widest text-white/25 mt-0.5">
            5 star ratings
          </p>
        </div>
      </div>
    </div>

    <!-- Bottom fade text -->
    <div class="absolute bottom-8 text-center">
      <p class="text-white/15 text-xs tracking-widest uppercase">
        v1.0.0 · patch notes pending
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { ArrowRight, PenLine } from "lucide-vue-next";
import { useRouter } from "vue-router";
import { fetchStats } from "../utils/api";
import type { ReviewStats } from "../utils/api";

const loading = ref(true);

const stats = ref<ReviewStats | null>(null);

onMounted(async () => {
  const data = await fetchStats();
  stats.value = data;
  loading.value = false;
});

const router = useRouter();

function goToReviewsAndOpen() {
  router.push({ path: "/reviews", query: { action: "review" } });
}
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
