import CryptoJS from 'crypto-js'

// Gateway 使用的 AES 密钥和 IV（与 server/internal/gateway/gateway.go 保持一致）
const AES_KEY = '202CB962AC59075B964B07152D234B70'
const AES_IV = '233FA6B19FDE8CAF'
// 签名密钥（与 gateway.go 的 SignKey 保持一致）
const SECURITY_KEY = 'D9840773233FA6B19FDE8CAF765402F5'

/**
 * 生成随机字符串（nonce）
 * @param length 字符串长度，默认 32
 * @returns 随机字符串
 */
function generateNonce(length: number = 32): string {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  let result = ''
  for (let i = 0; i < length; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  return result
}

/**
 * 对参数进行排序并生成签名字符串
 * @param params 参数对象
 * @returns 排序后的参数字符串（用于签名）
 */
function buildSignString(params: Record<string, any>): string {
  // 按 key 的 ASCII 码排序
  const sortedKeys = Object.keys(params).sort()
  // 构建 key=value 格式的字符串数组
  const paramPairs = sortedKeys.map(key => `${key}=${params[key]}`)
  // 拼接成字符串并添加 security key
  return paramPairs.join('&') + `&key=${SECURITY_KEY}`
}

/**
 * 生成签名
 * @param params 参数对象
 * @returns MD5 签名（大写）
 */
function generateSign(params: Record<string, any>): string {
  const signString = buildSignString(params)
  return CryptoJS.MD5(signString).toString().toUpperCase()
}

/**
 * 加密请求数据（完整流程：添加时间戳、随机串、签名，然后 AES 加密）
 * @param data 原始请求数据
 * @returns 16进制编码的密文字符串
 */
export function encryptData(data: any): string {
  // 创建参数对象
  let params: Record<string, any> = {}

  // 如果传入的数据是对象，则将其属性复制到 params
  if (typeof data === 'object' && data !== null) {
    params = { ...data }
  }

  // 添加时间戳（Unix 秒）
  const timestamp = Math.floor(Date.now() / 1000)
  params['timestamp'] = timestamp

  // 添加随机字符串
  const nonce = generateNonce()
  params['nonce_str'] = nonce

  // 生成签名
  const sign = generateSign(params)
  params['sign'] = sign

  // AES 加密
  const jsonStr = JSON.stringify(params)
  const key = CryptoJS.enc.Utf8.parse(AES_KEY)
  const iv = CryptoJS.enc.Utf8.parse(AES_IV)
  const encrypted = CryptoJS.AES.encrypt(jsonStr, key, {
    iv: iv,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7,
  })

  // 转换为 16 进制字符串
  return encrypted.ciphertext.toString(CryptoJS.enc.Hex)
}

/**
 * AES-CBC 加密（基础加密函数，不添加签名参数）
 * @param plainText 明文字符串
 * @returns 16进制编码的密文字符串
 */
export function encrypt(plainText: string): string {
  const key = CryptoJS.enc.Utf8.parse(AES_KEY)
  const iv = CryptoJS.enc.Utf8.parse(AES_IV)
  const encrypted = CryptoJS.AES.encrypt(plainText, key, {
    iv: iv,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7,
  })
  // 转换为 16 进制字符串（Gateway 使用 hex 编码）
  return encrypted.ciphertext.toString(CryptoJS.enc.Hex)
}

/**
 * AES-CBC 解密（匹配 Gateway 的解密方式）
 * @param cipherText 16进制编码的密文字符串
 * @returns 明文字符串
 */
export function decrypt(cipherText: string): string {
  const key = CryptoJS.enc.Utf8.parse(AES_KEY)
  const iv = CryptoJS.enc.Utf8.parse(AES_IV)
  // 从 16 进制字符串解析
  const encryptedHex = CryptoJS.enc.Hex.parse(cipherText)
  const encryptedBase64 = CryptoJS.enc.Base64.stringify(encryptedHex)
  const decrypted = CryptoJS.AES.decrypt(encryptedBase64, key, {
    iv: iv,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7,
  })
  return decrypted.toString(CryptoJS.enc.Utf8)
}
