<template>
  <div class="verify-list-page">
    <!-- 顶栏 -->
    <div class="nav-bar">
      <div class="nav-back" @click="router.back()">
        <span class="back-arrow">‹</span>
      </div>
      <div class="nav-title">核销记录</div>
    </div>
    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        :finished-text="list.length === 0 ? '' : '没有更多了'"
        @load="onLoad"
      >
        <div
          v-for="item in list"
          :key="item.id"
          class="record-card"
          @click="goToDetail(item.id)"
        >
          <div class="card-header">
            <span class="verify-time">{{ item.verifyTime }}</span>
            <span class="status-tag">已核销</span>
          </div>
          <div class="card-row">订单号：{{ item.orderSn }}</div>
          <div class="card-row">预约人：{{ item.orderInfo?.bookingName }} {{ item.orderInfo?.phoneArea }}{{ item.orderInfo?.bookingMobile }}</div>
          <div class="card-row">一日游：{{ item.productInfo?.title }}</div>
        </div>
        <!-- <van-empty v-if="finished && list.length === 0" description="暂无核销记录" /> -->
      </van-list>
      <div v-if="finished && list.length === 0" class="empty-wrap">
        <img src="https://oss.yeebok.net/static/attachment/2026-03-12/dh0gt1cd80q5glzqjh.png" alt="empty" />
        <div>暂无记录</div>
      </div>
    </van-pull-refresh>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { getVerifyLog, type VerifyLogItem } from '../api/travelStaff'

const router = useRouter()
const list = ref<VerifyLogItem[]>([])
const loading = ref(false)
const finished = ref(false)
const refreshing = ref(false)
const pageNum = ref(1)
const pageSize = 20

const onLoad = async () => {
  try {
    const res = await getVerifyLog({ pageNum: pageNum.value, pageSize })
    list.value.push(...res.list)
    pageNum.value++
    loading.value = false
    if (res.list.length < pageSize) {
      finished.value = true
    }
  } catch {
    loading.value = false
    finished.value = true
  }
}

const onRefresh = async () => {
  pageNum.value = 1
  list.value = []
  finished.value = false
  loading.value = true
  await onLoad()
  refreshing.value = false
}

const goToDetail = (id: number) => {
  router.push(`/verify/detail/${id}`)
}
</script>

<style scoped>
.verify-list-page {
  min-height: 100vh;
  padding-bottom: 24px;
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

.record-card {
  margin: 10px 16px 0;
  padding: 14px 16px;
  background: #fff;
  border-radius: 14px;
  border: 1px solid #e3e8ef;
  cursor: pointer;
  transition: box-shadow 0.15s, border-color 0.15s;
  box-shadow: 0 1px 3px rgba(60,66,87,0.06);
}

.record-card:active {
  box-shadow: 0 4px 12px rgba(60,66,87,0.1);
  border-color: #c8d2e0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.verify-time {
  font-size: 13px;
  font-weight: 600;
  color: #1a1f36;
}

.status-tag {
  display: inline-flex;
  align-items: center;
  padding: 2px 8px;
  border-radius: 100px;
  font-size: 11px;
  font-weight: 600;
  background: #e6faf4;
  color: #00875a;
}

.card-row {
  font-size: 12px;
  color: #697386;
  line-height: 1.6;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  margin-bottom: 2px;
}

.card-row:last-child { margin-bottom: 0; }

.empty-wrap {
  height: calc(100vh - 56px);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 12px;
}

.empty-wrap img {
  width: 64px;
  height: 64px;
  opacity: 0.45;
}

.empty-wrap div {
  font-size: 13px;
  color: #9da9bb;
}
</style>
