<template>
  <UiSheet
    :open="showModal"
    title="调整成长值"
    width="420px"
    @update:open="closeForm"
  >
    <UiSpinner :show="loading">
      <div class="space-y-4">
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">当前成长值</label>
          <p class="text-sm text-foreground py-1">{{ formValue.exp }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">调整值</label>
          <UiNumberInput v-model="formValue.value" placeholder="请输入调整值" />
          <span v-if="errors.value" class="text-xs text-destructive">{{ errors.value }}</span>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">操作人</label>
          <p class="text-sm text-foreground py-1">{{ formValue.operator }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">备注</label>
          <UiTextarea v-model="formValue.des" placeholder="请输入备注" :rows="3" />
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
import { ExpEdit } from '@/api/pmsMember';
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
  exp: '' as any,
  value: 0,
  operator: userStore.getUsername,
  des: '',
});
const formBtnLoading = ref(false);
const errors = reactive({ value: '' });

function validate(): boolean {
  errors.value = '';
  if ((Number(formValue.value.value) + Number(formValue.value.exp)) < 0) {
    errors.value = '调整值与当前成长值数相加不能小于0';
    return false;
  }
  return true;
}

function openModal(id: any, exp: any) {
  formValue.value.id = id;
  formValue.value.exp = exp;
  formValue.value.value = 0;
  formValue.value.des = '';
  errors.value = '';
  showModal.value = true;
}

function confirmForm() {
  if (!validate()) {
    message.error('验证错误');
    return;
  }
  formBtnLoading.value = true;
  ExpEdit(formValue.value)
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
    exp: '',
    value: 0,
    operator: userStore.getUsername,
    des: '',
  };
}

defineExpose({ openModal });
</script>
