<template>
  <div class="shadcn-dashboard">
    <n-spin :show="show" description="请稍候...">
      <div class="documents-frame hidden md:block">
        <header class="documents-header">
          <div class="documents-title">
            <button class="icon-button" type="button" aria-label="Sidebar">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="4" y="5" width="16" height="14" rx="2" />
                <path d="M9 5v14" />
              </svg>
            </button>
            <span class="header-separator"></span>
            <strong>主控台</strong>
          </div>
          <strong>{{ userStore.realName }}</strong>
        </header>

        <main class="documents-content">
          <section class="section-card-grid">
            <article v-for="card in featuredCards" :key="card.title" class="featured-card">
              <div class="featured-card-head">
                <span>{{ card.title }}</span>
                <span class="trend-badge">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path :d="card.up ? 'M7 17 17 7M9 7h8v8' : 'M7 7l10 10M17 9v8H9'" />
                  </svg>
                  {{ card.badge }}
                </span>
              </div>
              <div class="featured-value">{{ card.value }}</div>
              <div class="featured-caption">
                {{ card.caption }}
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path :d="card.up ? 'M7 17 17 7M9 7h8v8' : 'M7 7l10 10M17 9v8H9'" />
                </svg>
              </div>
              <p>{{ card.description }}</p>
            </article>
          </section>

          <section class="chart-card">
            <div class="chart-card-head">
              <div>
                <h2>七日订单趋势</h2>
                <p>{{ state.orderStatType === 'all' ? '全部渠道' : 'APP渠道' }}订单量变化</p>
              </div>
              <div class="range-tabs">
                <button type="button" :class="{ active: state.orderStatType === 'all' }" @click="handleUpdateValue('all')">全部</button>
                <button type="button" :class="{ active: state.orderStatType === 'hotel' }" @click="handleUpdateValue('hotel')">APP</button>
              </div>
            </div>
            <div ref="chartRef" class="featured-chart"></div>
          </section>

          <section class="table-toolbar">
            <div class="outline-tabs">
              <button type="button" :class="{ active: activePanel === 'orders' }" @click="activePanel = 'orders'">订单统计</button>
              <button type="button" :class="{ active: activePanel === 'property' }" @click="activePanel = 'property'">物业排行 <span>{{ state.dashboard.propertyList.length }}</span></button>
              <button type="button" :class="{ active: activePanel === 'ota' }" @click="activePanel = 'ota'">OTA排行 <span>{{ state.dashboard.oTAChannelList.length }}</span></button>
              <button type="button" :class="{ active: activePanel === 'country' }" @click="activePanel = 'country'">国家排行 <span>{{ state.dashboard.nationalityList.length }}</span></button>
            </div>
            <div class="toolbar-actions">
              <button type="button" class="toolbar-button" @click="Load">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <rect x="4" y="5" width="16" height="14" rx="2" />
                  <path d="M10 5v14" />
                </svg>
                {{ state.orderStatType === 'all' ? '全部' : 'APP' }}
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="m6 9 6 6 6-6" />
                </svg>
              </button>
              <button type="button" class="toolbar-button">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M12 5v14M5 12h14" />
                </svg>
                刷新数据
              </button>
            </div>
          </section>

          <section v-if="activePanel === 'orders'" class="data-shell">
            <table class="documents-table">
              <thead>
                <tr>
                  <th class="drag-cell"></th>
                  <th class="check-cell"><span class="fake-check"></span></th>
                  <th>指标</th>
                  <th>分类</th>
                  <th>状态</th>
                  <th>数值</th>
                  <th>同比</th>
                  <th>说明</th>
                  <th class="menu-cell"></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="row in documentRows" :key="row.name">
                  <td class="drag-cell">⋮⋮</td>
                  <td class="check-cell"><span class="fake-check"></span></td>
                  <td class="row-title">{{ row.name }}</td>
                  <td><span class="soft-pill">{{ row.category }}</span></td>
                  <td><span :class="['soft-pill', row.statusClass]">{{ row.status }}</span></td>
                  <td>{{ row.value }}</td>
                  <td>{{ row.trend }}</td>
                  <td>{{ row.remark }}</td>
                  <td class="menu-cell">⋮</td>
                </tr>
              </tbody>
            </table>
          </section>

          <section v-else class="data-shell">
            <table class="documents-table">
              <thead>
                <tr>
                  <th class="drag-cell"></th>
                  <th class="check-cell"><span class="fake-check"></span></th>
                  <th>排名</th>
                  <th>{{ rankingTitle }}</th>
                  <th>占比</th>
                  <th>说明</th>
                  <th class="menu-cell"></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="row in rankingRows" :key="`${activePanel}-${row.rank}`">
                  <td class="drag-cell">⋮⋮</td>
                  <td class="check-cell"><span class="fake-check"></span></td>
                  <td class="row-title">{{ row.rank }}</td>
                  <td><span class="soft-pill">{{ row.name }}</span></td>
                  <td>{{ row.rate }}%</td>
                  <td>{{ row.remark }}</td>
                  <td class="menu-cell">⋮</td>
                </tr>
              </tbody>
            </table>
          </section>
        </main>
      </div>

    </n-spin>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, reactive, Ref, ref } from 'vue';
import { useECharts } from '@/hooks/web/useECharts';
import { dashboardAll } from '@/api/comm';
import { hexToRgba } from '@/utils/artDesignUtils';
import { useUserStore } from '@/store/modules/user';
import * as echarts from 'echarts';

const show = ref(false);
const userStore = useUserStore();
const activePanel = ref<'orders' | 'property' | 'ota' | 'country'>('orders');
const chartRef = ref<HTMLDivElement | null>(null);
const chartRef30 = ref<HTMLDivElement | null>(null);
const chartRefMobile = ref<HTMLDivElement | null>(null);
const chartRef30Mobile = ref<HTMLDivElement | null>(null);
const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);
const { setOptions: setOptions30 } = useECharts(chartRef30 as Ref<HTMLDivElement>);
const { setOptions: setOptionsMobile } = useECharts(chartRefMobile as Ref<HTMLDivElement>);
const { setOptions: setOptions30Mobile } = useECharts(chartRef30Mobile as Ref<HTMLDivElement>);

const state = reactive({
  orderStatType: 'all',
  dashboard: {
    memberStat: {
      todayRegMemberNum: 0,
      todayMemberNumGrewPer: 0,
      totalMemberNum: 0,
      totalMemberGrewPer: 0,
      todayIncBal: 0,
      todayIncBalGrewPer: 0,
      todayConsumeBal: 0,
      todayConsumeBalGrewPer: 0,
    },
    orderStat: {
      allStat: {
        todayOrderNum: 0,
        todayOrderNumGrewPer: 0,
        totalOrderNum: 0,
        totalOrderNumGrewPer: 0,
        todayRefundNum: 0,
        todayRefundGrewPer: 0,
        totalRefundNum: 0,
        totalRefundGrewPer: 0,
        totalPropertyNum: 0,
        onlinePropertyNum: 0,
        todayOrderMoney: 0,
        todayOrderMoneyGrewPer: 0,
        totalOrderMoney: 0,
        totalOrderMoneyGrewPer: 0,
        todayRefundMoney: 0,
        todayRefundMoneyGrewPer: 0,
        totalRefundMoney: 0,
        totalRefundMoneyGrewPer: 0,
        totalRoomTypeNum: 0,
        onlineRoomTypeNum: 0,
      },
      hotelStat: {
        todayOrderNum: 0,
        todayOrderNumGrewPer: 0,
        totalOrderNum: 0,
        totalOrderNumGrewPer: 0,
        todayRefundNum: 0,
        todayRefundGrewPer: 0,
        totalRefundNum: 0,
        totalRefundGrewPer: 0,
        totalPropertyNum: 0,
        onlinePropertyNum: 0,
        todayOrderMoney: 0,
        todayOrderMoneyGrewPer: 0,
        totalOrderMoney: 0,
        totalOrderMoneyGrewPer: 0,
        todayRefundMoney: 0,
        todayRefundMoneyGrewPer: 0,
        totalRefundMoney: 0,
        totalRefundMoneyGrewPer: 0,
        totalRoomTypeNum: 0,
        onlineRoomTypeNum: 0,
      },
    },
    orderNumberList: [],
    newOrderNumberList: [],
    propertyList: [],
    oTAChannelList: [],
    nationalityList: [],
  },
});

const activeOrderStat = computed(() => {
  return state.orderStatType === 'hotel'
    ? state.dashboard.orderStat.hotelStat
    : state.dashboard.orderStat.allStat;
});

const formatTrend = (value: number) => {
  const num = Number(value || 0);
  return `${num >= 0 ? '+' : '-'}${Math.abs(num)}%`;
};

const featuredCards = computed(() => [
  {
    title: '新增会员',
    value: Number(state.dashboard.memberStat.todayRegMemberNum || 0).toLocaleString(),
    badge: formatTrend(state.dashboard.memberStat.todayMemberNumGrewPer),
    caption: '今日新增会员',
    description: `累计会员 ${Number(state.dashboard.memberStat.totalMemberNum || 0).toLocaleString()}`,
    up: Number(state.dashboard.memberStat.todayMemberNumGrewPer || 0) >= 0,
  },
  {
    title: '累计订单',
    value: Number(activeOrderStat.value.totalOrderNum || 0).toLocaleString(),
    badge: formatTrend(activeOrderStat.value.totalOrderNumGrewPer),
    caption: '全部订单累计',
    description: `今日订单 ${Number(activeOrderStat.value.todayOrderNum || 0).toLocaleString()} 单`,
    up: Number(activeOrderStat.value.totalOrderNumGrewPer || 0) >= 0,
  },
  {
    title: '订单总额',
    value: `¥${Number(activeOrderStat.value.totalOrderMoney || 0).toLocaleString()}`,
    badge: formatTrend(activeOrderStat.value.totalOrderMoneyGrewPer),
    caption: '累计订单总额',
    description: `今日订单总额 ¥${Number(activeOrderStat.value.todayOrderMoney || 0).toLocaleString()}`,
    up: Number(activeOrderStat.value.totalOrderMoneyGrewPer || 0) >= 0,
  },
  {
    title: '全部物业',
    value: Number(activeOrderStat.value.totalPropertyNum || 0).toLocaleString(),
    badge: `${Number(activeOrderStat.value.onlinePropertyNum || 0).toLocaleString()} 在线`,
    caption: '物业与房型覆盖',
    description: `全部房型 ${Number(activeOrderStat.value.totalRoomTypeNum || 0).toLocaleString()}，在线房型 ${Number(activeOrderStat.value.onlineRoomTypeNum || 0).toLocaleString()}`,
    up: true,
  },
]);

const documentRows = computed(() => [
  {
    name: '今日订单',
    category: '订单',
    status: '实时',
    statusClass: 'is-process',
    value: `${Number(activeOrderStat.value.todayOrderNum || 0).toLocaleString()} 单`,
    trend: formatTrend(activeOrderStat.value.todayOrderNumGrewPer),
    remark: '今日新增订单数量',
  },
  {
    name: '累计订单',
    category: '订单',
    status: '完成',
    statusClass: 'is-done',
    value: `${Number(activeOrderStat.value.totalOrderNum || 0).toLocaleString()} 单`,
    trend: formatTrend(activeOrderStat.value.totalOrderNumGrewPer),
    remark: '平台累计订单数量',
  },
  {
    name: '今日退款',
    category: '退款',
    status: '实时',
    statusClass: 'is-process',
    value: `${Number(activeOrderStat.value.todayRefundNum || 0).toLocaleString()} 单`,
    trend: formatTrend(activeOrderStat.value.todayRefundGrewPer),
    remark: '今日退款订单数量',
  },
  {
    name: '累计退款总额',
    category: '退款',
    status: '完成',
    statusClass: 'is-done',
    value: `¥${Number(activeOrderStat.value.totalRefundMoney || 0).toLocaleString()}`,
    trend: formatTrend(activeOrderStat.value.totalRefundMoneyGrewPer),
    remark: '平台累计退款金额',
  },
  {
    name: '新增积分',
    category: '会员',
    status: '实时',
    statusClass: 'is-process',
    value: Number(state.dashboard.memberStat.todayIncBal || 0).toLocaleString(),
    trend: formatTrend(state.dashboard.memberStat.todayIncBalGrewPer),
    remark: '会员今日新增积分',
  },
  {
    name: '消耗积分',
    category: '会员',
    status: '实时',
    statusClass: 'is-process',
    value: Number(state.dashboard.memberStat.todayConsumeBal || 0).toLocaleString(),
    trend: formatTrend(state.dashboard.memberStat.todayConsumeBalGrewPer),
    remark: '会员今日消耗积分',
  },
]);

const rankingTitle = computed(() => {
  if (activePanel.value === 'property') return '物业名称';
  if (activePanel.value === 'ota') return '渠道名称';
  return '国家名称';
});

const rankingRows = computed(() => {
  if (activePanel.value === 'property') {
    return state.dashboard.propertyList.map((item: any, index) => ({
      rank: index + 1,
      name: item?.propertyDetail?.name || '-',
      rate: item?.rate || 0,
      remark: '物业预订量占比',
    }));
  }
  if (activePanel.value === 'ota') {
    return state.dashboard.oTAChannelList.map((item: any, index) => ({
      rank: index + 1,
      name: item?.name || '-',
      rate: item?.rate || 0,
      remark: 'OTA渠道预订占比',
    }));
  }
  return state.dashboard.nationalityList.map((item: any, index) => ({
    rank: index + 1,
    name: item?.name || '-',
    rate: item?.rate || 0,
    remark: '国家预订量占比',
  }));
});

const Load = async () => {
  show.value = true;
  const res = await dashboardAll({});
  state.dashboard = res;
  await nextTick();
  const sevenDayOrders = (state.dashboard.orderNumberList || []).slice(0, 7) as any[];
  let daytimeArr = sevenDayOrders.map((item) => item?.orderDate || '');
  let percentArr = sevenDayOrders.map((item) => Number(item?.orderNumber || 0));
  setOptions({
    grid: {
      top: '9%',
      bottom: '9%',
      right: '20px',
      left: '60px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 2,
          color: '#111111',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#111111', 0.38).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#111111', 0.04).rgba,
            },
          ]),
        },
      },
    ],
  });

  // 处理30天订单趋势数据
  let daytimeArr30 = [];
  let percentArr30 = [];
  for (let i = 0; i < state.dashboard.newOrderNumberList.length; i++) {
    daytimeArr30.push(state.dashboard.newOrderNumberList[i]['orderDate']);
    percentArr30.push(state.dashboard.newOrderNumberList[i]['orderNumber']);
  }
  setOptions30({
    grid: {
      top: '9%',
      bottom: '9%',
      right: '20px',
      left: '60px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr30,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 13,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr30,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 3,
          color: '#5D87FF',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#5D87FF', 0.2).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#5D87FF', 0.01).rgba,
            },
          ]),
        },
      },
    ],
  });

  // 移动端七日订单趋势图表
  setOptionsMobile({
    grid: {
      top: '8%',
      bottom: '12%',
      right: '12px',
      left: '40px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        rotate: 0,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 2,
          color: '#5D87FF',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#5D87FF', 0.2).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#5D87FF', 0.01).rgba,
            },
          ]),
        },
      },
    ],
  });

  // 移动端30天订单趋势图表
  setOptions30Mobile({
    grid: {
      top: '8%',
      bottom: '12%',
      right: '12px',
      left: '40px',
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: daytimeArr30,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        rotate: 0,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        show: true,
        color: '#999',
        fontSize: 11,
        formatter: '{value} 单',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
        },
      },
      splitLine: {
        show: true,
        lineStyle: {
          color: '#e8e8e8',
          width: 1,
          type: 'dashed' as const,
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      valueFormatter: (value) => value + '单',
    },
    series: [
      {
        data: percentArr30,
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 2,
          color: '#5D87FF',
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: hexToRgba('#5D87FF', 0.2).rgba,
            },
            {
              offset: 1,
              color: hexToRgba('#5D87FF', 0.01).rgba,
            },
          ]),
        },
      },
    ],
  });
  show.value = false;
};

const handleUpdateValue = (value: string) => {
  state.orderStatType = value;
  console.log(value);
};

onMounted(() => {
  Load();
});
</script>

<style scoped lang="less">
.shadcn-dashboard {
  min-height: 100%;
  background: #f7f7f8;
  color: #0a0a0a;
  font-family:
    Inter,
    ui-sans-serif,
    system-ui,
    -apple-system,
    BlinkMacSystemFont,
    "Segoe UI",
    sans-serif;
}

.documents-frame {
  margin: 8px 12px 24px;
  overflow: hidden;
  border: 1px solid #e5e5e5;
  border-radius: 16px;
  background: #fff;
  box-shadow: 0 1px 2px rgb(0 0 0 / 0.05);
}

.documents-header {
  display: flex;
  height: 54px;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid #e5e5e5;
  padding: 0 24px;
  font-size: 16px;
}

.documents-title {
  display: flex;
  align-items: center;
  gap: 14px;
}

.icon-button {
  display: inline-flex;
  width: 28px;
  height: 28px;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  color: #111;

  svg {
    width: 18px;
    height: 18px;
  }
}

.header-separator {
  width: 1px;
  height: 24px;
  background: #e5e5e5;
}

.documents-content {
  padding: 24px;
}

.section-card-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
}

.featured-card,
.chart-card,
.data-shell {
  border: 1px solid #e5e5e5;
  border-radius: 14px;
  background: #fff;
  box-shadow: 0 2px 4px rgb(0 0 0 / 0.04);
}

.featured-card {
  min-height: 170px;
  padding: 24px;
  background: linear-gradient(180deg, #fff 0%, #f7f7f7 100%);
}

.featured-card-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  color: #737373;
  font-size: 14px;
}

.trend-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border: 1px solid #e5e5e5;
  border-radius: 999px;
  padding: 3px 9px;
  color: #111;
  font-size: 13px;
  font-weight: 700;
  white-space: nowrap;

  svg {
    width: 13px;
    height: 13px;
  }
}

.featured-value {
  margin-top: 14px;
  font-size: 32px;
  font-weight: 800;
  line-height: 1.1;
  letter-spacing: 0;
}

.featured-caption {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 32px;
  color: #111;
  font-size: 14px;
  font-weight: 700;

  svg {
    width: 15px;
    height: 15px;
  }
}

.featured-card p {
  margin-top: 8px;
  color: #737373;
  font-size: 14px;
}

.chart-card {
  margin-top: 24px;
  padding: 24px;
}

.chart-card-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 24px;

  h2 {
    margin: 0;
    color: #111;
    font-size: 16px;
    font-weight: 800;
  }

  p {
    margin-top: 6px;
    color: #737373;
    font-size: 14px;
  }
}

.range-tabs {
  display: inline-flex;
  overflow: hidden;
  border: 1px solid #e5e5e5;
  border-radius: 10px;

  button {
    height: 36px;
    border-left: 1px solid #e5e5e5;
    padding: 0 16px;
    color: #111;
    font-size: 14px;
    font-weight: 700;

    &:first-child {
      border-left: 0;
    }

    &.active {
      background: #f4f4f5;
    }
  }
}

.featured-chart {
  width: 100%;
  height: 320px;
  margin-top: 20px;
}

.table-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
  margin-top: 24px;
}

.outline-tabs {
  display: inline-flex;
  align-items: center;
  gap: 2px;
  border-radius: 12px;
  background: #f4f4f5;
  padding: 4px;

  button {
    display: inline-flex;
    height: 34px;
    align-items: center;
    gap: 8px;
    border-radius: 9px;
    padding: 0 12px;
    color: #737373;
    font-size: 14px;
    font-weight: 700;

    &.active {
      background: #fff;
      color: #111;
      box-shadow: 0 1px 2px rgb(0 0 0 / 0.08);
    }

    span {
      display: inline-flex;
      min-width: 24px;
      height: 22px;
      align-items: center;
      justify-content: center;
      border-radius: 999px;
      background: #d4d4d8;
      color: #111;
      font-size: 12px;
    }
  }
}

.toolbar-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.toolbar-button {
  display: inline-flex;
  height: 36px;
  align-items: center;
  gap: 10px;
  border: 1px solid #e5e5e5;
  border-radius: 10px;
  background: #fff;
  padding: 0 14px;
  color: #111;
  font-size: 14px;
  font-weight: 700;
  box-shadow: 0 1px 2px rgb(0 0 0 / 0.04);

  svg {
    width: 16px;
    height: 16px;
  }
}

.data-shell {
  margin-top: 20px;
  overflow: hidden;
}

.documents-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;

  th,
  td {
    height: 54px;
    border-bottom: 1px solid #e5e5e5;
    padding: 0 14px;
    text-align: left;
    white-space: nowrap;
  }

  th {
    background: #f7f7f7;
    color: #111;
    font-weight: 800;
  }

  tbody tr:last-child td {
    border-bottom: 0;
  }
}

.drag-cell {
  width: 44px;
  color: #737373;
  text-align: center !important;
}

.check-cell {
  width: 44px;
}

.menu-cell {
  width: 44px;
  color: #737373;
  text-align: center !important;
}

.fake-check {
  display: inline-flex;
  width: 18px;
  height: 18px;
  border: 1px solid #e5e5e5;
  border-radius: 5px;
  background: #fff;
  box-shadow: inset 0 1px 1px rgb(0 0 0 / 0.03);
}

.row-title {
  color: #111;
  font-weight: 800;
}

.soft-pill {
  display: inline-flex;
  align-items: center;
  border: 1px solid #e5e5e5;
  border-radius: 999px;
  padding: 4px 10px;
  color: #737373;
  font-size: 13px;
  font-weight: 600;
}

.soft-pill.is-done {
  color: #16a34a;
}

.soft-pill.is-process {
  color: #737373;
}

.soft-pill.is-muted {
  color: #a1a1aa;
}

.ranking-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.ranking-card {
  overflow: hidden;
  border: 1px solid #e5e5e5;
  border-radius: 18px;
  background: #fff;
  box-shadow: 0 2px 4px rgb(0 0 0 / 0.04);

  h3 {
    margin: 0;
    padding: 22px 28px;
    border-bottom: 1px solid #e5e5e5;
    color: #111;
    font-size: 20px;
    font-weight: 800;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    font-size: 16px;
  }

  th,
  td {
    height: 48px;
    border-bottom: 1px solid #eeeeee;
    padding: 0 18px;
    text-align: left;
    white-space: nowrap;
  }

  th {
    background: #f7f7f7;
    color: #111;
    font-weight: 800;
  }

  td {
    color: #1d2129;
  }

  tbody tr:last-child td {
    border-bottom: 0;
  }
}

@media (max-width: 1400px) {
  .section-card-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .ranking-grid {
    grid-template-columns: 1fr;
  }
}

.title-name {
  margin: 15px 0 19px 20px;
  font-weight: 500;
  font-size: 30px;
  color: #3d3d3d;
  line-height: 42px;
}

.room-stat-title {
  font-size: 24px;
  color: #3d3d3d;
  font-weight: 500;
  line-height: 34px;
}

.room-stat-title1 {
  font-weight: 500;
  font-size: 18px;
  color: #3d3d3d;
  line-height: 25px;
}

.room-stat-div {
  display: flex;
  align-items: center;
  width: 100%;
  margin-top: 5px;
}

.room-stat-time {
  font-size: 14px;
  color: #8b8b8b;
  font-weight: 400;
  line-height: 20px;
}

.room-stat-refresh {
  font-size: 14px;
  color: #156bff;
  line-height: 20px;
  margin-left: 5px;
}

.room-stat-gi-div {
  border-radius: 2px;
  display: flex;
  align-items: center;
  background: #f9fbff;
  padding: 20px 30px;

  .room-stat-gi-div-item {
    flex: 1;
    margin-left: 30px;

    .room-stat-gi-div-item-d1 {
      font-size: 14px;
      color: #3d3d3d;
      font-weight: 500;
      line-height: 20px;
    }

    .room-stat-gi-div-item-d2 {
      margin-top: 5px;
      font-size: 28px;
      color: #3d3d3d;
      font-weight: 700;
      line-height: 42px;
    }

    .room-stat-gi-div-item-d3 {
      display: flex;
      align-items: center;
      margin-top: 10px;
      font-size: 14px;
      color: #8b8b8b;
      font-weight: 400;
      line-height: 20px;

      img {
        display: block;
      }
    }
  }
}

.tab-div2 {
  padding: 12px 0;

  .tab-div2-d1 {
    font-weight: 500;
    font-size: 14px;
    color: #3d3d3d;
    line-height: 20px;

    span {
      color: #979797;
      font-weight: 400;
    }
  }

  .tab-div2-d2 {
    margin-top: 5px;
    font-weight: 700;
    font-size: 24px;
    color: #3d3d3d;
    line-height: 42px;
  }

  .tab-div2-d3 {
    display: flex;
    align-items: center;
    margin-top: 5px;
    font-weight: 400;
    font-size: 14px;
    color: #8b8b8b;
    line-height: 20px;

    img {
      display: block;
    }
  }
}

.data-table {
  .data-table-thead {
    background: #f2f3f8;
    display: flex;

    div {
      font-weight: 500;
      font-size: 14px;
      color: #1d2129;
      line-height: 40px;
    }

    .data-table-td-d1 {
      width: 17.3%;
      text-align: center;
    }

    .data-table-td-d2 {
      width: 55%;
      padding-left: 7.3%;
    }

    .data-table-td-d3 {
      width: 27.7%;
      text-align: center;
    }
  }

  .data-table-tbody {
    display: flex;
    border-bottom: 1px solid #f2f3f8;

    div {
      font-weight: 400;
      font-size: 14px;
      color: #1d2129;
      line-height: 44px;
    }

    .data-table-td-d1 {
      width: 17.3%;
      text-align: center;
    }

    .data-table-td-d2 {
      width: 55%;
      padding-left: 7.3%;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
    }

    .data-table-td-d3 {
      width: 27.7%;
      text-align: center;
    }
  }
}

// 移动端样式
.room-stat-gi-div-mobile {
  border-radius: 8px;
  background: #f9fbff;
  padding: 12px;
  min-height: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  .room-stat-gi-div-item-mobile {
    width: 100%;
    margin-top: 12px;
  }
}

.tab-div2-mobile {
  padding: 12px;
  background: #f9fbff;
  border-radius: 8px;
  min-height: 100px;
}

.data-table-mobile {
  .data-table-tbody-mobile {
    &:last-child {
      .border-b {
        border-bottom: none !important;
      }
    }
  }
}
</style>
