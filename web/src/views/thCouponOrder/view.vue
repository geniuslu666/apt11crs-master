<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
        padding: '20px',
      }" :body-content-style="{
        padding: '25px 0 0',
      }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">礼品券订单详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              会员礼品券信息
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">券号</div>
                  <div class="level-detail-div-item-content">{{ info.couponNo }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">核销状态</div>
                  <div class="level-detail-div-item-content">
                    <n-tag size="small" :bordered="false" type="success" v-if="info.state == 1">
                      待生效
                    </n-tag>
                    <n-tag size="small" :bordered="false" type="info" v-if="info.state == 2">
                      待使用
                    </n-tag>
                    <n-tag size="small" :bordered="false" v-if="info.state == 3">
                      已核销
                    </n-tag>
                    <n-tag size="small" :bordered="false" type="warning" v-if="info.state == 4">
                      已过期
                    </n-tag>
                    <n-tag size="small" :bordered="false" type="error" v-if="info.state == 5">
                      已失效
                    </n-tag>
                    <n-tag size="small" :bordered="false" type="error" v-if="info.state == 6">
                      已回收
                    </n-tag>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">发放时间</div>
                  <div class="level-detail-div-item-content">{{ info.createAt }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">有效期开始时间</div>
                  <div class="level-detail-div-item-content">{{ info.startTime }}</div>
                </div>
              </n-gi>

              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">有效期结束时间</div>
                  <div class="level-detail-div-item-content">{{ info.endTime }}</div>
                </div>
              </n-gi>

              <n-gi v-if="info.state == 5">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">失效时间</div>
                  <div class="level-detail-div-item-content">{{ info.invalidTime }}</div>
                </div>
              </n-gi>

              <n-gi v-if="info.state == 6">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">回收时间</div>
                  <div class="level-detail-div-item-content">{{ info.recoveryTime }}</div>
                </div>
              </n-gi>

            </n-grid>
          </div>

          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              领取会员信息
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员号</div>
                  <div class="level-detail-div-item-content">{{ info.member ? info.member.memberNo : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员名</div>
                  <div class="level-detail-div-item-content">{{ info.member ? info.member.fullName : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">手机号</div>
                  <div class="level-detail-div-item-content">{{ info.member ? info.member.phoneArea : '--' }}-{{
                    info.member ?
                      info.member.phone : '--' }}</div>
                </div>
              </n-gi>

            </n-grid>
          </div>
          <div class="level-detail-div" style="padding: 0 20px;" v-if="info.activityId > 0">
            <div class="level-detail-div-title">
              <div></div>
              员工福利活动
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">活动名称</div>
                  <div class="level-detail-div-item-content">{{ (info.employeeActivity &&
                    info.employeeActivity.nameLanguage) ?
                    info.employeeActivity.nameLanguage.content : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">员工名称</div>
                  <div class="level-detail-div-item-content">{{ info.employeeInfo ? info.employeeInfo.name : '--' }}
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>

          <div class="level-detail-div" style="padding: 0 20px;" v-if="info.indexActivityId > 0">
            <div class="level-detail-div-title">
              <div></div>
              首页活动
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">活动标题</div>
                  <div class="level-detail-div-item-content">{{ info.indexActivity ? info.indexActivity.title : '--' }}
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>

          <div class="level-detail-div" style="padding: 0 20px;" v-if="info.state == 3">
            <div class="level-detail-div-title">
              <div></div>
              核销信息
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi :span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">核销时间</div>
                  <div class="level-detail-div-item-content">{{ info.verifyTime }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">核销商户</div>
                  <div class="level-detail-div-item-content">{{ info.verifyMch ? info.verifyMch.name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">核销门店</div>
                  <div class="level-detail-div-item-content">{{ (info.verifyStore && info.verifyStore.nameLanguage) ?
                    info.verifyStore.nameLanguage.content : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">核销商品名</div>
                  <div class="level-detail-div-item-content">{{ info.thCouponMch ? info.thCouponMch.name : '--' }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
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
              适用商户
            </div>
            <n-grid y-gap="25" cols="1 600:3">
              <n-gi span="3" v-if="info.thCoupon && info.thCoupon.mchList">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title"></div>
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
                        <tr v-for="(item, index) of info.thCoupon.mchList" :key="index">
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

                            {{ item.mchInfo?.status }}
                          </td>
                        </tr>
                      </tbody>
                    </n-table>
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

        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { useMessage } from "naive-ui";
import { View } from "@/api/thMemberCoupon";
import { options, loadOptions } from './model'
import { adaModalWidth, adaTableScrollX, getOptionLabel, getOptionTag } from "@/utils/hotgo";
import { jsontoobj } from "@/utils/smjcomm";
import {
  options as mchOptions,
  loadDictOptions as mchLoadOptions
} from "@/views/thMch/model";
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

const getInfo = (id) => {
  loading.value = true;
  View({ id })
    .then((res) => {

      res.thCoupon.nameLanguage = jsontoobj(res.thCoupon.nameLanguage);
      res.thCoupon.subNameLanguage = jsontoobj(res.thCoupon.subNameLanguage);
      res.thCoupon.descLanguage = jsontoobj(res.thCoupon.descLanguage);

      info.value = res;

      info.value.zh_name = res.thCoupon.nameLanguage.zh && res.thCoupon.nameLanguage.zh.content ? res.thCoupon.nameLanguage.zh.content : ''
      info.value.en_name = res.thCoupon.nameLanguage.en && res.thCoupon.nameLanguage.en.content ? res.thCoupon.nameLanguage.en.content : ''
      info.value.ko_name = res.thCoupon.nameLanguage.ko && res.thCoupon.nameLanguage.ko.content ? res.thCoupon.nameLanguage.ko.content : ''
      info.value.ja_name = res.thCoupon.nameLanguage.ja && res.thCoupon.nameLanguage.ja.content ? res.thCoupon.nameLanguage.ja.content : ''
      if (res.thCoupon.nameLanguage.zh_CN) {
        info.value.zh_CN_name = res.thCoupon.nameLanguage.zh_CN.content ? res.thCoupon.nameLanguage.zh_CN.content : ''
      } else if (res.thCoupon.nameLanguage.zh_cn) {
        info.value.zh_CN_name = res.thCoupon.nameLanguage.zh_cn.content ? res.thCoupon.nameLanguage.zh_cn.content : ''
      } else {
        info.value.zh_CN_name = ''
      }

      info.value.zh_sub_name = res.thCoupon.subNameLanguage.zh && res.thCoupon.subNameLanguage.zh.content ? res.thCoupon.subNameLanguage.zh.content : ''
      info.value.en_sub_name = res.thCoupon.subNameLanguage.en && res.thCoupon.subNameLanguage.en.content ? res.thCoupon.subNameLanguage.en.content : ''
      info.value.ko_sub_name = res.thCoupon.subNameLanguage.ko && res.thCoupon.subNameLanguage.ko.content ? res.thCoupon.subNameLanguage.ko.content : ''
      info.value.ja_sub_name = res.thCoupon.subNameLanguage.ja && res.thCoupon.subNameLanguage.ja.content ? res.thCoupon.subNameLanguage.ja.content : ''
      if (res.thCoupon.nameLanguage.zh_CN) {
        info.value.zh_CN_sub_name = res.thCoupon.subNameLanguage.zh_CN.content ? res.thCoupon.subNameLanguage.zh_CN.content : ''
      } else if (res.thCoupon.nameLanguage.zh_cn) {
        info.value.zh_CN_sub_name = res.thCoupon.subNameLanguage.zh_cn.content ? res.thCoupon.subNameLanguage.zh_cn.content : ''
      } else {
        info.value.zh_CN_sub_name = ''
      }

      info.value.zh_desc = res.thCoupon.descLanguage.zh && res.thCoupon.descLanguage.zh.content ? res.thCoupon.descLanguage.zh.content : ''
      info.value.en_desc = res.thCoupon.descLanguage.en && res.thCoupon.descLanguage.en.content ? res.thCoupon.descLanguage.en.content : ''
      info.value.ko_desc = res.thCoupon.descLanguage.ko && res.thCoupon.descLanguage.ko.content ? res.thCoupon.descLanguage.ko.content : ''
      info.value.ja_desc = res.thCoupon.descLanguage.ja && res.thCoupon.descLanguage.ja.content ? res.thCoupon.descLanguage.ja.content : ''
      if (res.thCoupon.descLanguage.zh_CN) {
        info.value.zh_CN_desc = res.thCoupon.descLanguage.zh_CN.content ? res.thCoupon.descLanguage.zh_CN.content : ''
      } else if (res.thCoupon.descLanguage.zh_cn) {
        info.value.zh_CN_desc = res.thCoupon.descLanguage.zh_cn.content ? res.thCoupon.descLanguage.zh_cn.content : ''
      } else {
        info.value.zh_CN_desc = ''
      }

      console.log('info', info.value);
    })
    .finally(() => {
      loading.value = false;
    });
};

async function openModal(id) {
  couponId.value = id
  showModal.value = true;
  loading.value = true;
  loadOptions();
  mchLoadOptions();
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
