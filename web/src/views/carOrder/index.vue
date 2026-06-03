<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
        padding: '20px',
      }">
        <template #header>
          <span style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">预约订单</span>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
      padding: '0 20px 20px',
    }">
      <!--      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">-->
      <!--&lt;!&ndash;        <n-tab-pane name="WAIT_CONFIRM" tab="待接单">&ndash;&gt;-->
      <!--&lt;!&ndash;        </n-tab-pane>&ndash;&gt;-->
      <!--        <n-tab-pane name="WAIT_SERVE" tab="待服务">-->
      <!--        </n-tab-pane>-->
      <!--        <n-tab-pane name="SERVING" tab="服务中">-->
      <!--        </n-tab-pane>-->
      <!--        <n-tab-pane name="DONE" tab="已完成">-->
      <!--        </n-tab-pane>-->
      <!--        <n-tab-pane name="ALL" tab="全部">-->
      <!--        </n-tab-pane>-->
      <!--      </n-tabs>-->
      <div class="statusTab hs">
        <div :class="tabValue == 'WAIT_CONFIRM' ? 'active' : ''"><span
            @click="handleUpdateValue('WAIT_CONFIRM')">待接单</span>
        </div>
        <div :class="tabValue == 'WAIT_SERVE' ? 'active' : ''"><span @click="handleUpdateValue('WAIT_SERVE')">待服务</span>
        </div>
        <div :class="tabValue == 'SERVING' ? 'active' : ''"><span @click="handleUpdateValue('SERVING')">服务中</span></div>
        <div :class="tabValue == 'DONE' ? 'active' : ''"><span @click="handleUpdateValue('DONE')">已完成</span></div>
        <div :class="tabValue == 'ALL' ? 'active' : ''"><span @click="handleUpdateValue('ALL')">全部</span></div>
        <div :class="tabValue == 'ABNORMAL' ? 'active' : ''"><span @click="handleUpdateValue('ABNORMAL')">异常</span>
        </div>
        <div :class="tabValue == 'CANCEL' ? 'active' : ''"><span @click="handleUpdateValue('CANCEL')">已取消</span></div>
      </div>
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable"
        @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
        <template #bookStartTimeSlot="{ model, field }">
          <n-date-picker 
            v-model:formatted-value="model[field]" 
            type="datetimerange" 
            value-format="yyyy-MM-dd HH:mm:ss"
            clearable
            :shortcuts="defRangeShortcuts()"
            style="width: 100%"
          />
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
      <BasicTable ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id"
        :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"
        @update:sorter="handleUpdateSorter">
        <template #tableTitle>
          <n-button type="primary" @click="handleExport" class="min-left-space hidden md:inline-flex"
            v-if="hasPermission(['/carOrder/export'])">
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" @reloadTable="reloadTable" />
    <Status ref="statusRef" @reloadTable="reloadTable" />
    <Dispatch ref="dispatchRef" @reloadTable="reloadTable" />
    <Cancel ref="cancelRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { NButton, useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { List, ConfirmAgree, ExportOrder } from '@/api/carOrder';
import { columns, schemas, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import View from './view.vue';
import Status from "@/views/carOrder/order_confirm.vue";
import Dispatch from "@/views/carOrder/order_dispatch.vue";
import Cancel from "@/views/carOrder/cancel.vue";
import { PrinterCarOrder } from '@/api/printer';
import { ExportOutlined } from "@vicons/antd";
import { usePermission } from "@/hooks/web/usePermission";
import { useSorter } from "@/hooks/common";

const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const tabValue = ref('WAIT_CONFIRM')
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();
const statusRef = ref();
const dispatchRef = ref();
const cancelRef = ref();
const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadTable);

const actionColumn = reactive({
  width: 120,
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
          auth: ['/carOrder/view']
        },
        // {
        //   label: '确认',
        //   onClick: handleAgree.bind(null, record),
        //   ifShow: () => {
        //     return record.orderStatus === "WAIT_CONFIRM";
        //   },
        //   auth: ['/carOrder/confirmAgree'],
        //   type: 'success'
        // },
        // {
        //   label: '拒绝',
        //   onClick: handleDisagree.bind(null, record),
        //   ifShow: () => {
        //     return record.orderStatus === "WAIT_CONFIRM";
        //   },
        //   auth: ['/carOrder/confirmDisagree'],
        //   type: 'error'
        // },
        // {
        //   label: '订单调度',
        //   onClick: handleDispatch.bind(null, record),
        //   ifShow: () => {
        //     return record.orderStatus === "WAIT_SERVE" && record.dispatchStatus === "WAIT";
        //   },
        //   auth: ['/carOrder/dispatch'],
        //   type: 'error'
        // },
        // {
        //   label: '退款',
        //   onClick: handleRefund.bind(null, record),
        //   ifShow: () => {
        //     return record.payStatus === "HAVE_PAID" && record.orderStatus !== "DONE" && record.orderStatus !== "SERVING" && record.adminCancelNum == 0;
        //   },
        //   auth: ['/carOrder/refund'],
        //   type: 'error'
        // },
        // {
        //   label: '打印',
        //   onClick: handlePrinterOrder.bind(null, record),
        //   ifShow: () => {
        //     return record.payTime != null && record.payStatus === "HAVE_PAID";
        //   },
        //   auth: ['/printer/printCarOrder'],
        //   type: 'info'
        // },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  res.orderStatus = tabValue.value;

  if (sortStatesRef.value.length > 0) {
    var bookSort = "";
    var createSort = "";
    sortStatesRef.value.forEach(element => {
      if (element.columnKey == 'bookStartTime') {
        bookSort = element.order
      } else if (element.columnKey == 'createdAt') {
        createSort = element.order
      }

    });
    return await List({ ...searchFormRef.value?.formModel, ...res, bookSort: bookSort, createSort: createSort });
  } else {
    return await List({ ...searchFormRef.value?.formModel, ...res });
  }
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleUpdateValue(e) {
  tabValue.value = e;
  reloadTable()
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}


// 订单确认
function handleAgree(record: Recordable) {

  dialog.info({
    title: '提示',
    content: '你确定要确认该订单吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      ConfirmAgree({ id: record.id }).then((_res) => {
        message.success('操作成功');
        reloadTable();
      });
    },
    onNegativeClick: () => { },
  });

}

// 导出
function handleExport() {
  let messageReactive = message.loading('正在导出列表...', { duration: 0 });
  ExportOrder({ ...searchFormRef.value?.formModel, ...{ orderStatus: tabValue.value } }).then((_res) => {
    messageReactive.destroy()
    message.success('请稍后在导出记录中进行下载');
  });
}

// 修改状态
function handleDisagree(record: Recordable) {
  statusRef.value.openModal(record.id);
}

// 调度
function handleDispatch(record: Recordable) {
  dispatchRef.value.openModal(record.id);
}

/**
 * 退款
 * @param record
 */
function handleRefund(record: Recordable) {
  cancelRef.value.openModal(record);
}

// 订单确认
function handlePrinterOrder(record: Recordable) {

  dialog.info({
    title: '提示',
    content: '你确定要打印该订单吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      PrinterCarOrder({ orderId: record.id }).then((_res) => {
        message.success('操作成功');
        reloadTable();
      });
    },
    onNegativeClick: () => { },
  });

}


onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped>
.statusTab {
  display: flex;
  margin-bottom: 30px;
  max-width: 100%;
  overflow-x: auto;

  div {
    padding: 0 16px;
    height: 28px;
    line-height: 28px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    text-wrap: nowrap;
    white-space: nowrap;

    span {
      cursor: pointer;
    }

    &.active {
      background: #F2F3F8;
      border-radius: 28px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
