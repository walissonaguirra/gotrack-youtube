<script setup>
defineProps({
  url: { type: String, default: '' },
})

const emit = defineEmits(['close'])
</script>

<template>
  <div v-if="url" class="mb-6">
    <div class="relative bg-black rounded-lg overflow-hidden" style="padding-top: 56.25%">
      <iframe
        class="absolute inset-0 w-full h-full"
        :src="url"
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
        allowfullscreen
      ></iframe>
    </div>
    <button
      class="mt-2 px-3 py-1.5 text-xs text-secondary hover:text-dark bg-gray-100 rounded-md transition-colors"
      @click="emit('close')"
    >
      Fechar player
    </button>
  </div>
</template>
