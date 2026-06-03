<template>
  <UiDialog
    :open="showModal"
    title="会员注销申请"
    max-width="480px"
    @update:open="closeForm"
  >
    <div class="space-y-4 py-2">
      <div class="flex flex-col gap-1">
        <label class="text-sm font-medium text-foreground">注销原因</label>
        <UiTextarea
          v-model="formValue.cancelReason"
          placeholder="请填写注销原因"
          :rows="4"
        />
        <span v-if="errors.cancelReason" class="text-xs text-destructive">{{ errors.cancelReason }}</span>
      </div>
    </div>
    <template #footer>
      <UiButton variant="outline" @click="closeForm">取消</UiButton>
      <UiButton :disabled="formBtnLoading" @click="confirmForm">
        {{ formBtnLoading ? '提交中...' : '确定' }}
      </UiButton>
    </template>
  </UiDialog>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { Cancel } from '@/api/pmsMember';
import { useMessage } from 'naive-ui';
import { UiDialog, UiButton, UiTextarea } from '@/components/ui';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const showModal = ref(false);
const formValue = ref({
  cancelReason: '',
  id: 0,
});
const formBtnLoading = ref(false);
const errors = reactive({ cancelReason: '' });

function openModal(state: any) {
  showModal.value = true;
  formValue.value.id = state.id;
  formValue.value.cancelReason = '';
  errors.cancelReason = '';
}

function confirmForm() {
  errors.cancelReason = '';
  formBtnLoading.value = true;
  Cancel({
    id: formValue.value.id,
    cancelReason: formValue.value.cancelReason,
  })
    .then(() => {
      message.success('操作成功');
      closeForm();
      emit('reloadTable');
    })
    .catch(() => {})
    .finally(() => {
      formBtnLoading.value = false;
    });
}

function closeForm() {
  showModal.value = false;
  formValue.value.cancelReason = '';
  errors.cancelReason = '';
}

defineExpose({ openModal });
</script>
