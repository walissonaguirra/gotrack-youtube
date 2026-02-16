<script setup>
import { ref, computed } from 'vue'
import LessonItem from './LessonItem.vue'
import VideoPlayer from './VideoPlayer.vue'

const props = defineProps({
  chapter: { type: Object, required: true },
  isLocked: { type: Boolean, default: false },
})

const emit = defineEmits(['toggle-lesson'])

const videoUrl = ref('')
const watchingLessonId = ref('')

const completedCount = computed(() =>
  props.chapter.lessons.filter(l => l.completed).length
)

const progress = computed(() =>
  Math.round((completedCount.value / props.chapter.lessons.length) * 100)
)

function onToggle(lessonId) {
  emit('toggle-lesson', lessonId)
}

async function onWatch(lessonId) {
  try {
    videoUrl.value = await window.goGetYouTubeURL(lessonId)
    watchingLessonId.value = lessonId
  } catch (err) {
    console.error('Watch video error:', err)
  }
}

function closePlayer() {
  videoUrl.value = ''
  watchingLessonId.value = ''
}
</script>

<template>
  <!-- Locked state -->
  <div v-if="isLocked" class="flex flex-col items-center justify-center py-16 text-center">
    <div class="w-16 h-16 text-secondary mb-4">
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
        <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
      </svg>
    </div>
    <h3 class="text-lg text-secondary mb-2">Cap. {{ chapter.number }} - {{ chapter.title }}</h3>
    <p class="text-sm text-secondary">
      Complete todas as aulas do capitulo {{ chapter.requiredBy }} para desbloquear estes exercicios.
    </p>
  </div>

  <!-- Unlocked state -->
  <div v-else>
    <div class="mb-8">
      <h2 class="text-[28px] font-bold text-dark">Cap. {{ chapter.number }} - {{ chapter.title }}</h2>
      <div class="text-[13px] text-secondary mt-1">
        {{ completedCount }} de {{ chapter.lessons.length }} aulas completas ({{ progress }}%)
      </div>
    </div>

    <VideoPlayer :url="videoUrl" @close="closePlayer" />

    <div class="h-1 bg-gray-200 rounded-full overflow-hidden mb-6">
      <div
        class="h-full bg-primary rounded-full transition-all duration-400"
        :style="{ width: progress + '%' }"
      ></div>
    </div>

    <div class="flex flex-col gap-1.5 mb-8">
      <LessonItem
        v-for="lesson in chapter.lessons"
        :key="lesson.id"
        :lesson="lesson"
        :is-watching="watchingLessonId === lesson.id"
        @toggle="onToggle"
        @watch="onWatch"
      />
    </div>
  </div>
</template>
