<script setup lang="ts">
import { cn } from '@/utils/cn'

interface Props {
  open?: boolean
  title?: string
  maxWidth?: string
}
const props = withDefaults(defineProps<Props>(), { open: false, maxWidth: '520px' })
const emit = defineEmits<{ 'update:open': [val: boolean] }>()
function close() { emit('update:open', false) }
</script>
<template>
  <Teleport to="body">
    <Transition name="dialog-fade">
      <div v-if="open" class="fixed inset-0 z-[1000] flex items-center justify-center p-4">
        <div class="fixed inset-0 bg-black/40 backdrop-blur-[1px]" @click="close" />
        <div class="relative bg-white rounded-xl shadow-2xl w-full flex flex-col" :style="{ maxWidth, maxHeight: '90vh' }">
          <!-- Header -->
          <div v-if="title || $slots.header" class="flex items-center justify-between px-5 py-3 border-b border-border flex-shrink-0">
            <slot name="header">
              <h2 class="text-[14px] font-semibold text-foreground">{{ title }}</h2>
            </slot>
            <button @click="close" class="w-7 h-7 flex items-center justify-center rounded-md text-muted-foreground hover:bg-muted hover:text-foreground transition-colors">
              <svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
            </button>
          </div>
          <!-- Body -->
          <div class="overflow-y-auto flex-1 px-5 py-4 text-[13px]">
            <slot />
          </div>
          <!-- Footer -->
          <div v-if="$slots.footer" class="flex items-center justify-end gap-2 px-5 py-3 border-t border-border flex-shrink-0">
            <slot name="footer" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>
<style scoped>
.dialog-fade-enter-active, .dialog-fade-leave-active { transition: opacity 0.15s ease; }
.dialog-fade-enter-from, .dialog-fade-leave-to { opacity: 0; }
</style>
