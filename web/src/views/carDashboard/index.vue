<template>
  <div class="">
    <n-grid :cols="1" :y-gap="15">
      <!-- PC端实时概况卡片 -->
      <n-gi class="hidden md:block">
        <n-spin :show="show1" description="请稍候...">
          <n-card :bordered="false" :header-style="{
            padding: '25px 20px 23px',
          }" :content-style="{
            padding: '0 20px 26px',
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
            <n-grid :cols="3" x-gap="15" y-gap="15">
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日预约数</div>
                      <div class="room-stat-gi-div-item-d2">总数：{{ baseInfo.todayOrderTotalNum }}</div>
                      <div class="room-stat-gi-div-item-d3">接机：{{ baseInfo.todayPickUpOrderTotalNum }} / 送机：{{
                        baseInfo.todayDeliveryOrderTotalNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/car_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日营业额</div>
                      <div class="room-stat-gi-div-item-d2">总计：{{ baseInfo.todayOrderAmount }}</div>
                      <div class="room-stat-gi-div-item-d3">退款：{{ baseInfo.todayOrderRefund }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/car_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">累计预约数</div>
                      <div class="room-stat-gi-div-item-d2">全部：{{ baseInfo.totalOrderTotalNum }} / 有效：{{
                        baseInfo.effectiveOrderTotalNum }}</div>
                      <div class="room-stat-gi-div-item-d3">接机：{{ baseInfo.totalPickUpOrderTotalNum }} / 送机：{{
                        baseInfo.totalDeliveryOrderTotalNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/car_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">累计营业金额</div>
                      <div class="room-stat-gi-div-item-d2">总计：{{ baseInfo.totalOrderAmount }}</div>
                      <div class="room-stat-gi-div-item-d3">退款：{{ baseInfo.totalOrderRefund }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/car_stat_icon4.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">司机概况</div>
                      <div class="room-stat-gi-div-item-d2">总数：{{ baseInfo.totalDriverNum }}</div>
                      <div class="room-stat-gi-div-item-d3">正常：{{ baseInfo.onDriverTotal }} / 休息：{{
                        baseInfo.restDriverTotal }} /
                        停用：{{ baseInfo.offDriverTotal }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/car_stat_icon5.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">车辆概况</div>
                      <div class="room-stat-gi-div-item-d2">总数：{{ baseInfo.totalCarNum }}</div>
                      <div class="room-stat-gi-div-item-d3">正常：{{ baseInfo.onCarTotal }} / 暂不可用：{{ baseInfo.offCarTotal
                      }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/car_stat_icon6.png" />
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>
      <!-- 移动端实时概况卡片 -->
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
                    <div class="text-lg font-bold text-gray-900 mt-1">总数：{{ baseInfo.todayOrderTotalNum }}</div>
                    <div class="text-xs text-gray-500 mt-2">接机：{{ baseInfo.todayPickUpOrderTotalNum }} / 送机：{{
                      baseInfo.todayDeliveryOrderTotalNum }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/car_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">今日营业额</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">总计：{{ baseInfo.todayOrderAmount }}</div>
                    <div class="text-xs text-gray-500 mt-2">退款：{{ baseInfo.todayOrderRefund }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/car_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">累计预约数</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">全部：{{ baseInfo.totalOrderTotalNum }} / 有效：{{
                      baseInfo.effectiveOrderTotalNum }}</div>
                    <div class="text-xs text-gray-500 mt-2">接机：{{ baseInfo.totalPickUpOrderTotalNum }} / 送机：{{
                      baseInfo.totalDeliveryOrderTotalNum }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/car_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">累计营业金额</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">总计：{{ baseInfo.totalOrderAmount }}</div>
                    <div class="text-xs text-gray-500 mt-2">退款：{{ baseInfo.totalOrderRefund }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/car_stat_icon4.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">司机概况</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">总数：{{ baseInfo.totalDriverNum }}</div>
                    <div class="text-xs text-gray-500 mt-2">正常：{{ baseInfo.onDriverTotal }} / 休息：{{
                      baseInfo.restDriverTotal
                    }} / 停用：{{ baseInfo.offDriverTotal }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/car_stat_icon5.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">车辆概况</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">总数：{{ baseInfo.totalCarNum }}</div>
                    <div class="text-xs text-gray-500 mt-2">正常：{{ baseInfo.onCarTotal }} / 暂不可用：{{ baseInfo.offCarTotal
                    }}
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/car_stat_icon6.png" />
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>
      <!-- PC端今日订单卡片 -->
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
            <n-grid :cols="5" y-gap="28">
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
      <!-- 移动端今日订单卡片 -->
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
      <!-- PC端排行区域 -->
      <n-gi class="hidden md:block">
        <n-grid :cols="2" x-gap="15">
          <n-gi>
            <n-spin :show="show3" description="请稍候...">
              <div class="food-stat-div">
                <div class="food-stat-div-top">
                  <div class="food-stat-div-top-title">模式排行</div>
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
                    <div class="data-table-td-d2">服务模式</div>
                    <div class="data-table-td-d3">{{ rankTab1 == 'order_num' ? '预定量（单）' : '交易额（JPY）' }}</div>
                  </div>
                  <div class="data-table-tbody" v-for="(item, index) in serviceLine" :key="index">
                    <div class="data-table-td-d1">{{ index + 1 }}</div>
                    <div class="data-table-td-d2" v-if="(item as any).serviceType == 'PICKUP'">接机</div>
                    <div class="data-table-td-d2" v-else-if="(item as any).serviceType == 'DELIVERY'">送机</div>
                    <div class="data-table-td-d2" v-else>包车</div>
                    <div class="data-table-td-d3">{{ rankTab1 == 'order_num' ? (item as any).totalOrderNum : (item as
                      any).totalOrderAmount }}</div>
                  </div>
                </div>
              </div>
            </n-spin>
          </n-gi>
          <n-gi>
            <n-spin :show="show4" description="请稍候...">
              <div class="food-stat-div">
                <div class="food-stat-div-top">
                  <div class="food-stat-div-top-title">司机排行</div>
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
                    <div class="data-table-td-d2">司机</div>
                    <div class="data-table-td-d3">{{ rankTab2 == 'order_num' ? '有效服务量（单）' : '有效服务金额（JPY）' }}</div>
                  </div>
                  <div class="data-table-tbody" v-for="(item, index) in driverLine" :key="index">
                    <div class="data-table-td-d1">{{ index + 1 }}</div>
                    <div class="data-table-td-d2">
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          {{ (item as any).name }}
                        </template>
                        {{ (item as any).name }}
                      </n-tooltip>
                    </div>
                    <div class="data-table-td-d3">{{ rankTab2 == 'order_num' ? (item as any).payOrderNum : (item as
                      any).payOrderAmount }}</div>
                  </div>
                </div>
              </div>
            </n-spin>
          </n-gi>
        </n-grid>
      </n-gi>
      <!-- 移动端排行区域 -->
      <n-gi class="block md:hidden">
        <div class="space-y-4">
          <!-- 模式排行 -->
          <n-spin :show="show3" description="请稍候...">
            <div class="food-stat-div-mobile">
              <div class="food-stat-div-top-mobile">
                <div class="text-base font-medium text-gray-900 mb-3">模式排行</div>
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
                  <div class="flex-1 text-xs font-medium text-gray-700 ml-2">服务模式</div>
                  <div class="w-20 text-right text-xs font-medium text-gray-700">{{ rankTab1 == 'order_num' ? '预定量（单）' :
                    '交易额（JPY）' }}</div>
                </div>
                <div class="data-table-tbody-mobile" v-for="(item, index) in serviceLine" :key="index">
                  <div class="w-12 text-center text-sm text-gray-700">{{ index + 1 }}</div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate" v-if="(item as any).serviceType == 'PICKUP'">
                    接机
                  </div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate"
                    v-else-if="(item as any).serviceType == 'DELIVERY'">
                    送机</div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate" v-else>包车</div>
                  <div class="w-20 text-right text-sm text-gray-700">{{ rankTab1 == 'order_num' ? (item as
                    any).totalOrderNum
                    : (item as any).totalOrderAmount }}</div>
                </div>
              </div>
            </div>
          </n-spin>
          <!-- 司机排行 -->
          <n-spin :show="show4" description="请稍候...">
            <div class="food-stat-div-mobile">
              <div class="food-stat-div-top-mobile">
                <div class="text-base font-medium text-gray-900 mb-3">司机排行</div>
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
                  <div class="flex-1 text-xs font-medium text-gray-700 ml-2">司机</div>
                  <div class="w-20 text-right text-xs font-medium text-gray-700">{{ rankTab2 == 'order_num' ? '有效服务量（单）'
                    :
                    '有效服务金额（JPY）' }}</div>
                </div>
                <div class="data-table-tbody-mobile" v-for="(item, index) in driverLine" :key="index">
                  <div class="w-12 text-center text-sm text-gray-700">{{ index + 1 }}</div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate">
                    {{ (item as any).name }}
                  </div>
                  <div class="w-20 text-right text-sm text-gray-700">{{ rankTab2 == 'order_num' ? (item as
                    any).payOrderNum :
                    (item as any).payOrderAmount }}</div>
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
import { dashboard } from "@/api/carDashboard";

const nowTime = ref(formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss'));
const show1 = ref(false);
const show2 = ref(false);
const show3 = ref(false);
const show4 = ref(false);
const rankTab1 = ref('order_num');
const rankTab2 = ref('order_num');

const baseInfo = ref({
  todayOrderTotalNum: 0, //今日预约总数
  todayPickUpOrderTotalNum: 0, //今日预约接机总数
  todayDeliveryOrderTotalNum: 0, //今日预约接机送机总数
  todayCharteredCarOrderTotalNum: 0, //今日预约包车送机总数
  todayOrderAmount: 0, //今日营业额
  todayOrderRefund: 0, //今日退款额
  totalOrderTotalNum: 0, //全部累计预约数
  effectiveOrderTotalNum: 0, //有效累计预约数
  totalPickUpOrderTotalNum: 0, //全部预约接机总数
  totalDeliveryOrderTotalNum: 0, //全部预约接机送机总数
  totalCharteredCarOrderTotalNum: 0, //全部预约包车送机总数
  totalOrderAmount: 0, //总营业额
  totalOrderRefund: 0, //总退款额
  totalDriverNum: 0, // 司机总数
  onDriverTotal: 0, // 司机正常总数
  offDriverTotal: 0,  //停用司机数
  restDriverTotal: 0,  //休息中司机数
  totalCarNum: 0, // 车辆总数
  onCarTotal: 0,  //启用车辆数
  offCarTotal: 0,  //停用车辆数
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
const driverLine = ref([])

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
    type: 'driverLine',
    driver_line_type: rankTab2.value == 'order_num' ? 'num' : 'amount'
  })
    .then((res) => {
      driverLine.value = res.driverLine
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
    right: 8px;
    bottom: 8px;
    width: 70px;
    height: 70px;
    z-index: 0;
    opacity: 0.6;
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
