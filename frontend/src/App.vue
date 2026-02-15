<script setup>
import { ref, computed, onMounted } from 'vue'
import AppSidebar from './components/AppSidebar.vue'
import AppDashboard from './components/AppDashboard.vue'
import ChapterView from './components/ChapterView.vue'
import PomodoroTimer from './components/PomodoroTimer.vue'

const modules = ref([])
const currentChapter = ref(0)
const stats = ref({ progress: 0, completedLessons: 0, totalLessons: 0, totalMinutes: 0 })
const view = ref('loading')

const activeChapter = computed(() => {
  for (const mod of modules.value) {
    for (const ch of mod.chapters) {
      if (ch.number === currentChapter.value) return ch
    }
  }
  return null
})

const isChapterLocked = computed(() => {
  const ch = activeChapter.value
  return ch ? !ch.unlocked : false
})

async function loadModules() {
  const data = await window.goGetModules()
  modules.value = JSON.parse(data)
}

async function loadStats() {
  const data = await window.goGetStats()
  stats.value = JSON.parse(data)
}

async function showDashboard() {
  await loadStats()
  currentChapter.value = 0
  view.value = 'dashboard'
}

function navigateToChapter(chapterNum) {
  currentChapter.value = chapterNum
  view.value = 'chapter'
}

async function toggleLesson(lessonId) {
  const data = await window.goToggleLesson(lessonId)
  modules.value = JSON.parse(data)
}

function getCurrentChapter() {
  return currentChapter.value
}

onMounted(async () => {
  try {
    await loadModules()
    await showDashboard()
    if ('Notification' in window && Notification.permission === 'default') {
      Notification.requestPermission()
    }
  } catch (err) {
    console.error('Init error:', err)
  }
})
</script>

<template>
  <div class="bg-light text-dark min-h-screen flex font-['Segoe_UI',system-ui,-apple-system,sans-serif]">
    <AppSidebar
      :modules="modules"
      :active-chapter="currentChapter"
      @navigate="navigateToChapter"
      @show-dashboard="showDashboard"
    />

    <main class="flex-1 px-10 py-8 overflow-y-auto h-screen">
      <AppDashboard
        v-if="view === 'dashboard'"
        :modules="modules"
        :stats="stats"
      />

      <ChapterView
        v-else-if="view === 'chapter' && activeChapter"
        :chapter="activeChapter"
        :is-locked="isChapterLocked"
        @toggle-lesson="toggleLesson"
      />

      <div v-else class="flex items-center justify-center h-full">
        <p class="text-secondary">Carregando...</p>
      </div>
    </main>

    <PomodoroTimer :get-current-chapter="getCurrentChapter" />
  </div>
</template>
