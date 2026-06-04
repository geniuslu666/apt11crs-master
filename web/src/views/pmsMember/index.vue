<template>
  <div class="member-admin-page member-page space-y-4">
    <div class="member-page-header">
      <div>
        <div class="member-eyebrow">Members</div>
        <h1 class="member-title">会员列表</h1>
      </div>
      <div class="member-header-actions">
        <UiButton size="sm" variant="outline" @click="loadData" :disabled="loading">
          <RefreshCw :size="14" :class="loading ? 'animate-spin' : ''" />
          刷新
        </UiButton>
        <UiButton size="sm" variant="outline" @click="handleExport" v-if="hasPermission(['/pmsMember/export'])">
          <Download :size="14" />
          导出
        </UiButton>
      </div>
    </div>

    <!-- Search card -->
    <div class="member-panel">
      <div class="member-panel-section">
        <div class="member-filter-grid">
          <div class="member-field">
            <div class="member-label">会员号</div>
            <UiInput v-model="searchForm.memberNo" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div class="member-field">
            <div class="member-label">手机号</div>
            <UiInput v-model="searchForm.phone" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div class="member-field">
            <div class="member-label">昵称</div>
            <UiInput v-model="searchForm.fullName" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div class="member-field">
            <div class="member-label">邮箱</div>
            <UiInput v-model="searchForm.mail" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div class="member-field">
            <div class="member-label">会员等级</div>
            <UiSelect
              v-model="searchForm.level"
              :options="levelList"
              label-field="levelName"
              value-field="id"
              placeholder="全部等级"
              clearable
              filterable
            />
          </div>
        </div>

        <!-- More filters (date ranges) -->
        <div v-show="showMore" class="member-more-grid">
          <div class="member-field">
            <div class="member-label">注册时间</div>
            <div class="flex items-center gap-1.5">
              <input
                type="datetime-local"
                v-model="searchForm.createdAtStart"
                class="member-date-input"
              />
              <Minus :size="12" class="text-muted-foreground flex-shrink-0" />
              <input
                type="datetime-local"
                v-model="searchForm.createdAtEnd"
                class="member-date-input"
              />
            </div>
          </div>
          <div class="member-field">
            <div class="member-label">最后登录</div>
            <div class="flex items-center gap-1.5">
              <input
                type="datetime-local"
                v-model="searchForm.lastLoginStart"
                class="member-date-input"
              />
              <Minus :size="12" class="text-muted-foreground flex-shrink-0" />
              <input
                type="datetime-local"
                v-model="searchForm.lastLoginEnd"
                class="member-date-input"
              />
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center justify-between mt-4">
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
          <button
            type="button"
            @click="showMore = !showMore"
            class="member-disclosure"
          >
            {{ showMore ? '收起筛选' : '展开筛选' }}
            <ChevronDown :size="13" :class="['transition-transform duration-200', showMore ? 'rotate-180' : '']" />
          </button>
        </div>
      </div>
    </div>

    <!-- Table card -->
    <div class="member-panel member-table-panel">
      <!-- Toolbar -->
      <div class="member-table-toolbar">
        <div class="flex items-center gap-2">
          <Users :size="16" class="text-primary" />
          <span class="text-[13px] font-semibold text-foreground">会员列表</span>
          <span
            v-if="totalCount > 0"
            class="member-count"
          >{{ totalCount }}</span>
        </div>
        <button
          type="button"
          @click="loadData"
          :disabled="loading"
          class="member-refresh"
        >
          <RefreshCw :size="14" :class="loading ? 'animate-spin' : ''" />
          刷新
        </button>
      </div>

      <!-- Table -->
      <div :class="['member-table-scroll', activeActionId !== null ? 'member-table-scroll-open' : '']">
        <!-- Loading overlay -->
        <div v-if="loading" class="absolute inset-0 bg-white/70 z-10 flex items-center justify-center">
          <Loader2 class="animate-spin text-primary" :size="22" />
        </div>

        <table class="member-table w-full text-[13px]">
          <colgroup>
            <col class="member-colgroup-name" />
            <col class="member-colgroup-id" />
            <col class="member-colgroup-level" />
            <col class="member-colgroup-phone" />
            <col class="member-colgroup-email" />
            <col class="member-colgroup-status" />
            <col class="member-colgroup-source" />
            <col class="member-colgroup-date" />
            <col class="member-colgroup-date" />
            <col class="member-colgroup-action" />
          </colgroup>
          <thead>
            <tr class="border-b border-border bg-[#f8fafc]">
              <th class="member-col-name py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">会员昵称</th>
              <th class="member-col-id py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">会员ID</th>
              <th class="member-col-level py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">等级</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">手机信息</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">邮箱信息</th>
              <th class="member-col-status py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">状态</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">注册来源</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">上次登录</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">创建时间</th>
              <th class="member-col-action py-2.5 text-center text-[11px] font-medium text-muted-foreground whitespace-nowrap">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-border">
            <tr v-if="data.length === 0 && !loading">
              <td colspan="10" class="px-3 py-14 text-center">
                <div class="member-empty">
                  <SearchX :size="24" />
                  <span>暂无数据</span>
                </div>
              </td>
            </tr>
            <tr v-for="row in data" :key="row.id" class="hover:bg-[#f8fafc] transition-colors">

              <!-- 会员昵称 -->
              <td class="member-col-name py-2">
                <div
                  class="member-name-cell"
                  :data-full-name="getMemberName(row)"
                >
                  <span class="member-name-text">{{ getMemberName(row) }}</span>
                </div>
              </td>

              <!-- 会员ID -->
              <td class="member-col-id py-2 text-muted-foreground whitespace-nowrap">
                {{ row.memberNo || row.id || '—' }}
              </td>

              <!-- 会员等级 -->
              <td class="member-col-level py-2 whitespace-nowrap">
                <span class="member-level-badge" :title="row.memberLevel?.levelName || '未分级'">
                  <img
                    v-if="row.memberLevel?.levelBadge"
                    :src="row.memberLevel.levelBadge"
                    alt=""
                    class="member-level-badge-img"
                  />
                  <Award v-else :size="14" />
                </span>
              </td>

              <!-- 手机信息 -->
              <td class="px-3 py-2 text-muted-foreground whitespace-nowrap">
                {{ row.phone ? `${row.phoneArea}-${row.phone}` : '—' }}
              </td>

              <!-- 邮箱信息 -->
              <td class="px-3 py-2 text-muted-foreground">
                <div class="member-single-line">{{ row.mail || '—' }}</div>
              </td>

              <!-- 状态 -->
              <td class="member-col-status py-2">
                <span :class="[
                  'member-status inline-flex items-center px-1.5 py-0.5 rounded text-[11px] font-medium',
                  row.status === 1
                    ? 'member-status-ok'
                    : 'member-status-danger'
                ]">
                  {{ getOptionLabel(options.sys_normal_disable, row.status) || '—' }}
                </span>
              </td>

              <!-- 注册来源 -->
              <td class="px-3 py-2 text-muted-foreground whitespace-nowrap">
                {{ row.source || '—' }}
              </td>

              <!-- 上次登录 -->
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">
                {{ formatDateOnly(row.lastLogin) }}
              </td>

              <!-- 创建时间 -->
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">
                {{ formatDateOnly(row.createdAt) }}
              </td>

              <!-- 操作 -->
              <td class="member-col-action py-2">
                <div class="member-action-menu">
                  <button
                    type="button"
                    class="member-kebab"
                    @click.stop="toggleActionMenu(row.id)"
                  >
                    <MoreHorizontal :size="19" />
                  </button>

                  <div v-if="activeActionId === row.id" class="member-action-popover">
                    <div class="member-menu-section">
                      <div class="member-menu-title">操作</div>
                      <button
                        v-if="hasPermission(['/pmsMember/view'])"
                        type="button"
                        class="member-menu-item"
                        @click="runAction('view', row)"
                      >查看详情</button>
                      <button
                        v-if="row.status === 1 && hasPermission(['/pmsMember/status'])"
                        type="button"
                        class="member-menu-item member-menu-danger"
                        @click="runAction('disable', row)"
                      >禁用会员</button>
                      <button
                        v-if="row.status === 2 && hasPermission(['/pmsMember/status'])"
                        type="button"
                        class="member-menu-item"
                        @click="runAction('enable', row)"
                      >启用会员</button>
                      <button
                        v-if="row.memberCancelArr == null && hasPermission(['/pmsWithdraw/disagreeStaff'])"
                        type="button"
                        class="member-menu-item member-menu-danger"
                        @click="runAction('cancel', row)"
                      >注销会员</button>
                    </div>
                    <div class="member-menu-divider"></div>
                    <div class="member-menu-section">
                      <div class="member-menu-title">连接</div>
                      <button type="button" class="member-menu-item" @click="runAction('mes', row)">发通知</button>
                      <button type="button" class="member-menu-item" @click="runAction('mail', row)">发邮件</button>
                      <button type="button" class="member-menu-item" @click="runAction('telmes', row)">发短信</button>
                      <button type="button" class="member-menu-item" @click="runAction('coupon', row)">发优惠券</button>
                      <button type="button" class="member-menu-item" @click="runAction('thcoupon', row)">发礼品券</button>
                      <button type="button" class="member-menu-item" @click="runAction('h5_link', row)">H5分销链接</button>
                    </div>
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="pageCount > 0" class="px-4 py-3 border-t border-border flex items-center justify-between text-[13px] text-muted-foreground">
        <span>共 {{ totalCount }} 条</span>
        <div class="flex items-center gap-1">
          <button
            @click="changePage(page - 1)"
            :disabled="page <= 1"
            class="h-7 px-2 rounded border border-border hover:bg-muted disabled:opacity-40 disabled:cursor-not-allowed transition-colors"
          >
            <ChevronLeft :size="13" />
          </button>
          <template v-for="p in pages" :key="p">
            <button
              v-if="p !== '...'"
              @click="changePage(p as number)"
              :class="['h-7 min-w-[1.75rem] px-1.5 rounded border text-[13px] transition-colors', page === p ? 'bg-primary text-primary-foreground border-primary' : 'border-border hover:bg-muted']"
            >{{ p }}</button>
            <span v-else class="h-7 flex items-center px-1">…</span>
          </template>
          <button
            @click="changePage(page + 1)"
            :disabled="page >= pageCount"
            class="h-7 px-2 rounded border border-border hover:bg-muted disabled:opacity-40 disabled:cursor-not-allowed transition-colors"
          >
            <ChevronRight :size="13" />
          </button>
          <select
            :value="pageSize"
            @change="changePageSize(Number(($event.target as HTMLSelectElement).value))"
            class="h-7 border border-border rounded px-2 text-[13px] bg-background hover:bg-muted transition-colors cursor-pointer"
          >
            <option v-for="s in [10, 20, 50]" :key="s" :value="s">{{ s }}/页</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Sub-components -->
    <Edit ref="editRef" @reloadTable="loadData" />
    <sendemail ref="sendemailref" />
    <sendnotify ref="sendnotifyref" />
    <sendsms ref="sendsmsref" />
    <h5Link ref="h5Linkref" />
    <ChooseCoupon ref="chooseCouponRef" />
    <ChooseThCoupon ref="chooseThCouponRef" />
    <MemberCancel ref="memberCancelRef" @reloadTable="loadData" />
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, computed, onMounted, onBeforeUnmount } from 'vue';
import { useMessage } from 'naive-ui';
import { UiButton, UiInput, UiSelect } from '@/components/ui';
import {
  Award,
  ChevronDown,
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
  Users,
} from '@lucide/vue';
import { List, Status, Export } from '@/api/pmsMember';
import { options, loadOptions, levelList } from './model';
import { getOptionLabel } from '@/utils/hotgo';
import Edit from './edit.vue';
import sendemail from '@/views/smjcomm/sendemail.vue';
import sendnotify from '@/views/pmsMember/sendnotify.vue';
import sendsms from '@/views/smjcomm/sendsms.vue';
import h5Link from '@/views/smjcomm/h5Link.vue';
import { useRouter } from 'vue-router';
import ChooseCoupon from '@/views/pmsMember/chooseCouponType.vue';
import ChooseThCoupon from '@/views/pmsMember/chooseThCoupon.vue';
import MemberCancel from '@/views/pmsMember/member_cancel.vue';
import { usePermission } from '@/hooks/web/usePermission';

const { hasPermission } = usePermission();
const router = useRouter();
const message = useMessage();

// Modal refs
const editRef = ref();
const sendemailref = ref();
const sendnotifyref = ref();
const sendsmsref = ref();
const h5Linkref = ref();
const chooseCouponRef = ref();
const chooseThCouponRef = ref();
const memberCancelRef = ref();

// UI state
const showMore = ref(true);
const activeActionId = ref<number | string | null>(null);

// Table state
const loading = ref(false);
const data = ref<any[]>([]);
const page = ref(1);
const pageSize = ref(10);
const pageCount = ref(0);
const totalCount = ref(0);

// Search form
const searchForm = reactive({
  memberNo: '',
  phone: '',
  fullName: '',
  mail: '',
  level: null as number | null,
  createdAtStart: '',
  createdAtEnd: '',
  lastLoginStart: '',
  lastLoginEnd: '',
});

function buildParams() {
  const p: Record<string, any> = {};
  if (searchForm.memberNo) p.memberNo = searchForm.memberNo;
  if (searchForm.phone) p.phone = searchForm.phone;
  if (searchForm.fullName) p.fullName = searchForm.fullName;
  if (searchForm.mail) p.mail = searchForm.mail;
  if (searchForm.level !== null && searchForm.level !== undefined) p.level = searchForm.level;
  if (searchForm.createdAtStart || searchForm.createdAtEnd) {
    p.createdAt = [
      searchForm.createdAtStart ? searchForm.createdAtStart.replace('T', ' ') : '',
      searchForm.createdAtEnd ? searchForm.createdAtEnd.replace('T', ' ') : '',
    ];
  }
  if (searchForm.lastLoginStart || searchForm.lastLoginEnd) {
    p.lastLogin = [
      searchForm.lastLoginStart ? searchForm.lastLoginStart.replace('T', ' ') : '',
      searchForm.lastLoginEnd ? searchForm.lastLoginEnd.replace('T', ' ') : '',
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

function changePage(p: number) {
  page.value = p;
  loadData();
}

function changePageSize(s: number) {
  pageSize.value = s;
  page.value = 1;
  loadData();
}

const pages = computed(() => {
  if (pageCount.value <= 7) return Array.from({ length: pageCount.value }, (_, i) => i + 1);
  const arr: (number | '...')[] = [1];
  if (page.value > 3) arr.push('...');
  for (let i = Math.max(2, page.value - 1); i <= Math.min(pageCount.value - 1, page.value + 1); i++) arr.push(i);
  if (page.value < pageCount.value - 2) arr.push('...');
  if (pageCount.value > 1) arr.push(pageCount.value);
  return arr;
});

function handleSearch() {
  page.value = 1;
  loadData();
}

function handleReset() {
  searchForm.memberNo = '';
  searchForm.phone = '';
  searchForm.fullName = '';
  searchForm.mail = '';
  searchForm.level = null;
  searchForm.createdAtStart = '';
  searchForm.createdAtEnd = '';
  searchForm.lastLoginStart = '';
  searchForm.lastLoginEnd = '';
  page.value = 1;
  loadData();
}

function getMemberName(row: any): string {
  const name = `${row.lastName || ''}${row.firstName || ''}`.trim();
  return name || row.fullName || row.nickname || row.nickName || '—';
}

function formatDateOnly(value: string | null | undefined): string {
  if (!value) return '—';
  return String(value).slice(0, 10) || '—';
}

function toggleActionMenu(id: number | string) {
  activeActionId.value = activeActionId.value === id ? null : id;
}

function closeActionMenu() {
  activeActionId.value = null;
}

function handleDocumentClick(event: MouseEvent) {
  const target = event.target as HTMLElement;
  if (!target.closest('.member-action-menu')) closeActionMenu();
}

function runAction(key: string, record: any) {
  closeActionMenu();
  if (key === 'view') return handleView(record);
  if (key === 'disable') return handleStatus(record, 2);
  if (key === 'enable') return handleStatus(record, 1);
  if (key === 'cancel') return handleCancel(record);
  return handleDropdownSelect(key, record);
}

function handleDropdownSelect(key: string, record: any) {
  if (key === 'mes') return tomessage(record);
  if (key === 'mail') return goemail(record);
  if (key === 'telmes') return gosend(record);
  if (key === 'coupon') return sendCoupon(record);
  if (key === 'thcoupon') return sendThCoupon(record);
  if (key === 'h5_link') return goLink(record);
}

function gosend(aa: any) {
  if (aa.phone) {
    sendsmsref.value.openModal({ phone: aa.phone, area_no: aa.phoneArea });
  } else {
    message.error('未找到联络方式');
  }
}
function goemail(aa: any) {
  if (aa.mail) {
    sendemailref.value.openModal(aa.mail);
  } else {
    message.error('该会员无邮箱');
  }
}
function goLink(aa: any) {
  if (aa.id) {
    h5Linkref.value.openModal({ id: aa.id });
  } else {
    message.error('未找到联络方式');
  }
}
function sendCoupon(record: any) { chooseCouponRef.value.openModal(record.id); }
function sendThCoupon(record: any) { chooseThCouponRef.value.openModal(record.id); }
function tomessage(aa: any) { sendnotifyref.value.openModal(aa.id); }

function handleView(record: any) {
  router.push({ name: 'pmsMember_view', params: { id: record.id } });
}

function handleStatus(record: any, status: number) {
  Status({ id: record.id, status }).then(() => {
    message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
    loadData();
  });
}

function handleCancel(record: any) {
  memberCancelRef.value.openModal(record);
}

function handleExport() {
  message.loading('正在导出列表...', { duration: 1200 });
  Export(buildParams());
}

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
.member-page {
  --crs-primary: #38aeea;
  --crs-primary-strong: #128fc8;
  --crs-primary-soft: #eaf7ff;
  --crs-border: #dce7f2;
  --crs-surface: #ffffff;
  --crs-surface-soft: #f7fbff;
  --crs-text: #152033;
  --crs-muted: #6b7c93;
  --crs-danger: #ef4444;
  --crs-danger-soft: #fff1f2;
  --crs-success: #16a34a;
  --crs-success-soft: #eefbf3;
  --crs-shadow: 0 1px 2px rgba(15, 23, 42, 0.04), 0 0 0 1px rgba(220, 231, 242, 0.5);
  padding: 16px;
  color: var(--crs-text);
  background: #f6f9fc;
  min-height: calc(100vh - var(--header-height, 56px));

  .member-page-header {
    display: flex;
    align-items: flex-end;
    justify-content: space-between;
    gap: 12px;
  }

  .member-eyebrow {
    color: var(--crs-primary-strong);
    font-size: 11px;
    font-weight: 700;
    line-height: 16px;
    text-transform: uppercase;
  }

  .member-title {
    color: var(--crs-text);
    font-size: 20px;
    font-weight: 700;
    line-height: 28px;
    margin: 0;
  }

  .member-header-actions {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .member-panel {
    background: var(--crs-surface);
    border: 1px solid var(--crs-border);
    border-radius: 8px;
    box-shadow: var(--crs-shadow);
  }

  .member-table-panel {
    overflow: visible;
  }

  .member-panel-section {
    padding: 16px;
  }

  .member-filter-grid {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 12px;
  }

  .member-field {
    min-width: 0;
  }

  .member-label {
    color: var(--crs-muted);
    font-size: 12px;
    font-weight: 600;
    line-height: 18px;
    margin-bottom: 4px;
  }

  .member-more-grid {
    display: grid;
    gap: 12px;
    grid-template-columns: repeat(1, minmax(0, 1fr));
    margin-top: 12px;
  }

  .member-date-input {
    background: hsl(var(--background));
    border: 1px solid hsl(var(--input));
    border-radius: 6px;
    color: var(--crs-text);
    flex: 1 1 0;
    font-size: 13px;
    height: 32px;
    min-width: 0;
    padding: 0 8px;
    transition: border-color 0.15s ease, box-shadow 0.15s ease;
  }

  .member-disclosure,
  .member-refresh {
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

  .member-disclosure:hover,
  .member-refresh:hover {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
  }

  .member-refresh:disabled {
    cursor: not-allowed;
    opacity: 0.55;
  }

  .member-table-toolbar {
    align-items: center;
    border-bottom: 1px solid var(--crs-border);
    display: flex;
    justify-content: space-between;
    padding: 10px 16px 10px 24px;
  }

  .member-table-scroll {
    overflow-x: auto;
    overflow-y: visible;
    position: relative;
  }

  .member-table-scroll-open {
    padding-bottom: 360px;
  }

  .member-table {
    min-width: 1040px;
    table-layout: fixed;
  }

  .member-col-name {
    padding-left: 46px;
    padding-right: 12px;
  }

  .member-col-id {
    padding-left: 10px;
    padding-right: 10px;
  }

  .member-col-level,
  .member-col-status {
    padding-left: 8px;
    padding-right: 8px;
    text-align: center;
  }

  .member-col-action {
    background: var(--crs-surface);
    box-shadow: -10px 0 16px -16px rgba(15, 23, 42, 0.35);
    padding-left: 8px;
    padding-right: 8px;
    position: sticky;
    right: 0;
    text-align: center;
    z-index: 6;
  }

  thead .member-col-action {
    background: #f8fafc;
    z-index: 8;
  }

  tbody tr:hover .member-col-action {
    background: #f8fafc;
  }

  .member-colgroup-name {
    width: 176px;
  }

  .member-colgroup-id {
    width: 92px;
  }

  .member-colgroup-level {
    width: 64px;
  }

  .member-colgroup-phone {
    width: 144px;
  }

  .member-colgroup-email {
    width: 180px;
  }

  .member-colgroup-status {
    width: 74px;
  }

  .member-colgroup-source {
    width: 96px;
  }

  .member-colgroup-date {
    width: 112px;
  }

  .member-colgroup-action {
    width: 64px;
  }

  .member-count {
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

  .member-empty {
    align-items: center;
    color: var(--crs-muted);
    display: inline-flex;
    flex-direction: column;
    font-size: 13px;
    gap: 8px;
  }

  .member-single-line {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .member-name-cell {
    display: inline-block;
    max-width: 118px;
    position: relative;
    vertical-align: middle;
  }

  .member-name-text {
    display: block;
    font-weight: 600;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .member-name-cell:hover::after {
    background: var(--crs-text);
    border-radius: 6px;
    box-shadow: 0 10px 24px rgba(15, 23, 42, 0.16);
    color: #fff;
    content: attr(data-full-name);
    font-size: 12px;
    font-weight: 500;
    left: 0;
    line-height: 18px;
    max-width: 240px;
    overflow-wrap: break-word;
    padding: 6px 8px;
    position: absolute;
    top: calc(100% + 8px);
    white-space: normal;
    width: max-content;
    z-index: 90;
  }

  .member-level-badge {
    align-items: center;
    background: #f8fbff;
    border: 1px solid rgba(18, 143, 200, 0.14);
    border-radius: 999px;
    color: var(--crs-primary-strong);
    display: inline-flex;
    height: 26px;
    justify-content: center;
    line-height: 1;
    overflow: hidden;
    padding: 0 7px;
    width: 34px;
  }

  .member-level-badge-img {
    display: block;
    max-height: 18px;
    max-width: 22px;
    object-fit: contain;
  }

  .member-action-menu {
    display: inline-flex;
    justify-content: center;
    position: relative;
    width: 100%;
  }

  .member-kebab {
    align-items: center;
    background: transparent;
    border: 0;
    border-radius: 8px;
    box-shadow: none;
    color: var(--crs-muted);
    display: inline-flex;
    height: 32px;
    justify-content: center;
    padding: 0;
    transition: background 0.15s ease, color 0.15s ease, box-shadow 0.15s ease;
    width: 36px;
  }

  .member-kebab:hover,
  .member-kebab:focus-visible {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
    box-shadow: 0 0 0 3px rgba(56, 174, 234, 0.14);
  }

  .member-action-popover {
    background: var(--crs-surface);
    border: 1px solid var(--crs-border);
    border-radius: 8px;
    box-shadow: 0 18px 36px rgba(15, 23, 42, 0.12);
    color: var(--crs-text);
    min-width: 176px;
    padding: 8px 0;
    position: absolute;
    right: 0;
    text-align: left;
    top: calc(100% + 8px);
    z-index: 80;
  }

  .member-menu-section {
    display: grid;
    gap: 2px;
    padding: 0 0;
  }

  .member-menu-title {
    color: var(--crs-muted);
    font-size: 12px;
    font-weight: 700;
    line-height: 18px;
    padding: 4px 14px 5px;
  }

  .member-menu-item {
    background: transparent;
    border: 0;
    color: var(--crs-text);
    display: block;
    font-size: 13px;
    font-weight: 500;
    line-height: 20px;
    min-height: 32px;
    padding: 6px 14px;
    text-align: left;
    transition: background 0.15s ease, color 0.15s ease;
    width: 100%;
  }

  .member-menu-item:hover {
    background: #f4f8fd;
    color: var(--crs-primary-strong);
  }

  .member-menu-danger {
    color: var(--crs-danger);
  }

  .member-menu-danger:hover {
    background: var(--crs-danger-soft);
    color: #dc2626;
  }

  .member-menu-divider {
    border-top: 1px solid var(--crs-border);
    margin: 6px 0;
  }

  @media (min-width: 768px) {
    .member-filter-grid {
      grid-template-columns: repeat(3, minmax(0, 1fr));
    }

    .member-more-grid {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  @media (min-width: 1200px) {
    .member-filter-grid {
      grid-template-columns: repeat(5, minmax(0, 1fr));
    }
  }

  @media (max-width: 640px) {
    padding: 12px;

    .member-page-header {
      align-items: flex-start;
      flex-direction: column;
    }

    .member-header-actions {
      width: 100%;
    }

    .member-header-actions :deep(button) {
      flex: 1;
    }

    .member-panel-section {
      padding: 12px;
    }

    .member-filter-grid {
      grid-template-columns: repeat(1, minmax(0, 1fr));
    }
  }

  :deep(.border-border),
  .border-border {
    border-color: var(--crs-border);
  }

  :deep(.text-muted-foreground),
  .text-muted-foreground {
    color: var(--crs-muted);
  }

  :deep(.text-foreground),
  .text-foreground {
    color: var(--crs-text);
  }

  :deep(.bg-white),
  .bg-white {
    background: var(--crs-surface);
  }

  :deep(.bg-muted),
  .bg-muted {
    background: var(--crs-primary-soft);
    color: var(--crs-primary-strong);
  }

  :deep(.bg-primary),
  .bg-primary {
    background: var(--crs-primary);
  }

  :deep(.text-primary) {
    color: var(--crs-primary-strong);
  }

  :deep(.border-primary) {
    border-color: var(--crs-primary);
  }

  :deep(input),
  :deep(button),
  select {
    border-color: var(--crs-border);
  }

  :deep(input:focus),
  :deep(button:focus-visible),
  select:focus {
    border-color: var(--crs-primary);
    box-shadow: 0 0 0 2px rgba(56, 174, 234, 0.16);
  }

  thead tr {
    background: var(--crs-surface-soft) !important;
  }

  tbody tr:hover {
    background: var(--crs-surface-soft) !important;
  }

  .member-link,
  .member-action-primary {
    color: var(--crs-primary-strong);
  }

  .member-action {
    align-items: center;
    border-radius: 6px;
    display: inline-flex;
    font-size: 13px;
    font-weight: 600;
    gap: 4px;
    height: 28px;
    padding: 0 7px;
    transition: background 0.15s ease, color 0.15s ease;
    white-space: nowrap;
  }

  :deep([class*='text-stripe']) {
    color: var(--crs-primary-strong);
  }

  .member-link:hover,
  .member-action-primary:hover,
  :deep([class*='text-stripe']:hover) {
    color: #0879ad;
  }

  .member-action-primary:hover {
    background: var(--crs-primary-soft);
  }

  .member-action-danger {
    color: var(--crs-danger);
  }

  .member-action-danger:hover {
    background: var(--crs-danger-soft);
    color: #dc2626;
  }

  .member-status {
    border: 1px solid transparent;
    box-shadow: none;
  }

  .member-status-ok {
    background: var(--crs-success-soft);
    border-color: rgba(22, 163, 74, 0.18);
    color: var(--crs-success);
  }

  .member-status-danger {
    background: var(--crs-danger-soft);
    border-color: rgba(239, 68, 68, 0.18);
    color: var(--crs-danger);
  }
}
</style>
