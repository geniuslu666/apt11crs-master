<script setup lang="ts">
import { cn } from '@/utils/cn'
interface Props { modelValue?: string; placeholder?: string; rows?: number; class?: string; disabled?: boolean }
const props = withDefaults(defineProps<Props>(), { rows: 3 })
const emit = defineEmits<{ 'update:modelValue': [v: string] }>()
</script>
<template>
  <textarea
    :value="modelValue"
    @input="emit('update:modelValue', ($event.target as HTMLTextAreaElement).value)"
    :placeholder="placeholder"
    :rows="rows"
    :disabled="disabled"
    :class="cn('flex w-full rounded-md border border-input bg-background px-3 py-2 text-sm shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-1 disabled:cursor-not-allowed disabled:opacity-50 resize-none', props.class)"
  />
</template>
