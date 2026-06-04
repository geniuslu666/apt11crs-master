<template>
  <div class="dashboard-shell flex flex-1 flex-col bg-background text-foreground">
    <header class="site-header sticky top-0 z-10 flex h-14 shrink-0 items-center gap-2 border-b bg-background/95 px-4 backdrop-blur lg:px-6">
      <button class="icon-button" type="button" aria-label="Toggle sidebar">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M4 6h16M4 12h16M4 18h16" />
        </svg>
      </button>
      <div class="h-4 w-px bg-border"></div>
      <nav class="hidden items-center gap-2 text-sm text-muted-foreground md:flex">
        <span>{{ translang('主控台') }}</span>
        <svg class="size-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="m9 18 6-6-6-6" />
        </svg>
        <span class="font-medium text-foreground">Dashboard</span>
      </nav>
      <div class="ml-auto flex items-center gap-2">
        <div class="hidden text-right sm:block">
          <p class="text-sm font-medium leading-none">{{ namegetlang(userStore.userPms.nameLanguage) }}</p>
          <p class="mt-1 text-xs text-muted-foreground">{{ currentDateLabel }}</p>
        </div>
        <input
          v-model="dateValue"
          type="date"
          class="date-input"
          @change="onDateChange"
        />
      </div>
    </header>

    <main class="flex flex-1 flex-col gap-4 py-4 md:gap-6 md:py-6">
      <section class="px-4 lg:px-6">
        <div class="flex flex-col gap-1">
          <h1 class="text-lg font-semibold tracking-tight md:text-xl">
            {{ translang('欢迎') }}，{{ namegetlang(userStore.userPms.nameLanguage) }}！
          </h1>
          <p class="text-sm text-muted-foreground">A dashboard with sidebar, charts and data table.</p>
        </div>
      </section>

      <section class="grid grid-cols-1 gap-4 px-4 sm:grid-cols-2 lg:grid-cols-4 lg:px-6">
        <article v-for="card in sectionCards" :key="card.title" class="metric-card">
          <div class="grid grid-cols-[1fr_auto] gap-1 px-6">
            <p class="text-sm text-muted-foreground">{{ card.title }}</p>
            <p class="text-2xl font-semibold tabular-nums lg:text-3xl">{{ card.value }}</p>
            <span class="status-pill">
              <svg class="size-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path :d="card.up ? 'M23 6 13.5 15.5 8.5 10.5 1 18M17 6h6v6' : 'M23 18 13.5 8.5 8.5 13.5 1 6M17 18h6v-6'" />
              </svg>
              {{ card.badge }}
            </span>
          </div>
          <div class="px-6">
            <div class="flex items-center gap-2 text-sm font-medium">
              {{ card.caption }}
              <svg class="size-4 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path :d="card.up ? 'M23 6 13.5 15.5 8.5 10.5 1 18M17 6h6v6' : 'M23 18 13.5 8.5 8.5 13.5 1 6M17 18h6v-6'" />
              </svg>
            </div>
            <p class="mt-1.5 text-sm text-muted-foreground">{{ card.description }}</p>
          </div>
        </article>
      </section>

      <section class="px-4 lg:px-6">
        <article class="dashboard-card">
          <div class="flex flex-col gap-3 p-5 pb-0 sm:flex-row sm:items-center sm:justify-between sm:p-6 sm:pb-0">
            <div>
              <h2 class="text-base font-semibold">Total Visitors</h2>
              <p class="text-sm text-muted-foreground">Total for the last 3 months</p>
            </div>
            <div class="range-tabs">
              <button
                v-for="r in timeRanges"
                :key="r.value"
                type="button"
                :class="['range-tab', activeRange === r.value && 'is-active']"
                @click="activeRange = r.value; updateChart()"
              >
                {{ r.label }}
              </button>
            </div>
          </div>
          <div ref="chartRef" class="h-[250px] w-full px-2 pt-4 sm:px-4 sm:pt-6"></div>
        </article>
      </section>

      <section class="px-4 lg:px-6">
        <article class="dashboard-card overflow-hidden">
          <div class="flex flex-col gap-3 border-b px-5 py-4 sm:flex-row sm:items-center sm:justify-between sm:px-6">
            <div>
              <h2 class="text-base font-semibold">Tasks</h2>
              <p class="text-sm text-muted-foreground">Drag and drop your tasks to update your workflow.</p>
            </div>
            <button class="outline-button" type="button">
              <svg class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 5v14M5 12h14" />
              </svg>
              Add Section
            </button>
          </div>
          <div class="overflow-x-auto">
            <table class="data-table">
              <thead>
                <tr>
                  <th class="w-10"><input type="checkbox" /></th>
                  <th>Header</th>
                  <th>Section Type</th>
                  <th>Status</th>
                  <th>Target</th>
                  <th>Limit</th>
                  <th>Reviewer</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="row in tableRows" :key="row.id">
                  <td><input type="checkbox" /></td>
                  <td class="min-w-[220px] font-medium">{{ row.header }}</td>
                  <td>{{ row.type }}</td>
                  <td><span :class="['table-badge', row.statusClass]">{{ row.status }}</span></td>
                  <td>{{ row.target }}</td>
                  <td>{{ row.limit }}</td>
                  <td>{{ row.reviewer }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </article>
      </section>

      <section class="grid grid-cols-1 gap-4 px-4 lg:grid-cols-2 lg:px-6">
        <article class="dashboard-card">
          <div class="border-b px-5 py-4 sm:px-6">
            <h2 class="text-base font-semibold">自助入住报告</h2>
          </div>
          <div class="grid grid-cols-1 gap-3 p-5 sm:grid-cols-3 sm:p-6">
            <div v-for="item in checkinItems" :key="item.label" class="report-cell">
              <div class="flex items-center gap-2">
                <span :class="['size-2 rounded-full', item.dotColor]"></span>
                <p class="text-xs text-muted-foreground">{{ item.label }}</p>
              </div>
              <p class="text-xl font-semibold tabular-nums">{{ item.value }}</p>
              <div class="h-1.5 overflow-hidden rounded-full bg-border">
                <div :class="['h-full rounded-full', item.barColor]" :style="{ width: item.pct + '%' }"></div>
              </div>
              <p class="text-xs text-muted-foreground">总数: {{ item.total }}</p>
            </div>
          </div>
        </article>

        <article class="dashboard-card">
          <div class="border-b px-5 py-4 sm:px-6">
            <h2 class="text-base font-semibold">清洁任务</h2>
          </div>
          <div class="grid grid-cols-2 gap-3 p-5 sm:p-6">
            <div v-for="item in cleaningItems" :key="item.label" class="cleaning-cell">
              <span :class="['absolute left-0 top-0 rounded-br-md px-2.5 py-0.5 text-xs font-medium', item.badgeClass]">
                {{ item.label }}
              </span>
              <p class="mt-4 text-xl font-semibold tabular-nums">暂无数据</p>
            </div>
          </div>
        </article>
      </section>
    </main>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref, watch } from 'vue';
import type { Ref } from 'vue';
import { dashboard } from '@/api/comm';
import { useECharts } from '@/hooks/web/useECharts';
import { useUserStore } from '@/store/modules/user';
import { formatToDateTime } from '@/utils/dateUtil';
import { getlang, translang } from '@/utils/smjcomm';

const userStore = useUserStore();
const today = new Date();
const dateValue = ref(formatToDateTime(today, 'yyyy-MM-dd'));

const currentDateLabel = computed(() => {
  const d = new Date(`${dateValue.value}T00:00:00`);
  return d.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric', weekday: 'long' });
});

const stats = reactive({
  import_bookings: 0,
  export_bookings: 0,
  checked_bookings: 0,
  checkout_bookings: 0,
  all_room: 0,
  today_stays: 0,
});

const sectionCards = computed(() => [
  {
    title: '今日预抵',
    value: stats.import_bookings,
    badge: '+12.5%',
    caption: `已办理入住 ${stats.checked_bookings} 间`,
    description: 'Trending up this month',
    up: true,
  },
  {
    title: '今日预离',
    value: stats.export_bookings,
    badge: '-20%',
    caption: `已办理退房 ${stats.checkout_bookings} 间`,
    description: 'Down 20% this period',
    up: false,
  },
  {
    title: '在住客房',
    value: stats.checked_bookings,
    badge: '+12.5%',
    caption: `客房总数 ${stats.all_room} 间`,
    description: 'Strong user retention',
    up: true,
  },
  {
    title: '今日新增预定',
    value: stats.today_stays,
    badge: '+4.5%',
    caption: '预定持续增长',
    description: 'Meets growth projections',
    up: true,
  },
]);

const tableRows = computed(() => [
  { id: 1, header: '今日预抵房态核对', type: 'Section', status: 'In Process', statusClass: 'is-process', target: stats.import_bookings, limit: stats.all_room, reviewer: 'Assign reviewer' },
  { id: 2, header: '自助入住完成率', type: 'Table of Contents', status: 'Done', statusClass: 'is-done', target: stats.checked_bookings, limit: stats.import_bookings, reviewer: namegetlang(userStore.userPms.nameLanguage) },
  { id: 3, header: '今日预离订单清单', type: 'Narrative', status: 'Done', statusClass: 'is-done', target: stats.checkout_bookings, limit: stats.export_bookings, reviewer: 'Assign reviewer' },
  { id: 4, header: '客房清洁任务同步', type: 'Checklist', status: 'In Process', statusClass: 'is-process', target: 0, limit: stats.all_room, reviewer: 'Assign reviewer' },
  { id: 5, header: '今日新增预订复盘', type: 'Performance', status: 'Not Started', statusClass: 'is-muted', target: stats.today_stays, limit: stats.all_room, reviewer: 'Assign reviewer' },
]);

async function Load() {
  const res = await dashboard({
    date: dateValue.value,
    puid: userStore.getuserPms?.uid ?? '',
  });
  if (res?.Details) {
    Object.assign(stats, res.Details);
  }
}

function onDateChange() {
  Load();
}

const namegetlang = (data: any) => {
  if (data) return getlang(data, userStore.language)?.content ?? '暂无物业名';
  return '暂无物业名';
};

const checkinItems = [
  { label: '已提前入住', value: '暂无', total: '暂无', pct: 0, dotColor: 'bg-primary', barColor: 'bg-primary' },
  { label: '已入住', value: '暂无', total: '暂无', pct: 0, dotColor: 'bg-emerald-500', barColor: 'bg-emerald-500' },
  { label: '已退房', value: '暂无', total: '暂无', pct: 0, dotColor: 'bg-rose-500', barColor: 'bg-rose-500' },
];

const cleaningItems = [
  { label: '未分配', badgeClass: 'bg-rose-100 text-rose-700' },
  { label: '待定', badgeClass: 'bg-amber-100 text-amber-700' },
  { label: '准备清洁', badgeClass: 'bg-teal-100 text-teal-700' },
  { label: '已确认', badgeClass: 'bg-blue-100 text-blue-700' },
];

const chartRef = ref<HTMLDivElement | null>(null);
const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);

const timeRanges = [
  { label: 'Last 3 months', value: '90d' },
  { label: 'Last 30 days', value: '30d' },
  { label: 'Last 7 days', value: '7d' },
];
const activeRange = ref('90d');

const chartRawData = [
  { date: '2024-04-01', desktop: 42, mobile: 24 },
  { date: '2024-04-08', desktop: 58, mobile: 37 },
  { date: '2024-04-15', desktop: 35, mobile: 29 },
  { date: '2024-04-22', desktop: 67, mobile: 42 },
  { date: '2024-04-29', desktop: 53, mobile: 38 },
  { date: '2024-05-06', desktop: 72, mobile: 45 },
  { date: '2024-05-13', desktop: 89, mobile: 58 },
  { date: '2024-05-20', desktop: 61, mobile: 40 },
  { date: '2024-05-27', desktop: 95, mobile: 64 },
  { date: '2024-06-03', desktop: 78, mobile: 51 },
  { date: '2024-06-10', desktop: 112, mobile: 76 },
  { date: '2024-06-17', desktop: 88, mobile: 66 },
  { date: '2024-06-24', desktop: 134, mobile: 82 },
  { date: '2024-06-30', desktop: 105, mobile: 71 },
];

function updateChart() {
  const days = activeRange.value === '7d' ? 7 : activeRange.value === '30d' ? 30 : 90;
  const cutoff = new Date('2024-06-30');
  cutoff.setDate(cutoff.getDate() - days);
  const filtered = chartRawData.filter((d) => new Date(d.date) >= cutoff);
  const xData = filtered.map((d) => {
    const dt = new Date(d.date);
    return `${dt.getMonth() + 1}/${dt.getDate()}`;
  });

  setOptions({
    color: ['#2563eb', '#60a5fa'],
    grid: { left: '1%', right: '1%', top: '8%', bottom: '8%', containLabel: true },
    tooltip: {
      trigger: 'axis',
      axisPointer: { type: 'line', lineStyle: { color: 'hsl(var(--border))' } },
      backgroundColor: 'hsl(var(--background))',
      borderColor: 'hsl(var(--border))',
      textStyle: { color: 'hsl(var(--foreground))', fontSize: 12 },
    },
    legend: { show: false },
    xAxis: {
      type: 'category',
      data: xData,
      axisLine: { show: false },
      axisTick: { show: false },
      axisLabel: { color: 'hsl(var(--muted-foreground))', fontSize: 11, margin: 10 },
      boundaryGap: false,
    },
    yAxis: {
      type: 'value',
      axisLine: { show: false },
      axisTick: { show: false },
      axisLabel: { color: 'hsl(var(--muted-foreground))', fontSize: 11 },
      splitLine: { lineStyle: { color: 'hsl(var(--border))' } },
    },
    series: [
      {
        name: 'Desktop',
        type: 'line',
        data: filtered.map((d) => d.desktop),
        smooth: true,
        symbol: 'none',
        lineStyle: { width: 2 },
        areaStyle: { opacity: 0.14 },
      },
      {
        name: 'Mobile',
        type: 'line',
        data: filtered.map((d) => d.mobile),
        smooth: true,
        symbol: 'none',
        lineStyle: { width: 2 },
        areaStyle: { opacity: 0.1 },
      },
    ],
  });
}

watch(
  () => userStore.getuserPms,
  () => Load(),
  { immediate: true, deep: true }
);

onMounted(() => {
  Load();
  updateChart();
});
</script>

<style scoped>
.dashboard-shell {
  --background: 0 0% 100%;
  --foreground: 222.2 84% 4.9%;
  --muted: 210 40% 96.1%;
  --muted-foreground: 215.4 16.3% 46.9%;
  --card: 0 0% 100%;
  --border: 214.3 31.8% 91.4%;
  --input: 214.3 31.8% 91.4%;
  --primary: 221.2 83.2% 53.3%;
  background-color: hsl(var(--background));
  color: hsl(var(--foreground));
}

.text-muted-foreground {
  color: hsl(var(--muted-foreground));
}

.border-b {
  border-bottom: 1px solid hsl(var(--border));
}

.border {
  border: 1px solid hsl(var(--border));
}

.bg-background {
  background-color: hsl(var(--background));
}

.bg-card {
  background-color: hsl(var(--card));
}

.bg-border {
  background-color: hsl(var(--border));
}

.icon-button {
  display: inline-flex;
  width: 2rem;
  height: 2rem;
  align-items: center;
  justify-content: center;
  border-radius: 0.375rem;
  color: hsl(var(--foreground));
}

.icon-button svg {
  width: 1rem;
  height: 1rem;
}

.date-input,
.outline-button {
  display: inline-flex;
  height: 2rem;
  align-items: center;
  gap: 0.5rem;
  border: 1px solid hsl(var(--input));
  border-radius: 0.375rem;
  background: transparent;
  padding: 0 0.75rem;
  font-size: 0.8125rem;
  font-weight: 500;
}

.outline-button svg {
  width: 1rem;
  height: 1rem;
}

.metric-card,
.dashboard-card {
  border: 1px solid hsl(var(--border));
  border-radius: 0.75rem;
  background: hsl(var(--card));
  box-shadow: 0 1px 2px rgb(15 23 42 / 0.04);
}

.metric-card {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  padding: 1.5rem 0;
  background-image: linear-gradient(to top, hsl(var(--primary) / 0.05), hsl(var(--card)));
}

.status-pill {
  grid-column-start: 2;
  grid-row: 1 / span 2;
  display: inline-flex;
  align-self: start;
  align-items: center;
  gap: 0.25rem;
  justify-self: end;
  border: 1px solid hsl(var(--border));
  border-radius: 999px;
  padding: 0.125rem 0.625rem;
  font-size: 0.75rem;
  font-weight: 500;
}

.range-tabs {
  display: inline-flex;
  align-items: center;
  gap: 0.125rem;
  border: 1px solid hsl(var(--border));
  border-radius: 0.5rem;
  padding: 0.125rem;
}

.range-tab {
  border-radius: 0.375rem;
  padding: 0.375rem 0.75rem;
  color: hsl(var(--muted-foreground));
  font-size: 0.8125rem;
  font-weight: 500;
  transition: color 0.15s ease, background-color 0.15s ease, box-shadow 0.15s ease;
}

.range-tab.is-active {
  background: hsl(var(--background));
  color: hsl(var(--foreground));
  box-shadow: 0 1px 2px rgb(15 23 42 / 0.08);
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.875rem;
}

.data-table th,
.data-table td {
  border-bottom: 1px solid hsl(var(--border));
  padding: 0.75rem 1rem;
  text-align: left;
  white-space: nowrap;
}

.data-table th {
  color: hsl(var(--muted-foreground));
  font-weight: 500;
}

.data-table tbody tr:hover {
  background: hsl(var(--muted) / 0.5);
}

.table-badge {
  display: inline-flex;
  align-items: center;
  border-radius: 999px;
  border: 1px solid hsl(var(--border));
  padding: 0.125rem 0.5rem;
  font-size: 0.75rem;
  font-weight: 500;
}

.table-badge.is-done {
  color: #047857;
}

.table-badge.is-process {
  color: #1d4ed8;
}

.table-badge.is-muted {
  color: hsl(var(--muted-foreground));
}

.report-cell {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  border-radius: 0.5rem;
  background: hsl(var(--muted) / 0.55);
  padding: 0.75rem;
}

.cleaning-cell {
  position: relative;
  display: flex;
  min-height: 6rem;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border-radius: 0.5rem;
  background: hsl(var(--muted) / 0.45);
}
</style>
