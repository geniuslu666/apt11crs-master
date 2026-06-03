<template>
  <div class="">
    <!-- PC端标题 -->
    <div class="title-name hidden md:block">欢迎回来，{{ userStore.realName }}！</div>
    <!-- 移动端标题 -->
    <div class="text-lg font-semibold text-gray-900 mb-4 md:hidden">欢迎回来，{{ userStore.realName }}！</div>
    <n-spin :show="show" description="请稍候...">
      <n-grid :cols="1" y-gap="15">
        <n-gi>
          <!-- PC端会员统计卡片 -->
          <n-card :bordered="false" class="hidden md:block" :header-style="{
            padding: '25px 20px 26px',
          }" :content-style="{
            padding: '0 20px 30px',
          }">
            <template #header>
              <div class="room-stat-title">会员统计</div>
            </template>
            <n-grid :cols="4" x-gap="15">
              <n-gi>
                <div class="room-stat-gi-div">
                  <img style="width: 65px" src="@/assets/images/index/img1.png" />
                  <div class="room-stat-gi-div-item">
                    <div class="room-stat-gi-div-item-d1">新增会员</div>
                    <div class="room-stat-gi-div-item-d2">{{
                      state.dashboard.memberStat.todayRegMemberNum
                    }}</div>
                    <div class="room-stat-gi-div-item-d3">
                      同比：
                      <img v-if="state.dashboard.memberStat.todayMemberNumGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img5.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.todayMemberNumGrewPer >= 0 ?
                        state.dashboard.memberStat.todayMemberNumGrewPer :
                        -state.dashboard.memberStat.todayMemberNumGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <img style="width: 65px" src="@/assets/images/index/img2.png" />
                  <div class="room-stat-gi-div-item">
                    <div class="room-stat-gi-div-item-d1">累计会员</div>
                    <div class="room-stat-gi-div-item-d2">{{
                      state.dashboard.memberStat.totalMemberNum
                    }}</div>
                    <div class="room-stat-gi-div-item-d3">
                      较昨日：
                      <img v-if="state.dashboard.memberStat.totalMemberGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img5.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.totalMemberGrewPer >= 0 ?
                        state.dashboard.memberStat.totalMemberGrewPer :
                        -state.dashboard.memberStat.totalMemberGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <img style="width: 65px" src="@/assets/images/index/img3.png" />
                  <div class="room-stat-gi-div-item">
                    <div class="room-stat-gi-div-item-d1">新增积分</div>
                    <div class="room-stat-gi-div-item-d2">{{
                      state.dashboard.memberStat.todayIncBal
                    }}</div>
                    <div class="room-stat-gi-div-item-d3">
                      较昨日：
                      <img v-if="state.dashboard.memberStat.todayIncBalGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img5.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.todayIncBalGrewPer >= 0 ?
                        state.dashboard.memberStat.todayIncBalGrewPer :
                        -state.dashboard.memberStat.todayIncBalGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <img style="width: 65px" src="@/assets/images/index/img4.png" />
                  <div class="room-stat-gi-div-item">
                    <div class="room-stat-gi-div-item-d1">消耗积分</div>
                    <div class="room-stat-gi-div-item-d2">{{
                      state.dashboard.memberStat.todayConsumeBal
                    }}</div>
                    <div class="room-stat-gi-div-item-d3">
                      较昨日：
                      <img v-if="state.dashboard.memberStat.todayConsumeBalGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img5.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.todayConsumeBalGrewPer >= 0 ?
                        state.dashboard.memberStat.todayConsumeBalGrewPer :
                        -state.dashboard.memberStat.todayConsumeBalGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
          <!-- 移动端会员统计卡片 -->
          <n-card :bordered="false" class="block md:hidden mb-4" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <div class="text-base font-medium text-gray-900">会员统计</div>
            </template>
            <n-grid :cols="2" x-gap="12" y-gap="12">
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <img class="w-12" src="@/assets/images/index/img1.png" />
                  <div class="room-stat-gi-div-item-mobile w-full">
                    <div class="text-sm text-gray-700 font-medium">新增会员</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{
                      state.dashboard.memberStat.todayRegMemberNum
                    }}</div>
                    <div class="flex items-center mt-2 text-xs text-gray-500">
                      同比：
                      <img v-if="state.dashboard.memberStat.todayMemberNumGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img5.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.todayMemberNumGrewPer >= 0 ?
                        state.dashboard.memberStat.todayMemberNumGrewPer
                        : -state.dashboard.memberStat.todayMemberNumGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <img class="w-12" src="@/assets/images/index/img2.png" />
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm text-gray-700 font-medium">累计会员</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{
                      state.dashboard.memberStat.totalMemberNum
                    }}</div>
                    <div class="flex items-center mt-2 text-xs text-gray-500">
                      较昨日：
                      <img v-if="state.dashboard.memberStat.totalMemberGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img5.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.totalMemberGrewPer >= 0 ?
                        state.dashboard.memberStat.totalMemberGrewPer :
                        -state.dashboard.memberStat.totalMemberGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <img class="w-12" src="@/assets/images/index/img3.png" />
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm text-gray-700 font-medium">新增积分</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{
                      state.dashboard.memberStat.todayIncBal
                    }}</div>
                    <div class="flex items-center mt-2 text-xs text-gray-500">
                      较昨日：
                      <img v-if="state.dashboard.memberStat.todayIncBalGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img5.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.todayIncBalGrewPer >= 0 ?
                        state.dashboard.memberStat.todayIncBalGrewPer :
                        -state.dashboard.memberStat.todayIncBalGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <img class="w-12" src="@/assets/images/index/img4.png" />
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm text-gray-700 font-medium">消耗积分</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{
                      state.dashboard.memberStat.todayConsumeBal
                    }}</div>
                    <div class="flex items-center mt-2 text-xs text-gray-500">
                      较昨日：
                      <img v-if="state.dashboard.memberStat.todayConsumeBalGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img5.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img6.png" />
                      {{ state.dashboard.memberStat.todayConsumeBalGrewPer >= 0 ?
                        state.dashboard.memberStat.todayConsumeBalGrewPer
                        : -state.dashboard.memberStat.todayConsumeBalGrewPer }}%
                    </div>
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-gi>
        <n-gi>
          <!-- PC端订单统计卡片 -->
          <n-card :bordered="false" class="hidden md:block" :header-style="{
            padding: '25px 20px 18px',
          }" :content-style="{
            padding: '0 20px 23px',
          }">
            <template #header>
              <div style="display: inline-block">
                <n-tabs type="segment" animated pane-wrapper-style="display:none"
                  :tab-style="{ width: '50px', borderRadius: '2px' }" @update:value="handleUpdateValue">
                  <n-tab-pane name="all" tab="全部"></n-tab-pane>
                  <n-tab-pane name="hotel" tab="APP"></n-tab-pane>
                </n-tabs>
              </div>
            </template>
            <n-grid :cols="5" x-gap="28" y-gap="15">
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">今日订单<span>(单)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayOrderNum
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayOrderNum
                  }}</div>
                  <div class="tab-div2-d3">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayOrderNumGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayOrderNumGrewPer :
                        -state.dashboard.orderStat.allStat.todayOrderNumGrewPer }}</span>%
                    </template>

                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">全量订单<span>(单)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalOrderNum
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalOrderNum
                  }}</div>
                  <div class="tab-div2-d3">同比：
                    <!--                    <img style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />-->

                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalOrderNumGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalOrderNumGrewPer :
                        -state.dashboard.orderStat.allStat.totalOrderNumGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">今日退款<span>(单)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayRefundNum
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayRefundNum
                  }}</div>
                  <div class="tab-div2-d3">
                    同比：
                    <!--                    <img style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />-->
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayRefundGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayRefundGrewPer :
                        -state.dashboard.orderStat.allStat.todayRefundGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayRefundGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayRefundGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayRefundGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">全量退款<span>(单)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalRefundNum
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalRefundNum
                  }}</div>
                  <div class="tab-div2-d3">
                    同比：
                    <!--                    <img style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />-->
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalRefundGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalRefundGrewPer :
                        -state.dashboard.orderStat.allStat.totalRefundGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalRefundGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalRefundGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalRefundGrewPer }}</span>%
                    </template>

                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">全部物业<span>(幢)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalPropertyNum
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalPropertyNum
                  }}</div>
                  <div class="tab-div2-d3">
                    在线物业：
                    <span v-if="state.orderStatType == 'all'">{{
                      state.dashboard.orderStat.allStat.onlinePropertyNum
                    }}</span>
                    <span v-if="state.orderStatType == 'hotel'">{{
                      state.dashboard.orderStat.hotelStat.onlinePropertyNum
                    }}</span>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">今日订单总额<span>(日元)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayOrderMoney
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayOrderMoney
                  }}</div>
                  <div class="tab-div2-d3">
                    同比：

                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer }}</span>%
                    </template>

                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">全量订单总额<span>(日元)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalOrderMoney
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalOrderMoney
                  }}</div>
                  <div class="tab-div2-d3">
                    同比：

                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer }}</span>%
                    </template>

                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">今日退款总额<span>(日元)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayRefundMoney
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayRefundMoney
                  }}</div>
                  <div class="tab-div2-d3">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer }}</span>%
                    </template>

                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">全量退款总额<span>(日元)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalRefundMoney
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalRefundMoney
                  }}</div>
                  <div class="tab-div2-d3">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer >= 0"
                        style="width: 16px; margin-right: 2px" src="@/assets/images/index/img7.png" />
                      <img v-else style="width: 16px; margin-right: 2px" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2">
                  <div class="tab-div2-d1">全部房型<span>(个)</span></div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalRoomTypeNum
                  }}</div>
                  <div class="tab-div2-d2" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalRoomTypeNum
                  }}</div>
                  <div class="tab-div2-d3">
                    在线房型：
                    <span v-if="state.orderStatType == 'all'">{{
                      state.dashboard.orderStat.allStat.onlineRoomTypeNum
                    }}</span>
                    <span v-if="state.orderStatType == 'hotel'">{{
                      state.dashboard.orderStat.hotelStat.onlineRoomTypeNum
                    }}</span>
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
          <!-- 移动端订单统计卡片 -->
          <n-card :bordered="false" class="block md:hidden mb-4" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <div class="inline-block">
                <n-tabs type="segment" animated pane-wrapper-style="display:none"
                  :tab-style="{ width: '50px', borderRadius: '2px' }" @update:value="handleUpdateValue">
                  <n-tab-pane name="all" tab="全部"></n-tab-pane>
                  <n-tab-pane name="hotel" tab="APP"></n-tab-pane>
                </n-tabs>
              </div>
            </template>
            <n-grid :cols="2" x-gap="12" y-gap="12">
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">今日订单<span class="text-gray-500 font-normal">(单)</span>
                  </div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayOrderNum
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayOrderNum
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayOrderNumGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayOrderNumGrewPer :
                        -state.dashboard.orderStat.allStat.todayOrderNumGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayOrderNumGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">全量订单<span class="text-gray-500 font-normal">(单)</span>
                  </div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalOrderNum
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalOrderNum
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalOrderNumGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalOrderNumGrewPer :
                        -state.dashboard.orderStat.allStat.totalOrderNumGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalOrderNumGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">今日退款<span class="text-gray-500 font-normal">(单)</span>
                  </div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayRefundNum
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayRefundNum
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayRefundGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayRefundGrewPer :
                        -state.dashboard.orderStat.allStat.todayRefundGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayRefundGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayRefundGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayRefundGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">全量退款<span class="text-gray-500 font-normal">(单)</span>
                  </div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalRefundNum
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalRefundNum
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalRefundGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalRefundGrewPer :
                        -state.dashboard.orderStat.allStat.totalRefundGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalRefundGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalRefundGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalRefundGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalRefundGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">全部物业<span class="text-gray-500 font-normal">(幢)</span>
                  </div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalPropertyNum
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalPropertyNum
                  }}</div>
                  <div class="text-xs text-gray-500 mt-1">
                    在线物业：
                    <span v-if="state.orderStatType == 'all'">{{
                      state.dashboard.orderStat.allStat.onlinePropertyNum
                    }}</span>
                    <span v-if="state.orderStatType == 'hotel'">{{
                      state.dashboard.orderStat.hotelStat.onlinePropertyNum
                    }}</span>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">今日订单总额<span
                      class="text-gray-500 font-normal">(日元)</span></div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayOrderMoney
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayOrderMoney
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.todayOrderMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayOrderMoneyGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">全量订单总额<span
                      class="text-gray-500 font-normal">(日元)</span></div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalOrderMoney
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalOrderMoney
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.totalOrderMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalOrderMoneyGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">今日退款总额<span
                      class="text-gray-500 font-normal">(日元)</span></div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.todayRefundMoney
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.todayRefundMoney
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.todayRefundMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.todayRefundMoneyGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">全量退款总额<span
                      class="text-gray-500 font-normal">(日元)</span></div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalRefundMoney
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalRefundMoney
                  }}</div>
                  <div class="flex items-center mt-1 text-xs text-gray-500">
                    同比：
                    <template v-if="state.orderStatType == 'all'">
                      <img v-if="state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer :
                        -state.dashboard.orderStat.allStat.totalRefundMoneyGrewPer }}</span>%
                    </template>
                    <template v-if="state.orderStatType == 'hotel'">
                      <img v-if="state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer >= 0" class="w-4 mr-1"
                        src="@/assets/images/index/img7.png" />
                      <img v-else class="w-4 mr-1" src="@/assets/images/index/img8.png" />
                      <span>{{ state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer >= 0 ?
                        state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer :
                        -state.dashboard.orderStat.hotelStat.totalRefundMoneyGrewPer }}</span>%
                    </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="tab-div2-mobile">
                  <div class="text-sm font-medium text-gray-700">全部房型<span class="text-gray-500 font-normal">(个)</span>
                  </div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'all'">{{
                    state.dashboard.orderStat.allStat.totalRoomTypeNum
                  }}</div>
                  <div class="text-lg font-bold text-gray-900 mt-1" v-if="state.orderStatType == 'hotel'">{{
                    state.dashboard.orderStat.hotelStat.totalRoomTypeNum
                  }}</div>
                  <div class="text-xs text-gray-500 mt-1">
                    在线房型：
                    <span v-if="state.orderStatType == 'all'">{{
                      state.dashboard.orderStat.allStat.onlineRoomTypeNum
                    }}</span>
                    <span v-if="state.orderStatType == 'hotel'">{{
                      state.dashboard.orderStat.hotelStat.onlineRoomTypeNum
                    }}</span>
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-gi>
        <n-gi>
          <!-- PC端图表区域 -->
          <div class="hidden md:flex" style="gap: 8px 15px">
            <div style="flex: 1; height: 375px">
              <n-card :bordered="false" style="height: 375px" :header-style="{
                padding: '25px 20px 20px',
              }" :content-style="cardContStyle">
                <template #header>
                  <span class="room-stat-title1">七日订单趋势</span>
                </template>
                <div ref="chartRef" style="width: 100%; height: 80%"></div>
              </n-card>
            </div>
            <div style="width: 20.5%; height: 375px">
              <n-card :bordered="false" style="height: 375px" :header-style="{
                padding: '25px 20px 20px',
              }" :content-style="cardContStyle">
                <template #header>
                  <span class="room-stat-title1">物业预订量排行</span>
                </template>
                <div class="data-table">
                  <div class="data-table-thead">
                    <div class="data-table-td-d1">排名</div>
                    <div class="data-table-td-d2">物业名称</div>
                    <div class="data-table-td-d3">占比</div>
                  </div>

                  <div v-for="(item, index) in state.dashboard.propertyList" :key="index">
                    <div class="data-table-tbody">
                      <div class="data-table-td-d1">{{ index + 1 }}</div>
                      <div class="data-table-td-d2">
                        <n-tooltip trigger="hover">
                          <template #trigger>
                            {{ (item as any)?.propertyDetail?.name }}
                          </template>
                          {{ (item as any)?.propertyDetail?.name }}
                        </n-tooltip>
                      </div>
                      <div class="data-table-td-d3">{{ (item as any)?.rate }}%</div>
                    </div>
                  </div>
                </div>
              </n-card>
            </div>
            <div style="width: 20.5%; height: 375px">
              <n-card :bordered="false" style="height: 375px" :header-style="{
                padding: '25px 20px 20px',
              }" :content-style="cardContStyle">
                <template #header>
                  <span class="room-stat-title1">OTA渠道排名排行</span>
                </template>
                <div class="data-table">
                  <div class="data-table-thead">
                    <div class="data-table-td-d1">排名</div>
                    <div class="data-table-td-d2">渠道名称</div>
                    <div class="data-table-td-d3">占比</div>
                  </div>
                  <div v-for="(item, index) in state.dashboard.oTAChannelList" :key="index">
                    <div class="data-table-tbody">
                      <div class="data-table-td-d1">{{ index + 1 }}</div>
                      <div class="data-table-td-d2">
                        <n-tooltip trigger="hover">
                          <template #trigger>
                            {{ (item as any)?.name }}
                          </template>
                          {{ (item as any)?.name }}
                        </n-tooltip>
                      </div>
                      <div class="data-table-td-d3">{{ (item as any)?.rate }}%</div>
                    </div>
                  </div>
                </div>
              </n-card>
            </div>
            <div style="width: 20.5%; height: 375px">
              <n-card :bordered="false" style="height: 375px" :header-style="{
                padding: '25px 20px 20px',
              }" :content-style="cardContStyle">
                <template #header>
                  <span class="room-stat-title1">国家预订量排行</span>
                </template>
                <div class="data-table">
                  <div class="data-table-thead">
                    <div class="data-table-td-d1">排名</div>
                    <div class="data-table-td-d2">国家名称</div>
                    <div class="data-table-td-d3">占比</div>
                  </div>
                  <div class="data-table-tbody" v-for="(item, index) in state.dashboard.nationalityList" :key="index">
                    <div class="data-table-td-d1">{{ index + 1 }}</div>
                    <div class="data-table-td-d2">
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          {{ (item as any)?.name }}
                        </template>
                        {{ (item as any)?.name }}
                      </n-tooltip>
                    </div>
                    <div class="data-table-td-d3">{{ (item as any)?.rate }}%</div>
                  </div>
                </div>
              </n-card>
            </div>
          </div>
          <!-- 移动端图表区域 -->
          <div class="flex flex-col gap-4 md:hidden mb-4">
            <!-- 七日订单趋势 -->
            <n-card :bordered="false" :header-style="{
              padding: '16px',
            }" :content-style="{
              padding: '0 16px 16px',
            }">
              <template #header>
                <span class="text-base font-medium text-gray-900">七日订单趋势</span>
              </template>
              <div ref="chartRefMobile" style="width: 100%; height: 250px"></div>
            </n-card>
            <!-- 物业预订量排行 -->
            <n-card :bordered="false" :header-style="{
              padding: '16px',
            }" :content-style="{
              padding: '0 16px 16px',
            }">
              <template #header>
                <span class="text-base font-medium text-gray-900">物业预订量排行</span>
              </template>
              <div class="data-table-mobile">
                <div v-for="(item, index) in state.dashboard.propertyList" :key="index" class="data-table-tbody-mobile">
                  <div class="flex items-center justify-between py-2 border-b border-gray-100">
                    <div class="flex items-center flex-1 min-w-0">
                      <div class="w-8 text-center text-sm font-medium text-gray-700">{{ index + 1 }}</div>
                      <div class="flex-1 min-w-0 text-sm text-gray-900 truncate ml-2">{{ (item as
                        any)?.propertyDetail?.name }}
                      </div>
                    </div>
                    <div class="text-sm font-medium text-gray-700 ml-2">{{ (item as any)?.rate }}%</div>
                  </div>
                </div>
              </div>
            </n-card>
            <!-- OTA渠道排名排行 -->
            <n-card :bordered="false" :header-style="{
              padding: '16px',
            }" :content-style="{
              padding: '0 16px 16px',
            }">
              <template #header>
                <span class="text-base font-medium text-gray-900">OTA渠道排名排行</span>
              </template>
              <div class="data-table-mobile">
                <div v-for="(item, index) in state.dashboard.oTAChannelList" :key="index"
                  class="data-table-tbody-mobile">
                  <div class="flex items-center justify-between py-2 border-b border-gray-100">
                    <div class="flex items-center flex-1 min-w-0">
                      <div class="w-8 text-center text-sm font-medium text-gray-700">{{ index + 1 }}</div>
                      <div class="flex-1 min-w-0 text-sm text-gray-900 truncate ml-2">{{ (item as any)?.name }}</div>
                    </div>
                    <div class="text-sm font-medium text-gray-700 ml-2">{{ (item as any)?.rate }}%</div>
                  </div>
                </div>
              </div>
            </n-card>
            <!-- 国家预订量排行 -->
            <n-card :bordered="false" :header-style="{
              padding: '16px',
            }" :content-style="{
              padding: '0 16px 16px',
            }">
              <template #header>
                <span class="text-base font-medium text-gray-900">国家预订量排行</span>
              </template>
              <div class="data-table-mobile">
                <div v-for="(item, index) in state.dashboard.nationalityList" :key="index"
                  class="data-table-tbody-mobile">
                  <div class="flex items-center justify-between py-2 border-b border-gray-100">
                    <div class="flex items-center flex-1 min-w-0">
                      <div class="w-8 text-center text-sm font-medium text-gray-700">{{ index + 1 }}</div>
                      <div class="flex-1 min-w-0 text-sm text-gray-900 truncate ml-2">{{ (item as any)?.name }}</div>
                    </div>
                    <div class="text-sm font-medium text-gray-700 ml-2">{{ (item as any)?.rate }}%</div>
                  </div>
                </div>
              </div>
            </n-card>
          </div>
        </n-gi>
        <n-gi>
          <!-- PC端近三十日订单趋势 -->
          <n-card :bordered="false" class="hidden md:block" :header-style="{
            padding: '25px 20px 20px',
          }" :content-style="cardContStyle">
            <template #header>
              <span class="room-stat-title1">近三十日订单趋势</span>
            </template>
            <div ref="chartRef30" style="width: 100%; height: 300px"></div>
          </n-card>
          <!-- 移动端近三十日订单趋势 -->
          <n-card :bordered="false" class="block md:hidden" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <span class="text-base font-medium text-gray-900">近三十日订单趋势</span>
            </template>
            <div ref="chartRef30Mobile" style="width: 100%; height: 250px"></div>
          </n-card>
        </n-gi>
      </n-grid>
    </n-spin>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, Ref, ref } from 'vue';
import { useECharts } from '@/hooks/web/useECharts';
import { dashboardAll } from '@/api/comm';
import { hexToRgba } from '@/utils/artDesignUtils';
import { useUserStore } from '@/store/modules/user';
import * as echarts from 'echarts';

const show = ref(false);
const userStore = useUserStore();
const cardContStyle = ref({
  padding: '0 20px 20px',
});
const chartRef = ref<HTMLDivElement | null>(null);
const chartRef30 = ref<HTMLDivElement | null>(null);
const chartRefMobile = ref<HTMLDivElement | null>(null);
const chartRef30Mobile = ref<HTMLDivElement | null>(null);
const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);
const { setOptions: setOptions30 } = useECharts(chartRef30 as Ref<HTMLDivElement>);
const { setOptions: setOptionsMobile } = useECharts(chartRefMobile as Ref<HTMLDivElement>);
const { setOptions: setOptions30Mobile } = useECharts(chartRef30Mobile as Ref<HTMLDivElement>);

const state = reactive({
  orderStatType: 'all',
  dashboard: {
    memberStat: {
      todayRegMemberNum: 0,
      todayMemberNumGrewPer: 0,
      totalMemberNum: 0,
      totalMemberGrewPer: 0,
      todayIncBal: 0,
      todayIncBalGrewPer: 0,
      todayConsumeBal: 0,
      todayConsumeBalGrewPer: 0,
    },
    orderStat: {
      allStat: {
        todayOrderNum: 0,
        todayOrderNumGrewPer: 0,
        totalOrderNum: 0,
        totalOrderNumGrewPer: 0,
        todayRefundNum: 0,
        todayRefundGrewPer: 0,
        totalRefundNum: 0,
        totalRefundGrewPer: 0,
        totalPropertyNum: 0,
        onlinePropertyNum: 0,
        todayOrderMoney: 0,
        todayOrderMoneyGrewPer: 0,
        totalOrderMoney: 0,
        totalOrderMoneyGrewPer: 0,
        todayRefundMoney: 0,
        todayRefundMoneyGrewPer: 0,
        totalRefundMoney: 0,
        totalRefundMoneyGrewPer: 0,
        totalRoomTypeNum: 0,
        onlineRoomTypeNum: 0,
      },
      hotelStat: {
        todayOrderNum: 0,
        todayOrderNumGrewPer: 0,
        totalOrderNum: 0,
        totalOrderNumGrewPer: 0,
        todayRefundNum: 0,
        todayRefundGrewPer: 0,
        totalRefundNum: 0,
        totalRefundGrewPer: 0,
        totalPropertyNum: 0,
        onlinePropertyNum: 0,
        todayOrderMoney: 0,
        todayOrderMoneyGrewPer: 0,
        totalOrderMoney: 0,
        totalOrderMoneyGrewPer: 0,
        todayRefundMoney: 0,
        todayRefundMoneyGrewPer: 0,
        totalRefundMoney: 0,
        totalRefundMoneyGrewPer: 0,
        totalRoomTypeNum: 0,
        onlineRoomTypeNum: 0,
      },
    },
    orderNumberList: [],
    newOrderNumberList: [],
    propertyList: [],
    oTAChannelList: [],
    nationalityList: [],
  },
});

const Load = async () => {
  show.value = true;
  const res = await dashboardAll({});
  state.dashboard = res;
  console.log(state.dashboard);
  let daytimeArr = [];
  let percentArr = [];
  for (let i = 0; i < 7; i++) {
    daytimeArr.push(state.dashboard.orderNumberList[i]['orderDate']);
    percentArr.push(state.dashboard.orderNumberList[i]['orderNumber']);
  }
  setOptions({
    grid: {
      top: '9%',
      bottom: '9%',
      right: '20px',
      left: '60px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 3,
          color: '#5D87FF',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#5D87FF', 0.2).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#5D87FF', 0.01).rgba,
            },
          ]),
        },
      },
    ],
  });

  // 处理30天订单趋势数据
  let daytimeArr30 = [];
  let percentArr30 = [];
  for (let i = 0; i < state.dashboard.newOrderNumberList.length; i++) {
    daytimeArr30.push(state.dashboard.newOrderNumberList[i]['orderDate']);
    percentArr30.push(state.dashboard.newOrderNumberList[i]['orderNumber']);
  }
  setOptions30({
    grid: {
      top: '9%',
      bottom: '9%',
      right: '20px',
      left: '60px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr30,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr30,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 3,
          color: '#5D87FF',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#5D87FF', 0.2).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#5D87FF', 0.01).rgba,
            },
          ]),
        },
      },
    ],
  });

  // 移动端七日订单趋势图表
  setOptionsMobile({
    grid: {
      top: '8%',
      bottom: '12%',
      right: '12px',
      left: '40px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        rotate: 0,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 2,
          color: '#5D87FF',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#5D87FF', 0.2).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#5D87FF', 0.01).rgba,
            },
          ]),
        },
      },
    ],
  });

  // 移动端30天订单趋势图表
  setOptions30Mobile({
    grid: {
      top: '8%',
      bottom: '12%',
      right: '12px',
      left: '40px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr30,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        rotate: 0,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr30,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 2,
          color: '#5D87FF',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#5D87FF', 0.2).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#5D87FF', 0.01).rgba,
            },
          ]),
        },
      },
    ],
  });
  show.value = false;
};

const handleUpdateValue = (value: string) => {
  state.orderStatType = value;
  console.log(value);
};

onMounted(() => {
  Load();
});
</script>

<style scoped lang="less">
.title-name {
  margin: 15px 0 19px 20px;
  font-weight: 500;
  font-size: 30px;
  color: #3d3d3d;
  line-height: 42px;
}

.room-stat-title {
  font-size: 24px;
  color: #3d3d3d;
  font-weight: 500;
  line-height: 34px;
}

.room-stat-title1 {
  font-weight: 500;
  font-size: 18px;
  color: #3d3d3d;
  line-height: 25px;
}

.room-stat-div {
  display: flex;
  align-items: center;
  width: 100%;
  margin-top: 5px;
}

.room-stat-time {
  font-size: 14px;
  color: #8b8b8b;
  font-weight: 400;
  line-height: 20px;
}

.room-stat-refresh {
  font-size: 14px;
  color: #156bff;
  line-height: 20px;
  margin-left: 5px;
}

.room-stat-gi-div {
  border-radius: 2px;
  display: flex;
  align-items: center;
  background: #f9fbff;
  padding: 20px 30px;

  .room-stat-gi-div-item {
    flex: 1;
    margin-left: 30px;

    .room-stat-gi-div-item-d1 {
      font-size: 14px;
      color: #3d3d3d;
      font-weight: 500;
      line-height: 20px;
    }

    .room-stat-gi-div-item-d2 {
      margin-top: 5px;
      font-size: 28px;
      color: #3d3d3d;
      font-weight: 700;
      line-height: 42px;
    }

    .room-stat-gi-div-item-d3 {
      display: flex;
      align-items: center;
      margin-top: 10px;
      font-size: 14px;
      color: #8b8b8b;
      font-weight: 400;
      line-height: 20px;

      img {
        display: block;
      }
    }
  }
}

.tab-div2 {
  padding: 12px 0;

  .tab-div2-d1 {
    font-weight: 500;
    font-size: 14px;
    color: #3d3d3d;
    line-height: 20px;

    span {
      color: #979797;
      font-weight: 400;
    }
  }

  .tab-div2-d2 {
    margin-top: 5px;
    font-weight: 700;
    font-size: 24px;
    color: #3d3d3d;
    line-height: 42px;
  }

  .tab-div2-d3 {
    display: flex;
    align-items: center;
    margin-top: 5px;
    font-weight: 400;
    font-size: 14px;
    color: #8b8b8b;
    line-height: 20px;

    img {
      display: block;
    }
  }
}

.data-table {
  .data-table-thead {
    background: #f2f3f8;
    display: flex;

    div {
      font-weight: 500;
      font-size: 14px;
      color: #1d2129;
      line-height: 40px;
    }

    .data-table-td-d1 {
      width: 17.3%;
      text-align: center;
    }

    .data-table-td-d2 {
      width: 55%;
      padding-left: 7.3%;
    }

    .data-table-td-d3 {
      width: 27.7%;
      text-align: center;
    }
  }

  .data-table-tbody {
    display: flex;
    border-bottom: 1px solid #f2f3f8;

    div {
      font-weight: 400;
      font-size: 14px;
      color: #1d2129;
      line-height: 44px;
    }

    .data-table-td-d1 {
      width: 17.3%;
      text-align: center;
    }

    .data-table-td-d2 {
      width: 55%;
      padding-left: 7.3%;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
    }

    .data-table-td-d3 {
      width: 27.7%;
      text-align: center;
    }
  }
}

// 移动端样式
.room-stat-gi-div-mobile {
  border-radius: 8px;
  background: #f9fbff;
  padding: 12px;
  min-height: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  .room-stat-gi-div-item-mobile {
    width: 100%;
    margin-top: 12px;
  }
}

.tab-div2-mobile {
  padding: 12px;
  background: #f9fbff;
  border-radius: 8px;
  min-height: 100px;
}

.data-table-mobile {
  .data-table-tbody-mobile {
    &:last-child {
      .border-b {
        border-bottom: none !important;
      }
    }
  }
}
</style>
