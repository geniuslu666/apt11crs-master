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
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable"
        @keyup.enter="reloadTable">
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
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { NButton, useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { List, ExportOrder } from '@/api/cabinetOrder';
import { columns, schemas, loadOptions } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import { defRangeShortcuts } from '@/utils/dateUtil';
import View from './view.vue';
import { ExportOutlined } from "@vicons/antd";
import { usePermission } from "@/hooks/web/usePermission";
import { useSorter } from "@/hooks/common";

const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();
const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadTable);

const actionColumn = reactive({
  width: 130,
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
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 90,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  // console.log("searchInfo", { ...searchFormRef.value?.formModel, ...res });
  if (sortStatesRef.value.length > 0) {
    var bookSort = "";
    var createSort = "";
    sortStatesRef.value.forEach(element => {
      if (element.columnKey == 'createdAt') {
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

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}

// 导出
function handleExport() {
  let messageReactive = message.loading('正在导出列表...', { duration: 0 });
  ExportOrder({ ...searchFormRef.value?.formModel }).then((_res) => {
    messageReactive.destroy()
    message.success('请稍后在导出记录中进行下载');
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

  div {
    padding: 0 16px;
    height: 28px;
    line-height: 28px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;

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
