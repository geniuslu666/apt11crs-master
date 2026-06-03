import { createI18n } from 'vue-i18n'
import zhCN from './zh-CN'
import en from './en'
import ja from './ja'

const i18n = createI18n({
  legacy: false,
  locale: localStorage.getItem('lang') || 'zh-CN',
  fallbackLocale: 'zh-CN',
  messages: {
    'zh-CN': zhCN,
    en,
    ja,
  },
})

export default i18n
