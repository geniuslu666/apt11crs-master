<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">导出记录</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
      <BasicTable  ref="actionRef" :columns="columns" :actionColumn="actionColumn" :request="loadDataTable" :scroll-x="scrollX" :resizeHeightOffset="-10000">
        <template #tableTitle>
        </template>
      </BasicTable>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed, onMounted, reactive, h} from 'vue';
import {BasicTable, TableAction} from '@/components/Table';
import { ExportList } from '@/api/cabinetOrder';
import { columns } from './model';
import { adaTableScrollX } from '@/utils/hotgo';

const actionRef = ref();
const searchFormRef = ref<any>({});

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
          label: '下载',
          onClick: handleDowload.bind(null, record),
          ifShow: () => {
            return record.status === 1;
          },
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

// 加载表格数据
const loadDataTable = async (res) => {
  return await ExportList({ ...searchFormRef.value?.formModel, ...res });
};

// 下载
function handleDowload(record: Recordable) {
  window.open(record.path)
}

onMounted(() => {
});
</script>

<style lang="less" scoped></style>

