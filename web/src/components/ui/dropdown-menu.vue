<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { ChevronDown } from '@lucide/vue'
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
      class="inline-flex h-7 items-center gap-1 rounded-md px-2 text-[13px] font-medium text-primary hover:bg-accent hover:text-accent-foreground transition-colors cursor-pointer">
      {{ label }}
      <ChevronDown :size="13" :class="open ? 'rotate-180 transition-transform' : 'transition-transform'" />
    </button>
    <Transition name="select-drop">
      <div v-if="open" class="absolute right-0 top-full mt-1 z-[200] bg-popover text-popover-foreground rounded-md border border-border shadow-lg min-w-[128px] py-1">
        <button v-for="item in items" :key="item.key" type="button" @click="select(item.key)" :disabled="item.disabled"
          :class="['w-full text-left px-3 py-1.5 text-[13px] transition-colors', item.danger ? 'text-destructive hover:bg-destructive/10' : 'text-popover-foreground hover:bg-accent hover:text-accent-foreground', item.disabled ? 'opacity-50 cursor-not-allowed' : '']"
        >{{ item.label }}</button>
      </div>
    </Transition>
  </div>
</template>
<style scoped>
.select-drop-enter-active, .select-drop-leave-active { transition: opacity 0.1s ease, transform 0.1s ease; }
.select-drop-enter-from, .select-drop-leave-to { opacity: 0; transform: translateY(-4px); }
</style>
