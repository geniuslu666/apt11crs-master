<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="">
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">

      </BasicTable>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import {carOrderLogList} from '@/api/orderLog';
import { columns, schemas, options, loadOptions } from './order_model';
import { adaTableScrollX, getOptionLabel } from '@/utils/hotgo';
import {useRouter} from "vue-router";

const dialog = useDialog();
const router = useRouter();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const checkedIds = ref([]);

const actionColumn = reactive({
  /*width: 360,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [

      ],
    });
  },*/
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, 0);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

// 加载表格数据
const loadDataTable = async (res) => {
  return await carOrderLogList({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}


onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

