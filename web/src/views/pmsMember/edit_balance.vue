<template>
  <UiSheet
    :open="showModal"
    title="调整积分"
    width="420px"
    @update:open="closeForm"
  >
    <UiSpinner :show="loading">
      <div class="space-y-4">
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">当前积分</label>
          <p class="text-sm text-foreground py-1">{{ formValue.balance }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">调整值</label>
          <UiNumberInput v-model="formValue.value" placeholder="请输入调整值" />
          <span v-if="errors.value" class="text-xs text-destructive">{{ errors.value }}</span>
          <span class="text-xs text-red-500">调整后：{{ Number(formValue.value) + Number(formValue.balance) }} 积分</span>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">操作人</label>
          <p class="text-sm text-foreground py-1">{{ formValue.operator }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">备注 <span class="text-destructive">*</span>(后端展示)</label>
          <UiTextarea v-model="formValue.des" placeholder="请输入备注" :rows="3" />
          <span v-if="errors.des" class="text-xs text-destructive">{{ errors.des }}</span>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">原因 <span class="text-destructive">*</span>(前端展示)</label>
          <UiTextarea v-model="formValue.reason" placeholder="请输入原因" :rows="3" />
          <span v-if="errors.reason" class="text-xs text-destructive">{{ errors.reason }}</span>
        </div>
      </div>
    </UiSpinner>
    <template #footer>
      <UiButton variant="outline" @click="closeForm">取消</UiButton>
      <UiButton :disabled="formBtnLoading" @click="confirmForm">
        {{ formBtnLoading ? '保存中...' : '保存' }}
      </UiButton>
    </template>
  </UiSheet>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { BalanceEdit } from '@/api/pmsMember';
import { useUserStore } from '@/store/modules/user';
import { useMessage } from 'naive-ui';
import { UiSheet, UiButton, UiNumberInput, UiTextarea, UiSpinner } from '@/components/ui';

const emit = defineEmits(['reloadInfo']);
const message = useMessage();
const userStore = useUserStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: '' as any,
  balance: '' as any,
  value: 0,
  operator: userStore.getUsername,
  des: '',
  reason: '',
});
const formBtnLoading = ref(false);
const errors = reactive({ value: '', des: '', reason: '' });

function validate(): boolean {
  errors.value = '';
  errors.des = '';
  errors.reason = '';
  let valid = true;

  if ((Number(formValue.value.value) + Number(formValue.value.balance)) < 0) {
    errors.value = '调整值与当前积分数相加不能小于0';
    valid = false;
  }
  if (!formValue.value.des) {
    errors.des = '备注必填';
    valid = false;
  }
  if (!formValue.value.reason) {
    errors.reason = '原因必填';
    valid = false;
  }
  return valid;
}

function openModal(id: any, balance: any) {
  formValue.value.id = id;
  formValue.value.balance = balance;
  formValue.value.value = 0;
  formValue.value.des = '';
  formValue.value.reason = '';
  errors.value = '';
  errors.des = '';
  errors.reason = '';
  showModal.value = true;
}

function confirmForm() {
  if (!validate()) {
    message.error('验证错误');
    return;
  }
  formBtnLoading.value = true;
  BalanceEdit(formValue.value)
    .then(() => {
      message.success('操作成功');
      closeForm();
      emit('reloadInfo');
    })
    .catch(() => {})
    .finally(() => {
      formBtnLoading.value = false;
    });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formValue.value = {
    id: '',
    balance: '',
    value: 0,
    operator: userStore.getUsername,
    des: '',
    reason: '',
  };
}

defineExpose({ openModal });
</script>
