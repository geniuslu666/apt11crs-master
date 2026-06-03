<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <BasicTable  ref="actionRef" openChecked :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/foodRestaurantNotice/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/foodRestaurantNotice/delete'])">
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" :restaurantId="props.restaurantId"/>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Delete, Status } from '@/api/foodRestaurantNotice';
import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
import { columns, options, loadOptions } from './model';
import { adaTableScrollX, getOptionLabel } from '@/utils/hotgo';
import Edit from './edit.vue';
import View from './view.vue';

const props = defineProps({
  restaurantId: {
    type: Number,
    default: 0,
  },
});

const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const viewRef = ref();
const checkedIds = ref([]);

const actionColumn = reactive({
  width: 288,
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
          auth: ['/foodRestaurantNotice/edit'],
        },

        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1;
          },
          auth: ['/foodRestaurantNotice/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2;
          },
          auth: ['/foodRestaurantNotice/status'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/foodRestaurantNotice/delete'],
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

// 加载表格数据
const loadDataTable = async (res) => {
  res.restaurantId = props.restaurantId
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 添加数据
function addTable() {
  editRef.value.openModal(null);
}

// 编辑数据
function handleEdit(record: Recordable) {
  editRef.value.openModal(record);
}

// 单个删除
function handleDelete(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete(record).then((_res) => {
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

// 批量删除
function handleBatchDelete() {
  if (checkedIds.value.length < 1){
    message.error('请至少选择一项要删除的数据');
    return;
  }

  dialog.warning({
    title: '警告',
    content: '你确定要批量删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({ id: checkedIds.value }).then((_res) => {
        checkedIds.value = [];
        message.success('删除成功');
        reloadTable();
      });
    },
  });
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

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

