<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }"
              :content-style="{
                    padding: '0 20px 20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">区域管理</text>
        </template>
        <n-space vertical :size="12">
          <n-space>
            <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/foodArea/edit'])">
              <template #icon>
                <n-icon>
                  <PlusOutlined />
                </n-icon>
              </template>
              添加
            </n-button>
          </n-space>
          <n-data-table
            v-if="dataSource.length > 0 || !loading"
            :columns="columns.concat(actionColumn)"
            :data="dataSource"
            :row-key="(row) => row.id"
            :loading="loading"
            :resizeHeightOffset="-20000"
            default-expand-all
          />
        </n-space>
      </n-card>
    </n-spin>
    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, onMounted } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { TableAction } from '@/components/Table';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Status, Delete } from '@/api/foodArea';
import { PlusOutlined } from '@vicons/antd';
import { columns, loadOptions, options } from './model';
import Edit from './edit.vue';
import { getOptionLabel } from '@/utils/hotgo';

const show = ref(false);
const loading = ref(false);
const dataSource = ref<any>([]);
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const editRef = ref();

const actionColumn = reactive({
  width: 144,
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
          auth: ['/foodArea/edit'],
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.areaStatus === 1;
          },
          auth: ['/foodArea/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.areaStatus === 2;
          },
          auth: ['/foodArea/status'],
        },

        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/foodArea/delete'],
        },
      ],

    });
  },
});

// 加载表格数据
function loadDataTable()  {
  loading.value = true;
  List({ Pagination: false }).then((res) => {
    dataSource.value = res.list ?? [];
    loading.value = false;
    show.value = false;
  });
};

// 重新加载表格数据
function reloadTable() {
  loadDataTable();
}

// 添加数据
function addTable() {
  editRef.value.openModal(null);
}

// 编辑数据
function handleEdit(record: Recordable) {
  editRef.value.openModal(record);
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

onMounted(() => {
  show.value = true;
  loadDataTable();
  loadOptions();
});
</script>

<style lang="less" scoped></style>

