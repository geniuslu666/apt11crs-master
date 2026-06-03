<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">结算概况</text>
        </template>
        <n-row>
          <n-col :span="6">
            <n-statistic label="订单总额">
              {{ data.totalOrderAmount }}
            </n-statistic>
          </n-col>
          <n-col :span="6">
            <n-statistic label="待结算">
              {{ data.waitSettlementAmount }}
            </n-statistic>
          </n-col>
          <n-col :span="6">
            <n-statistic label="已结算">
              {{ data.doneSettlementAmount }}
            </n-statistic>
          </n-col>
          <n-col :span="6">
            <n-statistic label="已核账">
              {{ data.hadVerifySettlementAmount }}
            </n-statistic>
          </n-col>
        </n-row>
      </n-card>
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">结算明细</text>
        </template>
        <div class="statusTab">
          <div :class="tabValue=='TECHNICIAN' ? 'active' : ''"><span @click="handleUpdateValue('TECHNICIAN')">技师</span></div>
          <div :class="tabValue=='ISP' ? 'active' : ''"><span @click="handleUpdateValue('ISP')">服务商</span></div>
        </div>
        <template v-if="tabValue=='TECHNICIAN'">
          <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
            <template #verifyTimeSlot="{ model, field }">
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
          <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000">
          </BasicTable>
        </template>
        <template v-if="tabValue=='ISP'">
          <BasicForm  ref="searchIspFormRef" @register="registerIsp" @submit="reloadIspTable" @reset="reloadIspTable" @keyup.enter="reloadIspTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
            <template #verifyTimeSlot="{ model, field }">
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
          <BasicTable  ref="actionIspRef" :columns="ispColumns" :request="loadIspDataTable" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000">
          </BasicTable>
        </template>
      </n-card>
    </n-spin>
    <Verify ref="verifyRef" @reloadTable="reloadTable" @reloadIspTable="reloadIspTable"/>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, h, reactive} from 'vue';
import { useRouter } from 'vue-router';
import {NIcon, NTag, NTooltip, useMessage} from 'naive-ui';
import { Stat,List } from '@/api/spaSettlementOrder';
import {BasicTable, TableAction} from "@/components/Table";
import {adaTableScrollX, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import { defRangeShortcuts } from '@/utils/dateUtil';
import {BasicForm, useForm} from "@/components/Form";
import {ispSchemas, loadTechnicianOptions, schemas} from "@/views/spaSettlementOrder/model";
import Verify from "@/views/spaSettlementOrder/verify.vue";

const tabValue = ref('TECHNICIAN')
const message = useMessage();
const router = useRouter();
const params = router.currentRoute.value.params;
const loading = ref(false);
const verifyRef = ref();
const data = ref({
  totalOrderAmount: 0,
  waitSettlementAmount: 0,
  doneSettlementAmount: 0,
  hadVerifySettlementAmount: 0,
});
const actionRef = ref();
const searchFormRef = ref<any>({});
const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

const actionIspRef = ref();
const searchIspFormRef = ref<any>({});
const [registerIsp, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: ispSchemas,
});

const actionColumn = reactive({
  width: 150,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleOrderView.bind(null, record),
          auth: ['/spaOrder/technicianOrderList'],
        },
        {
          label: '核账',
          onClick: handleVerify.bind(null, record),
          ifShow: () => {
            return record.verifyStatus === "WAIT_VERIFY";
          },
          auth: ['/spaSettlementOrder/verify'],
          type: 'success'
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const getInfo = () => {
  loading.value = true;
  Stat(params)
    .then((res) => {
      data.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
};

// 表格列
const columns = [
  {
    title: '结算单号',
    key: 'orderSn',
    align: 'left',
    width: 180,
  },
  {
    title: '技师',
    key: 'technician',
    align: 'left',
    width: 110,
    render(row) {
      return row.technicianDetail.name
    },
  },
  {
    title: '订单总额',
    key: 'orderAmount',
    align: 'left',
    width: 110,
  },
  {
    title: '结算金额',
    key: 'settlementAmount',
    align: 'left',
    width: 110,
  },
  {
    title: '结算状态',
    key: 'status',
    align: 'left',
    width: 110,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: "success",
          bordered: false,
        },
        {
          default: () => "已结算",
        }
      );
    },
  },
  {
    title: '账期时间',
    key: 'startTime',
    align: 'left',
    width: 230,
    render(row) {
      return h(
        'div',
        {
          style: {
          },
        },
        [
          h(
            'div',
            {},
            {
              default: () => "开始时间：" + row.startTime
            }
          ),
          h(
            'div',
            {},
            {
              default: () => "结束时间：" + row.endTime
            }
          ),
        ],
      );
    }
  },
  {
    title: '核账状态',
    key: 'verifyStatus',
    align: 'left',
    width: 110,
    render(row) {
      if(row.verifyStatus == "WAIT_VERIFY"){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "warning",
            bordered: false,
          },
          {
            default: () => "待核账",
          }
        );
      }
      if(row.verifyStatus == "VERIFIED"){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "success",
            bordered: false,
          },
          {
            default: () => "已核账",
          }
        );
      }
      return "--"
    },
  },
  {
    title: '核账时间',
    key: 'verifyTime',
    align: 'left',
    width: 200,
    render(row) {
      if (row.verifyTime != null){
        return h(
          'div',
          {
            style: {
            },
          },
          [
            h(
              'div',
              {},
              {
                default: () => row.verifyTime
              }
            ),
            h(
              'div',
              {},
              {
                default: () => "操作员：" + row.operateDetail.username
              }
            ),
          ],
        );
      }
      return "--"
    }
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: 180,
  },
];

// 表格列
const ispColumns = [
  {
    title: '结算单号',
    key: 'orderSn',
    align: 'left',
    width: 180,
  },
  {
    title: '服务商',
    key: 'isp',
    align: 'left',
    width: 110,
    render(row) {
      return row.ispDetail.name
    },
  },
  {
    title: '订单总额',
    key: 'orderAmount',
    align: 'left',
    width: 110,
  },
  {
    title: '结算金额',
    key: 'settlementAmount',
    align: 'left',
    width: 110,
  },
  {
    title: '结算状态',
    key: 'status',
    align: 'left',
    width: 110,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: "success",
          bordered: false,
        },
        {
          default: () => "已结算",
        }
      );
    },
  },
  {
    title: '账期时间',
    key: 'startTime',
    align: 'left',
    width: 230,
    render(row) {
      return h(
        'div',
        {
          style: {
          },
        },
        [
          h(
            'div',
            {},
            {
              default: () => "开始时间：" + row.startTime
            }
          ),
          h(
            'div',
            {},
            {
              default: () => "结束时间：" + row.endTime
            }
          ),
        ],
      );
    }
  },
  {
    title: '核账状态',
    key: 'verifyStatus',
    align: 'left',
    width: 110,
    render(row) {
      if(row.verifyStatus == "WAIT_VERIFY"){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "warning",
            bordered: false,
          },
          {
            default: () => "待核账",
          }
        );
      }
      if(row.verifyStatus == "VERIFIED"){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "success",
            bordered: false,
          },
          {
            default: () => "已核账",
          }
        );
      }
      return "--"
    },
  },
  {
    title: '核账时间',
    key: 'verifyTime',
    align: 'left',
    width: 200,
    render(row) {
      if (row.verifyTime != null){
        return h(
          'div',
          {
            style: {
            },
          },
          [
            h(
              'div',
              {},
              {
                default: () => row.verifyTime
              }
            ),
            h(
              'div',
              {},
              {
                default: () => "操作员：" + row.operateDetail.username
              }
            ),
          ],
        );
      }
      return "--"
    }
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: 180,
  },
];

// 加载表格数据
const loadDataTable = async (res) => {
  res.settlementObject = "TECHNICIAN"
  return await List({ ...{

    },...searchFormRef.value?.formModel, ...res });
};

// 加载表格数据
const loadIspDataTable = async (res) => {
  res.settlementObject = "ISP"
  return await List({ ...searchIspFormRef.value?.formModel, ...res });
};

// const scrollX = computed(() => {
//   return adaTableScrollX(columns, 0);
// });

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 重新加载表格数据
function reloadIspTable() {
  actionIspRef.value?.reload();
}

// 结算订单列表
function handleOrderView(record: Recordable) {
  // editRef.value.openModal(record);
  router.push({ name: 'spaSettlementOrderDetail', params: { id: record.id } });
}

// 修改状态
function handleVerify(record: Recordable) {
  verifyRef.value.openModal(record);
}

function handleUpdateValue(e){
  tabValue.value = e;
}

onMounted(() => {
  getInfo();
  loadTechnicianOptions();
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}

.statusTab{
  display: flex;
  margin-bottom: 30px;
  div{
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    span{
      cursor: pointer;
    }
    &.active{
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>


