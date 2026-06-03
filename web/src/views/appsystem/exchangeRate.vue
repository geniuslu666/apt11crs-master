<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              汇率配置
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="日元汇率" path="jaRate">
                <n-input-group>
                  <n-input-group-label>1日元≈</n-input-group-label>
                  <n-input v-model:value="formValue.jaRate" :style="{ width: '80px' }" />
                  <n-input-group-label>人民币</n-input-group-label>
                </n-input-group>
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
import {useDialog, useMessage} from 'naive-ui';
import {getConfig, updateConfig} from '@/api/sys/config';

const formBtnLoading = ref(false);
const group = ref('rmbExchangeRate');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();

  const formValue = ref({
    jaRate: 0,
  });

  // const payMode = ref([])

  const rules = {
  /*  basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },*/
  };

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        // formValue.value.paymode = payMode.value.join(",")
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
          // payMode.value = res.list.paymode.split(",").map(item => String(item));
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
