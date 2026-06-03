<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{ padding: '20px' }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">核销记录</text>
        </template>
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{ padding: '0 20px 20px' }">
      <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" />
      <BasicTable
        ref="actionRef"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      />
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {computed, h, ref} from 'vue';
import { BasicTable } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { List } from '@/api/travelVerifyRecord';
import { schemas } from './model';
import { adaTableScrollX } from '@/utils/hotgo';
import {NTooltip} from "naive-ui";

const actionRef = ref();
const searchFormRef = ref<any>({});

const columns = [
  { title: '预约单号', key: 'orderSn', align: 'left', width: 180 },
  {
    title: '产品',
    key: 'productName',
    align: 'left',
    width: 200,
    render(row){
      return row.productInfo.title
    }
  },
  {
    title: '预约信息',
    key: 'orderSn',
    align: 'left',
    width: 180,
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
              default: () => row.orderInfo.bookDate,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.orderInfo.bookingName + ' / ' + row.pmsMemberMemberNo
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.orderInfo.phoneArea+row.orderInfo.bookingMobile,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.orderInfo.bookingNum + '人',
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
    title: '核销人员',
    key: 'verifyStaffName',
    align: 'left',
    width: 180,
    render(row){
      return row.verifyStaffInfo.name
    }
  },
  { title: '核销时间', key: 'verifyTime', align: 'left', width: 180 },
  { title: '创建时间', key: 'createdAt', align: 'left', width: 180 },
];

const scrollX = computed(() => adaTableScrollX(columns, 0));

const [register] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

const loadDataTable = async (res) => {
  return await List({ ...searchFormRef.value?.formModel, ...res });
};

function reloadTable() {
  actionRef.value?.reload();
}
</script>

<style lang="less" scoped></style>
