<template>
  <n-loading-bar-provider>
    <div class="flex-column" style="height: calc(100vh - 118px)">
      <n-card :bordered="false" title="App入住订单">
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
        <div class="flex-row">
          <n-input
            v-model:value="state.search.orderSn"
            type="text"
            placeholder="订单号"
            style="width: 180px"
            @keyup.enter="loadDataTable"
          />
          <n-input
            v-model:value="state.search.outOrderSn"
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
            v-model:value="state.search.orderStatus"
            :options="optionsstatus.order_status"
            clearable
            style="width: 180px"
          />
          <n-button type="primary" class="ml-1" @click="loadDataTable"> 搜索 </n-button>
          <!-- <a-button
            type="primary"
            class="ml-1"
            :loading="state.btnloading"
            @click="syncAllOrderLoad"
          >
            同步订单
          </a-button> -->
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
              <div class="flex-row" v-if="record.propertyDetail">
                <img :src="record.propertyDetail.cover" style="width: 50px; height: 65px" />
                <div class="flex-item" style="padding-left: 10px">
                  <div
                    >{{ record.orderSn }}&nbsp;&nbsp;<span class="fw csuccess">{{
                      record.outOrderSn
                    }}</span></div
                  >
                  <div v-if="record.guestProfileDetail">
                    <span v-if="record.guestProfileDetail.fullName" class="cblue fw f14">
                      {{ record.guestProfileDetail.fullName }}</span
                    >
                    <span class="c333" style="margin-left: 10px"
                      >tel:{{ record.guestProfileDetail.phone }}</span
                    >
                  </div>
                  <div class="c999">{{ record.propertyDetail.name }}</div>
                </div>
              </div>
            </template>
            <template v-if="column.key === 'status'">
              <div>
                <n-tag
                  :type="getOptionTag(optionsstatus.order_status, record?.orderStatus)"
                  size="small"
                  class="min-left-space"
                >
                  {{ getOptionLabel(optionsstatus.order_status, record?.orderStatus) }}
                </n-tag>
              </div>
            </template>
            <template v-if="column.key === 'checkInDate'">
              {{ record.checkinDate }}&nbsp;{{ record.checkinTime }}
            </template>
            <template v-if="column.key === 'checkOutDate'">
              {{ record.checkoutDate }}&nbsp;{{ record.checkoutTime }}
            </template>
            <template v-if="column.key === 'orderAmount'">
              <div> {{ record.orderAmount }}<span class="c999"> JPY</span> </div>
            </template>

            <template v-if="column.key === 'roommsg'">
              <div class="flex-row">
                <div class="c999">房间总数：</div>
                <div class="flex-item">{{ record.roomNum }}</div>
              </div>

              <div class="flex-row">
                <div class="c999">入住人数：</div>
                <div class="flex-item">{{ record.guestNum }}</div>
              </div>
            </template>
            <template v-if="column.key === 'work'">
              <div>
                <n-button type="primary" size="small" @click="handleView(record)"> 详情 </n-button>
                <n-button
                  v-if="record.isRefund"
                  size="small"
                  class="ml-1"
                  type="warning"
                  @click="openreturn(record)"
                >
                  退款
                </n-button>
                <!-- <n-button  v-if="record.orderStatus !== 'CANCEL'" @click="syncOrderLoad(record)" class="ml-1" type="success" size="small"
                  >同步</n-button
                > -->
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

      <Edit ref="editRef" @reloadTable="reloadTable" />
      <View ref="viewRef" />
      <n-modal
        v-model:show="showreturn"
        :mask-closable="false"
        preset="dialog"
        title="退款"
        positive-text="提交"
        negative-text="算了"
        @positive-click="onPositiveClick"
        @negative-click="onNegativeClick"
      >
        <div style="width: 500px" v-if="showreturn">
          <div class="flex-row">
            <n-select
              v-model:value="state.returnAmount.refundType"
              :options="state.returnoptions"
              style="width: 150px; margin-right: 10px"
              @update:value="selectchange"
            >
            </n-select>
            <n-input-number
              v-model:value="state.returnAmount.refundAmount"
              style="width: 220px"
              :max="state.returnAmount.maxAmount"
              :placeholder="'最大退款金额' + state.returnAmount.maxAmount"
            ></n-input-number>
          </div>
        </div>
      </n-modal>
    </div>
  </n-loading-bar-provider>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { List, Delete, paycloud, syncAllOrder, syncOrder } from '@/api/pmsAppReservation';
  import Edit from './edit.vue';
  import View from './view.vue';
  import { Dicts } from '@/api/dict/dict';
  import { useLoadingBar } from 'naive-ui';
  import { PersonOutline, Accessibility } from '@vicons/ionicons5';
  import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { formatToDate } from '@/utils/dateUtil';

  const dialog = useDialog();
  const message = useMessage();
  const actionRef = ref();
  const editRef = ref();
  const viewRef = ref();
  const showreturn = ref(false);
  const loadingBar = useLoadingBar();
  const optionsstatus = ref({
    checkin_status: [] as Option[],
    order_status: [] as Option[],
  });
  async function loadOptions() {
    optionsstatus.value = await Dicts({
      types: ['checkin_status', 'order_status'],
    });
  }

  loadOptions();
  function selectchange(key, options) {
    state.returnAmount.maxAmount = options.maxAmount;
    console.log(key, options);
  }
  const syncOrderLoad = async (row) => {
    loadingBar.start();
    await syncOrder({ orderSn: row.orderSn, outOrderSn: row.outOrderSn });

    message.success('已同步');

    loadingBar.finish();
  };
  const syncAllOrderLoad = async () => {
    loadingBar.start();
    await syncAllOrder({});

    message.success('已同步');

    loadingBar.finish();
  };
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
      orderSn: '',
      outOrderSn: '',
      checkoutDate: '',
      checkinDate: '',
      orderStatus: '',
      page: 1,
      pageSize: 10,
    },
    dataSource: [],
    columns: [
      {
        title: '订单信息',
        dataIndex: 'orderSn',
        key: 'orderSn',
        width: 530
      },
      {
        title: '订单状态',
        dataIndex: 'status',
        key: 'status',
      },

      {
        title: '人员信息',
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
        title: '费用',
        dataIndex: 'orderAmount',
        key: 'orderAmount',
      },
      {
        title: '操作',
        dataIndex: 'work',
        key: 'work',
      },
    ],
  });
  const onNegativeClick = () => {
    showreturn.value = false;
  };

  const openreturn = (row) => {
    state.returnAmount.ordersn = row.orderSn;
    state.returnoptions = row.transactionDetail.map((m) => {
      return {
        label: m.payType,
        value: m.payType,
        maxAmount: m.payAmount,
      };
    });

    showreturn.value = true;
  };
  const onPositiveClick = async () => {
    await paycloud(state.returnAmount);

    message.success('已提交');
    showreturn.value = false;
    loadDataTable();
  };
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

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 查看详情
  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
  }
  function changepage(page, pageSize) {
    state.search.page = page;
    loadDataTable();
  }

  onMounted(() => {
    loadDataTable();
  });
</script>

<style lang="less" scoped></style>
