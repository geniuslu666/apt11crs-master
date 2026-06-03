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
  padding-bottom: 20px;
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

.record-card {
  margin: 0 20px 10px;
  padding: 16px 12px 18px;
  background: #fff;
  border-radius: 10px;
  border: 1px solid #D8DCE5;
  cursor: pointer;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 7px;
}

.verify-time {
  font-size: 14px;
  font-weight: 600;
  line-height: 20px;
  color: #3D3D3D;
}

.status-tag {
  font-size: 12px;
  color: #269C74;
  font-weight: 600;
  line-height: 17px;
}

.card-row {
  margin-bottom: 5px;
  font-size: 12px;
  color: #929292;
  line-height: 17px;
  font-weight: 400;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  margin-bottom: 5px;
}

.card-row:last-child{
  margin-bottom: 0;
}

.empty-wrap {
  height: calc(100vh - 44px);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.empty-wrap img {
  display: block;
  margin: 0 auto 11px;
  width: 70px;
  height: 70px;
}

.empty-wrap div {
  font-weight: 400;
  font-size: 12px;
  color: #C9CDD4;
  line-height: 17px;
  text-align: center;
}
</style>
