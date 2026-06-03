<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '30px 20px',
                  }">
        <template #header>
          <n-divider title-placement="left" style="margin: 0">
            <span style="font-size: 18px;color: #3D3D3D;line-height: 25px;font-weight: 500">员工管理</span>
          </n-divider>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
        <template #createdAtSlot="{ model, field }">
          <n-date-picker 
            v-model:formatted-value="model[field]" 
            type="datetimerange" 
            value-format="yyyy-MM-dd HH:mm:ss"
            clearable
            :shortcuts="defRangeShortcuts()"
            style="width: 100%"
          />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" openChecked :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/pmsStaff/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/pmsStaff/delete'])">
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
    <Edit ref="editRef" @reloadTable="reloadTable" />
    <Bind ref="bindRef" @reloadTable="reloadTable" />
    <View ref="viewRef" @reloadTable="reloadTable"/>
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Delete, Status, Unbind } from '@/api/pmsStaff';
import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
import { columns, schemas, options, loadOptions } from './model';
import { adaTableScrollX, getOptionLabel } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import Edit from '@/views/pmsStaff/edit.vue';
import Bind from '@/views/pmsStaff/bind.vue';
import {useRouter} from "vue-router";
import View from "@/views/pmsStaff/view.vue";

const dialog = useDialog();
const router = useRouter();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const bindRef = ref();
const viewRef = ref();
const checkedIds = ref([]);

const actionColumn = reactive({
  width: 360,
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
          auth: ['/pmsStaff/edit'],
        },

        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1;
          },
          auth: ['/pmsStaff/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2;
          },
          auth: ['/pmsStaff/status'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/pmsStaff/delete'],
        },
        {
          label: '绑定会员',
          onClick: handleBind.bind(null, record),
          ifShow: () => {
            return record.pmsMemberId === 0;
          },
          auth: ['/pmsStaff/bind'],
          color: '#3F9EFF'
        },
        {
          label: '解绑会员',
          onClick: handleUnbind.bind(null, record),
          ifShow: () => {
            return record.pmsMemberId > 0;
          },
          auth: ['/pmsStaff/unbind'],
          type: 'warning'
        },
        {
          label: '详情',
          onClick: handleView.bind(null, record),
          auth: ['/pmsStaff/view'],
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
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

// 绑定会员
function handleBind(record: Recordable) {
  bindRef.value.openModal(record.id, record.pmsMemberId);
}

// 解绑会员
function handleUnbind(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要解绑？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Unbind(record).then((_res) => {
        message.success('解绑成功');
        reloadTable();
      });
    },
  });
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record.id);
  // router.push({ name: 'pmsStaff_view', params: { id: record.id } });
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

