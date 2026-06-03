<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">手动核销</div>
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
            <div class="dispatch-form-item-title">券号</div>
            <div class="dispatch-form-item-textarea">{{formValue.couponNo}}</div>
          </div>
          <div class="dispatch-form-item">
            <div class="dispatch-form-item-title">核销时间<span>*</span></div>
            <div class="dispatch-form-item-textarea">
              <n-date-picker  placeholder="请选择核销时间" v-model:formatted-value="formValue.verifyTime" type="datetime" clearable style="width: 300px" />
            </div>
          </div>
          <div class="dispatch-form-item">
            <div class="dispatch-form-item-title">核销门店<span>*</span></div>
            <div class="dispatch-form-item-textarea">
              <n-select
                placeholder="请选择核销门店"
                v-model:value="formValue.storeId"
                :options="storeList"
                label-field="storeName"
                value-field="id"
                clearable
                filterable
                style="width: 400px"
              />
            </div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import { ManualVerify } from '@/api/thMemberCoupon';
import { useMessage } from 'naive-ui';
import { options, State } from './model';
import {adaModalWidth} from '@/utils/hotgo';
import { All as storeAll } from '@/api/thMchStore';

const emit = defineEmits(['reloadInfo']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  storeId: null,
  verifyTime: null,
  couponNo: '',
});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(560);
});

const storeList = ref([]);

async function openModal(state: State) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = state.id;
  formValue.value.couponNo = state.couponNo;
  await loadStoreList(state.couponId);
  loading.value = false;
}

async function loadStoreList(couponId){
  let storeArrList = await storeAll({
    couponId: couponId
  })
  storeList.value = storeArrList.list
}

function confirmForm(e) {
  e.preventDefault();
  if (!formValue.value.verifyTime) {
    message.error('请选择核销时间');
    return;
   }
  if (!formValue.value.storeId) {
    message.error('请选择核销门店');
    return;
  }
  formBtnLoading.value = true;
  ManualVerify({
    id: formValue.value.id,
    verifyTime: formValue.value.verifyTime,
    storeId: formValue.value.storeId,
  }).then((_res) => {
    message.success('确认成功');
    formBtnLoading.value = false;
    setTimeout(() => {
      closeForm();
      emit('reloadTable');
    });
  }).catch((err) => {
    formBtnLoading.value = false;
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formValue.value.id = 0;
  formValue.value.couponNo = '';
  formValue.value.storeId = null;
  formValue.value.verifyTime = null;
}

defineExpose({
  openModal,
});
</script>

<style lang="less">
.dispatch-form-item{
  margin-bottom: 24px;
  &-title{
    font-weight: 400;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
    span{
      color: #F73314;
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
}
</style>


