<script setup>
import { computed } from 'vue'

const props = defineProps({
  modules: { type: Array, required: true },
  activeChapter: { type: Number, default: 0 },
})

const emit = defineEmits(['navigate', 'show-dashboard'])

const totalLessons = computed(() => {
  let total = 0
  props.modules.forEach(mod => {
    mod.chapters.forEach(ch => { total += ch.lessons.length })
  })
  return total
})

const completedLessons = computed(() => {
  let completed = 0
  props.modules.forEach(mod => {
    mod.chapters.forEach(ch => {
      completed += ch.lessons.filter(l => l.completed).length
    })
  })
  return completed
})

const globalPct = computed(() =>
  totalLessons.value > 0 ? Math.round((completedLessons.value / totalLessons.value) * 100) : 0
)

function chapterProgress(ch) {
  const done = ch.lessons.filter(l => l.completed).length
  return ch.lessons.length > 0 ? Math.round((done / ch.lessons.length) * 100) : 0
}

function isChapterComplete(ch) {
  return ch.lessons.every(l => l.completed)
}
</script>

<template>
  <aside class="w-[260px] bg-white border-r border-border flex flex-col shrink-0 h-screen sticky top-0">
    <div class="px-5 py-6 border-b border-border cursor-pointer" @click="emit('show-dashboard')">
      <h1 class="text-[22px] font-bold text-primary">GoTrack</h1>
      <div class="text-xs text-secondary mt-1">Aprenda Go - Painel de Estudos</div>
    </div>

    <div class="px-5 py-4 border-b border-border">
      <div class="flex justify-between text-xs text-secondary mb-2">
        <span>Progresso geral</span>
        <span>{{ completedLessons }}/{{ totalLessons }} ({{ globalPct }}%)</span>
      </div>
      <div class="h-1.5 bg-gray-200 rounded-full overflow-hidden">
        <div
          class="h-full bg-primary rounded-full transition-all duration-400"
          :style="{ width: globalPct + '%' }"
        ></div>
      </div>
    </div>

    <nav class="flex-1 overflow-y-auto py-3">
      <div v-for="mod in modules" :key="mod.name" class="px-3 mb-1">
        <div class="px-3 py-2.5 text-[11px] font-semibold uppercase tracking-wider text-secondary">
          {{ mod.name }}
        </div>
        <div
          v-for="ch in mod.chapters"
          :key="ch.number"
          class="flex items-center gap-2.5 px-3 py-2 rounded-md cursor-pointer text-[13px] transition-all duration-200"
          :class="{
            'bg-blue-50 text-primary font-medium': ch.number === activeChapter,
            'opacity-40 cursor-not-allowed text-secondary': !ch.unlocked && ch.number !== activeChapter,
            'text-secondary hover:bg-gray-100 hover:text-dark': ch.unlocked && ch.number !== activeChapter,
          }"
          @click="emit('navigate', ch.number)"
        >
          <span
            class="w-7 h-7 rounded-md flex items-center justify-center text-[11px] font-bold shrink-0"
            :class="isChapterComplete(ch) ? 'bg-success text-white' : 'bg-gray-100 text-secondary'"
          >
            <svg v-if="!ch.unlocked" class="w-3 h-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="11" width="18" height="11" rx="2"></rect>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
            </svg>
            <template v-else>{{ ch.number }}</template>
          </span>
          <span class="truncate">{{ ch.title }}</span>
          <span class="w-8 text-[11px] text-right ml-auto shrink-0 text-secondary">
            {{ chapterProgress(ch) }}%
          </span>
        </div>
      </div>
    </nav>
  </aside>
</template>
