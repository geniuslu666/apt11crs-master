<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              日本客服中心
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="电话" path="japan.phone">
                <n-input v-model:value="formValue.japan.phone" placeholder="请输入日本客服电话" style="width: 400px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="邮箱" path="japan.email">
                <n-input v-model:value="formValue.japan.email" placeholder="请输入日本客服邮箱" style="width: 400px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="工作时间" path="japan.workTime">
                <n-input v-model:value="formValue.japan.workTime" placeholder="例如：09:00-18:00" style="width: 400px;"/>
                <template #feedback>
                  <span style="color: #999;">请填写日本时间（JST，UTC+9）</span>
                </template>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <n-divider title-placement="left">
              中国客服中心
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="电话" path="china.phone">
                <n-input v-model:value="formValue.china.phone" placeholder="请输入中国客服电话" style="width: 400px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="邮箱" path="china.email">
                <n-input v-model:value="formValue.china.email" placeholder="请输入中国客服邮箱" style="width: 400px;"/>
              </n-form-item>
            </div>
            <div style="margin-left: 40px">
              <n-form-item label="工作时间" path="china.workTime">
                <n-input v-model:value="formValue.china.workTime" placeholder="例如：09:00-18:00" style="width: 400px;"/>
                <template #feedback>
                  <span style="color: #999;">请填写中国时间（CST，UTC+8）</span>
                </template>
              </n-form-item>
            </div>
          </n-gi>
        </n-grid>
        <div style="text-align: center; margin-top: 20px;">
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
import {useMessage} from 'naive-ui';
import {getConfig, updateConfig} from '@/api/sys/config';

const formBtnLoading = ref(false);
const group = ref('contactCenter');
const show = ref(false);
const formRef: any = ref(null);
const message = useMessage();

interface ContactInfo {
  phone: string;
  email: string;
  workTime: string;
}

const formValue = ref({
  japan: {
    phone: '',
    email: '',
    workTime: ''
  } as ContactInfo,
  china: {
    phone: '',
    email: '',
    workTime: ''
  } as ContactInfo
});

const rules = {};

function formSubmit() {
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      // 将数据转换为 JSON 字符串存储
      const submitData = {
        contactData: JSON.stringify(formValue.value)
      };
      updateConfig({ group: group.value, list: submitData }).then((_res) => {
        formBtnLoading.value = false;
        message.success('更新成功');
        load();
      }).catch((_err) => {
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
  getConfig({ group: group.value })
    .then((res) => {
      if (res.list && res.list.contactData) {
        // 从 JSON 字符串解析数据
        try {
          const data = JSON.parse(res.list.contactData);
          formValue.value = {
            japan: data.japan || { phone: '', email: '', workTime: '' },
            china: data.china || { phone: '', email: '', workTime: '' }
          };
        } catch (e) {
          console.error('解析客服中心配置失败', e);
        }
      }
    })
    .finally(() => {
      show.value = false;
    });
}
</script>
