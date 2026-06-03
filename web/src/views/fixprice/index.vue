<template>
  <!--选择物业-->
  <div class="title-name">
    <!--      <div>{{ namegetlang(userStore.getuserPms.nameLanguage) }}</div>-->
  
    <div class="dropdown-container" style="width: 500px">
      <!-- 自定义触发元素 -->
      <n-popover trigger="click" placement="bottom-start" :show-arrow="false" :overlap="false" raw>
        <template #trigger>
          <div class="trigger-div" @click="handleClick">
            {{ propertyName }}
            <!--              {{ namegetlang(userStore.getuserPms.nameLanguage) || '全部物业' }}-->
            <n-icon :component="isOpen ? ChevronUp : ChevronDown" />
          </div>
        </template>
  
        <!-- 下拉选择器 -->
        <n-select style="width: 500px" v-model:value="selectedValue" filterable :filter="filterHandler"
          :options="sortedOptions" :render-label="renderLabel" value-field="id"
          :menu-props="{ style: { width: '500px' } }" @update:value="handleSelect" />
      </n-popover>
    </div>
  
  </div>
  <!-- 操作加载中 -->
  <div v-if="state.loading" style="height: 80vh">
    <div class="first-loading-wrap" style="padding-top: 200px">
      <div class="loading-wrap">
        <span class="dot dot-spin"><i></i><i></i><i></i><i></i></span>
        <span class="fw f22 ml-3">Loading...</span>
      </div>
    </div>
  </div>
  <!-- 操作加载中 -->
  <div v-else class="bgfff" style="width: 100%; padding: 10px; border-radius: 4px; overflow-y: hidden">
    <div class="mt-2">
      <div class="flex-row f12 p-2" style="position: relative">
        <div class="c333 fw f18" style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{
          translang('定价管理') }}</div>
        <div class="flex-item text-r">
          <a-date-picker v-model:value="state.selectdate" :placeholder="translang('初始为当天')"
            style="width: 220px;border-radius: 4px" @change="selectdateChange"></a-date-picker>
          <!-- <n-button type="primary" class="mr-1" @click="state.yuyueopen = true">预约订房</n-button>
                                                              <n-button type="warning" class="mr-1">关房</n-button> -->
        </div>
      </div>
    </div>
  
    <div class="flex-column bgfff brt1 brl1" style="height: 78vh; position: relative">
      <!-- 向左向右 操作 -->
      <!-- 向左 -->
      <div class="goleft hover:bg-white transition-opacity duration-300" @click="maindivgo('left')">
        <n-icon color="#333" size="18">
          <ChevronBackOutline />
        </n-icon>
      </div>
      <!-- 向右 -->
      <div class="goright hover:bg-white transition-opacity duration-300" @click="maindivgo('right')">
        <n-icon color="#333" size="18">
          <ChevronForwardOutline />
        </n-icon>
      </div>
      <!-- 日期 -->
      <div class="flex-row">
        <div class="dyg rq" style="width: 240px"> </div>
        <div id="rqdiv" class="flex-item" style="
                                                                width: 100%;
                                                                flex-grow: 1;
                                                                display: flex;
                                                                z-index: 1;
                                                                line-height: 25px;
                                                                height: 86px;
                                                                overflow-x: auto;
                                                                overflow-y: hidden;
                                                              ">
          <div v-for="(item, index) in datelist" :key="index" class="dyg rq text-c">
            <div :class="item.today ? 'todaybg' : ''">
              <div :class="'周日，周六'.indexOf(item.week) != -1 ? 'cred' : 'c999'" style="font-size: 12px">
                {{ item.week }}
              </div>
              <div class="c333 fw" style="font-size: 16px">
                <div :class="item.today ? 'today' : ''">{{ item.day }}</div>
              </div>
              <div class="text-c c999" style="font-size: 11px">
                <span v-if="item.today" style="color: #053dc8ff">今日</span>
                <span v-else>{{ item.year }}年{{ item.mm }}月</span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="flex-item" style="overflow-y: hidden">
        <n-infinite-scroll :distance="20" style="height: 69.6vh; overflow-y: auto" @load="DataLoad">
          <div class="flex-row">
            <div style="width: 240px">
              <div v-for="(item, index) in state.RoomTypeList" :key="index">
                <div class="dyg999" style="line-height: 23px; height: 50px">
                  <span class="ml-2 fw"> {{ item.name }}</span>
                  <div class="ml-8">
                    <span>
                      <n-icon color="#999" size="15">
                        <BusinessOutline />
                      </n-icon>
                    </span>
                    <span class="ml-1 f12 c999">{{ item.property.name }}</span>
                  </div>
                </div>
                <div class="dyg999" style="line-height: 50px; height: 50px">
                  <span class="ml-8 f16">Standard Rate</span>
                </div>
              </div>
            </div>
            <div id="maindiv" class="flex-item" style="overflow-x: auto; overflow-y: hidden;" @scroll="optimizedScroll">
              <div v-for="(item, index) in state.RoomTypeList" :key="index">
                <div class="righthx">
                  <div class="kucun">
                    <div v-for="(item2, index2) in state.Availabilities[item.uid]" class="dyg999 c999"
                      style="width: 70px; height: 50px; padding: 1px 0; line-height: 20px">
                      <div v-if="item2[0].allotment" style="color: #333; font-size: 14px">
                        {{ item2[0].allotment }}
                        <div class="kyg">{{ translang('可预订') }}</div>
                      </div>
                      <div v-else style="color: #333; font-size: 14px">
                        0
                        <div class="manfang">{{ translang('无库存') }}</div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="righthx">
                  <div class="kucun">
                    <div v-for="(item2, index2) in state.Availabilities[item.uid]" class="dyg999 c999" style="
                                                                            width: 70px;
                                                                            height: 50px;
                                                                            padding: 5px;
                                                                            line-height: 20px;
                                                                            text-align: left;
                                                                          ">
                      <div class="f12">{{ item2[0].currency }}</div>
                      <div class="f14" style="color: #007d70">{{ item2[0].price }} </div>
                      <div class="f14" style="color: #999999;text-decoration: line-through">{{item2[0].basePrice }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </n-infinite-scroll>
      </div>
    </div>
  
    <a-drawer v-model:open="state.yuyueopen" :title="translang('预约')" class="custom-class" placement="right"
      root-class-name="root-class-name" width="580">
      <Appointment :selectdays="state.selectdays" />
    </a-drawer>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
import { onMounted, reactive, ref, watch, h, computed } from 'vue';
import {
  BusinessOutline,
  ChevronBackOutline, ChevronDown,
  ChevronForwardOutline,
  ChevronUp
} from '@vicons/ionicons5';
import { priceCalendar } from '@/api/comm';
import { List as pmsRoomTypeList } from '@/api/pmsRoomType';

import { formatToDateTime, formatWeek } from '@/utils/dateUtil';
import { options } from '../curdDemo/model';
import { getlang, translang } from '@/utils/smjcomm';

import { Dicts } from '@/api/dict/dict';
import { useUserStore } from '@/store/modules/user';
import Appointment from './comm/appointment.vue';
import View from '@/views/pmsAppReservation/view.vue';
import { useMessage } from 'naive-ui';
import { storage } from "@/utils/Storage";

const viewRef = ref();

const message = useMessage();
const userStore = useUserStore();

const state = reactive({
  page: 1,
  limit: 100,
  mainkey: 1,
  noMore: false,
  yuyueopen: false,
  roomindex: null,
  fxlabel: '所有房型',
  roomTypeUid: '',
  roomTypeoptions: [],
  loading: false,
  Reservation: [],
  RoomTypeList: [],
  options: [],
  reservationobj: {}, // 订单房间数据详情
  selectdays: [], //选的 天数
  sx_fx: '',
  selectdate: '',
  Availabilities: {},
});
const namegetlang = (data) => {
  if (data) {
    return getlang(data, userStore.language).content;
  } else {
    return '所有物业';
  }
};

// 物业名称
const propertyName = ref(userStore.getuserPms.id > 0 ? namegetlang(userStore.getuserPms.nameLanguage) : '全部物业');

// 物业列表
const wylists = ref(storage.get('wylists'));

// 控制箭头
const isOpen = ref(false)

// 下拉选中的值
const selectedValue = ref(userStore.getuserPms.id)

// 排序
const sortedOptions = computed(() => {

  // console.log("wylists.value222", wylists.value)

  const selected = wylists.value.find(opt => opt.id === selectedValue.value)
  const others = wylists.value.filter(opt => opt.id !== selectedValue.value)
  if (selectedValue.value > 0) {
    return selected ? [wylists.value.find(opt => opt.id === 0), selected, ...others.filter(opt => opt.id !== 0)] : wylists.value
  }
  return wylists.value
})

// 自定义选项渲染
const renderLabel = (option) => {
  // console.log('option', namegetlang(option.nameLanguage))
  if (!option.id) return h('div', '全部物业')
  return h('div', { class: 'flex items-center gap-2' }, [
    h('span', namegetlang(option.nameLanguage))
  ])
}

// 筛选处理
const filterHandler = (inputValue, option) => {
  // 同时匹配label和value字段
  return namegetlang(option.nameLanguage).includes(inputValue) ||
    option.id.toString().includes(inputValue)
}

// 选择处理
const handleSelect = (value: string) => {
  // console.log('value', value)
  const selected = wylists.value.find(opt => opt.id === value)

  console.log('selected', selected)
  storage.set('userPms', selected);
  userStore.setuserPms(selected);
  propertyName.value = value > 0 ? namegetlang(selected.nameLanguage) : '全部物业'

  isOpen.value = false

}

// 点击触发元素
const handleClick = () => {
  // 可添加点击逻辑
  isOpen.value = !isOpen.value
}

let nowdate = reactive(new Date());
let sdate = state.selectdate ? state.selectdate.$d : reactive(new Date());
let edate = reactive(new Date());
let datelist = reactive([]);

// 计算365天的总宽度，用于优化滚动性能
const totalWidth = computed(() => {
  return 365 * 70; // 365天 * 每列70px宽度
});
const weekload = () => {
  datelist = [];

  //上方日历列表 - 支持365天数据查询
  for (let i = 0; i <= 364; i++) {
    const a = formatToDateTime(new Date(), 'yyyy-MM-dd');
    const date1 = new Date();
    date1.setDate(sdate.getDate() + i);
    const b = formatToDateTime(date1, 'dd');
    const aa = formatToDateTime(date1, 'yyyy-MM-dd');

    const m = formatToDateTime(date1, 'MM');
    const y = formatToDateTime(date1, 'yyyy'); // 添加年份
    const c = formatToDateTime(date1, 'yyyy-MM-dd 23:59:59'); //时间 当天最晚

    datelist.push({
      allday: c,
      day: b,
      timenum: new Date(c).getTime(),
      week: '周' + formatWeek(date1),
      today: aa == a ? true : false,
      mm: m,
      year: y, // 添加年份数据
    });
  }
};
weekload();
const orderClick = (item, detail) => {
  //订单点击事件 上方父级任务， 详细订单信息

  viewRef.value.openModal(detail);
};
const selectdateChange = (date: dayjs | string, dateString: string) => {
  sdate = new Date(dateString);
  edate = new Date(dateString);
  weekload();
  loadOptions();
  tenListLoad();
};
const maindivscroll = (e) => {
  var element = document.getElementById('rqdiv');
  element.scrollLeft = e.target.scrollLeft;
}

// 优化365天滚动的性能处理
const optimizedScroll = (e) => {
  // 使用节流来优化滚动性能
  if (window.scrollTimeout) {
    clearTimeout(window.scrollTimeout);
  }
  window.scrollTimeout = setTimeout(() => {
    maindivscroll(e);
  }, 16); // 约60fps的节流
}
//控制向左向右
const maindivgo = (type) => {
  //向左
  if (type == 'left') {
    scrollToLeftSmooth();
    // 将滚动条滚动到最左边
  } else {
    //向右
    // 获取目标元素
    scrollToRightSmooth();
  }
};
function scrollToLeftSmooth() {
  var element = document.getElementById('maindiv');
  var element2 = document.getElementById('rqdiv');
  var startScrollLeft = element.scrollLeft;
  var startScrollLeft2 = element2.scrollLeft;
  var startTime = null;
  var duration = 600; // 减少滚动持续时间，提高响应性

  // 计算一周的距离（7天 * 70px每列 = 490px）
  var scrollDistance = 7 * 70; // 490px
  var endScrollLeft = Math.max(0, startScrollLeft - scrollDistance);
  var endScrollLeft2 = Math.max(0, startScrollLeft2 - scrollDistance);

  function scroll(timestamp) {
    if (!startTime) startTime = timestamp;
    var progress = timestamp - startTime;
    var newScrollLeft = startScrollLeft - (progress / duration) * (startScrollLeft - endScrollLeft);
    var newScrollLeft2 = startScrollLeft2 - (progress / duration) * (startScrollLeft2 - endScrollLeft2);
    element.scrollLeft = newScrollLeft;
    element2.scrollLeft = newScrollLeft2;
    if (progress < duration) {
      requestAnimationFrame(scroll);
    }
  }

  requestAnimationFrame(scroll);
}

function scrollToRightSmooth() {
  var element = document.getElementById('maindiv');
  var element2 = document.getElementById('rqdiv');

  var startScrollLeft = element.scrollLeft;
  var startScrollLeft2 = element2.scrollLeft;
  var startTime = null;
  var duration = 600; // 减少滚动持续时间，提高响应性

  // 计算一周的距离（7天 * 70px每列 = 490px）
  var scrollDistance = 7 * 70; // 490px
  var maxScrollLeft = element.scrollWidth - element.clientWidth;
  var maxScrollLeft2 = element2.scrollWidth - element2.clientWidth;
  var endScrollLeft = Math.min(maxScrollLeft, startScrollLeft + scrollDistance);
  var endScrollLeft2 = Math.min(maxScrollLeft2, startScrollLeft2 + scrollDistance);

  function scroll(timestamp) {
    if (!startTime) startTime = timestamp;
    var progress = timestamp - startTime;
    var newScrollLeft = startScrollLeft + (progress / duration) * (endScrollLeft - startScrollLeft);
    var newScrollLeft2 = startScrollLeft2 + (progress / duration) * (endScrollLeft2 - startScrollLeft2);
    element.scrollLeft = newScrollLeft;
    element2.scrollLeft = newScrollLeft2;
    if (progress < duration) {
      requestAnimationFrame(scroll);
    }
  }

  requestAnimationFrame(scroll);
}
//日历中物业的收起和缩放
const updown = (roomitem) => {
  state.RoomList.forEach((f) => {
    if (roomitem.rtUid == f.rtUid) {
      f.updown = !f.updown;
    }
  });
};
//正序
function sortDownDate(a, b) {
  return Date.parse(a.day) - Date.parse(b.day);
}

const selectDay = (item, detail, day, index) => {
  if (state.selectdays.length == 0) {
    state.roomindex = null;
  }
  if (state.roomindex == null || state.roomindex == index) {
    state.roomindex = index;
    detail.selected = !detail.selected;
    detail.day = day;
    detail.roomdata = item;

    //选中状态下
    if (detail.selected) {
      state.selectdays.push(detail);
    } else {
      //反选状态下
      state.selectdays = state.selectdays.filter((f) => {
        if (f.day != day) {
          return f;
        }
      });
    }
  } else {
    message.error(`只能连续选择当前房间`);
  }
  state.selectdays = state.selectdays.sort(sortDownDate);
  console.log(state.selectdays, '选择哪些天');
};
// 加载房型房号数据
const pmsRoomTypeListLoad = () => {
  const all = [
    {
      label: '全部',
      value: '',
    },
  ];

  pmsRoomTypeList({
    page: 1,
    pageSize: 200,
    puid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
  }).then((res) => {
    state.roomTypeoptions = res.list.map((m) => {
      return {
        label: m.name,
        value: m.uid + '',
      };
    });

    state.roomTypeoptions = all.concat(state.roomTypeoptions);
  });
};

// 加载表格数据
const tenListLoad = async () => {
  state.noMore = false;
  state.loading = true;
  state.Availabilities = {};
  state.RoomTypeList = [];
  state.page = 1;
  await DataLoad();

  state.loading = false;
};
const DataLoad = async () => {
  if (state.noMore) {
    return; //没有更多为真，return 跳出
  }

  // 365天按30天分割，计算需要多少个请求
  const totalDays = 365;
  const daysPerRequest = 30;
  const requestCount = Math.ceil(totalDays / daysPerRequest);

  console.log(`开始加载365天数据，分为${requestCount}个请求，每个请求${daysPerRequest}天`);

  // 创建所有请求的Promise数组
  const requests = [];
  for (let i = 0; i < requestCount; i++) {
    const startDate = new Date(sdate);
    startDate.setDate(sdate.getDate() + i * daysPerRequest);

    console.log(`请求${i + 1}: 从${formatToDateTime(startDate, 'yyyy-MM-dd')}开始`);

    const requestPromise = priceCalendar({
      date: formatToDateTime(startDate, 'yyyy-MM-dd'),
      puid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
      ruid: state.roomTypeUid,
      page: state.page,
      limit: state.limit,
    });

    requests.push(requestPromise);
  }

  try {
    // 并发执行所有请求
    const responses = await Promise.all(requests);
    let availableList = {};
    responses.forEach((res, index) => {
      // 取第一个数据的房型用来展示就行，理论上所有请求的房型都一致
      if (!index) {
        const romm = res.RoomTypeList || [];
        state.RoomTypeList = state.RoomTypeList.concat(romm);
      }
      for (const key in res.Availabilities) {
        if (Object.prototype.hasOwnProperty.call(res.Availabilities, key)) {
          const element = res.Availabilities[key] || {};
          if (Object.prototype.hasOwnProperty.call(availableList, key)) {
            Object.assign(availableList[key], element);
          } else {
            availableList[key] = element;
          }
        }
      }
      state.Availabilities = Object.assign(state.Availabilities, availableList);
    });

    // 检查是否还有更多数据（使用第一个响应的Count）
    if (responses[0] && responses[0].Count <= state.RoomTypeList.length) {
      state.noMore = true;
    } else {
      state.page++;
      state.noMore = false;
    }

    console.log(`365天数据加载完成，共${state.RoomTypeList.length}个房型`);

  } catch (error) {
    console.error('加载365天数据失败:', error);
    state.noMore = true;
  }
};
// const userPms = computed(() => userStore.getuserPms);
watch(
  () => userStore.getuserPms,
  () => {
    state.page = 1;
    state.noMore = false;
    tenListLoad();
    state.mainkey++;
  },
  {
    immediate: true,
    deep: true,
  }
);

async function loadOptions() {
  state.options = await Dicts({
    types: ['checkin_status', 'status'],
  });
  console.log('options', options);
}
onMounted(() => {
  wylists.value.unshift({ name: '全部', id: 0, uid: '' })
  state.roomTypeUid = '';
  state.fxlabel = '所有房型';
  // pmsRoomTypeListLoad();
  loadOptions();
  // tenListLoad();
});
</script>

<style lang="less" scoped>
.title-name {
  margin: 15px 0 19px 20px;
  font-weight: 500;
  font-size: 30px;
  color: #3D3D3D;
  line-height: 42px;
}

.todaybg {
  background: #f7f7f7ff;
  border-radius: 20px;
  margin: 0 6px;
}

.today {
  width: 28px;
  background: #053dc8ff;
  color: #fff;
  border-radius: 10px;
  margin: auto;
  font-weight: 500;
}

.selectddiv {
  height: 91%;
  background-color: #a3cfff;
  margin: 2px;
}

.br1 {
  border: 1px solid #e7e7e7ff;
}

.brl1 {
  border-left: 1px solid #e7e7e7ff;
}

.brr1 {
  border-right: 1px solid #e7e7e7ff;
}

.brt1 {
  border-top: 1px solid #e7e7e7ff;
}

.brb1 {
  border-bottom: 1px solid #e7e7e7ff;
}

.dyg {
  border-right: 1px solid #e7e7e7ff;
  border-bottom: 1px solid #e7e7e7ff;
  padding: 5px;
}

.dyg2 {
  border-right: 1px solid #e7e7e7ff;
  border-bottom: 1px solid #e7e7e7ff;
}

.dyg999 {
  border-right: 1px solid #e7e7e7ff;
  border-bottom: 1px solid #e7e7e7ff;
}

.lrx {
  text-align: center;
  display: flex;
  width: 188px;
  flex-shrink: 0;
  flex-direction: row;
  flex-wrap: wrap;
  overflow: hidden;
}

.hxrq {
  width: 2170px;
  flex-grow: 1;
  display: flex;
  z-index: 1;
  overflow-x: hidden;
  overflow-y: hidden;
  box-sizing: border-box;
  justify-content: flex-start;
}

.hxrq2 {
  display: flex;
  z-index: 1;
  overflow-x: hidden;
  overflow-y: hidden;
  box-sizing: border-box;
  justify-content: flex-start;
}

.rq {
  width: 70px;
  font-size: 14px;
  box-sizing: border-box;
  flex-shrink: 0;
  text-align: center;
  min-width: 70px;
  /* 确保最小宽度，支持365天显示 */
}

.righthx {
  width: 100%;
  flex-grow: 1;
  display: flex;
  z-index: 1;
  height: 50px;
}

.rqrighthx {
  height: 50px;
  font-size: 14px;
  line-height: 3;
  box-sizing: border-box;
  flex-shrink: 0;
  text-align: center;
  display: flex;
}

.kucun {
  height: 50px;
  font-size: 14px;
  line-height: 3;
  box-sizing: border-box;
  flex-shrink: 0;
  text-align: center;
  display: flex;
}

.order {
  width: 100%;
  height: 100%;
  cursor: pointer;
  background-color: #446ea5;
}

.order2 {
  width: 100%;
  height: 100%;
  cursor: pointer;
  background-color: #c39228;
}

.order3 {
  width: 100%;
  height: 100%;
  cursor: pointer;
  background-color: #74b0ff;
}

.nullorder {
  width: 100%;
  height: 100%;
  cursor: pointer;
}

.nullorder:hover {
  background-color: #e7e7e7ff;
}

.orderdyg {
  cursor: pointer;
}

.xzxinxi {
  height: 32px;
  background: #053dc83b;
  color: #053dc8;
  overflow: hidden;
  padding: 8px 0;
}

.kyg {
  color: rgb(18, 197, 132);
  background-color: rgb(231, 249, 243);
  border-radius: 5px;
  width: 50px;
  margin: auto;
  font-size: 11px;
}

.manfang {
  color: #333;
  background-color: rgba(152, 152, 152, 0.386);
  border-radius: 5px;
  width: 50px;
  margin: auto;
  font-size: 11px;
}

.roomclear {
  cursor: pointer;
  background-color: #00000040;
  width: 100%;
  height: 100%;
  padding-top: 15%;
  position: relative;
  z-index: 0;

  img {
    position: absolute;
    bottom: 5px;
    right: 2px;
    width: 20px;
    z-index: 99;
  }
}

.opup {
  cursor: pointer;
}

.goleft {
  position: absolute;
  left: 240px;
  line-height: 86px;
  width: 30px;
  text-align: center;
  z-index: 2;
  height: 85px;
  cursor: pointer;
  // box-shadow: 5px 0px 5px 0px #65656533;
}

.goright {
  position: absolute;
  right: 0px;
  line-height: 86px;
  width: 30px;
  text-align: center;
  z-index: 2;
  height: 85px;
  cursor: pointer;
  // box-shadow: -5px 0px 5px 0px #65656533;
}
</style>
