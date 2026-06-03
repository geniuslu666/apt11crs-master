<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">异常处理</div>
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
          <div class="dispatch-form-item" style="margin-top: 35px">
            <div class="dispatch-form-item-title">原因</div>
            <div class="dispatch-form-item-textarea">
              <n-input
                type="textarea"
                placeholder="请输入原因"
                autosize style="min-height: 100px"
                v-model:value="formValue.abnormalReason"
              />
            </div>
          </div>
          <div class="dispatch-form-item" style="margin-top: 20px">
            <div class="dispatch-form-item-title">操作人</div>
            <div class="dispatch-form-item-user">{{ userStore.username }}</div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import {Abnormal} from '@/api/spaOrder';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {useUserStore} from "@/store/modules/user";

const emit = defineEmits(['reloadInfo']);
const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(560);
});

const message = useMessage();
const userStore = useUserStore();
const formValue = ref({
  id: 0,
  abnormalReason: '',
  abnormalOperatorId: 0
});
const formBtnLoading = ref(false);

function openModal(id) {
  showModal.value = true;
  // 编辑
  loading.value = true;
  formValue.value.id = id;
  loading.value = false;
}


function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  if(!formValue.value.abnormalReason){
    formBtnLoading.value = false;
    message.error('请填写原因');
    return false;
  }
  formValue.value.abnormalOperatorId = userStore.info ? userStore.info.id : 0;
  Abnormal(formValue.value).then((_res) => {
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
  formValue.value.abnormalReason = '';
  formValue.value.abnormalOperatorId = 0;
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
  &-driver{
    margin-top: 8px;
    width: 405px;
    height: 80px;
    border: 1px solid #E0E0E6;
    border-radius: 2px;
    background: #FAFAFC;
    padding: 20px 14px;
    display: flex;
    justify-content: space-between;
    &-info{
      height: 100%;
      font-weight: 400;
      font-size: 14px;
      color: #C2C2C2;
      line-height: 20px;
      &.choose{
        font-weight: 500;
        display: flex;
        flex-direction: column;
        justify-content: center;
        color: #3D3D3D;
      }
    }
    &-btn{
      height: 100%;
      display: flex;
      align-items: center;
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
  &-user{
    margin-top: 8px;
    font-weight: 600;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
  }
}
</style>


