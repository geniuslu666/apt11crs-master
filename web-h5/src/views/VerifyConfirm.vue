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
  background: #FFFFFF;
  padding-bottom: 80px;
}

.nav-bar {
  height: 44px;
  background: #12C584;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  padding: 0 16px;
}

.nav-back {
  position: absolute;
  left: 16px;
  cursor: pointer;
}

.back-arrow {
  font-size: 28px;
  color: #FFFFFF;
  font-weight: 300;
}

.nav-title {
  font-size: 18px;
  font-weight: 500;
  color: #FFFFFF;
  line-height: 25px;
}

.bg{
  width: 100%;
  height: 67px;
  background: #12C584;
}

.content {
  background: #FFFFFF;
  margin-top: -67px;
  padding: 20px;
  border-radius: 10px 10px 0 0;
}

.section-title {
  font-size: 16px;
  font-weight: 500;
  color: #3D3D3D;
  line-height: 26px;
  margin-bottom: 10px;
}

.info-card {
  background: #FFFFFF;
  border: 1px solid #D8DCE5;
  border-radius: 10px;
  padding: 16px 12px;
  margin-bottom: 20px;
}

.product-title {
  font-size: 14px;
  font-weight: 600;
  color: #3D3D3D;
  line-height: 20px;
  margin-bottom: 9px;
}

.info-row {
  font-size: 12px;
  font-weight: 400;
  color: #929292;
  line-height: 17px;
  margin-bottom: 5px;
  display: flex;
}

.info-row:last-child {
  margin-bottom: 0;
}

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
  margin-bottom: 10px;
}

.order-row:last-child {
  margin-bottom: 0;
}

.order-label {
  font-size: 12px;
  font-weight: 400;
  color: #3D3D3D;
  line-height: 17px;
}

.order-value {
  font-size: 12px;
  font-weight: 600;
  color: #3D3D3D;
  line-height: 17px;
  text-align: right;
}

.order-value.price {
  font-size: 12px;
  font-weight: 600;
  color: #FF5A60;
}

.bottom-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: #FFFFFF;
  padding: 12px;
  display: flex;
  gap: 12px;
  border-top: 1px solid #F4F4F4;
}

.btn-cancel {
  flex: 1;
  height: 42px;
  background: #FFFFFF;
  border: 1px solid #C9CDD4;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 500;
  line-height: 42px;
  color: #3D3D3D;
  cursor: pointer;
  padding: 0;
}

.btn-confirm {
  flex: 2;
  height: 42px;
  background: #12C584;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 500;
  line-height: 42px;
  color: #FFFFFF;
  cursor: pointer;
  padding: 0;
}

.btn-confirm:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>

