<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
        padding: '20px',
      }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">预约日历</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>

    <n-card :bordered="false" class="proCard" :content-style="{
      padding: '0 20px 20px',
    }">
      <div class="statusTab">
        <div class="tab__item" :class="tabValue == TabTypeVo.DAY ? 'active' : ''"><span
            @click="handleUpdateValue(TabTypeVo.DAY)">日视图</span></div>
        <div class="tab__item" :class="tabValue == TabTypeVo.WEEK ? 'active' : ''"><span
            @click="handleUpdateValue(TabTypeVo.WEEK)">周视图</span></div>
        <div class="tab__item" :class="tabValue == TabTypeVo.MONTH ? 'active' : ''"><span
            @click="handleUpdateValue(TabTypeVo.MONTH)">月视图</span></div>
      </div>

      <div class="flex items-center justify-between mb-4">
        <div class="">
          <n-date-picker v-model:value="selectedDay" type="date" />
        </div>
        <div class="flex items-center gap-5">
          <div v-for="(item, index) in orderStatusList" :key="index" class="flex items-center gap-1">
            <div class="w-3 h-3 rounded-sm" :style="{ backgroundColor: item.color }"></div>
            <div class="text-sm">{{ item.label }}</div>
          </div>
        </div>
      </div>

      <!-- 日视图 -->
      <div v-if="tabValue === TabTypeVo.DAY" class="w-full border border-[#E7E7E7]">
        <div class="w-full h-14 flex items-center justify-start pl-5 border-b border-[#E7E7E7]">{{ selectDayText }}
        </div>
        <div class="w-full py-6 flex flex-col gap-4 h-[calc(100vh-388px)] overflow-y-auto hs">
          <div v-for="item in halfHourTimes" :key="item.time"
            class="flex-shrink-0 w-full h-14 flex items-center justify-start overflow-hidden">
            <div class="w-32 h-full flex-shrink-0 flex items-center justify-start pl-5">{{ item.title }}</div>
            <div
              class="w-full flex-1 h-full flex items-center justify-start gap-5 flex-nowrap overflow-x-auto hs relative after:content-[''] after:absolute after:left-0 after:top-1/2 after:w-full after:h-[1px] after:bg-[#E7E7E7]">
              <div @click="handleView(order)" v-for="order in item.orderList" :key="order.id"
                class="h-full px-3 py-2.5 pl-2 z-10 relative text-xs flex-shrink-0 cursor-pointer hover:opacity-80 transition-all duration-300"
                :style="{ borderLeft: '5px solid transparent', borderLeftColor: caculateBorderColor(order.orderStatus)[0], backgroundColor: caculateBorderColor(order.orderStatus)[1] }">
                <div>{{ caculateServiceType(order.serviceType) }}：{{ order.startServiceAddress?.name || '--' }} {{
                  order.bookTime }}</div>
                <div class="mt-1">
                  <span>司机：</span>
                  <span v-if="order.orderType == 'INNN'">--</span>
                  <span class="text-[#3264FF] font-medium" v-else-if="order.dispatchStatus == 'WAIT'">待指派</span>
                  <span v-else>{{ order.driverDetail?.nickname }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 周视图 -->
      <div v-if="tabValue === TabTypeVo.WEEK" class="w-full border border-[#E7E7E7]">
        <div class="w-full h-14 flex text-base">
          <div class="h-full flex-1 flex-shrink-0 flex items-center justify-center border-b text-base font-medium">
            {{ halfWeekTimes[0].monthText }}
          </div>
          <div class="h-full flex-1 flex-shrink-0 flex items-center justify-center border-l border-b border-[#E7E7E7]"
            v-for="(item, index) in halfWeekTimes" :key="item.day">
            <span>{{ item.dayNum }}</span>
            <span :class="['ml-1', index > 4 && 'text-[#E13E3E]']">{{
              item.weekText }}</span>
          </div>
        </div>
        <div class="w-full flex flex-col h-[calc(100vh-388px)] overflow-y-auto hs">
          <div class="w-full flex-shrink-0 border-b border-[#E7E7E7] flex"
            :class="[index === halfHourTimes.length - 1 && 'border-b-0']" v-for="(item, index) in halfHourTimes"
            :key="item.time">
            <div
              class="w-[12.5%] flex-shrink-0 whitespace-nowrap flex items-center justify-center p-2.5 text-base font-medium">
              {{
                item.title }}</div>
            <div
              class="w-[12.5%] flex-shrink-0 p-2.5 border-l border-[#E7E7E7] flex flex-col gap-1.5 items-start justify-start"
              v-for="(week, nextIndex) in halfWeekTimes" :key="week.day">
              <div class="w-full flex" v-for="(order, orderIndex) in week.times[index].orderList" :key="orderIndex">
                <n-popover trigger="hover">
                  <template #trigger>
                    <div @click="handleView(order)"
                      class="w-auto max-w-full hd px-3 py-1.5 pl-2 text-xs flex-shrink-0 cursor-pointer hover:opacity-80 transition-all duration-300"
                      :style="{ borderLeft: '5px solid transparent', borderLeftColor: caculateBorderColor(order.orderStatus)[0], backgroundColor: caculateBorderColor(order.orderStatus)[1] }">
                      <span class="mr-1">{{ caculateServiceType(order.serviceType) }}：{{ order.startServiceAddress?.name
                        ||
                        '--'
                      }}</span>
                      <span v-if="order.orderType == 'INNN'">--</span>
                      <span class="text-[#3264FF] font-medium" v-else-if="order.dispatchStatus == 'WAIT'">待指派</span>
                      <span v-else>{{ order.driverDetail?.nickname }}</span>
                    </div>
                  </template>
                  <div>
                    <span class="mr-1">{{ caculateServiceType(order.serviceType) }}：{{ order.startServiceAddress?.name
                      ||
                      '--'
                      }}</span>
                    <span v-if="order.orderType == 'INNN'">--</span>
                    <span class="text-[#3264FF]" v-else-if="order.dispatchStatus == 'WAIT'">待指派</span>
                    <span v-else>{{ order.driverDetail?.nickname }}</span>
                  </div>
                </n-popover>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 月视图 -->
      <div v-if="tabValue === TabTypeVo.MONTH" class="w-full border border-[#E7E7E7]">
        <div class="w-full h-14 flex text-base">
          <div class="h-full flex-1 flex-shrink-0 flex items-center justify-center border-l border-b border-[#E7E7E7]"
            v-for="(item, index) in ['周一', '周二', '周三', '周四', '周五', '周六', '周日']" :key="item">
            <span :class="[index > 4 && 'text-[#E13E3E]', !index && 'border-l-0']">{{ item }}</span>
          </div>
        </div>
        <div class="w-full flex flex-col h-[calc(100vh-388px)] overflow-y-auto hs">
          <div class="w-full flex-1 h-full overflow-hidden flex-shrink-0 border-b border-[#E7E7E7] flex"
            :class="[index === monthTimes.length - 1 && 'border-b-0']" v-for="(item, index) in monthTimes"
            :key="item.time">
            <div class="w-[14.28%] h-full flex-shrink-0 p-3 border-r border-[#E7E7E7] flex flex-col gap-1.5"
              :class="[dayIndex === item.length - 1 && 'border-r-0']" v-for="(day, dayIndex) in item" :key="dayIndex">
              <template v-if="day.week">
                <div class="w-full flex-shrink-0 flex items-center justify-end ">
                  <n-popover trigger="hover">
                    <template #trigger>
                      <div class="cursor-pointer hover:opacity-80 transition-all duration-300">
                        <span class="font-medium">{{ day.dayNum }}</span>
                        <span class="text-[#ccc]" v-if="day.orderList?.length">【{{
                          day.orderList?.length }}】</span>
                      </div>
                    </template>
                    <div class="w-[240px] flex flex-col gap-2">
                      <div class="w-full flex items-start justify-start" v-for="(item, index) in orderStatusList"
                        :key="index">
                        <div class="flex-1 flex-shrink-0">{{ item.label }}</div>
                        <div class="w-16 text-right">{{ getOrderStatusCount(day.orderList, item.value) }}</div>
                      </div>
                    </div>
                  </n-popover>
                </div>
                <div class="w-full h-full flex-1 overflow-y-auto hs flex gap-1 flex-col">
                  <div class="w-full flex" v-for="(order, orderIndex) in day.orderList" :key="orderIndex">
                    <n-popover trigger="hover">
                      <template #trigger>
                        <div @click="handleView(order)"
                          class="w-full hd px-3 py-1 pl-2 text-xs flex-shrink-0 cursor-pointer hover:opacity-80 transition-all duration-300"
                          :style="{ borderLeft: '5px solid transparent', borderLeftColor: caculateBorderColor(order.orderStatus)[0], backgroundColor: caculateBorderColor(order.orderStatus)[1] }">
                          <span class="mr-1">{{ caculateServiceType(order.serviceType) }}：{{
                            order.startServiceAddress?.name
                            ||
                            '--'
                          }}</span>
                          <span v-if="order.orderType == 'INNN'">--</span>
                          <span class="text-[#3264FF] font-medium" v-else-if="order.dispatchStatus == 'WAIT'">待指派</span>
                          <span v-else>{{ order.driverDetail?.nickname }}</span>
                        </div>
                      </template>
                      <div>
                        <span class="mr-1">{{ caculateServiceType(order.serviceType) }}：{{
                          order.startServiceAddress?.name
                          ||
                          '--'
                        }}</span>
                        <span v-if="order.orderType == 'INNN'">--</span>
                        <span class="text-[#3264FF]" v-else-if="order.dispatchStatus == 'WAIT'">待指派</span>
                        <span v-else>{{ order.driverDetail?.nickname }}</span>
                      </div>
                    </n-popover>
                  </div>
                </div>
              </template>
            </div>
          </div>
        </div>
      </div>
    </n-card>

    <!-- 详情弹窗 -->
    <DetailView ref="viewRef" @reloadTable="getOrderList" />
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref, watchEffect } from 'vue';
import { List } from '@/api/carOrder';
import { formatTime } from '@/utils/date';
import DetailView from './view.vue';

enum TabTypeVo {
  DAY = 'DAY', // 日视图
  WEEK = 'WEEK', // 周视图
  MONTH = 'MONTH', // 月视图
}

// 订单状态枚举
enum OrderStatusVo {
  WAIT_PAY = 'WAIT_PAY', // 待支付
  WAIT_CONFIRM = 'WAIT_CONFIRM', // 待确认
  WAIT_SERVE = 'WAIT_SERVE', // 待服务
  SERVING = 'SERVING', // 服务中
  DONE = 'DONE', // 已完成
  CANCEL = 'CANCEL', // 已取消
  OVERDUE = 'OVERDUE', // 已逾期
}

// 预约类型
enum ServiceTypeVo {
  PICKUP = 'PICKUP', // 接机
  DELIVERY = 'DELIVERY', // 送机
  CAR = 'CAR', // 包车
}

const generateHalfHourTimes = () => {
  const times: any[] = [];
  for (let hour = 0; hour < 24; hour++) {
    for (let minute of [0, 30]) {
      let period = "";
      if (hour < 12) {
        period = "上午";
      } else if (hour < 18) {
        period = "下午";
      } else {
        period = "晚上";
      }

      let displayHour = hour < 10 ? `0${hour}` : hour;
      times.push(
        minute === 0
          ? { title: `${period}${displayHour}时`, time: `${displayHour}:00`, orderList: [] }
          : { title: `${period}${displayHour}时${minute}分`, time: `${displayHour}:${minute}`, orderList: [] }
      );
    }
  }
  return times
}

const formatDateTime = (date: Date, time: string) => {
  const y = date.getFullYear();
  const m = String(date.getMonth() + 1).padStart(2, "0");
  const d = String(date.getDate()).padStart(2, "0");
  return `${y}-${m}-${d} ${time}`;
}

const getWeekDays = (dateStr: Date | string | number) => {
  const inputDate = new Date(dateStr);

  // 获取当天是星期几 (0=周日, 1=周一,...6=周六)
  let dayOfWeek = inputDate.getDay();
  if (dayOfWeek === 0) dayOfWeek = 7; // 把周日改为 7

  // 找到周一 (当前日期 - (dayOfWeek - 1)天)
  const monday = new Date(inputDate);
  monday.setDate(inputDate.getDate() - (dayOfWeek - 1));

  const result: any[] = [];

  for (let i = 0; i < 7; i++) {
    const d = new Date(monday);
    d.setDate(monday.getDate() + i);

    let timePart = "00:00:00";
    if (i === 6) timePart = "23:59:59"; // 周日结尾

    result.push({
      day: formatDateTime(d, timePart),
      weekText: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'][i],
      weekNum: i + 1,
      dayNum: d.getDate(),
      monthText: d.getMonth() + 1 + '月',
      times: generateHalfHourTimes()
    });
  }

  return result;
}

const getMonthDays = (dateStr: Date | string | number) => {
  const inputDate = new Date(dateStr);
  const year = inputDate.getFullYear();
  const month = inputDate.getMonth() + 1; // JS 月份从0开始

  // 当月第一天和最后一天
  const firstDay = new Date(year, month - 1, 1);
  const lastDay = new Date(year, month, 0); // 当月最后一天

  const result: any[] = [];

  for (let d = 1; d <= lastDay.getDate(); d++) {
    const current = new Date(year, month - 1, d);

    // 获取周几 (0=周日,1=周一,...6=周六)
    let dayOfWeek = current.getDay();
    if (dayOfWeek === 0) dayOfWeek = 7; // 周日=7

    let timePart = "00:00:00";
    if (d === lastDay.getDate()) timePart = "23:59:59"; // 当月最后一天结尾

    result.push({
      day: formatDateTime(current, timePart),
      week: dayOfWeek,
      dayNum: d,
      month: month,
      orderList: []
    });
  }

  return result;
}

const getMonthDaysGrid = (days: any[]) => {
  const grid: any = [];
  let week: any[] = [];

  // 补前置空对象
  const firstDayOfWeek = days[0].week;
  for (let i = 1; i < firstDayOfWeek; i++) {
    week.push({});
  }

  // 遍历当月天数，按周分组
  days.forEach(day => {
    week.push(day);
    if (week.length === 7) {
      grid.push(week);
      week = [];
    }
  });

  // 补后置空对象
  if (week.length > 0) {
    while (week.length < 7) {
      week.push({});
    }
    grid.push(week);
  }

  return grid;
}

const viewRef = ref();
const tabValue = ref(TabTypeVo.DAY);
const selectedDay = ref<any>(new Date());

// 将selectedDay格式化为 7月22日全天
const selectDayText = computed(() => {
  const date = new Date(selectedDay.value);
  return `${date.getMonth() + 1}月${date.getDate()}日全天`;
});

// 日视图数据
const halfHourTimes = ref(generateHalfHourTimes());
// 周视图数据
const halfWeekTimes = ref(getWeekDays(selectedDay.value));
// 月视图数据
const monthTimes = ref(getMonthDaysGrid(getMonthDays(selectedDay.value)));

const orderStatusList = ref([
  {
    label: '待支付',
    value: [OrderStatusVo.WAIT_PAY],
    color: '#F73314',
    bgColor: '#FFECEA'
  },
  {
    label: '待确认',
    value: [OrderStatusVo.WAIT_CONFIRM],
    color: '#FF6F00',
    bgColor: '#FFEBDC'
  },
  {
    label: '待服务',
    value: [OrderStatusVo.WAIT_SERVE],
    color: '#19A158',
    bgColor: '#E3F4EB'
  },
  {
    label: '服务中',
    value: [OrderStatusVo.SERVING],
    color: '#227FF0',
    bgColor: '#E3F0FD'
  },
  {
    label: '已完成/已取消/已逾期',
    value: [OrderStatusVo.DONE, OrderStatusVo.CANCEL, OrderStatusVo.OVERDUE],
    color: '#343639',
    bgColor: '#EEEEEE'
  },
]);

const caculateBorderColor = (orderStatus: OrderStatusVo) => {
  const orderStatusItem = orderStatusList.value.find(item => item.value.includes(orderStatus));
  return [orderStatusItem?.color || '#E7E7E7', orderStatusItem?.bgColor || '#E7E7E7'];
}

const caculateServiceType = (serviceType: ServiceTypeVo) => {
  if (serviceType == ServiceTypeVo.PICKUP) {
    return '接机'
  } else if (serviceType == ServiceTypeVo.DELIVERY) {
    return '送机'
  } else {
    return '包车'
  }
}

const getOrderStatusCount = (orderList: any[], statusList: OrderStatusVo[]) => {
  return orderList.filter(order => statusList.includes(order.orderStatus)).length;
}

// 切换视图
const handleUpdateValue = (value: TabTypeVo) => {
  tabValue.value = value;
};

// 获取日视图订单列表
const getOrderList = async () => {
  try {
    const startTime = new Date(selectedDay.value);
    startTime.setHours(0, 0, 0, 0);
    const endTime = new Date(startTime.getTime() + 24 * 60 * 60 * 1000);
    const bookStartTime = [formatTime(startTime.getTime()), formatTime(endTime.getTime())];
    const res = await List({
      page: 1,
      pageSize: 500,
      bookStartTime
    });
    halfHourTimes.value.forEach(item => {
      item.orderList = res.list.filter(order => {
        // 解析订单的预约开始时间
        const orderTime = new Date(order.bookStartTime);
        const orderHour = orderTime.getHours();
        const orderMinute = orderTime.getMinutes();

        // 解析当前时间段
        const [itemHour, itemMinute] = item.time.split(':').map(Number);

        // 计算时间差（分钟）
        const timeDiff = Math.abs((orderHour * 60 + orderMinute) - (itemHour * 60 + itemMinute));

        // 如果时间差在15分钟以内，则包含此订单
        return timeDiff <= 15;
      });
    });
  } catch (error) {
    console.log(error);
  }
};

// 获取周视图订单列表
const getWeekOrderList = async () => {
  try {
    halfWeekTimes.value = getWeekDays(selectedDay.value);
    const bookStartTime = [halfWeekTimes.value[0].day, halfWeekTimes.value[6].day];
    const res = await List({
      page: 1,
      pageSize: 500,
      bookStartTime
    });
    halfWeekTimes.value.forEach(week => {
      week.times.forEach(item => {
        item.orderList = res.list.filter(order => {
          // 解析订单的预约开始时间
          const orderTime = new Date(order.bookStartTime);
          const orderHour = orderTime.getHours();
          const orderMinute = orderTime.getMinutes();

          // 解析当前时间段
          const [itemHour, itemMinute] = item.time.split(':').map(Number);

          // 计算时间差（分钟）
          const timeDiff = Math.abs((orderHour * 60 + orderMinute) - (itemHour * 60 + itemMinute));

          // 如果时间差在15分钟以内，则包含此订单
          return timeDiff <= 15 && week.dayNum === orderTime.getDate();
        });
      });
    })
  } catch (error) {
    console.log(error);
  }
};

// 获取月视图订单列表
const getMonthOrderList = async () => {
  try {
    const list = getMonthDays(selectedDay.value);
    const bookStartTime = [list[0].day, list[list.length - 1].day];
    const res = await List({
      page: 1,
      pageSize: 1000,
      pagination: false,
      bookStartTime
    });
    list.forEach(item => {
      item.orderList = res.list.filter(order => {
        const orderTime = new Date(order.bookStartTime);
        return item.dayNum === orderTime.getDate();
      });
    });
    monthTimes.value = getMonthDaysGrid(list);
  } catch (error) {
    console.log(error);
  }
};

// 监听视图切换
watchEffect(() => {
  if (tabValue.value === TabTypeVo.DAY && selectedDay.value) {
    getOrderList();
  }
  if (tabValue.value === TabTypeVo.WEEK && selectedDay.value) {
    getWeekOrderList();
  }
  if (tabValue.value === TabTypeVo.MONTH && selectedDay.value) {
    getMonthOrderList();
  }
});

// 查看详情
function handleView(record: Recordable) {
  viewRef.value.openModal(record);
}
</script>

<style lang="less" scoped>
.hd {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

// 隐藏滚动条通用样式
.hs {
  &::-webkit-scrollbar {
    display: none;
  }

  &::-webkit-scrollbar-track-piece {
    display: none;
  }

  &::-webkit-scrollbar-thumb {
    display: none;
  }

  &::-webkit-scrollbar-button {
    display: none;
  }

  &::-webkit-scrollbar-corner {
    display: none;
  }

  &::-webkit-scrollbar-track {
    display: none;
  }
}


.statusTab {
  display: flex;
  margin-bottom: 16px;

  .tab__item {
    padding: 0 16px;
    height: 28px;
    line-height: 28px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;

    span {
      cursor: pointer;
    }

    &.active {
      background: #F2F3F8;
      border-radius: 28px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
