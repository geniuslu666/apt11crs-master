<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }" :content-style="{
                    paddingBottom: '0px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">财务报表</text>
        </template>
        <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" style="margin-top: 15px">
          <template #statusSlot="{ model, field }">
            <n-input v-model:value="model[field]" />
          </template>
          <template #timeRangeSlot="{ model, field }">
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
      </n-card>
      <n-card :bordered="false" :content-style="{
                    padding: '0 20px',
                  }">
        <div class="statistic-title">
          <span class="statistic-title-t">数据统计</span>
          <span class="statistic-title-c" v-if="!searchFormValue.payChannel && !searchFormValue.payType && (!searchFormValue.timeRange || searchFormValue.timeRange.length <= 0)">/ 全部</span>
          <span class="statistic-title-c" v-else>/ {{ searchFormValue.payChannel ? getOptionLabel(options.pay_channel, searchFormValue.payChannel)+' ' : '' }}{{ searchFormValue.payType ? getOptionLabel(options.pay_type, searchFormValue.payType)+' ' : '' }}{{ searchFormValue.timeRange.length > 0 ? format(new Date(searchFormValue.timeRange[0]), 'yyyy-MM-dd HH:mm:ss')+'~'+format(new Date(searchFormValue.timeRange[1]), 'yyyy-MM-dd HH:mm:ss') : '' }}</span>
        </div>
        <n-row>
          <n-col :span="8">
            <div class="statistic-item">
              <div>总交易金额</div>
              <div>
                <n-number-animation :from="0.0" :to="statInfo.TransactionAmount" :precision="2" />
              </div>
            </div>
          </n-col>
          <n-col :span="8">
            <div class="statistic-item">
              <div>总支付金额</div>
              <div>
                <n-number-animation :from="0.0" :to="statInfo.PayAmount" :precision="2" />
              </div>
            </div>
          </n-col>
          <n-col :span="8">
            <div class="statistic-item">
              <div>总退款金额</div>
              <div>
                <n-number-animation :from="0.0" :to="statInfo.RefundAmount" :precision="2" />
              </div>
            </div>
          </n-col>
        </n-row>
        <div style="width: 100%;height: 1px;background: #F1F1F6;margin-top: 20px"></div>
      </n-card>
      <n-card :bordered="false" :content-style="{
                    padding: '20px',
                  }">
        <n-button type="primary" @click="handleExport('pay')">
          <template #icon>
            <n-icon>
              <CloudDownloadOutline />
            </n-icon>
          </template>
          支付报表下载
        </n-button>
        <n-button type="success" @click="handleExport('refund')" class="min-left-space">
          <template #icon>
            <n-icon>
              <CloudDownloadOutline />
            </n-icon>
          </template>
          退款报表下载
        </n-button>
        <div class="statusTab">
          <div :class="tabValue=='PAY' ? 'active' : ''"><span @click="handleUpdateValue('PAY')">交易记录</span></div>
          <div :class="tabValue=='REFUND' ? 'active' : ''"><span @click="handleUpdateValue('REFUND')">退款记录</span></div>
        </div>
        <template v-if="tabValue=='PAY'">
          <BasicTable  ref="actionPayRef" :columns="paycolumns" :request="loadDataTable" :scroll-x="scrollPayX" :resizeHeightOffset="-10000">
          </BasicTable>
        </template>
        <template v-else-if="tabValue=='REFUND'">
          <BasicTable  ref="actionRefundRef" :columns="refundcolumns" :request="loadDataTable" :scroll-x="scrollRefundX" :resizeHeightOffset="-10000">
          </BasicTable>
        </template>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {BasicTable} from "@/components/Table";
import {BasicForm, useForm} from "@/components/Form";
import {schemas, loadOptions, options} from "@/views/finance/report_model";
import {computed, onMounted, ref} from "vue";
import {ExportOutlined} from "@vicons/antd";
import {CloudDownloadOutline} from "@vicons/ionicons5"
import {adaTableScrollX, getOptionLabel} from "@/utils/hotgo";
import { defRangeShortcuts } from '@/utils/dateUtil';
import {List,Export} from "@/api/finance";
import {useMessage} from "naive-ui";
import {format} from "date-fns";

const tabValue = ref('PAY')
const message = useMessage();
const loading = ref(false);
const tabsType = ref('pay');
const statInfo = ref({
  TransactionAmount: 0.00,
  PayAmount: 0.00,
  RefundAmount: 0.00,
})
const actionPayRef = ref();
const actionRefundRef = ref();
const searchFormRef = ref<any>({});
const searchFormValue = ref({
  payChannel: '',
  payType: '',
  timeRange: []
})
// 表格列
const paycolumns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 100,
  },
  {
    title: '场景',
    key: 'scene',
    align: 'left',
    width: 100,
    render(row){
      if(row.Scene == 'HOTEL'){
        return '住宿'
      }else if(row.Scene == 'CAR'){
        return '接送机'
      }else if(row.Scene == 'FOOD'){
        return '餐厅'
      }else if(row.Scene == 'SPA'){
        return '按摩'
      }else if(row.Scene == 'CABINET'){
        return '储物柜'
      }else{
        return '未知'
      }
    }
  },
  {
    title: '支付单号',
    key: 'TypeSn',
    align: 'left',
    width: 150,
  },
  {
    title: '支付平台',
    key: 'payChannel',
    align: 'left',
    width: 100,
    render(row){
      return getOptionLabel(options.value.pay_channel, row.payChannel)
    }
  },
  {
    title: '支付方式',
    key: 'payType',
    align: 'left',
    width: 100,
    render(row){
      return getOptionLabel(options.value.pay_type, row.payType)
    }
  },
  {
    title: '支付金额',
    key: 'amount',
    align: 'left',
    width: 150,
    render(row){
      return row.amount + 'JPY'
    }
  },
  {
    title: '支付时间',
    key: 'payTime',
    align: 'left',
    width: 150,
  },
];

// 表格列
const refundcolumns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 90,
  },
  {
    title: '场景',
    key: 'scene',
    align: 'left',
    width: 90,
    render(row){
      if(row.Scene == 'HOTEL'){
        return '住宿'
      }else if(row.Scene == 'CAR'){
        return '接送机'
      }else if(row.Scene == 'FOOD'){
        return '餐厅'
      }else if(row.Scene == 'SPA'){
        return '按摩'
      }else if(row.Scene == 'CABINET'){
        return '储物柜'
      }else{
        return '未知'
      }
    }
  },
  {
    title: '退款单号',
    key: 'TypeSn',
    align: 'left',
    width: 150,
  },
  {
    title: '支付平台',
    key: 'payChannel',
    align: 'left',
    width: 90,
    render(row){
      return getOptionLabel(options.value.pay_channel, row.payChannel)
    }
  },
  {
    title: '支付方式',
    key: 'payType',
    align: 'left',
    width: 90,
    render(row){
      return getOptionLabel(options.value.pay_type, row.payType)
    }
  },
  {
    title: '退款金额',
    key: 'amount',
    align: 'left',
    width: 90,
    render(row){
      return row.amount + 'JPY'
    }
  },
  {
    title: '退款时间',
    key: 'payTime',
    align: 'left',
    width: 150,
  },
];
const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});
const scrollPayX = computed(() => {
  return adaTableScrollX(paycolumns, 0);
});
const scrollRefundX = computed(() => {
  return adaTableScrollX(refundcolumns, 0);
});

// 加载表格数据
const loadDataTable = async (res) => {
  searchFormValue.value = { ...searchFormValue.value, ...searchFormRef.value?.formModel }
  console.log('searchFormRef',searchFormRef.value?.formModel)
  console.log('searchFormValue',searchFormValue.value)
  loading.value = true;
  console.log('tabsType',tabsType.value)
  res.type = tabsType.value
  let data = await List({ ...searchFormRef.value?.formModel,...res });
  statInfo.value.PayAmount = data.payAmount
  statInfo.value.RefundAmount = data.refundAmount
  statInfo.value.TransactionAmount = data.transactionAmount
  loading.value = false;
  if(tabsType.value == 'pay'){
    return data.pay_list_item;
  }else{
    return data.refund_list_item;
  }
};

function handleUpdateValue(e){
  tabValue.value = e;
  if(e == 'PAY'){
    tabsType.value = 'pay'
  }else{
    tabsType.value = 'refund'
  }
}

// 重新加载表格数据
function reloadTable() {
  if(tabsType.value == 'pay'){
    actionPayRef.value?.reload();
  }else{
    actionRefundRef.value?.reload();
  }
}

function handleChange(e){
  tabsType.value = e
}

// 导出
function handleExport(type) {
  message.loading('正在导出...', { duration: 1200 });
  var querryParams = {
    type: type,
  }

  if(searchFormRef.value?.formModel.payChannel){
    querryParams.payChannel = searchFormRef.value?.formModel.payChannel
  }

  if(searchFormRef.value?.formModel.payType){
    querryParams.payType = searchFormRef.value?.formModel.payType
  }

  if(searchFormRef.value?.formModel.page){
    querryParams.page = searchFormRef.value?.formModel.page
  }

  if(searchFormRef.value?.formModel.pageSize){
    querryParams.pageSize = searchFormRef.value?.formModel.pageSize
  }

  if (searchFormRef.value?.formModel.timeRange && searchFormRef.value?.formModel.timeRange.length > 0) {
    querryParams.startTime = searchFormRef.value?.formModel.timeRange[0];
    querryParams.endTime = searchFormRef.value?.formModel.timeRange[1];
  }

  Export(querryParams);
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped>
.statistic-title{
  margin-bottom: 10px;
  .statistic-title-t{
    font-weight: 500;
    font-size: 18px;
    color: #3D3D3D;
    line-height: 25px;
    margin-right: 7px;
  }
  .statistic-title-c{
    font-weight: 400;
    font-size: 14px;
    color: #979797;
    line-height: 20px;
  }
}
.statistic-item{
  padding: 12px 0;
  div{
    &:first-child{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      margin-bottom: 5px;
    }
    &:last-child{
      font-weight: 700;
      font-size: 24px;
      color: #3D3D3D;
      line-height: 42px;
    }
  }
}

.statusTab{
  display: flex;
  margin-top: 30px;
  margin-bottom: 4px;
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


