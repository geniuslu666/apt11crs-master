<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{ padding: '20px' }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">一日游产品</text>
        </template>
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{ padding: '0 20px 20px' }">
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane :name="0" tab="全部" />
        <n-tab-pane :name="1" tab="已启用" />
        <n-tab-pane :name="2" tab="已禁用" />
        <n-tab-pane :name="3" tab="回收站" />
      </n-tabs>
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" />
      <BasicTable
        ref="actionRef"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :actionColumn="actionColumn"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable" class="min-left-space" v-if="hasPermission(['/travel/product/edit'])">
            <template #icon><n-icon><PlusOutlined /></n-icon></template>
            添加
          </n-button>
        </template>
      </BasicTable>
    </n-card>

    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { NButton, NTag, useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { List, RecycleList, Status, Delete, Restore } from '@/api/travelProduct';
import { PlusOutlined } from '@vicons/antd';
import { schemas, loadOptions, options } from './model';
import {adaTableScrollX, getOptionLabel, getOptionTag} from '@/utils/hotgo';
import Edit from "@/views/travelProduct/edit.vue";
import {isNullObject} from "@/utils/is";
import { useRouter } from 'vue-router';

const router = useRouter();

const tabValue = ref(0);
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();

const columns = [
  // {
  //   title: '列表图',
  //   key: 'listImage',
  //   align: 'left',
  //   width: 100,
  //   render(row) {
  //     return row.listImage
  //       ? h(NImage, { width: 80, src: row.listImage })
  //       : '--';
  //   },
  // },
  { title: '产品标题', key: 'title', align: 'left', width: 220 },
  {
    title: '售价（JPY）',
    key: 'minPrice',
    align: 'left',
    width: 120,
    render(row) { 
      // 检查是否有SKU数据（可能是skuList或hg_travel_product_sku字段）
      const skuData = row.skuList || row.hg_travel_product_sku || [];
      if (skuData && skuData.length > 0) {
        // 提取价格字段（可能是price或_price）
        const prices = skuData.map(sku => sku.price || sku._price).filter(price => price != null);
        if (prices.length > 0) {
          const minPrice = Math.min(...prices);
          return minPrice + ' JPY~';
        }
      }
      return '暂无价格';
    },
  },
  { title: '总库存', key: 'stock', align: 'left', width: 120 },
  { title: '排序', key: 'sort', align: 'left', width: 80 },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },
  { title: '创建时间', key: 'createdAt', align: 'left', width: 180 },
];

const actionColumn = reactive({
  width: 220,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          ifShow: () => tabValue.value !== 3,
          auth: ['/travel/product/edit'],
        },
        {
          label: '车型管理',
          onClick: handleSku.bind(null, record),
          ifShow: () => tabValue.value !== 3,
          auth: ['/travel/product/edit'],
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1 && tabValue.value !== 3;
          },
          auth: ['/travel/product/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2 && tabValue.value !== 3;
          },
          auth: ['/travel/product/status'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          ifShow: () => tabValue.value !== 3,
          auth: ['/travel/product/delete'],
        },
        {
          label: '恢复',
          type: 'warning',
          onClick: handleRestore.bind(null, record),
          ifShow: () => tabValue.value === 3,
          auth: ['/travel/product/restore'],
        },
      ],
    });
  },
});

const scrollX = computed(() => adaTableScrollX(columns, actionColumn.width));

const [register] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

function handleUpdateValue(v) {
  tabValue.value = v;
  reloadTable();
}

const loadDataTable = async (res) => {
  if (tabValue.value === 3) {
    return await RecycleList({ ...searchFormRef.value?.formModel, ...res });
  }
  if (tabValue.value > 0) res.status = tabValue.value;
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

function reloadTable() {
  actionRef.value?.reload();
}

// 修改状态
function handleStatus(record: Recordable, status: number) {
  Status({ id: record.id, status: status }).then((_res) => {
    message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}

function addTable() {
  editRef.value.openModal({ id: 0 });
}

function handleEdit(record) {
  editRef.value.openModal(record);
}

function handleSku(record) {
  router.push({ name: 'travelProductSkuIndex', params: { id: record.id } });
}

function handleDelete(record) {
  dialog.warning({
    title: '警告',
    content: '确定要删除该产品吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({ id: [record.id] }).then(() => {
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

function handleRestore(record) {
  dialog.warning({
    title: '提示',
    content: '确定要恢复该产品吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Restore({ id: [record.id] }).then(() => {
        message.success('恢复成功');
        reloadTable();
      });
    },
  });
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>
