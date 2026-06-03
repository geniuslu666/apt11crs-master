<template>
  <div class="px-0">
    <n-grid :cols="1" :y-gap="15">
      <!-- PC端会员信息卡片 -->
      <n-gi class="hidden md:block">
        <n-spin :show="show1" description="请稍候...">
          <n-card :bordered="false" :header-style="{
            padding: '25px 20px',
          }" :content-style="{
            padding: '0 20px 20px',
          }">
            <template #header>
              <div class="room-stat-title">会员信息</div>
              <div class="room-stat-div">
                <span class="room-stat-time">在{{ nowTime }}刷新</span>
                <n-button text style="margin-left: 20px" @click="refreshBaseInfo">
                  <img style="width: 16px;" src="@/assets/images/pms_dashboard_refresh_icon.png" />
                  <span class="room-stat-refresh">刷新</span>
                </n-button>
              </div>
            </template>
            <n-grid :cols="4" x-gap="15">
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">会员总数</div>
                      <div class="room-stat-gi-div-item-d2">{{ memberBaseStat.memberTotal }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/member_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日增加会员</div>
                      <div class="room-stat-gi-div-item-d2">{{ memberBaseStat.todayMemberAdd }}</div>
                      <div class="room-stat-gi-div-item-d3">昨日新增：<span>{{ memberBaseStat.yesterdayMemberAdd }}</span>
                      </div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/member_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">会员积分池</div>
                      <div class="room-stat-gi-div-item-d2">{{ memberBaseStat.memberBalanceTotal }}</div>
                      <div class="room-stat-gi-div-item-d3">今日消耗：<span>{{ memberBaseStat.todayConsumedBalance }}</span>
                      </div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/member_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">今日订单数</div>
                      <div class="room-stat-gi-div-item-d2">{{ memberBaseStat.todayHotelOrderNum }}</div>
                      <div class="room-stat-gi-div-item-d3">昨日订单数：<span>{{ memberBaseStat.yesterdayHotelOrderNum
                          }}</span></div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image" style="margin-right: 10px">
                    <MemberOrderPie :memberOrderPercent="memberOrderPie.memberOrderPercent"
                      :memberNotOrderPercent="memberOrderPie.memberNotOrderPercent" />
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>
      <!-- 移动端会员信息卡片 -->
      <n-gi class="block md:hidden">
        <n-spin :show="show1" description="请稍候...">
          <n-card :bordered="false" :header-style="{
            padding: '16px',
          }" :content-style="{
            padding: '0 16px 16px',
          }">
            <template #header>
              <div class="text-base font-medium text-gray-900 mb-2">会员信息</div>
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
                    <div class="text-sm font-medium text-gray-700">会员总数</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ memberBaseStat.memberTotal }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/member_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">今日增加会员</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ memberBaseStat.todayMemberAdd }}</div>
                    <div class="text-xs text-gray-500 mt-2">昨日新增：<span class="text-gray-700 font-medium">{{
                      memberBaseStat.yesterdayMemberAdd }}</span></div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/member_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">会员积分池</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ memberBaseStat.memberBalanceTotal }}</div>
                    <div class="text-xs text-gray-500 mt-2">今日消耗：<span class="text-gray-700 font-medium">{{
                      memberBaseStat.todayConsumedBalance }}</span></div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain opacity-50" src="@/assets/images/member_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">今日订单数</div>
                    <div class="text-xl font-bold text-gray-900 mt-1">{{ memberBaseStat.todayHotelOrderNum }}</div>
                    <div class="text-xs text-gray-500 mt-2">昨日订单数：<span class="text-gray-700 font-medium">{{
                      memberBaseStat.yesterdayHotelOrderNum }}</span></div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <MemberOrderPie :memberOrderPercent="memberOrderPie.memberOrderPercent"
                      :memberNotOrderPercent="memberOrderPie.memberNotOrderPercent" />
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>
      <!-- PC端图表区域 -->
      <n-gi class="hidden md:block">
        <n-grid x-gap="15" :cols="2">
          <n-gi>
            <n-spin :show="show2" description="请稍候...">
              <div class="member-stat-div">
                <div class="member-stat-div-title">新增会员趋势</div>
                <MemberCreateLine :dateList="memberCreateTrendDateList" :numList="memberCreateTrendNumList" />
              </div>
            </n-spin>
          </n-gi>
          <n-gi>
            <n-spin :show="show3" description="请稍候...">
              <div class="member-stat-div">
                <div class="member-stat-div-title">会员消费趋势</div>
                <MemberOrderLine :dateList="memberOrderTrendDateList" :moneyList="memberOrderTrendMoneyList" />
              </div>
            </n-spin>
          </n-gi>
        </n-grid>
      </n-gi>
      <!-- 移动端图表区域 -->
      <n-gi class="block md:hidden">
        <div class="space-y-4">
          <n-spin :show="show2" description="请稍候...">
            <div class="member-stat-div-mobile">
              <div class="member-stat-div-title-mobile">新增会员趋势</div>
              <MemberCreateLine :dateList="memberCreateTrendDateList" :numList="memberCreateTrendNumList" />
            </div>
          </n-spin>
          <n-spin :show="show3" description="请稍候...">
            <div class="member-stat-div-mobile">
              <div class="member-stat-div-title-mobile">会员消费趋势</div>
              <MemberOrderLine :dateList="memberOrderTrendDateList" :moneyList="memberOrderTrendMoneyList" />
            </div>
          </n-spin>
        </div>
      </n-gi>
      <!-- PC端饼图区域 -->
      <n-gi class="hidden md:block">
        <n-grid x-gap="15" :cols="2">
          <n-gi>
            <n-spin :show="show4" description="请稍候...">
              <div class="member-stat-div">
                <div class="member-stat-div-title">会员来源</div>
                <MemberFromPie :percentList="memberSourcePercentList" />
              </div>
            </n-spin>
          </n-gi>
          <n-gi>
            <n-spin :show="show5" description="请稍候...">
              <div class="member-stat-div">
                <div class="member-stat-div-title">会员等级比率</div>
                <MemberLevelPie :percentList="memberLevelPercentList" />
              </div>
            </n-spin>
          </n-gi>
        </n-grid>
      </n-gi>
      <!-- 移动端饼图区域 -->
      <n-gi class="block md:hidden">
        <div class="space-y-4">
          <n-spin :show="show4" description="请稍候...">
            <div class="member-stat-div-mobile">
              <div class="member-stat-div-title-mobile">会员来源</div>
              <MemberFromPie :percentList="memberSourcePercentList" />
            </div>
          </n-spin>
          <n-spin :show="show5" description="请稍候...">
            <div class="member-stat-div-mobile">
              <div class="member-stat-div-title-mobile">会员等级比率</div>
              <MemberLevelPie :percentList="memberLevelPercentList" />
            </div>
          </n-spin>
        </div>
      </n-gi>
    </n-grid>
  </div>
</template>

<script setup lang="ts">
import MemberOrderPie from './components/MemberOrderPie.vue';
import MemberCreateLine from './components/MemberCreateLine.vue';
import MemberOrderLine from './components/MemberOrderLine.vue';
import MemberFromPie from './components/MemberFromPie.vue';
import MemberLevelPie from './components/MemberLevelPie.vue';
import { onMounted, ref } from "vue";
import { Stat } from '@/api/pmsMember';
import { formatToDateTime } from "@/utils/dateUtil";
import { dashboard } from "@/api/comm";

const show1 = ref(false);
const show2 = ref(false);
const show3 = ref(false);
const show4 = ref(false);
const show5 = ref(false);
const today = ref('');
const memberBaseStat = ref({
  memberTotal: 0,
  todayMemberAdd: 0,
  yesterdayMemberAdd: 0,
  memberBalanceTotal: 0,
  todayConsumedBalance: 0,
  todayMemberOrderCount: 0,
  todayHotelOrderNum: 0,
  yesterdayHotelOrderNum: 0,
});
const memberOrderPie = ref({
  memberOrder: 0,
  memberOrderPercent: 0,
  memberNotOrder: 0,
  memberNotOrderPercent: 0,
})
const memberCreateTrendDateList = ref([])
const memberCreateTrendNumList = ref([])
const memberOrderTrendDateList = ref([])
const memberOrderTrendMoneyList = ref([])
const memberSourcePercentList = ref([])
const memberLevelPercentList = ref([])
const nowTime = ref(formatToDateTime(new Date(), 'yyyy-MM-dd hh:mm:ss'));

function refreshBaseInfo() {
  nowTime.value = formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss');
  Load1()
}

const Load1 = () => {
  show1.value = true;
  Stat({
    type: "memberBaseStat"
  }).then((res) => {
    memberBaseStat.value = res.memberBaseStat
    Load1More();
  }).catch((err) => {
    Load1More();
  })
};

const Load1More = () => {
  Stat({
    type: "memberOrderPie"
  }).then((res) => {
    memberOrderPie.value = res.memberOrderPie
    show1.value = false;
  }).catch((err) => {
    show1.value = false;
  })
};

const Load2 = () => {
  show2.value = true;
  Stat({
    type: "memberCreateTrend"
  }).then((res) => {
    memberCreateTrendDateList.value = res.memberCreateTrend.map((item) => {
      return item.date
    })
    memberCreateTrendNumList.value = res.memberCreateTrend.map((item) => {
      return item.num
    })
    show2.value = false;
  }).catch((err) => {
    show2.value = false;
  })
};

const Load3 = () => {
  show3.value = true;
  Stat({
    type: "memberOrderTrend"
  }).then((res) => {
    memberOrderTrendDateList.value = res.memberOrderTrend.map((item) => {
      return item.date
    })
    memberOrderTrendMoneyList.value = res.memberOrderTrend.map((item) => {
      return item.money
    })
    show3.value = false;
  }).catch((err) => {
    show3.value = false;
  })
};

const Load4 = () => {
  show4.value = true;
  Stat({
    type: "memberSource"
  }).then((res) => {
    memberSourcePercentList.value = res.memberSource.map((item) => {
      return {
        value: item.percent,
        name: item.source
      }
    })
    show4.value = false;
  }).catch((err) => {
    show4.value = false;
  })
};

const Load5 = () => {
  show5.value = true;
  Stat({
    type: "memberLevel"
  }).then((res) => {
    memberLevelPercentList.value = res.memberLevel.map((item) => {
      return {
        value: item.percent,
        name: item.levelName
      }
    })
    show5.value = false;
  }).catch((err) => {
    show5.value = false;
  })
};

onMounted(() => {
  Load1();
  Load2();
  Load3();
  Load4();
  Load5();
});
</script>

<style scoped lang="less">
.room-stat-title {
  font-size: 20px;
  color: #3D3D3D;
  font-weight: 500;
  line-height: 28px
}

.room-stat-div {
  display: flex;
  align-items: center;
  width: 100%;
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
        margin-top: 5px;
        font-size: 36px;
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

  .room-stat-gi-div-image {
    width: 140px;
    height: 137px;
  }
}

.member-stat-div {
  background: #FFFFFF;
  border-radius: 4px;
  height: 375px;
  padding: 25px 16px 16px 29px;

  .member-stat-div-title {
    font-weight: 500;
    font-size: 18px;
    color: #3D3D3D;
    line-height: 25px;
    margin-bottom: 15px;
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

.member-stat-div-mobile {
  background: #FFFFFF;
  border-radius: 12px;
  padding: 16px;
  border: 1px solid #f0f0f0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column;
  height: 320px;

  .member-stat-div-title-mobile {
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 24px;
    margin-bottom: 12px;
    flex-shrink: 0;
  }

  // 确保图表容器可以自适应
  >*:not(.member-stat-div-title-mobile) {
    flex: 1;
    min-height: 0;
    overflow: hidden;
  }
}
</style>
