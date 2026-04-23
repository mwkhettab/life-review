<template>
  <div class="w-full min-h-screen px-6 pt-16 pb-16 text-white font-space">
    <div class="grain-overlay" />

    <div class="max-w-3xl mx-auto space-y-8">
      <!-- Header -->
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold tracking-tight">Admin Console</h1>
          <p class="text-white/30 text-xs mt-0.5 uppercase tracking-widest">
            Life Review · Review Queue
          </p>
        </div>
        <button
          v-if="authed"
          @click="logout"
          class="text-xs text-white/25 hover:text-white/50 transition border border-white/8 px-3 py-1.5 rounded-lg"
        >
          Sign out
        </button>
      </div>

      <!-- ── Auth gate ── -->
      <div
        v-if="!authed"
        class="border border-white/10 rounded-2xl p-10 max-w-sm mx-auto space-y-5"
      >
        <div class="text-center space-y-1">
          <div
            class="w-10 h-10 rounded-xl bg-white/5 border border-white/10 flex items-center justify-center mx-auto mb-3"
          >
            <Lock :size="16" class="text-white/40" />
          </div>
          <p class="font-semibold text-sm">Enter admin key</p>
          <p class="text-white/30 text-xs">
            Stored in sessionStorage — gone when you close the tab.
          </p>
        </div>
        <div class="space-y-2">
          <input
            v-model="keyInput"
            type="password"
            placeholder="••••••••••••••••"
            @keydown.enter="attemptAuth"
            class="w-full bg-white/4 border rounded-xl px-4 py-3 text-sm placeholder-white/20 focus:outline-none transition"
            :class="
              authError
                ? 'border-red-400/50'
                : 'border-white/10 focus:border-white/30'
            "
          />
          <p v-if="authError" class="text-red-400/60 text-xs">
            {{ authError }}
          </p>
        </div>
        <button
          @click="attemptAuth"
          :disabled="!keyInput.trim()"
          class="w-full py-2.5 rounded-xl text-sm font-semibold transition"
          :class="
            keyInput.trim()
              ? 'bg-white text-black hover:bg-white/90'
              : 'bg-white/8 text-white/20 cursor-not-allowed'
          "
        >
          Unlock
        </button>
      </div>

      <!-- ── Dashboard ── -->
      <template v-else>
        <!-- Stats row -->
        <div class="grid grid-cols-3 gap-3">
          <div class="border border-white/8 rounded-xl p-4 space-y-1">
            <p class="text-white/30 text-xs uppercase tracking-widest">Total</p>
            <p class="text-3xl font-bold">{{ reviews.length }}</p>
          </div>
          <div
            class="border border-yellow-400/15 rounded-xl p-4 space-y-1 bg-yellow-400/3"
          >
            <p class="text-yellow-400/50 text-xs uppercase tracking-widest">
              Pending
            </p>
            <p class="text-3xl font-bold text-yellow-400/80">
              {{ pendingCount }}
            </p>
          </div>
          <div
            class="border border-green-400/15 rounded-xl p-4 space-y-1 bg-green-400/3"
          >
            <p class="text-green-400/50 text-xs uppercase tracking-widest">
              Approved
            </p>
            <p class="text-3xl font-bold text-green-400/80">
              {{ approvedCount }}
            </p>
          </div>
        </div>

        <!-- Filter tabs -->
        <div class="flex gap-2">
          <button
            v-for="f in filters"
            :key="f.value"
            @click="activeFilter = f.value"
            class="text-xs px-3 py-1.5 rounded-lg border transition"
            :class="
              activeFilter === f.value
                ? 'border-white/30 text-white bg-white/8'
                : 'border-white/8 text-white/35 hover:border-white/15'
            "
          >
            {{ f.label }}
            <span class="ml-1.5 opacity-50">{{ f.count }}</span>
          </button>
        </div>

        <!-- Loading -->
        <template v-if="loadingInitial">
          <div
            v-for="i in 3"
            :key="i"
            class="border border-white/8 rounded-2xl p-5 space-y-3 animate-pulse"
          >
            <div class="h-4 bg-white/8 rounded w-1/3" />
            <div class="h-3 bg-white/5 rounded w-2/3" />
          </div>
        </template>

        <div
          v-else-if="fetchError"
          class="border border-red-400/20 rounded-2xl p-6 text-center text-red-400/60 text-sm"
        >
          {{ fetchError }}
        </div>

        <template v-else>
          <!-- Empty state -->
          <div
            v-if="filteredReviews.length === 0"
            class="border border-white/8 rounded-2xl p-10 text-center text-white/25 text-sm"
          >
            Nothing here.
          </div>

          <!-- Review cards -->
          <div
            v-for="review in filteredReviews"
            :key="review.id"
            class="border rounded-2xl p-5 space-y-3 transition"
            :class="
              review.approved
                ? 'border-white/8 bg-white/2'
                : 'border-yellow-400/15 bg-yellow-400/2'
            "
          >
            <!-- Top row -->
            <div class="flex items-start justify-between gap-4">
              <div class="space-y-0.5">
                <div class="flex items-center gap-2">
                  <p class="font-semibold text-sm">{{ review.name }}</p>
                  <span
                    class="text-[10px] px-1.5 py-0.5 rounded border"
                    :class="
                      review.approved
                        ? 'border-green-400/25 text-green-400/50'
                        : 'border-yellow-400/30 text-yellow-400/60'
                    "
                  >
                    {{ review.approved ? "approved" : "pending" }}
                  </span>
                </div>
                <div class="flex gap-0.5">
                  <span
                    v-for="i in 5"
                    :key="i"
                    class="text-xs"
                    :class="
                      i <= review.rating ? 'text-green-400' : 'text-white/15'
                    "
                    >★</span
                  >
                </div>
              </div>
              <p class="text-white/20 text-xs shrink-0">{{ review.date }}</p>
            </div>

            <!-- Body -->
            <p v-if="review.body" class="text-white/45 text-sm leading-relaxed">
              {{ review.body }}
            </p>

            <!-- Pros / Cons -->
            <div
              v-if="review.pros?.length || review.cons?.length"
              class="flex flex-wrap gap-1.5"
            >
              <span
                v-for="p in review.pros"
                :key="'p' + p"
                class="text-xs px-2 py-0.5 rounded-md bg-green-400/5 text-green-300/60 border border-green-400/10"
                >+ {{ p }}</span
              >
              <span
                v-for="c in review.cons"
                :key="'c' + c"
                class="text-xs px-2 py-0.5 rounded-md bg-red-400/5 text-red-300/60 border border-red-400/10"
                >− {{ c }}</span
              >
            </div>

            <!-- doItAgain -->
            <p
              v-if="review.doItAgain !== undefined"
              class="text-white/25 text-xs"
            >
              {{
                review.doItAgain
                  ? "🔁 Would do it again"
                  : "🚪 One life was enough"
              }}
            </p>

            <!-- Actions -->
            <div class="flex gap-2 pt-1">
              <button
                v-if="!review.approved"
                @click="handleApprove(review)"
                :disabled="actioning === review.id"
                class="flex items-center gap-1.5 text-xs px-3 py-1.5 rounded-lg border border-green-400/30 text-green-400/70 bg-green-400/5 hover:bg-green-400/10 hover:border-green-400/50 transition disabled:opacity-40 disabled:cursor-not-allowed"
              >
                <Check :size="11" />
                {{ actioning === review.id ? "Approving…" : "Approve" }}
              </button>
              <button
                @click="handleDelete(review)"
                :disabled="actioning === review.id"
                class="flex items-center gap-1.5 text-xs px-3 py-1.5 rounded-lg border border-red-400/20 text-red-400/50 bg-red-400/3 hover:bg-red-400/8 hover:border-red-400/40 transition disabled:opacity-40 disabled:cursor-not-allowed"
              >
                <Trash2 :size="11" />
                {{ actioning === review.id ? "Deleting…" : "Delete" }}
              </button>
            </div>
          </div>

          <!-- Load more -->
          <div v-if="hasMore" class="text-center pt-2">
            <button
              @click="loadMore"
              :disabled="loadingMore"
              class="px-6 py-2.5 rounded-xl border border-white/10 text-white/40 text-sm hover:border-white/20 hover:text-white/60 transition disabled:opacity-40"
            >
              {{ loadingMore ? "Loading..." : "Load more" }}
            </button>
          </div>
        </template>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { Lock, Check, Trash2 } from "lucide-vue-next";
import {
  fetchAllReviews,
  approveReview,
  deleteReview,
  type Review,
} from "../utils/api";

// ── Auth ──────────────────────────────────────────────────────────────────────
const SESSION_KEY = "admin_key";

const keyInput = ref("");
const authError = ref("");
const authed = ref(false);
const adminKey = ref("");

function attemptAuth() {
  const key = keyInput.value.trim();
  if (!key) return;
  // Optimistically store and try to fetch — if it 401s we'll know
  adminKey.value = key;
  sessionStorage.setItem(SESSION_KEY, key);
  authed.value = true;
  authError.value = "";
  loadReviews(true);
}

function logout() {
  sessionStorage.removeItem(SESSION_KEY);
  authed.value = false;
  adminKey.value = "";
  keyInput.value = "";
  reviews.value = [];
}

// ── Data ──────────────────────────────────────────────────────────────────────
const PAGE_SIZE = 20;
const reviews = ref<Review[]>([]);
const offset = ref(0);
const hasMore = ref(true);
const loadingInitial = ref(false);
const loadingMore = ref(false);
const fetchError = ref("");
const actioning = ref<number | null>(null);

async function loadReviews(reset = false) {
  if (reset) {
    loadingInitial.value = true;
    offset.value = 0;
    reviews.value = [];
  } else {
    loadingMore.value = true;
  }
  fetchError.value = "";

  try {
    const page = await fetchAllReviews(adminKey.value, PAGE_SIZE, offset.value);
    reviews.value = reset ? page : [...reviews.value, ...page];
    offset.value += page.length;
    hasMore.value = page.length === PAGE_SIZE;
  } catch (e) {
    if (
      e instanceof Error &&
      e.message.toLowerCase().includes("unauthorized")
    ) {
      authError.value = "Invalid admin key.";
      authed.value = false;
      sessionStorage.removeItem(SESSION_KEY);
    } else {
      fetchError.value =
        e instanceof Error ? e.message : "Failed to load reviews.";
    }
  } finally {
    loadingInitial.value = false;
    loadingMore.value = false;
  }
}

async function loadMore() {
  await loadReviews(false);
}

async function handleApprove(review: Review) {
  actioning.value = review.id;
  try {
    const updated = await approveReview(adminKey.value, review.id);
    const idx = reviews.value.findIndex((r) => r.id === review.id);
    if (idx !== -1) reviews.value[idx] = updated;
  } catch (e) {
    alert(e instanceof Error ? e.message : "Failed to approve.");
  } finally {
    actioning.value = null;
  }
}

async function handleDelete(review: Review) {
  if (!confirm(`Delete review from "${review.name}"?`)) return;
  actioning.value = review.id;
  try {
    await deleteReview(adminKey.value, review.id);
    reviews.value = reviews.value.filter((r) => r.id !== review.id);
  } catch (e) {
    alert(e instanceof Error ? e.message : "Failed to delete.");
  } finally {
    actioning.value = null;
  }
}

// ── Filters ───────────────────────────────────────────────────────────────────
const activeFilter = ref<"all" | "pending" | "approved">("pending");

const pendingCount = computed(
  () => reviews.value.filter((r) => !r.approved).length,
);
const approvedCount = computed(
  () => reviews.value.filter((r) => r.approved).length,
);

const filters = computed(() => [
  { value: "all" as const, label: "All", count: reviews.value.length },
  { value: "pending" as const, label: "Pending", count: pendingCount.value },
  { value: "approved" as const, label: "Approved", count: approvedCount.value },
]);

const filteredReviews = computed(() => {
  if (activeFilter.value === "pending")
    return reviews.value.filter((r) => !r.approved);
  if (activeFilter.value === "approved")
    return reviews.value.filter((r) => r.approved);
  return reviews.value;
});

// ── Init ──────────────────────────────────────────────────────────────────────
onMounted(() => {
  const saved = sessionStorage.getItem(SESSION_KEY);
  if (saved) {
    adminKey.value = saved;
    authed.value = true;
    loadReviews(true);
  }
});
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
