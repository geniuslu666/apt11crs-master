<template>
  <UiDialog
    :open="showModal"
    :title="formValue.id > 0 ? '编辑会员信息 #' + formValue.id : '添加会员信息'"
    max-width="600px"
    @update:open="closeForm"
  >
    <UiSpinner :show="loading">
      <div class="space-y-4 py-2">
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">会员号</label>
          <UiInput v-model="formValue.memberNo" placeholder="请输入会员号" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">名</label>
          <UiInput v-model="formValue.firstName" placeholder="请输入名" />
          <span v-if="errors.firstName" class="text-xs text-destructive">{{ errors.firstName }}</span>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">姓</label>
          <UiInput v-model="formValue.lastName" placeholder="请输入姓" />
          <span v-if="errors.lastName" class="text-xs text-destructive">{{ errors.lastName }}</span>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">全名</label>
          <UiInput v-model="formValue.fullName" placeholder="请输入全名" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">区号</label>
          <UiInput v-model="formValue.phoneArea" placeholder="请输入手机区号" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">手机号</label>
          <UiInput v-model="formValue.phone" placeholder="请输入手机号" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">邮箱</label>
          <UiInput v-model="formValue.mail" placeholder="请输入邮箱" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">生日</label>
          <input
            type="date"
            v-model="formValue.birthday"
            class="flex h-8 w-full rounded-md border border-input bg-background px-3 py-1 text-[13px] shadow-sm transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
          />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">积分</label>
          <p class="text-sm text-foreground">{{ formValue.balance }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">推荐人</label>
          <p class="text-sm text-foreground">{{ formValue.referrer }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">注册来源</label>
          <p class="text-sm text-foreground">{{ formValue.source }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">上次登录IP</label>
          <p class="text-sm text-foreground">{{ formValue.lastLoginIp }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">上次登录时间</label>
          <p class="text-sm text-foreground">{{ formValue.lastLogin }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">注册IP</label>
          <p class="text-sm text-foreground">{{ formValue.registerIp }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">注册时间</label>
          <p class="text-sm text-foreground">{{ formValue.registerTime }}</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">地址</label>
          <UiInput v-model="formValue.address" placeholder="请输入地址" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-foreground">国籍</label>
          <p class="text-sm text-foreground">{{ formValue.nationality }}</p>
        </div>
      </div>
    </UiSpinner>
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
import { Edit, View } from '@/api/pmsMember';
import { State, newState } from './model';
import { useMessage } from 'naive-ui';
import { UiDialog, UiButton, UiInput, UiSpinner } from '@/components/ui';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formBtnLoading = ref(false);
const errors = reactive({ firstName: '', lastName: '' });

function validate(): boolean {
  errors.firstName = '';
  errors.lastName = '';
  let valid = true;
  if (!formValue.value.firstName) {
    errors.firstName = '请输入名';
    valid = false;
  }
  if (!formValue.value.lastName) {
    errors.lastName = '请输入姓';
    valid = false;
  }
  return valid;
}

function openModal(state: State) {
  showModal.value = true;
  errors.firstName = '';
  errors.lastName = '';

  if (!state || state.id < 1) {
    formValue.value = newState(state);
    return;
  }

  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm() {
  if (!validate()) {
    message.error('请填写完整信息');
    return;
  }
  formBtnLoading.value = true;
  Edit(formValue.value)
    .then(() => {
      message.success('操作成功');
      closeForm();
      emit('reloadTable');
    })
    .finally(() => {
      formBtnLoading.value = false;
    });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  errors.firstName = '';
  errors.lastName = '';
}

defineExpose({ openModal });
</script>
