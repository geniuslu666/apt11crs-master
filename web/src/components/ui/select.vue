<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { cn } from '@/utils/cn'

interface Option { label: string; value: any; [key: string]: any }
interface Props {
  modelValue?: any
  options?: any[]
  placeholder?: string
  clearable?: boolean
  filterable?: boolean
  disabled?: boolean
  labelField?: string
  valueField?: string
  class?: string
}
const props = withDefaults(defineProps<Props>(), {
  options: () => [],
  placeholder: '请选择',
  labelField: 'label',
  valueField: 'value',
})
const emit = defineEmits<{ 'update:modelValue': [val: any] }>()

const open = ref(false)
const search = ref('')

const normalizedOptions = computed(() =>
  props.options.map(o => ({
    label: o[props.labelField] ?? o.label,
    value: o[props.valueField] ?? o.value,
    disabled: o.disabled,
  }))
)

const filtered = computed(() =>
  search.value
    ? normalizedOptions.value.filter(o => String(o.label).toLowerCase().includes(search.value.toLowerCase()))
    : normalizedOptions.value
)

const selectedLabel = computed(() => {
  if (props.modelValue === null || props.modelValue === undefined || props.modelValue === '') return ''
  const found = normalizedOptions.value.find(o => o.value === props.modelValue)
  return found?.label ?? ''
})

function select(val: any) {
  emit('update:modelValue', val)
  open.value = false
  search.value = ''
}

function clear(e: Event) {
  e.stopPropagation()
  emit('update:modelValue', null)
}

function toggle() {
  if (!props.disabled) { open.value = !open.value; search.value = '' }
}

const containerRef = ref<HTMLElement>()
function handleOutside(e: MouseEvent) {
  if (containerRef.value && !containerRef.value.contains(e.target as Node)) open.value = false
}
onMounted(() => document.addEventListener('mousedown', handleOutside))
onBeforeUnmount(() => document.removeEventListener('mousedown', handleOutside))
</script>
<template>
  <div ref="containerRef" class="relative w-full">
    <button type="button" @click="toggle" :disabled="disabled"
      :class="cn('w-full h-8 px-3 text-left text-[13px] rounded-md border border-input bg-background flex items-center justify-between gap-1 transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-1 disabled:opacity-50 disabled:cursor-not-allowed hover:border-ring/50', props.class)"
    >
      <span :class="selectedLabel ? 'text-foreground' : 'text-muted-foreground'">
        {{ selectedLabel || placeholder }}
      </span>
      <div class="flex items-center gap-0.5 flex-shrink-0">
        <button v-if="clearable && modelValue !== null && modelValue !== undefined && modelValue !== ''" type="button" @click="clear"
          class="w-4 h-4 flex items-center justify-center rounded-sm hover:bg-muted text-muted-foreground">
          <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
        </button>
        <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" :class="['transition-transform', open ? 'rotate-180' : '']"><path d="m6 9 6 6 6-6"/></svg>
      </div>
    </button>
    <Transition name="select-drop">
      <div v-if="open" class="absolute z-[200] w-full mt-1 bg-white rounded-md border border-border shadow-lg overflow-hidden">
        <div v-if="filterable" class="p-2 border-b border-border">
          <input v-model="search" class="w-full h-7 px-2 text-sm rounded border border-input outline-none focus:ring-1 focus:ring-ring bg-background" placeholder="搜索..." @keydown.esc="open = false" />
        </div>
        <div class="max-h-52 overflow-y-auto py-1">
          <div v-if="filtered.length === 0" class="px-3 py-2 text-sm text-muted-foreground text-center">暂无选项</div>
          <button v-for="opt in filtered" :key="opt.value" type="button" @click="select(opt.value)"
            :class="['w-full text-left px-3 py-1.5 text-[13px] transition-colors', opt.value === modelValue ? 'bg-primary/10 text-primary font-medium' : 'hover:bg-muted text-foreground', opt.disabled ? 'opacity-50 cursor-not-allowed' : '']"
            :disabled="opt.disabled"
          >{{ opt.label }}</button>
        </div>
      </div>
    </Transition>
  </div>
</template>
<style scoped>
.select-drop-enter-active, .select-drop-leave-active { transition: opacity 0.1s ease, transform 0.1s ease; }
.select-drop-enter-from, .select-drop-leave-to { opacity: 0; transform: translateY(-4px); }
</style>
