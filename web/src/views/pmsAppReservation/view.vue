<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content :native-scrollbar="false" closable :header-style="{
        padding: '20px',
      }" :body-content-style="{
        padding: '0',
      }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ viewtype }}</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="order-detail-info">
            <div class="order-detail-info-top">
              <div class="order-detail-info-top-info">
                <div class="order-detail-info-top-info-t flex-col !items-start gap-4 md:flex-row">
                  <div class="order-detail-info-top-info-t-l max-w-[410px]">{{ formValue.propertyDetail?.name }}</div>
                  <div class="order-detail-info-top-info-t-r">
                    <div class="order-detail-info-top-info-t-r-v1"
                      v-if="formValue.appReservation && formValue.appReservation[0].sourceName != ''">{{
                        formValue.appReservation[0].sourceName }}</div>
                    <div class="order-detail-info-top-info-t-r-v2"
                      :class="getOptionTag(options.order_status, formValue.orderStatus)" v-if="viewtype != 'AIRHOST订单'">
                      {{
                        getOptionLabel(options.order_status, formValue.orderStatus) }}</div>
                    <div class="order-detail-info-top-info-t-r-v2"
                      :class="getOptionTag(options.status, formValue.appReservation[0].status)"
                      v-if="viewtype == 'AIRHOST订单'">
                      {{ getOptionLabel(options.status, formValue.appReservation[0].status) }}</div>
                  </div>
                </div>
                <div class="order-detail-info-top-info-ordersn">
                  <div class="order-detail-info-top-info-ordersn-item">
                    系统单号：{{ formValue.orderSn }}
                    <div @click="copyTextAuto(formValue.orderSn)"></div>
                  </div>
                  <div class="order-detail-info-top-info-ordersn-item">
                    预定单号：<span>{{ formValue.outOrderSn }}</span>
                    <div @click="copyTextAuto(formValue.outOrderSn)"></div>
                  </div>
                </div>
              </div>
            </div>
            <div class="order-detail-info-center flex-col gap-4 md:flex-row">
              <div class="order-detail-info-center-v" style="width: 151px" v-if="formValue.source != 'AIRHOST'">
                <div class="order-detail-info-center-v-label">预订者：</div>
                <div class="order-detail-info-center-v-cont">{{ formValue.guestProfileDetail ?
                  formValue.guestProfileDetail.fullName : '--' }} </div>
                <div class="order-detail-info-center-v-cont"><span style="color:#00B42A">{{ formValue.memberDetail ?
                  formValue.memberDetail.memberNo : '' }}</span></div>
                <div class="order-detail-info-center-v-cont" v-if="formValue.memberDeleted"><span style="color:red">会员已注销</span></div>
                <!-- <div class="order-detail-info-center-v-btn btn1" v-if="formValue.guestProfileDetail" @click="tomessage()">
                  <div class="img"></div>
                  <div class="btn-text">进入对话</div>
                </div> -->
              </div>
              <div class="order-detail-info-center-v" style="width: 195px">
                <div class="order-detail-info-center-v-label">Tel：</div>
                <div class="order-detail-info-center-v-cont">{{
                  formValue.guestProfileDetail ? formValue.guestProfileDetail.areaNo : '--' }}-{{
                    formValue.guestProfileDetail ? formValue.guestProfileDetail.phone : '--' }}</div>
                <!-- <div class="order-detail-info-center-v-btn btn2" @click="sendSmsClick()">
                  <div class="img"></div>
                  <div class="btn-text">发短信</div>
                </div> -->
              </div>
              <div class="order-detail-info-center-v" style="flex: 1;">
                <div class="order-detail-info-center-v-label">E-mail：</div>
                <div class="order-detail-info-center-v-cont">{{
                  formValue.guestProfileDetail ? formValue.guestProfileDetail.email : '--' }}</div>
                <!-- <div class="order-detail-info-center-v-btn btn3" @click="goemail()">
                  <div class="img"></div>
                  <div class="btn-text">发邮件</div>
                </div> -->
              </div>
            </div>
            <div class="order-detail-info-bottom  flex-col gap-4 md:flex-row">
              <div class="order-detail-info-bottom-v" style="width: 151px">
                <div>到店日期</div>
                <div>{{ formValue.checkInDate }}</div>
              </div>
              <div class="order-detail-info-bottom-v" style="width: 195px">
                <div>离店日期</div>
                <div>{{ formValue.checkOutDate }}</div>
              </div>
              <div class="order-detail-info-bottom-v" style="flex: 1;">
                <div>天数</div>
                <div>{{ nightNum }}晚</div>
              </div>
            </div>
          </div>
          <div class="order-detail-tabs">
            <div class="order-detail-tabs-item" @click="nowTab = 1" :class="nowTab == 1 ? 'active' : ''">
              基础信息
              <div></div>
            </div>
            <div class="order-detail-tabs-item" @click="nowTab = 2" :class="nowTab == 2 ? 'active' : ''">
              订单信息
              <div></div>
            </div>
            <div class="order-detail-tabs-item" @click="nowTab = 3" :class="nowTab == 3 ? 'active' : ''">
              订单日志
              <div></div>
            </div>
          </div>
          <div class="basic-info" v-if="nowTab == 1">
            <div class="basic-info-room">
              <div class="basic-info-room-title">
                <div></div>房间信息
              </div>
              <div class="basic-info-room-list flex-col gap-4 md:flex-row"
                v-for="(item, index) in formValue.appReservation" :key="index">
                <n-image :src="item.roomTypeDetail.cover" :object-fit="'cover'"
                  style="width: 75px; height: 75px; border-radius: 2px" />
                <div class="basic-info-room-list-info">
                  <div class="basic-info-room-list-info-t">
                    <div class="basic-info-room-list-info-t-tag"
                      :class="getOptionTag(options.checkin_status, item.checkinStatus)">{{
                        getOptionLabel(options.checkin_status, item.checkinStatus) }}</div>
                    <div class="basic-info-room-list-info-t-name">{{ item.roomTypeDetail.name }}<template
                        v-if="item.roomUnitDetail">#{{ item.roomUnitDetail.roomNo }}</template></div>
                  </div>
                  <div class="basic-info-room-list-info-c">{{ item.checkinDate }} {{ item.checkinTime }} ~ {{
                    item.checkoutDate
                  }} {{ item.checkoutTime }}</div>
                  <div class="basic-info-room-list-info-b flex-col !items-start gap-4 md:flex-row">
                    <div>下单时间：{{ item.createdAt }}</div>
                    <div>客人信息：<span>{{ item.adultCount }}成人/{{ item.childCount }}儿童</span></div>
                  </div>

                  <template v-if="item.reservationChangeList && item.reservationChangeList.length > 0">
                    <template v-for="(changeItem, index1) in item.reservationChangeList" :key="index1">
                      <div class="payitem" v-if="changeItem.changeStatus == 'DONE'">
                        <div v-if="changeItem.changeType == 'PEOPLE'">
                          <div>入住人数变更，变更时间:{{ changeItem.doneDate }}，变更后订单金额:{{ changeItem.newOrderPrice }}，变动金额:{{
                            changeItem.changeAmount }}</div>
                          <div class="flex-item">
                            <div class="flex-row">
                              <span class="c999 f12 mr-1">新人数：</span>
                              <img src="@/assets/images/cr.png" style="width: 16px; height: 16px; margin-right: 5px" />
                              <div class="c333" style="margin-right: 5px">成人：{{ changeItem.newAdultCount }}</div>
                              <img src="@/assets/images/ye.png" style="width: 16px; height: 16px; margin-right: 5px" />
                              <div class="c333">儿童：{{ changeItem.newChildCount }}</div>
                              <img src="@/assets/images/ye.png" style="width: 16px; height: 16px; margin-right: 5px" />
                              <div class="c333">婴儿：{{ changeItem.newInfantCount }}</div>
                            </div>
                            <div class="flex-row">
                              <span class="c999 f12 mr-1">旧人数：</span>
                              <img src="@/assets/images/cr.png" style="width: 16px; height: 16px; margin-right: 5px" />
                              <div class="c333" style="margin-right: 5px">成人：{{ changeItem.oldAdultCount }}</div>
                              <img src="@/assets/images/ye.png" style="width: 16px; height: 16px; margin-right: 5px" />
                              <div class="c333">儿童：{{ changeItem.oldChildCount }}</div>
                              <img src="@/assets/images/ye.png" style="width: 16px; height: 16px; margin-right: 5px" />
                              <div class="c333">婴儿：{{ changeItem.oldInfantCount }}</div>
                            </div>
                          </div>
                        </div>
                        <div v-if="changeItem.changeType == 'DATE'">
                          <div>日期变更，变更时间:{{ changeItem.doneDate }}，变更后订单金额:{{ changeItem.newOrderPrice }}，变动金额:{{
                            changeItem.changeAmount }}</div>
                          <div class="flex-item">
                            <div class="flex-row">
                              <span class="c999 f12 mr-1">新日期：</span>
                              <div class="c333" style="margin-right: 5px">{{ changeItem.newCheckinDate }} ~ {{
                                changeItem.newCheckoutDate }}</div>
                            </div>
                            <div class="flex-row">
                              <span class="c999 f12 mr-1">原日期：</span>
                              <div class="c333" style="margin-right: 5px">{{ changeItem.oldCheckinDate }} ~ {{
                                changeItem.oldCheckoutDate }}</div>
                            </div>
                          </div>
                        </div>
                        <div v-if="changeItem.changeType == 'GUEST'">
                          <div>预订人信息变更，变更时间:{{ changeItem.doneDate }}</div>
                          <div class="flex-item">
                            <div class="flex-row">
                              <span class="c999 f12 mr-1">新预定人：</span>
                              <div class="c333" style="margin-right: 5px">姓名:{{ changeItem.newMainGuest.full_name }} |
                                电话:{{ changeItem.newMainGuest.phone }} | 邮箱: {{ changeItem.newMainGuest.email }}</div>
                            </div>
                            <div class="flex-row">
                              <span class="c999 f12 mr-1">原预定人：</span>
                              <div class="c333" style="margin-right: 5px">姓名:{{ changeItem.oldMainGuest.full_name }} |
                                电话:{{ changeItem.oldMainGuest.phone }} | 邮箱: {{ changeItem.oldMainGuest.email }}</div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </template>
                  </template>
                </div>
              </div>
            </div>
            <div class="basic-info-message">
              <div class="basic-info-message-title">
                <div>用户留言</div>
                <div @click="changeTranslateTab(translateTab)">
                  <img style="width: 16px;margin-right: 3px" src="@/assets/images/translate_icon.png" />
                  <span v-if="translateTab == 'zh'">翻译</span>
                  <span v-else>原文</span>
                </div>
              </div>
              <div class="basic-info-message-content">{{ translateTab == 'zh' ? (guestRemarks || '--') : (guestRemarksJa
                ||
                '--') }}</div>
            </div>
            <div class="basic-info-pay">
              <div class="basic-info-pay-title">
                <div></div>付款信息
              </div>
              <div class="basic-info-pay-content">
                <n-grid y-gap="25" :cols="2">
                  <n-gi :span="2">
                    <div class="basic-info-pay-content-item">
                      <div class="basic-info-pay-content-item-title">价格计划名</div>
                      <div class="basic-info-pay-content-item-content">{{ formValue.pricePlanName }}</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="basic-info-pay-content-item">
                      <div class="basic-info-pay-content-item-title">原价房费（{{ formValue.roomNum }}间{{ nightNum }}晚）</div>
                      <div class="basic-info-pay-content-item-content">{{ formValue.orgAmount }}JPY</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="basic-info-pay-content-item">
                      <div class="basic-info-pay-content-item-title">活动优惠金额</div>
                      <div class="basic-info-pay-content-item-content">{{ parseFloat(formValue.totalChangeAmount) < 0 ? '+'+Math.abs(parseFloat(formValue.totalChangeAmount)) : '-'+Math.abs(parseFloat(formValue.totalChangeAmount)) }}JPY</div>
                    </div>
                  </n-gi>
                  <n-gi>
                    <div class="basic-info-pay-content-item">
                      <div class="basic-info-pay-content-item-title">实际付款</div>
                      <div class="basic-info-pay-content-item-content orange">{{ formValue.orderAmount }}JPY</div>
                    </div>
                  </n-gi>
                </n-grid>
              </div>
            </div>
          </div>
          <div class="order-info" v-else-if="nowTab == 2">
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
                <n-gi v-if="formValue.appReservation && formValue.appReservation[0].sourceName != ''">
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单渠道</div>
                    <div>
                      <div class="level-detail-div-item-content1 default">{{ formValue.appReservation[0].sourceName }}
                      </div>
                    </div>
                  </div>
                </n-gi>
                <n-gi v-if="viewtype != 'AIRHOST订单'">
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单状态</div>
                    <div>
                      <div class="level-detail-div-item-content1"
                        :class="getOptionTag(options.order_status, formValue.orderStatus)">{{
                          getOptionLabel(options.order_status, formValue.orderStatus) }}</div>
                    </div>
                  </div>
                </n-gi>
                <n-gi v-if="viewtype == 'AIRHOST订单'">
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">订单状态</div>
                    <div>
                      <div class="level-detail-div-item-content1"
                        :class="getOptionTag(options.status, formValue.appReservation[0].status)">{{
                          getOptionLabel(options.status, formValue.appReservation[0].status) }}</div>
                    </div>
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
            <div class="transaction-info-contain" v-if="formValue.transactionDetail">
              <div class="transaction-info-contain-title">
                <div></div>交易流水
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
            <div class="transaction-refund-info-contain" v-if="formValue.transactionRefundDetail">
              <div class="transaction-refund-info-contain-title">
                <div></div>退款流水
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
            <div class="transaction-refund-info-contain" v-if="formValue.cancelOrder">
              <div class="transaction-refund-info-contain-title">
                <div></div>订单取消记录
              </div>
              <div class="transaction-div">
                <div class="transaction-div-top">
                  <span class="title">订单号：{{ formValue.cancelOrder.cancelOrderSn }}</span>
                </div>
                <div class="transaction-div-item">
                  取消时间：{{ formValue.cancelOrder.cancelAt }}
                  <div>
                    <span><span style="color: #3D3D3D">取消金额：{{ formValue.cancelOrder.cancelAmount }}</span>JPY</span>
                  </div>
                </div>
              </div>
            </div>

            <div class="cancel-policy-info-contain" v-if="formValue.cancelOrder">
              <div class="cancel-policy-info-contain-title">
                <div></div>取消政策
              </div>
              <div class="cancel-policy-info-contain-item">
                <div class="cancel-policy-info-contain-item-list"
                  v-for="(item, index) in formValue.cancelOrder.cancelRate" :key="index">
                  <template v-if="item.mode == 'before'">
                    {{ item.date }}前可免费取消
                  </template>
                  <template v-if="item.mode == 'middle'">
                    {{ item.date }}内取消，将收取订单金额的50%作为取消费用
                  </template>
                  <template v-if="item.mode == 'after'">
                    {{ item.date }}后取消，将收取订单金额的100%作为取消费用
                  </template>
                </div>
              </div>
            </div>
          </div>
          <div style="padding: 25px 20px" v-else-if="nowTab == 3">
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
        </n-spin>
      </n-drawer-content>
    </n-drawer>
    <sendemail ref="sendemailref" />
    <sendnotify ref="sendnotifyref" />
    <sendsms ref="sendsmsref" />
    <hisorderview ref="hisorderviewRef" />
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, reactive, Ref } from 'vue';
import { useMessage } from 'naive-ui';
import { View } from '@/api/pmsAppReservation';
import hisorderview from '@/views/pmsAppReservation/hisorderview.vue';


import { pmsRoomReservationView } from '@/api/pmsRoomReservation';

import { State, options, loadOptions } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import sendemail from '@/views/smjcomm/sendemail.vue';
import sendnotify from '@/views/smjcomm/sendnotify.vue';
import sendsms from '@/views/smjcomm/sendsms.vue';

const nowTab = ref(1)
import { Text } from "@/api/translate";

const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const sendemailref = ref();
const sendnotifyref = ref();
const sendsmsref = ref();
const hisorderviewRef = ref();
const guestRemarks = ref('');
const guestRemarksJa = ref('');
const translateTab = ref('zh')

const formValue: Ref<any> = ref({});
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const state = reactive({
  appReservationitemobj: null,
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});
/*async function loadOptions() {
  options.value = await Dicts({
    types: ['checkin_status', 'order_status', 'status'],
  });
}*/

loadOptions();

//发短信
const sendSmsClick = () => {
  if (formValue.value.guestProfileDetail && formValue.value.guestProfileDetail.phone != "") {
    sendsmsref.value.openModal({
      phone: formValue.value.guestProfileDetail.phone,
      area_no: formValue.value.guestProfileDetail.areaNo,
    });
  } else {
    message.warning('没有号码，不可发送');
  }

};

const nightNum = computed(() => {
  let start = new Date(formValue.value.checkInDate)
  let end = new Date(formValue.value.checkOutDate)
  start.setHours(0, 0, 0, 0)
  end.setHours(0, 0, 0, 0)
  let diffTime = end.getTime() - start.getTime()
  return diffTime / (1000 * 60 * 60 * 24)
});

function copyTextAuto(text) {
  if (navigator.clipboard && navigator.clipboard.writeText) {
    navigator.clipboard.writeText(text)
      .then(() => message.success('复制成功'))
      .catch(err => {
        copyTextFallback(text);
      });
  } else {
    copyTextFallback(text);
  }
}

function copyTextFallback(text) {
  const textarea = document.createElement('textarea');
  textarea.value = text;
  textarea.style.position = 'fixed';  // 避免页面滚动
  textarea.style.opacity = '0';
  document.body.appendChild(textarea);
  textarea.focus();
  textarea.select();

  try {
    const success = document.execCommand('copy');
    console.log(success ? message.success('复制成功') : message.error('复制失败'));
  } catch (err) {
    message.error('复制异常:' + err)
  }

  document.body.removeChild(textarea);
}

function goemail() {
  if (formValue.value.guestProfileDetail && formValue.value.guestProfileDetail.email != "") {
    sendemailref.value.openModal(formValue.value.guestProfileDetail.email);
  } else {
    message.warning('没有邮箱，不可发送');
  }
}
function tomessage() {
  sendnotifyref.value.openModal(formValue.value.memberId);
}
function tohis() {
  hisorderviewRef.value.openModal(formValue.value.memberId);
}
let viewtype = ref('');

function changeTranslateTab(tab: string) {
  if (tab == 'zh') {
    // 翻译
    if (guestRemarks.value && !guestRemarksJa.value) {
      Text({
        text: guestRemarks.value
      }).then((_res) => {
        guestRemarksJa.value = _res.ja
        translateTab.value = 'ja'
      }).catch((err) => {
        translateTab.value = 'ja'
      });
    } else {
      translateTab.value = 'ja'
    }
  } else {
    translateTab.value = 'zh'
  }
}

function openModal(state: State) {
  nowTab.value = 1;
  showModal.value = true;
  loading.value = true;
  viewtype.value = state.viewtype;
  if (state.viewtype == 'AIRHOST订单') {
    pmsRoomReservationView({ out_order_sn: state.outOrderSn, order_sn: state.orderSn })
      .then((res) => {
        formValue.value = res;
        showModal.value = true;
      })
      .finally(() => {
        loading.value = false;
      });
  } else if (state.viewtype == '订单详情') {
    View({ orderSn: state.content })
      .then((res) => {
        formValue.value = res;

        formValue.value.roomNum = formValue.value.appReservation.length
        formValue.value.pricePlanName = parseInt(formValue.value.appReservation[0].pricePlanId) > 0 ? formValue.value.appReservation[0].pricePlanInfo.planShowName : '--';

        let totalChangeAmount = 0
        formValue.value.appReservation.forEach((item) => {
          totalChangeAmount += item.changeAmount
        })
        formValue.value.totalChangeAmount = parseFloat(totalChangeAmount);
        formValue.value.orgAmount = parseFloat(parseFloat(totalChangeAmount) + parseFloat(formValue.value.orderAmount));

        guestRemarks.value = formValue.value?.appReservation[0]?.guestRemarks || ''

        showModal.value = true;
      })
      .finally(() => {
        loading.value = false;
      });
  } else {
    viewtype.value = '订单详情'
    let param: any = {}
    param.orderSn = state.orderSn
    param.outOrderSn = state.outOrderSn
    View(param)
      .then((res) => {
        console.log(res);
        res.appReservation.map(function (item) {
          item.checkinDate = item.checkinDate.split(" ")[0]
          item.checkoutDate = item.checkoutDate.split(" ")[0]
          return item;
        });

        formValue.value = res;
        formValue.value.roomNum = formValue.value.appReservation.length
        formValue.value.pricePlanName = formValue.value.appReservation[0].pricePlanInfo.planShowName;

        let totalChangeAmount = 0
        formValue.value.appReservation.forEach((item) => {
          totalChangeAmount += item.changeAmount
        })
        formValue.value.totalChangeAmount = parseFloat(totalChangeAmount);
        formValue.value.orgAmount = parseFloat(parseFloat(totalChangeAmount) + parseFloat(formValue.value.orderAmount));

        guestRemarks.value = formValue.value?.appReservation[0]?.guestRemarks || ''

        showModal.value = true;
      })
      .finally(() => {
        loading.value = false;
      });
  }
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

.order-detail-info {
  margin: 20px 20px 30px;
  background: #F9FBFC;
  border-radius: 13px;

  &-top {
    display: flex;
    padding: 29px 20px 18px;
    border-bottom: 1px solid #EEEEEE;

    &-info {
      flex: 1;
      padding-top: 7px;

      &-t {
        display: flex;
        align-items: center;
        justify-content: space-between;

        &-l {
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
          font-weight: 500;
          font-size: 18px;
          color: #3D3D3D;
          line-height: 25px;
        }

        &-r {
          display: flex;
          align-items: center;

          &-v1 {
            padding: 0 5px;
            height: 22px;
            line-height: 22px;
            background: #F4F4F5;
            border-radius: 2px;
            font-weight: 400;
            font-size: 14px;
            color: #919399;
          }

          &-v2 {
            margin-left: 3px;
            padding: 0 5px;
            height: 22px;
            border-radius: 2px;
            line-height: 22px;
            font-weight: 400;
            font-size: 14px;

            &.success {
              background: #E3F4EB;
              color: #26A763;
            }

            &.primary {
              background: #ECF5FF;
              color: #3F9EFF;
            }

            &.info {
              background: #ECF5FF;
              color: #3F9EFF;
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

      &-address {
        margin-top: 20px;
        font-weight: 400;
        font-size: 14px;
        color: #979797;
        line-height: 20px;
        text-align: left;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
      }

      &-ordersn {
        margin-top: 10px;

        &-item {
          display: flex;
          align-items: center;
          font-weight: 400;
          font-size: 14px;
          color: #979797;
          line-height: 24px;

          span {
            color: #00B42A;
          }

          div {
            cursor: pointer;
            width: 16px;
            height: 16px;
            margin-left: 7px;
            background-image: url('../../assets/images/new/property_detail_copy1.png');
            background-repeat: no-repeat;
            background-position: 100%;
            background-size: cover;

            &:hover {
              background-image: url('../../assets/images/new/property_detail_copy2.png');
            }
          }
        }
      }

    }
  }

  &-center {
    padding: 20px;
    border-bottom: 1px solid #EEEEEE;
    display: flex;

    &-v {
      &-label {
        font-weight: 500;
        font-size: 12px;
        color: #979797;
        line-height: 17px;
      }

      &-cont {
        margin-top: 9px;
        font-weight: 500;
        font-size: 14px;
        color: #3D3D3D;
        line-height: 20px;
      }

      &-btn {
        margin-top: 10px;
        border: 1px solid #053DC8;
        height: 32px;
        display: flex;
        align-items: center;
        width: fit-content;
        border-radius: 2px;
        padding: 0 10px;
        cursor: pointer;

        .img {
          display: inline-block;
          width: 16px;
          height: 16px;
          margin-right: 4px;
          background-size: cover;
        }

        .btn-text {
          display: inline-block;
          font-weight: 400;
          font-size: 14px;
          color: #053DC8;
          line-height: 32px;
        }

        &.btn1 {
          .img {
            background-image: url("@/assets/images/icon_hotel_order_btn1.png");
          }
        }

        &.btn2 {
          .img {
            background-image: url("@/assets/images/icon_hotel_order_btn2.png");
          }
        }

        &.btn3 {
          .img {
            background-image: url("@/assets/images/icon_hotel_order_btn3.png");
          }
        }

        &:hover {
          background: #053DC8;

          &.btn1 {
            .img {
              background-image: url("@/assets/images/icon_hotel_order_btn1_active.png");
            }
          }

          &.btn2 {
            .img {
              background-image: url("@/assets/images/icon_hotel_order_btn2_active.png");
            }
          }

          &.btn3 {
            .img {
              background-image: url("@/assets/images/icon_hotel_order_btn3_active.png");
            }
          }

          .btn-text {
            color: #FFFFFF;
          }
        }
      }
    }
  }

  &-bottom {
    padding: 20px 20px 24px;
    display: flex;

    &-v {
      div {
        &:first-child {
          font-weight: 500;
          font-size: 12px;
          color: #979797;
          line-height: 17px;
        }

        &:last-child {
          margin-top: 9px;
          font-weight: 500;
          font-size: 14px;
          color: #3D3D3D;
          line-height: 20px;
        }
      }
    }
  }
}

.order-detail-tabs {
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

.basic-info {
  padding: 25px 20px;

  &-room {
    &-title {
      display: flex;
      align-items: center;
      font-weight: 600;
      font-size: 16px;
      color: #3D3D3D;
      line-height: 22px;
      margin-bottom: 15px;

      div {
        width: 6px;
        height: 15px;
        background: #053DC8;
        margin-right: 5px;
      }
    }

    &-list {
      border-bottom: 1px solid #F4F4F4;
      padding: 10px 0 15px;
      display: flex;

      &-info {
        flex: 1;
        padding-top: 4px;

        &-t {
          display: flex;
          align-items: center;

          &-tag {
            padding: 0 5px;
            height: 22px;
            line-height: 22px;
            border-radius: 2px;
            font-weight: 400;
            font-size: 14px;

            &.success {
              background: #E3F4EB;
              color: #26A763;
            }

            &.primary {
              background: #ECF5FF;
              color: #3F9EFF;
            }

            &.info {
              background: #ECF5FF;
              color: #3F9EFF;
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

          &-name {
            margin-left: 5px;
            font-weight: 500;
            font-size: 16px;
            color: #3D3D3D;
            line-height: 22px;
          }
        }

        &-c {
          margin-top: 4px;
          font-weight: 400;
          font-size: 12px;
          color: #979797;
          line-height: 17px;
        }

        &-b {
          margin-top: 9px;
          display: flex;
          align-items: center;

          div {
            &:first-child {
              font-weight: 400;
              font-size: 14px;
              color: #979797;
              line-height: 20px;
              margin-right: 65px;
            }

            &:last-child {
              display: flex;
              align-items: center;
              font-weight: 400;
              font-size: 14px;
              color: #979797;
              line-height: 20px;

              span {
                font-weight: 400;
                font-size: 14px;
                color: #3D3D3D;
                line-height: 20px;
              }
            }
          }
        }
      }
    }
  }

  &-message {
    padding: 25px 0 30px;

    &-title {
      display: flex;
      align-items: center;

      div {
        &:first-child {
          font-weight: 400;
          font-size: 12px;
          color: #707070;
          line-height: 17px;
          margin-right: 8px;
        }

        &:last-child {
          display: flex;
          align-items: center;
          cursor: pointer;

          span {
            font-weight: 400;
            font-size: 12px;
            color: #1664FF;
            line-height: 17px;
          }
        }
      }
    }

    &-content {
      margin-top: 8px;
      font-weight: 600;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
    }
  }

  &-pay {
    &-title {
      display: flex;
      align-items: center;
      font-weight: 600;
      font-size: 16px;
      color: #3D3D3D;
      line-height: 22px;
      margin-bottom: 15px;

      div {
        width: 6px;
        height: 15px;
        background: #053DC8;
        margin-right: 5px;
      }
    }

    &-content {
      margin-top: 25px;

      &-item {
        &-title {
          font-weight: 400;
          font-size: 12px;
          color: #707070;
          line-height: 17px;
          margin-bottom: 8px;
        }

        &-content {
          font-weight: 600;
          font-size: 14px;
          color: #3D3D3D;
          line-height: 20px;

          &.orange {
            color: #FF6F00;
          }
        }
      }
    }
  }
}

.order-info {
  padding: 25px 20px;
}

.level-detail-div {
  margin-bottom: 25px;

  &-title {
    display: flex;
    align-items: center;
    font-weight: 600;
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
      font-weight: 600;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
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

.transaction-info-contain {
  &-title {
    display: flex;
    align-items: center;
    font-weight: 600;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 15px;

    div {
      width: 6px;
      height: 15px;
      background: #053DC8;
      margin-right: 5px;
    }
  }
}

.transaction-refund-info-contain {
  margin-top: 25px;

  &-title {
    display: flex;
    align-items: center;
    font-weight: 600;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 15px;

    div {
      width: 6px;
      height: 15px;
      background: #053DC8;
      margin-right: 5px;
    }
  }
}

.transaction-div {
  background: #F9FBFC;
  padding: 20px;
  margin-top: 20px;
  border-radius: 2px;

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
      font-size: 14px;
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

.cancel-policy-info-contain {
  margin-top: 25px;

  &-title {
    display: flex;
    align-items: center;
    font-weight: 600;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 15px;

    div {
      width: 6px;
      height: 15px;
      background: #053DC8;
      margin-right: 5px;
    }
  }

  &-item {
    margin-top: 25px;

    &-top {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 8px;

      div {
        font-weight: 600;
        font-size: 14px;
        color: #3D3D3D;
        line-height: 20px;
      }
    }

    &-list {
      margin-top: 8px;
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;

      &:first-child {
        margin-top: 0;
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
