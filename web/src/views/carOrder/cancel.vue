<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">退款</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;">
            取消
          </n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;margin-left: 10px">
            确认
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="dispatch-form-item">
            <div class="dispatch-form-item-title">退款方式<span>*</span></div>
            <div class="dispatch-form-item-radio">
              <a-radio-group v-model:value="formValue.refundType" name="refundType">
                <a-radio :value="1">仅退款</a-radio>
                <a-radio :value="2">退款并取消订单</a-radio>
              </a-radio-group>
            </div>
            <div class="dispatch-form-item-tips">
              <div>仅退款：订单会处于部分退款状态，订单状态不会改变。</div>
              <div>退款并取消订单：退款后，订单会处于已取消状态。</div>
            </div>
          </div>
          <div class="dispatch-form-item" style="margin-top: 25px">
            <div class="dispatch-form-item-title">退款/取消原因<span>*</span></div>
            <div class="dispatch-form-item-textarea">
              <n-input
                type="textarea"
                placeholder="请输入退款/取消原因"
                autosize style="min-height: 100px"
                v-model:value="formValue.adminCancelReason"
              />
            </div>
          </div>
          <div class="dispatch-form-item" style="margin-top: 25px">
            <div class="dispatch-form-item-title">退款金额<span>*</span></div>
            <div class="dispatch-form-item-input">
                <n-input-group>
                  <n-input-number placeholder="请输入" :show-button="false" :min="0" :max="formValue.maxRefundMoney" :precision="0" v-model:value="formValue.refundMoney" style="width: 245px" />
                  <n-input-group-label>JPY</n-input-group-label>
                </n-input-group>
            </div>
            <div class="dispatch-form-item-tips">
              <div style="color: red">订单金额为{{formValue.orderAmount}}JPY，最多可退款{{formValue.maxRefundMoney}}JPY</div>
            </div>
          </div>

          <div class="dispatch-form-item" style="margin-top: 25px">
            <div class="dispatch-form-item-title">图片</div>
            <div class="dispatch-form-item-input">
              <FileChooser1 v-model:value="formValue.images" :maxNumber="1"
                            fileType="default"/>
            </div>
          </div>

        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { RefundOrder } from '@/api/carOrder';
import {useMessage} from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const emit = defineEmits(['reloadInfo']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(560);
});
const formValue = ref({
  adminCancelReason: "",
  id: 0,
  refundType: 1,
  refundMoney: null,
  orderAmount: 0,
  maxRefundMoney: 0,
  images: '',
})

function openModal(state) {
  showModal.value = true;

  formValue.value.id = state.id
  formValue.value.orderAmount = state.orderAmount
  formValue.value.maxRefundMoney = parseFloat(state.orderAmount - state.couponAmount - state.refundAmount);
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(!formValue.value.adminCancelReason){
    formBtnLoading.value = false;
    message.error('请输入取消原因');
    return false;
  }
  if(!formValue.value.refundMoney || formValue.value.refundMoney < 0 || formValue.value.refundMoney > formValue.value.maxRefundMoney){
    formBtnLoading.value = false;
    message.error('退款金额不正确');
    return false;
  }
  RefundOrder({
    id: formValue.value.id,
    refundType: formValue.value.refundType,
    refundMoney: formValue.value.refundMoney,
    adminCancelReason: formValue.value.adminCancelReason,
    images: formValue.value.images,
  }).then((_res) => {
    message.success('操作成功');
    formBtnLoading.value = false;
    setTimeout(() => {
      closeForm();
      emit('reloadInfo');
    });
  }).catch((err)=>{
    formBtnLoading.value = false;
  });
}

function closeForm() {
  showModal.value = false;
  formValue.value.adminCancelReason = ""
  formValue.value.id = 0
  formValue.value.refundType = 1
  formValue.value.refundMoney = null
  formValue.value.orderAmount = 0
  formValue.value.maxRefundMoney = 0
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less">
.dispatch-form-item{
  &-title{
    font-weight: 400;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
    span{
      color: #F73314;
    }
  }
  &-radio{
    margin: 15px 0 10px;
  }
  &-tips{
    div{
      font-weight: 400;
      font-size: 12px;
      color: #9EA4AA;
      line-height: 17px;
      margin-bottom: 5px;
      &:last-child{
        margin-bottom: 0;
      }
    }
  }
  &-textarea{
    margin-top: 8px;
    textarea{
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
    }
  }
  &-input{
    margin-top: 8px;
    margin-bottom: 10px;
  }
}
</style>


