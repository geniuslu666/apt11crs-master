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
  background: #FFFFFF;
}

.nav-bar {
  height: 44px;
  background: #FFFFFF;
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
  color: #000000;
  font-weight: 300;
}

.nav-title {
  font-size: 18px;
  font-weight: 500;
  color: #000000;
  line-height: 26px;
}

.content {
  padding: 16px 20px 0;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.card {
  border: 1px solid #D8DCE5;
  background: #fff;
  border-radius: 10px;
  padding: 16px 12px;
}

.status-text {
  font-size: 20px;
  font-weight: 600;
  color: #269C74;
  line-height: 28px;
  margin-bottom: 10px;
}

.product-title {
  font-size: 14px;
  font-weight: 600;
  color: #3D3D3D;
  line-height: 20px;
  margin-bottom: 10px;
}

.info-row {
  display: flex;
  font-size: 12px;
  margin-bottom: 5px;
  color: #929292;
  line-height: 17px;
}

.card-title {
  font-size: 16px;
  line-height: 22px;
  font-weight: 600;
  color: #3D3D3D;
  margin-bottom: 16px;
}

.detail-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  line-height: 17px;
  margin-bottom: 10px;
  color: #3D3D3D;
}

.footer {
  padding: 35px 20px;
}
.btn-home{
  width: 100%;
  height: 42px;
  line-height: 42px;
  background: #12C584;
  border: none;
  border-radius: 10px;
  font-size: 16px;
  font-weight: 600;
  color: #FFFFFF;
  cursor: pointer;
  padding: 0;
}
</style>
