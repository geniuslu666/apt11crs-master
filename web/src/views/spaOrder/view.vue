<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">订单详情</div>
        </template>
        <template #footer v-if="!loading">
          <template v-if="formValue.orderStatus == 'WAIT_CONFIRM'">
            <n-button type="primary" v-if="hasPermission(['/spaOrder/confirmAgree'])" @click="agreeOrder" :loading="formBtnLoading1" style="width: 70px;height: 35px;margin-right: 10px">
              确认
            </n-button>
            <n-button type="warning" v-if="hasPermission(['/spaOrder/confirmDisagree'])" @click="disAgreeOrder" style="width: 70px;height: 35px;margin-right: 10px">
              拒绝
            </n-button>
          </template>
          <template v-if="formValue.orderStatus == 'WAIT_SERVE' && formValue.dispatchStatus == 'WAIT'">
            <n-button type="primary" v-if="hasPermission(['/spaOrder/dispatch'])" @click="dispatchOrder" style="width: 98px;height: 35px;margin-right: 10px">
              指派技师
            </n-button>
          </template>
          <template v-if="formValue.orderStatus == 'WAIT_SERVE' && formValue.serviceType == 2 && formValue.dispatchStatus == 'DONE' && !formValue.technicianGoTime">
            <n-button type="primary" v-if="hasPermission(['/spaOrder/goOut'])" @click="confirmGoOut" :loading="formBtnLoading" style="width: 98px;height: 35px;margin-right: 10px">
              技师出发
            </n-button>
          </template>
          <template v-if="formValue.orderStatus == 'WAIT_SERVE' && formValue.serviceType == 1 && formValue.dispatchStatus == 'DONE' && !formValue.memberArriveTime">
            <n-button type="primary" v-if="hasPermission(['/spaOrder/arrive'])" @click="confirmArrive" :loading="formBtnLoading" style="width: 98px;height: 35px;margin-right: 10px">
              客人到店
            </n-button>
          </template>
          <template v-if="formValue.orderStatus == 'WAIT_SERVE' && ((formValue.serviceType == 2 && formValue.technicianGoTime) || formValue.serviceType == 1 && formValue.memberArriveTime) && formValue.dispatchStatus == 'DONE'">
            <n-button type="primary" v-if="hasPermission(['/spaOrder/startService'])" @click="confirmStart" :loading="formBtnLoading" style="width: 98px;height: 35px;margin-right: 10px">
              开始服务
            </n-button>
          </template>
          <template v-if="formValue.orderStatus == 'SERVING'">
            <n-button type="primary" v-if="hasPermission(['/spaOrder/endService'])" @click="confirmEnd" :loading="formBtnLoading" style="width: 98px;height: 35px;margin-right: 10px">
              结束服务
            </n-button>
          </template>
          <template v-if="formValue.abnormalStatus == 2">
            <n-button type="error" @click="abnormalOrder" v-if="hasPermission(['/spaOrder/abnormal'])" style="width: 98px;height: 35px;margin-right: 10px">
              异常处理
            </n-button>
          </template>
          <template v-if="formValue.payStatus == 'HAVE_PAID' && formValue.orderStatus != 'DONE' && formValue.orderStatus != 'SERVING' && formValue.orderStatus != 'CANCEL' && formValue.adminCancelNum == 0">
            <n-button type="error" v-if="hasPermission(['/spaOrder/cancelPay'])" @click="refundOrder" style="width: 70px;height: 35px;margin-right: 10px">
              退款
            </n-button>
          </template>
          <template v-if="( formValue.orderStatus == 'CANCEL'|| formValue.orderStatus == 'SERVING' ) && formValue.refundAmount == 0 && (formValue.payStatus == 'HAVE_PAID' || (formValue.payTime!=null&&formValue.payTime != '')) && formValue.adminCancelNum == 0">
            <n-button type="error" v-if="hasPermission(['/spaOrder/cancelPay'])" @click="refundOrder" style="width: 70px;height: 35px;margin-right: 10px">
              退款
            </n-button>
          </template>
          <n-button v-if="hasPermission(['/printer/printSpaOrder'])" @click="printOrder" :loading="formBtnLoading4" style="width: 98px;height: 35px;">
            打印订单
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div style="padding: 0 20px">
            <div class="car-order-detail-top">
              {{ formValue.serviceType == 1 ? '到店' : '上门' }}服务
              <n-tag size="small" :bordered="false" type="error" style="margin-left: 10px" v-if="formValue.abnormalStatus == 2">
                异常待处理
              </n-tag>
              <n-tooltip trigger="hover" v-if="formValue.abnormalStatus == 3">
                <template #trigger>
                  <n-tag style="margin-left: 10px" :bordered="false" size="small" type="warning">
                    异常已处理
                  </n-tag>
                </template>
                {{ formValue.abnormalReason }}
              </n-tooltip>
            </div>
            <div class="car-order-detail-address" v-if="formValue.serviceType == 1">
              <div class="car-order-detail-address-item">
                <div class="car-order-detail-address-item-d1">门店</div>
                <div class="car-order-detail-address-item-d2">{{ formValue.storeDetail.name }}</div>
                <n-tooltip trigger="hover" :arrow-point-to-center="true">
                  <template #trigger>
                    <div class="car-order-detail-address-item-d3">
                      {{ formValue.storeDetail.detailAddress }}
                    </div>
                  </template>
                  {{ formValue.storeDetail.detailAddress }}
                </n-tooltip>
              </div>
            </div>
            <div class="car-order-detail-time" v-if="formValue.serviceType == 2">
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">物业</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.propertyDetail.name }}</div>
              </div>
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">房间号</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.roomNo }}</div>
              </div>
            </div>
            <div class="car-order-detail-time">
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">服务时间</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.bookStartTime }}</div>
              </div>
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">服务状态</div>
                <div class="car-order-detail-time-item-status" :class="getOptionTag(options.spa_order_status, formValue.orderStatus)">{{ getOptionLabel(options.spa_order_status, formValue.orderStatus) }}</div>
              </div>
            </div>
            <div class="car-order-detail-time">
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">下单会员</div>
                <div class="car-order-detail-time-item-d2">
                  <div>{{ formValue.memberDetail ? formValue.memberDetail.fullName : '-' }} / {{ formValue.memberDetail ? formValue.memberDetail.memberNo : '-' }}</div>
                  <div v-if="formValue.memberDeleted" style="color:red;font-size: 14px;margin-top: 5px">会员已注销</div>
                </div>
              </div>
            </div>
          </div>
          <div class="car-order-detail-tabs">
            <div class="car-order-detail-tabs-item" @click="nowTab = 1" :class="nowTab == 1 ? 'active' : ''">
              预约信息
              <div></div>
            </div>
            <div class="car-order-detail-tabs-item" @click="nowTab = 2" :class="nowTab == 2 ? 'active' : ''">
              订单信息
              <div></div>
            </div>
            <div class="car-order-detail-tabs-item" @click="nowTab = 3" :class="nowTab == 3 ? 'active' : ''">
              订单日志
              <div></div>
            </div>
          </div>
          <div style="padding: 30px 20px" v-if="nowTab == 1">
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                客人信息
              </div>
              <n-grid y-gap="25" :cols="2">
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">预订人</div>
                    <div class="level-detail-div-item-content">
                      <div>{{ formValue.bookingName }}</div>
                      <div>{{ formValue.phoneArea }} {{ formValue.bookingMobile }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">人数</div>
                    <div class="level-detail-div-item-content">
                      <div>{{ formValue.goodsNum }}人</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">邮箱</div>
                    <div class="level-detail-div-item-content">{{ formValue.bookingEmail }}</div>
                  </div>
                </n-gi>
                <n-gi span="2">
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">
                      用户留言
                      <div class="translate-div" @click="translateTab = translateTab == 'zh' ? 'ja' : 'zh'">
                        <img
                          style="width: 16px;margin-right: 3px"
                          src="@/assets/images/translate_icon.png"
                        />
                        <span v-if="translateTab == 'zh'">翻译</span>
                        <span v-else>原文</span>
                      </div>
                    </div>
                    <div class="level-detail-div-item-content" v-if="translateTab == 'zh'">{{ formValue.memberMessage ? formValue.memberMessage : '--' }}</div>
                    <div class="level-detail-div-item-content" v-else>{{ formValue.memberMessageJa ? formValue.memberMessageJa : '--' }}</div>
                  </div>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                服务详情
              </div>
              <n-grid y-gap="25" :cols="2">
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">预定服务</div>
                    <div class="level-detail-div-item-content">{{ formValue.serviceDetail.name }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">预定项目</div>
                    <div class="level-detail-div-item-content">{{ formValue.goodsDetail.goodsName }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">预约时间</div>
                    <div class="level-detail-div-item-content">{{ formValue.bookStartTime }}</div>
                  </div>
                </n-gi>
                <n-gi span="2">
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">指派技师/订单调度备注</div>
                    <div class="level-detail-div-item-content">{{ formValue.dispatchStatus == 'DONE' ? (formValue.dispatchDesc ? formValue.dispatchDesc : '无') : '--' }}</div>
                  </div>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                服务进度
              </div>
              <n-timeline>
                <n-timeline-item v-for="(item, key) in formValue.logList" :type="item.actionWay == 'DONE' ? 'info' : (item.actionWay == 'REFUND' || item.actionWay == 'OVERDUE' ? 'error' : (item.actionWay == 'CANCEL' ? 'warning' : 'default'))" :key="key" :title="item.remark" :time="item.createdAt"/>
              </n-timeline>
            </div>
          </div>
          <div style="padding: 30px 20px" v-else-if="nowTab == 2">
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                订单信息
              </div>
              <n-grid y-gap="25" :cols="2">
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单号</div>
                    <div class="level-detail-div-item-content">{{ formValue.orderSn }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">服务类型</div>
                    <div class="level-detail-div-item-content">{{ formValue.serviceType == 1 ? '到店' : '上门' }}服务</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单金额</div>
                    <div class="level-detail-div-item-content">{{ formValue.orderAmount }} JPY</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">支付状态</div>
                    <div>
                      <div class="level-detail-div-item-content1" :class="getOptionTag(options.spa_order_pay_status, formValue.payStatus)">{{ getOptionLabel(options.spa_order_pay_status, formValue.payStatus) }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单状态</div>
                    <div>
                      <div class="level-detail-div-item-content1" :class="getOptionTag(options.spa_order_status, formValue.orderStatus)">{{ getOptionLabel(options.spa_order_status, formValue.orderStatus) }}</div>
                    </div>
                  </div>
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
                <div class="transaction-div-item">第三方支付流水号：{{ item.paymentRequestId ? item.paymentRequestId : '--' }}</div>
                <div class="transaction-div-item">创建时间：{{ item.createdAt }}</div>
                <div class="transaction-div-item">
                  过期时间：{{ item.expiredTime }}
                  <div>
                    <n-tooltip trigger="hover" :arrow-point-to-center="true">
                      <template #trigger>
                        <span><span style="color: #3D3D3D">{{ item.amount }}</span>{{ item.priceCurrency ? item.priceCurrency : 'JPY' }}</span>
                      </template>
                      总金额
                    </n-tooltip>
                    <span style="margin: 0 5px">/</span>
                    <n-tooltip trigger="hover" :arrow-point-to-center="true">
                      <template #trigger>
                        <span><span style="color: #FF6F00">{{ item.payAmount }}</span>{{ item.priceCurrency ? item.priceCurrency : 'JPY' }}</span>
                      </template>
                      支付金额
                    </n-tooltip>
                  </div>
                </div>
                <div class="transaction-div-bottom">
                  <div>
                    支付平台：<span style="color: #1664FF">{{item.payChannel}}/{{item.payType}}</span>
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
                    退款方式：<span style="color: #1664FF">{{item.refundType}}</span>
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
                      <div style="width: 100%;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;">{{ item.actionWay == 'CREATE' ? '订单创建' : item.remark }}</div>
                    </template>
                    {{ item.remark }}
                  </n-tooltip>
                </div>
                <div style="width: 27.9%">{{ item.createdAt }}</div>
              </div>
            </div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <Status ref="statusRef" @reloadInfo="reloadInfo"/>
    <Dispatch ref="dispatchRef" @reloadInfo="reloadInfo"/>
    <Cancel ref="cancelRef" @reloadInfo="reloadInfo" />
    <Abnormal ref="abnormalRef" @reloadInfo="reloadInfo"/>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import { View, ConfirmAgree, GoOut, StartService, EndService, Arrive } from '@/api/spaOrder';
import { State, newState } from '@/views/spaOrder/model';
import {adaModalWidth, getOptionLabel, getOptionTag, Option} from '@/utils/hotgo';
import {Dicts} from "@/api/dict/dict";
import Status from "@/views/spaOrder/order_confirm.vue";
import Dispatch from "@/views/spaOrder/order_dispatch.vue";
import Cancel from "@/views/spaOrder/cancel.vue";
import {PrinterSpaOrder} from "@/api/printer";
import {usePermission} from "@/hooks/web/usePermission";
import Abnormal from "@/views/spaOrder/abnormal.vue";

const { hasPermission } = usePermission();
const statusRef = ref();
const dispatchRef = ref();
const cancelRef = ref();
const abnormalRef = ref();
const dialog = useDialog();
const emit = defineEmits(['reloadTable']);
const nowTab = ref(1)
const translateTab = ref('zh')
const formBtnLoading = ref(false);
const formBtnLoading1 = ref(false);
const formBtnLoading2 = ref(false);
const formBtnLoading3 = ref(false);
const formBtnLoading4 = ref(false);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});
const options = ref({
  spa_order_status: [] as Option[],
  spa_order_pay_status: [] as Option[],
  work_status: [] as Option[],
});

function openModal(state: State) {
  nowTab.value = 1;
  showModal.value = true;
  loading.value = true;
  Dicts({
    types: ['spa_order_status','spa_order_pay_status','work_status'],
  }).then((res) => {
    options.value = res;
    View({ orderSn: state.orderSn })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  });
}

function reloadInfo(){
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

// 确认订单
function agreeOrder(e) {
  e.preventDefault();
  formBtnLoading1.value = true;
  if(formValue.value.orderStatus != 'WAIT_CONFIRM'){
    formBtnLoading1.value = false;
    message.error('订单状态不正确');
    return false;
  }
  dialog.info({
    title: '提示',
    content: '你确定要确认该订单吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      ConfirmAgree({id: formValue.value.id}).then((_res) => {
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

// 拒绝订单
function disAgreeOrder(e) {
  e.preventDefault();
  if(formValue.value.orderStatus != 'WAIT_CONFIRM'){
    message.error('订单状态不正确');
    return false;
  }
  statusRef.value.openModal(formValue.value.id);
}

// 派单
function dispatchOrder(e){
  e.preventDefault();
  dispatchRef.value.openModal(formValue.value.id, formValue.value.goodsNum);
}

// 出发
function confirmGoOut(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(formValue.value.serviceType == 1){
    formBtnLoading.value = false;
    message.error('服务方式不正确');
    return false;
  }
  if(formValue.value.dispatchStatus == 'WAIT'){
    formBtnLoading.value = false;
    message.error('订单还未调度');
    return false;
  }
  if(formValue.value.technicianGoTime){
    formBtnLoading.value = false;
    message.error('技师已出发');
    return false;
  }
  dialog.info({
    title: '提示',
    content: '你确定要操作吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      GoOut({id: formValue.value.id}).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
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
        formBtnLoading.value = false;
      });
    },
    onNegativeClick: () => {
      formBtnLoading.value = false;
    },
  })
}

// 客人到店
function confirmArrive(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(formValue.value.serviceType != 1){
    formBtnLoading.value = false;
    message.error('服务方式不正确');
    return false;
  }
  if(formValue.value.dispatchStatus == 'WAIT'){
    formBtnLoading.value = false;
    message.error('订单还未调度');
    return false;
  }
  if(formValue.value.memberArriveTime){
    formBtnLoading.value = false;
    message.error('客人已到店');
    return false;
  }
  dialog.info({
    title: '提示',
    content: '你确定要操作吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Arrive({id: formValue.value.id}).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
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
        formBtnLoading.value = false;
      });
    },
    onNegativeClick: () => {
      formBtnLoading.value = false;
    },
  })
}


// 开始服务
function confirmStart(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(formValue.value.dispatchStatus == 'WAIT'){
    formBtnLoading.value = false;
    message.error('订单还未调度');
    return false;
  }
  dialog.info({
    title: '提示',
    content: '你确定要开始服务吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      StartService({id: formValue.value.id}).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
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
        formBtnLoading.value = false;
      });
    },
    onNegativeClick: () => {
      formBtnLoading.value = false;
    },
  })
}

// 结束服务
function confirmEnd(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(formValue.value.orderStatus != 'SERVING'){
    formBtnLoading.value = false;
    message.error('订单状态不正确');
    return false;
  }
  dialog.info({
    title: '提示',
    content: '你确定要结束服务吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      EndService({id: formValue.value.id}).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
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
        formBtnLoading.value = false;
      });
    },
    onNegativeClick: () => {
      formBtnLoading.value = false;
    },
  })
}


// 退款
function refundOrder(e){
  e.preventDefault();
  cancelRef.value.openModal(formValue.value);
}

// 异常处理
function abnormalOrder(e){
  e.preventDefault();
  abnormalRef.value.openModal(formValue.value.id);
}

// 打印订单
function printOrder(e){
  e.preventDefault();
  dialog.info({
    title: '提示',
    content: '你确定要打印该订单吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      PrinterSpaOrder({orderId: formValue.value.id}).then((_res) => {
        message.success('操作成功');
      });
    },
    onNegativeClick: () => {},
  });
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

.car-order-detail-top{
  display: flex;
  align-items: center;
  font-weight: 500;
  font-size: 18px;
  color: #3D3D3D;
  line-height: 25px;
  /*div{
    padding: 0 6px;
    font-weight: 400;
    font-size: 14px;
    color: #FFFFFF;
    line-height: 24px;
    border-radius: 2px;
    margin-left: 8px;
    &.crs{
      background: #12C584;
    }
    &.innn{
      background: #FC752F;
    }
  }*/
}
.car-order-detail-address{
  margin-top: 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  .car-order-detail-address-item{
    width: 230px;
    &-d1{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
    }
    &-d2{
      margin-top: 8px;
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      white-space: nowrap; /* 防止文本换行 */
      overflow: hidden; /* 隐藏溢出的内容 */
      text-overflow: ellipsis; /* 使用省略号表示溢出的文本 */
    }
    &-d3{
      margin-top: 2px;
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      white-space: nowrap; /* 防止文本换行 */
      overflow: hidden; /* 隐藏溢出的内容 */
      text-overflow: ellipsis; /* 使用省略号表示溢出的文本 */
      span{
        display: block;
        width: 100%;
        white-space: nowrap; /* 防止文本换行 */
        overflow: hidden; /* 隐藏溢出的内容 */
        text-overflow: ellipsis; /* 使用省略号表示溢出的文本 */
      }
    }
  }
}
.car-order-detail-time{
  margin-top: 26px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  .car-order-detail-time-item{
    width: 230px;
    &-d1{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
    }
    &-d2{
      margin-top: 8px;
      font-weight: 500;
      font-size: 16px;
      color: #3D3D3D;
      line-height: 22px;
    }
    &-status{
      margin-top: 8px;
      border-radius: 2px;
      font-weight: 400;
      font-size: 14px;
      line-height: 22px;
      display: inline-block;
      padding: 0 5px;
      &.success{
        color: #19A158;
        background: #E3F4EB;
      }
      &.info{
        color: #3F9EFF;
        background: #ECF5FF;
      }
      &.default{
        color: #919399;
        background: #F4F4F5;
      }
      &.warning{
        color: #EFA020;
        background: #FDF1DD;
      }
      &.error{
        color: #F56C6C;
        background: #FEF0F0;
      }
    }
  }
}
.car-order-detail-tabs{
  background: #F9FBFC;
  height: 40px;
  line-height: 40px;
  padding-left: 20px;
  display: flex;
  margin-top: 44px;
  &-item{
    height: 40px;
    line-height: 40px;
    position: relative;
    font-weight: 500;
    font-size: 14px;
    color: #3D3D3D;
    margin-right: 36px;
    cursor: pointer;
    &.active{
      color: #053DC8;
      div{
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
.level-detail-div{
  margin-bottom: 25px;
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
      display: flex;
      align-items: center;
      .translate-div{
        cursor: pointer;
        margin-left: 8px;
        display: flex;
        align-items: center;
        span{
          font-weight: 400;
          font-size: 12px;
          color: #1664FF;
          line-height: 17px;
        }
      }
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
    }
    &-content1{
      padding: 0 6px;
      height: 24px;
      line-height: 24px;
      font-weight: 400;
      font-size: 14px;
      display: inline-block;
      &.success{
        color: #19A158;
        background: #E3F4EB;
      }
      &.info{
        color: #3F9EFF;
        background: #ECF5FF;
      }
      &.default{
        color: #919399;
        background: #F4F4F5;
      }
      &.warning{
        color: #EFA020;
        background: #FDF1DD;
      }
      &.error{
        color: #F56C6C;
        background: #FEF0F0;
      }
    }
  }
}
.transaction-div{
  background: #F9FBFC;
  padding: 20px;
  margin-bottom: 20px;
  &:last-child{
    margin-bottom: 0;
  }
  &-top{
    display: flex;
    justify-content: space-between;
    align-items: center;
    span{
      font-weight: 500;
      font-size: 14px;
      line-height: 20px;
      &.title{
        color: #3D3D3D;
      }
    }
  }
  &-item{
    margin-top: 10px;
    font-weight: 400;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    div{
      font-weight: 500;
      font-size: 16px;
      color: #979797;
      line-height: 22px;
      display: flex;
      align-items: center;
    }
  }
  &-bottom{
    margin-top: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    div{
      &:first-child{
        font-weight: 400;
        font-size: 14px;
        color: #3D3D3D;
        line-height: 20px;
      }
      &:last-child{
        font-weight: 400;
        font-size: 14px;
        color: #979797;
        line-height: 20px;
      }
    }
  }
}
.order-table-header{
  display: flex;
  align-items: center;
  height: 40px;
  line-height: 40px;
  background: #F7F8FA;
  div{
    font-weight: 500;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 40px;
  }
}
.order-table-body{
  &-item{
    display: flex;
    align-items: center;
    border-bottom: 1px solid #EEEEEE;
    height: 56px;
    &:hover{
      background: #FAFAFC;
    }
    div{
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
    }
    background: #FFFFFF;
    &:nth-child(even) {
      background: #FAFAFC; /* 例如，将背景色设置为灰色 */
    }
  }
}
</style>


