<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">礼品券列表</text>
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
        <template #createAtSlot="{ model, field }">
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
      <BasicTable  ref="actionRef" :openChecked="false" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/thCoupon/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加礼品券
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref, computed, onMounted} from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import {List, Delete, Status, UseStatus} from '@/api/thCoupon';
import { PlusOutlined } from '@vicons/antd';
import { columns, schemas, options, loadOptions } from './model';
import {adaTableScrollX, getOptionLabel} from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import Edit from '@/views/thCoupon/edit.vue';
import View from '@/views/thCoupon/view.vue';
import {useRouter} from "vue-router";

const router = useRouter();
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const editRef = ref();
const viewRef = ref();
const checkedIds = ref([]);

const actionColumn = reactive({
  width: 290,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleView.bind(null, record),
          auth: ['/thCoupon/view'],
        },
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          auth: ['/thCoupon/edit'],
        },

        {
          label: '禁用发行',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1;
          },
          auth: ['/thCoupon/status'],
          type: 'warning'
        },
        {
          label: '启用发行',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2;
          },
          auth: ['/thCoupon/status'],
          type: 'success'
        },

        {
          label: '禁用使用',
          onClick: handleUseStatus.bind(null, record, 2),
          ifShow: () => {
            return record.useStatus === 1;
          },
          auth: ['/thCoupon/useStatus'],
          type: 'warning'
        },
        {
          label: '启用使用',
          onClick: handleUseStatus.bind(null, record, 1),
          ifShow: () => {
            return record.useStatus === 2;
          },
          auth: ['/thCoupon/useStatus'],
          type: 'success'
        },

        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          ifShow: () => {
            return record.status === -1;
          },
          auth: ['/thCoupon/delete'],
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
  labelWidth: 100,
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
  // router.push({ name: 'thCoupon_edit', params: { id: null } });
}

// 编辑数据
function handleEdit(record: Recordable) {
  editRef.value.openModal(record);
  // router.push({ name: 'thCoupon_edit', params: { id: record.id } });
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record.id);
  // router.push({ name: 'thCoupon_view', params: { id: record.id } });
}

// 修改状态
function handleStatus(record: Recordable, status: number) {
  Status({ id: record.id, status: status }).then((_res) => {
    message.success('设为' + getOptionLabel(options.value.th_coupon_status, status) + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });

}

// 修改状态
function handleUseStatus(record: Recordable, status: number) {
  UseStatus({ id: record.id, status: status }).then((_res) => {
    message.success('设为' + getOptionLabel(options.value.th_coupon_use_status, status) + '成功');
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
  loadOptions();
});
</script>

<style lang="less" scoped></style>

