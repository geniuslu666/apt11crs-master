<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content :title="viewtype" closable>
        <div class="c333" style="font-size: 13px">
          <!-- 历史订单入口 -->
          <div
            style="
              position: absolute;
              padding: 5px;
              right: 55px;
              top: 11px;
              text-align: center;
              color: rgb(30, 148, 252);
              cursor: pointer;
            "
            v-if="viewtype == '最新订单'"
            @click="tohis()"
          >
            历史订单
            <n-icon color="#1E94FCFF" size="13">
              <ChevronForward />
            </n-icon>
          </div>

          <!-- 客人信息 -->
          <div style="border-radius: 5px; margin: 10px 20px; background-color: #f9fbfcff">
            <div class="lh-2" style="padding: 0; border-bottom: 1px solid #eeeeeeff">
              <div class="flex-row p-15">
                <n-image
                  class="carousel-img"
                  :src="formValue.propertyDetail.cover"
                  width="100%"
                  object-fit="cover"
                  style="height: 85px; width: 60px; margin-right: 10px"
                />
                <div class="flex-item">
                  <div class="fw f18">{{ formValue.propertyDetail.name }}</div>
                  <div class="c999">{{ formValue.propertyDetail.address }}</div>
                  <div class="">
                    订单编号：{{ formValue.orderSn }}
                    <span class="fw csuccess">{{ formValue.outOrderSn }}</span>
                  </div>
                </div>
                <div v-if="viewtype != 'AIRHOST订单'">
                  <n-tag
                    :type="getOptionTag(options.order_status, formValue.orderStatus)"
                    size="small"
                    class="min-left-space"
                  >
                    {{ getOptionLabel(options.order_status, formValue.orderStatus) }}
                  </n-tag>
                </div>
                <div v-else>
                  <n-tag
                    v-if="formValue.appReservation"
                    :type="getOptionTag(options.status, formValue.appReservation[0].status)"
                    size="small"
                    class="min-left-space"
                  >
                    {{ getOptionLabel(options.status, formValue.appReservation[0].status) }}
                  </n-tag>
                </div>
              </div>
              <div class="flex-row p-15" style="border-top: 1px solid #eeeeeeff; margin-top: 20px">
                <div style="width: 120px">
                  <div class="c999 f12">预定者</div>
                  <div> {{ formValue.guestProfileDetail.fullName }}</div>
                  <div>
                    <n-button icon-placement="left" type="primary" size="tiny" @click="tomessage()">
                      <template #icon>
                        <n-icon>
                          <ChatbubblesOutline />
                        </n-icon>
                      </template>
                      发送站内信
                    </n-button>
                  </div>
                </div>
                <div style="width: 150px">
                  <div class="c999 f12">Tel:</div>
                  <div>
                    {{ formValue.guestProfileDetail.areaNo }} -
                    {{ formValue.guestProfileDetail.phone }}</div
                  >
                  <div>
                    <n-button
                      icon-placement="left"
                      type="primary"
                      size="tiny"
                      @click="sendSmsClick()"
                    >
                      <template #icon>
                        <n-icon>
                          <LogoWhatsapp />
                        </n-icon>
                      </template>
                      发短信
                    </n-button>
                  </div>
                </div>
                <div class="flex-item">
                  <div class="c999 f12">E-mail:</div>
                  <div>{{ formValue.guestProfileDetail.email }}</div>
                  <div>
                    <n-button icon-placement="left" type="primary" size="tiny" @click="goemail()">
                      <template #icon>
                        <n-icon>
                          <MailUnreadOutline />
                        </n-icon>
                      </template>
                      发邮件
                    </n-button>
                  </div>
                </div>
              </div>
            </div>
            <div class="p-10 lh-2">
              <div class="flex-row">
                <div class="flex-item">
                  <n-icon size="15" color="#053DC8FF" style="margin-right: 5px">
                    <CalendarNumberOutline />
                  </n-icon>
                  <span class="c999">日期：</span>
                  <span class="c333">{{ formValue.checkInDate }}</span>
                  <span class="c999">&nbsp;→&nbsp;</span>
                  <span class="c333">{{ formValue.checkOutDate }}</span>
                </div>
                <div class="text-r f16 c333">
                  ￥{{ formValue.orderAmount }} <span class="c999 f14">JPY</span>
                </div>
              </div>
            </div>
          </div>
          <!-- f订单信息 -->
          <!-- 房间信息 -->
          <div
            style="
              padding: 20px 0px;
              border-top: 1px solid #eeeeeeff;
              border-bottom: 1px solid #eeeeeeff;
              margin-top: 30px;
            "
            class="lh-2"
            v-if="formValue.guestProfileDetail"
          >
            <div style="margin-bottom: 10px">
              <div class="f16">房间信息</div>
              <div
                v-for="(item, index) in formValue.appReservation"
                :key="index"
                style="
                  margin: 10px 0;
                  font-size: 12px;
                  padding: 10px;
                  border-radius: 5px;
                  line-height: 22px;
                "
              >
                <div class="flex-row">
                  <n-image
                    :src="item.roomTypeDetail.cover"
                    :object-fit="'cover'"
                    width="65px"
                    style="height: 65px; margin-right: 10px"
                  />
                  <div class="flex-item">
                    <div class="flex-row">
                      <div class="c333 f14">
                        {{ item.roomTypeDetail.name }}
                        <span class="c333 f14 ml-2 fw" v-if="item.roomUnitDetail"
                          >#{{ item.roomUnitDetail.roomNo }}</span
                        ></div
                      >
                      <div class="flex-item text-r c999 f12">
                        {{ item.roomTypeDetail.createAt }}
                      </div>
                    </div>
                    <div class="f12 c999">
                      {{ item.roomTypeDetail.checkinAt }} {{ item.roomTypeDetail.checkoutAt }}</div
                    >
                    <div class="flex-row">
                      <div>
                        <span class="c999 f12 mr-1">入住状态：</span>
                        <n-tag
                          :type="getOptionTag(options.checkin_status, item.checkinStatus)"
                          size="small"
                          class="min-left-space"
                        >
                          {{ getOptionLabel(options.checkin_status, item.checkinStatus) }}
                        </n-tag>
                      </div>
                      <div class="flex-item">
                        <div class="flex-row">
                          <div class="c999 f12 ml-16 mr-1">客人信息：</div>
                          <img
                            src="@/assets/images/cr.png"
                            style="width: 16px; height: 16px; margin-right: 5px"
                          />
                          <div class="c333" style="margin-right: 5px"
                            >成人：{{ item.adultCount }}</div
                          >
                          <img
                            src="@/assets/images/ye.png"
                            style="width: 16px; height: 16px; margin-right: 5px"
                          />

                          <div class="c333">儿童：{{ item.childCount }}</div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- 交易流水 -->
          <div style="padding: 20px 0" class="lh-2" v-if="formValue.transactionDetail">
            <div class="f16">交易流水</div>
            <div class="payitem" v-for="(item, index) in formValue.transactionDetail" :key="index">
              <div class="flex-row">
                <div> <span class="c333">支付流水号：</span>{{ item.transactionSn }}</div>
                <div class="flex-item text-r">
                  <span v-if="item.payStatus == 'WAIT'" class="cblue">等待支付</span>
                  <span v-else-if="item.payStatus == 'DONE'" class="csuccess">完成支付</span>
                  <span v-else-if="item.payStatus == 'CANCEL'" class="c999">取消支付</span>
                </div>
              </div>
              <div><span class="c333">第三方支付流水号：</span>{{ item.paymentRequestId }}</div>
              <div class="flex-row">
                <div><span class="c333">创建时间：</span>{{ item.createdAt }}</div>
              </div>
              <div class="flex-row">
                <div><span class="c333">过期时间：</span>{{ item.expiredTime }}</div>

                <div class="flex-item text-r">
                  {{ item.amount }} <span class="c999">{{ item.priceCurrency }} </span>
                  /
                  <span class="cwarning" style="font-size: 14px">{{ item.payAmount }}</span>
                  <span class="c999 f12">{{ item.priceCurrency }}</span>
                </div>
              </div>
              <div class="flex-row">
                <div>
                  <span class="c333">支付平台：</span>
                  <span v-if="item.payChannel == 'SYSTEM'" class="cblue">系统积分</span>
                  <span v-else-if="item.payChannel == 'PAYCLOUD'" class="cblue"
                    >paycloud第三方支付平台
                  </span>
                </div>
                <div class="flex-item text-r">
                  <span class="c999">
                    【总退款金额：{{ item.refundAmount }}
                    <span class="c999 f12">{{ item.priceCurrency }} </span>
                    】
                  </span>
                </div>
              </div>
            </div>
          </div>
          <!-- 退款流水 -->
          <div style="padding: 20px 0" class="lh-2" v-if="formValue.transactionRefundDetail">
            <div class="f16">退款流水</div>
            <div
              class="payitem"
              v-for="(item, index) in formValue.transactionRefundDetail"
              :key="index"
            >
              <div>
                <span class="c333">退款单号: </span>
                <span>{{ item.refundSn }}</span>
              </div>
              <div class="flex-row">
                <div>
                  <span class="c333">退款方式： </span>
                  <span>{{ item.refundType }}</span>
                </div>
                <div>
                  <span class="c333">退款时间： </span>
                  <span>{{ item.refundTime }}</span>
                </div>
                <div class="text-r flex-item"
                  >{{ item.refundAmount }}<span class="c999 f12">JPY</span></div
                >
              </div>
            </div>
          </div>

          <div style="padding: 20px 0" class="lh-2" v-if="formValue.cancelOrder">
            <div class="f16">取消订单</div>
            <div class="payitem">
              <div>
                <span class="c333">订单号： </span>
                <span>{{ formValue.cancelOrder.cancelOrderSn }}</span>
              </div>
              <div>
                <span class="c333">取消时间： </span>
                <span>{{ formValue.cancelOrder.cancelAt }}</span>
              </div>
              <div>
                <span class="c333">取消金额： </span>
                <span class="c333">{{ formValue.cancelOrder.cancelAmount }}</span>
                <span class="c999 ml-1">JPY</span>
              </div>
            </div>
            <a-timeline class="mt-3">
              <a-timeline-item
                v-for="(item, index) in formValue.cancelOrder.cancelRate"
                :key="index"
                :color="index == formValue.cancelOrder.cancelRate.length - 1 ? 'blue' : 'gray'"
              >
                <div
                  :class="index == formValue.cancelOrder.cancelRate.length - 1 ? 'cblue' : 'c999'"
                  >{{ item.name }}</div
                >
                <div
                  ><span>{{ item.date }}</span>
                  <span v-if="item.mode == 'before'" class="ml-1">前可免费取消</span>
                  <span v-if="item.mode == 'middle'" class="ml-1"
                    >内取消，将收取订单金额的50%作为取消费用</span
                  >
                  <span v-if="item.mode == 'after'" class="ml-1"
                    >后取消，将收取订单金额的100%作为取消费用</span
                  >
                </div>
              </a-timeline-item>
            </a-timeline>
          </div>

          <div style="padding: 20px 0" class="lh-2" v-if="formValue.referrerDetail">
            <div class="flex-row">
              <div class="flex-item f16">
                <span v-if="formValue.referrerDetail.rebateMode == 'MEMBER'">会员</span>
                <span v-if="formValue.referrerDetail.rebateMode == 'STAFF'">员工</span>
                <span v-if="formValue.referrerDetail.rebateMode == 'CHANNEL'">渠道</span>
                分销
              </div>
              <div class="text-r">
                <span>{{ formValue.rebateAmount }}</span>
                <span class="c999">JPY</span>/
                <span class="cwarning mr-5">{{ formValue.rebateRate }}%</span>
                <a-tag color="blue" v-if="formValue.rebateStatus == 'WAIT'">等待中</a-tag>
                <a-tag color="green" v-if="formValue.rebateStatus == 'SUCCESS'">已成功</a-tag>
                <a-tag color="red" v-if="formValue.rebateStatus == 'FAIL'">失败</a-tag>
              </div>
            </div>
            <div class="fle-row" v-if="formValue.referrerDetail">
              <div class="payitem" v-if="formValue.referrerDetail.channel">
                <span class="c999">推荐人信息：</span>{{ formValue.referrerDetail.channel.name }}

                <span>tel:{{ formValue.referrerDetail.channel.phone }}</span>
                <span>mail:{{ formValue.referrerDetail.channel.email }}</span>
              </div>
            </div>
          </div>
        </div>
      </n-drawer-content>
    </n-drawer>
    <sendemail ref="sendemailref" />
    <sendnotify ref="sendnotifyref" />
    <sendsms ref="sendsmsref" />
    <hisorderview ref="hisorderviewRef" />
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive } from 'vue';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/pmsAppReservation';
  import hisorderview from '@/views/pmsAppReservation/hisorderview.vue';


  import { pmsRoomReservationView } from '@/api/pmsRoomReservation';

  import { State, options } from '@/views/pmsAppReservation/model';
  import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import sendemail from '@/views/smjcomm/sendemail.vue';
  import sendnotify from '@/views/smjcomm/sendnotify.vue';
  import sendsms from '@/views/smjcomm/sendsms.vue';

  import {
    CalendarNumberOutline,
    ChatbubblesOutline,
    LogoWhatsapp,
    ChevronForward,
    MailUnreadOutline,
  } from '@vicons/ionicons5';
  import { Dicts } from '@/api/dict/dict';
  const message = useMessage();
  const loading = ref(false);
  const showModal = ref(false);
  const sendemailref = ref();
  const sendnotifyref = ref();
  const sendsmsref = ref();
  const hisorderviewRef = ref();

  const formValue = ref({});
  const dialogWidth = computed(() => {
    return adaModalWidth(620);
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
  async function loadOptions() {
    options.value = await Dicts({
      types: ['checkin_status', 'order_status', 'status'],
    });
  }

  loadOptions();

  //发短信
  const sendSmsClick = () => {
    sendsmsref.value.openModal({
      phone: formValue.value.guestProfileDetail.phone,
      area_no: formValue.value.guestProfileDetail.areaNo,
    });
  };
  function goemail() {
    sendemailref.value.openModal(formValue.value.guestProfileDetail.email);
  }
  function tomessage() {
    sendnotifyref.value.openModal(formValue.value.memberId);
  }
  function tohis() {
    hisorderviewRef.value.openModal(formValue.value.memberId);
  }
  let viewtype = ref('');
  function openModal(state: State) {

    console.log(state)
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
    } else if (state.viewtype == '最新订单') {
      View({ orderSn: state.content.orderSn })
        .then((res) => {
          formValue.value = res;

          showModal.value = true;
        })
        .finally(() => {
          loading.value = false;
        });
    } else {

      View({ orderSn: state.orderSn })
        .then((res) => {
          formValue.value = res;


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
</style>
