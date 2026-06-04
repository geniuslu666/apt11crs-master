<template>
  <div class="reservation-page space-y-4">
    <div class="reservation-toolbar">
      <div class="reservation-property">
        <div class="reservation-label">物业范围</div>
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
      <div class="reservation-header-actions">
        <UiButton size="sm" variant="outline" @click="loadData" :disabled="loading">
          <RefreshCw :size="14" :class="loading ? 'animate-spin' : ''" />
          刷新
        </UiButton>
        <UiButton
          size="sm"
          variant="outline"
          @click="handleExport"
          v-if="hasPermission(['/pmsAppReservation/export'])"
        >
          <Download :size="14" />
          导出
        </UiButton>
      </div>
    </div>

    <div class="reservation-panel">
      <div class="reservation-panel-section">
        <div class="reservation-filter-grid">
          <div class="reservation-field">
            <div class="reservation-label">订单号</div>
            <UiInput v-model="searchForm.snKeyword" placeholder="系统单号/外部单号" @keyup.enter="handleSearch" />
          </div>
          <div class="reservation-field">
            <div class="reservation-label">预订状态</div>
            <UiSelect v-model="searchForm.orderStatus" :options="options.order_status" placeholder="全部状态" clearable />
          </div>
          <div class="reservation-field">
            <div class="reservation-label">退款状态</div>
            <UiSelect v-model="searchForm.refund_status" :options="options.refund_status" placeholder="全部退款状态" clearable />
          </div>
          <div class="reservation-field">
            <div class="reservation-label">会员号</div>
            <UiInput v-model="searchForm.MemberNo" placeholder="请输入会员号" @keyup.enter="handleSearch" />
          </div>
        </div>

        <div v-show="showMore" class="reservation-more-grid">
          <div class="reservation-field">
            <div class="reservation-label">入住日期</div>
            <div class="reservation-date-range">
              <input type="date" v-model="searchForm.checkinDateStart" class="reservation-date-input" />
              <Minus :size="12" class="text-muted-foreground flex-shrink-0" />
              <input type="date" v-model="searchForm.checkinDateEnd" class="reservation-date-input" />
            </div>
          </div>
          <div class="reservation-field">
            <div class="reservation-label">退房日期</div>
            <div class="reservation-date-range">
              <input type="date" v-model="searchForm.checkoutDateStart" class="reservation-date-input" />
              <Minus :size="12" class="text-muted-foreground flex-shrink-0" />
              <input type="date" v-model="searchForm.checkoutDateEnd" class="reservation-date-input" />
            </div>
          </div>
          <div class="reservation-field">
            <div class="reservation-label">下单时间</div>
            <div class="reservation-date-range">
              <input type="date" v-model="searchForm.createdAtStart" class="reservation-date-input" />
              <Minus :size="12" class="text-muted-foreground flex-shrink-0" />
              <input type="date" v-model="searchForm.createdAtEnd" class="reservation-date-input" />
            </div>
          </div>
        </div>

        <div class="reservation-filter-actions">
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
          <button type="button" class="reservation-disclosure" @click="showMore = !showMore">
            {{ showMore ? '收起筛选' : '展开筛选' }}
            <ChevronDownIcon :size="13" :class="['transition-transform duration-200', showMore ? 'rotate-180' : '']" />
          </button>
        </div>
      </div>
    </div>

    <div class="reservation-panel reservation-table-panel">
      <div class="reservation-table-toolbar">
        <div class="flex items-center gap-2">
          <CalendarCheck :size="16" class="text-primary" />
          <span class="text-[13px] font-semibold text-foreground">预订订单</span>
          <span v-if="totalCount > 0" class="reservation-count">{{ totalCount }}</span>
        </div>
        <button type="button" class="reservation-refresh" :disabled="loading" @click="loadData">
          <RefreshCw :size="14" :class="loading ? 'animate-spin' : ''" />
          刷新
        </button>
      </div>

      <div :class="['reservation-table-scroll', activeActionId !== null ? 'reservation-table-scroll-open' : '']">
        <div v-if="loading" class="absolute inset-0 bg-white/70 z-10 flex items-center justify-center">
          <Loader2 class="animate-spin text-primary" :size="22" />
        </div>

        <table class="reservation-table w-full text-[13px]">
          <colgroup>
            <col class="reservation-col-order" />
            <col class="reservation-col-status" />
            <col class="reservation-col-guest" />
            <col class="reservation-col-contact" />
            <col class="reservation-col-property" />
            <col class="reservation-col-date" />
            <col class="reservation-col-date" />
            <col class="reservation-col-amount" />
            <col class="reservation-col-date" />
            <col class="reservation-col-refund-status" />
            <col class="reservation-col-amount" />
            <col class="reservation-col-action" />
          </colgroup>
          <thead>
            <tr class="border-b border-border bg-[#f8fafc]">
              <th class="reservation-cell-order py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">订单号</th>
              <th class="reservation-cell-center py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">预订状态</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">预定人</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">联系方式</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">预定物业</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">入住日期</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">退房日期</th>
              <th class="reservation-cell-amount py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">订单金额</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">下单时间</th>
              <th class="reservation-cell-center py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">退款状态</th>
              <th class="reservation-cell-amount py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">退款金额</th>
              <th class="reservation-cell-action py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-border">
            <tr v-if="data.length === 0 && !loading">
              <td colspan="12" class="px-3 py-14 text-center">
                <div class="reservation-empty">
                  <SearchX :size="24" />
                  <span>暂无数据</span>
                </div>
              </td>
            </tr>
            <tr v-for="row in data" :key="row.id" class="hover:bg-[#f8fafc] transition-colors">
              <td class="reservation-cell-order py-2">
                <div class="reservation-hover-tip" :data-full-text="row.orderSn || row.outOrderSn || '—'">
                  <span class="reservation-single-line">{{ row.orderSn || row.outOrderSn || '—' }}</span>
                </div>
              </td>
              <td class="reservation-cell-center py-2">
                <span :class="['reservation-status', getStatusClass(row.orderStatus)]">
                  {{ getOptionLabel(options.order_status, row.orderStatus) || '—' }}
                </span>
              </td>
              <td class="px-3 py-2">
                <div class="reservation-hover-tip" :data-full-text="getGuestName(row)">
                  <span class="reservation-single-line">{{ getGuestName(row) }}</span>
                </div>
              </td>
              <td class="px-3 py-2 text-muted-foreground whitespace-nowrap">
                <div class="reservation-hover-tip" :data-full-text="getGuestPhone(row)">
                  <span class="reservation-single-line">{{ getGuestPhone(row) }}</span>
                </div>
              </td>
              <td class="px-3 py-2">
                <div class="reservation-hover-tip" :data-full-text="getPropertyName(row)">
                  <span class="reservation-single-line">{{ getPropertyName(row) }}</span>
                </div>
              </td>
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">
                {{ formatDateOnly(row.checkInDate || row.checkinDate) }}
              </td>
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">
                {{ formatDateOnly(row.checkOutDate || row.checkoutDate) }}
              </td>
              <td class="reservation-cell-amount py-2 font-medium whitespace-nowrap">
                {{ formatMoney(row.orderAmount) }}
              </td>
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">
                {{ formatDateOnly(row.createdAt) }}
              </td>
              <td class="reservation-cell-center py-2">
                <span :class="['reservation-status', getRefundClass(row.refundStatus)]">
                  {{ getOptionLabel(options.refund_status, row.refundStatus) || '—' }}
                </span>
              </td>
              <td class="reservation-cell-amount py-2 font-medium whitespace-nowrap">
                {{ formatMoney(row.refundAmount) }}
              </td>
              <td class="reservation-cell-action py-2">
                <div class="reservation-action-menu">
                  <button type="button" class="reservation-kebab" @click.stop="toggleActionMenu(row.id)">
                    <MoreHorizontal :size="19" />
                  </button>
                  <div v-if="activeActionId === row.id" class="reservation-action-popover">
                    <div class="reservation-menu-section">
                      <div class="reservation-menu-title">操作</div>
                      <button type="button" class="reservation-menu-item" @click="runAction('view', row)">查看详情</button>
                      <button
                        v-if="canRefund(row)"
                        type="button"
                        class="reservation-menu-item"
                        @click="runAction('refund', row)"
                      >退款</button>
                    </div>
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="pageCount > 0" class="reservation-pagination">
        <span>共 {{ totalCount }} 条</span>
        <div class="flex items-center gap-1">
          <button @click="changePage(page - 1)" :disabled="page <= 1" class="reservation-page-btn">
            <ChevronLeft :size="13" />
          </button>
          <template v-for="p in pages" :key="p">
            <button
              v-if="p !== '...'"
              @click="changePage(p as number)"
              :class="['reservation-page-num', page === p ? 'reservation-page-num-active' : '']"
            >{{ p }}</button>
            <span v-else class="h-7 flex items-center px-1">...</span>
          </template>
          <button @click="changePage(page + 1)" :disabled="page >= pageCount" class="reservation-page-btn">
            <ChevronRight :size="13" />
          </button>
          <select
            :value="pageSize"
            class="reservation-page-size"
            @change="changePageSize(Number(($event.target as HTMLSelectElement).value))"
          >
            <option v-for="s in [10, 20, 50]" :key="s" :value="s">{{ s }}/页</option>
          </select>
        </div>
      </div>
    </div>

    <n-modal
      v-model:show="state.showreturn"
      :mask-closable="false"
      preset="dialog"
      title="退款"
      positive-text="提交"
      negative-text="算了"
      @positive-click="onPositiveClick"
      @negative-click="onNegativeClick"
      :style="{ width: dialogWidth }"
    >
      <div style="width: 500px;max-width: 100%;" v-if="state.showreturn">
        <n-form ref="formRef" :model="state.returnAmount" :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="120" class="py-4">
          <n-form-item label="订单号" path="maxAmount">{{ state.returnAmount.ordersn }}</n-form-item>
          <n-form-item label="最大退款金额" path="maxAmount">{{ state.returnAmount.maxAmount }} (JPY)</n-form-item>
          <n-form-item label="主动退款" path="maxAmount">
            <n-input-number
              v-model:value="state.returnAmount.refundAmount"
              style="width: 220px"
              :max="state.returnAmount.maxAmount"
              :placeholder="'最大退款金额' + state.returnAmount.maxAmount"
            />
          </n-form-item>
          <n-form-item label="退款方式" path="isCancelOrder">
            <a-radio-group v-model:value="state.returnAmount.isCancelOrder" name="radioGroup">
              <a-radio value="N">仅退款</a-radio>
              <a-radio value="Y" v-if="state.orderStatus != 'CANCEL'">退款并取消订单</a-radio>
              <a-radio value="Y" v-else disabled="true">退款并取消订单</a-radio>
            </a-radio-group>
            <template #feedback>
              <div style="font-size: 12px">仅退款：订单会处于部分退款状态，订单状态不会改变</div>
              <div style="font-size: 12px; margin-bottom: 8px">退款并取消订单：退款后，订单会处于已取消状态</div>
            </template>
          </n-form-item>
          <n-form-item label="退款取消说明" path="reason">
            <n-input type="textarea" placeholder="请输入退款说明" v-model:value="state.returnAmount.reason" />
          </n-form-item>
        </n-form>
      </div>
    </n-modal>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
import { computed, h, onBeforeUnmount, onMounted, reactive, ref, watch } from 'vue';
import { useMessage } from 'naive-ui';
import { UiButton, UiInput, UiSelect } from '@/components/ui';
import {
  CalendarCheck,
  ChevronDown as ChevronDownIcon,
  ChevronLeft,
  ChevronRight,
  Download,
  Loader2,
  Minus,
  MoreHorizontal,
  RefreshCw,
  RotateCcw,
  Search,
  SearchX,
} from '@lucide/vue';
import { Export, List, paycloud } from '@/api/pmsAppReservation';
import { loadOptions, options } from './model';
import { adaModalWidth, getOptionLabel } from '@/utils/hotgo';
import View from './view.vue';
import { useUserStore } from '@/store/modules/user';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { storage } from '@/utils/Storage';
import { getlang } from '@/utils/smjcomm';
import { usePermission } from '@/hooks/web/usePermission';

const userStore = useUserStore();
const message = useMessage();
const { hasPermission } = usePermission();
const settingStore = useProjectSettingStore();
const viewRef = ref();
const loading = ref(false);
const data = ref<any[]>([]);
const page = ref(1);
const pageSize = ref(10);
const pageCount = ref(0);
const totalCount = ref(0);
const showMore = ref(true);
const activeActionId = ref<number | string | null>(null);
const selectedPropertyId = ref(userStore.getuserPms.id || 0);
const rawPropertyList = ref<any[]>(storage.get('wylists') || []);

const dialogWidth = computed(() => adaModalWidth(620));

const state = reactive({
  showreturn: false,
  btnloading: false,
  orderStatus: null as string | null,
  returnAmount: {
    ordersn: '',
    refundAmount: null as number | null,
    maxAmount: 0,
    isCancelOrder: 'N',
    reason: '',
  },
});

const searchForm = reactive({
  snKeyword: '',
  orderStatus: null as string | null,
  refund_status: null as string | null,
  MemberNo: '',
  checkinDateStart: '',
  checkinDateEnd: '',
  checkoutDateStart: '',
  checkoutDateEnd: '',
  createdAtStart: '',
  createdAtEnd: '',
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
  if (searchForm.snKeyword) p.snKeyword = searchForm.snKeyword;
  if (searchForm.orderStatus) p.orderStatus = searchForm.orderStatus;
  if (searchForm.refund_status) p.refund_status = searchForm.refund_status;
  if (searchForm.MemberNo) p.MemberNo = searchForm.MemberNo;
  if (searchForm.checkinDateStart || searchForm.checkinDateEnd) {
    p.checkinDate = [searchForm.checkinDateStart, searchForm.checkinDateEnd];
  }
  if (searchForm.checkoutDateStart || searchForm.checkoutDateEnd) {
    p.checkoutDate = [searchForm.checkoutDateStart, searchForm.checkoutDateEnd];
  }
  if (searchForm.createdAtStart || searchForm.createdAtEnd) {
    p.createdAt = [
      searchForm.createdAtStart ? `${searchForm.createdAtStart} 00:00:00` : '',
      searchForm.createdAtEnd ? `${searchForm.createdAtEnd} 23:59:59` : '',
    ];
  }
  return p;
}

async function loadData() {
  loading.value = true;
  try {
    const res = await List({ page: page.value, pageSize: pageSize.value, ...buildParams() });
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
  searchForm.snKeyword = '';
  searchForm.orderStatus = null;
  searchForm.refund_status = null;
  searchForm.MemberNo = '';
  searchForm.checkinDateStart = '';
  searchForm.checkinDateEnd = '';
  searchForm.checkoutDateStart = '';
  searchForm.checkoutDateEnd = '';
  searchForm.createdAtStart = '';
  searchForm.createdAtEnd = '';
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

function getGuestPhone(row: any): string {
  return row.guestProfileDetail?.phone || row.mainGuestDetail?.phone || row.guestProfile?.phone || '—';
}

function getPropertyName(row: any): string {
  return row.propertyDetail?.name || row.propertyName || '—';
}

function getStatusClass(status: string | null | undefined): string {
  if (status === 'DONE' || status === 'CONFIRMED') return 'reservation-status-ok';
  if (status === 'CANCEL' || status === 'FAILED') return 'reservation-status-danger';
  return 'reservation-status-muted';
}

function getRefundClass(status: string | null | undefined): string {
  if (status === 'DONE') return 'reservation-status-ok';
  if (status === 'WAIT') return 'reservation-status-warn';
  if (status === 'FAILED') return 'reservation-status-danger';
  return 'reservation-status-muted';
}

function canRefund(record: any): boolean {
  return record.isRefund === true && Number(record.refundAmount || 0) < Number(record.orderAmount || 0);
}

function toggleActionMenu(id: number | string) {
  activeActionId.value = activeActionId.value === id ? null : id;
}

function closeActionMenu() {
  activeActionId.value = null;
}

function handleDocumentClick(event: MouseEvent) {
  const target = event.target as HTMLElement;
  if (!target.closest('.reservation-action-menu')) closeActionMenu();
}

function runAction(key: string, record: any) {
  closeActionMenu();
  if (key === 'refund') return handleRefund(record);
  return handleEdit(record);
}

function handleRefund(record: Recordable) {
  state.returnAmount.ordersn = record.orderSn;
  state.returnAmount.isCancelOrder = 'N';
  state.returnAmount.maxAmount = 0;
  state.orderStatus = record.orderStatus;
  let payAmount = 0;
  (record.transactionDetail || []).map((m) => {
    if (m.payStatus === 'DONE' && m.payType !== 'COUPON') payAmount += m.payAmount;
  });

  let refundAmount = 0;
  if (record.transactionRefundDetail != null) {
    for (let i = 0; i < record.transactionRefundDetail.length; i++) {
      const item = record.transactionRefundDetail[i];
      if (item.refundStatus === 'DONE' || item.refundStatus === 'WAIT') refundAmount += item.refundAmount;
    }
  }

  state.returnAmount.maxAmount += payAmount - refundAmount;
  state.showreturn = true;
}

function handleEdit(record: Recordable) {
  viewRef.value.openModal(record);
}

const onPositiveClick = async () => {
  paycloud(state.returnAmount).then((_res) => {
    message.success('已提交');
    state.showreturn = false;
    state.returnAmount.refundAmount = null;
    state.returnAmount.maxAmount = 0;
    state.returnAmount.reason = '';
    state.returnAmount.isCancelOrder = 'N';
    loadData();
  });
};

const onNegativeClick = () => {
  state.returnAmount.refundAmount = null;
  state.returnAmount.maxAmount = 0;
  state.showreturn = false;
};

function handleExport() {
  message.loading('正在导出列表...', { duration: 1200 });
  Export(buildParams());
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
.reservation-page {
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
  --crs-warning: #d97706;
  --crs-warning-soft: #fff7ed;
  --crs-shadow: 0 1px 2px rgba(15, 23, 42, 0.04), 0 0 0 1px rgba(220, 231, 242, 0.5);
  background: #f6f9fc;
  color: var(--crs-text);
  min-height: calc(100vh - var(--header-height, 56px));
  padding: 16px;

  .reservation-toolbar {
    align-items: flex-end;
    display: flex;
    gap: 12px;
    justify-content: space-between;
  }

  .reservation-property {
    max-width: 360px;
    min-width: 260px;
    width: 32vw;
  }

  .reservation-header-actions {
    align-items: center;
    display: flex;
    gap: 8px;
  }

  .reservation-panel {
    background: var(--crs-surface);
    border: 1px solid var(--crs-border);
    border-radius: 8px;
    box-shadow: var(--crs-shadow);
  }

  .reservation-panel-section {
    padding: 16px;
  }

  .reservation-filter-grid {
    display: grid;
    gap: 12px;
    grid-template-columns: repeat(4, minmax(0, 1fr));
  }

  .reservation-more-grid {
    display: grid;
    gap: 12px;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    margin-top: 12px;
  }

  .reservation-field {
    min-width: 0;
  }

  .reservation-label {
    color: var(--crs-muted);
    font-size: 12px;
    font-weight: 600;
    line-height: 18px;
    margin-bottom: 4px;
  }

  .reservation-date-range {
    align-items: center;
    display: flex;
    gap: 6px;
  }

  .reservation-date-input {
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

  .reservation-filter-actions {
    align-items: center;
    display: flex;
    justify-content: space-between;
    margin-top: 14px;
  }

  .reservation-disclosure,
  .reservation-refresh {
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

  .reservation-disclosure:hover,
  .reservation-refresh:hover {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
  }

  .reservation-table-panel {
    overflow: visible;
  }

  .reservation-table-toolbar {
    align-items: center;
    border-bottom: 1px solid var(--crs-border);
    display: flex;
    justify-content: space-between;
    padding: 10px 16px 10px 24px;
  }

  .reservation-table-scroll {
    overflow-x: auto;
    overflow-y: visible;
    position: relative;
  }

  .reservation-table-scroll-open {
    padding-bottom: 220px;
  }

  .reservation-table {
    min-width: 1320px;
    table-layout: fixed;
  }

  .reservation-col-order {
    width: 160px;
  }

  .reservation-col-status,
  .reservation-col-refund-status {
    width: 86px;
  }

  .reservation-col-guest {
    width: 104px;
  }

  .reservation-col-contact {
    width: 126px;
  }

  .reservation-col-property {
    width: 150px;
  }

  .reservation-col-date {
    width: 104px;
  }

  .reservation-col-amount {
    width: 104px;
  }

  .reservation-col-action {
    width: 64px;
  }

  .reservation-cell-order {
    padding-left: 24px;
    padding-right: 10px;
  }

  .reservation-cell-center {
    padding-left: 8px;
    padding-right: 8px;
    text-align: center;
  }

  .reservation-cell-amount {
    padding-left: 8px;
    padding-right: 12px;
    text-align: left;
  }

  .reservation-cell-action {
    background: var(--crs-surface);
    box-shadow: -10px 0 16px -16px rgba(15, 23, 42, 0.35);
    padding-left: 8px;
    padding-right: 8px;
    position: sticky;
    right: 0;
    text-align: center;
    z-index: 6;
  }

  thead .reservation-cell-action {
    background: #f8fafc;
    z-index: 8;
  }

  tbody tr:hover .reservation-cell-action {
    background: #f8fafc;
  }

  .reservation-hover-tip {
    display: block;
    max-width: 100%;
    position: relative;
  }

  .reservation-single-line {
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .reservation-hover-tip:hover::after {
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

  .reservation-status {
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

  .reservation-status-ok {
    background: var(--crs-success-soft);
    color: var(--crs-success);
  }

  .reservation-status-danger {
    background: var(--crs-danger-soft);
    color: var(--crs-danger);
  }

  .reservation-status-warn {
    background: var(--crs-warning-soft);
    color: var(--crs-warning);
  }

  .reservation-status-muted {
    background: #f1f5f9;
    color: var(--crs-muted);
  }

  .reservation-count {
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

  .reservation-empty {
    align-items: center;
    color: var(--crs-muted);
    display: inline-flex;
    flex-direction: column;
    font-size: 13px;
    gap: 8px;
  }

  .reservation-action-menu {
    display: inline-flex;
    justify-content: center;
    position: relative;
    width: 100%;
  }

  .reservation-kebab {
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

  .reservation-kebab:hover,
  .reservation-kebab:focus-visible {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
    box-shadow: 0 0 0 3px rgba(56, 174, 234, 0.14);
  }

  .reservation-action-popover {
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

  .reservation-menu-section {
    display: grid;
    gap: 2px;
  }

  .reservation-menu-title {
    color: var(--crs-muted);
    font-size: 12px;
    font-weight: 700;
    line-height: 18px;
    padding: 2px 14px 4px;
  }

  .reservation-menu-item {
    color: var(--crs-text);
    display: block;
    font-size: 13px;
    line-height: 20px;
    padding: 7px 14px;
    text-align: left;
    transition: background 0.15s ease, color 0.15s ease;
    width: 100%;
  }

  .reservation-menu-item:hover {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
  }

  .reservation-pagination {
    align-items: center;
    border-top: 1px solid var(--crs-border);
    color: var(--crs-muted);
    display: flex;
    font-size: 13px;
    justify-content: space-between;
    padding: 12px 16px;
  }

  .reservation-page-btn,
  .reservation-page-num {
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

  .reservation-page-btn:hover,
  .reservation-page-num:hover {
    background: #f1f5f9;
  }

  .reservation-page-btn:disabled {
    cursor: not-allowed;
    opacity: 0.45;
  }

  .reservation-page-num-active {
    background: var(--crs-primary) !important;
    border-color: var(--crs-primary);
    color: #fff;
  }

  .reservation-page-size {
    background: #fff;
    border: 1px solid var(--crs-border);
    border-radius: 6px;
    height: 28px;
    padding: 0 8px;
  }

  @media (max-width: 1024px) {
    .reservation-filter-grid,
    .reservation-more-grid {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  @media (max-width: 640px) {
    .reservation-toolbar,
    .reservation-filter-actions,
    .reservation-pagination {
      align-items: stretch;
      flex-direction: column;
      gap: 10px;
    }

    .reservation-property {
      max-width: none;
      min-width: 0;
      width: 100%;
    }

    .reservation-filter-grid,
    .reservation-more-grid {
      grid-template-columns: 1fr;
    }
  }
}
</style>
