<template>
  <div>
    <n-grid :cols="1" y-gap="15">
      <!-- PC端 实时概况 -->
      <n-gi class="hidden md:block">
        <n-spin :show="show1" description="请稍候...">
          <n-card :bordered="false" :header-style="{
            padding: '25px 20px 23px',
          }" :content-style="{
            padding: '0 20px 23px',
          }">
            <template #header>
              <div class="room-stat-title">实时概况</div>
              <div class="room-stat-div">
                <span class="room-stat-time">在{{ nowTime }}刷新</span>
                <n-button text style="margin-left: 20px" @click="refreshBaseInfo">
                  <img style="width: 15px;" src="@/assets/images/pms_dashboard_refresh_icon.png" />
                  <span class="room-stat-refresh">刷新</span>
                </n-button>
              </div>
            </template>
            <n-grid :cols="4" x-gap="15">
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日预约数</div>
                      <div class="room-stat-gi-div-item-d2">总数：{{ baseInfo.todayOrderTotalNum }}</div>
                      <div class="room-stat-gi-div-item-d3">到店：{{ baseInfo.todayDdOrderTotalNum }} / 上门：{{
                        baseInfo.todaySmOrderTotalNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/spa_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">营业额</div>
                      <div class="room-stat-gi-div-item-d2">今日：{{ baseInfo.todayOrderAmount }}</div>
                      <div class="room-stat-gi-div-item-d3">累计：{{ baseInfo.totalOrderAmount }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/spa_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">累计预约数</div>
                      <div class="room-stat-gi-div-item-d2">有效：{{ baseInfo.effectiveTotalOrderTotalNum }}</div>
                      <div class="room-stat-gi-div-item-d3">全部：{{ baseInfo.totalOrderTotalNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/spa_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">在岗技师</div>
                      <div class="room-stat-gi-div-item-d2">空闲：{{ baseInfo.canOrderTechnicianTotal }}</div>
                      <div class="room-stat-gi-div-item-d3">工作中：{{ baseInfo.workingTechnicianTotal }} / 休息：{{
                        baseInfo.restTechnicianTotal }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/spa_stat_icon4.png" />
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>

      <!-- 移动端 实时概况 -->
      <n-gi class="block md:hidden">
        <n-spin :show="show1" description="请稍候...">
          <n-card :bordered="false" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <div class="text-base font-medium text-gray-900 mb-2">实时概况</div>
              <div class="flex items-center justify-between">
                <span class="text-xs text-gray-500">在{{ nowTime }}刷新</span>
                <n-button text class="min-h-[44px]" @click="refreshBaseInfo">
                  <img class="w-4" src="@/assets/images/pms_dashboard_refresh_icon.png" />
                  <span class="text-sm text-primary ml-1">刷新</span>
                </n-button>
              </div>
            </template>
            <n-grid :cols="2" x-gap="12" y-gap="12">
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">今日预约数</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">
                      总数：{{ baseInfo.todayOrderTotalNum }}
                    </div>
                    <div class="text-xs text-gray-500 mt-2">
                      到店：{{ baseInfo.todayDdOrderTotalNum }} / 上门：{{ baseInfo.todaySmOrderTotalNum }}
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/spa_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">营业额</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">
                      今日：{{ baseInfo.todayOrderAmount }}
                    </div>
                    <div class="text-xs text-gray-500 mt-2">
                      累计：{{ baseInfo.totalOrderAmount }}
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/spa_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">累计预约数</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">
                      有效：{{ baseInfo.effectiveTotalOrderTotalNum }}
                    </div>
                    <div class="text-xs text-gray-500 mt-2">
                      全部：{{ baseInfo.totalOrderTotalNum }}
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/spa_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">在岗技师</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">
                      空闲：{{ baseInfo.canOrderTechnicianTotal }}
                    </div>
                    <div class="text-xs text-gray-500 mt-2">
                      工作中：{{ baseInfo.workingTechnicianTotal }} / 休息：{{ baseInfo.restTechnicianTotal }}
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/spa_stat_icon4.png" />
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>

      <!-- PC端 今日订单 -->
      <n-gi class="hidden md:block">
        <n-spin :show="show2" description="请稍候...">
          <n-card :bordered="false" :header-style="{
            padding: '25px 20px 18px',
          }" :content-style="{
            padding: '0 20px 23px',
          }">
            <template #header>
              <div style="display: flex">
                <div class="room-stat-title">今日订单</div>
              </div>
            </template>
            <n-grid :cols="5" x-gap="28">
              <n-gi>
                <div class="today-stat-div">
                  <div>待付款订单</div>
                  <div>{{ todayInfo.waitPayOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div">
                  <div>待确认订单</div>
                  <div>{{ todayInfo.waitConfirmOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div">
                  <div>待服务订单</div>
                  <div>{{ todayInfo.waitServeOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div">
                  <div>服务中订单</div>
                  <div>{{ todayInfo.serviceDoingOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div">
                  <div>服务完成订单</div>
                  <div>{{ todayInfo.serviceCompleteOrder }}</div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>

      <!-- 移动端 今日订单 -->
      <n-gi class="block md:hidden">
        <n-spin :show="show2" description="请稍候...">
          <n-card :bordered="false" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <div class="text-base font-medium text-gray-900">今日订单</div>
            </template>
            <n-grid :cols="2" x-gap="12" y-gap="12">
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">待付款订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.waitPayOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">待确认订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.waitConfirmOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">待服务订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.waitServeOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">服务中订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.serviceDoingOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">服务完成订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.serviceCompleteOrder }}</div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>

      <!-- PC端 排行区域 -->
      <n-gi class="hidden md:block">
        <n-grid :cols="2" x-gap="15">
          <n-gi>
            <n-spin :show="show3" description="请稍候...">
              <div class="food-stat-div">
                <div class="food-stat-div-top">
                  <div class="food-stat-div-top-title">服务套餐排行</div>
                  <div class="food-stat-div-top-tab">
                    <div @click="handleUpdateValue1('order_num')" :class="rankTab1 == 'order_num' ? 'active' : ''">预定量
                    </div>
                    <div @click="handleUpdateValue1('order_amount')"
                      :class="rankTab1 == 'order_amount' ? 'active' : ''">预定金额
                    </div>
                  </div>
                </div>
                <div class="data-table">
                  <div class="data-table-thead">
                    <div class="data-table-td-d1">排名</div>
                    <div class="data-table-td-d2">服务套餐名称</div>
                    <div class="data-table-td-d3">{{ rankTab1 == 'order_num' ? '预定量（单）' : '交易额（JPY）' }}</div>
                  </div>
                  <div class="data-table-tbody" v-for="(item, index) in serviceLine" :key="index">
                    <div class="data-table-td-d1">{{ index + 1 }}</div>
                    <div class="data-table-td-d2">
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          {{ item.name }}
                        </template>
                        {{ item.name }}
                      </n-tooltip>
                    </div>
                    <div class="data-table-td-d3">{{ rankTab1 == 'order_num' ? item.totalOrderNum :
                      item.totalOrderAmount }}
                    </div>
                  </div>
                </div>
              </div>
            </n-spin>
          </n-gi>
          <n-gi>
            <n-spin :show="show4" description="请稍候...">
              <div class="food-stat-div">
                <div class="food-stat-div-top">
                  <div class="food-stat-div-top-title">服务人员排行</div>
                  <div class="food-stat-div-top-tab">
                    <div @click="handleUpdateValue2('order_num')" :class="rankTab2 == 'order_num' ? 'active' : ''">服务量
                    </div>
                    <div @click="handleUpdateValue2('order_amount')"
                      :class="rankTab2 == 'order_amount' ? 'active' : ''">服务金额
                    </div>
                  </div>
                </div>
                <div class="data-table">
                  <div class="data-table-thead">
                    <div class="data-table-td-d1">排名</div>
                    <div class="data-table-td-d2">服务人员</div>
                    <div class="data-table-td-d3">{{ rankTab2 == 'order_num' ? '有效服务量（单）' : '有效服务金额（JPY）' }}</div>
                  </div>
                  <div class="data-table-tbody" v-for="(item, index) in technicianLine" :key="index">
                    <div class="data-table-td-d1">{{ index + 1 }}</div>
                    <div class="data-table-td-d2">
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          {{ item.name }}
                        </template>
                        {{ item.name }}
                      </n-tooltip>
                    </div>
                    <div class="data-table-td-d3">{{ rankTab2 == 'order_num' ? item.payOrderNum : item.payOrderAmount }}
                    </div>
                  </div>
                </div>
              </div>
            </n-spin>
          </n-gi>
        </n-grid>
      </n-gi>

      <!-- 移动端 排行区域 -->
      <n-gi class="block md:hidden">
        <div class="space-y-4">
          <!-- 服务套餐排行 -->
          <n-spin :show="show3" description="请稍候...">
            <div class="food-stat-div-mobile">
              <div class="food-stat-div-top-mobile">
                <div class="text-base font-medium text-gray-900 mb-3">服务套餐排行</div>
                <div class="food-stat-div-top-tab-mobile">
                  <div @click="handleUpdateValue1('order_num')" :class="rankTab1 == 'order_num' ? 'active' : ''">预定量
                  </div>
                  <div @click="handleUpdateValue1('order_amount')" :class="rankTab1 == 'order_amount' ? 'active' : ''">
                    预定金额
                  </div>
                </div>
              </div>
              <div class="data-table-mobile">
                <div class="data-table-thead-mobile">
                  <div class="w-12 text-center text-xs font-medium text-gray-700">排名</div>
                  <div class="flex-1 text-xs font-medium text-gray-700 ml-2">服务套餐名称</div>
                  <div class="w-24 text-right text-xs font-medium text-gray-700">
                    {{ rankTab1 == 'order_num' ? '预定量（单）' : '交易额（JPY）' }}
                  </div>
                </div>
                <div class="data-table-tbody-mobile" v-for="(item, index) in serviceLine" :key="index">
                  <div class="w-12 text-center text-sm text-gray-700">{{ index + 1 }}</div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate">
                    {{ item.name }}
                  </div>
                  <div class="w-24 text-right text-sm text-gray-700">
                    {{ rankTab1 == 'order_num' ? item.totalOrderNum : item.totalOrderAmount }}
                  </div>
                </div>
              </div>
            </div>
          </n-spin>

          <!-- 服务人员排行 -->
          <n-spin :show="show4" description="请稍候...">
            <div class="food-stat-div-mobile">
              <div class="food-stat-div-top-mobile">
                <div class="text-base font-medium text-gray-900 mb-3">服务人员排行</div>
                <div class="food-stat-div-top-tab-mobile">
                  <div @click="handleUpdateValue2('order_num')" :class="rankTab2 == 'order_num' ? 'active' : ''">服务量
                  </div>
                  <div @click="handleUpdateValue2('order_amount')" :class="rankTab2 == 'order_amount' ? 'active' : ''">
                    服务金额
                  </div>
                </div>
              </div>
              <div class="data-table-mobile">
                <div class="data-table-thead-mobile">
                  <div class="w-12 text-center text-xs font-medium text-gray-700">排名</div>
                  <div class="flex-1 text-xs font-medium text-gray-700 ml-2">服务人员</div>
                  <div class="w-28 text-right text-xs font-medium text-gray-700">
                    {{ rankTab2 == 'order_num' ? '有效服务量（单）' : '有效服务金额（JPY）' }}
                  </div>
                </div>
                <div class="data-table-tbody-mobile" v-for="(item, index) in technicianLine" :key="index">
                  <div class="w-12 text-center text-sm text-gray-700">{{ index + 1 }}</div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate">
                    {{ item.name }}
                  </div>
                  <div class="w-28 text-right text-sm text-gray-700">
                    {{ rankTab2 == 'order_num' ? item.payOrderNum : item.payOrderAmount }}
                  </div>
                </div>
              </div>
            </div>
          </n-spin>
        </div>
      </n-gi>
    </n-grid>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { formatToDateTime } from "@/utils/dateUtil";
import { dashboard } from "@/api/spaDashboard";

const nowTime = ref(formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss'));
const show1 = ref(false);
const show2 = ref(false);
const show3 = ref(false);
const show4 = ref(false);
const rankTab1 = ref('order_num');
const rankTab2 = ref('order_num');

const baseInfo = ref({
  todayOrderTotalNum: 0, //今日预约总数
  todayDdOrderTotalNum: 0, //今日预约到店总数
  todaySmOrderTotalNum: 0, //今日预约上门总数
  todayOrderAmount: 0, //今日营业额
  totalOrderAmount: 0, //总营业额
  totalOrderTotalNum: 0, //全部预约数
  effectiveTotalOrderTotalNum: 0, //有效预约数
  canOrderTechnicianTotal: 0,  //空闲中技师数
  workingTechnicianTotal: 0,  //工作中技师数
  restTechnicianTotal: 0,  //休息中技师数
})

const todayInfo = ref({
  waitPayOrder: 0,//待付款订单
  waitPayOrderAmount: 0,//待付款订单金额
  waitConfirmOrder: 0,//待确认订单
  waitConfirmOrderAmount: 0,//待确认订单金额
  waitServeOrder: 0,//待服务订单
  waitServeOrderAmount: 0,//待服务订单金额
  serviceDoingOrder: 0,//服务中订单
  serviceDoingOrderAmount: 0,//服务中订单金额
  serviceCompleteOrder: 0,//服务完成订单
  serviceCompleteOrderAmount: 0,//服务完成订单金额
})

const serviceLine = ref([])
const technicianLine = ref([])

function handleUpdateValue1(value) {
  rankTab1.value = value
  Load3()
}
function handleUpdateValue2(value) {
  rankTab2.value = value
  Load4()
}

const Load1 = () => {
  show1.value = true;
  dashboard({
    type: 'base',
  })
    .then((res) => {
      baseInfo.value = res.details
      console.log(baseInfo.value)
      show1.value = false;
    }).catch((err) => {
      show1.value = false;
    })
};

const Load2 = () => {
  show2.value = true;
  dashboard({
    type: 'todayOrder',
  })
    .then((res) => {
      todayInfo.value = res.todayOrder
      console.log(todayInfo.value)
      show2.value = false;
    }).catch((err) => {
      show2.value = false;
    })
};

const Load3 = () => {
  show3.value = true;
  dashboard({
    type: 'serviceLine',
    service_line_type: rankTab1.value == 'order_num' ? 'num' : 'amount'
  })
    .then((res) => {
      serviceLine.value = res.serviceLine
      show3.value = false;
    }).catch((err) => {
      show3.value = false;
    })
};

const Load4 = () => {
  show4.value = true;
  dashboard({
    type: 'technicianLine',
    technician_line_type: rankTab2.value == 'order_num' ? 'num' : 'amount'
  })
    .then((res) => {
      technicianLine.value = res.technicianLine
      show4.value = false;
    }).catch((err) => {
      show4.value = false;
    })
};

function refreshBaseInfo() {
  nowTime.value = formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss');
  Load1();
}

onMounted(() => {
  Load1();
  Load2();
  Load3();
  Load4();
});

</script>

<style scoped lang="less">
.base-title {
  font-size: 14px;
  font-weight: 400;
  color: rgb(118, 124, 130);
}

.base-cont {
  margin-top: 4px;
  font-size: 24px;
  font-weight: 400;
  color: rgb(51, 54, 57);
}

.base-sub {
  margin-top: 4px;
  font-size: 14px;
  color: #8B8B8B;
}

.room-stat-title {
  font-size: 20px;
  color: #3D3D3D;
  font-weight: 500;
  line-height: 28px
}

.room-stat-div {
  display: flex;
  align-items: center;
  margin-top: 11px
}

.room-stat-time {
  font-size: 14px;
  color: #8B8B8B;
  font-weight: 400;
  line-height: 20px
}

.room-stat-refresh {
  font-size: 14px;
  color: #156BFF;
  line-height: 20px;
  margin-left: 5px
}

.room-stat-gi-div {
  border-radius: 2px;
  background: #F9FBFF;
  display: flex;
  align-items: center;
  position: relative;
  justify-content: flex-end;

  .room-stat-gi-div-abso {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    align-items: center;
  }

  .room-stat-gi-div-item {
    flex: 1;
    margin-left: 20px;

    .room-stat-gi-div-item-d1 {
      font-size: 14px;
      color: #3D3D3D;
      font-weight: 500;
      line-height: 20px;
    }

    .room-stat-gi-div-item-d2 {
      margin-top: 8px;
      font-size: 28px;
      color: #3D3D3D;
      font-weight: 600;
      line-height: 42px;
    }

    .room-stat-gi-div-item-d3 {
      margin-top: 7px;
      font-weight: 400;
      font-size: 14px;
      color: #8B8B8B;
      line-height: 20px;
    }
  }

  .room-stat-gi-div-image {
    width: 140px;
    height: 137px;
  }
}

.today-stat-div {
  padding: 12px 0;

  div {
    &:first-child {
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
    }

    &:last-child {
      margin-top: 5px;
      font-weight: 700;
      font-size: 24px;
      color: #3D3D3D;
      line-height: 42px;
    }
  }
}

.food-stat-div {
  background: #FFFFFF;
  border-radius: 4px;
  padding: 25px 20px 19px;
  height: 354px;

  .food-stat-div-top {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .food-stat-div-top-title {
      margin-left: 9px;
      font-weight: 500;
      font-size: 18px;
      color: #3D3D3D;
      line-height: 25px;
    }

    .food-stat-div-top-tab {
      padding: 3px 6px;
      background: #F2F3F8;
      border-radius: 2px;
      display: flex;

      div {
        cursor: pointer;
        padding: 1px 12px;
        font-weight: 400;
        font-size: 12px;
        color: #4E5969;
        line-height: 20px;

        &.active {
          background: #FFFFFF;
          border-radius: 2px;
          color: #1664FF;
          font-weight: 500;
        }
      }
    }
  }
}

.data-table {
  margin-top: 17px;

  .data-table-thead {
    background: #F2F3F8;
    border-top-left-radius: 4px;
    border-top-right-radius: 4px;
    display: flex;

    div {
      font-weight: 500;
      font-size: 14px;
      color: #1D2129;
      line-height: 40px;
    }

    .data-table-td-d1 {
      width: 15.4%;
      padding-left: 2.1%;
    }

    .data-table-td-d2 {
      width: 60.3%;
      padding-left: 1.5%;
    }

    .data-table-td-d3 {
      width: 24.3%;
      padding-left: 1.5%;
    }
  }

  .data-table-tbody {
    display: flex;
    border-bottom: 1px solid #F2F3F8;

    div {
      font-weight: 400;
      font-size: 14px;
      color: #1D2129;
      line-height: 44px;
    }

    .data-table-td-d1 {
      width: 15.4%;
      padding-left: 2.6%;
    }

    .data-table-td-d2 {
      width: 60.3%;
      padding-left: 1.5%;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
    }

    .data-table-td-d3 {
      width: 24.3%;
      padding-left: 1.5%;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
    }
  }
}

// 移动端样式
.room-stat-gi-div-mobile {
  border-radius: 12px;
  background: linear-gradient(135deg, #f9fbff 0%, #f0f4f8 100%);
  display: flex;
  align-items: flex-start;
  position: relative;
  padding: 14px;
  min-height: 120px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  overflow: hidden;

  .room-stat-gi-div-item-mobile {
    flex: 1;
    z-index: 1;
    min-width: 0;
  }

  .room-stat-gi-div-image-mobile {
    position: absolute;
    right: 12px;
    bottom: 12px;
    width: 100px;
    height: 100px;
    z-index: 0;
    pointer-events: none;

    img {
      width: 100%;
      height: 100%;
      object-fit: contain;
    }
  }
}

.today-stat-div-mobile {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 12px;
  padding: 14px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  min-height: 90px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.food-stat-div-mobile {
  background: #FFFFFF;
  border-radius: 12px;
  padding: 16px;
  border: 1px solid #f0f0f0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);

  .food-stat-div-top-mobile {
    display: flex;
    flex-direction: column;
    margin-bottom: 12px;

    .food-stat-div-top-tab-mobile {
      padding: 4px;
      background: #F2F3F8;
      border-radius: 8px;
      display: flex;
      gap: 4px;
      margin-top: 8px;

      div {
        flex: 1;
        cursor: pointer;
        padding: 8px 12px;
        font-weight: 400;
        font-size: 13px;
        color: #4E5969;
        line-height: 20px;
        text-align: center;
        border-radius: 6px;
        transition: all 0.2s;
        min-height: 36px;
        display: flex;
        align-items: center;
        justify-content: center;

        &.active {
          background: #FFFFFF;
          color: #1664FF;
          font-weight: 500;
          box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
        }
      }
    }
  }
}

.data-table-mobile {
  .data-table-thead-mobile {
    background: #F2F3F8;
    border-radius: 8px;
    display: flex;
    align-items: center;
    padding: 10px 12px;
    margin-bottom: 8px;
  }

  .data-table-tbody-mobile {
    display: flex;
    align-items: center;
    padding: 12px;
    border-bottom: 1px solid #f0f0f0;
    min-height: 48px;

    &:last-child {
      border-bottom: none;
    }
  }
}
</style>
