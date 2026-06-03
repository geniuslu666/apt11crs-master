<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="130" :model="formValue" :rules="rules" ref="formRef" label-placement="left">
        <n-divider title-placement="left">基础配置</n-divider>

        <n-form-item label="启用极验" path="geeTestEnabled">
          <n-radio-group v-model:value="formValue.geeTestEnabled" name="geeTestEnabled">
            <n-space>
              <n-radio :value="true">开启</n-radio>
              <n-radio :value="false">关闭</n-radio>
            </n-space>
          </n-radio-group>
          <template #feedback>关闭后发送短信验证码不会校验滑块</template>
        </n-form-item>

        <n-form-item label="Captcha ID" path="geeTestCaptchaID">
          <n-input v-model:value="formValue.geeTestCaptchaID" placeholder="请输入 GeeTest captcha_id" />
          <template #feedback>GeeTest GT4 控制台提供的 captcha_id</template>
        </n-form-item>

        <n-form-item label="Private Key" path="geeTestCaptchaKey">
          <n-input
            v-model:value="formValue.geeTestCaptchaKey"
            type="password"
            placeholder="请输入 GeeTest private_key"
            show-password-on="click"
          >
            <template #password-visible-icon>
              <n-icon :size="16" :component="GlassesOutline" />
            </template>
            <template #password-invisible-icon>
              <n-icon :size="16" :component="Glasses" />
            </template>
          </n-input>
          <template #feedback>用于服务端二次校验 sign_token 生成</template>
        </n-form-item>

        <n-form-item label="校验地址" path="geeTestValidateURL">
          <n-input
            v-model:value="formValue.geeTestValidateURL"
            placeholder="https://gcaptcha4.geetest.com/validate"
          />
          <template #feedback>一般保持默认即可</template>
        </n-form-item>

        <n-alert type="info" :show-icon="false" style="margin-bottom: 20px">
          Flutter 端在滑块成功后，需要把 <n-text code>captcha_output</n-text>、
          <n-text code>lot_number</n-text>、<n-text code>pass_token</n-text>、
          <n-text code>gen_time</n-text> 连同手机号一起提交到
          <n-text code>/api/member/SendPhoneCode</n-text>。
        </n-alert>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import { GlassesOutline, Glasses } from '@vicons/ionicons5';

  const group = ref('geetest');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    geeTestEnabled: false,
    geeTestCaptchaID: '',
    geeTestCaptchaKey: '',
    geeTestValidateURL: 'https://gcaptcha4.geetest.com/validate',
  });

  const rules = {
    geeTestCaptchaID: {
      validator: (_rule, value) => {
        if (formValue.value.geeTestEnabled && !value) {
          return new Error('请输入 GeeTest captcha_id');
        }
        return true;
      },
      trigger: ['blur', 'input'],
    },
    geeTestCaptchaKey: {
      validator: (_rule, value) => {
        if (formValue.value.geeTestEnabled && !value) {
          return new Error('请输入 GeeTest private_key');
        }
        return true;
      },
      trigger: ['blur', 'input'],
    },
  };

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then(() => {
          message.success('更新成功');
          load();
        });
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  function load() {
    show.value = true;
    getConfig({ group: group.value })
      .then((res) => {
        formValue.value = {
          ...formValue.value,
          ...res.list,
        };
      })
      .finally(() => {
        show.value = false;
      });
  }

  onMounted(() => {
    load();
  });
</script>
