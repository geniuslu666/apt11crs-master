<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              ToretaApi配置
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="域名" path="cabinetDomain">
                <n-input v-model:value="formValue.toretaDomain" placeholder="域名" style="width: 400px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="API TOKEN" path="apiToken">
                <n-input v-model:value="formValue.apiToken" placeholder="API TOKEN" style="width: 400px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="ClientId" path="clientId">
                <n-input v-model:value="formValue.clientId" placeholder="ClientId" style="width: 400px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="ClientSecret" path="clientSecret">
                <n-input v-model:value="formValue.clientSecret" placeholder="ClientSecret" style="width: 750px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="AccessToken" path="accessToken">
                <n-input v-model:value="formValue.accessToken" disabled placeholder="AccessToken" style="width: 300px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="Token过期时间" path="expiresTime">
                <n-input v-model:value="formValue.expiresTime" disabled placeholder="Token过期时间" style="width: 300px;"/>
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
const group = ref('toretaApi');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    toretaDomain: '',
    apiToken: '',
    clientId: '',
    clientSecret: '',
    accessToken: '',
    refreshToken: '',
    expiresAt: 0,
    expiresTime: ''
  });

  const rules = {
  };

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        delete formValue.value.refreshToken;
        delete formValue.value.accessToken;
        delete formValue.value.expiresTime;
        delete formValue.value.expiresAt;
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
          formValue.value.expiresTime = new Date(res.list.expiresAt * 1000).toLocaleString();
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
