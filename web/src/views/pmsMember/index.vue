<template>
  <div class="p-4 space-y-3">

    <!-- Search card -->
    <div class="bg-white rounded-lg border border-border">
      <div class="p-4">
        <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-3">
          <div>
            <div class="text-[12px] font-medium text-muted-foreground mb-1">会员号</div>
            <UiInput v-model="searchForm.memberNo" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div>
            <div class="text-[12px] font-medium text-muted-foreground mb-1">手机号</div>
            <UiInput v-model="searchForm.phone" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div>
            <div class="text-[12px] font-medium text-muted-foreground mb-1">昵称</div>
            <UiInput v-model="searchForm.fullName" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div>
            <div class="text-[12px] font-medium text-muted-foreground mb-1">邮箱</div>
            <UiInput v-model="searchForm.mail" placeholder="请输入" @keyup.enter="handleSearch" />
          </div>
          <div>
            <div class="text-[12px] font-medium text-muted-foreground mb-1">会员等级</div>
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
        <div v-show="showMore" class="grid grid-cols-1 sm:grid-cols-2 gap-3 mt-3 pt-3 border-t border-border/40">
          <div>
            <div class="text-[12px] font-medium text-muted-foreground mb-1">注册时间</div>
            <div class="flex items-center gap-1.5">
              <input
                type="datetime-local"
                v-model="searchForm.createdAtStart"
                class="flex-1 h-8 rounded-md border border-input bg-background px-2 text-[13px] transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring min-w-0"
              />
              <span class="text-muted-foreground text-[12px] flex-shrink-0">—</span>
              <input
                type="datetime-local"
                v-model="searchForm.createdAtEnd"
                class="flex-1 h-8 rounded-md border border-input bg-background px-2 text-[13px] transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring min-w-0"
              />
            </div>
          </div>
          <div>
            <div class="text-[12px] font-medium text-muted-foreground mb-1">最后登录</div>
            <div class="flex items-center gap-1.5">
              <input
                type="datetime-local"
                v-model="searchForm.lastLoginStart"
                class="flex-1 h-8 rounded-md border border-input bg-background px-2 text-[13px] transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring min-w-0"
              />
              <span class="text-muted-foreground text-[12px] flex-shrink-0">—</span>
              <input
                type="datetime-local"
                v-model="searchForm.lastLoginEnd"
                class="flex-1 h-8 rounded-md border border-input bg-background px-2 text-[13px] transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring min-w-0"
              />
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center justify-between mt-3 pt-3 border-t border-border/40">
          <div class="flex items-center gap-2">
            <UiButton size="sm" @click="handleSearch">
              <svg class="w-3.5 h-3.5 mr-1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.35-4.35"/></svg>
              查询
            </UiButton>
            <UiButton size="sm" variant="outline" @click="handleReset">重置</UiButton>
            <UiButton size="sm" variant="outline" @click="handleExport" v-if="hasPermission(['/pmsMember/export'])">导出</UiButton>
          </div>
          <button
            type="button"
            @click="showMore = !showMore"
            class="inline-flex items-center gap-1 text-[12px] text-muted-foreground hover:text-foreground transition-colors"
          >
            {{ showMore ? '收起' : '更多筛选' }}
            <svg
              :class="['transition-transform duration-200', showMore ? 'rotate-180' : '']"
              xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24"
              fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
            ><path d="m6 9 6 6 6-6"/></svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Table card -->
    <div class="bg-white rounded-lg border border-border overflow-hidden">
      <!-- Toolbar -->
      <div class="flex items-center justify-between px-4 py-2.5 border-b border-border">
        <div class="flex items-center gap-2">
          <span class="text-[13px] font-semibold text-foreground">会员列表</span>
          <span
            v-if="totalCount > 0"
            class="inline-flex items-center h-5 px-1.5 text-[11px] font-medium bg-muted text-muted-foreground rounded"
          >{{ totalCount }}</span>
        </div>
        <button
          type="button"
          @click="loadData"
          :disabled="loading"
          class="inline-flex items-center gap-1 text-[12px] text-muted-foreground hover:text-foreground disabled:opacity-50 transition-colors"
        >
          <svg :class="['w-3.5 h-3.5', loading ? 'animate-spin' : '']" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12a9 9 0 1 1-6.219-8.56"/></svg>
          刷新
        </button>
      </div>

      <!-- Table -->
      <div class="overflow-x-auto relative">
        <!-- Loading overlay -->
        <div v-if="loading" class="absolute inset-0 bg-white/70 z-10 flex items-center justify-center">
          <svg class="animate-spin w-5 h-5 text-primary" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
          </svg>
        </div>

        <table class="w-full text-[13px]">
          <thead>
            <tr class="border-b border-border bg-[#f8fafc]">
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">会员信息</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">会员分组</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">会员等级</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">推荐信息</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">手机/邮箱</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">积分/经验</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">状态</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">注册来源</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">上次登录</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">创建时间</th>
              <th class="px-3 py-2.5 text-left text-[11px] font-medium text-muted-foreground whitespace-nowrap">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-border">
            <tr v-if="data.length === 0 && !loading">
              <td colspan="11" class="px-3 py-12 text-center text-muted-foreground text-[13px]">暂无数据</td>
            </tr>
            <tr v-for="row in data" :key="row.id" class="hover:bg-[#f8fafc] transition-colors">

              <!-- 会员信息 -->
              <td class="px-3 py-2 w-[180px]">
                <div class="flex items-center gap-2.5">
                  <img
                    :src="row.avatar || defaultImg"
                    @error="(e: any) => (e.target.src = defaultImg)"
                    class="w-8 h-8 rounded-full flex-shrink-0 object-cover bg-muted"
                  />
                  <div class="min-w-0">
                    <div class="text-stripe-600 font-medium truncate">{{ row.memberNo }}</div>
                    <div class="text-[12px] text-muted-foreground truncate">{{ (row.lastName || '') + (row.firstName || '') || '—' }}</div>
                  </div>
                </div>
              </td>

              <!-- 会员分组 -->
              <td class="px-3 py-2 text-muted-foreground whitespace-nowrap">
                {{ row.groupId > 0 && row.memberGroup ? row.memberGroup.memberGroup : '—' }}
              </td>

              <!-- 会员等级 -->
              <td class="px-3 py-2 whitespace-nowrap">
                {{ row.memberLevel?.levelName || '—' }}
              </td>

              <!-- 推荐信息 -->
              <td class="px-3 py-2 text-muted-foreground whitespace-nowrap">
                {{ getReferrerText(row) }}
              </td>

              <!-- 手机/邮箱 -->
              <td class="px-3 py-2">
                <div class="space-y-0.5">
                  <div class="flex items-center gap-1">
                    <span class="inline-block text-[10px] text-muted-foreground w-7 flex-shrink-0">tel</span>
                    <span class="text-[13px]">{{ row.phone ? `${row.phoneArea}-${row.phone}` : '—' }}</span>
                  </div>
                  <div class="flex items-center gap-1">
                    <span class="inline-block text-[10px] text-muted-foreground w-7 flex-shrink-0">mail</span>
                    <span class="text-[13px] truncate max-w-[140px]">{{ row.mail || '—' }}</span>
                  </div>
                </div>
              </td>

              <!-- 积分/经验 -->
              <td class="px-3 py-2 whitespace-nowrap tabular-nums">
                {{ row.balance }}<span class="text-muted-foreground mx-0.5 text-[11px]">/</span>{{ row.exp }}
              </td>

              <!-- 状态 -->
              <td class="px-3 py-2">
                <span :class="[
                  'inline-flex items-center px-1.5 py-0.5 rounded text-[11px] font-medium',
                  row.status === 1
                    ? 'bg-green-50 text-green-700 ring-1 ring-green-600/20'
                    : 'bg-red-50 text-red-700 ring-1 ring-red-600/20'
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
                {{ row.lastLogin || '—' }}
              </td>

              <!-- 创建时间 -->
              <td class="px-3 py-2 text-[12px] text-muted-foreground whitespace-nowrap">
                {{ row.createdAt || '—' }}
              </td>

              <!-- 操作 -->
              <td class="px-3 py-2">
                <div class="flex items-center gap-2 flex-wrap">
                  <button
                    v-if="hasPermission(['/pmsMember/view'])"
                    type="button"
                    class="text-[13px] text-stripe-600 hover:text-stripe-800 transition-colors whitespace-nowrap"
                    @click="handleView(row)"
                  >详情</button>
                  <button
                    v-if="row.status === 1 && hasPermission(['/pmsMember/status'])"
                    type="button"
                    class="text-[13px] text-orange-500 hover:text-orange-700 transition-colors whitespace-nowrap"
                    @click="handleStatus(row, 2)"
                  >禁用</button>
                  <button
                    v-if="row.status === 2 && hasPermission(['/pmsMember/status'])"
                    type="button"
                    class="text-[13px] text-green-600 hover:text-green-800 transition-colors whitespace-nowrap"
                    @click="handleStatus(row, 1)"
                  >启用</button>
                  <button
                    v-if="row.memberCancelArr == null && hasPermission(['/pmsWithdraw/disagreeStaff'])"
                    type="button"
                    class="text-[13px] text-red-500 hover:text-red-700 transition-colors whitespace-nowrap"
                    @click="handleCancel(row)"
                  >注销</button>
                  <UiDropdownMenu
                    :items="dropdownItems"
                    label="更多"
                    @select="(key) => handleDropdownSelect(key, row)"
                  />
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
            <svg xmlns="http://www.w3.org/2000/svg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>
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
            <svg xmlns="http://www.w3.org/2000/svg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
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
import { reactive, ref, computed, onMounted } from 'vue';
import { useMessage } from 'naive-ui';
import { UiButton, UiInput, UiSelect, UiDropdownMenu } from '@/components/ui';
import { List, Status, Export } from '@/api/pmsMember';
import { options, loadOptions, levelList, recommendModel } from './model';
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
import defaultImg from '@/assets/images/mrtx.png';

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
const showMore = ref(false);

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

// Column helpers
function getReferrerText(row: any): string {
  if (row.referrer <= 0 && row.lastReferrer <= 0) return '无';
  if (recommendModel.value === 'FIRST') {
    if (!row.referrerDetail) return '无';
    const { rebateMode, memberNo } = row.referrerDetail;
    return rebateMode === 'CHANNEL' ? `渠道：${memberNo}` : rebateMode === 'STAFF' ? `员工：${memberNo}` : `会员：${memberNo}`;
  } else {
    if (!row.LastReferrerDetail) return '无';
    const { rebateMode, memberNo } = row.LastReferrerDetail;
    return rebateMode === 'CHANNEL' ? `渠道：${memberNo}` : rebateMode === 'STAFF' ? `员工：${memberNo}` : `会员：${memberNo}`;
  }
}

// Dropdown actions
const dropdownItems = [
  { label: '发通知', key: 'mes' },
  { label: '发邮件', key: 'mail' },
  { label: '发短信', key: 'telmes' },
  { label: '发优惠券', key: 'coupon' },
  { label: '发礼品券', key: 'thcoupon' },
  { label: 'H5分销链接', key: 'h5_link' },
];

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
});
</script>
