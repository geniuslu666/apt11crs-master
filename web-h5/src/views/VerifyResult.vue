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
  background: #f6f9fc;
  position: relative;
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

.result-content {
  padding: 0 16px;
}

/* 成功/失败 icon */
.result-icon {
  width: 72px;
  height: 72px;
  border-radius: 50%;
  margin: 40px auto 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #e6faf4;
}

.result-icon.error-icon { background: #fef2f2; }

.result-icon img {
  width: 40px;
  height: 40px;
  object-fit: contain;
}

.result-title {
  font-size: 20px;
  font-weight: 700;
  color: #1a1f36;
  text-align: center;
  margin-bottom: 6px;
  letter-spacing: -0.03em;
}

.result-desc {
  font-size: 14px;
  color: #697386;
  line-height: 1.5;
  text-align: center;
  margin-bottom: 28px;
}

.result-desc.error { color: #ef4444; }

.result-card {
  background: #ffffff;
  border: 1px solid #e3e8ef;
  border-radius: 14px;
  padding: 16px;
  margin-bottom: 16px;
  box-shadow: 0 1px 3px rgba(60,66,87,0.06);
}

.card-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #f0f3f7;
}

.card-row:last-child { border-bottom: none; }

.card-row .label {
  font-size: 13px;
  color: #697386;
}

.card-row .value {
  font-size: 13px;
  font-weight: 600;
  color: #1a1f36;
  text-align: right;
  max-width: 60%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.btn-primary {
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
  margin-bottom: 10px;
  font-family: inherit;
  transition: background 0.15s, transform 0.1s;
}

.btn-primary:hover { background: #4f46e5; }
.btn-primary:active { transform: scale(0.98); }

.btn-secondary {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 48px;
  background: #ffffff;
  border: 1px solid #e3e8ef;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 600;
  color: #697386;
  cursor: pointer;
  font-family: inherit;
  transition: background 0.12s;
}

.btn-secondary:active { background: #f8fafc; }
</style>
