<template>
  <div class="confirm-page">
    <!-- 顶栏 -->
    <div class="nav-bar">
      <div class="nav-back" @click="router.back()">
        <span class="back-arrow">‹</span>
      </div>
      <div class="nav-title">核销确认</div>
    </div>

    <div class="bg"></div>

    <div v-if="orderInfo" class="content">
      <!-- 出游信息 -->
      <div class="section-title">出游信息</div>
      <div class="info-card">
        <div class="product-title">{{ orderInfo.productInfo?.title }}</div>
        <div class="info-row">出游时间：{{ orderInfo.bookDate }}</div>
        <div class="info-row">预约人：{{ orderInfo.bookingName }} {{ orderInfo.phoneArea }}{{ orderInfo.bookingMobile }}</div>
        <div class="info-row">人数：{{ orderInfo.bookingNum }}人</div>
        <div class="info-row" v-if="orderInfo.productInfo?.meetingTime">集合时间：{{ orderInfo.bookDate }} {{ orderInfo.productInfo?.meetingTime }}</div>
        <div class="info-row" v-if="orderInfo.productInfo?.meetingPlace">集合地点：<div>{{ orderInfo.productInfo?.meetingPlace }}</div></div>
      </div>

      <!-- 订单信息 -->
      <div class="section-title">订单信息</div>
      <div class="info-card">
        <div class="order-row">
          <span class="order-label">订单编号</span>
          <span class="order-value">{{ orderInfo.orderSn }}</span>
        </div>
        <div class="order-row">
          <span class="order-label">下单时间</span>
          <span class="order-value">{{ orderInfo.createdAt }}</span>
        </div>
        <div class="order-row">
          <span class="order-label">支付金额</span>
          <span class="order-value price">{{ orderInfo.orderAmount }}JPY</span>
        </div>
      </div>
    </div>

    <!-- 底部按钮 -->
    <div class="bottom-bar" v-if="orderInfo">
      <button class="btn-cancel" @click="router.back()">取消</button>
      <button class="btn-confirm" @click="handleConfirm" :disabled="confirming">
        {{ confirming ? '核销中...' : '确认核销' }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showToast, showLoadingToast, closeToast } from 'vant'
import { getCodeView, verifyCode } from '../api/travelStaff'

const router = useRouter()
const route = useRoute()
const confirming = ref(false)
const orderInfo = ref<any>(null)

onMounted(async () => {
  const code = route.query.code as string
  if (!code) {
    showToast('缺少核销码')
    router.back()
    return
  }
  showLoadingToast({ message: '加载中...', forbidClick: true, duration: 0 })
  try {
    orderInfo.value = await getCodeView(code)
  } catch (error: any) {
    showToast(error.message || '加载订单信息失败')
    router.back()
  } finally {
    closeToast()
  }
})

const handleConfirm = async () => {
  confirming.value = true
  try {
    const res = await verifyCode(orderInfo.value.orderSn)
    router.replace({
      path: '/verify/result',
      query: { status: 'success', orderSn: res.orderSn, verifyTime: res.verifyTime },
    })
  } catch (error: any) {
    router.replace({
      path: '/verify/result',
      query: { status: 'fail', msg: error.message || '核销失败' },
    })
  } finally {
    confirming.value = false
  }
}
</script>

<style scoped>
.confirm-page {
  min-height: 100vh;
  background: #f6f9fc;
  padding-bottom: 88px;
}

.nav-bar {
  height: 56px;
  background: #635bff;
  display: flex;
  align-items: center;
  justify-content: center;
  position: sticky;
  top: 0;
  z-index: 10;
  padding: 0 16px;
}

.nav-back {
  position: absolute;
  left: 12px;
  width: 36px;
  height: 36px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.12s;
}

.nav-back:active { background: rgba(255,255,255,0.15); }

.back-arrow {
  font-size: 24px;
  color: #ffffff;
  font-weight: 400;
  line-height: 1;
}

.nav-title {
  font-size: 16px;
  font-weight: 700;
  color: #ffffff;
  letter-spacing: -0.02em;
}

.bg {
  width: 100%;
  height: 48px;
  background: #635bff;
}

.content {
  background: transparent;
  margin-top: -48px;
  padding: 0 16px 16px;
  border-radius: 20px 20px 0 0;
  overflow: hidden;
}

.section-title {
  font-size: 14px;
  font-weight: 700;
  color: #1a1f36;
  margin-bottom: 10px;
  letter-spacing: -0.01em;
  margin-top: 16px;
}

.info-card {
  background: #ffffff;
  border: 1px solid #e3e8ef;
  border-radius: 14px;
  padding: 16px;
  margin-bottom: 12px;
  box-shadow: 0 1px 3px rgba(60,66,87,0.06);
}

.product-title {
  font-size: 14px;
  font-weight: 600;
  color: #1a1f36;
  margin-bottom: 8px;
  line-height: 1.4;
}

.info-row {
  font-size: 12px;
  color: #697386;
  line-height: 1.6;
  margin-bottom: 3px;
  display: flex;
}

.info-row:last-child { margin-bottom: 0; }

.info-row div {
  flex: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.order-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #f0f3f7;
}

.order-row:last-child { border-bottom: none; }

.order-label {
  font-size: 13px;
  color: #697386;
}

.order-value {
  font-size: 13px;
  font-weight: 600;
  color: #1a1f36;
  text-align: right;
}

.order-value.price {
  color: #ef4444;
}

.bottom-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: #ffffff;
  padding: 12px 16px;
  display: flex;
  gap: 10px;
  border-top: 1px solid #e3e8ef;
  box-shadow: 0 -4px 12px rgba(60,66,87,0.06);
}

.btn-cancel {
  flex: 1;
  height: 48px;
  background: #ffffff;
  border: 1px solid #e3e8ef;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  color: #697386;
  cursor: pointer;
  font-family: inherit;
  transition: background 0.12s;
}

.btn-cancel:active { background: #f8fafc; }

.btn-confirm {
  flex: 2;
  height: 48px;
  background: #635bff;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  color: #ffffff;
  cursor: pointer;
  font-family: inherit;
  transition: background 0.15s, transform 0.1s;
}

.btn-confirm:hover { background: #4f46e5; }
.btn-confirm:active { transform: scale(0.98); }
.btn-confirm:disabled { opacity: 0.6; cursor: not-allowed; transform: none; }
</style>

