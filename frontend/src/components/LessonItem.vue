<script setup>
defineProps({
  lesson: { type: Object, required: true },
  isWatching: { type: Boolean, default: false },
})

const emit = defineEmits(['toggle', 'watch'])
</script>

<template>
  <div
    class="flex items-center gap-3.5 px-4 py-3.5 border rounded-lg transition-all duration-200 cursor-pointer"
    :class="[
      isWatching
        ? 'border-primary bg-blue-50 dark:bg-blue-900/20 ring-2 ring-primary/20'
        : lesson.completed
          ? 'border-success bg-green-50 dark:bg-green-900/20'
          : 'border-border dark:border-gray-700 bg-white dark:bg-gray-800 hover:border-primary'
    ]"
  >
    <div
      class="w-[22px] h-[22px] rounded-full border-2 flex items-center justify-center shrink-0"
      :class="lesson.completed ? 'border-success bg-success text-white' : 'border-gray-300 dark:border-gray-600'"
      @click="emit('toggle', lesson.id)"
    >
      <svg
        v-if="lesson.completed"
        class="w-3.5 h-3.5"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="3"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <polyline points="20 6 9 17 4 12"></polyline>
      </svg>
    </div>

    <div class="flex-1 min-w-0 cursor-pointer" @click="emit('toggle', lesson.id)">
      <div class="text-[11px] text-secondary dark:text-gray-400">Aula {{ lesson.lessonNumber }}</div>
      <div class="text-sm font-medium truncate text-dark dark:text-gray-100">{{ lesson.title }}</div>
    </div>

    <button
      class="w-9 h-9 rounded-lg bg-primary text-white flex items-center justify-center shrink-0 hover:bg-blue-700 transition-colors duration-200"
      @click.stop="emit('watch', lesson.id)"
      title="Assistir"
    >
      <svg class="w-4 h-4" viewBox="0 0 24 24" fill="currentColor">
        <path d="M8 5v14l11-7z" />
      </svg>
    </button>
  </div>
</template>
