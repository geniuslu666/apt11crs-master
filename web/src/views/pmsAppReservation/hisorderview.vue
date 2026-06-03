<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      style="width: 680px"
    >
      <n-loading-bar-provider v-if="showModal">
        <div class="flex-column" style="height: calc(100vh - 118px)">
          <n-card :bordered="false" title="历史记录">
            <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
            <div class="flex-row">
              <n-input
                v-model:value="state.search.orderSn"
                type="text"
                placeholder="订单号"
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
                placeholder="订单状态"
                clearable
                style="width: 180px"
              />
              <n-button type="primary" class="ml-1" @click="loadDataTable"> 搜索 </n-button>
            </div>
          </n-card>

          <div class="flex-item pl-5 bgfff">
            <div
              v-for="(record, index) in state.dataSource"
              :key="index"
              style="line-height: 2; margin: 10px 5px; border-bottom: 1px solid #eeeeeeff"
            >
              <div class="flex-row">
                <div class="c333 f16 flex-item">{{ record.propertyDetail.name }}</div>
                <div class="text-r">
                  <n-tag
                    :type="getOptionTag(optionsstatus.order_status, record?.orderStatus)"
                    size="small"
                    class="min-left-space"
                  >
                    {{ getOptionLabel(optionsstatus.order_status, record?.orderStatus) }}
                  </n-tag>
                </div>
              </div>

              <div class="c333 f12">
                <span class="mr-2"> {{ record.checkinDate }}-{{ record.checkoutDate }}</span>
                <span class="mr-2">{{ record.roomNum }}间</span>
                <span class="mr-2">{{ record.guestNum }}人</span>
              </div>
              <div class="c999 f12">
                <span>订单编号：</span> {{ record.orderSn }}&nbsp;&nbsp;<span class="fw csuccess">{{
                  record.outOrderSn
                }}</span>
              </div>
              <div class="flex-row">
                <div class="flex-item">
                  <div class="c999 f12"> <span>下单时间：</span>{{ record.createdAt }} </div>
                </div>
                <div class="text-r">
                  <span class="mr-2 f12">费用</span>
                  <span class="f16"> {{ record.orderAmount }}JPY</span>
                </div>
              </div>
            </div>
          </div>
          <div class="text-r p-10">
            <a-pagination
              v-model:current="state.search.page"
              :total="state.totalCount"
              simple
              @change="changepage"
            />
          </div>
        </div>
      </n-loading-bar-provider>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted } from 'vue';
  import { List } from '@/api/pmsAppReservation';

  import { Dicts } from '@/api/dict/dict';
  import { useLoadingBar } from 'naive-ui';
  import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { formatToDate } from '@/utils/dateUtil';
  const showModal = ref(false);

  const loadingBar = useLoadingBar();
  const optionsstatus = ref({
    checkin_status: [] as Option[],
    order_status: [] as Option[],
  });
  async function loadOptions() {
    optionsstatus.value = await Dicts({
      types: ['checkin_status', 'status', 'order_status','refund_status'],
    });
  }

  loadOptions();

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
      member_id: '',
      page: 1,
      pageSize: 10,
    },
    dataSource: [],
  });
  function openModal(val) {
    state.search.member_id = val;
    loadDataTable();
    showModal.value = true;
  }
  function changepage(page, pageSize) {
    state.search.page = page;
    loadDataTable();
  }

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

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped></style>
