<script setup lang="ts">
interface Tab { key: string; label: string }
interface Props { tabs: Tab[]; modelValue: string }
defineProps<Props>()
const emit = defineEmits<{ 'update:modelValue': [key: string] }>()
</script>
<template>
  <div class="flex items-center gap-0.5 border-b border-border overflow-x-auto">
    <button v-for="tab in tabs" :key="tab.key" type="button"
      @click="emit('update:modelValue', tab.key)"
      :class="['px-3 py-2 text-sm font-medium whitespace-nowrap cursor-pointer transition-colors select-none rounded-t-md -mb-px', modelValue === tab.key ? 'text-primary border-b-2 border-primary' : 'text-muted-foreground hover:text-foreground hover:bg-muted/50']"
    >{{ tab.label }}</button>
  </div>
</template>
