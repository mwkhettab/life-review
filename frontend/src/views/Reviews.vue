<template>
  <div class="w-full min-h-screen px-6 pt-24 pb-16 text-white font-space">
    <div class="grain-overlay" />

    <div class="max-w-2xl mx-auto space-y-10">
      <!-- Header -->
      <div class="text-center space-y-3">
        <h1 class="text-4xl sm:text-5xl font-bold tracking-tight">
          Life Review
        </h1>
        <p class="text-white/40 text-sm uppercase tracking-widest">
          Developer feedback portal · {{ totalCount }} reviews filed
        </p>
        <p class="text-white/25 text-xs italic">{{ quote }}</p>
      </div>

      <!-- Aggregate Score -->
      <div
        class="border border-white/12 rounded-2xl p-8 flex flex-col sm:flex-row items-center justify-between gap-6 bg-white/1"
      >
        <div class="text-center sm:text-left space-y-1">
          <p class="text-white/40 text-xs uppercase tracking-widest">
            Overall Rating
          </p>
          <div class="flex items-end gap-3">
            <span class="text-6xl font-bold">{{ averageRating }}</span>
            <span class="text-white/30 text-lg mb-2">/ 5</span>
          </div>
          <div class="flex gap-1">
            <span v-for="i in 5" :key="i" class="relative text-xl inline-block">
              <span class="text-white/15">★</span>
              <span
                class="absolute inset-0 text-green-400 overflow-hidden"
                :style="{ width: starFillPercent(i, averageRatingNum) + '%' }"
                >★</span
              >
            </span>
          </div>
        </div>

        <div class="w-px h-16 bg-white/8 hidden sm:block" />

        <div class="text-center space-y-1">
          <p class="text-white/40 text-xs uppercase tracking-widest">
            Total Reviews
          </p>
          <p class="text-5xl font-bold">{{ totalCount }}</p>
          <p class="text-white/25 text-xs">
            from {{ totalCount }} conscious entities
          </p>
        </div>

        <div class="w-px h-16 bg-white/8 hidden sm:block" />

        <div class="space-y-1.5 text-sm w-36">
          <div
            v-for="star in [5, 4, 3, 2, 1]"
            :key="star"
            class="flex items-center gap-2"
          >
            <span class="text-white/35 w-2">{{ star }}</span>
            <span class="text-white/25 text-xs">★</span>
            <div class="flex-1 h-1.5 bg-white/8 rounded-full overflow-hidden">
              <div
                class="h-full bg-green-400 rounded-full transition-all duration-700"
                :style="{ width: starPercent(star) + '%' }"
              />
            </div>
            <span class="text-white/25 text-xs w-4 text-right">{{
              starCount(star)
            }}</span>
          </div>
        </div>
      </div>

      <!-- CTA -->
      <div class="text-center">
        <button
          @click="openModal"
          class="px-8 py-3 bg-white text-black font-semibold rounded-xl hover:bg-white/90 transition text-sm inline-flex items-center gap-2"
        >
          <PenLine :size="14" />
          Leave Your Review
        </button>
        <p class="text-white/25 text-xs mt-3">
          No sign-up required. Just honesty.
        </p>
      </div>

      <!-- Reviews List -->
      <div class="space-y-4">
        <div class="flex items-center justify-between">
          <h2 class="text-xl font-semibold">What Others Are Saying</h2>
          <span class="text-white/25 text-xs">{{ totalCount }} reviews</span>
        </div>

        <!-- Loading skeleton -->
        <template v-if="loadingInitial">
          <div
            v-for="i in 3"
            :key="i"
            class="border border-white/8 rounded-2xl p-6 space-y-3 animate-pulse"
          >
            <div class="h-4 bg-white/8 rounded w-1/4" />
            <div class="h-3 bg-white/5 rounded w-3/4" />
            <div class="h-3 bg-white/5 rounded w-1/2" />
          </div>
        </template>

        <!-- Error state -->
        <div
          v-else-if="fetchError"
          class="border border-red-400/20 rounded-2xl p-6 text-center text-red-400/60 text-sm"
        >
          {{ fetchError }}
        </div>

        <template v-else>
          <div
            v-for="review in reviews"
            :key="review.id"
            class="border border-white/8 rounded-2xl p-6 space-y-4 hover:border-white/15 transition bg-white/10"
          >
            <div class="flex items-start justify-between gap-4">
              <div class="space-y-1">
                <p class="font-semibold text-sm">{{ review.name }}</p>
                <div class="flex gap-0.5">
                  <span
                    v-for="i in 5"
                    :key="i"
                    class="text-sm"
                    :class="
                      i <= review.rating ? 'text-green-400' : 'text-white/15'
                    "
                    >★</span
                  >
                </div>
              </div>
              <div class="text-right h-min shrink-0">
                <span
                  v-if="review.doItAgain !== undefined"
                  class="text-xs px-2 py-0.5 rounded-md border"
                  :class="
                    review.doItAgain
                      ? 'border-green-400/30 text-green-400/70'
                      : 'border-red-400/30 text-red-400/70'
                  "
                >
                  {{
                    review.doItAgain
                      ? "🔁 Would do it again"
                      : "🚪 One life was enough"
                  }}
                </span>
                <p class="text-white/20 text-xs mt-1">{{ review.date }}</p>
              </div>
            </div>

            <p v-if="review.body" class="text-white/50 text-sm leading-relaxed">
              {{ review.body }}
            </p>

            <div
              v-if="review.pros?.length || review.cons?.length"
              class="grid grid-cols-2 gap-3 pt-1"
            >
              <div v-if="review.pros?.length" class="space-y-1.5">
                <p class="text-green-400/50 text-xs uppercase tracking-widest">
                  Highlights
                </p>
                <div class="flex flex-wrap gap-1.5">
                  <span
                    v-for="p in review.pros"
                    :key="p"
                    class="text-xs px-2 py-0.5 rounded-md bg-green-400/5 text-green-300/60 border border-green-400/10"
                  >
                    + {{ p }}
                  </span>
                </div>
              </div>
              <div v-if="review.cons?.length" class="space-y-1.5">
                <p class="text-red-400/50 text-xs uppercase tracking-widest">
                  Pain Points
                </p>
                <div class="flex flex-wrap gap-1.5">
                  <span
                    v-for="c in review.cons"
                    :key="c"
                    class="text-xs px-2 py-0.5 rounded-md bg-red-400/5 text-red-300/60 border border-red-400/10"
                  >
                    − {{ c }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Load more -->
          <div v-if="hasMore" class="text-center pt-2">
            <button
              @click="loadMore"
              :disabled="loadingMore"
              class="px-6 py-2.5 rounded-xl border border-white/10 text-white/40 text-sm hover:border-white/20 hover:text-white/60 transition disabled:opacity-40"
            >
              <span v-if="loadingMore">Loading...</span>
              <span v-else>Load more reviews</span>
            </button>
          </div>

          <p
            v-if="!hasMore && reviews.length > 0"
            class="text-center text-white/15 text-xs pt-2"
          >
            You've reached the end. That's all of them.
          </p>
        </template>
      </div>
    </div>

    <!-- ── Modal ── -->
    <Transition name="modal">
      <div
        v-if="modalOpen"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
      >
        <div
          class="absolute inset-0 bg-black/80 backdrop-blur-sm"
          @click="closeModal"
        />

        <div
          class="relative z-10 w-full max-w-lg bg-[#0d0d0d] border border-white/15 rounded-2xl max-h-[90vh] overflow-y-auto shadow-2xl"
        >
          <!-- Submitted state -->
          <div v-if="submitted" class="p-10 text-center space-y-4">
            <div
              class="w-14 h-14 rounded-2xl bg-green-400/10 border border-green-400/20 flex items-center justify-center mx-auto"
            >
              <CheckCircle :size="24" class="text-green-400" />
            </div>
            <p class="text-green-400 font-semibold text-lg">
              Review submitted.
            </p>
            <p class="text-white/45 text-sm leading-relaxed max-w-xs mx-auto">
              Your feedback is currently
              <span class="text-white/65 italic">under review</span> by the
              developers. Processing times may vary. Thank you for your
              patience, and for existing.
            </p>
            <div class="pt-2 flex flex-col gap-2 items-center">
              <button
                @click="resetForm"
                class="text-xs text-white/30 hover:text-white/60 transition underline underline-offset-4"
              >
                Submit another review
              </button>
              <button
                @click="closeModal"
                class="text-xs text-white/20 hover:text-white/40 transition"
              >
                Close
              </button>
            </div>
          </div>

          <!-- Form -->
          <div v-else class="p-8 space-y-6">
            <div class="flex items-start justify-between">
              <div>
                <h2 class="text-xl font-semibold">Leave Your Review</h2>
                <p class="text-white/30 text-xs mt-0.5">
                  Be honest. They can take it. Probably.
                </p>
              </div>
              <button
                @click="closeModal"
                class="text-white/30 hover:text-white/60 transition text-xl leading-none p-1"
              >
                <X :size="18" />
              </button>
            </div>

            <!-- Name -->
            <div class="space-y-1.5">
              <div class="flex items-center justify-between">
                <label class="text-xs uppercase tracking-widest text-white/40">
                  Your Name
                  <span class="text-red-400/70 normal-case tracking-normal"
                    >*</span
                  >
                </label>
                <span
                  class="text-[10px] tabular-nums"
                  :class="
                    form.name.length >= NAME_MAX
                      ? 'text-red-400/70'
                      : 'text-white/20'
                  "
                >
                  {{ form.name.length }}/{{ NAME_MAX }}
                </span>
              </div>
              <input
                v-model="form.name"
                type="text"
                :maxlength="NAME_MAX"
                placeholder="At least give us a name"
                class="w-full bg-white/4 border rounded-xl px-4 py-3 text-sm placeholder-white/20 focus:outline-none transition"
                :class="
                  nameError
                    ? 'border-red-400/50 focus:border-red-400/70'
                    : 'border-white/10 focus:border-green-400/50'
                "
                @blur="nameError = form.name.trim() === ''"
                @input="nameError = false"
              />
              <p v-if="nameError" class="text-red-400/70 text-xs">
                A name is required. Anonymous Entity is taken.
              </p>
            </div>

            <!-- Star Rating -->
            <div class="space-y-1.5">
              <label class="text-xs uppercase tracking-widest text-white/40">
                Overall Rating
                <span class="text-red-400/70 normal-case tracking-normal"
                  >*</span
                >
              </label>
              <div class="flex gap-2">
                <button
                  v-for="i in 5"
                  :key="i"
                  @click="form.rating = i"
                  @mouseover="hoverRating = i"
                  @mouseleave="hoverRating = 0"
                  class="text-3xl transition-transform hover:scale-110"
                  :class="
                    i <= (hoverRating || form.rating)
                      ? 'text-green-400'
                      : 'text-white/15'
                  "
                >
                  ★
                </button>
              </div>
              <p class="text-white/30 text-xs h-4">{{ ratingLabel }}</p>
            </div>

            <!-- Highlights / Pros -->
            <div class="space-y-2">
              <div class="flex items-center justify-between">
                <label
                  class="text-xs uppercase tracking-widest text-green-400/60"
                >
                  Highlights
                </label>
                <span
                  class="text-[10px] tabular-nums"
                  :class="
                    form.pros.length >= TAG_LIMIT
                      ? 'text-green-400/70'
                      : 'text-white/20'
                  "
                >
                  {{ form.pros.length }}/{{ TAG_LIMIT }}
                </span>
              </div>
              <p class="text-white/25 text-xs">
                The parts that made it worth logging in.
              </p>
              <div class="flex flex-wrap gap-1.5">
                <button
                  v-for="s in proSuggestions"
                  :key="s"
                  @click="toggleOrAdd(form.pros, s)"
                  :disabled="
                    !form.pros.includes(s) && form.pros.length >= TAG_LIMIT
                  "
                  class="text-xs px-2.5 py-1 rounded-lg border transition"
                  :class="
                    form.pros.includes(s)
                      ? 'border-green-400/50 bg-green-400/10 text-green-300'
                      : form.pros.length >= TAG_LIMIT
                        ? 'border-white/5 text-white/15 cursor-not-allowed'
                        : 'border-white/8 text-white/30 hover:border-green-400/30 hover:text-white/50'
                  "
                >
                  + {{ s }}
                </button>
              </div>
              <div class="flex gap-2">
                <div class="flex-1 relative">
                  <input
                    v-model="proInput"
                    @keydown.enter.prevent="addCustom(form.pros, 'pro')"
                    type="text"
                    :maxlength="TAG_MAX"
                    :disabled="form.pros.length >= TAG_LIMIT"
                    :placeholder="
                      form.pros.length >= TAG_LIMIT
                        ? 'Limit reached'
                        : 'Type your own and press ↵'
                    "
                    class="w-full bg-white/4 border border-white/8 rounded-xl px-3 py-2 text-xs placeholder-white/20 focus:outline-none focus:border-green-400/40 transition pr-12 disabled:opacity-30 disabled:cursor-not-allowed"
                  />
                  <span
                    v-if="form.pros.length < TAG_LIMIT"
                    class="absolute right-3 top-1/2 -translate-y-1/2 text-[10px] tabular-nums"
                    :class="
                      proInput.length >= TAG_MAX
                        ? 'text-red-400/70'
                        : 'text-white/15'
                    "
                  >
                    {{ proInput.length }}/{{ TAG_MAX }}
                  </span>
                </div>
                <button
                  @click="addCustom(form.pros, 'pro')"
                  :disabled="form.pros.length >= TAG_LIMIT"
                  class="px-3 py-2 rounded-xl border border-white/8 text-white/40 hover:border-green-400/40 hover:text-green-300 transition text-xs disabled:opacity-30 disabled:cursor-not-allowed"
                >
                  ↵
                </button>
              </div>
              <div v-if="form.pros.length" class="flex flex-wrap gap-1.5">
                <span
                  v-for="p in form.pros"
                  :key="p"
                  class="text-xs px-2.5 py-1 rounded-lg bg-green-400/10 text-green-300 border border-green-400/30 flex items-center gap-1.5"
                >
                  {{ p }}
                  <button
                    @click="remove(form.pros, p)"
                    class="opacity-50 hover:opacity-100 transition leading-none"
                  >
                    <X :size="10" />
                  </button>
                </span>
              </div>
            </div>

            <!-- Pain Points / Cons -->
            <div class="space-y-2">
              <div class="flex items-center justify-between">
                <label
                  class="text-xs uppercase tracking-widest text-red-400/60"
                >
                  Pain Points
                </label>
                <span
                  class="text-[10px] tabular-nums"
                  :class="
                    form.cons.length >= TAG_LIMIT
                      ? 'text-red-400/70'
                      : 'text-white/20'
                  "
                >
                  {{ form.cons.length }}/{{ TAG_LIMIT }}
                </span>
              </div>
              <p class="text-white/25 text-xs">
                Don't hold back. This is a safe space. File a bug report.
              </p>
              <div class="flex flex-wrap gap-1.5">
                <button
                  v-for="s in conSuggestions"
                  :key="s"
                  @click="toggleOrAdd(form.cons, s)"
                  :disabled="
                    !form.cons.includes(s) && form.cons.length >= TAG_LIMIT
                  "
                  class="text-xs px-2.5 py-1 rounded-lg border transition"
                  :class="
                    form.cons.includes(s)
                      ? 'border-red-400/50 bg-red-400/10 text-red-300'
                      : form.cons.length >= TAG_LIMIT
                        ? 'border-white/5 text-white/15 cursor-not-allowed'
                        : 'border-white/8 text-white/30 hover:border-red-400/30 hover:text-white/50'
                  "
                >
                  − {{ s }}
                </button>
              </div>
              <div class="flex gap-2">
                <div class="flex-1 relative">
                  <input
                    v-model="conInput"
                    @keydown.enter.prevent="addCustom(form.cons, 'con')"
                    type="text"
                    :maxlength="TAG_MAX"
                    :disabled="form.cons.length >= TAG_LIMIT"
                    :placeholder="
                      form.cons.length >= TAG_LIMIT
                        ? 'Limit reached'
                        : 'Type your own and press ↵'
                    "
                    class="w-full bg-white/4 border border-white/8 rounded-xl px-3 py-2 text-xs placeholder-white/20 focus:outline-none focus:border-red-400/40 transition pr-12 disabled:opacity-30 disabled:cursor-not-allowed"
                  />
                  <span
                    v-if="form.cons.length < TAG_LIMIT"
                    class="absolute right-3 top-1/2 -translate-y-1/2 text-[10px] tabular-nums"
                    :class="
                      conInput.length >= TAG_MAX
                        ? 'text-red-400/70'
                        : 'text-white/15'
                    "
                  >
                    {{ conInput.length }}/{{ TAG_MAX }}
                  </span>
                </div>
                <button
                  @click="addCustom(form.cons, 'con')"
                  :disabled="form.cons.length >= TAG_LIMIT"
                  class="px-3 py-2 rounded-xl border border-white/8 text-white/40 hover:border-red-400/40 hover:text-red-300 transition text-xs disabled:opacity-30 disabled:cursor-not-allowed"
                >
                  ↵
                </button>
              </div>
              <div v-if="form.cons.length" class="flex flex-wrap gap-1.5">
                <span
                  v-for="c in form.cons"
                  :key="c"
                  class="text-xs px-2.5 py-1 rounded-lg bg-red-400/10 text-red-300 border border-red-400/30 flex items-center gap-1.5"
                >
                  {{ c }}
                  <button
                    @click="remove(form.cons, c)"
                    class="opacity-50 hover:opacity-100 transition leading-none"
                  >
                    <X :size="10" />
                  </button>
                </span>
              </div>
            </div>

            <!-- Review body -->
            <div class="space-y-1.5">
              <div class="flex items-center justify-between">
                <label class="text-xs uppercase tracking-widest text-white/40">
                  Written Review
                </label>
                <span
                  class="text-[10px] tabular-nums"
                  :class="
                    form.body.length >= BODY_MAX
                      ? 'text-red-400/70'
                      : 'text-white/20'
                  "
                >
                  {{ form.body.length }}/{{ BODY_MAX }}
                </span>
              </div>
              <textarea
                v-model="form.body"
                rows="3"
                :maxlength="BODY_MAX"
                placeholder="Be specific. The developers are listening. Allegedly."
                class="w-full bg-white/4 border border-white/8 rounded-xl px-4 py-3 text-sm placeholder-white/20 focus:outline-none focus:border-green-400/50 transition resize-none"
              />
            </div>

            <!-- Would you do it again -->
            <div class="border border-white/8 rounded-xl px-4 py-3 space-y-2">
              <p class="text-sm text-white/50">
                Given the chance — would you do it again?
              </p>
              <div class="flex gap-2">
                <button
                  @click="form.doItAgain = true"
                  class="flex-1 text-xs px-3 py-2 rounded-lg border transition"
                  :class="
                    form.doItAgain === true
                      ? 'border-green-400/60 bg-green-400/10 text-green-300'
                      : 'border-white/8 text-white/35 hover:border-white/20'
                  "
                >
                  🔁 Yeah, sign me up again
                </button>
                <button
                  @click="form.doItAgain = false"
                  class="flex-1 text-xs px-3 py-2 rounded-lg border transition"
                  :class="
                    form.doItAgain === false
                      ? 'border-red-400/60 bg-red-400/10 text-red-300'
                      : 'border-white/8 text-white/35 hover:border-white/20'
                  "
                >
                  🚪 One run was enough
                </button>
              </div>
            </div>

            <!-- Submit error -->
            <p v-if="submitError" class="text-red-400/70 text-xs text-center">
              {{ submitError }}
            </p>

            <!-- Submit -->
            <button
              @click="submitReview"
              :disabled="!canSubmit || loading"
              class="w-full py-3 rounded-xl font-semibold text-sm transition flex items-center justify-center gap-2"
              :class="
                canSubmit && !loading
                  ? 'bg-white text-black hover:bg-white/90'
                  : 'bg-white/10 text-white/25 cursor-not-allowed'
              "
            >
              <template v-if="loading">
                <span class="spinner" /> Submitting to the void...
              </template>
              <template v-else> <Send :size="14" /> Submit Review </template>
            </button>

            <p v-if="!canSubmit" class="text-white/20 text-xs text-center">
              {{
                !form.name.trim() && !form.rating
                  ? "A name and star rating are required."
                  : !form.name.trim()
                    ? "A name is required. Even a fake one."
                    : "A star rating is required. It's one click."
              }}
            </p>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useRoute } from "vue-router";
import { PenLine, X, Send, CheckCircle } from "lucide-vue-next";
import {
  fetchReviews,
  createReview,
  fetchStats,
  type Review,
  type ReviewStats,
} from "../utils/api";
import { getRandomQuote } from "../utils/quotes";

// ── Limits (must match backend domain/review.go) ──────────────────────────────
const NAME_MAX = 25;
const BODY_MAX = 500;
const TAG_MAX = 25;
const TAG_LIMIT = 6;

// ── Stats from API ────────────────────────────────────────────────────────────
const stats = ref<ReviewStats>({
  totalReviews: 0,
  averageRating: 0,
  ratingCounts: { 1: 0, 2: 0, 3: 0, 4: 0, 5: 0 },
});

const totalCount = computed(() => stats.value.totalReviews);
const averageRatingNum = computed(() => stats.value.averageRating);
const averageRating = computed(() => averageRatingNum.value.toFixed(1));

function starCount(star: number) {
  return stats.value.ratingCounts[star] ?? 0;
}
function starPercent(star: number) {
  return totalCount.value ? (starCount(star) / totalCount.value) * 100 : 0;
}
function starFillPercent(i: number, rating: number) {
  if (rating >= i) return 100;
  if (rating <= i - 1) return 0;
  return (rating - (i - 1)) * 100;
}

async function loadStats() {
  try {
    stats.value = await fetchStats();
  } catch {
    // stats will stay at defaults — non-critical
  }
}

// Random quote for fun
const quote = ref("");

// ── Pagination ────────────────────────────────────────────────────────────────
const PAGE_SIZE = 10;
const reviews = ref<Review[]>([]);
const offset = ref(0);
const hasMore = ref(true);
const loadingInitial = ref(true);
const loadingMore = ref(false);
const fetchError = ref("");

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
    const page = await fetchReviews(PAGE_SIZE, offset.value);
    reviews.value = reset ? page : [...reviews.value, ...page];
    offset.value += page.length;
    hasMore.value = page.length === PAGE_SIZE;
  } catch (e) {
    fetchError.value = "Couldn't load reviews. Is the server running?";
  } finally {
    loadingInitial.value = false;
    loadingMore.value = false;
  }
}

async function loadMore() {
  await loadReviews(false);
}

// ── Suggestions ───────────────────────────────────────────────────────────────
const proSuggestions = [
  "Sunsets",
  "Dogs",
  "Pizza",
  "Music",
  "Naps",
  "Love",
  "Coffee",
  "Comedy",
  "The ocean",
  "Flow states",
];
const conSuggestions = [
  "Mortality",
  "Taxes",
  "Mondays",
  "Traffic",
  "Aging",
  "Existential dread",
  "Slow Wi-Fi",
  "Allergies",
  "Meetings",
  "The news",
];

// ── Form state ────────────────────────────────────────────────────────────────
const hoverRating = ref(0);
const loading = ref(false);
const submitted = ref(false);
const modalOpen = ref(false);
const nameError = ref(false);
const submitError = ref("");
const proInput = ref("");
const conInput = ref("");

const form = ref({
  name: "",
  rating: 0,
  body: "",
  pros: [] as string[],
  cons: [] as string[],
  doItAgain: undefined as boolean | undefined,
});

const ratingLabels = [
  "",
  "Zero stars if I could",
  "Needs serious patching",
  "Mixed feelings, honestly",
  "Pretty good all things considered",
  "Tell the devs I said thanks",
];
const ratingLabel = computed(
  () => ratingLabels[hoverRating.value || form.value.rating] ?? "",
);
const canSubmit = computed(
  () => form.value.name.trim() !== "" && form.value.rating > 0,
);

// ── Tag helpers ───────────────────────────────────────────────────────────────
function toggleOrAdd(list: string[], item: string) {
  const idx = list.indexOf(item);
  if (idx !== -1) list.splice(idx, 1);
  else if (list.length < TAG_LIMIT) list.push(item);
}
function remove(list: string[], item: string) {
  const idx = list.indexOf(item);
  if (idx !== -1) list.splice(idx, 1);
}
function addCustom(list: string[], type: "pro" | "con") {
  if (list.length >= TAG_LIMIT) return;
  const inputRef = type === "pro" ? proInput : conInput;
  const val = inputRef.value.trim().slice(0, TAG_MAX);
  if (val && !list.includes(val)) list.push(val);
  inputRef.value = "";
}

// ── Modal ─────────────────────────────────────────────────────────────────────
function openModal() {
  modalOpen.value = true;
  document.body.style.overflow = "hidden";
}
function closeModal() {
  if (loading.value) return;
  modalOpen.value = false;
  document.body.style.overflow = "";
}

// ── Submit ────────────────────────────────────────────────────────────────────
function resetForm() {
  submitted.value = false;
  nameError.value = false;
  submitError.value = "";
  proInput.value = "";
  conInput.value = "";
  form.value = {
    name: "",
    rating: 0,
    body: "",
    pros: [],
    cons: [],
    doItAgain: undefined,
  };
}

async function submitReview() {
  if (!canSubmit.value || loading.value) return;
  if (!form.value.name.trim()) {
    nameError.value = true;
    return;
  }

  loading.value = true;
  submitError.value = "";

  try {
    await createReview({
      name: form.value.name.trim(),
      rating: form.value.rating,
      body: form.value.body.trim() || undefined,
      pros: form.value.pros.length ? form.value.pros : undefined,
      cons: form.value.cons.length ? form.value.cons : undefined,
      doItAgain: form.value.doItAgain,
    });
    submitted.value = true;
    await loadStats();
  } catch (e) {
    submitError.value = "Failed to submit review. Please try again later.";
  } finally {
    loading.value = false;
  }
}

// ── Init ──────────────────────────────────────────────────────────────────────
const route = useRoute();
onMounted(async () => {
  quote.value = getRandomQuote();
  await Promise.all([loadReviews(true), loadStats()]);
  if (route.query.action === "review") openModal();
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
.spinner {
  display: inline-block;
  width: 14px;
  height: 14px;
  border: 2px solid rgba(255, 255, 255, 0.2);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
.modal-enter-active {
  transition: opacity 0.2s ease;
}
.modal-leave-active {
  transition: opacity 0.15s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
