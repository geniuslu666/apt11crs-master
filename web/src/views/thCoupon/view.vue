<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
        padding: '20px',
      }" :body-content-style="{
        padding: '25px 0 0',
      }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">礼品券详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              礼品券名称
            </div>
            <n-grid y-gap="25" :cols="1">
              <n-gi>
                <n-grid y-gap="25" cols="1 600:3">
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券名称-中文</div>
                      <div class="level-detail-div-item-content">{{ info.zh_name ? info.zh_name : '--' }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券名称-日本语</div>
                      <div class="level-detail-div-item-content">{{ info.ja_name ? info.ja_name : '--' }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券名称-繁体中文</div>
                      <div class="level-detail-div-item-content">{{ info.zh_CN_name ? info.zh_CN_name : '--' }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券名称-English</div>
                      <div class="level-detail-div-item-content">{{ info.en_name ? info.en_name : '--' }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券名称-한국어</div>
                      <div class="level-detail-div-item-content">{{ info.ko_name ? info.ko_name : '--' }}</div>
                    </div>
                  </n-gi>
                </n-grid>
              </n-gi>
              <n-gi>
                <n-grid y-gap="25" cols="1 600:3">
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券副标题-中文</div>
                      <div class="level-detail-div-item-content">{{ info.zh_sub_name ? info.zh_sub_name : '--' }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券副标题-日本语</div>
                      <div class="level-detail-div-item-content">{{ info.ja_sub_name ? info.ja_sub_name : '--' }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券副标题-繁体中文</div>
                      <div class="level-detail-div-item-content">{{ info.zh_CN_sub_name ? info.zh_CN_sub_name : '--' }}
                      </div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券副标题-English</div>
                      <div class="level-detail-div-item-content">{{ info.en_sub_name ? info.en_sub_name : '--' }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">礼品券副标题-한국어</div>
                      <div class="level-detail-div-item-content">{{ info.ko_sub_name ? info.ko_sub_name : '--' }}</div>
                    </div>
                  </n-gi>
                </n-grid>
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
                  <div class="level-detail-div-item-title">礼品券前缀编号</div>
                  <div class="level-detail-div-item-content">{{ info.couponNoPrefix }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">礼品券分类</div>
                  <div class="level-detail-div-item-content">{{ info.categoryInfo?.name }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">礼品券使用模式</div>
                  <div class="level-detail-div-item-content">{{ info.useMode == 'ARRIVE_VERIFY' ? '到店核销' : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">礼品券发行状态</div>
                  <div class="level-detail-div-item-content">
                    <n-tag :type="getOptionTag(options.th_coupon_status, info.status)" size="small"
                      class="min-left-space">
                      {{ getOptionLabel(options.th_coupon_status, info.status) }}
                    </n-tag>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">礼品券使用状态</div>
                  <div class="level-detail-div-item-content">
                    <n-tag :type="getOptionTag(options.th_coupon_use_status, info.useStatus)" size="small"
                      class="min-left-space">
                      {{ getOptionLabel(options.th_coupon_use_status, info.useStatus) }}
                    </n-tag>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">礼品券创建时间</div>
                  <div class="level-detail-div-item-content">{{ info.createAt }}</div>
                </div>
              </n-gi>
              <n-gi span="3" v-if="info.mchList">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">适用商户</div>
                  <div class="level-detail-div-item-content">
                    <n-table style="margin-bottom: 25px">
                      <thead>
                        <tr>
                          <th>商户名称</th>
                          <th>商户分类</th>
                          <th>商户状态</th>
                          <th>商户核销商品名</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="(item, index) of info.mchList" :key="index">
                          <td>{{ item.mchInfo?.name }}</td>
                          <td>{{ item.mchInfo?.categoryInfo?.name }}</td>
                          <td>
                            <n-tag :type="getOptionTag(mchOptions.sys_normal_disable, item.mchInfo?.status)"
                              size="small" class="min-left-space">
                              {{ getOptionLabel(mchOptions.sys_normal_disable, item.mchInfo?.status) }}
                            </n-tag>
                          </td>
                          <td>
                            {{ item.name }}
                          </td>
                        </tr>
                      </tbody>
                    </n-table>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">有效期</div>
                  <div class="level-detail-div-item-content">{{ '激活后' + info.fixedTerm + '天有效' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">需预约餐厅</div>
                  <div class="level-detail-div-item-content">{{ info.needReservation == 1 ? info.restaurantNames :
                    '无需预约' }}
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              使用说明
            </div>
            <n-grid y-gap="25" :cols="1">
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
                <div>已发放</div>
                <div>{{ info.count }}</div>
              </div>
              <div class="coupon-statist-div-item">
                <div>已核销</div>
                <div>{{ info.usedCount }}</div>
              </div>
              <div class="coupon-statist-div-item">
                <div>未使用</div>
                <div>{{ info.count - info.usedCount }}</div>
              </div>
            </div>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title" style="padding: 0 20px;margin-bottom: 8px">
              <div></div>
              核销记录
            </div>
            <div style="padding: 0 20px;">
              <BasicTable :bordered="false" ref="actionCouponRef" :columns="couponColumns"
                :request="loadCouponDataTable" :scroll-x="scrollCouponX" :resizeHeightOffset="-10000">
                <template #tableTitle>
                  <n-button type="primary" @click="handleExport" class="min-left-space"
                    v-if="hasPermission(['/thMemberCoupon/export'])">
                    <template #icon>
                      <n-icon>
                        <ExportOutlined />
                      </n-icon>
                    </template>
                    导出
                  </n-button>
                </template>
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
import { View } from "@/api/thCoupon";
import { BasicTable } from "@/components/Table";
import { options, loadOptions } from './model'
import { couponColumns } from './view_coupon_model';
import { adaModalWidth, adaTableScrollX, getOptionLabel, getOptionTag } from "@/utils/hotgo";
import { Export, List as CouponList } from "@/api/thMemberCoupon";
import { jsontoobj } from "@/utils/smjcomm";
import {
  options as mchOptions, loadDictOptions as mchLoadDictOptions
} from "@/views/thMch/model";
import { ExportOutlined } from "@vicons/antd";
import { usePermission } from "@/hooks/web/usePermission";

const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});

const couponId = ref(0);
const { hasPermission } = usePermission();
const message = useMessage();
const router = useRouter();
const info = ref({});

const actionCouponRef = ref();
const scrollCouponX = computed(() => {
  return adaTableScrollX(couponColumns, 0);
});

const getInfo = (id) => {
  loading.value = true;
  View({ id })
    .then((res) => {
      res.nameLanguage = jsontoobj(res.nameLanguage);
      res.subNameLanguage = jsontoobj(res.subNameLanguage);
      res.descLanguage = jsontoobj(res.descLanguage);



      info.value = res;

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

      info.value.zh_sub_name = res.subNameLanguage.zh && res.subNameLanguage.zh.content ? res.subNameLanguage.zh.content : ''
      info.value.en_sub_name = res.subNameLanguage.en && res.subNameLanguage.en.content ? res.subNameLanguage.en.content : ''
      info.value.ko_sub_name = res.subNameLanguage.ko && res.subNameLanguage.ko.content ? res.subNameLanguage.ko.content : ''
      info.value.ja_sub_name = res.subNameLanguage.ja && res.subNameLanguage.ja.content ? res.subNameLanguage.ja.content : ''
      if (res.nameLanguage.zh_CN) {
        info.value.zh_CN_sub_name = res.subNameLanguage.zh_CN.content ? res.subNameLanguage.zh_CN.content : ''
      } else if (res.nameLanguage.zh_cn) {
        info.value.zh_CN_sub_name = res.subNameLanguage.zh_cn.content ? res.subNameLanguage.zh_cn.content : ''
      } else {
        info.value.zh_CN_sub_name = ''
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
  res.couponId = couponId.value
  res.state = 3
  res.isVerifyTimeDesc = true
  return await CouponList({ ...res });
};

// 重新加载表格数据
function reloadCouponTable() {
  actionCouponRef.value?.reload();
}

async function openModal(id) {
  couponId.value = id
  showModal.value = true;
  loading.value = true;
  loadOptions();
  mchLoadDictOptions();
  getInfo(id);
}

defineExpose({
  openModal,
});

function handleExport() {
  message.loading('正在导出列表...', { duration: 1200 });
  Export({ couponId: couponId.value, state: 3 })
}
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
