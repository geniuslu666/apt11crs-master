<template>
  <div class="index-page">
    <!-- 用户信息区域 -->
    <div class="user-header">
      <div class="user-info">
        <div class="avatar">
          <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0gj79ic51dgkfrkz.png" alt="avatar" />
        </div>
        <div class="info">
          <div class="name">{{ staffInfo?.name || staffInfo?.username || 'admin' }}</div>
          <div class="role">管理员</div>
        </div>
      </div>
      <div class="logout-btn" @click="handleLogout">
         <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0glhxkq9kcznhx0u.png" alt="logout" />
      </div>
    </div>

    <!-- 扫码区域 -->
    <div class="scan-section">
      <!-- 摄像头预览 -->
      <div v-if="scanning" class="camera-box">
        <div id="inline-qr-reader"></div>
        <div class="cancel-scan" @click="stopScan">取消</div>
      </div>
      <!-- 扫码按钮 -->
      <div v-else class="scan-button" :class="{ loading: cameraLoading }" @click="!cameraLoading && startScan()">
        <div class="scan-icon">
          <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0gnuppf426nbnbeb.png" alt="scan" />
        </div>
        <div class="scan-text">{{ cameraLoading ? '启动中...' : '点击扫码' }}</div>
      </div>
    </div>

    <!-- 方案二：全屏摄像头覆盖层（已注释，需要时取消注释并注释掉方案一）
    <div v-if="scanning" class="camera-fullscreen">
      <div class="camera-header">
        <div class="camera-title">扫描二维码</div>
        <div class="camera-close" @click="stopScan">✕</div>
      </div>
      <div id="inline-qr-reader" class="camera-container"></div>
      <div class="camera-tip">将二维码放入框内，即可自动扫描</div>
    </div>
    -->

    <!-- 二维码识别失败弹窗 -->
    <div v-if="showFailDialog" class="dialog-mask">
      <div class="dialog-box">
        <div class="dialog-icon">
          <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0mdkboxlo5u5mnzi.png" alt="error" />
        </div>
        <div class="dialog-title">二维码识别失败</div>
        <div class="dialog-desc">无法识别二维码。请在继续验证前确认相关信息。</div>
        <button class="dialog-btn" @click="showFailDialog = false">好的</button>
      </div>
    </div>

    <!-- 核销记录 -->
    <div class="verify-section">
      <div class="section-header">
        <span class="title">核销记录</span>
        <span class="more" @click="goToList">
          查看全部
          <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0gqx2iwkbyoqjklc.png" alt="arrow" />
        </span>
      </div>
      <div class="verify-list">
        <div class="verify-list-empty" v-if="list.length === 0">
            <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0gt1cd80q5glzqjh.png" alt="empty" />
            <div>暂无记录</div>
        </div>
        <div v-else class="verify-item" v-for="item in list" :key="item.id" @click="goToDetail(item.id)">
          <div class="item-header">
            <span class="time">{{ item.verifyTime }}</span>
            <span class="status">已核销</span>
          </div>
          <div class="item-content">订单号：{{ item.orderSn }}</div>
          <div class="item-content">预约人：{{ item.orderInfo?.bookingName }} {{ item.orderInfo?.phoneArea }}{{ item.orderInfo?.bookingMobile }}</div>
          <div class="item-content">一日游：{{ item.productInfo?.title }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { showConfirmDialog } from 'vant'
import { Html5Qrcode } from 'html5-qrcode'
import { useTravelStaffStore } from '../store/travelStaff'
import { getVerifyLog, staffLogout, getCodeView, type VerifyLogItem } from '../api/travelStaff'

const router = useRouter()
const staffStore = useTravelStaffStore()
const staffInfo = computed(() => staffStore.staffInfo)
const list = ref<VerifyLogItem[]>([])

const scanning = ref(false)
const cameraLoading = ref(false)
const showFailDialog = ref(false)
let html5QrCode: Html5Qrcode | null = null

const loadVerifyLog = async () => {
  try {
    const res = await getVerifyLog({pageNum: 1, pageSize: 5})
    list.value = res.list
  } catch (error) {
    console.error('加载核销记录失败:', error)
  }
}

const startScan = async () => {
  cameraLoading.value = true
  scanning.value = true
  await nextTick()
  try {
    html5QrCode = new Html5Qrcode('inline-qr-reader')
    await html5QrCode.start(
      { facingMode: 'environment' },
      { fps: 10, qrbox: 220 },
      (decodedText) => handleScanResult(decodedText),
      () => {}
    )
  } catch {
    scanning.value = false
    showFailDialog.value = true
  } finally {
    cameraLoading.value = false
  }
}

const handleScanResult = async (code: string) => {
  await stopScan()
  try {
    await getCodeView(code)
    router.push({ path: '/verify/confirm', query: { code } })
  } catch {
    showFailDialog.value = true
  }
}

const stopScan = async () => {
  try {
    await html5QrCode?.stop()
    html5QrCode?.clear()
  } catch { /* ignore */ }
  html5QrCode = null
  scanning.value = false
}

const goToList = () => {
  router.push('/verify/list')
}

const goToDetail = (id: number) => {
  router.push(`/verify/detail/${id}`)
}

const handleLogout = async () => {
  try {
    await showConfirmDialog({
      title: '提示',
      message: '确定要退出登录吗？',
    })
    await staffLogout()
    staffStore.clear()
    router.replace('/login')
  } catch {
    // 用户取消
  }
}

onMounted(() => {
  loadVerifyLog()
})

onUnmounted(() => {
  stopScan()
})
</script>

<style scoped>
/* Stripe-style Index/Dashboard page */
.index-page {
  min-height: 100vh;
  background: #f6f9fc;
}

/* 用户信息区域 */
.user-header {
  padding: 52px 20px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #ffffff;
  border-bottom: 1px solid #e3e8ef;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  overflow: hidden;
  border: 2px solid #e3e8ef;
  flex-shrink: 0;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.info .name {
  font-size: 15px;
  font-weight: 600;
  color: #1a1f36;
  margin-bottom: 2px;
  letter-spacing: -0.01em;
}

.info .role {
  font-size: 12px;
  color: #697386;
  font-weight: 400;
}

.logout-btn {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  background: #f8fafc;
  border: 1px solid #e3e8ef;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.15s;
}

.logout-btn:active { background: #f0f3f7; }

.logout-btn img {
  width: 18px;
  height: 18px;
  opacity: 0.7;
}

/* 扫码按钮区域 */
.scan-section {
  padding: 40px 20px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.scan-button {
  width: 180px;
  height: 180px;
  border-radius: 50%;
  background: linear-gradient(135deg, #635bff 0%, #4f46e5 100%);
  box-shadow: 0 8px 32px rgba(99,91,255,0.4), 0 0 0 8px rgba(99,91,255,0.08);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: transform 0.15s, box-shadow 0.15s;
}

.scan-button:active {
  transform: scale(0.94);
  box-shadow: 0 4px 16px rgba(99,91,255,0.3), 0 0 0 8px rgba(99,91,255,0.06);
}

.scan-button.loading {
  opacity: 0.75;
  cursor: not-allowed;
}

.scan-icon {
  margin-bottom: 10px;
}

.scan-icon img {
  width: 56px;
  height: 56px;
  filter: brightness(0) invert(1);
}

.scan-text {
  font-size: 16px;
  color: #fff;
  font-weight: 600;
  letter-spacing: -0.01em;
}

/* 摄像头扫码区域 */
.camera-box {
  width: 100%;
  max-width: 400px;
  margin: 0 auto;
  position: relative;
  border-radius: 16px;
  overflow: hidden;
  background: #000;
  box-shadow: 0 8px 32px rgba(0,0,0,0.2);
}

#inline-qr-reader {
  width: 100%;
}

.cancel-scan {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  padding: 10px 28px;
  background: rgba(255,255,255,0.95);
  border-radius: 100px;
  font-size: 14px;
  font-weight: 500;
  color: #1a1f36;
  cursor: pointer;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(0,0,0,0.15);
}

/* 核销记录区域 */
.verify-section {
  padding: 0 16px 24px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding: 0 4px;
}

.section-header .title {
  font-size: 15px;
  font-weight: 700;
  color: #1a1f36;
  letter-spacing: -0.01em;
}

.section-header .more {
  font-size: 13px;
  color: #635bff;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 2px;
  cursor: pointer;
}

.section-header img {
  width: 14px;
  height: 14px;
}

.verify-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.verify-list-empty {
  padding: 60px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
}

.verify-list-empty img {
  width: 64px;
  height: 64px;
  opacity: 0.5;
}

.verify-list-empty div {
  font-size: 13px;
  color: #9da9bb;
}

.verify-item {
  background: #ffffff;
  border: 1px solid #e3e8ef;
  border-radius: 14px;
  padding: 14px 16px;
  cursor: pointer;
  transition: box-shadow 0.15s, border-color 0.15s;
  box-shadow: 0 1px 3px rgba(60,66,87,0.06);
}

.verify-item:active {
  box-shadow: 0 4px 12px rgba(60,66,87,0.1);
  border-color: #c8d2e0;
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.item-header .time {
  font-weight: 600;
  font-size: 13px;
  color: #1a1f36;
}

.item-header .status {
  display: inline-flex;
  align-items: center;
  padding: 2px 8px;
  border-radius: 100px;
  font-size: 11px;
  font-weight: 600;
  background: #e6faf4;
  color: #00875a;
}

.item-content {
  font-size: 12px;
  line-height: 1.6;
  color: #697386;
  margin-bottom: 2px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* 弹窗样式 */
.dialog-mask {
  position: fixed;
  inset: 0;
  background: rgba(10,37,64,0.5);
  backdrop-filter: blur(4px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  padding: 24px;
}

.dialog-box {
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
  background: #fef2f2;
  border-radius: 12px;
  margin: 0 auto 12px;
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
  font-size: 14px;
  color: #697386;
  line-height: 1.6;
  margin-bottom: 24px;
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
</style>
