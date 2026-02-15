import { ref, computed, onUnmounted } from 'vue'

const MODES = {
  focus: { label: 'Foco', minutes: 25 },
  short: { label: 'Pausa', minutes: 5 },
  long: { label: 'Longa', minutes: 15 },
}

export function usePomodoro(getCurrentChapter) {
  const mode = ref('focus')
  const running = ref(false)
  const remaining = ref(MODES.focus.minutes * 60)
  const minimized = ref(false)

  let intervalId = null

  const timeDisplay = computed(() => {
    const mins = Math.floor(remaining.value / 60)
    const secs = remaining.value % 60
    return `${String(mins).padStart(2, '0')}:${String(secs).padStart(2, '0')}`
  })

  const currentMode = computed(() => MODES[mode.value])

  const modes = computed(() =>
    Object.entries(MODES).map(([key, val]) => ({
      key,
      label: val.label,
      minutes: val.minutes,
      active: mode.value === key,
    }))
  )

  function toggle() {
    if (running.value) {
      clearInterval(intervalId)
      running.value = false
    } else {
      running.value = true
      intervalId = setInterval(() => {
        remaining.value--
        if (remaining.value <= 0) {
          clearInterval(intervalId)
          running.value = false
          onComplete()
        }
      }, 1000)
    }
  }

  function reset() {
    if (intervalId) clearInterval(intervalId)
    running.value = false
    remaining.value = MODES[mode.value].minutes * 60
  }

  function setMode(newMode) {
    if (running.value) return
    mode.value = newMode
    remaining.value = MODES[newMode].minutes * 60
  }

  function toggleMinimize() {
    minimized.value = !minimized.value
  }

  function onComplete() {
    const currentModeKey = mode.value
    const minutes = MODES[currentModeKey].minutes

    if (currentModeKey === 'focus') {
      const chapter = getCurrentChapter()
      window.goSaveTimerSession(minutes, chapter).catch(err =>
        console.error('Error saving timer:', err)
      )
    }

    if ('Notification' in window && Notification.permission === 'granted') {
      new Notification('GoTrack - Pomodoro', {
        body: currentModeKey === 'focus' ? 'Hora da pausa!' : 'Hora de focar!',
      })
    }

    setMode(currentModeKey === 'focus' ? 'short' : 'focus')
  }

  onUnmounted(() => {
    if (intervalId) clearInterval(intervalId)
  })

  return {
    mode,
    running,
    remaining,
    minimized,
    timeDisplay,
    currentMode,
    modes,
    toggle,
    reset,
    setMode,
    toggleMinimize,
  }
}
