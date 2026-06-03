<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
interface MenuItem { label: string; key: string; disabled?: boolean; danger?: boolean }
interface Props { items: MenuItem[]; label?: string }
withDefaults(defineProps<Props>(), { label: '更多' })
const emit = defineEmits<{ select: [key: string] }>()
const open = ref(false)
const containerRef = ref<HTMLElement>()
function handleOutside(e: MouseEvent) {
  if (containerRef.value && !containerRef.value.contains(e.target as Node)) open.value = false
}
onMounted(() => document.addEventListener('mousedown', handleOutside))
onBeforeUnmount(() => document.removeEventListener('mousedown', handleOutside))
function select(key: string) { emit('select', key); open.value = false }
</script>
<template>
  <div ref="containerRef" class="relative inline-flex">
    <button type="button" @click="open = !open"
      class="inline-flex items-center gap-1 text-[13px] font-medium text-stripe-600 hover:text-stripe-800 transition-colors cursor-pointer px-1">
      {{ label }}
      <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" :class="open ? 'rotate-180 transition-transform' : 'transition-transform'"><path d="m6 9 6 6 6-6"/></svg>
    </button>
    <Transition name="select-drop">
      <div v-if="open" class="absolute right-0 top-full mt-1 z-[200] bg-white rounded-md border border-border shadow-lg min-w-[120px] py-1">
        <button v-for="item in items" :key="item.key" type="button" @click="select(item.key)" :disabled="item.disabled"
          :class="['w-full text-left px-3 py-1.5 text-[13px] transition-colors', item.danger ? 'text-destructive hover:bg-destructive/10' : 'text-foreground hover:bg-muted', item.disabled ? 'opacity-50 cursor-not-allowed' : '']"
        >{{ item.label }}</button>
      </div>
    </Transition>
  </div>
</template>
<style scoped>
.select-drop-enter-active, .select-drop-leave-active { transition: opacity 0.1s ease, transform 0.1s ease; }
.select-drop-enter-from, .select-drop-leave-to { opacity: 0; transform: translateY(-4px); }
</style>
