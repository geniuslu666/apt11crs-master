<template>
  <div>
    <n-grid :cols="1" y-gap="15">
      <!-- PC端 实时概况 -->
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
            <n-grid :cols="4" y-gap="15" x-gap="15">
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">餐厅总数</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.restaurantTotalNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">营业中餐厅数量</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.restaurantTotalOnNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">累计餐厅预约量</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.restaurantTotalOrderNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">累计餐厅预约金额</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.restaurantTotalOrderAmount }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon4.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">待结算订单量</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.waitSettlementOrderNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon5.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">待结算金额</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.waitSettlementOrderAmount }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon6.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">待核账订单量</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.waitVerifyOrderNum }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon7.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div">
                  <div class="room-stat-gi-div-abso">
                    <div class="room-stat-gi-div-item">
                      <div class="room-stat-gi-div-item-d1">待核账金额</div>
                      <div class="room-stat-gi-div-item-d2">{{ baseInfo.waitVerifyOrderAmount }}</div>
                    </div>
                  </div>
                  <div class="room-stat-gi-div-image">
                    <img style="width: 100%;" src="@/assets/images/food_stat_icon8.png" />
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
                    <div class="text-sm font-medium text-gray-700">餐厅总数</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.restaurantTotalNum }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon1.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">营业中餐厅数量</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.restaurantTotalOnNum }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon2.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">累计餐厅预约量</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.restaurantTotalOrderNum }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon3.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">累计餐厅预约金额</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.restaurantTotalOrderAmount }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon4.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">待结算订单量</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.waitSettlementOrderNum }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon5.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">待结算金额</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.waitSettlementOrderAmount }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon6.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">待核账订单量</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.waitVerifyOrderNum }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon7.png" />
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="room-stat-gi-div-mobile">
                  <div class="room-stat-gi-div-item-mobile">
                    <div class="text-sm font-medium text-gray-700">待核账金额</div>
                    <div class="text-lg font-bold text-gray-900 mt-1">{{ baseInfo.waitVerifyOrderAmount }}</div>
                  </div>
                  <div class="room-stat-gi-div-image-mobile">
                    <img class="w-full h-full object-contain" src="@/assets/images/food_stat_icon8.png" />
                  </div>
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
            <n-spin :show="show2" description="请稍候...">
              <div class="food-stat-div">
                <div class="food-stat-div-top">
                  <div class="food-stat-div-top-title">餐厅排行</div>
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
                    <div class="data-table-td-d2">餐厅名称</div>
                    <div class="data-table-td-d3">{{ rankTab1 == 'order_num' ? '预定量' : '预定金额' }}</div>
                  </div>
                  <div class="data-table-tbody" v-for="(item, index) in restaurantLine" :key="index">
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
            <n-spin :show="show3" description="请稍候...">
              <div class="food-stat-div">
                <div class="food-stat-div-top">
                  <div class="food-stat-div-top-title">套餐排行</div>
                  <div class="food-stat-div-top-tab">
                    <div @click="handleUpdateValue2('order_num')" :class="rankTab2 == 'order_num' ? 'active' : ''">预定量
                    </div>
                    <div @click="handleUpdateValue2('order_amount')"
                      :class="rankTab2 == 'order_amount' ? 'active' : ''">预定金额
                    </div>
                  </div>
                </div>
                <div class="data-table">
                  <div class="data-table-thead">
                    <div class="data-table-td-d1">排名</div>
                    <div class="data-table-td-d2" style="width: 36.3%;padding-left: 1.5%">套餐名称</div>
                    <div class="data-table-td-d3" style="width: 23.8%;padding-left: 1.5%">餐厅名称</div>
                    <div class="data-table-td-d4" style="width: 24.5%;padding-left: 1.5%">{{ rankTab2 == 'order_num' ?
                      '预定量' :
                      '预定金额' }}</div>
                  </div>
                  <div class="data-table-tbody" v-for="(item, index) in goodsLine" :key="index">
                    <div class="data-table-td-d1">{{ index + 1 }}</div>
                    <div class="data-table-td-d2" style="width: 36.3%;padding-left: 1.5%">
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          {{ item.goodsName }}
                        </template>
                        {{ item.goodsName }}
                      </n-tooltip>
                    </div>
                    <div class="data-table-td-d3" style="width: 23.8%;padding-left: 1.5%">
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          {{ item.restaurantDetail.name }}
                        </template>
                        {{ item.restaurantDetail.name }}
                      </n-tooltip>
                    </div>
                    <div class="data-table-td-d4" style="width: 24.5%;padding-left: 1.5%">{{ rankTab2 == 'order_num' ?
                      item.totalOrderNum : item.totalOrderAmount }}</div>
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
          <!-- 餐厅排行 -->
          <n-spin :show="show2" description="请稍候...">
            <div class="food-stat-div-mobile">
              <div class="food-stat-div-top-mobile">
                <div class="text-base font-medium text-gray-900 mb-3">餐厅排行</div>
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
                  <div class="flex-1 text-xs font-medium text-gray-700 ml-2">餐厅名称</div>
                  <div class="w-20 text-right text-xs font-medium text-gray-700">
                    {{ rankTab1 == 'order_num' ? '预定量' : '预定金额' }}
                  </div>
                </div>
                <div class="data-table-tbody-mobile" v-for="(item, index) in restaurantLine" :key="index">
                  <div class="w-12 text-center text-sm text-gray-700">{{ index + 1 }}</div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate">
                    {{ item.name }}
                  </div>
                  <div class="w-20 text-right text-sm text-gray-700">
                    {{ rankTab1 == 'order_num' ? item.totalOrderNum : item.totalOrderAmount }}
                  </div>
                </div>
              </div>
            </div>
          </n-spin>

          <!-- 套餐排行 -->
          <n-spin :show="show3" description="请稍候...">
            <div class="food-stat-div-mobile">
              <div class="food-stat-div-top-mobile">
                <div class="text-base font-medium text-gray-900 mb-3">套餐排行</div>
                <div class="food-stat-div-top-tab-mobile">
                  <div @click="handleUpdateValue2('order_num')" :class="rankTab2 == 'order_num' ? 'active' : ''">预定量
                  </div>
                  <div @click="handleUpdateValue2('order_amount')" :class="rankTab2 == 'order_amount' ? 'active' : ''">
                    预定金额
                  </div>
                </div>
              </div>
              <div class="data-table-mobile">
                <div class="data-table-thead-mobile">
                  <div class="w-12 text-center text-xs font-medium text-gray-700">排名</div>
                  <div class="flex-1 text-xs font-medium text-gray-700 ml-2">套餐名称</div>
                  <div class="w-24 text-right text-xs font-medium text-gray-700">
                    {{ rankTab2 == 'order_num' ? '预定量' : '预定金额' }}
                  </div>
                </div>
                <div class="data-table-tbody-mobile" v-for="(item, index) in goodsLine" :key="index">
                  <div class="w-12 text-center text-sm text-gray-700">{{ index + 1 }}</div>
                  <div class="flex-1 text-sm text-gray-900 ml-2 truncate">
                    {{ item.goodsName }}（{{ item.restaurantDetail.name }}）
                  </div>
                  <div class="w-24 text-right text-sm text-gray-700">
                    {{ rankTab2 == 'order_num' ? item.totalOrderNum : item.totalOrderAmount }}
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
import { dashboard } from "@/api/foodDashboard";

const nowTime = ref(formatToDateTime(new Date(), 'yyyy-MM-dd HH:mm:ss'));
const show1 = ref(false);
const show2 = ref(false);
const show3 = ref(false);
const rankTab1 = ref('order_num');
const rankTab2 = ref('order_num');

const baseInfo = ref({
  restaurantTotalNum: 0, //餐厅总数
  restaurantTotalOnNum: 0, //营业中餐厅数量
  restaurantTotalOrderNum: 0, //累计餐厅预约量
  restaurantTotalOrderAmount: 0, //累计餐厅预约金额
  waitSettlementOrderNum: 0, //待结算订单量
  waitSettlementOrderAmount: 0, //待结算金额
  waitVerifyOrderNum: 0, //待核账订单量
  waitVerifyOrderAmount: 0,  //待核账金额
})

const restaurantLine = ref([])
const goodsLine = ref([])

function handleUpdateValue1(value) {
  rankTab1.value = value
  Load2()
}
function handleUpdateValue2(value) {
  rankTab2.value = value
  Load3()
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
    type: 'restaurantLine',
    restaurant_line_type: rankTab1.value == 'order_num' ? 'num' : 'amount'
  })
    .then((res) => {
      restaurantLine.value = res.restaurantLine
      show2.value = false;
    }).catch((err) => {
      show2.value = false;
    })
};

const Load3 = () => {
  show3.value = true;
  dashboard({
    type: 'goodsLine',
    goods_line_type: rankTab2.value == 'order_num' ? 'num' : 'amount'
  })
    .then((res) => {
      goodsLine.value = res.goodsLine
      show3.value = false;
    }).catch((err) => {
      show3.value = false;
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
      margin-top: 5px;
      font-size: 28px;
      color: #3D3D3D;
      font-weight: 700;
      line-height: 42px;
    }
  }

  .room-stat-gi-div-image {
    width: 140px;
    height: 137px;
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
