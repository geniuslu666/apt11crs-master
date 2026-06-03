<template>
  <div class="verify-detail-page">
    <!-- 顶栏 -->
    <div class="nav-bar">
      <div class="nav-back" @click="router.back()">
        <span class="back-arrow">‹</span>
      </div>
      <div class="nav-title">核销详情</div>
    </div>

    <div v-if="detail" class="content">
      <div class="card">
        <div class="status-text">已核销</div>
        <div class="product-title">{{ detail.productInfo?.title }}</div>
        <div class="info-row">订单号：{{ detail.orderSn }}</div>
        <div class="info-row">预约人：{{ detail.orderInfo?.bookingName }} {{ detail.orderInfo?.phoneArea }}{{ detail.orderInfo?.bookingMobile }}</div>
        <div class="info-row">出游时间：{{ detail.bookDate }}</div>
      </div>

      <div class="card">
        <div class="card-title">核销信息</div>
        <div class="detail-row">
          <span class="detail-label">核销时间</span>
          <span class="detail-value">{{ detail.verifyTime }}</span>
        </div>
        <div class="detail-row">
          <span class="detail-label">操作人</span>
          <span class="detail-value">{{ detail.staffName || '-' }}</span>
        </div>
      </div>
    </div>

    <div class="footer">
      <button class="btn-home" @click="router.replace('/')">返回首页</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showToast, showLoadingToast, closeToast } from 'vant'
import { getVerifyLogView, type VerifyLogDetailResponse } from '../api/travelStaff'

const router = useRouter()
const route = useRoute()
const detail = ref<VerifyLogDetailResponse | null>(null)

onMounted(async () => {
  showLoadingToast({ message: '加载中...', forbidClick: true, duration: 0 })
  try {
    const id = Number(route.params.id)
    detail.value = await getVerifyLogView(id)
  } catch (error: any) {
    showToast(error.message || '加载失败')
  } finally {
    closeToast()
  }
})
</script>

<style scoped>
.verify-detail-page {
  min-height: 100vh;
  padding-bottom: 100px;
  background: #f6f9fc;
}

.nav-bar {
  height: 56px;
  background: #ffffff;
  border-bottom: 1px solid #e3e8ef;
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

.nav-back:active { background: #f0f3f7; }

.back-arrow {
  font-size: 24px;
  color: #1a1f36;
  font-weight: 400;
  line-height: 1;
}

.nav-title {
  font-size: 16px;
  font-weight: 700;
  color: #1a1f36;
  letter-spacing: -0.02em;
}

.content {
  padding: 16px 16px 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.card {
  border: 1px solid #e3e8ef;
  background: #fff;
  border-radius: 14px;
  padding: 16px;
  box-shadow: 0 1px 3px rgba(60,66,87,0.06);
}

.status-text {
  font-size: 18px;
  font-weight: 700;
  color: #00875a;
  margin-bottom: 8px;
  letter-spacing: -0.02em;
}

.product-title {
  font-size: 14px;
  font-weight: 600;
  color: #1a1f36;
  margin-bottom: 8px;
  line-height: 1.4;
}

.info-row {
  display: flex;
  font-size: 12px;
  margin-bottom: 4px;
  color: #697386;
  line-height: 1.6;
}

.card-title {
  font-size: 14px;
  font-weight: 700;
  color: #1a1f36;
  margin-bottom: 14px;
  letter-spacing: -0.01em;
}

.detail-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 13px;
  padding: 8px 0;
  border-bottom: 1px solid #f0f3f7;
  color: #1a1f36;
}

.detail-row:last-child { border-bottom: none; }

.detail-row .label { color: #697386; }

.footer {
  padding: 24px 16px;
}

.btn-home {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 48px;
  background: #635bff;
  border: none;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 600;
  color: #ffffff;
  cursor: pointer;
  font-family: inherit;
  transition: background 0.15s, transform 0.1s;
}

.btn-home:hover { background: #4f46e5; }
.btn-home:active { transform: scale(0.98); }
</style>
