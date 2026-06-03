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
                v-model:value="formValue.applyRemark"
              />
            </div>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { DisagreeStaff } from '@/api/pmsWithdraw';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  applyRemark: '',
});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});

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
  if(!formValue.value.applyRemark){
    formBtnLoading.value = false;
    message.error('请输入拒绝原因');
    return false;
  }
  DisagreeStaff({
    id: formValue.value.id,
    applyRemark: formValue.value.applyRemark
  }).then((_res) => {
    message.success('操作成功');
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
  formValue.value.applyRemark = '';
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


