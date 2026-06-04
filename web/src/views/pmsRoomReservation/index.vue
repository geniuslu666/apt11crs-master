<template>
  <div class="room-order-page space-y-4">
    <div class="room-order-toolbar">
      <div class="room-order-property">
        <div class="room-order-label">物业范围</div>
        <n-select
          v-model:value="selectedPropertyId"
          filterable
          clearable
          :options="propertyOptions"
          label-field="label"
          value-field="id"
          placeholder="全部物业"
          @update:value="handlePropertySelect"
        />
      </div>
      <div class="room-order-header-actions">
        <UiButton size="sm" variant="outline" @click="loadData" :disabled="loading">
          <RefreshCw :size="14" :class="loading ? 'animate-spin' : ''" />
          刷新
        </UiButton>
      </div>
    </div>

    <div class="room-order-panel">
      <div class="room-order-panel-section">
        <div class="room-order-status-tabs">
          <button
            v-for="item in statusTabs"
            :key="item.value"
            type="button"
            :class="['room-order-status-tab', tabValue === item.value ? 'room-order-status-tab-active' : '']"
            @click="handleUpdateValue(item.value)"
          >
            {{ item.label }}
          </button>
        </div>

        <div class="room-order-filter-grid">
          <div class="room-order-field">
            <div class="room-order-label">订单来源</div>
            <UiSelect v-model="searchForm.source" :options="sourceOptions" placeholder="全部来源" clearable />
          </div>
          <div class="room-order-field">
            <div class="room-order-label">订单号</div>
            <UiInput v-model="searchForm.orderSn" placeholder="请输入订单号" @keyup.enter="handleSearch" />
          </div>
          <div class="room-order-field">
            <div class="room-order-label">Airhost订单号</div>
            <UiInput v-model="searchForm.outOrderSn" placeholder="请输入 Airhost 订单号" @keyup.enter="handleSearch" />
          </div>
          <div class="room-order-field">
            <div class="room-order-label">物业名称</div>
            <UiInput v-model="searchForm.propertyName" placeholder="请输入物业名称" @keyup.enter="handleSearch" />
          </div>
        </div>

        <div v-show="showMore" class="room-order-more-grid">
          <div class="room-order-field">
            <div class="room-order-label">订单状态</div>
            <UiSelect v-model="searchForm.orderStatus" :options="options.order_status" placeholder="全部订单状态" clearable />
          </div>
          <div class="room-order-field">
            <div class="room-order-label">预订状态</div>
            <UiSelect v-model="searchForm.status" :options="options.status" placeholder="全部预订状态" clearable />
          </div>
          <div class="room-order-field">
            <div class="room-order-label">入住状态</div>
            <UiSelect v-model="searchForm.checkinStatus" :options="options.checkin_status" placeholder="全部入住状态" clearable />
          </div>
          <div class="room-order-field">
            <div class="room-order-label">入住日期</div>
            <div class="room-order-date-range">
              <input type="date" v-model="searchForm.checkinDateStart" class="room-order-date-input" />
              <Minus :size="12" class="text-muted-foreground flex-shrink-0" />
              <input type="date" v-model="searchForm.checkinDateEnd" class="room-order-date-input" />
            </div>
          </div>
          <div class="room-order-field">
            <div class="room-order-label">退房日期</div>
            <div class="room-order-date-range">
              <input type="date" v-model="searchForm.checkoutDateStart" class="room-order-date-input" />
              <Minus :size="12" class="text-muted-foreground flex-shrink-0" />
              <input type="date" v-model="searchForm.checkoutDateEnd" class="room-order-date-input" />
            </div>
          </div>
        </div>

        <div class="room-order-filter-actions">
          <div class="flex items-center gap-2">
            <UiButton size="sm" @click="handleSearch">
              <Search :size="14" />
              查询
            </UiButton>
            <UiButton size="sm" variant="outline" @click="handleReset">
              <RotateCcw :size="14" />
              重置
            </UiButton>
          </div>
          <button type="button" class="room-order-disclosure" @click="showMore = !showMore">
            {{ showMore ? '收起筛选' : '展开筛选' }}
            <ChevronDownIcon :size="13" :class="['transition-transform duration-200', showMore ? 'rotate-180' : '']" />
          </button>
        </div>
      </div>
    </div>

    <div class="room-order-panel room-order-table-panel">
      <div class="room-order-table-toolbar">
        <div class="flex items-center gap-2">
          <BedDouble :size="16" class="text-primary" />
          <span class="text-[13px] font-semibold text-foreground">入住订单</span>
          <span v-if="totalCount > 0" class="room-order-count">{{ totalCount }}</span>
        </div>
        <button type="button" class="room-order-refresh" :disabled="loading" @click="loadData">
          <RefreshCw :size="14" :class="loading ? 'animate-spin' : ''" />
          刷新
        </button>
      </div>

      <div :class="['room-order-table-scroll', activeActionId !== null ? 'room-order-table-scroll-open' : '']">
        <div v-if="loading" class="absolute inset-0 bg-white/70 z-10 flex items-center justify-center">
          <Loader2 class="animate-spin text-primary" :size="22" />
        </div>

        <table class="room-order-table w-full text-[13px]">
          <colgroup>
            <col class="room-order-col-order" />
            <col class="room-order-col-out-order" />
            <col class="room-order-col-source" />
            <col class="room-order-col-status" />
            <col class="room-order-col-status" />
            <col class="room-order-col-guest" />
            <col class="room-order-col-property" />
            <col class="room-order-col-room" />
            <col class="room-order-col-date" />
            <col class="room-order-col-date" />
            <col class="room-order-col-status" />
            <col class="room-order-col-amount" />
            <col class="room-order-col-action" />
          </colgroup>
          <thead>
            <tr class="border-b border-border bg-[#f8fafc]">
              <th class="room-order-cell-order py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">订单号</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">Airhost订单号</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">订单来源</th>
              <th class="room-order-cell-center py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">订单状态</th>
              <th class="room-order-cell-center py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">预订状态</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">预定人</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">预定物业</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">房间信息</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">入住日期</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">退房日期</th>
              <th class="room-order-cell-center py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">入住状态</th>
              <th class="room-order-cell-amount py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">账款</th>
              <th class="room-order-cell-action py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-border">
            <tr v-if="data.length === 0 && !loading">
              <td colspan="13" class="px-3 py-14 text-center">
                <div class="room-order-empty">
                  <SearchX :size="24" />
                  <span>暂无数据</span>
                </div>
              </td>
            </tr>
            <tr v-for="row in data" :key="row.id" class="hover:bg-[#f8fafc] transition-colors">
              <td class="room-order-cell-order py-2">
                <div class="room-order-hover-tip" :data-full-text="row.orderSn || '—'">
                  <span class="room-order-single-line">{{ row.orderSn || '—' }}</span>
                </div>
              </td>
              <td class="px-3 py-2">
                <div class="room-order-hover-tip" :data-full-text="row.outOrderSn || '—'">
                  <span class="room-order-single-line">{{ row.outOrderSn || '—' }}</span>
                </div>
              </td>
              <td class="px-3 py-2 whitespace-nowrap">
                <span :class="['room-order-source-badge', getSourceClass(row)]">
                  {{ getSourceLabel(row) }}
                </span>
              </td>
              <td class="room-order-cell-center py-2">
                <span :class="['room-order-status', getStatusClass(row.orderStatus)]">
                  {{ getOptionLabel(options.order_status, row.orderStatus) || '—' }}
                </span>
              </td>
              <td class="room-order-cell-center py-2">
                <span :class="['room-order-status', getStatusClass(row.status)]">
                  {{ getOptionLabel(options.status, row.status) || '—' }}
                </span>
              </td>
              <td class="px-3 py-2">
                <div class="room-order-hover-tip" :data-full-text="getGuestName(row)">
                  <span class="room-order-single-line">{{ getGuestName(row) }}</span>
                </div>
              </td>
              <td class="px-3 py-2">
                <div class="room-order-hover-tip" :data-full-text="getPropertyName(row)">
                  <span class="room-order-single-line">{{ getPropertyName(row) }}</span>
                </div>
              </td>
              <td class="px-3 py-2">
                <div class="room-order-hover-tip" :data-full-text="getRoomInfo(row)">
                  <span class="room-order-single-line">{{ getRoomInfo(row) }}</span>
                </div>
              </td>
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">{{ formatDateOnly(row.checkinDate) }}</td>
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">{{ formatDateOnly(row.checkoutDate) }}</td>
              <td class="room-order-cell-center py-2">
                <span :class="['room-order-status', getStatusClass(row.checkinStatus)]">
                  {{ getOptionLabel(options.checkin_status, row.checkinStatus) || '—' }}
                </span>
              </td>
              <td class="room-order-cell-amount py-2 font-medium whitespace-nowrap">
                {{ formatMoney(row.bookingFee) }}
              </td>
              <td class="room-order-cell-action py-2">
                <div class="room-order-action-menu">
                  <button type="button" class="room-order-kebab" @click.stop="toggleActionMenu(row.id)">
                    <MoreHorizontal :size="19" />
                  </button>
                  <div v-if="activeActionId === row.id" class="room-order-action-popover">
                    <div class="room-order-menu-section">
                      <div class="room-order-menu-title">操作</div>
                      <button
                        v-if="row.orderSn"
                        type="button"
                        class="room-order-menu-item"
                        @click="runAction('view', row)"
                      >查看详情</button>
                      <button
                        v-if="canCancel(row)"
                        type="button"
                        class="room-order-menu-item room-order-menu-danger"
                        @click="runAction('cancel', row)"
                      >取消订单</button>
                    </div>
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="pageCount > 0" class="room-order-pagination">
        <span>共 {{ totalCount }} 条</span>
        <div class="flex items-center gap-1">
          <button @click="changePage(page - 1)" :disabled="page <= 1" class="room-order-page-btn">
            <ChevronLeft :size="13" />
          </button>
          <template v-for="p in pages" :key="p">
            <button
              v-if="p !== '...'"
              @click="changePage(p as number)"
              :class="['room-order-page-num', page === p ? 'room-order-page-num-active' : '']"
            >{{ p }}</button>
            <span v-else class="h-7 flex items-center px-1">...</span>
          </template>
          <button @click="changePage(page + 1)" :disabled="page >= pageCount" class="room-order-page-btn">
            <ChevronRight :size="13" />
          </button>
          <select
            :value="pageSize"
            class="room-order-page-size"
            @change="changePageSize(Number(($event.target as HTMLSelectElement).value))"
          >
            <option v-for="s in [10, 20, 50]" :key="s" :value="s">{{ s }}/页</option>
          </select>
        </div>
      </div>
    </div>

    <OrderView ref="orderViewRef" />
    <Cancel ref="cancelRef" @reloadTable="loadData" />
  </div>
</template>

<script lang="ts" setup>
import { computed, onBeforeUnmount, onMounted, reactive, ref, watch } from 'vue';
import { RoomList } from '@/api/pmsAppReservation';
import { loadOptions, options } from './model';
import { getOptionLabel } from '@/utils/hotgo';
import OrderView from '../pmsAppReservation/view.vue';
import Cancel from './cancel.vue';
import { useUserStore } from '@/store/modules/user';
import { storage } from '@/utils/Storage';
import { getlang } from '@/utils/smjcomm';
import { UiButton, UiInput, UiSelect } from '@/components/ui';
import {
  BedDouble,
  ChevronDown as ChevronDownIcon,
  ChevronLeft,
  ChevronRight,
  Loader2,
  Minus,
  MoreHorizontal,
  RefreshCw,
  RotateCcw,
  Search,
  SearchX,
} from '@lucide/vue';

const userStore = useUserStore();
const orderViewRef = ref();
const cancelRef = ref();
const loading = ref(false);
const data = ref<any[]>([]);
const page = ref(1);
const pageSize = ref(10);
const pageCount = ref(0);
const totalCount = ref(0);
const showMore = ref(true);
const activeActionId = ref<number | string | null>(null);
const tabValue = ref('CONFIRM');
const selectedPropertyId = ref(userStore.getuserPms.id || 0);
const rawPropertyList = ref<any[]>(storage.get('wylists') || []);

const statusTabs = [
  { label: '已确认', value: 'CONFIRM' },
  { label: '已取消', value: 'CANCEL' },
  { label: '全部', value: 'ALL' },
];

const sourceOptions = [
  { label: 'APP', value: 'APP' },
  { label: 'AIRHOST', value: 'AIRHOST' },
];

const searchForm = reactive({
  source: 'APP' as string | null,
  orderSn: '',
  outOrderSn: '',
  propertyName: '',
  orderStatus: null as string | null,
  status: null as string | null,
  checkinStatus: null as string | null,
  checkinDateStart: '',
  checkinDateEnd: '',
  checkoutDateStart: '',
  checkoutDateEnd: '',
});

const namegetlang = (data: any) => {
  if (!data) return '全部物业';
  return getlang(data, userStore.language)?.content || '全部物业';
};

const propertyOptions = computed(() => {
  const seen = new Set<number | string>();
  const list = [{ id: 0, uid: '', label: '全部物业' }];
  for (const item of rawPropertyList.value || []) {
    if (!item || seen.has(item.id) || item.id === 0) continue;
    seen.add(item.id);
    list.push({
      ...item,
      label: namegetlang(item.nameLanguage),
    });
  }
  return list;
});

const pages = computed(() => {
  if (pageCount.value <= 7) return Array.from({ length: pageCount.value }, (_, i) => i + 1);
  const arr: (number | '...')[] = [1];
  if (page.value > 3) arr.push('...');
  for (let i = Math.max(2, page.value - 1); i <= Math.min(pageCount.value - 1, page.value + 1); i++) arr.push(i);
  if (page.value < pageCount.value - 2) arr.push('...');
  if (pageCount.value > 1) arr.push(pageCount.value);
  return arr;
});

function buildParams() {
  const p: Record<string, any> = {};
  const currentPms = userStore.getuserPms;
  p.puid = currentPms?.uid || '';
  p.selectStatus = tabValue.value;
  if (searchForm.source) p.source = searchForm.source;
  if (searchForm.orderSn) p.orderSn = searchForm.orderSn;
  if (searchForm.outOrderSn) p.outOrderSn = searchForm.outOrderSn;
  if (searchForm.propertyName) p.propertyName = searchForm.propertyName;
  if (searchForm.orderStatus) p.orderStatus = searchForm.orderStatus;
  if (searchForm.status) p.status = searchForm.status;
  if (searchForm.checkinStatus) p.checkinStatus = searchForm.checkinStatus;
  if (searchForm.checkinDateStart || searchForm.checkinDateEnd) {
    p.checkinDate = [searchForm.checkinDateStart, searchForm.checkinDateEnd];
  }
  if (searchForm.checkoutDateStart || searchForm.checkoutDateEnd) {
    p.checkoutDate = [searchForm.checkoutDateStart, searchForm.checkoutDateEnd];
  }
  return p;
}

async function loadData() {
  loading.value = true;
  try {
    const res = await RoomList({ page: page.value, pageSize: pageSize.value, ...buildParams() });
    data.value = res.list || [];
    pageCount.value = res.pageCount || 0;
    totalCount.value = res.totalCount || 0;
  } finally {
    loading.value = false;
  }
}

function handleSearch() {
  page.value = 1;
  loadData();
}

function handleReset() {
  searchForm.source = 'APP';
  searchForm.orderSn = '';
  searchForm.outOrderSn = '';
  searchForm.propertyName = '';
  searchForm.orderStatus = null;
  searchForm.status = null;
  searchForm.checkinStatus = null;
  searchForm.checkinDateStart = '';
  searchForm.checkinDateEnd = '';
  searchForm.checkoutDateStart = '';
  searchForm.checkoutDateEnd = '';
  page.value = 1;
  loadData();
}

function changePage(p: number) {
  if (p < 1 || p > pageCount.value) return;
  page.value = p;
  loadData();
}

function changePageSize(s: number) {
  pageSize.value = s;
  page.value = 1;
  loadData();
}

function handleUpdateValue(value: string) {
  tabValue.value = value;
  page.value = 1;
  loadData();
}

function handlePropertySelect(value: number | null) {
  const nextValue = value || 0;
  selectedPropertyId.value = nextValue;
  const selected = propertyOptions.value.find((item) => item.id === nextValue) || propertyOptions.value[0];
  storage.set('userPms', selected);
  userStore.setuserPms(selected);
  page.value = 1;
  loadData();
}

function formatDateOnly(value: string | null | undefined): string {
  if (!value) return '—';
  return String(value).slice(0, 10) || '—';
}

function formatMoney(value: number | string | null | undefined): string {
  if (value === null || value === undefined || value === '') return '—';
  return `${Number(value).toLocaleString()} JPY`;
}

function getGuestName(row: any): string {
  return row.guestProfileDetail?.fullName || row.mainGuestDetail?.fullName || row.guestProfile?.fullName || '—';
}

function getPropertyName(row: any): string {
  return row.propertyDetail?.name || row.propertyName || '—';
}

function getRoomInfo(row: any): string {
  const roomNo = row.roomUnitDetail?.roomNo ? `#${row.roomUnitDetail.roomNo}` : '';
  const roomType = row.roomTypeDetail?.name || '';
  const text = `${roomNo} ${roomType}`.trim();
  return text || '—';
}

function getSourceLabel(row: any): string {
  const source = String(row.sourceName || row.source || '').toUpperCase();
  if (source.includes('AIRHOST')) return 'Air';
  if (source.includes('APP')) return 'APP';
  return source || '—';
}

function getSourceClass(row: any): string {
  const label = getSourceLabel(row);
  if (label === 'Air') return 'room-order-source-air';
  if (label === 'APP') return 'room-order-source-app';
  return 'room-order-source-muted';
}

function getStatusClass(status: string | null | undefined): string {
  if (status === 'DONE' || status === 'CONFIRM' || status === 'confirmed' || status === 'checked_in') return 'room-order-status-ok';
  if (status === 'CANCEL' || status === 'cancelled' || status === 'FAILED') return 'room-order-status-danger';
  return 'room-order-status-muted';
}

function canCancel(row: any): boolean {
  return row.status !== 'cancelled' && row.status !== 'CANCEL';
}

function toggleActionMenu(id: number | string) {
  activeActionId.value = activeActionId.value === id ? null : id;
}

function closeActionMenu() {
  activeActionId.value = null;
}

function handleDocumentClick(event: MouseEvent) {
  const target = event.target as HTMLElement;
  if (!target.closest('.room-order-action-menu')) closeActionMenu();
}

function runAction(key: string, record: any) {
  closeActionMenu();
  if (key === 'cancel') return handleCancel(record);
  return handleView(record);
}

function handleCancel(record: Recordable) {
  cancelRef.value.openModal(record);
}

function handleView(record: Recordable) {
  record.viewtype = '住宿订单';
  orderViewRef.value.openModal(record);
}

watch(
  () => userStore.getuserPms,
  () => {
    selectedPropertyId.value = userStore.getuserPms.id || 0;
  },
  { deep: true }
);

onMounted(() => {
  loadOptions();
  loadData();
  document.addEventListener('click', handleDocumentClick);
});

onBeforeUnmount(() => {
  document.removeEventListener('click', handleDocumentClick);
});
</script>

<style lang="less" scoped>
.room-order-page {
  --crs-primary: #38aeea;
  --crs-primary-strong: #128fc8;
  --crs-primary-soft: #eaf7ff;
  --crs-border: #dce7f2;
  --crs-surface: #ffffff;
  --crs-text: #152033;
  --crs-muted: #6b7c93;
  --crs-danger: #ef4444;
  --crs-danger-soft: #fff1f2;
  --crs-success: #16a34a;
  --crs-success-soft: #eefbf3;
  --crs-shadow: 0 1px 2px rgba(15, 23, 42, 0.04), 0 0 0 1px rgba(220, 231, 242, 0.5);
  background: #f6f9fc;
  color: var(--crs-text);
  min-height: calc(100vh - var(--header-height, 56px));
  padding: 16px;

  .room-order-toolbar {
    align-items: flex-end;
    display: flex;
    gap: 12px;
    justify-content: space-between;
  }

  .room-order-property {
    max-width: 360px;
    min-width: 260px;
    width: 32vw;
  }

  .room-order-header-actions {
    align-items: center;
    display: flex;
    gap: 8px;
  }

  .room-order-panel {
    background: var(--crs-surface);
    border: 1px solid var(--crs-border);
    border-radius: 8px;
    box-shadow: var(--crs-shadow);
  }

  .room-order-panel-section {
    padding: 16px;
  }

  .room-order-status-tabs {
    align-items: center;
    display: inline-flex;
    gap: 4px;
    margin-bottom: 14px;
  }

  .room-order-status-tab {
    border-radius: 999px;
    color: var(--crs-muted);
    font-size: 13px;
    height: 28px;
    padding: 0 12px;
    transition: background 0.15s ease, color 0.15s ease;
  }

  .room-order-status-tab:hover {
    background: #f1f5f9;
    color: var(--crs-text);
  }

  .room-order-status-tab-active {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
    font-weight: 700;
  }

  .room-order-filter-grid {
    display: grid;
    gap: 12px;
    grid-template-columns: repeat(4, minmax(0, 1fr));
  }

  .room-order-more-grid {
    display: grid;
    gap: 12px;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    margin-top: 12px;
  }

  .room-order-field {
    min-width: 0;
  }

  .room-order-label {
    color: var(--crs-muted);
    font-size: 12px;
    font-weight: 600;
    line-height: 18px;
    margin-bottom: 4px;
  }

  .room-order-date-range {
    align-items: center;
    display: flex;
    gap: 6px;
  }

  .room-order-date-input {
    background: hsl(var(--background));
    border: 1px solid hsl(var(--input));
    border-radius: 6px;
    color: var(--crs-text);
    flex: 1 1 0;
    font-size: 13px;
    height: 32px;
    min-width: 0;
    padding: 0 8px;
  }

  .room-order-filter-actions {
    align-items: center;
    display: flex;
    justify-content: space-between;
    margin-top: 14px;
  }

  .room-order-disclosure,
  .room-order-refresh {
    align-items: center;
    border-radius: 6px;
    color: var(--crs-muted);
    display: inline-flex;
    font-size: 12px;
    gap: 4px;
    height: 28px;
    padding: 0 8px;
    transition: color 0.15s ease, background 0.15s ease;
  }

  .room-order-disclosure:hover,
  .room-order-refresh:hover {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
  }

  .room-order-table-panel {
    overflow: visible;
  }

  .room-order-table-toolbar {
    align-items: center;
    border-bottom: 1px solid var(--crs-border);
    display: flex;
    justify-content: space-between;
    padding: 10px 16px 10px 24px;
  }

  .room-order-table-scroll {
    overflow-x: auto;
    overflow-y: visible;
    position: relative;
  }

  .room-order-table-scroll-open {
    padding-bottom: 180px;
  }

  .room-order-table {
    min-width: 1400px;
    table-layout: fixed;
  }

  .room-order-col-order {
    width: 150px;
  }

  .room-order-col-out-order {
    width: 150px;
  }

  .room-order-col-source,
  .room-order-col-status {
    width: 86px;
  }

  .room-order-col-guest {
    width: 104px;
  }

  .room-order-col-property,
  .room-order-col-room {
    width: 150px;
  }

  .room-order-col-date {
    width: 104px;
  }

  .room-order-col-amount {
    width: 96px;
  }

  .room-order-col-action {
    width: 64px;
  }

  .room-order-cell-order {
    padding-left: 24px;
    padding-right: 10px;
  }

  .room-order-cell-center {
    padding-left: 8px;
    padding-right: 8px;
    text-align: center;
  }

  .room-order-cell-amount {
    padding-left: 8px;
    padding-right: 12px;
    text-align: left;
  }

  .room-order-cell-action {
    background: var(--crs-surface);
    box-shadow: -10px 0 16px -16px rgba(15, 23, 42, 0.35);
    padding-left: 8px;
    padding-right: 8px;
    position: sticky;
    right: 0;
    text-align: center;
    z-index: 6;
  }

  thead .room-order-cell-action {
    background: #f8fafc;
    z-index: 8;
  }

  tbody tr:hover .room-order-cell-action {
    background: #f8fafc;
  }

  .room-order-hover-tip {
    display: block;
    max-width: 100%;
    position: relative;
  }

  .room-order-single-line {
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .room-order-hover-tip:hover::after {
    background: var(--crs-text);
    border-radius: 6px;
    box-shadow: 0 10px 24px rgba(15, 23, 42, 0.16);
    color: #fff;
    content: attr(data-full-text);
    font-size: 12px;
    font-weight: 500;
    left: 0;
    line-height: 18px;
    max-width: 280px;
    overflow-wrap: break-word;
    padding: 6px 8px;
    position: absolute;
    top: calc(100% + 8px);
    white-space: normal;
    width: max-content;
    z-index: 90;
  }

  .room-order-status {
    align-items: center;
    border-radius: 4px;
    display: inline-flex;
    font-size: 11px;
    font-weight: 600;
    height: 20px;
    justify-content: center;
    max-width: 72px;
    padding: 0 6px;
    white-space: nowrap;
  }

  .room-order-status-ok {
    background: var(--crs-success-soft);
    color: var(--crs-success);
  }

  .room-order-status-danger {
    background: var(--crs-danger-soft);
    color: var(--crs-danger);
  }

  .room-order-status-muted {
    background: #f1f5f9;
    color: var(--crs-muted);
  }

  .room-order-source-badge {
    align-items: center;
    border-radius: 4px;
    display: inline-flex;
    font-size: 11px;
    font-weight: 700;
    height: 20px;
    justify-content: center;
    padding: 0 7px;
    white-space: nowrap;
  }

  .room-order-source-air {
    background: #eaf7ff;
    color: #128fc8;
  }

  .room-order-source-app {
    background: #eefbf3;
    color: #16a34a;
  }

  .room-order-source-muted {
    background: #f1f5f9;
    color: var(--crs-muted);
  }

  .room-order-count {
    align-items: center;
    background: var(--crs-primary-soft);
    border-radius: 999px;
    color: var(--crs-primary-strong);
    display: inline-flex;
    font-size: 11px;
    font-weight: 700;
    height: 20px;
    padding: 0 7px;
  }

  .room-order-empty {
    align-items: center;
    color: var(--crs-muted);
    display: inline-flex;
    flex-direction: column;
    font-size: 13px;
    gap: 8px;
  }

  .room-order-action-menu {
    display: inline-flex;
    justify-content: center;
    position: relative;
    width: 100%;
  }

  .room-order-kebab {
    align-items: center;
    background: transparent;
    border: 0;
    border-radius: 8px;
    color: var(--crs-muted);
    display: inline-flex;
    height: 32px;
    justify-content: center;
    padding: 0;
    width: 36px;
  }

  .room-order-kebab:hover,
  .room-order-kebab:focus-visible {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
    box-shadow: 0 0 0 3px rgba(56, 174, 234, 0.14);
  }

  .room-order-action-popover {
    background: var(--crs-surface);
    border: 1px solid var(--crs-border);
    border-radius: 8px;
    box-shadow: 0 18px 36px rgba(15, 23, 42, 0.12);
    color: var(--crs-text);
    min-width: 148px;
    padding: 8px 0;
    position: absolute;
    right: 0;
    text-align: left;
    top: calc(100% + 8px);
    z-index: 80;
  }

  .room-order-menu-section {
    display: grid;
    gap: 2px;
  }

  .room-order-menu-title {
    color: var(--crs-muted);
    font-size: 12px;
    font-weight: 700;
    line-height: 18px;
    padding: 2px 14px 4px;
  }

  .room-order-menu-item {
    color: var(--crs-text);
    display: block;
    font-size: 13px;
    line-height: 20px;
    padding: 7px 14px;
    text-align: left;
    transition: background 0.15s ease, color 0.15s ease;
    width: 100%;
  }

  .room-order-menu-item:hover {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
  }

  .room-order-menu-danger {
    color: var(--crs-danger);
  }

  .room-order-menu-danger:hover {
    background: var(--crs-danger-soft);
    color: var(--crs-danger);
  }

  .room-order-pagination {
    align-items: center;
    border-top: 1px solid var(--crs-border);
    color: var(--crs-muted);
    display: flex;
    font-size: 13px;
    justify-content: space-between;
    padding: 12px 16px;
  }

  .room-order-page-btn,
  .room-order-page-num {
    align-items: center;
    border: 1px solid var(--crs-border);
    border-radius: 6px;
    display: inline-flex;
    height: 28px;
    justify-content: center;
    min-width: 28px;
    padding: 0 7px;
    transition: background 0.15s ease, color 0.15s ease;
  }

  .room-order-page-btn:hover,
  .room-order-page-num:hover {
    background: #f1f5f9;
  }

  .room-order-page-btn:disabled {
    cursor: not-allowed;
    opacity: 0.45;
  }

  .room-order-page-num-active {
    background: var(--crs-primary) !important;
    border-color: var(--crs-primary);
    color: #fff;
  }

  .room-order-page-size {
    background: #fff;
    border: 1px solid var(--crs-border);
    border-radius: 6px;
    height: 28px;
    padding: 0 8px;
  }

  @media (max-width: 1024px) {
    .room-order-filter-grid,
    .room-order-more-grid {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  @media (max-width: 640px) {
    .room-order-toolbar,
    .room-order-filter-actions,
    .room-order-pagination {
      align-items: stretch;
      flex-direction: column;
      gap: 10px;
    }

    .room-order-property {
      max-width: none;
      min-width: 0;
      width: 100%;
    }

    .room-order-filter-grid,
    .room-order-more-grid {
      grid-template-columns: 1fr;
    }
  }
}
</style>
