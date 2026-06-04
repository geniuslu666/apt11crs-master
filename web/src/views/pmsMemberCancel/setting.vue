<template>
  <div class="member-admin-page">
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="基础设置"
      >
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="150"
          class="py-4"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="是否允许注销" path="isOpen">
                <n-switch v-model:value="formValue.isEnableCancel" :unchecked-value="2" :checked-value="1"/>
                <template #feedback>设置为关闭，则无法注销会员账号</template>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="是否注销审核" path="isOpen">
                <n-switch v-model:value="formValue.isCancelAudit" :unchecked-value="2" :checked-value="1"/>
                <template #feedback>设置为关闭，会员将直接注销成功。</template>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="注销后不允许注册" path="isOpen">
                <n-switch v-model:value="formValue.isAllowedRegister" :unchecked-value="2" :checked-value="1"/>
                <template #feedback>设置为开启，会员注销后不允许继续注册，设置为关闭，会员注销后允许继续注册。</template>
              </n-form-item>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="formSubmit">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";

  const rules = ref({});
  const group = ref('membercancelsetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    isEnableCancel: 1,
    isCancelAudit: 1,
    isAllowedRegister: 2
  });

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          formBtnLoading.value = false;
          message.success('更新成功');
          load();
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证失败，请填写完整信息');
        formBtnLoading.value = false;
      }
    });
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          if(res.list){
            formValue.value = res.list;
          }
          show.value = false;
        }).catch((err)=>{
          show.value = false;
        })
    });
  }
</script>
