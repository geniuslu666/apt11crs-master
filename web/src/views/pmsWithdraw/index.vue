<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">提现申请表</text>
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
      <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle></template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" />
    <Disagree ref="disagreeRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import {StaffList, AgreeStaff, TransferStaff} from '@/api/pmsWithdraw';
import { columns, schemas } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import View from './view.vue';
import Disagree from './disagree.vue';

const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const disagreeRef = ref();
const viewRef = ref();
const checkedIds = ref([]);

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
          label: '同意',
          onClick: handleAgree.bind(null, record),
          ifShow: () => {
            return record.withdrawStatus === "WAIT";
          },
          auth: ['/pmsWithdraw/agreeStaff'],
          type: 'success'
        },
        {
          label: '拒绝',
          onClick: handleDisgree.bind(null, record),
          ifShow: () => {
            return record.withdrawStatus === "WAIT";
          },
          auth: ['/pmsWithdraw/disagreeStaff'],
          type: 'warning'
        },
        {
          label: '转账',
          onClick: handleTransfer.bind(null, record),
          ifShow: () => {
            return record.withdrawStatus === "SUCCESS" && record.transfer == 1;
          },
          auth: ['/pmsWithdraw/transferStaff'],
          type: 'info'
        },
        {
          label: '详情',
          onClick: handleView.bind(null, record),
          auth: ['/pmsWithdraw/view'],
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
  return await StaffList({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}

// 同意提现申请
function handleAgree(record: Recordable) {

  dialog.info({
    title: '提示',
    content: '你确定要同意该提现申请吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      AgreeStaff({id: record.id}).then((_res) => {
        message.success('操作成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {},
  });

}

// 拒绝提现申请
function handleDisgree(record: Recordable) {
  disagreeRef.value.openModal(record.id);

}


// 提现转账
function handleTransfer(record: Recordable) {

  dialog.info({
    title: '提示',
    content: '你确定已转账吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      TransferStaff({id: record.id}).then((_res) => {
        message.success('操作成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {},
  });

}

</script>

<style lang="less" scoped></style>

