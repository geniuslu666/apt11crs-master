<template>
<div class="login-container">
  <div class="login-container-left">
    <img src="../../assets/images/new/login_bg.png" class="login-container-left-bg" alt="Bg">
    <div class="login-container-left-logo">
      <img src="../../assets/images/new/login_logo.png" alt="Logo">
    </div>
  </div>
  <div class="login-container-right">
    <div class="login-container-right-login-box">
      <div class="login-container-right-login-box-title">住一CRS业务中台</div>
      <div class="login-container-right-login-box-subtitle">请输入您的账号密码登录</div>
      <div class="login-container-right-login-box-form">
        <n-form
          ref="formRef"
          label-placement="left"
          size="large"
          :model="formInline"
          :rules="rules"
        >
          <n-form-item path="username">
            <n-input
              @keyup.enter="handleSubmit"
              v-model:value="formInline.username"
              placeholder="请输入用户名"
              class="login-input"
            >
              <template #prefix>
                <img src="../../assets/images/new/login_user_icon.png" width="16">
              </template>
            </n-input>
          </n-form-item>
          <n-form-item path="pass">
            <n-input
              @keyup.enter="handleSubmit"
              v-model:value="formInline.pass"
              type="password"
              show-password-on="click"
              placeholder="请输入密码"
              class="login-input"
            >
              <template #prefix>
                <img src="../../assets/images/new/login_pass_icon.png" width="16">
              </template>
            </n-input>
          </n-form-item>

          <n-form-item path="code" v-show="codeBase64 !== ''">
            <n-input-group>
              <n-grid x-gap="9" :cols="4">
                <n-gi :span="3">
                  <n-input
                    :style="{ width: '100%', marginRight: '16px' }"
                    placeholder="验证码"
                    @keyup.enter="handleSubmit"
                    v-model:value="formInline.code"
                    class="login-input"
                  >
                    <template #prefix>
                      <img src="../../assets/images/new/login_code_icon.png" width="16">
                    </template>
                  </n-input>
                </n-gi>
                <n-gi :span="1">
                  <n-loading-bar-provider
                    :to="loadingBarTargetRef"
                    container-style="position: absolute;"
                  >
                    <img
                      ref="loadingBarTargetRef"
                      style="width: 100%; height: 100%"
                      :src="codeBase64"
                      @click="refreshCode"
                      loading="lazy"
                      alt="点击获取"
                      class="image"
                    />
                    <loading-bar-trigger />
                  </n-loading-bar-provider>
                </n-gi>
              </n-grid>
            </n-input-group>
          </n-form-item>

          <n-space :vertical="true" :size="32">
            <div style="display: flex;align-items: center;justify-content: space-between">
              <n-checkbox v-model:checked="autoLogin" style="--n-color-checked: #2398F5;--n-border-checked: 1px solid #2398F5;--n-border-focus: 1px solid #2398F5;">
                <span style="color: #3D3D3D">自动登录</span>
              </n-checkbox>
              <n-button :text="true" @click="handleResetPassword">
                <span style="color: #2398F5">忘记密码？</span></n-button>
            </div>
            <n-button
              type="primary"
              size="large"
              :block="true"
              :loading="loading"
              @click="handleLogin"
              color="#4473E8"
            >
              登录
            </n-button>

            <!-- <FormOther moduleKey="register" tag="注册账号" @updateActiveModule="updateActiveModule" /> -->
          </n-space>
        </n-form>
      </div>
    </div>
  </div>
</div>
</template>

<script lang="ts" setup>

import {onMounted, ref} from "vue";
import {aesEcb} from "@/utils/encrypt";
import {useLoadingBar, useMessage} from "naive-ui";
import {useUserStore} from "@/store/modules/user";
import {ResultEnum} from "@/enums/httpEnum";
import {GetCaptcha} from "@/api/base";
import {useRoute, useRouter} from "vue-router";
import {PageEnum} from "@/enums/pageEnum";

const formRef = ref();
const message = useMessage();
const userStore = useUserStore();
const loading = ref(false);
const codeBase64 = ref('');
const loadingBar = useLoadingBar();
const loadingBarTargetRef = ref<undefined | HTMLElement>(undefined);
const router = useRouter();
const route = useRoute();
const LOGIN_NAME = PageEnum.BASE_LOGIN_NAME;
const autoLogin = ref(true);

interface FormState {
  username: string;
  pass: string;
  cid: string;
  code: string;
  password: string;
}
const formInline = ref<FormState>({
  username: '',
  pass: '',
  cid: '',
  code: '',
  password: '',
});

const rules = {
  username: { required: true, message: '请输入用户名', trigger: 'blur' },
  pass: { required: true, message: '请输入密码', trigger: 'blur' },
};

const handleSubmit = (e) => {
  e.preventDefault();
  formRef.value.validate(async (errors) => {
    if (!errors) {
      if (userStore.loginConfig?.loginCaptchaSwitch === 1 && formInline.value.code === '') {
        message.error('请输入验证码');
        return;
      }

      const params = {
        username: formInline.value.username,
        password: aesEcb.encrypt(formInline.value.pass),
        cid: formInline.value.cid,
        code: formInline.value.code,
      };
      await handleLoginResp(userStore.login(params));
    } else {
      message.error('请填写完整信息，并且进行验证码校验');
    }
  });
};

function handleResetPassword() {
  message.info('如果你忘记了密码，请联系管理员找回');
}

async function refreshCode() {
  if (userStore.loginConfig?.loginCaptchaSwitch !== 1) {
    return;
  }
  loadingBar.start();
  const data = await GetCaptcha();
  codeBase64.value = data.base64;
  formInline.value.cid = data.cid;
  formInline.value.code = '';
  loadingBar.finish();
}

async function handleLoginResp(request: Promise<any>) {
  message.loading('登录中...');
  loading.value = true;
  try {
    const { code, message: msg } = await request;
    message.destroyAll();
    if (code == ResultEnum.SUCCESS) {
      const toPath = decodeURIComponent((route.query?.redirect || '/') as string);
      message.success('登录成功，即将进入系统');
      if (route.name === LOGIN_NAME) {
        await router.replace('/');
      } else await router.replace(toPath);
    } else {
      message.destroyAll();
      message.info(msg || '登录失败');
      await refreshCode();
    }
  } finally {
    loading.value = false;
  }
}

function handleLogin(e) {
  handleSubmit(e);
}

onMounted(() => {
  setTimeout(function () {
    refreshCode();
  });
});
</script>

<style lang="less" scoped>
.login-container {
  display: flex;
  width: 100%;
  height: 100vh;
  &-left {
    position: relative;
    flex: 1;
    background: #F3F4FB;
    display: flex;
    justify-content: center;
    align-items: center;
    min-width: 440px;
    &-bg{
      width: 440px;
    }
    &-logo{
      position: absolute;
      top: 25px;
      left: 25px;
      img{
        width: 110px;
      }
    }
  }
  &-right {
    flex: 1;
    background: white;
    display: flex;
    justify-content: center;
    align-items: center;
    min-width: 440px;
    &-login-box{
      width: 400px;
      &-title{
        font-weight: 600;
        font-size: 32px;
        color: #3D3D3D;
        line-height: 45px;
      }
      &-subtitle{
        font-weight: 400;
        font-size: 14px;
        color: #979797;
        line-height: 20px;
        margin-top: 5px;
      }
      &-form{
        margin-top: 32px;
      }
    }
  }
}
.login-input{
  color: #3D3D3D;
}

/* 当屏幕小于880px，隐藏左侧 */
@media (max-width: 880px) {
  .login-container-left {
    display: none;
  }
  .login-container-right {
    flex: 1 1 100%;
    justify-content: center;
    align-items: center;
  }
}

@media (max-width: 480px) {
  .login-container-right {
    min-width: auto;
  }
  .login-container-right-login-box {
    width: 90%;
  }
}
</style>
