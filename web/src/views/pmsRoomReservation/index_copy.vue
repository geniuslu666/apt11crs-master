<template>
  <div>
    <n-loading-bar-provider>
      <div class="flex-column" style="height: calc(100vh - 118px)">
        <n-card :bordered="false" title="AIRHOST订单">
          <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
          <div class="flex-row">
            <n-input
            v-model:value="state.search.order_sn"
            type="text"
            placeholder="订单号"
            style="width: 180px"
            @keyup.enter="loadDataTable"
          />
            <n-input
              v-model:value="state.search.out_order_sn"
              class="ml-1"
              type="text"
              placeholder="外部订单号"
              style="width: 180px"
              @keyup.enter="loadDataTable"
            />
            <a-date-picker
              v-model:value="state.search.checkinDate"
              style="width: 180px"
              class="ml-1"
              placeholder="入住日期"
              @change="loadDataTable"
            />
            <a-date-picker
              v-model:value="state.search.checkoutDate"
              style="width: 180px"
              class="ml-1"
              placeholder="退房日期"
              @change="loadDataTable"
            />
            <n-select
              class="ml-1"
              v-model:value="state.search.status"
              :options="optionsstatus.status"
              clearable
              style="width: 180px"
            />
            <n-select
              class="ml-1"
              v-model:value="state.search.checkinStatus"
              :options="optionsstatus.checkin_status"
              clearable
              style="width: 180px"
            />
            <n-button type="primary" class="ml-1" @click="loadDataTable"> 搜索 </n-button>
          </div>
        </n-card>

        <div class="flex-item pl-5 bgfff">
          <a-table
            :pagination="false"
            :dataSource="state.dataSource"
            :columns="state.columns"
            bordered
            size="middle"
          >
            <template #bodyCell="{ column, record }">
              <template v-if="column.key === 'orderSn'">
                <div class="flex-row" v-if="record.guest_profile_detail">
                  <img :src="record.property_detail.cover" style="width: 50px; height: 65px" />
                  <div class="flex-item" style="padding-left: 10px">
                    <div class="c999">{{ record.property_detail.name }}</div>
                    <div>
                      <span class="fw csuccess">{{ record.outOrderSn }}</span> &nbsp;&nbsp;{{
                        record.orderSn
                      }}
                    </div>
                    <div v-if="record.guest_profile_detail">
                      <span v-if="record.guest_profile_detail.fullName" class="cblue fw f14">
                        {{ record.guest_profile_detail.fullName }}</span
                      >
                    </div>
                    <div class="c999">{{ record.guest_profile_detail.name }}</div>
                  </div>
                </div>
              </template>

              <template v-if="column.key === 'status'">
                <n-tag
                  :type="getOptionTag(optionsstatus.status, record?.status)"
                  size="small"
                  class="min-left-space"
                >
                  {{ getOptionLabel(optionsstatus.status, record?.status) }}
                </n-tag>
              </template>
              <template v-if="column.key === 'checkinDate'">
                {{ record.checkinDate }}&nbsp;{{ record.checkinTime }}
              </template>
              <template v-if="column.key === 'checkoutDate'">
                {{ record.checkoutDate }}&nbsp;{{ record.checkoutTime }}
              </template>
              <template v-if="column.key === 'checkinStatus'">
                <n-tag
                  :type="getOptionTag(optionsstatus.checkin_status, record?.checkinStatus)"
                  size="small"
                  class="min-left-space"
                >
                  {{ getOptionLabel(optionsstatus.checkin_status, record?.checkinStatus) }}
                </n-tag>
              </template>
              <template v-if="column.key === 'bookingFee'">
                <div> {{ record.bookingFee }}<span class="c999"> JPY</span> </div>
              </template>

              <template v-if="column.key === 'roommsg'">
                <div class="flex-row">
                  <img
                    :src="record.room_type_detail.cover"
                    style="width: 50px; height: 50px; object-fit: cover"
                  />
                  <div class="flex-item ml-3">
                    <div class="c333">{{ record.room_type_detail.name }}</div>
                    <div class="flex-row">
                      <img
                        src="@/assets/images/cr.png"
                        style="width: 16px; height: 16px; margin-right: 10px"
                      />
                      <div class="c333" style="margin-right: 5px">{{ record.adultCount }}</div>
                      <img
                        src="@/assets/images/ye.png"
                        style="width: 16px; height: 16px; margin-right: 10px"
                      />

                      <div class="c333">{{ record.childCount }}</div>
                    </div>
                  </div>
                </div>
              </template>
              <template v-if="column.key === 'guestRemarks'">
                <a-popover trigger="hover">
                  <template #content>
                    <div style="max-width: 250px; color: #333">{{ record.guestRemarks }}</div>
                  </template>
                  <div class="textline2 f12"> {{ record.guestRemarks }} </div>
                </a-popover>
              </template>

              <template v-if="column.key === 'work'">
                <div>
                  <n-button type="primary" size="small" @click="handleView(record)">
                    详情
                  </n-button>
                </div>
              </template>
            </template>
          </a-table>
        </div>
        <div class="text-r p-10">
          <a-pagination
            v-model:current="state.search.page"
            :total="state.totalCount"
            simple
            @change="changepage"
          />
        </div>

        <View ref="viewRef" />
      </div>
    </n-loading-bar-provider>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, computed, onMounted } from 'vue';
  import { List } from '@/api/pmsRoomReservation';
  import { formatToDate } from '@/utils/dateUtil';
  import { useLoadingBar } from 'naive-ui';
  import { Dicts } from '@/api/dict/dict';
  import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

  import View from '../pmsAppReservation/view.vue';
  const loadingBar = useLoadingBar();

  const viewRef = ref();
  const optionsstatus = ref({
    checkin_status: [] as Option[],
    order_status: [] as Option[],
  });
  async function loadOptions() {
    optionsstatus.value = await Dicts({
      types: ['checkin_status', 'status'],
    });
  }

  const state = reactive({
    btnloading: false,
    returnoptions: [
      {
        label: 'BAL',
        value: 'BAL',
      },
      {
        label: 'ALIPAY',
        value: 'ALIPAY',
      },
    ],
    returnAmount: {
      ordersn: '',
      refundType: '',
      refundAmount: null,
      maxAmount: 0,
    },
    totalCount: 0,
    search: {
      order_sn:'',
      out_order_sn: '',
      checkoutDate: '',
      checkinDate: '',
      status: '',
      checkinStatus: '',
      page: 1,
      pageSize: 10,
    },
    dataSource: [],
    columns: [
      {
        title: '订单信息',
        dataIndex: 'orderSn',
        key: 'orderSn',
      },
      {
        title: '订单状态',
        dataIndex: 'status',
        key: 'status',
      },

      {
        title: '房间信息',
        dataIndex: 'roommsg',
        key: 'roommsg',
      },

      {
        title: '入住',
        dataIndex: 'checkinDate',
        key: 'checkinDate',
      },

      {
        title: '退房',
        dataIndex: 'checkoutDate',
        key: 'checkoutDate',
      },
      {
        title: '入住状态',
        dataIndex: 'checkinStatus',
        key: 'checkinStatus',
      },
      {
        title: '账款',
        dataIndex: 'bookingFee',
        key: 'bookingFee',
      },

      {
        title: '评论',
        dataIndex: 'guestRemarks',
        key: 'guestRemarks',
        width: 250,
      },
      {
        title: '操作',
        dataIndex: 'work',
        key: 'work',
      },
    ],
  });

  // 加载表格数据
  const loadDataTable = async () => {
    loadingBar.start();
    let data = JSON.parse(JSON.stringify(state.search));
    if (data.checkinDate) {
      data.checkinDate = formatToDate(data.checkinDate, 'yyyy-MM-dd');
    }
    if (data.checkoutDate) {
      data.checkoutDate = formatToDate(data.checkoutDate, 'yyyy-MM-dd');
    }
    const res = await List(data);
    state.dataSource = res.list;
    state.totalCount = res.totalCount;
    loadingBar.finish();
  };

  function changepage(page, pageSize) {
    state.search.page = page;
    loadDataTable();
  }

  // 查看详情
  function handleView(record: Recordable) {
    record.viewtype = 'AIRHOST订单';
    viewRef.value.openModal(record);
  }

  onMounted(() => {
    loadOptions();
    loadDataTable();
  });
</script>

<style lang="less" scoped></style>
