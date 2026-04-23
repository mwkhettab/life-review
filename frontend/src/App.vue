<template>
  <div class="min-h-screen text-white font-space">
    <Header />
    <router-view v-slot="{ Component }">
      <transition name="page" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>
    <Footer />
  </div>
</template>

<script setup>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
import { initStars } from "./utils/stars";
import { onMounted } from "vue";

onMounted(() => {
  initStars();
});
</script>

<style>
.page-enter-active {
  transition:
    opacity 0.25s ease,
    transform 0.25s ease;
}
.page-leave-active {
  transition:
    opacity 0.15s ease,
    transform 0.15s ease;
}
.page-enter-from {
  opacity: 0;
  transform: translateY(8px);
}
.page-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}
</style>
