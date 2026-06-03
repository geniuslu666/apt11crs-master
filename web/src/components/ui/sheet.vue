<script setup lang="ts">
interface Props {
  open?: boolean
  title?: string
  width?: string
  side?: 'left' | 'right'
}
const props = withDefaults(defineProps<Props>(), { open: false, width: '400px', side: 'right' })
const emit = defineEmits<{ 'update:open': [val: boolean] }>()
function close() { emit('update:open', false) }
</script>
<template>
  <Teleport to="body">
    <Transition name="sheet-fade">
      <div v-if="open" class="fixed inset-0 z-[1000] flex">
        <div class="fixed inset-0 bg-black/40" @click="close" />
        <Transition :name="side === 'right' ? 'sheet-slide-right' : 'sheet-slide-left'">
          <div v-if="open" :class="['fixed top-0 h-full bg-white shadow-xl flex flex-col', side === 'right' ? 'right-0' : 'left-0']" :style="{ width }">
            <div class="flex items-center justify-between px-4 py-3 border-b border-border flex-shrink-0">
              <slot name="header">
                <h2 class="text-[14px] font-semibold text-foreground">{{ title }}</h2>
              </slot>
              <button @click="close" class="w-7 h-7 flex items-center justify-center rounded-md text-muted-foreground hover:bg-muted hover:text-foreground transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
              </button>
            </div>
            <div class="flex-1 overflow-y-auto px-4 py-4 text-[13px]">
              <slot />
            </div>
            <div v-if="$slots.footer" class="flex items-center justify-end gap-2 px-4 py-3 border-t border-border flex-shrink-0">
              <slot name="footer" />
            </div>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>
<style scoped>
.sheet-fade-enter-active, .sheet-fade-leave-active { transition: opacity 0.2s ease; }
.sheet-fade-enter-from, .sheet-fade-leave-to { opacity: 0; }
.sheet-slide-right-enter-active, .sheet-slide-right-leave-active { transition: transform 0.25s ease; }
.sheet-slide-right-enter-from, .sheet-slide-right-leave-to { transform: translateX(100%); }
.sheet-slide-left-enter-active, .sheet-slide-left-leave-active { transition: transform 0.25s ease; }
.sheet-slide-left-enter-from, .sheet-slide-left-leave-to { transform: translateX(-100%); }
</style>
