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
.login-page {
  min-height: 100vh;
  background: url('https://oss.yeebok.net/static/attachment/2026-03-12/dh0ch3648pzw55clvy.png') center center / cover no-repeat;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 40px;
  position: relative;
}

/* 语言切换 */
.lang-switcher {
  position: absolute;
  top: 20px;
  right: 20px;
  display: flex;
  gap: 8px;
}

.lang-item {
  font-size: 12px;
  color: #999;
  cursor: pointer;
  padding: 2px 6px;
  border-radius: 4px;
  transition: all 0.2s;
}

.lang-item.active {
  color: #1a9b6e;
  font-weight: 600;
}

/* Logo 区域 */
.logo-area {
  margin-top: 120px;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 60px;
}

.logo-wrapper {
  margin-bottom: 13px;
}

.logo-wrapper img{
  display: block;
  width: 135px;
  height: 52px;
}

.logo-icon {
  width: 64px;
  height: 64px;
}

.logo-icon svg {
  width: 100%;
  height: 100%;
}

.logo-text {
  display: flex;
  flex-direction: column;
}

.logo-en {
  font-size: 13px;
  font-weight: 700;
  color: #1a6b4a;
  line-height: 1.3;
  letter-spacing: 0.5px;
}

.logo-zh {
  font-size: 14px;
  color: #1a6b4a;
  margin-top: 2px;
}

.platform-title {
  font-size: 28px;
  font-weight: 500;
  color: #125D43;
  line-height: 39px;
  margin: 0;
  letter-spacing: 2px;
}

/* 表单区域 */
.form-area {
  width: 100%;
  max-width: 400px;
}

.input-item {
  background: white;
  border-radius: 10px;
  padding: 0 16px;
  height: 54px;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  box-shadow: 0px 8px 24px 0px rgba(118,219,187,0.2);
}

.input-icon {
  width: 20px;
  height: 20px;
  margin-right: 16px;
}


.input-field {
  flex: 1;
  font-weight: 400;
  border: none;
  outline: none;
  font-size: 16px;
  color: #3D3D3D;
  background: transparent;
  line-height: 54px;
}

.input-field::placeholder {
  color: #ADAEB0;
}

.login-btn {
  width: 100%;
  height: 54px;
  background: linear-gradient( 90deg, #12C584 0%, #2BBAAE 100%);
  border: none;
  border-radius: 10px;
  color: white;
  font-size: 16px;
  font-weight: 600;
  margin-top: 20px;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0px 8px 24px 0px rgba(118,219,187,0.2);
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-btn:active {
  transform: scale(0.98);
}

.login-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loading-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.6s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.help-text {
  text-align: center;
  color: #979797;
  font-weight: 400;
  font-size: 14px;
  line-height: 20px;
  margin-top: 20px;
}

/* 弹窗样式 */
.dialog-mask {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.dialog-box {
  position: relative;
  width: 335px;
  background: #FFFFFF;
  border-radius: 10px;
  padding: 34px;
  text-align: center;
}

.dialog-icon {
  width: 35px;
  height: 35px;
  margin: 0 auto 5px;
}

.dialog-icon img{
  width: 100%;
  width: 100%;
}

.dialog-title {
  font-size: 18px;
  font-weight: 600;
  color: #3D3D3D;
  line-height: 25px;
  margin-bottom: 12px;
}

.dialog-desc {
  margin-bottom: 24px;
}

.dialog-desc div{
  font-size: 14px;
  font-weight: 400;
  color: #999999;
  line-height: 20px;
  text-align: center;
}

.dialog-btn {
  width: 180px;
  height: 42px;
  background: #12C584;
  border-radius: 4px;
  border: none;
  font-size: 16px;
  line-height: 42px;
  font-weight: 600;
  color: #FFFFFF;
  cursor: pointer;
  padding: 0;
}

.dialog-close{
  position: absolute;
  top: 20px;
  right: 20px;
  width: 24px;
  height: 24px;
}

.dialog-close img {
  width: 100%;
  height: 100%;
}
</style>

