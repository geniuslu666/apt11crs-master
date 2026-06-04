<template>
  <div class="member-admin-page">
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        size="small"
        :segmented="{ content: true }"
      >
        <n-form :label-width="110" :model="formValue" :rules="rules" ref="formRef" >
          <n-form-item path="appversion">
            <template #label>
              <span>全局积分汇率</span>
              <span style="color: red; padding-left: 8px">*请谨慎修改</span>
            </template>
            <a-input-number
              v-model:value="formValue.exchangeRate"
              placeholder="建议输入正整数"
              :min="0"
              style="width: 200px"
              :precision="2"
            />
            <template #feedback>全局积分汇率是指积分抵用比例（1积分抵多少金额），金额单位JPY，如设置为1，则1积分可抵用1JPY</template>
          </n-form-item>
          <div style="margin-top: 15px">
            <n-space>
              <n-button type="primary" @click="formSubmit">保存更新</n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';

  const settingStore = useProjectSettingStore();
  const group = ref('yyconfig');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();

  const formValue = ref({
    memberExpRate: '',
    recommendModel: '',
    exchangeRate: '',
    memberRegisterAward: '',
    channelRegisterAward: '',
    memberBrokerageRate: ''
  });

  const rules = {
    basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },
  };

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: {exchangeRate:formValue.value.exchangeRate} }).then((_res) => {
          message.success('更新成功');
          load();
        });
      } else {
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
