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
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.todayOrderTotalNum }}</div>
                      <div class="room-stat-gi-div-item-d3">有效：{{ baseInfo.todayEffectiveOrderNum }}</div>
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
                      <div class="room-stat-gi-div-item-d1">今日营业额</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.todayOrderAmount }}</div>
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
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.totalOrderNum }}</div>
                      <div class="room-stat-gi-div-item-d3">有效：{{ baseInfo.effectiveOrderNum }}</div>
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
                      <div class="room-stat-gi-div-item-d1">今日待核销</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.todayWaitVerifyNum }}</div>
                      <div class="room-stat-gi-div-item-d3">已核销：{{ baseInfo.todayVerifiedNum }}</div>
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
                      {{ baseInfo.todayOrderTotalNum }}
                    </div>
                    <div class="text-xs text-gray-500 mt-2">
                      有效：{{ baseInfo.todayEffectiveOrderNum }}
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
                    <div class="text-sm font-medium text-gray-700">今日营业额</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">
                      {{ baseInfo.todayOrderAmount }}
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
                      {{ baseInfo.totalOrderNum }}
                    </div>
                    <div class="text-xs text-gray-500 mt-2">
                      有效：{{ baseInfo.effectiveOrderNum }}
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
                    <div class="text-sm font-medium text-gray-700">今日待核销</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">
                      {{ baseInfo.todayWaitVerifyNum }}
                    </div>
                    <div class="text-xs text-gray-500 mt-2">
                      已核销：{{ baseInfo.todayVerifiedNum }}
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
              <div class="room-stat-title">今日订单</div>
            </template>
            <n-grid :cols="4" x-gap="28">
              <n-gi>
                <div class="today-stat-div">
                  <div>待支付订单</div>
                  <div>{{ todayInfo.waitPayOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div">
                  <div>待核销订单</div>
                  <div>{{ todayInfo.waitVerifyOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div">
                  <div>已完成订单</div>
                  <div>{{ todayInfo.doneOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div">
                  <div>已退款订单</div>
                  <div>{{ todayInfo.refundOrder }}</div>
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
                  <div class="text-sm font-medium text-gray-700">待支付订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.waitPayOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">待核销订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.waitVerifyOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">已完成订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.doneOrder }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="today-stat-div-mobile">
                  <div class="text-sm font-medium text-gray-700">已退款订单</div>
                  <div class="text-xl font-bold text-gray-900 mt-1">{{ todayInfo.refundOrder }}</div>
                </div>
              </n-gi>
            </n-grid>
          </n-card>
        </n-spin>
      </n-gi>

    </n-grid>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { formatToDateTime } from '@/utils/dateUtil';
import { dashboard } from '@/api/travelDashboard';

const nowTime = ref(formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss'));
const show1 = ref(false);
const show2 = ref(false);
const show3 = ref(false);
const rankTab = ref('order_num');

const baseInfo = ref({
  todayOrderTotalNum: 0,   // 今日预约总数
  todayEffectiveOrderNum: 0, // 今日有效预约数（非取消非退款）
  todayOrderAmount: 0,     // 今日营业额
  totalOrderAmount: 0,     // 累计营业额
  totalOrderNum: 0,        // 累计预约数
  effectiveOrderNum: 0,    // 累计有效预约数
  todayWaitVerifyNum: 0,   // 今日待核销数
  todayVerifiedNum: 0,     // 今日已核销数
});

const todayInfo = ref({
  waitPayOrder: 0,     // 待支付订单数
  waitVerifyOrder: 0,  // 待核销订单数
  doneOrder: 0,        // 已完成订单数
  refundOrder: 0,      // 已退款订单数
});

const loadBaseInfo = () => {
  show1.value = true;
  dashboard({ type: 'base' })
    .then((res) => {
      baseInfo.value = res.details;
      show1.value = false;
    })
    .catch(() => {
      show1.value = false;
    });
};

const loadTodayOrder = () => {
  show2.value = true;
  dashboard({ type: 'todayOrder' })
    .then((res) => {
      todayInfo.value = res.todayOrder;
      show2.value = false;
    })
    .catch(() => {
      show2.value = false;
    });
};

function refreshBaseInfo() {
  nowTime.value = formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss');
  loadBaseInfo();
}

onMounted(() => {
  loadBaseInfo();
  loadTodayOrder();
});
</script>

<style scoped lang="less">
.room-stat-title {
  font-size: 20px;
  color: #3D3D3D;
  font-weight: 500;
  line-height: 28px;
}

.room-stat-div {
  display: flex;
  align-items: center;
  margin-top: 11px;
}

.room-stat-time {
  font-size: 14px;
  color: #8B8B8B;
  font-weight: 400;
  line-height: 20px;
}

.room-stat-refresh {
  font-size: 14px;
  color: #156BFF;
  line-height: 20px;
  margin-left: 5px;
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
  min-height: 200px;

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
