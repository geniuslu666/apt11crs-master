<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
        padding: '20px',
      }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">优惠券详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              优惠券名称
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">优惠券名称-中文</div>
                  <div class="level-detail-div-item-content">{{ info.zh_name ? info.zh_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">优惠券名称-日本语</div>
                  <div class="level-detail-div-item-content">{{ info.ja_name ? info.ja_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">优惠券名称-繁体中文</div>
                  <div class="level-detail-div-item-content">{{ info.zh_CN_name ? info.zh_CN_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">优惠券名称-English</div>
                  <div class="level-detail-div-item-content">{{ info.en_name ? info.en_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">优惠券名称-한국어</div>
                  <div class="level-detail-div-item-content">{{ info.ko_name ? info.ko_name : '--' }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              基本信息
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">优惠券类型</div>
                  <div class="level-detail-div-item-content">
                    <div class="level-detail-div-item-content-tag success" v-if="info.type == 'reward'">满减</div>
                    <div class="level-detail-div-item-content-tag warning" v-else>折扣</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">使用场景</div>
                  <div class="level-detail-div-item-content">
                    <div class="level-detail-div-item-content-tag1 success" v-if="info.scene == 1">民宿</div>
                    <div class="level-detail-div-item-content-tag1 primary" v-else-if="info.scene == 2">订餐</div>
                    <div class="level-detail-div-item-content-tag1 warning" v-else-if="info.scene == 3">按摩</div>
                    <div class="level-detail-div-item-content-tag1 error" v-else-if="info.scene == 4">接送机/包车</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">{{ info.type == 'reward' ? '优惠面额' : '优惠折扣' }}</div>
                  <div class="level-detail-div-item-content">{{ info.type == 'reward' ? info.money + 'JPY' :
                    info.discount + '折'
                    }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">最大领取数量</div>
                  <div class="level-detail-div-item-content">{{ info.maxFetch }}张</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">使用门槛</div>
                  <div class="level-detail-div-item-content">{{ info.atLeast }}JPY</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">有效期</div>
                  <div class="level-detail-div-item-content">{{ info.validityType == 1 ? info.endUseTime :
                    (info.validityType ==
                      2 ? '领取后'+info.fixedTerm+'天有效' : '长期有效') }}</div>
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item" v-if="info.scene == 1">
                  <div class="level-detail-div-item-title">适用物业</div>
                  <div class="level-detail-div-item-content">{{ info.propertyNames }}</div>
                </div>
                <div class="level-detail-div-item" v-else-if="info.scene == 2">
                  <div class="level-detail-div-item-title">适用门店</div>
                  <div class="level-detail-div-item-content">{{ info.restaurantNames }}</div>
                </div>
                <div class="level-detail-div-item" v-else-if="info.scene == 3">
                  <div class="level-detail-div-item-title">适用服务类型</div>
                  <div class="level-detail-div-item-content">{{ info.spaServiceTypeStr }}</div>
                </div>
                <div class="level-detail-div-item" v-else-if="info.scene == 4">
                  <div class="level-detail-div-item-title">适用服务类型</div>
                  <div class="level-detail-div-item-content">{{ info.carServiceTypeStr }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              使用说明
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">简体中文</div>
                  <div class="level-detail-div-item-content">{{ info.zh_desc ? info.zh_desc : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">日本语</div>
                  <div class="level-detail-div-item-content">{{ info.ja_desc ? info.ja_desc : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">English</div>
                  <div class="level-detail-div-item-content">{{ info.en_desc ? info.en_desc : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">한국어</div>
                  <div class="level-detail-div-item-content">{{ info.ko_desc ? info.ko_desc : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">繁体中文</div>
                  <div class="level-detail-div-item-content">{{ info.zh_CN_desc ? info.zh_CN_desc : '--' }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              数据统计
            </div>
            <div class="coupon-statist-div">
              <div class="coupon-statist-div-item">
                <div>发放数</div>
                <div>{{ info.count }}</div>
              </div>
              <div class="coupon-statist-div-item">
                <div>领取数</div>
                <div>{{ info.leadCount }}</div>
              </div>
              <div class="coupon-statist-div-item">
                <div>发放数</div>
                <div>{{ info.usedCount }}</div>
              </div>
            </div>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title" style="padding: 0 20px;">
              <div></div>
              领取记录
            </div>
            <div class="car-order-detail-tabs">
              <div class="car-order-detail-tabs-item" @click="handleCouponState(0)"
                :class="searchState == 0 ? 'active' : ''">
                全部
                <div></div>
              </div>
              <div class="car-order-detail-tabs-item" @click="handleCouponState(1)"
                :class="searchState == 1 ? 'active' : ''">
                已领取
                <div></div>
              </div>
              <div class="car-order-detail-tabs-item" @click="handleCouponState(2)"
                :class="searchState == 2 ? 'active' : ''">
                已使用
                <div></div>
              </div>
              <div class="car-order-detail-tabs-item" @click="handleCouponState(3)"
                :class="searchState == 3 ? 'active' : ''">
                已过期
                <div></div>
              </div>
            </div>
            <div style="padding: 0 20px;">
              <BasicTable :bordered="false" ref="actionCouponRef" :columns="couponColumns"
                :request="loadCouponDataTable" :scroll-x="scrollCouponX" :resizeHeightOffset="-10000">
              </BasicTable>
            </div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { useMessage } from "naive-ui";
import { View } from "@/api/pmsCouponType";
import { BasicTable } from "@/components/Table";
import { couponColumns, loadCouponOptions } from './view_coupon_model';
import { adaModalWidth, adaTableScrollX } from "@/utils/hotgo";
import { List as CouponList } from "@/api/pmsCoupon";
import { jsontoobj } from "@/utils/smjcomm";

const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});

const couponTypeId = ref(0);
const message = useMessage();
const router = useRouter();
const info = ref({});
const searchState = ref(0);

const actionCouponRef = ref();
const scrollCouponX = computed(() => {
  return adaTableScrollX(couponColumns, 0);
});

const getInfo = (id) => {
  loading.value = true;
  View({ id })
    .then((res) => {
      res.nameLanguage = jsontoobj(res.nameLanguage);
      res.descLanguage = jsontoobj(res.descLanguage);



      info.value = res;
      console.log('res', info.value.carServiceTypeStr);
      if (res.scene == 1) {
        info.value.sceneName = '住宿'
      } else if (res.scene == 2) {
        info.value.sceneName = '订餐'
      } else if (res.scene == 3) {
        info.value.sceneName = '按摩'
      } else if (res.scene == 4) {
        info.value.sceneName = '接送机/包车'
      } else {
        info.value.sceneName = '未知'
      }

      info.value.zh_name = res.nameLanguage.zh && res.nameLanguage.zh.content ? res.nameLanguage.zh.content : ''
      info.value.en_name = res.nameLanguage.en && res.nameLanguage.en.content ? res.nameLanguage.en.content : ''
      info.value.ko_name = res.nameLanguage.ko && res.nameLanguage.ko.content ? res.nameLanguage.ko.content : ''
      info.value.ja_name = res.nameLanguage.ja && res.nameLanguage.ja.content ? res.nameLanguage.ja.content : ''
      if (res.nameLanguage.zh_CN) {
        info.value.zh_CN_name = res.nameLanguage.zh_CN.content ? res.nameLanguage.zh_CN.content : ''
      } else if (res.nameLanguage.zh_cn) {
        info.value.zh_CN_name = res.nameLanguage.zh_cn.content ? res.nameLanguage.zh_cn.content : ''
      } else {
        info.value.zh_CN_name = ''
      }

      info.value.zh_desc = res.descLanguage.zh && res.descLanguage.zh.content ? res.descLanguage.zh.content : ''
      info.value.en_desc = res.descLanguage.en && res.descLanguage.en.content ? res.descLanguage.en.content : ''
      info.value.ko_desc = res.descLanguage.ko && res.descLanguage.ko.content ? res.descLanguage.ko.content : ''
      info.value.ja_desc = res.descLanguage.ja && res.descLanguage.ja.content ? res.descLanguage.ja.content : ''
      if (res.descLanguage.zh_CN) {
        info.value.zh_CN_desc = res.descLanguage.zh_CN.content ? res.descLanguage.zh_CN.content : ''
      } else if (res.descLanguage.zh_cn) {
        info.value.zh_CN_desc = res.descLanguage.zh_cn.content ? res.descLanguage.zh_cn.content : ''
      } else {
        info.value.zh_CN_desc = ''
      }
    })
    .finally(() => {
      loading.value = false;
    });
};

// 加载表格数据
const loadCouponDataTable = async (res) => {
  res.couponTypeId = couponTypeId.value
  if (searchState.value > 0) {
    res.state = searchState.value
  }
  return await CouponList({ ...res });
};

// 重新加载表格数据
function reloadCouponTable() {
  actionCouponRef.value?.reload();
}

function handleCouponState(e) {
  searchState.value = e;
  reloadCouponTable()
}

async function openModal(id) {
  couponTypeId.value = id
  showModal.value = true;
  loading.value = true;
  loadCouponOptions();
  getInfo(id);
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}

.level-detail-div {
  margin-bottom: 25px;

  &-title {
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;

    div {
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }

  &-item {
    &-title {
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
    }

    &-content {
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;

      &-tag {
        width: 38px;
        height: 22px;
        line-height: 22px;
        text-align: center;
        font-size: 14px;
        border-radius: 2px;
        font-weight: 400;

        &.success {
          color: #17A158;
          border: 1px solid #17A158;
        }

        &.warning {
          color: #F48720;
          border: 1px solid #F48720;
        }
      }

      &-tag1 {
        padding: 0 5px;
        height: 22px;
        line-height: 22px;
        text-align: center;
        font-size: 14px;
        border-radius: 2px;
        font-weight: 400;

        &.success {
          color: #26A763;
          background: #E3F4EB;
        }

        &.warning {
          color: #EFA020;
          background: #FFECCE;
        }

        &.primary {
          color: #3F9EFF;
          background: #ECF5FF;
        }

        &.error {
          color: #F56C6C;
          background: #FEF0F0;
        }
      }
    }
  }
}

.coupon-statist-div {
  display: flex;
  align-items: center;
  justify-content: space-between;

  &-item {
    width: 340px;
    height: 88px;
    background: #F9FBFC;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 0 15px;

    div {
      &:first-child {
        font-weight: 400;
        font-size: 14px;
        color: #707070;
        line-height: 20px;
      }

      &:last-child {
        margin-top: 5px;
        font-weight: 500;
        font-size: 24px;
        color: #3D3D3D;
        line-height: 33px;
      }
    }
  }
}

.car-order-detail-tabs {
  background: #F9FBFC;
  height: 40px;
  line-height: 40px;
  padding-left: 20px;
  display: flex;

  &-item {
    height: 40px;
    line-height: 40px;
    position: relative;
    font-weight: 500;
    font-size: 14px;
    color: #3D3D3D;
    margin-right: 36px;
    cursor: pointer;

    &.active {
      color: #053DC8;

      div {
        position: absolute;
        left: 0;
        right: 0;
        bottom: -1px;
        background: #053DC8;
        height: 2px;
      }
    }
  }
}
</style>
