<template>
  <div class="result-page">
     <div class="nav-bar">
      <div class="nav-back" @click="router.back()">
        <span class="back-arrow">‹</span>
      </div>
    </div>

    <div class="result-content">
      <!-- 成功 -->
      <template v-if="isSuccess">
        <div class="result-icon">
          <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0o2jj1jbd6nevw7k.png" />
        </div>
        <div class="result-title">核销成功</div>
        <div class="result-card">
          <div class="card-row">
            <span class="label">订单编号</span>
            <span class="value">{{ orderSn }}</span>
          </div>
          <div class="card-row">
            <span class="label">核销时间</span>
            <span class="value">{{ verifyTime }}</span>
          </div>
        </div>
        <button class="btn-primary" @click="router.push('/')">返回首页</button>
      </template>

      <!-- 失败 -->
      <template v-else>
        <div class="result-icon">
          <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0o1cgfqn9lpqtqgi.png" />
        </div>
        <div class="result-title">核销失败</div>
        <div class="result-desc">{{ errorMsg }}</div>
        <button class="btn-primary" @click="router.back()">再试一次</button>
        <button class="btn-secondary" @click="router.push('/')">返回首页</button>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const isSuccess = computed(() => route.query.status === 'success')
const orderSn = route.query.orderSn as string
const verifyTime = route.query.verifyTime as string
const errorMsg = (route.query.msg as string) || '核销码已被使用，不能再次核销'
</script>

<style scoped>
.result-page {
  min-height: 100vh;
  background: #F4F4F4;
  position: relative;
}

.nav-bar {
  height: 44px;
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
  color: #3D3D3D;
  font-weight: 300;
}

.result-content {
  padding: 0 20px;
}

.result-icon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  margin: 38px auto 14px;
}

.result-icon img {
  width: 100%;
  height: 100%;
}

.result-title {
  font-size: 18px;
  font-weight: 500;
  line-height: 25px;
  color: #3D3D3D;
  text-align: center;
  margin-bottom: 6px;
}

.result-desc {
  font-weight: 400;
  font-size: 14px;
  color: #FF5A60;
  line-height: 20px;
  text-align: center;
  margin-bottom: 30px;
}

.result-card {
  background: #FFFFFF;
  border-radius: 10px;
  padding: 17px 12px 22px;
  margin: 40px 0 20px;
}

.card-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.card-row:last-child { margin-bottom: 0; }

.card-row .label {
  font-size: 12px;
  color: #3D3D3D;
  line-height: 17px;
  font-weight: 400;
}

.card-row .value {
  font-size: 12px;
  color: #3D3D3D;
  line-height: 17px;
  font-weight: 400;
}

.btn-primary {
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
  margin-bottom: 14px;
  padding: 0;
}

.btn-secondary {
  width: 100%;
  height: 42px;
  line-height: 42px;
  background: #FFFFFF;
  border: 1px solid #C9CDD4;
  border-radius: 10px;
  font-size: 16px;
  font-weight: 600;
  color: #3D3D3D;
  cursor: pointer;
  padding: 0;
}
</style>
