<script setup lang="ts">
interface Props { src?: string; name?: string; size?: 'sm' | 'md' | 'lg' | 'xl' }
const props = withDefaults(defineProps<Props>(), { size: 'md' })
const sizeMap = { sm: 'w-8 h-8 text-xs', md: 'w-10 h-10 text-sm', lg: 'w-14 h-14 text-base', xl: 'w-20 h-20 text-lg' }
</script>
<template>
  <div :class="['rounded-full overflow-hidden flex items-center justify-center bg-muted flex-shrink-0', sizeMap[size]]">
    <img v-if="src" :src="src" class="w-full h-full object-cover" :alt="name || 'avatar'" @error="($event.target as HTMLImageElement).style.display = 'none'" />
    <span v-else class="font-medium text-muted-foreground uppercase">{{ name ? name[0] : '?' }}</span>
  </div>
</template>
