<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 20px 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">支付流水详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              支付信息
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
                  <div class="level-detail-div-item-title">支付流水号</div>
                  <div class="level-detail-div-item-content">{{ formValue.transactionSn ? formValue.transactionSn : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">第三方支付流水号</div>
                  <div class="level-detail-div-item-content">{{ formValue.paymentRequestId ? formValue.paymentRequestId : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">总金额</div>
                  <div class="level-detail-div-item-content">{{ formValue.amount }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">币种</div>
                  <div class="level-detail-div-item-content">{{ formValue.priceCurrency ? formValue.priceCurrency : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">支付金额</div>
                  <div class="level-detail-div-item-content">{{ formValue.payAmount }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">支付状态</div>
                  <div>
                    <div class="level-detail-div-item-content1" :class="getOptionTag(options.pay_status, formValue?.payStatus)">{{ getOptionLabel(options.pay_status, formValue?.payStatus) }}</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">支付时间</div>
                  <div class="level-detail-div-item-content">{{ formValue.payTime ? formValue.payTime : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">支付渠道</div>
                  <div>
                    <div class="level-detail-div-item-content1" :class="getOptionTag(options.pay_channel, formValue?.payChannel)">{{ getOptionLabel(options.pay_channel, formValue?.payChannel) }}</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">支付方式</div>
                  <div>
                    <div class="level-detail-div-item-content1" :class="getOptionTag(options.pay_type, formValue?.payType)">{{ getOptionLabel(options.pay_type, formValue?.payType) }}</div>
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              积分信息
            </div>
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">场景积分抵扣比例</div>
                  <div class="level-detail-div-item-content">{{ formValue.scenePayRate }}%</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">积分汇率</div>
                  <div class="level-detail-div-item-content">{{ formValue.exchangeRate }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              退款信息
            </div>
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">退款金额</div>
                  <div class="level-detail-div-item-content">{{ formValue.refundAmount }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">退款状态</div>
                  <div>
                    <div v-if="formValue.refundStatus=='WAIT'" class="level-detail-div-item-content1 default">未退款</div>
                    <div v-else-if="formValue.refundStatus=='PART'" class="level-detail-div-item-content1 warning">部分退款</div>
                    <div v-else class="level-detail-div-item-content1 error">全额退款</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">过期时间</div>
                  <div class="level-detail-div-item-content">{{ formValue.expiredTime }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              其他
            </div>
            <n-grid y-gap="25" :cols="1">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">支付参数</div>
                  <div class="level-detail-div-item-content" v-if="formValue.payParams" v-html="formValue.payParams"></div>
                  <div class="level-detail-div-item-content" v-else>--</div>
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
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { View } from '@/api/pmsTransaction';
import { State, newState, options } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import { getFileExt } from '@/utils/urlUtils';

const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(580);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});

//下载
function download(url: string) {
  window.open(url);
}

function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
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
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;
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
</style>


