<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">司机设置</text>
        </template>
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="150"
          class="py-4"
          style="padding-top: 0"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="预估服务时长" path="carDriverPreServiceTime" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-number placeholder="请输入分钟" :min="0" :precision="0" :show-button="false" v-model:value="formValue.carDriverPreServiceTime" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
                <template #feedback>
                </template>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="退单服务费" path="carDriverCancelFee">
                <n-input-group>
                  <n-input-number placeholder="请输入JPY" :min="0" :precision="0" :show-button="false" v-model:value="formValue.carDriverCancelFee" style="width: 100px" />
                  <n-input-group-label>JPY</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="最低可提现金额" path="rate">
                <n-input-group>
                  <n-input-number v-model:value="formValue.minWithdrawalAmount" :show-button="false" style="width: 100px" />
                  <n-input-group-label>JPY</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="手续费" path="rate">
                <n-input-group>
                  <n-input-number v-model:value="formValue.serviceCharge" :show-button="false"  style="width: 100px" />
                  <n-input-group-label>%</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预计提现时间周期" path="rate">
                <n-input-group>
                  <n-input-number v-model:value="formValue.afterDay" :show-button="false"  style="width: 100px" />
                  <n-input-group-label>天</n-input-group-label>
                </n-input-group>
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
import {ref, onMounted, reactive} from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {Text} from "@/api/translate";

  const rules = ref({});
  const group = ref('cardriversetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    carDriverPreServiceTime: null,
    carDriverCancelFee: null,
    minWithdrawalAmount: null,
    serviceCharge: null,
    afterDay: null
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
