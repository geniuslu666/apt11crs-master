<template>
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
  <div class="bgfff" style="width: 100%" v-else>
    <div class="mt-2">
      <div class="flex-row f12 p-2">
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
            :class="'mr-1 bg-' + getOptionTag(state.options.checkin_status, 'checked_out')"
            style="float: left; width: 18px; height: 18px; border-radius: 2px"
          >
          </div>
          {{ getOptionLabel(state.options.checkin_status, 'checked_out') }}
        </div>
        <div class="flex-item text-r">
          <a-date-picker
            v-model:value="state.selectdate"
            placeholder="初始为当天"
            style="width: 220px"
            @change="selectdateChange"
          ></a-date-picker>
          <!-- <n-button type="primary" class="mr-1" @click="state.yuyueopen = true">预约订房</n-button>
          <n-button type="warning" class="mr-1">关房</n-button> -->
        </div>
      </div>
    </div>

    <div class="flex-row mt-2 bgfff brt1 brl1">
      <!--  左 -->
      <div style="width: 188px">
        <div
          class="text-c"
          :style="{ height: '64px' }"
          style="border-right: 1px solid #e0efff; border-bottom: 1px solid #e0efff"
        >
          <div class="lh-2" style="width: 100%; height: 32px; border-bottom: 1px solid #e0efff">
            <!-- {{ userStore.getuserPms }} -->
            {{ namegetlang(userStore.getuserPms.nameLanguage) }}
          </div>
          <div class="c999 f12 flex-row" style="width: 100%">
            <!-- <div class="flex-item xzxinxi">筛选</div> -->
            <div class="" style="width: 50%">
              <n-popselect
                v-model:value="state.roomTypeUid"
                :options="state.roomTypeoptions"
                trigger="click"
                @update:value="roomTypeoptionsupdate"
              >
                <n-button quaternary>
                  {{ state.fxlabel }}

                  <n-icon size="14" color="#053dc8" class="ml-1"> <FunnelOutline /> </n-icon>
                </n-button>
              </n-popselect>
            </div>
            <div class="flex-item">
              <n-popselect
                v-model:value="state.roomUtilUid"
                :options="state.roomUtiloptions"
                trigger="click"
                @update:value="roomUtiloptionsupdate"
              >
                <n-button quaternary>
                  {{ state.fjlabel }}
                  <n-icon size="14" color="#053dc8" class="ml-1"> <FunnelOutline /> </n-icon>
                </n-button>
              </n-popselect>
            </div>
          </div>
        </div>

        <div
          class="flex-row"
          style="height: 50px; line-height: 2"
          v-for="(roomitem, index) in state.RoomList"
          :key="index"
        >
          <div class="flex-item dyg2 text-c nohuanhang" style="overflow-y: hidden"
            >
            {{ roomitem.name }}
          </div>
          <div class="dyg2 text-c" style="width: 35%; padding: 2px">
            <!-- 正常状态 -->
            <div style="line-height: 45px">{{ roomitem.roomNo }}</div>
            <!-- 关房状态 -->
            <!-- <div class="roomclear">
              <span>{{ roomitem.roomNo }}</span>
              <a-tooltip placement="bottom">
                <template #title>
                  <span>转为净房</span>
                </template>
                <img style="object-fit: cover" src="@/assets/images/tiaozou.png" />
              </a-tooltip>
            </div> -->
            <!-- 需要清扫 -->
            <!-- <div class="roomclear">
              <span>{{ roomitem.roomNo }}</span>
              <a-tooltip placement="bottom">
                <template #title>
                  <span>转为净房</span>
                </template>
                <img style="object-fit: cover" src="@/assets/images/tiaozou.png" />
              </a-tooltip>
            </div> -->
          </div>
        </div>
      </div>

      <!-- 右边 主操作页面 -->
      <div class="flex-item" style="overflow-x: auto">
        <div class="hxrq">
          <div style="width: 100%; flex-grow: 1; display: flex; z-index: 1">
            <div class="dyg rq" v-for="(item, index) in datelist" :key="index">
              <span v-if="item.today" class="cblue">今天</span>
              <span class="ml-3">{{ item.day }}</span>
              <span
                v-if="!item.today"
                :class="'日，六'.indexOf(item.week) != -1 ? 'ml-2 cred' : 'ml-2'"
                >{{ item.week }}</span
              >
            </div>
          </div>
        </div>

        <div v-for="(roomitem, index) in state.RoomList" class="righthx">
          <div
            class="rqrighthx"
            v-for="(dygitem, index2) in state.ShowData[roomitem.uid].date"
            :key="index2"
          >
            <!-- 存在的订单数据 整个条数拉长 订单数据 -->

            <div v-if="dygitem.uid" style="width: 100%; height: 100%; padding: 1px 0">
              <div
                :class="
                  'orderdyg height100 bg-' +
                  getOptionTag(state.options.checkin_status, dygitem?.checkinStatus)
                "
                :style="{ width: dygitem.days * 100 + 'px' }"
                @click="orderClick(roomitem, dygitem)"
              >
                <div class="flex-row height100">
                  <div style="background-color: #006dc0; width: 5px"> </div>
                </div>
              </div>
              <!--   <n-popover trigger="hover" :show-arrow="false" placement="right-end">
                <template #trigger>
                 
                </template>
                <div class="bgfff c333 lh-2">
                  <div style="width: 300px">
                    <div class="text-r">
                     
                      <n-tag
                        :bordered="false"
                        :type="getOptionTag(state.options.checkin_status, dygitem?.checkinStatus)"
                      >
                        {{ getOptionLabel(state.options.checkin_status, dygitem?.checkinStatus) }}
                      </n-tag>
                    </div>
                    <div>
                      <n-icon size="18" color="#999" class="mr-1">
                        <PersonOutline />
                      </n-icon>
                      <span class="ml-2 mr-3" f>{{ state.reservationobj.adultCount }}</span>
                      <n-icon size="18" color="#999">
                        <Accessibility />
                      </n-icon>
                      <span class="ml-2">{{ state.reservationobj.childCount }}</span>
                    </div>

                    <div>
                      <n-icon size="18" color="#999" class="mr-1">
                        <CalendarOutline />
                      </n-icon>
                      <span class="ml-2 mr-2">
                        {{}}
                        入住&nbsp;&nbsp;{{ state.reservationobj.checkinDate }}
                        {{ state.reservationobj.checkinTime }}
                      </span>
                      |
                      <span class="ml-2 mr-3"
                        >离店&nbsp;&nbsp;{{ state.reservationobj.checkoutDate }}
                        {{ state.reservationobj.checkoutTime }}
                      </span>
                    </div>
                    <div
                      style="border-top: 1px solid #7878782b; border-bottom: 1px solid #7878782b"
                      class="mt-1 pt-1 pb-1"
                    >
                      <n-icon size="18" color="#999" class="mr-1">
                        <CalendarOutline />
                      </n-icon>
                      <span class="ml-2 mr-2"
                        >预定费&nbsp;&nbsp;￥{{ state.reservationobj.bookingFee }}</span
                      >
                      |
                      <span class="ml-2 mr-3"
                        >渠道费&nbsp;&nbsp;￥{{ state.reservationobj.channelFee }}</span
                      >
                      |
                      <span class="ml-2 mr-3"
                        >清洁费&nbsp;&nbsp;￥{{ state.reservationobj.cleaningFee }}</span
                      >
                    </div>
                    <div> 备注 </div>
                  </div>
                </div> 
              </n-popover>-->
            </div>

            <!-- 不存在订单 空天数 -->
            <div
              v-else
              class="nullorder dyg999"
              :style="{ width: dygitem.days * 100 + 'px' }"
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
    <a-drawer
      v-model:open="state.yuyueopen"
      class="custom-class"
      root-class-name="root-class-name"
      title="预约"
      width="580"
      placement="right"
    >
      <Appointment :selectdays="state.selectdays" />
    </a-drawer>
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref, watch, computed } from 'vue';
  import {
    CaretDownSharp,
    PersonOutline,
    Accessibility,
    CalendarOutline,
    CheckmarkSharp,
    FunnelOutline,
  } from '@vicons/ionicons5';
  import { roomStatus } from '@/api/comm';
  import { List as pmsRoomTypeList } from '@/api/pmsRoomType';
  import { List as pmsRoomUnitList } from '@/api/pmsRoomUnit';

  import { formatToDateTime, formatWeek } from '@/utils/dateUtil';
  import { State, options } from '../curdDemo/model';
  import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { getlang } from '@/utils/smjcomm';

  import { Dicts } from '@/api/dict/dict';
  import { useUserStore } from '@/store/modules/user';
  import { usePermission } from '@/hooks/web/usePermission';
  import loading from 'naive-ui/es/_internal/loading';
  import Appointment from './comm/appointment.vue';
  import View from '@/views/pmsAppReservation/view.vue';
  import { useMessage } from 'naive-ui';
  const viewRef = ref();

  const message = useMessage();
  const userStore = useUserStore();
  const state = reactive({
    mainkey: 1,
    yuyueopen: false,
    roomindex: null,
    fjlabel: '房间',
    fxlabel: '房型',
    roomTypeUid: '',
    roomUtilUid: '',
    roomTypeoptions: [],
    roomUtiloptions: [],
    loading: false,
    Reservation: [],
    RoomList: [],
    ShowData: [],
    options: [],
    reservationobj: {}, // 订单房间数据详情
    selectdays: [], //选的 天数
    sx_fx: '',
    sx_fh: '',
    selectdate: '',
  });
  const namegetlang = (data) => {
    if (data) {
      return getlang(data, userStore.language).content;
    } else {
      return '暂无物业名';
    }
  };
  let nowdate = reactive(new Date());
  let sdate = reactive(new Date());
  let edate = reactive(new Date());
  let datelist = reactive([]);

  const weekload = () => {
    datelist = [];
    sdate.setDate(sdate.getDate() - 1);

    edate.setDate(sdate.getDate() + 30);

    //上方日历列表
    for (let i = 0; i <= 30; i++) {
      const a = formatToDateTime(new Date(), 'MM-dd');
      const date1 = new Date();
      date1.setDate(sdate.getDate() + i);

      const b = formatToDateTime(date1, 'MM-dd');

      const c = formatToDateTime(date1, 'yyyy-MM-dd 23:59:59'); //时间 当天最晚

      datelist.push({
        allday: c,
        day: b,
        timenum: new Date(c).getTime(),
        week: formatWeek(date1),
        today: b == a ? true : false,
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
    pmsRoomTypeListLoad();
    loadOptions();
    tenListLoad();
    state.mainkey++;
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
    state.loading = true;
    const res = await roomStatus({
      startDate: formatToDateTime(sdate, 'yyyy-MM-dd'),
      endDate: formatToDateTime(edate, 'yyyy-MM-dd'),
      p_uid: userStore.getuserPms.uid ? userStore.getuserPms.uid : '',
      roomTypeUid: state.roomTypeUid,
      roomUtilUid: state.roomUtilUid,
      page: 1,
      limit: 100,
    });
    state.RoomList = res.RoomList;
    state.Reservation = res.Reservation;
    state.ShowData = res.ShowData;
    for (let k in state.ShowData) {
      const aa = JSON.stringify(state.ShowData[k].date);
      if (aa == '{}') {
        for (let i = 0; i <= 30; i++) {
          const date1 = new Date();
          date1.setDate(sdate.getDate() + i);
          const b = formatToDateTime(date1, 'yyyy-MM-dd');
          state.ShowData[k].date[b] = {
            checkinStatus: null,
            days: 1,
            selected: false,
            status: null,
            uid: null,
          };
        }
      }
    }
    state.loading = false;
    console.log('所有时间区域 ', state.ShowData);
  };
  // const userPms = computed(() => userStore.getuserPms);
  watch(
    () => userStore.getuserPms,
    () => {
      pmsRoomTypeListLoad();
      tenListLoad();
    },
    {
      immediate: true,
      deep: true,
    }
  );

  const roomTypeoptionsupdate = (value, option) => {
    state.fxlabel = option.label.slice(0, 8);
    pmsRoomTypeListLoad();
    tenListLoad();
    state.roomUtilUid = '';
    state.fjlabel = '房间';
  };
  const roomUtiloptionsupdate = (value, option) => {
    state.fjlabel = option.label.slice(0, 8);
    tenListLoad();
  };
  async function loadOptions() {
    state.options = await Dicts({
      types: ['checkin_status', 'status'],
    });
    console.log('options', options);
  }
  onMounted(() => {
    state.roomUtilUid = '';
    state.fjlabel = '房间';
    state.roomTypeUid = '';
    state.fxlabel = '房型';
    pmsRoomTypeListLoad();
    loadOptions();
    tenListLoad();
  });
</script>

<style lang="less" scoped>
  .selectddiv {
    height: 91%;
    background-color: #a3cfff;
    margin: 2px;
  }
  .br1 {
    border: 1px solid #e0efff;
  }
  .brl1 {
    border-left: 1px solid #e0efff;
  }
  .brr1 {
    border-right: 1px solid #e0efff;
  }
  .brt1 {
    border-top: 1px solid #e0efff;
  }
  .dyg {
    border-right: 1px solid #e0efff;
    border-bottom: 1px solid #e0efff;
    padding: 5px;
    box-shadow: inset -5px 0px 8px #4d91d90f;
  }
  .dyg2 {
    border-right: 1px solid #e0efff;
    border-bottom: 1px solid #e0efff;
    padding: 10px;
    box-shadow: inset -5px 0px 8px #4d91d90f;
  }
  .dyg999 {
    border-right: 1px solid hsla(0, 0%, 47%, 0.17);
    border-bottom: 1px solid hsla(0, 0%, 47%, 0.17);
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
    width: 3100px;
    flex-grow: 1;
    display: flex;
    z-index: 1;
    overflow-x: hidden;
    overflow-y: hidden;
    box-sizing: border-box;
    justify-content: flex-start;
    box-shadow: 0 4px 16px 0 rgba(26, 111, 255, 8%);
  }

  .rq {
    padding: 10px;
    height: 64px;
    width: 100px;
    font-size: 14px;
    line-height: 3;
    box-sizing: border-box;
    flex-shrink: 0;
    text-align: center;
    display: flex;
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
    background-color: #e0efff;
  }
  .orderdyg {
    cursor: pointer;
    border-radius: 5px;
  }
  .orderdyg:hover {
    box-shadow: 1px 1px 10px #9ca3af;
  }
  .xzxinxi {
    height: 32px;
    background: #053dc83b;
    color: #053dc8;
    overflow: hidden;
    padding: 8px 0;
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
</style>
