<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title='注销申请详情'
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="会员账号" path="memberNo">
                  {{ formValue.memberNo }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="手机" path="phone">
                  {{ formValue.phoneArea }}-{{ formValue.phone }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="邮箱" path="mail">
                  {{ formValue.mail }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="注销原因" path="cancelReason">
                  {{ formValue.cancelReason }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="操作员" path="cancelReason">
                  {{ formValue.adminMemberUsername }}
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.auditStatus == 1">
                <n-form-item label="审核备注" path="auditReason">
                  <n-input
                    type="textarea"
                    placeholder="请填写通过/拒绝原因"
                    v-model:value="formValue.auditReason"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-else>
                <n-form-item label="审核备注" >
                  {{ formValue.auditReason }}
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space v-if="formValue.auditStatus == '1'">
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="handleAgree">
            通过
          </n-button>
          <n-button type="warning" :loading="formBtnLoading" @click="handleDisagree">
            拒绝
          </n-button>
        </n-space>
        <n-space v-else>
          <n-button @click="closeForm">
            关闭
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Agree, Disagree } from '@/api/pmsMemberCancel';
import { State, newState } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import {useDialog, useMessage} from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';
import {useUserStore} from "@/store/modules/user";

const userStore = useUserStore();
// 表单验证规则
const rules = {
  auditReason: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入审核备注',
  },
};
const dialog = useDialog();
const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref({
  id: 0,
  auditReason: '',
  memberNo: '',
  phone: '',
  phoneArea: '',
  mail: '',
  cancelReason: '',
  auditStatus: '',
  adminMemberUsername: '',
});
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});

function openModal(state: State) {
  // userStore.getUsername
  showModal.value = true;
  if (state.auditStatus == '1') {
    state.adminMemberUsername = userStore.getUsername;
  }
  formValue.value = state;
  console.log("formvalue", formValue.value);
}

// 同意注销
function handleAgree(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      dialog.info({
        title: '提示',
        content: '你确定同意注销该会员吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          formBtnLoading.value = true;
          Agree({
            id: formValue.value.id,
            auditReason: formValue.value.auditReason
          }).then((_res) => {
            message.success('操作成功');
            formBtnLoading.value = false;
            setTimeout(() => {
              closeForm();
              emit('reloadTable');
            });
          }).catch((err)=>{
            formBtnLoading.value = false;
          });
        },
        onNegativeClick: () => {
          formBtnLoading.value = false;
        },
      });

    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }

  });
}

// 拒绝注销
function handleDisagree(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      dialog.info({
        title: '提示',
        content: '你确定拒绝注销该会员吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          Disagree({
            id: formValue.value.id,
            auditReason: formValue.value.auditReason
          }).then((_res) => {
            message.success('操作成功');
            formBtnLoading.value = false;
            setTimeout(() => {
              closeForm();
              emit('reloadTable');
            });
          });
        },
        onNegativeClick: () => {
          formBtnLoading.value = false;
        },
      });

    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
    // formBtnLoading.value = false;
  });
}


/*function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      console.log(formValue.value)
      Cancel({
        id: formValue.value.id,
        auditReason: formValue.value.auditReason
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
}*/

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formValue.value.auditReason = ""
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


