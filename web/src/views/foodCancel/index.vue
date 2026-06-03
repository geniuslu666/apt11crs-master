<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">取消政策</text>
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
              <n-form-item label="是否允许取消" path="cancelPolicyOpen">
                <n-radio-group v-model:value="formValue.cancelPolicyOpen" name="cancelPolicyOpen">
                  <n-space>
                    <n-radio :value="1">
                      允许
                    </n-radio>
                    <n-radio :value="2">
                      不允许
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
<!--              <n-form-item label="订单确认前取消费用" path="beforeConfirmCancelRate" v-if="formValue.cancelPolicyOpen == 1" style="margin-bottom: 24px">-->
<!--                <n-input-group>-->
<!--                  <n-input-group-label>订单金额的</n-input-group-label>-->
<!--                  <n-input-number placeholder="请输入" :min="0" :max="100" v-model:value="formValue.beforeConfirmCancelRate" style="width: 100px" />-->
<!--                  <n-input-group-label>%</n-input-group-label>-->
<!--                </n-input-group>-->
<!--                <template #feedback>-->
<!--                  填0则免费取消-->
<!--                </template>-->
<!--              </n-form-item>-->
              <n-form-item label="订单确认后取消费用" path="afterConfirmCancel" v-if="formValue.cancelPolicyOpen == 1" >
                <n-input-group>
                  <n-input-group-label>距离到店时间</n-input-group-label>
                  <n-input-number placeholder="请输入" :min="0" v-model:value="formValue.afterConfirmCancelDay" style="width: 100px" />
                  <n-input-group-label>天前，收取订单金额的</n-input-group-label>
                  <n-input-number placeholder="请输入" :min="0" :max="100" v-model:value="formValue.afterConfirmCancelRate1" style="width: 100px" />
                  <n-input-group-label>%</n-input-group-label>
                </n-input-group>
              </n-form-item>
              <n-form-item label=" " path="afterConfirmCancel" v-if="formValue.cancelPolicyOpen == 1" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-group-label>之后收取订单金额的</n-input-group-label>
                  <n-input-number placeholder="请输入" :min="0" :max="100" v-model:value="formValue.afterConfirmCancelRate2" style="width: 100px" />
                  <n-input-group-label>%</n-input-group-label>
                </n-input-group>
                <template #feedback>
                  两个比例都填0则免费取消
                </template>
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
  import {Text} from "@/api/translate";

  const rules = ref({});
  const group = ref('foodcancelpolicy');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    cancelPolicyOpen: 1,
    beforeConfirmCancelRate: 0,
    afterConfirmCancelDay: 0,
    afterConfirmCancelRate1: 0,
    afterConfirmCancelRate2: 0,
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
