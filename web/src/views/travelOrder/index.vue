<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{ padding: '20px' }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">一日游订单</text>
        </template>
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{ padding: '0 20px 20px' }">
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane name="" tab="全部" />
        <n-tab-pane name="WAIT_PAY" tab="待支付" />
        <n-tab-pane name="WAIT_VERIFY" tab="待核销" />
        <n-tab-pane name="DONE" tab="已完成" />
<!--        <n-tab-pane name="CANCEL" tab="已取消" />-->
        <n-tab-pane name="REFUND" tab="已退款" />
      </n-tabs>
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" >
        <template #bookDateSlot="{ model, field }">
          <n-date-picker
            v-model:formatted-value="model[field]"
            type="daterange"
            value-format="yyyy-MM-dd"
            clearable
            :shortcuts="defRangeShortcuts()"
            style="width: 100%"
          />
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
      <BasicTable
        ref="actionRef"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :actionColumn="actionColumn"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      />
    </n-card>
    <OrderView ref="viewRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref, computed, onMounted} from 'vue';
import {NTag, NTooltip} from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { List } from '@/api/travelOrder';
import { schemas, options, loadOptions } from './model';
import {adaTableScrollX, getOptionLabel, getOptionTag} from '@/utils/hotgo';
import OrderView from './view.vue';
import {defRangeShortcuts} from "@/utils/dateUtil";
import {isNullObject} from "@/utils/is";

const tabValue = ref('');
const actionRef = ref();
const searchFormRef = ref<any>({});
const viewRef = ref();

const columns = [
  { title: '预约单号', key: 'orderSn', align: 'left', width: 180 },
  {
    title: '产品',
    key: 'productName',
    align: 'left',
    width: 200,
    render(row){
      return h(
        'div',
        {
          class:'flex-item',
          style: {
            // paddingLeft: '8px',
          }
        },
        [
          h(
            'div',
            {
              style: {
                fontWeight: '400',
                fontSize: '14px',
                // color: '#1664FF',
                lineHeight: '20px'
              }
            },
            {
              default: () => row.productInfo.title,
            }
          ),
          h(
            'div',
            {
              style: {
                fontWeight: '400',
                fontSize: '14px',
                color: 'grey', //#3D3D3D
                lineHeight: '20px'
              }
            },
            {
              default: () => row.skuInfo.name,
            }
          ),
        ]
      )
    }
  },
  {
    title: '预约信息',
    key: 'orderSn',
    align: 'left',
    width: 150,
    render(row){
      let memberHtmlStr = ''
      if(row.memberDeleted){
        memberHtmlStr = h(
          'div',
          {
            style:{
              color: 'red'
            }
          },
          {
            default: () => '用户已注销'
          }
        )
      }
      let htmlStr = h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.bookDate ? row.bookDate.slice(0, 10) : '',
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.bookingName + ' / ' + row.pmsMemberMemberNo
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.phoneArea+row.bookingMobile,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.bookingNum + '人',
            }
          ),
          memberHtmlStr
        ]
      )

      return h(
        NTooltip,
        null,
        {
          trigger:()=>
            htmlStr,
          default: () => htmlStr,
        },
      )
    }
  },
  {
    title: '订单金额',
    key: 'orderAmount',
    align: 'left',
    width: 120,
    render(row) { return row.orderAmount + ' JPY'; },
  },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: 90,
    render(row) {
      if (isNullObject(row.orderStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.travel_order_status, row.orderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.travel_order_status, row.orderStatus),
        }
      );
    },
  },
  {
    title: '核销信息',
    key: 'verify',
    align: 'left',
    width: 200,
    render(row){
      if (row.orderStatus == 'DONE'){
        let htmlStr = h(
          'div',
          null,
          [
            h(
              'div',
              null,
              {
                default: () => row.verifyStaffInfo.name,
              }
            ),
            h(
              'div',
              null,
              {
                default: () => row.verifyTime
              }
            ),
          ]
        )

        return h(
          NTooltip,
          null,
          {
            trigger:()=>
              htmlStr,
            default: () => htmlStr,
          },
        )
      }else{
        return '--';
      }

    }
  },
  { title: '下单时间', key: 'createdAt', align: 'left', width: 160 },
];

const actionColumn = reactive({
  width: 100,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        { label: '详情', onClick: handleView.bind(null, record) },
      ],
    });
  },
});

const scrollX = computed(() => adaTableScrollX(columns, actionColumn.width));

const [register] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

function handleUpdateValue(v) {
  tabValue.value = v;
  reloadTable();
}

const loadDataTable = async (res) => {
  if (tabValue.value) res.orderStatus = tabValue.value;
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

function reloadTable() {
  actionRef.value?.reload();
}

// 查看详情
function handleView(record: Recordable) {
  // console.log('查看详情', record);
  viewRef.value.openModal(record);
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>
