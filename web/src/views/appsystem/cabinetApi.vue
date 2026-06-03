<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              储物柜Api配置
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="域名" path="cabinetDomain">
                <n-input v-model:value="formValue.cabinetDomain" placeholder="域名" style="width: 400px;"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="AppId" path="cabinetAppid">
                <n-input v-model:value="formValue.cabinetAppid" placeholder="AppId" style="width: 200px;"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="AppKey" path="cabinetAppkey">
                <n-input v-model:value="formValue.cabinetAppkey" placeholder="AppKey" style="width: 200px;"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="AppSecret" path="cabinetApiSecret">
                <n-input v-model:value="formValue.cabinetApiSecret" placeholder="AppSecret" style="width: 200px;"/>
              </n-form-item>
            </div>
          </n-gi>
        </n-grid>
        <div style="text-align: center">
          <n-space justify="center">
            <n-button type="primary" :loading="formBtnLoading" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {NButton, useMessage} from 'naive-ui';
import {getConfig, updateConfig} from '@/api/sys/config';

const formBtnLoading = ref(false);
const group = ref('cabinetApi');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    cabinetDomain: '',
    cabinetAppid: '',
    cabinetAppkey: '',
    cabinetApiSecret: ''
  });

  const rules = {
  };

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
        formBtnLoading.value = false;
        message.error('验证失败，请填写完整信息');
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
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
