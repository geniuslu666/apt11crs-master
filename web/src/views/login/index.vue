<template>
<div class="login-container">
  <!-- Left decorative panel (Stripe-style gradient) -->
  <div class="login-panel-left hidden md:flex">
    <div class="login-panel-left-content">
      <div class="login-brand">
        <img src="../../assets/images/new/login_logo.png" alt="Logo" class="login-brand-logo">
        <span class="login-brand-name">住一CRS</span>
      </div>
      <div class="login-panel-tagline">
        <h1>业务中台</h1>
        <p>统一管理 · 高效协作 · 实时数据</p>
      </div>
      <img src="../../assets/images/new/login_bg.png" class="login-panel-bg" alt="">
    </div>
  </div>

  <!-- Right form panel -->
  <div class="login-form-panel">
    <div class="login-form-box">
      <!-- Logo (mobile only) -->
      <div class="flex items-center gap-2 mb-8 md:hidden">
        <img src="../../assets/images/new/login_logo.png" alt="Logo" class="h-8 w-8 rounded-lg">
        <span class="text-base font-bold text-[#1a1f36] tracking-tight">住一CRS</span>
      </div>

      <h2 class="login-form-title">欢迎回来</h2>
      <p class="login-form-subtitle">请输入您的账号密码登录</p>

      <n-form
        ref="formRef"
        label-placement="left"
        :model="formInline"
        :rules="rules"
        class="login-form"
      >
        <n-form-item path="username" :show-label="false">
          <div class="login-field">
            <label class="login-label">用户名</label>
            <n-input
              @keyup.enter="handleSubmit"
              v-model:value="formInline.username"
              placeholder="请输入用户名"
              size="large"
              class="login-input"
            />
          </div>
        </n-form-item>

        <n-form-item path="pass" :show-label="false">
          <div class="login-field">
            <div class="flex items-center justify-between mb-1.5">
              <label class="login-label">密码</label>
              <button type="button" class="login-forgot" @click="handleResetPassword">忘记密码？</button>
            </div>
            <n-input
              @keyup.enter="handleSubmit"
              v-model:value="formInline.pass"
              type="password"
              show-password-on="click"
              placeholder="请输入密码"
              size="large"
              class="login-input"
            />
          </div>
        </n-form-item>

        <n-form-item path="code" v-show="codeBase64 !== ''" :show-label="false">
          <div class="login-field">
            <label class="login-label">验证码</label>
            <div class="flex gap-3">
              <n-input
                placeholder="请输入验证码"
                @keyup.enter="handleSubmit"
                v-model:value="formInline.code"
                size="large"
                class="login-input flex-1"
              />
              <n-loading-bar-provider :to="loadingBarTargetRef" container-style="position:absolute;">
                <img
                  ref="loadingBarTargetRef"
                  class="h-10 w-28 rounded-md cursor-pointer border border-[#e3e8ef] object-cover"
                  :src="codeBase64"
                  @click="refreshCode"
                  loading="lazy"
                  alt="验证码"
                />
                <loading-bar-trigger />
              </n-loading-bar-provider>
            </div>
          </div>
        </n-form-item>

        <div class="flex items-center justify-between mb-6">
          <n-checkbox
            v-model:checked="autoLogin"
            style="--n-color-checked:#635bff;--n-border-checked:1px solid #635bff;--n-border-focus:1px solid #635bff;"
          >
            <span class="text-sm text-[#697386]">记住登录状态</span>
          </n-checkbox>
        </div>

        <n-button
          type="primary"
          size="large"
          block
          :loading="loading"
          @click="handleLogin"
          class="login-btn"
          style="--n-color:#635bff;--n-color-hover:#4f46e5;--n-color-pressed:#3e35d9;--n-color-focus:#635bff;--n-border:none;--n-border-hover:none;--n-border-pressed:none;height:42px;font-size:14px;font-weight:600;border-radius:8px;"
        >
          登录
        </n-button>
      </n-form>
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
  background: #ffffff;
}

// Left decorative panel
.login-panel-left {
  flex: 1;
  background: linear-gradient(160deg, #0a2540 0%, #1a3a5c 60%, #0d2f50 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
  min-width: 420px;

  &::before {
    content: '';
    position: absolute;
    inset: 0;
    background: radial-gradient(ellipse at 30% 40%, rgba(99,91,255,0.3) 0%, transparent 60%),
                radial-gradient(ellipse at 80% 80%, rgba(99,91,255,0.15) 0%, transparent 50%);
  }

  &-content {
    position: relative;
    z-index: 1;
    padding: 48px;
    width: 100%;
    max-width: 480px;
    display: flex;
    flex-direction: column;
    height: 100%;
  }
}

.login-brand {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: auto;

  &-logo {
    width: 36px;
    height: 36px;
    border-radius: 8px;
  }

  &-name {
    font-size: 16px;
    font-weight: 700;
    color: #ffffff;
    letter-spacing: -0.02em;
  }
}

.login-panel-tagline {
  padding-bottom: 80px;

  h1 {
    font-size: 40px;
    font-weight: 800;
    color: #ffffff;
    letter-spacing: -0.04em;
    line-height: 1.1;
    margin-bottom: 12px;
  }

  p {
    font-size: 15px;
    color: rgba(200, 210, 224, 0.8);
    letter-spacing: 0.02em;
  }
}

.login-panel-bg {
  position: absolute;
  bottom: -20px;
  right: -40px;
  width: 360px;
  opacity: 0.12;
  pointer-events: none;
}

// Right form panel
.login-form-panel {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 48px;
  min-width: 360px;
}

.login-form-box {
  width: 100%;
  max-width: 400px;
}

.login-form-title {
  font-size: 28px;
  font-weight: 700;
  color: #1a1f36;
  letter-spacing: -0.03em;
  margin-bottom: 6px;
  line-height: 1.2;
}

.login-form-subtitle {
  font-size: 14px;
  color: #697386;
  margin-bottom: 32px;
}

.login-form {
  :deep(.n-form-item) {
    margin-bottom: 0;
  }

  :deep(.n-form-item-feedback-wrapper) {
    min-height: 20px;
  }
}

.login-field {
  width: 100%;
  margin-bottom: 16px;
}

.login-label {
  display: block;
  font-size: 13px;
  font-weight: 500;
  color: #1a1f36;
  margin-bottom: 6px;
}

.login-forgot {
  font-size: 13px;
  color: #635bff;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  transition: color 0.15s;

  &:hover {
    color: #4f46e5;
  }
}

.login-input {
  :deep(.n-input__input-el) {
    font-size: 14px;
    color: #1a1f36;
  }

  :deep(.n-input) {
    border-radius: 8px;
    border: 1px solid #e3e8ef;

    &:hover {
      border-color: #9da9bb;
    }

    &.n-input--focus {
      border-color: #635bff;
      box-shadow: 0 0 0 3px rgba(99,91,255,0.15);
    }
  }
}

@media (max-width: 880px) {
  .login-panel-left {
    display: none !important;
  }
  .login-form-panel {
    flex: 1 1 100%;
  }
}

@media (max-width: 480px) {
  .login-form-panel {
    padding: 32px 24px;
    min-width: auto;
  }
}
</style>
