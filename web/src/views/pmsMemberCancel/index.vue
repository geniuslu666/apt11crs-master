<template>
  <div>
    <div class="n-layout-page-header" style="margin: 0">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">会员注销列表</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane name="ALL" tab="全部">
        </n-tab-pane>
        <n-tab-pane name="1" tab="待审核">
        </n-tab-pane>
        <n-tab-pane name="2" tab="注销成功">
        </n-tab-pane>
        <n-tab-pane name="3" tab="审核拒绝">
        </n-tab-pane>
      </n-tabs>
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
      <BasicTable  ref="actionRef"  :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"   >
        <template #tableTitle></template>
      </BasicTable>
    </n-card>
    <Status ref="statusRef" @reloadTable="reloadTable"/>
    <Cancel ref="cancelRef" @reloadTable="reloadTable" />
    <MemberCancelAudit ref="memberCancelAuditRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref, computed, onMounted} from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import {BasicForm, useForm} from '@/components/Form/index';
import { List } from '@/api/pmsMemberCancel';
import { columns, schemas, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';

import Status from "@/views/foodOrder/order_confirm.vue";
import Cancel from "@/views/foodOrder/cancel.vue";
import {useRouter} from "vue-router";
import MemberCancelAudit from "@/views/pmsMemberCancel/audit_view.vue";

const router = useRouter();
const dialog = useDialog();
const message = useMessage();
const tabValue = ref('ALL')
const actionRef = ref();
const searchFormRef = ref<any>({});
const statusRef = ref();
const cancelRef = ref();
const memberCancelAuditRef = ref();

const actionColumn = reactive({
  width: 250,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '查看',
          onClick: handleCancel.bind(null, record),
          auth: ['/pmsWithdraw/disagreeStaff'],
        },
        {
          label: '会员信息',
          onClick: handleMemberView.bind(null, record),
          ifShow: () => {
            return record.auditStatus != 2;
          },
          auth: ['/pmsMember/view'],
          type: 'success'
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
  if(searchFormRef.value?.formModel.auditStatus){
    tabValue.value = searchFormRef.value?.formModel.auditStatus
  }
  return await List({ ...searchFormRef.value?.formModel, ...res, auditStatus: tabValue.value });
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleUpdateValue(e){
  tabValue.value = e;
  searchFormRef.value?.setFieldsValue({
    auditStatus: e
  });
  reloadTable()
}

// 查看会员详情
function handleMemberView(record: Recordable) {
  router.push({ name: 'pmsMember_view', params: { id: record.memberId } });
}


// 查看
function handleCancel(record: Recordable){
  memberCancelAuditRef.value.openModal(record);
}


onMounted(() => {
  // loadOptions();
});
</script>

<style lang="less" scoped></style>

