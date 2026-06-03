<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">预约订单</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane name="WAIT_CONFIRM" tab="待确认">
        </n-tab-pane>
        <n-tab-pane name="WAIT_SERVE" tab="待服务">
        </n-tab-pane>
        <n-tab-pane name="SERVING" tab="服务中">
        </n-tab-pane>
        <n-tab-pane name="DONE" tab="已完成">
        </n-tab-pane>
        <n-tab-pane name="ALL" tab="全部">
        </n-tab-pane>
        <n-tab-pane name="ABNORMAL" tab="异常">
        </n-tab-pane>
        <n-tab-pane name="CANCEL" tab="取消">
        </n-tab-pane>
      </n-tabs>
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef"  :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"   >
        <template #tableTitle>
          <n-button type="primary" @click="handleExport" class="min-left-space"
                    v-if="hasPermission(['/spaOrder/export'])">
            <template #icon>
              <n-icon>
                <ExportOutlined/>
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" @reloadTable="reloadTable"/>
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref, computed, onMounted} from 'vue';
import {NButton, useDialog, useMessage} from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import {BasicForm, useForm} from '@/components/Form/index';
import { List, ConfirmAgree, ExportOrder } from '@/api/spaOrder';
import { PrinterSpaOrder } from '@/api/printer';
import { columns, schemas, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import View from './view.vue';
import {ExportOutlined} from "@vicons/antd";
import {usePermission} from "@/hooks/web/usePermission";

const dialog = useDialog();
const {hasPermission} = usePermission();
const message = useMessage();
const tabValue = ref('WAIT_CONFIRM')
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();
const statusRef = ref();
const dispatchRef = ref();
const cancelRef = ref();

const actionColumn = reactive({
  width: 100,
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
        },
        // {
        //   label: '确认',
        //   onClick: handleAgree.bind(null, record),
        //   ifShow: () => {
        //     return record.orderStatus === "WAIT_CONFIRM";
        //   },
        //   auth: ['/spaOrder/confirmAgree'],
        //   type: 'success'
        // },
        // {
        //   label: '拒单',
        //   onClick: handleDisagree.bind(null, record),
        //   ifShow: () => {
        //     return record.orderStatus === "WAIT_CONFIRM";
        //   },
        //   auth: ['/spaOrder/confirmDisagree'],
        //   type: 'error'
        // },
        // {
        //   label: '订单调度',
        //   onClick: handleDispatch.bind(null, record),
        //   ifShow: () => {
        //     return record.orderStatus === "WAIT_SERVE" && record.dispatchStatus === "WAIT";
        //   },
        //   auth: ['/spaOrder/dispatch'],
        //   type: 'error'
        // },
        // {
        //   label: '退款',
        //   onClick: handleCancel.bind(null, record),
        //   ifShow: () => {
        //     return record.payStatus === "HAVE_PAID" && record.orderStatus !== "DONE" && record.orderStatus !== "SERVING" && record.adminCancelNum == 0;
        //   },
        //   auth: ['/spaOrder/cancelPay'],
        //   type: 'error'
        // },
        // {
        //   label: '打印',
        //   onClick: handlePrinterOrder.bind(null, record),
        //   ifShow: () => {
        //     return record.payTime != null && record.payStatus === "HAVE_PAID";
        //   },
        //   auth: ['/printer/printSpaOrder'],
        //   type: 'info'
        // },
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
  res.orderStatus = tabValue.value;
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleUpdateValue(e){
  tabValue.value = e;
  reloadTable()
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}

// 导出
function handleExport() {
  let messageReactive = message.loading('正在导出列表...', {duration: 0});
  ExportOrder({ ...searchFormRef.value?.formModel, ...{orderStatus: tabValue.value} }).then((_res) => {
    messageReactive.destroy()
    message.success('请稍后在导出记录中进行下载');
  });
}

// 订单确认
function handleAgree(record: Recordable) {

  dialog.info({
    title: '提示',
    content: '你确定要接单吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      ConfirmAgree({id: record.id}).then((_res) => {
        message.success('操作成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {},
  });

}

// 修改状态
function handleDisagree(record: Recordable) {
  statusRef.value.openModal(record.id);
}

// 调度
function handleDispatch(record: Recordable){
  dispatchRef.value.openModal(record.id,record.goodsNum);
}

// 取消
function handleCancel(record: Recordable){
  cancelRef.value.openModal(record);
}

// 打印
function handlePrinterOrder(record: Recordable) {

  dialog.info({
    title: '提示',
    content: '你确定要打印该订单吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      PrinterSpaOrder({orderId: record.id}).then((_res) => {
        message.success('操作成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {},
  });

}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

