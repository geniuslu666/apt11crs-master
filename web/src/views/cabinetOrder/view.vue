<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
        padding: '20px',
      }" :body-content-style="{
        padding: '25px 0 0',
      }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">订单详情-编号{{ formValue.orderSn }}
          </div>
        </template>
        <template #footer v-if="!loading">
          <template
            v-if="formValue.orderStatus == 'ING' || formValue.orderStatus == 'OVERTIME' || formValue.orderStatus == 'GRACE'">
            <n-button type="primary" v-if="hasPermission(['/cabinetOrder/complete'])" @click="completeOrder"
              :loading="formBtnLoading1" style="width: 85px;height: 35px;margin-right: 10px">
              强制完成
            </n-button>
          </template>
          <template
            v-if="formValue.orderStatus == 'ING' || formValue.orderStatus == 'DONE' || formValue.orderStatus == 'OVERTIME' || formValue.orderStatus == 'GRACE' || formValue.orderStatus == 'HAVE_PAID'">
            <n-button type="error" v-if="hasPermission(['/cabinetOrder/refund']) && formValue.refundStatus != 'DONE'"
              @click="refundOrder" style="width: 70px;height: 35px;margin-right: 10px">
              退款
            </n-button>
          </template>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div style="padding: 0 20px">
            <div class="car-order-detail-time">
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">柜门号</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.boxNo }}({{ formValue.boxTypeName }})</div>
              </div>
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">订单状态</div>
                <div class="car-order-detail-time-item-status"
                  :class="getOptionTag(options.cabinet_order_status, formValue.orderStatus)">{{
                    getOptionLabel(options.cabinet_order_status, formValue.orderStatus) }}</div>
              </div>
            </div>
            <div class="car-order-detail-time">
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">网点名称</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.mchBranchName }}</div>
              </div>
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">使用时长</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.useTimeStr }}</div>
              </div>
            </div>
            <div class="car-order-detail-time">
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">基础费用</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.baseAmount }}</div>
              </div>
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">超时支付费用</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.overtimeFee }}</div>
              </div>
            </div>
          </div>

          <div class="car-order-detail-tabs hs">
            <div class="car-order-detail-tabs-item" @click="nowTab = 1" :class="nowTab == 1 ? 'active' : ''">
              订单信息
              <div></div>
            </div>
            <div class="car-order-detail-tabs-item" @click="nowTab = 2" :class="nowTab == 2 ? 'active' : ''">
              机柜信息
              <div></div>
            </div>
            <div class="car-order-detail-tabs-item" @click="nowTab = 3" :class="nowTab == 3 ? 'active' : ''">
              订单日志
              <div></div>
            </div>
            <div class="car-order-detail-tabs-item" @click="nowTab = 4" :class="nowTab == 4 ? 'active' : ''">
              开柜记录
              <div></div>
            </div>
          </div>
          <div style="padding: 30px 20px" v-if="nowTab == 1">
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                订单信息
              </div>
              <n-grid y-gap="25" :cols="2">
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单编号</div>
                    <div class="level-detail-div-item-content">{{ formValue.orderSn }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">支付状态</div>
                    <div>
                      <div class="level-detail-div-item-content1"
                        :class="getOptionTag(options.spa_order_pay_status, formValue.payStatus)">{{
                          getOptionLabel(options.spa_order_pay_status, formValue.payStatus) }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">下单时间</div>
                    <div class="level-detail-div-item-content">{{ formValue.createdAt }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单金额</div>
                    <div class="level-detail-div-item-content">{{ formValue.orderAmount }} JPY</div>
                  </div>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                用户信息
              </div>
              <n-grid y-gap="25" :cols="2">
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">会员昵称</div>
                    <div class="level-detail-div-item-content">{{ formValue.memberInfo.fullName }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">会员号</div>
                    <div>
                      <div class="level-detail-div-item-content1">{{ formValue.memberInfo.memberNo }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">手机号</div>
                    <div class="level-detail-div-item-content">{{ formValue.memberInfo.phoneArea }} {{
                      formValue.memberInfo.phone }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">邮箱</div>
                    <div class="level-detail-div-item-content">{{ formValue.memberInfo.mail }}</div>
                  </div>
                </n-gi>
                <n-gi :span="2" v-if="formValue.memberDeleted">
                  <div style="color: red;font-size: 14px;">会员已注销</div>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div" v-if="formValue.transactionDetail">
              <div class="level-detail-div-title">
                <div></div>
                交易流水
              </div>
              <div class="transaction-div" v-for="(item, index) in formValue.transactionDetail" :key="index">
                <div class="transaction-div-top">
                  <span class="title">支付流水号：{{ item.transactionSn ? item.transactionSn : '--' }}</span>
                  <span>
                    <span v-if="item.payStatus == 'WAIT'" class="cblue">等待支付</span>
                    <span v-else-if="item.payStatus == 'DONE'" class="csuccess">完成支付</span>
                    <span v-else-if="item.payStatus == 'CANCEL'" class="c999">取消支付</span>
                  </span>
                </div>
                <div class="transaction-div-item">第三方支付流水号：{{ item.paymentRequestId ? item.paymentRequestId : '--' }}
                </div>
                <div class="transaction-div-item">创建时间：{{ item.createdAt }}</div>
                <div class="transaction-div-item">
                  过期时间：{{ item.expiredTime }}
                  <div>
                    <n-tooltip trigger="hover" :arrow-point-to-center="true">
                      <template #trigger>
                        <span><span style="color: #3D3D3D">{{ item.amount }}</span>{{ item.priceCurrency ?
                          item.priceCurrency : 'JPY' }}</span>
                      </template>
                      总金额
                    </n-tooltip>
                    <span style="margin: 0 5px">/</span>
                    <n-tooltip trigger="hover" :arrow-point-to-center="true">
                      <template #trigger>
                        <span><span style="color: #FF6F00">{{ item.payAmount }}</span>{{ item.priceCurrency ?
                          item.priceCurrency : 'JPY' }}</span>
                      </template>
                      支付金额
                    </n-tooltip>
                  </div>
                </div>
                <div class="transaction-div-bottom">
                  <div>
                    支付平台：<span style="color: #1664FF">{{ item.payChannel }}/{{ item.payType }}</span>
                  </div>
                  <div>(总退款金额：{{ item.refundAmount }}{{ item.priceCurrency ? item.priceCurrency : 'JPY' }})</div>
                </div>
              </div>
            </div>
            <div class="level-detail-div" v-if="formValue.transactionRefundDetail">
              <div class="level-detail-div-title">
                <div></div>
                退款流水
              </div>
              <div class="transaction-div" v-for="(item, index) in formValue.transactionRefundDetail" :key="index">
                <div class="transaction-div-top">
                  <span class="title">退款单号：{{ item.refundSn }}</span>
                </div>
                <div class="transaction-div-bottom">
                  <div style="color: #3D3D3D;">
                    退款方式：<span style="color: #1664FF">{{ item.refundType }}</span>
                  </div>
                </div>
                <div class="transaction-div-item">
                  退款时间：{{ item.refundTime ? item.refundTime : item.updatedAt }}
                  <div>
                    <span><span style="color: #3D3D3D">{{ item.refundAmount }}</span>JPY</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div style="padding: 30px 20px" v-else-if="nowTab == 2">
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                机柜信息
              </div>
              <n-grid y-gap="25" :cols="2">
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">机柜编号</div>
                    <div class="level-detail-div-item-content">
                      <div>{{ formValue.boxNo }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">柜门号</div>
                    <div class="level-detail-div-item-content">
                      <div>{{ formValue.boxTypeName }}</div>
                      <div>{{ formValue.boxNo }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">归属网点</div>
                    <div class="level-detail-div-item-content">
                      <div>{{ formValue.mchBranchName }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">储物柜</div>
                    <div class="level-detail-div-item-content">
                      <div>{{ formValue.cabinetName }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">PIN码</div>
                    <div class="level-detail-div-item-content">
                      <div>{{ formValue.pin }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">基础费用</div>
                    <div class="level-detail-div-item-content">{{ formValue.boxTypePrice }} / 时</div>
                  </div>
                </n-gi>
                <n-gi span="2">
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">
                      机柜地址
                    </div>
                    <div class="level-detail-div-item-content">{{ formValue.address }}</div>
                  </div>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                收费规则
              </div>

              <div class="order-table-header">
                <div style="padding-left: 3.3%;width: 25%">柜门类型</div>
                <div style="width: 35%">基础费用</div>
                <div style="width: 35%">超时费用</div>
              </div>
              <div class="order-table-body">
                <div class="order-table-body-item" v-for="(item, index) in formValue.boxTypeJson" :key="index">
                  <div style="padding-left: 3.3%;width: 25%">{{ item.name }}</div>
                  <div style="width: 35%">{{ item.price }}/小时</div>
                  <div style="width: 35%">{{ item.price }}/小时</div>
                </div>
              </div>
            </div>

          </div>
          <div style="padding: 30px 20px" v-else-if="nowTab == 3">
            <div class="order-table-header">
              <div style="padding-left: 3.3%;width: 20%">操作人</div>
              <div style="width: 21.8%">操作角色</div>
              <div style="width: 30.3%">操作内容</div>
              <div style="width: 27.9%">操作时间</div>
            </div>
            <div class="order-table-body">
              <div class="order-table-body-item" v-for="(item, index) in formValue.logList" :key="index">
                <div style="padding-left: 3.3%;width: 20%">{{ item.operateName }}</div>
                <div style="width: 21.8%">{{ item.operateType }}</div>
                <div style="width: 30.3%;padding-right: 2%">
                  <n-tooltip trigger="hover" :arrow-point-to-center="true">
                    <template #trigger>
                      <div style="width: 100%;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;">{{
                        item.actionWay == 'CREATE' ? '订单创建' : item.remark }}</div>
                    </template>
                    {{ item.remark }}
                  </n-tooltip>
                </div>
                <div style="width: 27.9%">{{ item.createdAt }}</div>
              </div>
            </div>
          </div>

          <div style="padding: 30px 20px;text-align: center" v-else-if="nowTab == 4">
            <div class="order-table-header">
              <div style="width: 50%">时间</div>
              <div style="width: 50%">事件</div>
            </div>
            <div class="order-table-body">
              <div class="order-table-body-item" v-for="(item, index) in formValue.openLogList" :key="index">
                <div style="width: 50%">{{ item.datetime }}</div>
                <div style="width: 50%">{{ item.type == 1 ? '存物开箱' : item.type == 2 ? '临时开柜' : '结束使用' }}</div>
              </div>
            </div>
          </div>

        </n-spin>
      </n-drawer-content>
    </n-drawer>
    <Cancel ref="cancelRef" @reloadInfo="reloadInfo" />
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import { View, OrderComplete } from '@/api/cabinetOrder';
import { State, newState } from '@/views/cabinetOrder/model';
import { adaModalWidth, getOptionLabel, getOptionTag, Option } from '@/utils/hotgo';
import { Dicts } from "@/api/dict/dict";
import { usePermission } from "@/hooks/web/usePermission";
import Cancel from "@/views/cabinetOrder/cancel.vue";

const { hasPermission } = usePermission();
const cancelRef = ref();
const dialog = useDialog();
const emit = defineEmits(['reloadTable']);
const nowTab = ref(1)
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const formBtnLoading1 = ref(false);
const options = ref({
  cabinet_order_status: [] as Option[],
  spa_order_pay_status: [] as Option[],
});

function openModal(state: State) {
  nowTab.value = 1;
  showModal.value = true;
  loading.value = true;
  Dicts({
    types: ['cabinet_order_status', 'spa_order_pay_status'],
  }).then((res) => {
    options.value = res;
    View({ orderSn: state.orderSn })
      .then((res) => {
        res.useTimeStr = "--";
        if (res && res.useTime > 0) {
          const hours = Math.floor(res.useTime / 3600);
          const minutes = Math.floor((res.useTime % 3600) / 60);
          const seconds = res.useTime % 60;
          res.useTimeStr = `${hours}时${minutes}分${seconds}秒`;
        }
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  });
}

// 退款
function refundOrder(e) {
  e.preventDefault();
  cancelRef.value.openModal(formValue.value);
}

// 强制完成
function completeOrder(e) {
  e.preventDefault();
  formBtnLoading1.value = true;
  if (formValue.value.orderStatus != 'ING' && formValue.value.orderStatus == 'OVERTIME' && formValue.value.orderStatus == 'GRACE') {
    formBtnLoading1.value = false;
    message.error('订单状态不正确');
    return false;
  }
  dialog.info({
    title: '提示',
    content: '你确定要强制完成该订单吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      OrderComplete({ id: formValue.value.id }).then((_res) => {
        message.success('操作成功');
        formBtnLoading1.value = false;
        setTimeout(() => {
          loading.value = true;
          View({ orderSn: formValue.value.orderSn })
            .then((res) => {
              formValue.value = res;
            })
            .finally(() => {
              loading.value = false;
            });
          emit('reloadTable');
        }, 500);
      }).catch((err) => {
        formBtnLoading1.value = false;
      });
    },
    onNegativeClick: () => {
      formBtnLoading1.value = false;
    },
  });
}

function reloadInfo() {
  loading.value = true;
  View({ orderSn: formValue.value.orderSn })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
  emit('reloadTable');
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
.payitem {
  border-radius: 5px;
  padding: 5px 10px;
  line-height: 25px;
  background: #f9fbfcff;
  margin-top: 10px;
}

.car-order-detail-top {
  display: flex;
  align-items: center;
  font-weight: 500;
  font-size: 18px;
  color: #3D3D3D;
  line-height: 25px;

  &-type {
    padding: 0 6px;
    font-weight: 400;
    font-size: 14px;
    color: #FFFFFF;
    line-height: 24px;
    border-radius: 2px;
    margin-left: 8px;

    &.crs {
      background: #12C584;
    }

    &.innn {
      background: #FC752F;
    }
  }
}

.car-order-detail-address {
  margin-top: 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .car-order-detail-address-item {
    width: 230px;

    &-d1 {
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
    }

    &-d2 {
      margin-top: 8px;
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      white-space: nowrap;
      /* 防止文本换行 */
      overflow: hidden;
      /* 隐藏溢出的内容 */
      text-overflow: ellipsis;
      /* 使用省略号表示溢出的文本 */
    }

    &-d3 {
      margin-top: 2px;
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      white-space: nowrap;
      /* 防止文本换行 */
      overflow: hidden;
      /* 隐藏溢出的内容 */
      text-overflow: ellipsis;

      /* 使用省略号表示溢出的文本 */
      span {
        display: block;
        width: 100%;
        white-space: nowrap;
        /* 防止文本换行 */
        overflow: hidden;
        /* 隐藏溢出的内容 */
        text-overflow: ellipsis;
        /* 使用省略号表示溢出的文本 */
      }
    }
  }
}

.car-order-detail-time {
  margin-top: 26px;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .car-order-detail-time-item {
    width: 230px;

    &-d1 {
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
    }

    &-d2 {
      margin-top: 8px;
      font-weight: 500;
      font-size: 16px;
      color: #3D3D3D;
      line-height: 22px;
    }

    &-status {
      margin-top: 8px;
      border-radius: 2px;
      font-weight: 400;
      font-size: 14px;
      line-height: 22px;
      display: inline-block;
      padding: 0 5px;

      &.success {
        color: #19A158;
        background: #E3F4EB;
      }

      &.info {
        color: #3F9EFF;
        background: #ECF5FF;
      }

      &.default {
        color: #919399;
        background: #F4F4F5;
      }

      &.warning {
        color: #EFA020;
        background: #FDF1DD;
      }

      &.error {
        color: #F56C6C;
        background: #FEF0F0;
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
  margin-top: 44px;
  flex-wrap: nowrap;
  overflow-x: auto;
  gap: 36px;

  &-item {
    height: 40px;
    line-height: 40px;
    position: relative;
    font-weight: 500;
    font-size: 14px;
    color: #3D3D3D;
    margin-right: 36px;
    cursor: pointer;
    white-space: nowrap;

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
      display: flex;
      align-items: center;

      .translate-div {
        cursor: pointer;
        margin-left: 8px;
        display: flex;
        align-items: center;

        span {
          font-weight: 400;
          font-size: 12px;
          color: #1664FF;
          line-height: 17px;
        }
      }
    }

    &-content {
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;

      .color {
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
    }

    &-content1 {
      padding: 0 6px;
      height: 24px;
      line-height: 24px;
      font-weight: 400;
      font-size: 14px;
      display: inline-block;

      &.success {
        color: #19A158;
        background: #E3F4EB;
      }

      &.info {
        color: #3F9EFF;
        background: #ECF5FF;
      }

      &.default {
        color: #919399;
        background: #F4F4F5;
      }

      &.warning {
        color: #EFA020;
        background: #FDF1DD;
      }

      &.error {
        color: #F56C6C;
        background: #FEF0F0;
      }
    }
  }
}

.transaction-div {
  background: #F9FBFC;
  padding: 20px;
  margin-bottom: 20px;

  &:last-child {
    margin-bottom: 0;
  }

  &-top {
    display: flex;
    justify-content: space-between;
    align-items: center;

    span {
      font-weight: 500;
      font-size: 14px;
      line-height: 20px;

      &.title {
        color: #3D3D3D;
      }
    }
  }

  &-item {
    margin-top: 10px;
    font-weight: 400;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    div {
      font-weight: 500;
      font-size: 16px;
      color: #979797;
      line-height: 22px;
      display: flex;
      align-items: center;
    }
  }

  &-bottom {
    margin-top: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    div {
      &:first-child {
        font-weight: 400;
        font-size: 14px;
        color: #3D3D3D;
        line-height: 20px;
      }

      &:last-child {
        font-weight: 400;
        font-size: 14px;
        color: #979797;
        line-height: 20px;
      }
    }
  }
}

.order-table-header {
  display: flex;
  align-items: center;
  height: 40px;
  line-height: 40px;
  background: #F7F8FA;

  div {
    font-weight: 500;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 40px;
  }
}

.order-table-body {
  &-item {
    display: flex;
    align-items: center;
    border-bottom: 1px solid #EEEEEE;
    height: 56px;

    &:hover {
      background: #FAFAFC;
    }

    div {
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
    }

    background: #FFFFFF;

    &:nth-child(even) {
      background: #FAFAFC;
      /* 例如，将背景色设置为灰色 */
    }
  }
}
</style>
