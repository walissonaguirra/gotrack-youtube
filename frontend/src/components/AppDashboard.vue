<script setup>
import { computed } from 'vue'

const props = defineProps({
  modules: { type: Array, required: true },
  stats: { type: Object, required: true },
})

const globalPct = computed(() => Math.round(props.stats.progress * 100))

const timeStr = computed(() => {
  const hours = Math.floor(props.stats.totalMinutes / 60)
  const mins = props.stats.totalMinutes % 60
  return hours > 0 ? `${hours}h ${mins}m` : `${mins}m`
})
</script>

<template>
  <div>
    <div class="mb-8">
      <h2 class="text-[28px] font-bold text-dark">Dashboard</h2>
      <div class="text-[13px] text-secondary mt-1">Seu progresso no curso Aprenda Go</div>
    </div>

    <div class="grid grid-cols-[repeat(auto-fit,minmax(280px,1fr))] gap-5 mb-8">
      <div class="bg-white border border-border rounded-lg p-6 shadow-sm">
        <div class="text-4xl font-bold text-primary">{{ globalPct }}%</div>
        <div class="text-[13px] text-secondary mt-1">Progresso geral</div>
      </div>
      <div class="bg-white border border-border rounded-lg p-6 shadow-sm">
        <div class="text-4xl font-bold text-success">{{ stats.completedLessons }}/{{ stats.totalLessons }}</div>
        <div class="text-[13px] text-secondary mt-1">Aulas completas</div>
      </div>
      <div class="bg-white border border-border rounded-lg p-6 shadow-sm">
        <div class="text-4xl font-bold text-warning">{{ timeStr }}</div>
        <div class="text-[13px] text-secondary mt-1">Tempo estudado</div>
      </div>
    </div>

    <h3 class="text-base font-semibold text-secondary mb-4">Modulos</h3>

    <div v-for="mod in modules" :key="mod.name" class="bg-white border border-border rounded-lg p-6 shadow-sm mb-3">
      <div class="flex justify-between items-center mb-2">
        <span class="font-semibold text-dark">{{ mod.name }}</span>
        <span class="text-secondary text-[13px]">{{ Math.round(mod.progress * 100) }}%</span>
      </div>
      <div class="h-1.5 bg-gray-200 rounded-full overflow-hidden">
        <div
          class="h-full bg-primary rounded-full transition-all duration-400"
          :style="{ width: Math.round(mod.progress * 100) + '%' }"
        ></div>
      </div>
    </div>
  </div>
</template>
