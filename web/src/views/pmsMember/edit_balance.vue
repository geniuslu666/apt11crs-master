<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">调整积分</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            label-placement="top"
            label-width="auto"
          >
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="当前积分" path="balance">
                  {{formValue.balance}}
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="调整值" path="value">
                  <n-input-number placeholder="请输入调整值" v-model:value="formValue.value" />
                  <!--调整后数值是-->
                  <template #feedback><span style="color: red">调整后：{{ formValue.value + formValue.balance }} 积分</span></template>
                </n-form-item>
              </n-gi>
              <n-gi style="margin-top: 24px">
                <n-form-item label="操作人" path="operator">
                  {{formValue.operator}}
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="备注(后端展示)" path="des" :show-require-mark="true">
                  <n-input type="textarea" placeholder="请输入备注" v-model:value="formValue.des" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="原因(前端展示)" path="reason" :show-require-mark="true">
                  <n-input type="textarea" placeholder="请输入原因" v-model:value="formValue.reason" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
  import { BalanceEdit } from '@/api/pmsMember';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useUserStore } from '@/store/modules/user';
  import { useMessage, FormItemRule } from 'naive-ui';
  import Verify from "@/views/spaSettlementOrder/verify.vue";
import {adaModalWidth} from "@/utils/hotgo";

  const emit = defineEmits(['reloadInfo']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const userStore = useUserStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref({
    id: '',
    balance: '',
    value: 0,
    operator: userStore.getUsername,
    des: '',
    reason: '',
  })
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(650);
  });

  const rules = {
    value: {
      required: true,
      validator(rule: FormItemRule, value: string) {
        if ((Number(value) + Number(formValue.value.balance)) < 0) {
          return new Error('调整值与当前积分数相加不能小于0')
        }
        return true
      },
      trigger: ['blur', 'input'],
    },
    des: {
      required: true,
      trigger: ['blur', 'input'],
      validator(rule: FormItemRule, value: string) {
        if (!value) {
          return new Error('备注必填')
        }
      }
    },
    reason: {
      required: true,
      trigger: ['blur', 'input'],
      validator(rule: FormItemRule, value: string) {
        if (!value) {
          return new Error('原因必填')
        }
      }
    },
  };

  function openModal(id, balance) {
    formValue.value.id = id
    formValue.value.balance = balance

    showModal.value = true;
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        BalanceEdit(formValue.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            closeForm();
            formBtnLoading.value = false;
            emit('reloadInfo');
          });
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证错误');
        formBtnLoading.value = false;
      }
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
    formValue.value = {
      id: '',
      balance: '',
      value: 0,
      des: '',
      reason: '',
    }
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
