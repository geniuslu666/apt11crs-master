<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">拒绝原因</div>
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
            <div class="dispatch-form-item-title">拒绝原因<span>*</span></div>
            <div class="dispatch-form-item-textarea">
              <n-input
                type="textarea"
                placeholder="请输入拒绝原因"
                autosize style="min-height: 100px"
                v-model:value="formValue.desc"
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
import { ConfirmDisagree } from '@/api/foodOrder';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { options } from './model';
import {adaModalWidth, getOptionLabel} from '@/utils/hotgo';

const emit = defineEmits(['reloadInfo']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  bookingStatus: '',
  desc: '',
});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(560);
});


function openModal(id, status) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = id;
  formValue.value.bookingStatus = status;
  loading.value = false;
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(!formValue.value.desc){
    formBtnLoading.value = false;
    message.error('请输入拒绝原因');
    return false;
  }
  ConfirmDisagree({
    id: formValue.value.id,
    bookingStatus: formValue.value.bookingStatus,
    bookingDisagreeReason: formValue.value.desc
  }).then((_res) => {
    message.success('操作成功');
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
  formValue.value.bookingStatus = '';
  formValue.value.desc = '';
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


