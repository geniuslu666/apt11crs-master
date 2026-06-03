<script setup lang="ts">
import { cn } from '@/utils/cn'
interface Props { modelValue?: number; placeholder?: string; min?: number; max?: number; step?: number; class?: string; showButtons?: boolean }
const props = withDefaults(defineProps<Props>(), { step: 1, showButtons: false })
const emit = defineEmits<{ 'update:modelValue': [v: number] }>()
function handleInput(e: Event) {
  const val = (e.target as HTMLInputElement).value
  if (val === '' || val === '-') { emit('update:modelValue', 0); return }
  const num = Number(val)
  if (!isNaN(num)) emit('update:modelValue', num)
}
function inc() { emit('update:modelValue', (props.modelValue ?? 0) + (props.step ?? 1)) }
function dec() { emit('update:modelValue', (props.modelValue ?? 0) - (props.step ?? 1)) }
</script>
<template>
  <div class="flex items-center w-full">
    <button v-if="showButtons" type="button" @click="dec" class="h-9 w-9 flex items-center justify-center border border-r-0 border-input rounded-l-md bg-muted hover:bg-muted/80 text-muted-foreground flex-shrink-0">−</button>
    <input type="number" :value="modelValue" @input="handleInput"
      :placeholder="placeholder" :min="min" :max="max" :step="step"
      :class="cn('flex h-9 w-full border border-input bg-background px-3 py-1 text-sm rounded-md shadow-sm transition-colors placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-1', showButtons ? 'rounded-none' : '', props.class)"
    />
    <button v-if="showButtons" type="button" @click="inc" class="h-9 w-9 flex items-center justify-center border border-l-0 border-input rounded-r-md bg-muted hover:bg-muted/80 text-muted-foreground flex-shrink-0">+</button>
  </div>
</template>
