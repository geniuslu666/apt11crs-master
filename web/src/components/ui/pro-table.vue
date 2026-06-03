<script setup lang="ts" generic="T extends Record<string, any>">
import { ref, computed, onMounted, watch } from 'vue'

interface Column {
  key: string
  title: string | (() => string)
  width?: number
  minWidth?: number
  render?: (row: T, index: number) => any
  ellipsis?: boolean
}

interface Props {
  columns: Column[]
  request: (params: Record<string, any>) => Promise<{ list: T[]; pageCount: number; totalCount?: number }>
  rowKey?: (row: T) => string | number
  extraParams?: Record<string, any>
  pageSizes?: number[]
  defaultPageSize?: number
  flat?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  pageSizes: () => [10, 20, 50],
  defaultPageSize: 10,
  flat: false,
})

const loading = ref(false)
const data = ref<T[]>([])
const page = ref(1)
const pageSize = ref(props.defaultPageSize)
const pageCount = ref(0)
const totalCount = ref(0)

async function fetchData() {
  loading.value = true
  try {
    const res = await props.request({ page: page.value, pageSize: pageSize.value, ...(props.extraParams || {}) })
    data.value = res.list || []
    pageCount.value = res.pageCount || 0
    totalCount.value = res.totalCount || 0
  } finally {
    loading.value = false
  }
}

function reload() {
  page.value = 1
  fetchData()
}

function changePage(p: number) {
  page.value = p
  fetchData()
}

function changePageSize(s: number) {
  pageSize.value = s
  page.value = 1
  fetchData()
}

onMounted(fetchData)

watch(() => props.extraParams, () => {
  page.value = 1
  fetchData()
}, { deep: true })

defineExpose({ reload })

function getColumnTitle(col: Column) {
  return typeof col.title === 'function' ? col.title() : col.title
}

function getCellValue(row: T, key: string) {
  return key.split('.').reduce((obj: any, k) => obj?.[k], row) ?? '--'
}

const pages = computed(() => {
  if (pageCount.value <= 7) return Array.from({ length: pageCount.value }, (_, i) => i + 1)
  const arr: (number | '...')[] = [1]
  if (page.value > 3) arr.push('...')
  for (let i = Math.max(2, page.value - 1); i <= Math.min(pageCount.value - 1, page.value + 1); i++) arr.push(i)
  if (page.value < pageCount.value - 2) arr.push('...')
  if (pageCount.value > 1) arr.push(pageCount.value)
  return arr
})
</script>

<template>
  <div class="space-y-3">
    <!-- Toolbar -->
    <div v-if="$slots.tableTitle" class="flex items-center justify-between">
      <div class="flex items-center gap-2">
        <slot name="tableTitle" />
      </div>
    </div>

    <!-- Table -->
    <div :class="['relative overflow-hidden', flat ? '' : 'rounded-lg border border-border bg-white']">
      <!-- Loading overlay -->
      <div v-if="loading" class="absolute inset-0 bg-white/70 z-10 flex items-center justify-center">
        <svg class="animate-spin w-6 h-6 text-primary" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
        </svg>
      </div>

      <!-- PC Table -->
      <div class="hidden md:block overflow-x-auto">
        <table class="w-full text-[13px]">
          <thead>
            <tr class="border-b border-border bg-[#f8fafc]">
              <th v-for="col in columns" :key="col.key"
                class="text-left px-3 py-2 text-[11px] font-medium text-muted-foreground whitespace-nowrap"
                :style="col.width ? { width: col.width + 'px', minWidth: col.width + 'px' } : {}">
                {{ getColumnTitle(col) }}
              </th>
              <th v-if="$slots.action" class="text-left px-3 py-2 text-[11px] font-medium text-muted-foreground whitespace-nowrap w-[1%]">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-border">
            <tr v-if="data.length === 0 && !loading">
              <td :colspan="columns.length + ($slots.action ? 1 : 0)" class="px-3 py-8 text-center text-muted-foreground text-[13px]">暂无数据</td>
            </tr>
            <tr v-for="(row, index) in data" :key="rowKey ? rowKey(row) : index"
              class="hover:bg-[#f8fafc] transition-colors">
              <td v-for="col in columns" :key="col.key"
                class="px-3 py-2 text-foreground"
                :style="col.width ? { width: col.width + 'px', minWidth: col.width + 'px' } : {}">
                <div :class="col.ellipsis ? 'truncate max-w-[200px]' : ''">
                  <component v-if="col.render" :is="() => col.render!(row, index)" />
                  <span v-else>{{ getCellValue(row, col.key) }}</span>
                </div>
              </td>
              <td v-if="$slots.action" class="px-3 py-2">
                <div class="flex items-center gap-1.5 flex-wrap">
                  <slot name="action" :row="row" :index="index" />
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Mobile cards -->
      <div class="block md:hidden divide-y divide-border">
        <div v-if="data.length === 0 && !loading" class="px-4 py-10 text-center text-muted-foreground text-sm">暂无数据</div>
        <div v-for="(row, index) in data" :key="rowKey ? rowKey(row) : index" class="px-4 py-3 space-y-2">
          <div v-for="col in columns" :key="col.key" class="flex justify-between items-start gap-2 text-sm">
            <span class="text-muted-foreground flex-shrink-0">{{ getColumnTitle(col) }}</span>
            <span class="text-foreground text-right">
              <component v-if="col.render" :is="() => col.render!(row, index)" />
              <span v-else>{{ getCellValue(row, col.key) }}</span>
            </span>
          </div>
          <div v-if="$slots.action" class="flex items-center gap-2 pt-2 border-t border-border/50 flex-wrap">
            <slot name="action" :row="row" :index="index" />
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="pageCount > 0" class="flex items-center justify-between text-[13px] text-muted-foreground">
      <span>共 {{ totalCount }} 条</span>
      <div class="flex items-center gap-1">
        <button @click="changePage(page - 1)" :disabled="page <= 1"
          class="h-7 px-2 rounded border border-border hover:bg-muted disabled:opacity-40 disabled:cursor-not-allowed transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>
        </button>
        <template v-for="p in pages" :key="p">
          <button v-if="p !== '...'" @click="changePage(p as number)"
            :class="['h-7 min-w-[1.75rem] px-1.5 rounded border text-[13px] transition-colors', page === p ? 'bg-primary text-primary-foreground border-primary' : 'border-border hover:bg-muted']">
            {{ p }}
          </button>
          <span v-else class="h-7 flex items-center px-1">…</span>
        </template>
        <button @click="changePage(page + 1)" :disabled="page >= pageCount"
          class="h-7 px-2 rounded border border-border hover:bg-muted disabled:opacity-40 disabled:cursor-not-allowed transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
        </button>
        <select :value="pageSize" @change="changePageSize(Number(($event.target as HTMLSelectElement).value))"
          class="h-7 border border-border rounded px-2 text-[13px] bg-background hover:bg-muted transition-colors cursor-pointer">
          <option v-for="s in pageSizes" :key="s" :value="s">{{ s }}/页</option>
        </select>
      </div>
    </div>
  </div>
</template>
