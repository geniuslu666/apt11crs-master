<template>
  <div>
    <div class="n-layout-page-header">
      <n-spin :show="show" description="请稍候...">
        <n-card :bordered="false" :title="driverName+ ' - 司机结算'">
          <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
          <template #header-extra>
            账期时间：{{ settlementOrderInfo.startTime }} ~ {{ settlementOrderInfo.endTime }}
          </template>
          <n-row>
            <n-col :span="6">
              <n-statistic label="订单总额">
                {{ settlementOrderInfo.orderAmount }}
              </n-statistic>
            </n-col>
            <n-col :span="6">
              <n-statistic label="结算金额">
                {{ settlementOrderInfo.settlementAmount }}
              </n-statistic>
            </n-col>
          </n-row>
        </n-card>
      </n-spin>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>

        </template>
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
import { SettleOrderList } from '@/api/carOrder';
import { columns, schemas } from './order';
import { adaTableScrollX, getOptionLabel } from '@/utils/hotgo';
import {useRouter} from "vue-router";
import {View as SettlementOrderView} from "@/api/carSettlementOrder";

const show = ref(false);
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const router = useRouter();
const params = router.currentRoute.value.params;
const actionRef = ref();
const searchFormRef = ref<any>({});
const checkedIds = ref([]);
const driverName = ref('');
const settlementOrderInfo = ref({});

const actionColumn = reactive({
  // width: 288,
  // title: '操作',
  // key: 'action',
  // fixed: 'right',
  // render(record) {
  //   return h(TableAction as any, {
  //     style: 'button',
  //     actions: [
  //
  //     ],
  //   });
  // },
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
  return await SettleOrderList({ ...searchFormRef.value?.formModel, ...res, settlementOrderId: params.id });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

async function loadSettlementView(){
  let info = await SettlementOrderView({ id: params.id })
  driverName.value = info.driverDetail.name
  settlementOrderInfo.value = info
}

onMounted(async() => {
  show.value = true;
  await loadSettlementView()
  show.value = false;
});
</script>

<style lang="less" scoped></style>

