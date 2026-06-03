<template>
  <div class="title-name">
    <div class="dropdown-container" style="width: 500px">
      <!-- 自定义触发元素 -->
      <n-popover
        trigger="click"
        placement="bottom-start"
        :show-arrow="false"
        :overlap="false"
        raw
      >
        <template #trigger>
          <div class="trigger-div" @click="handleClick">
            {{ propertyName }}
            <n-icon :component="isOpen ? ChevronUp : ChevronDown" />
          </div>
        </template>

        <!-- 下拉选择器 -->
        <n-select
          style="width: 500px"
          v-model:value="selectedValue"
          filterable
          :filter="filterHandler"
          :options="sortedOptions"
          :render-label="renderLabel"
          value-field="id"
          :menu-props="{ style: { width: '500px' } }"
          @update:value="handleSelect"
        />
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
  <div
    class="bgfff"
    style="width: 100%; padding: 10px; border-radius: 4px; overflow-y: hidden"
    v-else
  >
    <div class="mt-2">
      <div class="flex-row f12 p-2" style="position: relative">
        <div class="mr-2 mt-2">
          <div
            :class="'mr-1 bg-' + getOptionTag(state.options.checkin_status, 'before_checkin')"
            style="float: left; width: 18px; height: 18px; border-radius: 2px"
          >
          </div>
          {{ getOptionLabel(state.options.checkin_status, 'before_checkin') }}
        </div>
        <div class="mr-2 mt-2">
          <div
            :class="'mr-1 bg-' + getOptionTag(state.options.checkin_status, 'checked_in')"
            style="float: left; width: 18px; height: 18px; border-radius: 2px"
          >
          </div>
          {{ getOptionLabel(state.options.checkin_status, 'checked_in') }}
        </div>
        <div class="mr-2 mt-2">
          <div
            class="mr-1 bg-info"
            style="float: left; width: 18px; height: 18px; border-radius: 2px"
          >
          </div>
          {{ getOptionLabel(state.options.checkin_status, 'checked_out') }}
        </div>
        <div class="flex-item text-r">
          <a-date-picker
            v-model:value="state.selectdate"
            :placeholder="translang('初始为当天')"
            style="width: 220px;border-radius: 4px;"
            @change="selectdateChange"
          ></a-date-picker>
          <!-- <n-button type="primary" class="mr-1" @click="state.yuyueopen = true">预约订房</n-button>
          <n-button type="warning" class="mr-1">关房</n-button> -->
        </div>
      </div>
    </div>

    <div class="flex-column bgfff brt1 brl1" style="height: 78vh; position: relative">
      <!-- 向左向右 操作 -->
      <!-- 向左 -->
      <div class="goleft" @click="maindivgo('left')">
        <n-icon size="18" color="#333">
          <ChevronBackOutline />
        </n-icon>
      </div>
      <!-- 向右 -->
      <div class="goright" @click="maindivgo('right')">
        <n-icon size="18" color="#333">
          <ChevronForwardOutline />
        </n-icon>
      </div>
      <!-- 日期 -->
      <div class="flex-row">
        <div class="dyg rq" style="width: 240px"> </div>
        <div
          class="flex-item"
          id="rqdiv"
          style="
            width: 100%;
            flex-grow: 1;
            display: flex;
            z-index: 1;
            line-height: 25px;
            height: 86px;
            overflow: hidden;
          "
        >
          <div class="dyg rq text-c" v-for="(item, index) in datelist" :key="index">
            <div :class="item.today ? 'todaybg' : ''">
              <div
                :class="'周日，周六'.indexOf(item.week) != -1 ? 'cred' : 'c999'"
                style="font-size: 12px"
              >
                {{ item.week }}
              </div>
              <div class="c333 fw" style="font-size: 16px">
                <div :class="item.today ? 'today' : ''">{{ item.day }}</div>
              </div>
              <div class="text-c c999" style="font-size: 11px">
                <span v-if="item.today" style="color: #053dc8ff">今日</span>
                <span v-else>{{ item.mm }}月</span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="flex-item" style="overflow-y: hidden">
        <n-infinite-scroll style="height: 69.6vh; overflow-y: auto" :distance="20" @load="DataLoad">
          <div class="flex-row">
            <div style="width: 240px">
              <div v-for="(roomitem, index) in state.RoomList" :key="index">
                <div
                  v-if="
                    index == 0 ||
                    roomitem.room_type.uid != state.RoomList[index - 1].room_type.uid
                  "
                  style="line-height: 23px; height: 50px"
                  class="dyg999"
                >
                  <div
                    class="f14 c333 fw ml-2"
                    @click="updown(roomitem)"
                    style="overflow: hidden; text-overflow: ellipsis; white-space: nowrap"
                  >
                    <!-- 收起 打开 -->
                    <span class="opup" v-if="!roomitem.updown">
                      <n-icon size="16" color="#333">
                        <ChevronUp />
                      </n-icon>
                    </span>
                    <span class="opup" v-else>
                      <n-icon size="16" color="#333">
                        <ChevronDown />
                      </n-icon>
                    </span>
                    <span class="ml-2">{{ roomitem.room_type.name }}</span>
                  </div>
                  <div class="ml-8">
                    <span>
                      <n-icon size="15" color="#999">
                        <BusinessOutline />
                      </n-icon>
                    </span>
                    <span class="ml-1 f12 c999">{{ roomitem.room_type.property.name }}</span>
                  </div>
                </div>
                <div
                  class="dyg999 c333"
                  style="
                    width: 100%;
                    height: 50px;
                    line-height: 46px;
                    padding-left: 45px;
                    font-weight: 550;
                  "
                  v-show="!roomitem.updown"
                >
                  {{ roomitem.roomNo }}
                </div>
              </div>
            </div>
            <div id="maindiv" class="flex-item" @scroll="maindivscroll">
              <div v-for="(roomitem, index) in state.RoomList">
                <!-- 存在状态  库存 -->
                <div
                  v-if="
                    index == 0 ||
                    roomitem.room_type.uid != state.RoomList[index - 1].room_type.uid
                  "
                  class="righthx"
                >
                  <div class="kucun">
                    <div
                      v-for="(dygitem, index2) in state.ShowData[roomitem.uid].date"
                      :key="index2"
                      style="width: 70px; height: 50px; padding: 1px 0; line-height: 20px"
                      class="dyg999 c999"
                    >
                      <div
                        style="color: #333; font-size: 14px"
                        v-if="Object.keys(state.Availabilities).length > 0 &&state.Availabilities[roomitem.rtUid] && state.Availabilities[roomitem.rtUid][index2]"
                      >
                        {{ state.Availabilities[roomitem.rtUid][index2][0].allotment }}
                        <div class="kyg">{{ translang('可预订') }}</div>
                      </div>
                      <div v-else style="color: #333; font-size: 14px">
                        0
                        <div class="manfang">{{ translang('无库存') }}</div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="righthx" v-show="!roomitem.updown">
                  <div
                    class="rqrighthx"
                    v-for="(dygitem, index2) in state.ShowData[roomitem.uid].date"
                    :key="index2"
                    style="position: relative"
                  >
                    <!-- 存在的订单数据 整个条数拉长 订单数据 -->
                    <div
                      v-if="dygitem.uid"
                      style="
                        width: 70px;
                        height: 100%;
                        padding: 10px 0;
                        position: relative;
                      "
                      class="dyg999"
                    >
                      <div
                        class="orderdyg height100"
                        style="position: absolute; left: 50%; z-index: 99; height: 30px"
                        :style="{ width: dygitem.days * 69 + 'px' }"
                        @click="orderClick(roomitem, dygitem)"
                      >
                        <div class="flex-row height100">
                          <div
                            :class="
                              dygitem.checkinStatus == 'checked_out'
                                ? 'bg-info'
                                : 'bg-' +
                                  getOptionTag(state.options.checkin_status, dygitem?.checkinStatus)
                            "
                            style="width: 5px"
                          >
                          </div>
                          <div
                            :class="
                              dygitem.checkinStatus == 'checked_out'
                                ? 'bg-info-dan'
                                : 'bg-' +
                                  getOptionTag(
                                    state.options.checkin_status,
                                    dygitem?.checkinStatus
                                  ) +
                                  '-dan'
                            "
                            style="width: calc(100% - 10px); text-align: left; line-height: 2;
                        white-space: nowrap;
                        overflow: hidden;
                        text-overflow: ellipsis;"
                          >
                            <span
                              v-if="dygitem.days > 0"
                              :class="
                                dygitem.checkinStatus == 'checked_out'
                                  ? 'info'
                                  : getOptionTag(
                                      state.options.checkin_status,
                                      dygitem?.checkinStatus
                                    )
                              "
                              style="font-size: 12px; margin-left: 2px"
                              >{{ dygitem.reservation.guest_profile_detail.fullName }}</span
                            >
                          </div>
                          <div class="flex-item"></div>
                        </div>
                      </div>
                    </div>

                    <!-- 不存在订单 空天数 -->
                    <div
                      v-else
                      class="nullorder dyg999"
                      :style="{ width: dygitem.days * 70 + 'px' }"
                      @click="selectDay(roomitem, dygitem, index2, index)"
                    >
                      <div v-if="dygitem.selected" class="selectddiv">
                        <n-icon size="25" color="#fff" class="ml-1">
                          <CheckmarkSharp />
                        </n-icon>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </n-infinite-scroll>
      </div>
    </div>

    <a-drawer
      v-model:open="state.yuyueopen"
      class="custom-class"
      root-class-name="root-class-name"
      :title="translang('预约')"
      width="580"
      placement="right"
    >
      <Appointment :selectdays="state.selectdays" />
    </a-drawer>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, reactive, ref, watch, h, computed } from 'vue';
  import {
    BusinessOutline,
    CheckmarkSharp,
    ChevronUp,
    ChevronDown,
    ChevronForwardOutline,
    ChevronBackOutline,
  } from '@vicons/ionicons5';
  import { roomStatus } from '@/api/comm';
  import { List as pmsRoomTypeList } from '@/api/pmsRoomType';
  import { List as pmsRoomUnitList } from '@/api/pmsRoomUnit';

  import { formatToDateTime, formatWeek } from '@/utils/dateUtil';
  import { State, options } from '../curdDemo/model';
  import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { getlang, translang } from '@/utils/smjcomm';

  import { Dicts } from '@/api/dict/dict';
  import { useUserStore } from '@/store/modules/user';
  import Appointment from './comm/appointment.vue';
  import View from '@/views/pmsAppReservation/view.vue';
  import { useMessage } from 'naive-ui';
  import {storage} from "@/utils/Storage";
  const viewRef = ref();

  const message = useMessage();
  const userStore = useUserStore();
  const state = reactive({
    page: 1,
    limit: 20,
    mainkey: 1,
    noMore: false,
    yuyueopen: false,
    roomindex: null,
    fxlabel: '所有房型',
    roomTypeUid: '',
    roomUtilUid: '',
    roomTypeoptions: [],
    roomUtiloptions: [],
    loading: false,
    Reservation: [],
    RoomList: [],
    ShowData: {},
    options: [],
    reservationobj: {}, // 订单房间数据详情
    selectdays: [], //选的 天数
    sx_fx: '',
    sx_fh: '',
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
  let nowdate = reactive(new Date());
  let sdate = null;
  let edate = reactive(new Date());
  let datelist = reactive([]);
  const weekload = () => {
    datelist = [];
    sdate = state.selectdate ? state.selectdate.$d : reactive(new Date());
    sdate.setDate(sdate.getDate() - 1);

    edate.setDate(sdate.getDate() + 30);

    //上方日历列表
    for (let i = 0; i <= 30; i++) {
      const a = formatToDateTime(new Date(), 'MM-dd');
      const date1 = new Date();
      date1.setDate(sdate.getDate() + i);

      const b = formatToDateTime(date1, 'dd');
      const d = formatToDateTime(date1, 'MM-dd');
      const m = formatToDateTime(date1, 'MM');
      const c = formatToDateTime(date1, 'yyyy-MM-dd 23:59:59'); //时间 当天最晚

      datelist.push({
        allday: c,
        day: b,
        timenum: new Date(c).getTime(),
        week: '周' + formatWeek(date1),
        today: d == a ? true : false,
        mm: m,
      });
    }
  };
  weekload();
  const orderClick = (item, detail) => {
    //订单点击事件 上方父级任务， 详细订单信息
    viewRef.value.openModal(detail.reservation);
  };
  const selectdateChange = (date: dayjs | string, dateString: string) => {
    sdate = new Date(dateString);
    edate = new Date(dateString);
    weekload();
    loadOptions();
    tenListLoad();
    state.mainkey++;
  };

  const maindivscroll = (e) => {
    var element = document.getElementById('rqdiv');
    element.scrollLeft = e.target.scrollLeft;
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
    var duration = 800; // 滚动持续时间，单位为毫秒

    function scroll(timestamp) {
      if (!startTime) startTime = timestamp;
      var progress = timestamp - startTime;
      var newScrollLeft = startScrollLeft - (progress / duration) * startScrollLeft;
      var newScrollLeft2 = startScrollLeft2 - (progress / duration) * startScrollLeft2;
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
    var endScrollLeft = element.scrollWidth - element.clientWidth;
    var startScrollLeft2 = element2.scrollLeft;
    var endScrollLeft2 = element2.scrollWidth - element2.clientWidth;
    var startTime = null;
    var duration = 1000; // 滚动持续时间，单位为毫秒

    function scroll(timestamp) {
      if (!startTime) startTime = timestamp;
      var progress = timestamp - startTime;
      var newScrollLeft =
        startScrollLeft + (progress / duration) * (endScrollLeft - startScrollLeft);
      var newScrollLeft2 =
        startScrollLeft2 + (progress / duration) * (endScrollLeft2 - startScrollLeft2);
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
    pmsRoomUnitList({
      page: 1,
      pageSize: 200,
      puid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
      rt_uid: state.roomTypeUid ? state.roomTypeUid + '' : '',
    }).then((res) => {
      state.roomUtiloptions = res.list.map((m) => {
        return {
          label: m.roomNo,
          value: m.uid + '',
        };
      });

      state.roomUtiloptions = all.concat(state.roomUtiloptions);
    });
  };

  // 加载表格数据
  const tenListLoad = async () => {
    state.noMore = false;
    state.page = 1;
    state.loading = true;
    state.RoomList = [];
    state.Reservation = [];
    state.Availabilities = {};
    await DataLoad();
    state.loading = false;
    console.log('所有时间区域 ', state.ShowData);
  };
  const DataLoad = async () => {
    if (state.noMore) {
      return; //没有更多为真，return 跳出
    }
    const res = await roomStatus({
      startDate: formatToDateTime(sdate, 'yyyy-MM-dd'),
      endDate: formatToDateTime(edate, 'yyyy-MM-dd'),
      p_uid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
      roomTypeUid: state.roomTypeUid,
      roomUtilUid: state.roomUtilUid,
      page: state.page,
      limit: state.limit,
    });
    const romm = res.RoomList.map((m) => {
      return {
        ...m,
        updown: false, //每一个节点的收起和打开
      };
    });
    state.RoomList = state.RoomList.concat(romm);
    state.RoomList.forEach((f) => {
      f.updown = false;
    });
    state.Reservation = state.Reservation.concat(res.Reservation);
    state.Availabilities = Object.assign(state.Availabilities, res.Availabilities);
    for (let k in res.ShowData) {
      const aa = JSON.stringify(res.ShowData[k].date);
      if (aa == '{}') {
        for (let i = 0; i <= 30; i++) {
          const date1 = new Date();
          date1.setDate(sdate.getDate() + i);
          const b = formatToDateTime(date1, 'yyyy-MM-dd');
          res.ShowData[k].date[b] = {
            checkinStatus: null,
            days: 1,
            selected: false,
            status: null,
            uid: null,
          };
        }
      }
    }
    state.ShowData = Object.assign(state.ShowData, res.ShowData);

    if (res.Count <= state.RoomList.length) {
      state.noMore = true;
    } else {
      state.page++;
      state.noMore = false;
    }
  };
  // const userPms = computed(() => userStore.getuserPms);
  watch(
    () => userStore.getuserPms,
    () => {
      // pmsRoomTypeListLoad();
      tenListLoad();
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

    // console.log("wylists.value", wylists.value)

    const selected = wylists.value.find(opt => opt.id === selectedValue.value)
    const others = wylists.value.filter(opt => opt.id !== selectedValue.value)
    if(selectedValue.value > 0){
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
    console.log('value', value)
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


  onMounted(() => {
    wylists.value.unshift({ name: '全部', id: 0, uid: '' })
    console.log("wylists.value111", wylists.value)
    state.roomTypeUid = '';
    state.fxlabel = '所有房型';
    // pmsRoomTypeListLoad();
    loadOptions();
    // tenListLoad();
  });
</script>

<style lang="less" scoped>
  .title-name{
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
    background: #fff;
    height: 85px;
    cursor: pointer;
    box-shadow: 5px 0px 5px 0px #65656533;
  }
  .goright {
    position: absolute;
    right: 0px;
    line-height: 86px;
    width: 30px;
    text-align: center;
    z-index: 2;
    background: #fff;
    height: 85px;
    cursor: pointer;
    box-shadow: -5px 0px 5px 0px #65656533;
  }
</style>
