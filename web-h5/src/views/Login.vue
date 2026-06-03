<template>
  <div class="login-page">
    <!-- 语言切换 -->
    <!-- <div class="lang-switcher">
      <span
        v-for="lang in langs"
        :key="lang.value"
        :class="['lang-item', { active: currentLang === lang.value }]"
        @click="switchLang(lang.value)"
      >{{ lang.label }}</span>
    </div> -->

    <!-- Logo 区域 -->
    <div class="logo-area">
      <div class="logo-wrapper">
        <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0ck0nh1ohx7bgre1.png" />
      </div>
      <h1 class="platform-title">{{ t('login.title') }}</h1>
    </div>

    <!-- 表单区域 -->
    <div class="form-area">
      <div class="input-item">
        <img class="input-icon" src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0etkvpl7coylbaxu.png" />
        <input
          v-model="form.username"
          type="text"
          :placeholder="t('login.usernamePlaceholder')"
          class="input-field"
          autocomplete="username"
        />
      </div>

      <div class="input-item">
        <img class="input-icon" src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0evczubml89vs36w.png" />
        <input
          v-model="form.password"
          type="password"
          :placeholder="t('login.passwordPlaceholder')"
          class="input-field"
          autocomplete="current-password"
        />
      </div>

      <button class="login-btn" :disabled="loading" @click="onSubmit">
        <span v-if="loading" class="loading-spinner"></span>
        <span v-else>{{ t('login.loginBtn') }}</span>
      </button>

      <p class="help-text" @click="showDialog = true">{{ t('login.helpText') }}</p>
    </div>

    <!-- 联系客服弹窗 -->
    <div v-if="showDialog" class="dialog-mask">
      <div class="dialog-box">
        <div class="dialog-icon">
          <img src="https://oss.yeebok.net/static/attachment/2026-03-13/dh1i0u2vqtcotk0rjy.png" alt="tel" />
        </div>
        <div class="dialog-title">联系管理员</div>
        <div class="dialog-desc">
          <div>登录遇到困难？请拨打电话</div>
          <div>{{ configInfo?.contactMobile || '' }}</div>
        </div>
        <button class="dialog-btn" @click="makeCall">立即拨打</button>

        <div class="dialog-close" @click="showDialog = false">
          <img src="https://oss.yeebok.net/static/attachment/2026-03-13/dh1i18m4h734akjeyo.png" alt="close" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {closeToast, showLoadingToast, showToast} from 'vant'
import { useI18n } from 'vue-i18n'
import { useTravelStaffStore } from '../store/travelStaff'
import {type ConfigResponse, getConfig, staffLogin} from '../api/travelStaff'

const { t } = useI18n()
const router = useRouter()
const route = useRoute()
const staffStore = useTravelStaffStore()
const showDialog = ref(false)

// 拨打电话
const makeCall = () => {
  if(configInfo.value?.contactMobile){
    window.location.href = `tel:${configInfo.value?.contactMobile}`
  }else{
    showToast('未设置客服电话')
  }
}

// const langs = [
//   { label: '中文', value: 'zh-CN' },
//   { label: 'EN', value: 'en' },
//   { label: '日本語', value: 'ja' },
// ]
// const currentLang = ref(locale.value)

// const switchLang = (lang: string) => {
//   locale.value = lang
//   currentLang.value = lang
//   localStorage.setItem('lang', lang)
// }

const configInfo = ref<ConfigResponse | null>(null)
const form = ref({ username: '', password: '' })
const loading = ref(false)

const onSubmit = async () => {
  if (!form.value.username) {
    showToast(t('login.usernameRequired'))
    return
  }
  if (!form.value.password) {
    showToast(t('login.passwordRequired'))
    return
  }
  loading.value = true
  try {
    const res = await staffLogin(form.value)
    console.log(res)
    staffStore.setToken(res.token)
    staffStore.setStaffInfo({ username: res.username, name: res.name, mobile: res.mobile })
    const redirect = (route.query.redirect as string) || '/'
    router.replace(redirect)
  } catch (error: any) {
    showToast(error.message || t('login.loginFailed'))
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  showLoadingToast({ message: '加载中...', forbidClick: true, duration: 0 })
  try {
    configInfo.value = await getConfig()
  } catch (error: any) {
    showToast(error.message || '加载失败')
  } finally {
    closeToast()
  }
})
</script>

<style scoped>
/* Stripe-style Login page */
.login-page {
  min-height: 100vh;
  background: linear-gradient(160deg, #0a2540 0%, #1a3a5c 60%, #0d2f50 100%);
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 24px;
  position: relative;
  overflow: hidden;
}

/* Background decoration */
.login-page::before {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(ellipse at 20% 30%, rgba(99,91,255,0.25) 0%, transparent 55%),
              radial-gradient(ellipse at 85% 85%, rgba(99,91,255,0.12) 0%, transparent 50%);
  pointer-events: none;
}

.lang-switcher {
  position: absolute;
  top: 20px;
  right: 20px;
  display: flex;
  gap: 8px;
  z-index: 1;
}

.lang-item {
  font-size: 12px;
  color: rgba(200,210,224,0.7);
  cursor: pointer;
  padding: 4px 10px;
  border-radius: 100px;
  border: 1px solid rgba(255,255,255,0.12);
  transition: all 0.15s;
}

.lang-item.active {
  color: #ffffff;
  background: rgba(99,91,255,0.3);
  border-color: rgba(99,91,255,0.5);
}

/* Logo 区域 */
.logo-area {
  margin-top: 80px;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 48px;
  position: relative;
  z-index: 1;
}

.logo-wrapper {
  margin-bottom: 16px;
  width: 80px;
  height: 80px;
  background: rgba(255,255,255,0.08);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255,255,255,0.12);
}

.logo-wrapper img {
  display: block;
  width: 100px;
  height: auto;
  max-height: 44px;
  object-fit: contain;
}

.platform-title {
  font-size: 26px;
  font-weight: 700;
  color: #ffffff;
  margin: 0;
  letter-spacing: -0.03em;
  text-align: center;
}

/* 表单区域 */
.form-area {
  width: 100%;
  max-width: 380px;
  background: #ffffff;
  border-radius: 20px;
  padding: 28px 24px 32px;
  box-shadow: 0 20px 60px rgba(0,0,0,0.3), 0 0 0 1px rgba(255,255,255,0.05);
  position: relative;
  z-index: 1;
}

.form-area-title {
  font-size: 18px;
  font-weight: 700;
  color: #1a1f36;
  margin: 0 0 4px;
  letter-spacing: -0.02em;
}

.form-area-subtitle {
  font-size: 13px;
  color: #697386;
  margin: 0 0 24px;
}

.input-item {
  background: #f8fafc;
  border: 1px solid #e3e8ef;
  border-radius: 10px;
  padding: 0 14px;
  height: 48px;
  margin-bottom: 12px;
  display: flex;
  align-items: center;
  transition: border-color 0.15s, box-shadow 0.15s;
}

.input-item:focus-within {
  border-color: #635bff;
  box-shadow: 0 0 0 3px rgba(99,91,255,0.12);
  background: #ffffff;
}

.input-icon {
  width: 18px;
  height: 18px;
  margin-right: 12px;
  opacity: 0.5;
  flex-shrink: 0;
}

.input-field {
  flex: 1;
  border: none;
  outline: none;
  font-size: 15px;
  color: #1a1f36;
  background: transparent;
  font-family: inherit;
}

.input-field::placeholder {
  color: #9da9bb;
}

.login-btn {
  width: 100%;
  height: 48px;
  background: #635bff;
  border: none;
  border-radius: 10px;
  color: white;
  font-size: 15px;
  font-weight: 600;
  margin-top: 20px;
  cursor: pointer;
  transition: background 0.15s, transform 0.1s;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: inherit;
  letter-spacing: -0.01em;
}

.login-btn:hover { background: #4f46e5; }
.login-btn:active { transform: scale(0.98); }
.login-btn:disabled { opacity: 0.6; cursor: not-allowed; transform: none; }

.loading-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.6s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.help-text {
  text-align: center;
  color: #697386;
  font-size: 13px;
  margin-top: 16px;
  cursor: pointer;
}

.help-text:hover { color: #635bff; }

/* 弹窗样式 */
.dialog-mask {
  position: fixed;
  inset: 0;
  background: rgba(10,37,64,0.6);
  backdrop-filter: blur(4px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  padding: 24px;
}

.dialog-box {
  position: relative;
  width: 100%;
  max-width: 320px;
  background: #ffffff;
  border-radius: 20px;
  padding: 32px 24px 28px;
  text-align: center;
  box-shadow: 0 20px 60px rgba(0,0,0,0.3);
}

.dialog-icon {
  width: 48px;
  height: 48px;
  margin: 0 auto 12px;
  background: #f0efff;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.dialog-icon img {
  width: 28px;
  height: 28px;
  object-fit: contain;
}

.dialog-title {
  font-size: 17px;
  font-weight: 700;
  color: #1a1f36;
  margin-bottom: 8px;
  letter-spacing: -0.02em;
}

.dialog-desc {
  margin-bottom: 24px;
}

.dialog-desc div {
  font-size: 14px;
  color: #697386;
  line-height: 1.6;
}

.dialog-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 180px;
  height: 44px;
  background: #635bff;
  border-radius: 10px;
  border: none;
  font-size: 15px;
  font-weight: 600;
  color: #ffffff;
  cursor: pointer;
  font-family: inherit;
  transition: background 0.15s;
}

.dialog-btn:hover { background: #4f46e5; }

.dialog-close {
  position: absolute;
  top: 16px;
  right: 16px;
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: #f8fafc;
  cursor: pointer;
}

.dialog-close img {
  width: 16px;
  height: 16px;
}
</style>

