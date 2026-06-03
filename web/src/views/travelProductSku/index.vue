<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{ padding: '20px' }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">
            {{ productName ? productName + '-车型管理' : '车型管理' }}
          </text>
        </template>
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{ padding: '0 20px 20px' }">
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" />
      <BasicTable
        ref="actionRef"
        openChecked
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :actionColumn="actionColumn"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="handleOnCheckedRow"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable" class="min-left-space" v-if="hasPermission(['/travel/productSku/edit'])">
            <template #icon><n-icon><PlusOutlined /></n-icon></template>
            添加
          </n-button>
          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/travel/productSku/delete'])">
            <template #icon><n-icon><DeleteOutlined /></n-icon></template>
            批量删除
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Delete, Status } from '@/api/travelProductSku';
import { View as ProductView } from '@/api/travelProduct';
import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
import { columns, schemas } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import { useRouter } from 'vue-router';
import Edit from './edit.vue';
import { jsontoobj } from '@/utils/smjcomm';

const productName = ref('');
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const checkedIds = ref([]);
const router = useRouter();
const params = router.currentRoute.value.params;

const actionColumn = reactive({
  width: 200,
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
          auth: ['/travel/productSku/edit'],
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          type: 'warning',
          ifShow: () => record.status === 1,
          auth: ['/travel/productSku/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          type: 'success',
          ifShow: () => record.status === 2,
          auth: ['/travel/productSku/status'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/travel/productSku/delete'],
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

const loadDataTable = async (res) => {
  return await List({ ...searchFormRef.value?.formModel, ...res, productId: params.id });
};

function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

function reloadTable() {
  actionRef.value?.reload();
}

function addTable() {
  editRef.value.openModal(null, params.id);
}

function handleEdit(record: Recordable) {
  editRef.value.openModal(record, record.productId);
}

function handleDelete(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete(record).then(() => {
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

function handleBatchDelete() {
  if (checkedIds.value.length < 1) {
    message.error('请至少选择一项要删除的数据');
    return;
  }
  dialog.warning({
    title: '警告',
    content: '你确定要批量删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({ id: checkedIds.value }).then(() => {
        checkedIds.value = [];
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

function handleStatus(record: Recordable, status: number) {
  Status({ id: record.id, status }).then(() => {
    message.success('操作成功');
    reloadTable();
  });
}

onMounted(() => {
  ProductView({ id: params.id, isLanguage: true }).then((res) => {
    if (res.titleLanguage) {
      const tl = jsontoobj(res.titleLanguage);
      productName.value = tl?.zh?.content || tl?.en?.content || '';
    }
  });
});
</script>

<style lang="less" scoped></style>
