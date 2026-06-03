<template>
  <n-drawer v-model:show="showModal" :width="960">
    <n-drawer-content closable :header-style="{ padding: '20px' }" :body-content-style="{ padding: '20px' }">
      <template #header>
        <div style="font-weight: 500; font-size: 18px; color: #3D3D3D; line-height: 25px;">
          券领取记录 - {{ activityName }}
        </div>
      </template>

      <div style="margin-bottom: 16px; display: flex; gap: 12px; flex-wrap: wrap; align-items: center;">
        <n-select
          v-model:value="searchForm.issueType"
          :options="issueTypeOptions"
          placeholder="发放类型"
          clearable
          style="width: 140px"
          @update:value="handleSearch"
        />
        <n-select
          v-model:value="searchForm.state"
          :options="stateOptions"
          placeholder="券状态"
          clearable
          style="width: 140px"
          @update:value="handleSearch"
        />
        <n-input
          v-model:value="searchForm.keyword"
          placeholder="员工姓名/券号"
          clearable
          style="width: 180px"
          @keyup.enter="handleSearch"
        />
        <n-button type="primary" @click="handleSearch">搜索</n-button>
      </div>
      <!-- PLACEHOLDER_TABLE -->
      <n-data-table
        :columns="tableColumns"
        :data="tableData"
        :loading="loading"
        :pagination="pagination"
        :row-key="(row) => row.id"
        :scroll-x="1200"
        :row-props="() => ({ style: 'height: 52px' })"
        remote
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </n-drawer-content>
  </n-drawer>
</template>

<script lang="ts" setup>
import { h, ref, reactive } from 'vue';
import { NTag } from 'naive-ui';
import { CouponRecordList } from '@/api/employeeActivity';

const emit = defineEmits(['reloadTable']);
const showModal = ref(false);
const loading = ref(false);
const activityName = ref('');
const activityId = ref(0);
const tableData = ref<any[]>([]);

const searchForm = reactive({
  issueType: null as number | null,
  state: null as number | null,
  keyword: '',
});

const pagination = reactive({
  page: 1,
  pageSize: 10,
  pageCount: 1,
  itemCount: 0,
  showSizePicker: true,
  pageSizes: [10, 20, 50],
});

const issueTypeOptions = [
  { label: '自动领取', value: 1 },
  { label: '批量发放', value: 2 },
];

const stateOptions = [
  { label: '待生效', value: 1 },
  { label: '未使用', value: 2 },
  { label: '已核销', value: 3 },
  { label: '已过期', value: 4 },
  { label: '已失效', value: 5 },
  { label: '已回收', value: 6 },
];

const stateMap: Record<number, { label: string; type: string }> = {
  1: { label: '待生效', type: 'warning' },
  2: { label: '未使用', type: 'info' },
  3: { label: '已核销', type: 'success' },
  4: { label: '已过期', type: 'default' },
  5: { label: '已失效', type: 'error' },
  6: { label: '已回收', type: 'default' },
};

// Table columns
const tableColumns = [
  { title: 'ID', key: 'id', width: 70 },
  { title: '券号', key: 'couponNo', width: 140, ellipsis: { tooltip: true } },
  { title: '券名称', key: 'couponName', width: 130, ellipsis: { tooltip: true } },
  { title: '员工', key: 'employeeName', width: 100 },
  {
    title: '发放类型',
    key: 'issueType',
    width: 100,
    render(row: any) {
      return h(NTag, {
        type: row.issueType === 2 ? 'warning' : 'info',
        bordered: false,
        size: 'small',
      }, { default: () => row.issueType === 2 ? '批量发放' : '自动领取' });
    },
  },
  {
    title: '状态',
    key: 'state',
    width: 90,
    render(row: any) {
      const s = stateMap[row.state] || { label: '未知', type: 'default' };
      return h(NTag, { type: s.type as any, bordered: false, size: 'small' }, { default: () => s.label });
    },
  },
  { title: '有效期开始', key: 'startTime', width: 160 },
  { title: '有效期结束', key: 'endTime', width: 160 },
  { title: '核销时间', key: 'verifyTime', width: 160 },
  { title: '核销门店', key: 'storeName', width: 120, ellipsis: { tooltip: true } },
  { title: '领取时间', key: 'createAt', width: 160 },
];

function openModal(record: any) {
  activityId.value = record.id;
  activityName.value = record.name || '';
  // 重置搜索
  searchForm.issueType = null;
  searchForm.state = null;
  searchForm.keyword = '';
  pagination.page = 1;
  showModal.value = true;
  loadData();
}

async function loadData() {
  loading.value = true;
  try {
    const params: any = {
      activityId: activityId.value,
      page: pagination.page,
      pageSize: pagination.pageSize,
    };
    if (searchForm.issueType) params.issueType = searchForm.issueType;
    if (searchForm.state) params.state = searchForm.state;
    if (searchForm.keyword) params.keyword = searchForm.keyword;

    const res = await CouponRecordList(params);
    tableData.value = res.list || [];
    pagination.itemCount = res.totalCount || 0;
    pagination.pageCount = res.pageCount || 1;
  } finally {
    loading.value = false;
  }
}

function handleSearch() {
  pagination.page = 1;
  loadData();
}

function handlePageChange(page: number) {
  pagination.page = page;
  loadData();
}

function handlePageSizeChange(pageSize: number) {
  pagination.pageSize = pageSize;
  pagination.page = 1;
  loadData();
}

defineExpose({ openModal });
</script>
