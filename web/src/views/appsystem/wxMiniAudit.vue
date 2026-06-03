<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              小程序审核
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="是否开启" path="isOpen">
                <n-switch v-model:value="formValue.isOpen" :unchecked-value="2" :checked-value="1"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi v-if="formValue.isOpen === 1">
            <div style="margin-left: 40px">
              <n-button type="primary" :loading="formBtnLoading1" @click="handleUpdateMember">更新订单用户</n-button>
              <div style="font-size: 14px;color: red;margin-top: 5px;">再点击保存前确保点击过更新订单用户按钮操作</div>
            </div>
          </n-gi>
          <!-- <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="接送机订单用户ID" path="carOrderMemberIds">
                <n-input v-model:value="formValue.carOrderMemberIds" placeholder="接送机订单用户ID" style="width: 600px;"/>
              </n-form-item>
            </div>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-form-item label="按摩订单用户ID" path="spaOrderMemberIds">
                <n-input v-model:value="formValue.spaOrderMemberIds" placeholder="按摩订单用户ID" style="width: 600px;"/>
              </n-form-item>
            </div>
          </n-gi> -->
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
import {getConfig, updateConfig, updateOrderMember} from '@/api/sys/config';

const formBtnLoading = ref(false);
const formBtnLoading1 = ref(false);
const group = ref('wxminAudit');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    isOpen: 1,
    carOrderMemberIds: '',
    spaOrderMemberIds: ''
  });

  const rules = {
  };

  function handleUpdateMember() {
    formBtnLoading1.value = true;
    updateOrderMember().then((_res) => {
      formBtnLoading1.value = false;
      message.success('更新成功');
    }).catch((err) => {
      formBtnLoading1.value = false;
      message.error('更新失败');
    });
  }

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
