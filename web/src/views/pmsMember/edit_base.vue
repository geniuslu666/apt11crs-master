<template>
  <UiSheet
    :open="showModal"
    :title="baseTitle"
    width="420px"
    @update:open="closeForm"
  >
    <UiSpinner :show="loading">
      <div class="member-admin-page member-edit-form space-y-3">
        <!-- Group -->
        <template v-if="formValue.type === 'group'">
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-foreground">会员分组</label>
            <UiSelect
              v-model="formValue.groupId"
              :options="groupList"
              label-field="memberGroup"
              value-field="id"
              placeholder="请选择会员分组"
              clearable
              filterable
            />
          </div>
        </template>
        <!-- Level -->
        <template v-if="formValue.type === 'level'">
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-foreground">会员等级</label>
            <UiSelect
              v-model="formValue.level"
              :options="levelList"
              label-field="levelName"
              value-field="id"
              placeholder="请选择会员等级"
              clearable
              filterable
            />
          </div>
        </template>
        <!-- Last Name -->
        <template v-if="formValue.type === 'lastName'">
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-foreground">会员名</label>
            <UiInput v-model="formValue.lastName" placeholder="请输入会员名" />
          </div>
        </template>
        <!-- First Name -->
        <template v-if="formValue.type === 'firstName'">
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-foreground">会员姓</label>
            <UiInput v-model="formValue.firstName" placeholder="请输入会员姓" />
          </div>
        </template>
        <!-- Phone -->
        <template v-if="formValue.type === 'phone'">
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-foreground">区号</label>
            <UiInput v-model="formValue.phoneArea" placeholder="请输入区号" />
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-foreground">手机号</label>
            <UiInput v-model="formValue.phone" placeholder="请输入手机号" />
          </div>
        </template>
        <!-- Mail -->
        <template v-if="formValue.type === 'mail'">
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-foreground">邮箱</label>
            <UiInput v-model="formValue.mail" placeholder="请输入邮箱" />
            <span v-if="mailError" class="text-xs text-destructive">{{ mailError }}</span>
          </div>
        </template>
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
import { ref } from 'vue';
import { BaseEdit } from '@/api/pmsMember';
import { useMessage } from 'naive-ui';
import { All } from "@/api/pmsMemberLevel/index";
import { GroupAll } from "@/api/pmsMemberGroup/index";
import { UiSheet, UiButton, UiInput, UiSelect, UiSpinner } from '@/components/ui';

const baseTitle = ref('');
const levelList = ref<any[]>([]);
const groupList = ref<any[]>([]);
const emit = defineEmits(['reloadInfo']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const mailError = ref('');
const formValue = ref({
  id: '' as any,
  type: '',
  level: '' as any,
  groupId: '' as any,
  firstName: '',
  lastName: '',
  phoneArea: '',
  phone: '',
  mail: '',
});
const formBtnLoading = ref(false);

function validateMail(): boolean {
  mailError.value = '';
  if (formValue.value.type === 'mail') {
    if (!formValue.value.mail) {
      mailError.value = '邮箱必填';
      return false;
    }
    if (!/^\w+(\.)?(\w+)?@[0-9a-z]+(\.[a-z]+){1,3}$/.test(formValue.value.mail)) {
      mailError.value = '邮箱格式不正确';
      return false;
    }
  }
  return true;
}

function openModal(id: any, baseType: string, value: any) {
  formValue.value.id = id;
  formValue.value.type = baseType;
  formValue.value.level = value.level;
  formValue.value.groupId = value.groupId > 0 ? value.groupId : '';
  formValue.value.firstName = value.firstName;
  formValue.value.lastName = value.lastName;
  formValue.value.phoneArea = value.phoneArea;
  formValue.value.phone = value.phone;
  formValue.value.mail = value.mail;
  mailError.value = '';

  if (baseType === 'level') {
    loading.value = true;
    All({}).then((res) => {
      levelList.value = res.list;
      baseTitle.value = '编辑会员等级';
      loading.value = false;
    });
  } else if (baseType === 'group') {
    loading.value = true;
    GroupAll({}).then((res) => {
      groupList.value = res.list;
      baseTitle.value = '编辑会员分组';
      loading.value = false;
    });
  } else if (baseType === 'firstName') {
    baseTitle.value = '编辑会员姓';
  } else if (baseType === 'lastName') {
    baseTitle.value = '编辑会员名';
  } else if (baseType === 'phone') {
    baseTitle.value = '编辑手机号';
  } else if (baseType === 'mail') {
    baseTitle.value = '编辑邮箱';
  }
  showModal.value = true;
}

function confirmForm() {
  if (!validateMail()) {
    message.error('验证错误');
    return;
  }
  formBtnLoading.value = true;
  BaseEdit(formValue.value)
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
  mailError.value = '';
  formValue.value = {
    id: '',
    type: '',
    level: '',
    groupId: '',
    firstName: '',
    lastName: '',
    phoneArea: '',
    phone: '',
    mail: '',
  };
}

defineExpose({ openModal });
</script>
