<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title='取消预约'
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :rules="rules"
            :model="formValue"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">

              <n-gi span="1">
                <n-form-item label="取消原因" path="reason">
                  <n-input
                    type="textarea"
                    placeholder="请填写取消原因"
                    v-model:value="formValue.reason"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <div style="color: red; padding-left: 100px; ">此操作将取消预约，如果已支付，预定费将全部退款，优惠券不可退。</div>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Cancel } from '@/api/pmsRoomReservation';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import {FormItemRule, useMessage} from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});
const formValue = ref({
  reason: "",
  id: "",
  orderSn: "",
})

// 表单验证规则
const rules = {
  reason: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入取消原因',
  },
};

function openModal(state) {
  showModal.value = true;
  // 编辑

  formValue.value.id = state.id
  formValue.value.orderSn = state.orderSn

}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      // console.log(formValue.value)
      Cancel({
        id: formValue.value.id,
        orderSn: formValue.value.orderSn,
        reason: formValue.value.reason,
      }).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
        });
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


