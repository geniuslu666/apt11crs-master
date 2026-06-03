<template>
  <div>
    <n-spin :show="loading" description="请稍候...">

      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="统计报表"
      >

        <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable" style="margin-top: 15px">
          <template #statusSlot="{ model, field }">
            <n-input v-model:value="model[field]" />
          </template>
          <template #searchTimeSlot="{ model, field }">
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

        <div style="padding: 20px 0px;background: #f2f3f5;">
        <n-flex justify="space-around">
          <n-statistic label="预计收入(JPY)">
           {{data.overviewStat.expectIncome}}
          </n-statistic>
          <n-statistic label="收入总额(JPY)">
            {{data.overviewStat.totalIncome}}
          </n-statistic>
          <n-statistic label="支出总额(JPY)">
            {{data.overviewStat.totalExpend}}
          </n-statistic>
        </n-flex>
        </div>
      </n-card>
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="收入概况"
      >
        <n-row>
          <n-col :span="24">
            <div class="survey">
              <div class="layui-card layui-input-inline">
                <div class="layui-card-header">
                  <span class="hint">收入总额</span>
                </div>
                <div class="layui-card-body">
                  <span class="num total_income">{{data.overviewStat.totalIncome}}</span>
                </div>
              </div>
              <div class="symbol"> = </div>
              <div class="stat-text" >民宿订单</div>
              <div class="symbol"> + </div>
              <div class="stat-text" >餐饮订单</div>
              <div class="symbol"> + </div>
              <div class="stat-text" >车队订单</div>
              <div class="symbol"> + </div>
              <div class="stat-text" >按摩订单</div>
            </div>
          </n-col>
        </n-row>

        <n-grid x-gap="12" :cols="4">
          <n-gi>
            <n-card title="民宿订单" :header-style="cardHeadStyle" :content-style="cardContentStyle">
              <template #header-extra>
                <n-button
                  text
                  tag="a"
                  @click="handleHomeStayOrder"
                  target="_blank"
                  type="primary"
                >
                  明细
                </n-button>
              </template>
              <n-statistic>
                {{data.incomeStat.HomeStayOrderIncome}}
              </n-statistic>
            </n-card>
          </n-gi>
          <n-gi>
            <n-card title="餐饮订单" :header-style="cardHeadStyle" :content-style="cardContentStyle">
              <template #header-extra>
                <n-button
                  text
                  tag="a"
                 @click="handleFoodOrder"
                  type="primary"
                >
                  明细
                </n-button>
              </template>
              <n-statistic>
                {{data.incomeStat.foodOrderIncome}}
              </n-statistic>
            </n-card>
          </n-gi>
          <n-gi>
            <n-card title="车队订单" :header-style="cardHeadStyle" :content-style="cardContentStyle">
              <template #header-extra>
                <n-button
                  text
                  tag="a"
                  @click="handleCarOrder"
                  type="primary"
                >
                  明细
                </n-button>
              </template>
              <n-statistic>
                {{data.incomeStat.carOrderIncome}}
              </n-statistic>
            </n-card>
          </n-gi>
          <n-gi>
            <n-card title="按摩订单" :header-style="cardHeadStyle" :content-style="cardContentStyle">
              <template #header-extra>
                <n-button
                  text
                  tag="a"
                  @click="handleMassageOrder"
                  type="primary"
                >
                  明细
                </n-button>
              </template>
              <n-statistic>
                {{data.incomeStat.massageOrderIncome}}
              </n-statistic>
            </n-card>
          </n-gi>
        </n-grid>

      </n-card>
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="支出概况"
      >
        <n-row>
          <n-col :span="24">
            <div class="survey">
              <div class="layui-card layui-input-inline">
                <div class="layui-card-header">
                  <span class="hint">支出总额</span>
                </div>
                <div class="layui-card-body">
                  <span class="num total_income">{{data.overviewStat.totalExpend}}</span>
                </div>
              </div>
              <div class="symbol"> = </div>
              <div class="stat-text" >订单退款</div>
              <div class="symbol"> + </div>
              <div class="stat-text" >员工提现</div>
              <div class="symbol"> + </div>
              <div class="stat-text" >渠道提现</div>
            </div>
          </n-col>
        </n-row>

        <n-grid x-gap="12" :cols="4">
          <n-gi>
            <n-card title="订单退款" :header-style="cardHeadStyle" :content-style="cardContentStyle">
              <template #header-extra>
                <n-button
                  text
                  tag="a"
                  type="primary"
                >
                  明细
                </n-button>
              </template>
              <n-statistic>
                {{data.expendStat.orderRefund}}
              </n-statistic>
            </n-card>
          </n-gi>
          <n-gi>
            <n-card title="员工提现" :header-style="cardHeadStyle" :content-style="cardContentStyle">
              <template #header-extra>
                <n-button
                  text
                  tag="a"
                  @click="handleStaffWithdraw"
                  type="primary"
                >
                  明细
                </n-button>
              </template>
              <n-statistic>
                {{data.expendStat.staffWithdraw}}
              </n-statistic>
            </n-card>
          </n-gi>
          <n-gi>
            <n-card title="渠道提现" :header-style="cardHeadStyle" :content-style="cardContentStyle">
              <template #header-extra>
                <n-button
                  text
                  tag="a"
                  @click="handleChannelWithdraw"
                  target="_blank"
                  type="primary"
                >
                  明细
                </n-button>
              </template>
              <n-statistic>
                {{data.expendStat.channelWithdraw}}
              </n-statistic>
            </n-card>
          </n-gi>
        </n-grid>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, h} from 'vue';
import { useRouter } from 'vue-router';
import {NIcon, NTooltip, useMessage} from 'naive-ui';
import { Stat } from '@/api/finance';
import {BasicForm, useForm} from "@/components/Form";
import { defRangeShortcuts } from '@/utils/dateUtil';
import {schemas} from "@/views/finance/model";

const message = useMessage();
const router = useRouter();
const params = router.currentRoute.value.params;
const loading = ref(false);
const data = ref({
  overviewStat:{},
  incomeStat:{},
  expendStat:{}
});
const searchFormRef = ref<any>({});
const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

const cardHeadStyle = ref({
  padding: '16px 24px 0',
})

const cardContentStyle = ref({
  padding: '0 24px 16px',
})

const getInfo = (reqParams) => {
  loading.value = true;
  Stat(reqParams)
    .then((res) => {
      console.log('res',res)
      data.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
};


// 重新加载表格数据
function reloadTable() {
  // console.log('searchFormRef', { ...params,...searchFormRef.value?.formModel })
  getInfo({ ...params, ...searchFormRef.value?.formModel })
}

// 民宿订单跳转
function handleHomeStayOrder(){
  router.push({ name: 'pmsAppReservationIndex', params: {} });
}

// 餐饮订单跳转
function handleFoodOrder(){

}

// 车队订单跳转
function handleCarOrder(){

}

// 按摩订单跳转
function handleMassageOrder(){

}

// 员工提现跳转
function handleStaffWithdraw(){
  router.push({ name: 'staffApplyIndex', params: {} });
}

// 渠道提现跳转
function handleChannelWithdraw(){
  router.push({ name: 'channelWithdrawIndex', params: {} });
}

onMounted(() => {
  getInfo(params);
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}

.survey {
  margin: 20px 0px;
  background: #f2f3f5;
  display: flex;
  .layui-card{
    //width: 200px !important;
    background: #f2f3f5;
    vertical-align: top;
    display: inline-block;
    border-radius: 0;
    float: left;
    margin-top: 10px;
    box-shadow: unset;
    .layui-card-header {
      line-height: 42px;
      font-size: 14px;
      box-shadow: unset;
      border: 0;
      span{
        font-size: 16px;
      }
    }
  }
  .layui-input-inline {
    padding-left: 40px;
    box-sizing: border-box;
  }

  .layui-card-body {
    font-size: 24px;
    padding-top: 0px;
    cursor: pointer;
  }
  .symbol {
    padding: 35px 20px;
    font-size: 18px;
  }
  .stat-text {
    padding: 35px 10px;
    font-size: 14px;
  }
}

</style>


