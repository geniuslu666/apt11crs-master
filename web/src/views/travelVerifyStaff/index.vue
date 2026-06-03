<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{ padding: '20px' }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">核销人员</text>
        </template>
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{ padding: '0 20px 20px' }">
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
          <n-button type="primary" @click="handleAdd" class="min-left-space" v-if="hasPermission(['/travel/verifyStaff/edit'])">
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
import {h, reactive, ref, computed, onMounted} from 'vue';
import { NButton, NTag, useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Delete, Status } from '@/api/travelVerifyStaff';
import { PlusOutlined } from '@vicons/antd';
import { schemas, rules, State, newState, loadOptions, options } from './model';
import {adaTableScrollX, getOptionLabel, getOptionTag} from '@/utils/hotgo';
import Edit from "@/views/travelVerifyStaff/edit.vue";
import {isNullObject} from "@/utils/is";

const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const formRef = ref();
const submitLoading = ref(false);
const formValue = ref<State>(newState(null));
const editRef = ref();

const columns = [
  { title: '姓名', key: 'name', align: 'left', width: 150 },
  { title: '电话', key: 'mobile', align: 'left', width: 150 },
  { title: '登录账号', key: 'username', align: 'left', width: 180 },
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
  width: 140,
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
          auth: ['/travel/verifyStaff/edit'],
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1;
          },
          auth: ['/travel/verifyStaff/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2;
          },
          auth: ['/travel/verifyStaff/status'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/travel/verifyStaff/delete'],
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
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

function reloadTable() {
  actionRef.value?.reload();
}

function handleAdd() {
  editRef.value.openModal({ id: 0 });
}

function handleEdit(record) {
  editRef.value.openModal(record);
}

function handleDelete(record) {
  dialog.warning({
    title: '警告',
    content: '确定要删除该核销人员吗？',
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
