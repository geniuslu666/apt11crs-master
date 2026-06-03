<template>
  <div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '20px',
                  }">
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane name="WAIT_CONFIRM" tab="待确认">
        </n-tab-pane>
        <n-tab-pane name="CONFIRMED" tab="已确认">
        </n-tab-pane>
        <n-tab-pane name="CANCEL" tab="已取消">
        </n-tab-pane>
        <n-tab-pane name="ALL" tab="全部">
        </n-tab-pane>
      </n-tabs>
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
        <template #bookDateTimeSlot="{ model, field }">
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
      <BasicTable  ref="actionRef"  :columns="toretaColumns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"
        @update:sorter="handleUpdateSorter">
        <template #tableTitle>
          <n-button type="primary" @click="handleExport" class="min-left-space"
                    v-if="hasPermission(['/carOrder/export'])">
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
import { List, ExportOrder } from '@/api/foodOrder';
import { toretaColumns, crsAllSchemas, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import View from './view.vue';
import {ExportOutlined} from "@vicons/antd";
import {usePermission} from "@/hooks/web/usePermission";
import {useSorter} from "@/hooks/common";

const dialog = useDialog();
const message = useMessage();
const {hasPermission} = usePermission();
const tabValue = ref('WAIT_CONFIRM')
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();
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
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(toretaColumns, actionColumn.width);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:5 xl:5 2xl:5' },
  labelWidth: 80,
  schemas: crsAllSchemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  res.orderType = 'CRSALL';
  if(searchFormRef.value?.formModel.bookingStatus){
    tabValue.value = searchFormRef.value?.formModel.bookingStatus
  }

  if(sortStatesRef.value.length > 0){
    var bookSort = "";
    var createSort = "";
    sortStatesRef.value.forEach(element => {
      if(element.columnKey == 'bookDateTime'){
        bookSort = element.order
      }else if(element.columnKey == 'createdAt'){
        createSort = element.order
      }

    });
    return await List({ ...searchFormRef.value?.formModel, ...res, bookSort: bookSort, createSort: createSort });
  }else{
    return await List({ ...searchFormRef.value?.formModel, ...res });
  }
};

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleUpdateValue(e){
  tabValue.value = e;
  searchFormRef.value?.setFieldsValue({
    bookingStatus: e
  });
  if(e == 'WAIT_CONFIRM'){
    searchFormRef.value?.setFieldsValue({
      orderStatus: 'HAVE_PAID'
    });
  }
  reloadTable()
}

// 导出
function handleExport() {
  let messageReactive = message.loading('正在导出列表...', {duration: 0});
  ExportOrder({ ...searchFormRef.value?.formModel }).then((_res) => {
    messageReactive.destroy()
    message.success('请稍后在导出记录中进行下载');
  });
}

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}


onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

