<script setup lang="ts">
import { cn } from '@/utils/cn';

interface Props {
  class?: string;
  placeholder?: string;
  type?: string;
  disabled?: boolean;
  modelValue?: string | number;
}

const props = withDefaults(defineProps<Props>(), {
  type: 'text',
});

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
  (e: 'change', value: string): void;
}>();
</script>

<template>
  <input
    :type="type"
    :placeholder="placeholder"
    :disabled="disabled"
    :value="modelValue"
    :class="cn(
      'flex h-8 w-full rounded-md border border-input bg-background px-3 py-1 text-[13px] shadow-sm',
      'transition-colors placeholder:text-muted-foreground',
      'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring',
      'disabled:cursor-not-allowed disabled:opacity-50',
      props.class
    )"
    @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
    @change="emit('change', ($event.target as HTMLInputElement).value)"
  />
</template>
