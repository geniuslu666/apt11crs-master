<template>
  <div>
    <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }" :content-style="{
                    padding: '0 20px 20px',
                  }">
          <template #header>
            <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">部门管理</text>
          </template>
          <n-space vertical :size="12">
            <n-space>
              <n-button type="primary" @click="addTable" v-if="hasPermission(['/employeeDepartment/edit'])">
                <template #icon>
                  <n-icon>
                    <PlusOutlined />
                  </n-icon>
                </template>
                添加部门
              </n-button>
            </n-space>

            <n-data-table
              :columns="columns.concat(actionColumn)"
              :data="dataSource"
              :row-key="(row) => row.id"
              :loading="loading"
              :resizeHeightOffset="-20000"
              :default-expand-all="true"
            />
          </n-space>
      </n-card>

    <Edit @reloadTable="reloadTable" ref="editRef" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { TableAction } from '@/components/Table';
  import { usePermission } from '@/hooks/web/usePermission';
  import { List as DepartmentList, Delete, Switch } from '@/api/employeeDepartment';
  import { columns, loadOptions } from './model';
  import {
    PlusOutlined
  } from '@vicons/antd';
  import Edit from './edit.vue';

  const { hasPermission } = usePermission();
  const dialog = useDialog();
  const message = useMessage();
  const loading = ref(false);
  const dataSource = ref<any>([]);

  const editRef = ref();

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
            ifShow: () => {
              return hasPermission(['/employeeDepartment/edit']);
            },
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/employeeDepartment/switch'],
          },
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
            auth: ['/employeeDepartment/switch'],
          },
          {
            label: '添加子部门',
            onClick: handleAddChild.bind(null, record),
            ifShow: () => {
              return hasPermission(['/employeeDepartment/edit']);
            },
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            ifShow: () => {
              return hasPermission(['/employeeDepartment/delete']);
            },
          },
        ],
      });
    },
  });

  async function loadDataTable() {
    const res = await DepartmentList({ pageSize: 100, page: 1, withChildrenTotal: true });
    dataSource.value = res.list ?? [];
  }

  async function reloadTable() {
    await loadDataTable();
  }

  function addTable() {
    editRef.value.openModal(null, '添加部门');
  }

  function handleEdit(record: Recordable) {
    editRef.value.openModal(record, '编辑部门');
  }

  function handleAddChild(record: Recordable) {
    editRef.value.openModal(null, '添加子部门', record);
  }

  // 修改状态
function handleStatus(record: Recordable, status: number) {
  Switch({ id: record.id, status: status }).then((_res) => {
    message.success('修改成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: `您确定要删除部门：${record.name}？`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: record.id }).then(async (_res) => {
          message.success('删除成功');
          await reloadTable();
        });
      },
    });
  }

  onMounted(async () => {
    loading.value = true;
    await loadOptions();
    await loadDataTable();
    loading.value = false;
  });
</script>
