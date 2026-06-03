<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">确认订单</div>
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
            <div class="dispatch-form-item-title">预定日期<span>*</span></div>
            <div class="dispatch-form-item-textarea">
              <n-date-picker placeholder="请选择预定日期" v-model:formatted-value="formValue.bookDate" value-format="yyyy-MM-dd" type="date" style="width: 300px"/>
            </div>
          </div>
          <div class="dispatch-form-item">
            <div class="dispatch-form-item-title">预定时间<span>*</span></div>
            <div class="dispatch-form-item-textarea">
              <n-time-picker placeholder="请选择" format="HH:mm" v-model:formatted-value="formValue.bookTime" style="width: 300px"/>
            </div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import { ConfirmAgree } from '@/api/foodOrder';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { options, State } from './model';
import {adaModalWidth, getOptionLabel} from '@/utils/hotgo';

const emit = defineEmits(['reloadInfo']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  bookDate: null,
  bookTime: null,
});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(560);
});

function openModal(state: State) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = state.id;
  formValue.value.bookDate = state.bookDate;
  formValue.value.bookTime = state.bookTime;
  loading.value = false;
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  ConfirmAgree({
    id: formValue.value.id,
    bookDate: formValue.value.bookDate,
    bookTime: formValue.value.bookTime
  }).then((_res) => {
    message.success('确认成功');
    formBtnLoading.value = false;
    setTimeout(() => {
      closeForm();
      emit('reloadInfo');
    });
  }).catch((err) => {
    formBtnLoading.value = false;
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formValue.value.id = 0;
  formValue.value.bookDate = null;
  formValue.value.bookTime = null;
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


