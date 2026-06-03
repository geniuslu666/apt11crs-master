<template>
  <div class="px-0">
    <!-- PC端标题和下拉框 -->
    <div class="title-name hidden md:block">
      <!--      <div>{{ namegetlang(userStore.getuserPms.nameLanguage) }}</div>-->

      <div class="dropdown-container" style="width: 500px">
        <!-- 自定义触发元素 -->
        <n-popover trigger="click" placement="bottom-start" :show-arrow="false" :overlap="false" raw>
          <template #trigger>
            <div class="trigger-div" @click="handleClick">
              {{ propertyName }}
              <!--              {{ namegetlang(userStore.getuserPms.nameLanguage) || '全部物业' }}-->
              <n-icon :component="isOpen ? ChevronUp : ChevronDown" />
            </div>
          </template>

          <!-- 下拉选择器 -->
          <n-select style="width: 500px" v-model:value="selectedValue" filterable :filter="filterHandler"
            :options="sortedOptions" :render-label="renderLabel" value-field="id"
            :menu-props="{ style: { width: '500px' } }" @update:value="handleSelect" />
        </n-popover>
      </div>

    </div>
    <!-- 移动端标题和下拉框 -->
    <div class="block md:hidden mb-4">
      <div class="text-xl font-bold text-gray-900 mb-4">仪表板</div>
      <div class="dropdown-container w-full">
        <!-- 自定义触发元素 -->
        <n-popover trigger="click" placement="bottom-start" :show-arrow="false" :overlap="false" raw>
          <template #trigger>
            <div class="trigger-div-mobile" @click="handleClick">
              <span class="flex-1 text-left truncate">{{ propertyName }}</span>
              <n-icon :component="isOpen ? ChevronUp : ChevronDown" class="ml-2 flex-shrink-0" />
            </div>
          </template>

          <!-- 下拉选择器 -->
          <n-select class="w-full" v-model:value="selectedValue" filterable :filter="filterHandler"
            :options="sortedOptions" :render-label="renderLabel" value-field="id"
            :menu-props="{ style: { width: '100%' } }" @update:value="handleSelect" />
        </n-popover>
      </div>
    </div>
    <n-grid :cols="1" y-gap="15">
      <n-gi>
        <!-- PC端房态统计卡片 -->
        <n-spin :show="show1" description="请稍候...">
          <n-card :bordered="false" class="hidden md:block" :header-style="{
            padding: '25px 20px 26px',
          }" :content-style="{
            padding: '0 20px 30px',
          }">
            <template #header>
              <div class="room-stat-title">房态统计</div>
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
                <div class="room-stat-gi-div"
                  style="background: linear-gradient( 180deg, rgba(228,232,241,0.5) 0%, rgba(249,250,251,0.5) 100%);">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日预抵</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.import_bookings }}</div>
                      <div class="room-stat-gi-div-item-d3">已办理入住：<span>{{ baseInfo.checkin_bookings }}</span></div>
                    </div>
                  </div>
                  <img style="width: 140px;" src="@/assets/images/room_stat_icon1.png" />
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div"
                  style="background: linear-gradient( 180deg, rgba(222,234,255,0.5) 0%, rgba(248,250,255,0.5) 100%);">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日预离</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.export_bookings }}</div>
                      <div class="room-stat-gi-div-item-d3">已办理退房：<span>{{ baseInfo.checkout_bookings }}</span></div>
                    </div>
                  </div>
                  <img style="width: 140px;" src="@/assets/images/room_stat_icon2.png" />
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div"
                  style="background: linear-gradient( 180deg, rgba(216,241,255,0.5) 0%, rgba(247,251,255,0.5) 100%);">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">在住客房数</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.checked_bookings }}</div>
                      <div class="room-stat-gi-div-item-d3">客房总数：<span>{{ baseInfo.all_room }}</span></div>
                    </div>
                  </div>
                  <img style="width: 140px;" src="@/assets/images/room_stat_icon3.png" />
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div"
                  style="background: linear-gradient( 180deg, rgba(246,232,240,0.5) 0%, rgba(251,249,250,0.5) 100%);">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日预订</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.today_stays }}</div>
                      <div class="room-stat-gi-div-item-d3">APP预定：<span>{{ baseInfo.today_app_stays }}</span></div>
                    </div>
                  </div>
                  <img style="width: 140px;" src="@/assets/images/room_stat_icon4.png" />
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
        <!-- 移动端房态统计卡片 -->
        <n-spin :show="show1" description="请稍候...">
          <n-card :bordered="false" class="block md:hidden mb-4" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <div class="text-base font-medium text-gray-900 mb-2">房态统计</div>
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
                <div class="room-stat-gi-div-mobile"
                  style="background: linear-gradient( 180deg, rgba(228,232,241,0.5) 0%, rgba(249,250,251,0.5) 100%);">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">今日预抵</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ baseInfo.import_bookings }}</div>
                    <div class="text-xs text-gray-500 mt-2">已办理入住：<span class="text-gray-700 font-medium">{{
                      baseInfo.checkin_bookings }}</span></div>
                  </div>
                  <img class="w-20 absolute right-2 bottom-2 opacity-50" src="@/assets/images/room_stat_icon1.png" />
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile"
                  style="background: linear-gradient( 180deg, rgba(222,234,255,0.5) 0%, rgba(248,250,255,0.5) 100%);">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">今日预离</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ baseInfo.export_bookings }}</div>
                    <div class="text-xs text-gray-500 mt-2">已办理退房：<span class="text-gray-700 font-medium">{{
                      baseInfo.checkout_bookings }}</span></div>
                  </div>
                  <img class="w-20 absolute right-2 bottom-2 opacity-50" src="@/assets/images/room_stat_icon2.png" />
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile"
                  style="background: linear-gradient( 180deg, rgba(216,241,255,0.5) 0%, rgba(247,251,255,0.5) 100%);">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">在住客房数</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ baseInfo.checked_bookings }}</div>
                    <div class="text-xs text-gray-500 mt-2">客房总数：<span class="text-gray-700 font-medium">{{
                      baseInfo.all_room }}</span></div>
                  </div>
                  <img class="w-20 absolute right-2 bottom-2 opacity-50" src="@/assets/images/room_stat_icon3.png" />
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile"
                  style="background: linear-gradient( 180deg, rgba(246,232,240,0.5) 0%, rgba(251,249,250,0.5) 100%);">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">今日预订</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ baseInfo.today_stays }}</div>
                    <div class="text-xs text-gray-500 mt-2">APP预定：<span class="text-gray-700 font-medium">{{
                      baseInfo.today_app_stays }}</span></div>
                  </div>
                  <img class="w-20 absolute right-2 bottom-2 opacity-50" src="@/assets/images/room_stat_icon4.png" />
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>
      <n-gi>
        <!-- PC端营业额概况卡片 -->
        <n-spin :show="show2" description="请稍候...">
          <n-card :bordered="false" class="hidden md:block" :header-style="{
            padding: '25px 20px 18px',
          }" :content-style="{
            padding: '0 20px 23px',
          }">
            <template #header>
              <span class="room-stat-title">营业额概况</span><span
                style="font-weight: 400;font-size: 14px;color: #8B8B8B;line-height: 34px">（单位：日元）</span>
            </template>
            <template #header-extra>
              <div style="width: 156px;">
                <n-tabs type="segment" animated default-value="today" pane-wrapper-style="display:none"
                  @update:value="handleUpdateValue">
                  <n-tab-pane name="today" tab="今日"></n-tab-pane>
                  <n-tab-pane name="week" tab="本周"></n-tab-pane>
                  <n-tab-pane name="month" tab="本月"></n-tab-pane>
                </n-tabs>
              </div>
            </template>
            <div style="display: flex">
              <div style="width: 25%;height: 115px">
                <div class="amount-stat-title">总收入额</div>
                <div class="amount-stat-amount">{{ amountInfo.totalIncome }}</div>
                <div class="amount-stat-div">
                  有效收入：<span>{{ amountInfo.effectIncome }}</span>
                  <n-tooltip trigger="hover">
                    <template #trigger>
                      <n-button text style="margin-left: 6px">
                        <img style="width: 12px;" src="@/assets/images/pms_dashboard_question_icon.png" />
                      </n-button>
                    </template>
                    <span>有效收入=总收入-积分抵扣-优惠券抵扣-所有退款</span>
                  </n-tooltip>
                </div>
              </div>
              <div style="flex: 1">
                <n-grid :cols="4" x-gap="11">
                  <n-gi>
                    <div style="overflow: hidden">
                      <img style="float: left;width: 32px;margin-top: 29px"
                        src="@/assets/images/pms_dashboard_amount_icon1.png" />
                      <div style="float: left;flex: 1;margin-left: 8px">
                        <div class="amount-stat-title">PAYCLOUD</div>
                        <div class="amount-stat-amount">{{ amountInfo.payCloudIncome }}</div>
                        <div class="amount-stat-div">退：<span>{{ amountInfo.payCloudRefund }}</span></div>
                      </div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div style="overflow: hidden">
                      <img style="float: left;width: 32px;margin-top: 29px"
                        src="@/assets/images/pms_dashboard_amount_icon2.png" />
                      <div style="float: left;flex: 1;margin-left: 8px">
                        <div class="amount-stat-title">Strip信用卡</div>
                        <div class="amount-stat-amount">{{ amountInfo.stripIncome }}</div>
                        <div class="amount-stat-div">退：<span>{{ amountInfo.stripRefund }}</span></div>
                      </div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div style="overflow: hidden">
                      <img style="float: left;width: 32px;margin-top: 29px"
                        src="@/assets/images/pms_dashboard_amount_icon3.png" />
                      <div style="float: left;flex: 1;margin-left: 8px">
                        <div class="amount-stat-title">
                          积分抵扣支付
                          <n-tooltip trigger="hover">
                            <template #trigger>
                              <n-button text style="margin-left: 2px">
                                <img style="width: 12px;" src="@/assets/images/pms_dashboard_question_icon.png" />
                              </n-button>
                            </template>
                            <span>积分兑换支付金额比例为1:{{ amountInfo.exchangeRate }}，即1积分={{ amountInfo.exchangeRate }}日元</span>
                          </n-tooltip>
                        </div>
                        <div class="amount-stat-amount">{{ amountInfo.balIncome }}</div>
                        <div class="amount-stat-div">退：<span>{{ amountInfo.balRefund }}</span></div>
                      </div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div style="overflow: hidden">
                      <img style="float: left;width: 32px;margin-top: 29px"
                        src="@/assets/images/pms_dashboard_amount_icon4.png" />
                      <div style="float: left;flex: 1;margin-left: 8px">
                        <div class="amount-stat-title">优惠券抵扣</div>
                        <div class="amount-stat-amount">{{ amountInfo.couponIncome }}</div>
                        <div class="amount-stat-div">退：<span>{{ amountInfo.couponRefund }}</span></div>
                      </div>
                    </div>
                  </n-gi>
                </n-grid>
              </div>
            </div>

          </n-card>
        </n-spin>
        <!-- 移动端营业额概况卡片 -->
        <n-spin :show="show2" description="请稍候...">
          <n-card :bordered="false" class="block md:hidden mb-4" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <div class="mb-4">
                <div class="flex items-center justify-between mb-3">
                  <div>
                    <span class="text-lg font-semibold text-gray-900">营业额概况</span>
                    <span class="text-xs text-gray-500 ml-1">（单位：日元）</span>
                  </div>
                </div>
                <div class="w-full">
                  <n-tabs type="segment" animated default-value="today" pane-wrapper-style="display:none"
                    @update:value="handleUpdateValue">
                    <n-tab-pane name="today" tab="今日"></n-tab-pane>
                    <n-tab-pane name="week" tab="本周"></n-tab-pane>
                    <n-tab-pane name="month" tab="本月"></n-tab-pane>
                  </n-tabs>
                </div>
              </div>
            </template>
            <div class="space-y-4">
              <!-- 总收入额 -->
              <div class="bg-gradient-to-br from-blue-50 to-indigo-50 rounded-xl p-4 border border-blue-100">
                <div class="text-sm font-medium text-gray-600 mb-1">总收入额</div>
                <div class="text-2xl font-bold text-gray-900 mb-3">{{ amountInfo.totalIncome }}</div>
                <div class="flex items-center text-xs text-gray-600">
                  有效收入：<span class="text-gray-900 font-semibold ml-1">{{ amountInfo.effectIncome }}</span>
                  <n-tooltip trigger="hover">
                    <template #trigger>
                      <n-button text class="min-h-[44px] ml-1 p-0">
                        <img class="w-3.5 h-3.5" src="@/assets/images/pms_dashboard_question_icon.png" />
                      </n-button>
                    </template>
                    <span>有效收入=总收入-积分抵扣-优惠券抵扣-所有退款</span>
                  </n-tooltip>
                </div>
              </div>
              <!-- 支付方式统计 -->
              <n-grid :cols="2" x-gap="12" y-gap="12">
                <n-gi>
                  <div
                    class="bg-white rounded-xl p-3 border border-gray-100 shadow-sm hover:shadow-md transition-shadow">
                    <div class="flex items-start">
                      <div class="w-8 h-8 rounded-lg bg-orange-50 flex items-center justify-center flex-shrink-0">
                        <img class="w-5 h-5" src="@/assets/images/pms_dashboard_amount_icon1.png" />
                      </div>
                      <div class="ml-2 flex-1 min-w-0">
                        <div class="text-xs font-medium text-gray-600 mb-1">PAYCLOUD</div>
                        <div class="text-lg font-bold text-gray-900 mb-1">{{ amountInfo.payCloudIncome }}</div>
                        <div class="text-xs text-gray-500">退：<span class="text-red-500 font-medium">{{
                          amountInfo.payCloudRefund
                            }}</span></div>
                      </div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div
                    class="bg-white rounded-xl p-3 border border-gray-100 shadow-sm hover:shadow-md transition-shadow">
                    <div class="flex items-start">
                      <div class="w-8 h-8 rounded-lg bg-teal-50 flex items-center justify-center flex-shrink-0">
                        <img class="w-5 h-5" src="@/assets/images/pms_dashboard_amount_icon2.png" />
                      </div>
                      <div class="ml-2 flex-1 min-w-0">
                        <div class="text-xs font-medium text-gray-600 mb-1">Strip信用卡</div>
                        <div class="text-lg font-bold text-gray-900 mb-1">{{ amountInfo.stripIncome }}</div>
                        <div class="text-xs text-gray-500">退：<span class="text-red-500 font-medium">{{
                          amountInfo.stripRefund
                            }}</span></div>
                      </div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div
                    class="bg-white rounded-xl p-3 border border-gray-100 shadow-sm hover:shadow-md transition-shadow">
                    <div class="flex items-start">
                      <div class="w-8 h-8 rounded-lg bg-blue-50 flex items-center justify-center flex-shrink-0">
                        <img class="w-5 h-5" src="@/assets/images/pms_dashboard_amount_icon3.png" />
                      </div>
                      <div class="ml-2 flex-1 min-w-0">
                        <div class="flex items-center text-xs font-medium text-gray-600 mb-1">
                          <span class="truncate">积分抵扣支付</span>
                          <n-tooltip trigger="hover">
                            <template #trigger>
                              <n-button text class="min-h-[44px] ml-0.5 p-0 flex-shrink-0">
                                <img class="w-3 h-3" src="@/assets/images/pms_dashboard_question_icon.png" />
                              </n-button>
                            </template>
                            <span>积分兑换支付金额比例为1:{{ amountInfo.exchangeRate }}，即1积分={{ amountInfo.exchangeRate }}日元</span>
                          </n-tooltip>
                        </div>
                        <div class="text-lg font-bold text-gray-900 mb-1">{{ amountInfo.balIncome }}</div>
                        <div class="text-xs text-gray-500">退：<span class="text-red-500 font-medium">{{
                          amountInfo.balRefund
                            }}</span></div>
                      </div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div
                    class="bg-white rounded-xl p-3 border border-gray-100 shadow-sm hover:shadow-md transition-shadow">
                    <div class="flex items-start">
                      <div class="w-8 h-8 rounded-lg bg-purple-50 flex items-center justify-center flex-shrink-0">
                        <img class="w-5 h-5" src="@/assets/images/pms_dashboard_amount_icon4.png" />
                      </div>
                      <div class="ml-2 flex-1 min-w-0">
                        <div class="text-xs font-medium text-gray-600 mb-1">优惠券抵扣</div>
                        <div class="text-lg font-bold text-gray-900 mb-1">{{ amountInfo.couponIncome }}</div>
                        <div class="text-xs text-gray-500">退：<span class="text-red-500 font-medium">{{
                          amountInfo.couponRefund
                            }}</span></div>
                      </div>
                    </div>
                  </div>
                </n-gi>
              </n-grid>
            </div>
          </n-card>
        </n-spin>
      </n-gi>
      <n-gi>
        <!-- PC端近七天入住率图表 -->
        <n-spin :show="show3" description="请稍候...">
          <n-card :bordered="false" class="hidden md:block" :header-style="{
            padding: '25px 20px 20px',
          }">
            <template #header>
              <span class="room-stat-title">近七天入住率</span>
            </template>
            <div ref="chartRef" style="width: 100%;height: 300px"></div>
          </n-card>
        </n-spin>
        <!-- 移动端近七天入住率图表 -->
        <n-spin :show="show3" description="请稍候...">
          <n-card :bordered="false" class="block md:hidden" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <span class="text-base font-medium text-gray-900">近七天入住率</span>
            </template>
            <div ref="chartRefMobile" style="width: 100%;height: 250px"></div>
          </n-card>
        </n-spin>
      </n-gi>
    </n-grid>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, Ref, ref, watch, h, computed } from "vue";
import { useECharts } from "@/hooks/web/useECharts";
import { dashboard } from "@/api/comm";
import { formatToDateTime } from "@/utils/dateUtil";
import { useUserStore } from "@/store/modules/user";
import { getlang } from "@/utils/smjcomm";
import * as echarts from "echarts";
import { hexToRgba } from "@/utils/artDesignUtils";
import { ChevronDown, ChevronUp } from '@vicons/ionicons5'
import { storage } from "@/utils/Storage";

const show1 = ref(false);
const show2 = ref(false);
const show3 = ref(false);
const userStore = useUserStore();
const nowTime = ref(formatToDateTime(new Date(), 'yyyy-MM-dd hh:mm:ss'));
const startDate = ref(formatToDateTime(new Date(), 'yyyy-MM-dd 00:00:00'));
const endDate = ref(formatToDateTime(new Date(), 'yyyy-MM-dd 23:59:59'));
const baseInfo = ref({
  all_room: 0, //客房总数
  checked_bookings: 0, //在住客房数
  checkin_bookings: 0, //已抵数
  checkout_bookings: 0, //已离数
  export_bookings: 0, //预离数
  import_bookings: 0, //预抵数
  today_app_stays: 0, //app预定
  today_stays: 0,  //今日预定
})
const amountInfo = ref({
  totalIncome: 0, // 总收入额
  effectIncome: 0, // 有效收入
  payCloudIncome: 0, // payCloud支付
  payCloudRefund: 0, //payCloud退款
  stripIncome: 0, //stripCloud支付
  stripRefund: 0, // strip退款
  balIncome: 0, // 积分抵扣支付
  balRefund: 0,// 积分抵扣退款
  couponIncome: 0, // 优惠券抵扣
  couponRefund: 0, // 优惠券退款
  exchangeRate: 1, // 积分兑换比例
})

const chartRef = ref<HTMLDivElement | null>(null);
const chartRefMobile = ref<HTMLDivElement | null>(null);
const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);
const { setOptions: setOptionsMobile } = useECharts(chartRefMobile as Ref<HTMLDivElement>);

const namegetlang = (data) => {
  if (data) {
    return getlang(data, userStore.language).content;
  } else {
    return '暂无物业名';
  }
};

// 物业名称
const propertyName = ref(userStore.getuserPms.id > 0 ? namegetlang(userStore.getuserPms.nameLanguage) : '全部物业');

// 物业列表
const wylists = ref(storage.get('wylists'));

// 控制箭头
const isOpen = ref(false)

// 下拉选中的值
const selectedValue = ref(userStore.getuserPms.id)

// 排序
const sortedOptions = computed(() => {

  // console.log("wylists.value", wylists.value)

  const selected = wylists.value.find(opt => opt.id === selectedValue.value)
  const others = wylists.value.filter(opt => opt.id !== selectedValue.value)
  if (selectedValue.value > 0) {
    return selected ? [wylists.value.find(opt => opt.id === 0), selected, ...others.filter(opt => opt.id !== 0)] : wylists.value
  }
  return wylists.value
})

// 自定义选项渲染
const renderLabel = (option) => {
  // console.log('option', namegetlang(option.nameLanguage))
  if (!option.id) return h('div', '全部物业')
  return h('div', { class: 'flex items-center gap-2' }, [
    h('span', namegetlang(option.nameLanguage))
  ])
}

// 筛选处理
const filterHandler = (inputValue, option) => {
  // 同时匹配label和value字段
  return namegetlang(option.nameLanguage).includes(inputValue) ||
    option.id.toString().includes(inputValue)
}

// 选择处理
const handleSelect = (value: string) => {
  console.log('value', value)
  const selected = wylists.value.find(opt => opt.id === value)

  console.log('selected', selected)
  storage.set('userPms', selected);
  userStore.setuserPms(selected);
  propertyName.value = value > 0 ? namegetlang(selected.nameLanguage) : '全部物业'

  isOpen.value = false

}

// 点击触发元素
const handleClick = () => {
  // 可添加点击逻辑
  isOpen.value = !isOpen.value
}


const Load1 = () => {
  show1.value = true;
  dashboard({
    type: 'base',
    puid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
  })
    .then((res) => {
      baseInfo.value = res.Details
      show1.value = false;
    }).catch((err) => {
      show1.value = false;
    })
};

const Load2 = () => {
  show2.value = true;
  dashboard({
    type: 'amount',
    puid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
    start_date: startDate.value,
    end_date: endDate.value
  })
    .then((res) => {
      amountInfo.value = res.amountStat
      show2.value = false;
    }).catch((err) => {
      show2.value = false;
    })
};

const Load3 = () => {
  show3.value = true;
  var date = formatToDateTime(new Date(), 'yyyy-MM-dd');
  var base = Date.parse(date); // 转换为时间戳
  var oneDay = 24 * 3600 * 1000
  var daytimeArr = []
  for (var i = 1; i < 7; i++) { //前七天的时间
    var now = new Date(base -= oneDay);
    var myear = now.getFullYear();
    var month = now.getMonth() + 1;
    var mday = now.getDate()
    daytimeArr.push([myear, month >= 10 ? month : '0' + month, mday >= 10 ? mday : '0' + mday].join('-'))
  }
  daytimeArr = daytimeArr.reverse()
  daytimeArr.push(date)

  dashboard({
    type: 'echarts',
    puid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
  })
    .then((res) => {
      let CheckInTrend = [];
      let percentArr = [];
      if (res.CheckInTrend) {
        res.CheckInTrend.forEach((item, index) => {
          CheckInTrend[item.date] = item.percent
        })
        daytimeArr.forEach((item, index) => {
          percentArr[index] = CheckInTrend[item] ? CheckInTrend[item] : 0;
        })
      } else {
        percentArr = [0, 0, 0, 0, 0, 0, 0]
      }

      setOptions({
        xAxis: {
          type: 'category',
          data: daytimeArr,
          boundaryGap: false,
          axisTick: {
            show: false
          },
          axisLine: {
            show: true,
            lineStyle: {
              color: '#e8e8e8',
              width: 1
            }
          },
          axisLabel: {
            show: true,
            color: '#999',
            fontSize: 13
          }
        },
        yAxis: {
          type: 'value',
          axisLabel: {
            show: true,
            color: '#999',
            fontSize: 13,
            formatter: '{value} %'
          },
          axisLine: {
            show: true,
            lineStyle: {
              color: '#e8e8e8',
              width: 1
            }
          },
          splitLine: {
            show: true,
            lineStyle: {
              color: '#e8e8e8',
              width: 1,
              type: 'dashed' as const
            }
          }
        },
        tooltip: {
          trigger: 'axis',
          valueFormatter: (value) => value + '%'
        },
        series: [
          {
            data: percentArr,
            type: 'line',
            smooth: true,
            symbol: 'none',
            lineStyle: {
              width: 3,
              color: '#5D87FF'
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: hexToRgba('#5D87FF', 0.2).rgba
                },
                {
                  offset: 1,
                  color: hexToRgba('#5D87FF', 0.01).rgba
                }
              ])
            }
          }
        ]
      });

      // 移动端图表设置
      setOptionsMobile({
        grid: {
          top: '8%',
          bottom: '12%',
          right: '12px',
          left: '40px',
        },
        xAxis: {
          type: 'category',
          data: daytimeArr,
          boundaryGap: false,
          axisTick: {
            show: false
          },
          axisLine: {
            show: true,
            lineStyle: {
              color: '#e8e8e8',
              width: 1
            }
          },
          axisLabel: {
            show: true,
            color: '#999',
            fontSize: 11
          }
        },
        yAxis: {
          type: 'value',
          axisLabel: {
            show: true,
            color: '#999',
            fontSize: 11,
            formatter: '{value} %'
          },
          axisLine: {
            show: true,
            lineStyle: {
              color: '#e8e8e8',
              width: 1
            }
          },
          splitLine: {
            show: true,
            lineStyle: {
              color: '#e8e8e8',
              width: 1,
              type: 'dashed' as const
            }
          }
        },
        tooltip: {
          trigger: 'axis',
          valueFormatter: (value) => value + '%'
        },
        series: [
          {
            data: percentArr,
            type: 'line',
            smooth: true,
            symbol: 'none',
            lineStyle: {
              width: 2,
              color: '#5D87FF'
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: hexToRgba('#5D87FF', 0.2).rgba
                },
                {
                  offset: 1,
                  color: hexToRgba('#5D87FF', 0.01).rgba
                }
              ])
            }
          }
        ]
      });

      show3.value = false;
    }).catch((err) => {
      show3.value = false;
    })
};

function refreshBaseInfo() {
  nowTime.value = formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss');
  Load1();
}

function handleUpdateValue(value) {
  if (value == 'today') {
    startDate.value = formatToDateTime(new Date(), 'yyyy-MM-dd 00:00:00');
    endDate.value = formatToDateTime(new Date(), 'yyyy-MM-dd 23:59:59');
  } else if (value == 'week') {
    let today = new Date();
    let dayOfWeek = today.getDay();
    let weekStart = new Date(today);
    if (dayOfWeek !== 0) {
      weekStart.setDate(weekStart.getDate() - dayOfWeek + 1);
    } else {
      weekStart.setDate(weekStart.getDate() - dayOfWeek - 6);
    }
    let weekEnd = new Date(today);
    if (dayOfWeek !== 0) {
      weekEnd.setDate(weekEnd.getDate() + (7 - dayOfWeek));
    } else {
      weekEnd.setDate(weekEnd.getDate());
    }
    let start = `${weekStart.getFullYear()}-${(weekStart.getMonth() + 1) > 9 ? weekStart.getMonth() + 1 : '0' + (weekStart.getMonth() + 1)}-${weekStart.getDate() > 9 ? weekStart.getDate() : '0' + weekStart.getDate()}`
    let end = `${weekEnd.getFullYear()}-${(weekEnd.getMonth() + 1) > 9 ? weekEnd.getMonth() + 1 : '0' + (weekEnd.getMonth() + 1)}-${weekEnd.getDate() > 9 ? weekEnd.getDate() : '0' + weekEnd.getDate()}`
    startDate.value = start + ' 00:00:00';
    endDate.value = end + ' 23:59:59';
  } else if (value == 'month') {
    let today = new Date();
    let currentMonth = today.getMonth() + 1;
    const startTime = `${today.getFullYear()}-${currentMonth.toString().padStart(2, '0')}-01 00:00:00`;
    const lastDay = new Date(today.getFullYear(), currentMonth, 0).getDate();
    const endTime = `${today.getFullYear()}-${currentMonth.toString().padStart(2, '0')}-${lastDay.toString().padStart(2, '0')} 23:59:59`;
    startDate.value = startTime;
    endDate.value = endTime;
  }
  Load2();
}

//页面刷新  根据物业选择
watch(
  () => userStore.getuserPms,
  () => {
    Load1();  //当前页面重新加载的数据
    Load2();  //当前页面重新加载的数据
    Load3();  //当前页面重新加载的数据
  },
  {
    immediate: false,
    deep: true,
  }
);

onMounted(() => {
  wylists.value.unshift({ name: '全部', id: 0, uid: '' })
  // console.log("wylists.value", wylists.value)
  Load1();
  Load2();
  Load3();
});
</script>

<style scoped lang="less">
.title-name {
  margin: 15px 0 19px 20px;
  font-weight: 500;
  font-size: 30px;
  color: #3D3D3D;
  line-height: 42px;
}

.room-stat-title {
  font-size: 24px;
  color: #3D3D3D;
  font-weight: 500;
  line-height: 34px
}

.room-stat-div {
  display: flex;
  align-items: center;
  width: 100%;
  margin-top: 5px
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
  display: flex;
  align-items: center;
  justify-content: flex-end;
  position: relative;

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
    margin-left: 30px;

    .room-stat-gi-div-item-d1 {
      font-size: 14px;
      color: #3D3D3D;
      font-weight: 500;
      line-height: 20px;
    }

    .room-stat-gi-div-item-d2 {
      margin-top: 5px;
      font-size: 28px;
      color: #3D3D3D;
      font-weight: 700;
      line-height: 42px;
    }

    .room-stat-gi-div-item-d3 {
      margin-top: 10px;
      font-size: 14px;
      color: #8B8B8B;
      font-weight: 400;
      line-height: 20px;

      span {
        color: #3D3D3D;
        font-weight: 500;
      }
    }
  }
}

.amount-stat-title {
  display: flex;
  align-items: center;
  font-size: 12px;
  color: #1D2129;
  line-height: 20px;
  font-weight: 500;
}

.amount-stat-amount {
  margin-top: 8px;
  font-size: 24px;
  color: #1D2129;
  line-height: 42px;
  font-weight: 700;
}

.amount-stat-div {
  display: flex;
  align-items: center;
  margin-top: 4px;
  font-size: 14px;
  color: #8B8B8B;
  font-weight: 400;
  line-height: 20px;

  span {
    color: #D13820;
  }
}

// 移动端样式
.trigger-div-mobile {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 14px 16px;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  cursor: pointer;
  min-height: 48px;
  font-size: 15px;
  color: #1e293b;
  font-weight: 500;
  transition: all 0.2s ease;

  &:active {
    background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%);
    transform: scale(0.98);
  }
}

.room-stat-gi-div-mobile {
  border-radius: 12px;
  display: flex;
  align-items: flex-start;
  position: relative;
  padding: 14px;
  min-height: 110px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);

  .room-stat-gi-div-item-mobile {
    flex: 1;
    z-index: 1;
  }
}
</style>
