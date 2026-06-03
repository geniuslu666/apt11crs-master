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
.index-page {
  min-height: 100vh;
  background: #FFFFFF;
}

/* 用户信息区域 */
.user-header {
  padding: 50px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  overflow: hidden;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.info {
  padding-top: 3px;
}

.info .name {
  font-size: 16px;
  font-weight: 500;
  color: #3D3D3D;
  line-height: 22px;
  margin-bottom: 2px;
}

.info .role {
  font-size: 12px;
  color: #979797;
  line-height: 17px;
  font-weight: 400;
}

.logout-btn {
  cursor: pointer;
}

.logout-btn img {
  width: 20px;
  height: 20px;
}

/* 扫码按钮区域 */
.scan-section {
  padding: 0px 20px 40px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.scan-button {
  width: 195px;
  height: 195px;
  border-radius: 50%;
  background: #12C584;
  box-shadow: 0px 5px 10px 0px rgba(0,149,45,0.3);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: all 0.3s;
}

.scan-button:active {
  transform: scale(0.95);
}

.scan-icon {
  margin-bottom: 10px;
}

.scan-icon img {
  width: 66px;
  height: 66px;
}

.scan-text {
  font-size: 18px;
  color: #fff;
  font-weight: 400;
  line-height: 25px;
}

/* 核销记录区域 */
.verify-section {
  padding: 0 20px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.section-header .title {
  font-size: 16px;
  font-weight: 600;
  color: #3D3D3D;
  line-height: 22px;
}

.section-header .more {
  font-size: 14px;
  color: #979797;
  font-weight: 400;
  line-height: 20px;
  display: flex;
  align-items: center;
  cursor: pointer;
}

.section-header img {
  width: 16px;
  height: 16px;
}

.verify-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.verify-list-empty {
  padding: 100px 0;
}

.verify-list-empty img {
  display: block;
  margin: 0 auto 11px;
  width: 70px;
  height: 70px;
}

.verify-list-empty div {
  font-weight: 400;
  font-size: 12px;
  color: #C9CDD4;
  line-height: 17px;
  text-align: center;
}

.verify-item {
  border: 1px solid #D8DCE5;
  border-radius: 10px;
  padding: 16px 12px 13px;
  cursor: pointer;
  margin-bottom: 10px;
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 7px;
}

.item-header .time {
  font-weight: 600;
  font-size: 14px;
  color: #3D3D3D;
  line-height: 20px;
}

.item-header .status {
  font-weight: 600;
  font-size: 12px;
  color: #269C74;
  line-height: 17px;
}

.item-content {
  font-weight: 400;
  font-size: 12px;
  line-height: 17px;
  color: #929292;
  margin-bottom: 5px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.scan-button.loading {
  opacity: 0.8;
  cursor: not-allowed;
}

/* 方案一：内嵌摄像头扫码区域（当前使用） */
.camera-box {
  width: 100%;
  max-width: 400px;
  margin: 0 auto;
  position: relative;
  border-radius: 12px;
  overflow: hidden;
  background: #000;
}

#inline-qr-reader {
  width: 100%;
}

.cancel-scan {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  padding: 10px 30px;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 20px;
  font-size: 14px;
  color: #333;
  cursor: pointer;
  z-index: 10;
}

/* 方案二：全屏摄像头覆盖层（已注释，需要时取消注释并注释掉方案一）
.camera-fullscreen {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: #000;
  z-index: 999;
  display: flex;
  flex-direction: column;
}

.camera-header {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 50px 20px 20px;
  position: relative;
}

.camera-title {
  font-size: 18px;
  font-weight: 500;
  color: #FFFFFF;
}

.camera-close {
  position: absolute;
  right: 20px;
  font-size: 20px;
  color: #FFFFFF;
  cursor: pointer;
  padding: 4px 8px;
}

.camera-container {
  flex: 1;
  width: 100%;
}

.camera-tip {
  padding: 20px;
  text-align: center;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.7);
}
*/

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
  font-size: 14px;
  font-weight: 400;
  color: #999999;
  line-height: 20px;
  margin-bottom: 24px;
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
</style>
