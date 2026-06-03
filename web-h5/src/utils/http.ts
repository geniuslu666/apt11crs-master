import axios, { type AxiosInstance, type AxiosResponse, type InternalAxiosRequestConfig } from 'axios'
import { showToast } from 'vant'
import { encryptData, decrypt } from './encrypt'

const instance: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
    'x-language': 'zh-CN'
  },
})

// 请求拦截器
instance.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('H5_STAFF_TOKEN')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    console.log('request',config.data)

    // 对请求体进行加密（POST/PUT/PATCH 请求）
    if (config.data && ['post', 'put', 'patch'].includes(config.method?.toLowerCase() || '')) {
      try {
        // 使用 encryptData 自动添加 timestamp、nonce_str、sign 并加密
        const encryptedHex = encryptData(config.data)
        // Gateway 期望接收加密后的字符串（用引号包裹）
        config.data = JSON.stringify(encryptedHex)
      } catch (error) {
        console.error('加密请求失败:', error)
      }
    }

    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
instance.interceptors.response.use(
  (response: AxiosResponse) => {
    try {
      // Gateway 返回的数据结构: { resCode, resMessage, resData }
      const { resCode, resMessage, resData } = response.data

      // 检查 Gateway 层的响应码
      if (resCode !== 0) {
        showToast(resMessage || '请求失败')
        return Promise.reject(new Error(resMessage || '请求失败'))
      }

      // 解密响应数据
      if (resData && typeof resData === 'string') {
        const decryptedStr = decrypt(resData)
        const decryptedData = JSON.parse(decryptedStr)
        console.log('decryptedData',decryptedData)

        // 检查解密后的业务响应码
        if (decryptedData.code === 0) {
          return decryptedData.data
        } else {
          showToast(decryptedData.message || '请求失败')
          return Promise.reject(new Error(decryptedData.message || '请求失败'))
        }
      }

      // 如果没有加密数据，直接返回
      return resData
    } catch (error) {
      console.error('解密响应失败:', error)
      showToast('数据解析失败')
      return Promise.reject(error)
    }
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('H5_STAFF_TOKEN')
      localStorage.removeItem('H5_STAFF_INFO')
      window.location.href = '/daytrip/login'
    } else {
      showToast(error.message || '网络错误')
    }
    return Promise.reject(error)
  }
)

export default instance
