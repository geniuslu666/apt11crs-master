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
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">调整成长值</div>
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
                <n-form-item label="当前成长值" path="exp">
                  {{formValue.exp}}
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="调整值" path="value">
                  <n-input-number placeholder="请输入调整值" v-model:value="formValue.value" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="操作人" path="operator">
                  {{formValue.operator}}
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="备注" path="des">
                  <n-input type="textarea" placeholder="请输入备注" v-model:value="formValue.des" />
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
  import { ExpEdit } from '@/api/pmsMember';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useUserStore } from '@/store/modules/user';
  import { useMessage, FormItemRule } from 'naive-ui';
import {adaModalWidth} from "@/utils/hotgo";

  const emit = defineEmits(['reloadInfo']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const userStore = useUserStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref({
    id: '',
    exp: '',
    value: 0,
    operator: userStore.getUsername,
    des: '',
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
        if ((Number(value) + Number(formValue.value.exp)) < 0) {
          return new Error('调整值与当前成长值数相加不能小于0')
        }
        return true
      },
      trigger: ['blur', 'input'],
    },
  };

  function openModal(id, exp) {
    formValue.value.id = id
    formValue.value.exp = exp

    showModal.value = true;
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        ExpEdit(formValue.value).then((_res) => {
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
      exp: '',
      value: 0,
      des: '',
    }
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
