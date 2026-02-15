<script setup>
import { usePomodoro } from '../composables/usePomodoro.js'

const props = defineProps({
  getCurrentChapter: { type: Function, required: true },
})

const {
  running,
  minimized,
  timeDisplay,
  currentMode,
  modes,
  toggle,
  reset,
  setMode,
  toggleMinimize,
} = usePomodoro(props.getCurrentChapter)
</script>

<template>
  <div
    class="fixed bottom-6 right-6 z-50"
    :class="minimized
      ? 'bg-white border border-border rounded-xl px-4 py-3 shadow-lg'
      : 'bg-white border border-border rounded-2xl p-5 shadow-lg w-60'"
  >
    <!-- Minimized -->
    <div v-if="minimized" class="flex items-center gap-3">
      <span class="text-[13px] font-semibold text-secondary uppercase tracking-wide">Pomodoro</span>
      <span class="text-[13px] font-bold text-primary tabular-nums">{{ timeDisplay }}</span>
      <button
        class="w-6 h-6 flex items-center justify-center text-secondary hover:text-dark rounded transition-colors"
        @click="toggleMinimize"
        title="Expandir"
      >
        <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <polyline points="18 15 12 9 6 15"></polyline>
        </svg>
      </button>
    </div>

    <!-- Expanded -->
    <template v-else>
      <div class="flex justify-between items-center mb-4">
        <h4 class="text-[13px] font-semibold text-secondary uppercase tracking-wide">Pomodoro</h4>
        <button
          class="w-6 h-6 flex items-center justify-center text-secondary hover:text-dark rounded transition-colors"
          @click="toggleMinimize"
          title="Minimizar"
        >
          <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="6 9 12 15 18 9"></polyline>
          </svg>
        </button>
      </div>

      <div class="text-center mb-4">
        <div class="text-5xl font-bold text-dark tabular-nums leading-none">{{ timeDisplay }}</div>
        <div class="text-xs text-secondary mt-1.5">{{ currentMode.label }}</div>
      </div>

      <div class="flex gap-1 mb-4">
        <button
          v-for="m in modes"
          :key="m.key"
          class="flex-1 py-1.5 rounded-md text-[11px] font-medium cursor-pointer transition-all border"
          :class="m.active
            ? 'bg-primary text-white border-primary'
            : 'bg-white text-secondary border-border hover:border-primary'"
          @click="setMode(m.key)"
        >
          {{ m.label }} ({{ m.minutes }}m)
        </button>
      </div>

      <div class="flex gap-2">
        <button
          class="flex-1 py-2.5 rounded-lg text-[13px] font-semibold cursor-pointer transition-all bg-primary text-white hover:bg-blue-700"
          @click="toggle"
        >
          {{ running ? 'Pausar' : 'Iniciar' }}
        </button>
        <button
          class="flex-1 py-2.5 rounded-lg text-[13px] font-semibold cursor-pointer transition-all bg-gray-100 text-secondary hover:text-dark"
          @click="reset"
        >
          Reset
        </button>
      </div>
    </template>
  </div>
</template>
