<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">订单详情-编号{{ formValue.orderSn }}</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div style="padding: 0 20px">
            <div class="car-order-detail-address">
              <div class="car-order-detail-address-item">
                <div class="car-order-detail-address-item-d1">柜门号</div>
                <div class="car-order-detail-address-item-d2">{{ formValue.boxNo }}({{ formValue.boxTypeName }})</div>
              </div>


              <div class="car-order-detail-address-item">
                <div class="car-order-detail-address-item-d1">订单状态</div>
                <div class="car-order-detail-address-item-d2" :class="getOptionTag(options.cabinet_order_status, formValue.orderStatus)">{{ getOptionLabel(options.cabinet_order_status, formValue.orderStatus) }}</div>

              </div>

            </div>
            <div class="car-order-detail-time">
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">网点名称</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.mchBranchName }}</div>
              </div>
              <div class="car-order-detail-time-item">
                <div class="car-order-detail-time-item-d1">使用时长</div>
                <div class="car-order-detail-time-item-d2">{{ formValue.mchBranchName }}</div>
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

          <div class="car-order-detail-tabs">
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
          </div>
          <div style="padding: 30px 20px" v-if="nowTab == 1" >
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
                  <n-gi>
                    <div class="level-detail-div-item">
                      <div class="level-detail-div-item-title">支付状态</div>
                      <div>
                        <div class="level-detail-div-item-content1" :class="getOptionTag(options.spa_order_pay_status, formValue.payStatus)">{{ getOptionLabel(options.spa_order_pay_status, formValue.payStatus) }}</div>
                      </div>
                    </div>
                  </n-gi>
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
                <n-tooltip trigger="hover" :arrow-point-to-center="true">
                  <template #trigger>
                    <div class="car-order-detail-address-item-d3">
                      超时不满一小时按一小时计费
                    </div>
                  </template>
                  超时不满一小时按一小时计费
                </n-tooltip>
              </div>

              <div class="order-table-header">
                <div style="padding-left: 3.3%;width: 25%">柜门类型</div>
                <div style="width: 35%">基础费用</div>
                <div style="width: 35%">超时费用</div>
              </div>
              <div class="order-table-body">
                <div class="order-table-body-item" v-for="(item, index) in formValue.boxTypeJson" :key="index">
                  <div style="padding-left: 3.3%;width: 20%">{{ item.name }}</div>
                  <div style="width: 21.8%">{{ item.price }}/小时</div>
                  <div style="width: 27.9%">{{ item.price }}/小时</div>
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
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {View} from '@/api/cabinetOrder';
import {State, newState} from '@/views/cabinetOrder/model';
import {adaModalWidth, getOptionLabel, getOptionTag, Option} from '@/utils/hotgo';
import {Dicts} from "@/api/dict/dict";
import {usePermission} from "@/hooks/web/usePermission";

const { hasPermission } = usePermission();
const dialog = useDialog();
const emit = defineEmits(['reloadTable']);
const nowTab = ref(1)
const formBtnLoading = ref(false);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});

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
  &-type{
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
  }
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


