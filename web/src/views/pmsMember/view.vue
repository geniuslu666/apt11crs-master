<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <div style="margin-bottom: 15px">
        <n-grid x-gap="15" y-gap="15" cols="3 600:4">
          <n-gi span="3" style="background: #FFF">
            <n-card :bordered="false" :header-style="{
              padding: '25px 20px 20px',
            }" :content-style="{
              padding: '0 20px 8px',
            }">
              <template #header>
                <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">基本信息</text>
              </template>
              <div class="member-info">
                <div class="headimg mr-4 w-12 h-12 md:w-[100px] md:h-[100px] md:mr-14">
                  <img v-if="data.avatar" :src="data.avatar" loading="lazy"
                    style="width: 100%; height: 100%; border-radius: 100%" :onerror="avatarErr" />
                  <img v-else src="~@/assets/images/mrtx.png" style="width: 100%; height: 100%; border-radius: 100%" />
                </div>
                <div class="info">
                  <div class="data-list">
                    <div class="data-list-item">
                      <div class="data-list-item-title">会员号：</div>
                      <div class="data-list-item-cont">{{ data.memberNo }}</div>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">会员姓：</div>
                      <div class="data-list-item-cont">{{ data.lastName }}</div>
                      <a class="data-list-item-edit" href="javascript:void(0);"
                        v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('firstName', data)">修改</a>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">会员分组：</div>
                      <div class="data-list-item-cont">{{ data.groupName }}</div>
                      <a class="data-list-item-edit" href="javascript:void(0);"
                        v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('group', data)">修改</a>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">会员名：</div>
                      <div class="data-list-item-cont">{{ data.firstName }}</div>
                      <a class="data-list-item-edit" href="javascript:void(0);"
                        v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('lastName', data)">修改</a>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">会员等级：</div>
                      <div class="data-list-item-cont">{{ data.levelName }}</div>
                      <a class="data-list-item-edit" href="javascript:void(0);"
                        v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('level', data)">修改</a>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">会员全名：</div>
                      <div class="data-list-item-cont">{{ data.fullName }}</div>
                    </div>
                  </div>
                  <div class="data-list">
                    <div class="data-list-item">
                      <div class="data-list-item-title">手机号码：</div>
                      <div class="data-list-item-cont">{{ data.phoneArea ? data.phoneArea + ' ' : '' }}{{ data.phone ?
                        data.phone :
                        '&#45;&#45;' }}</div>
                      <a class="data-list-item-edit" href="javascript:void(0);"
                        v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('phone', data)">修改</a>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">注册来源：</div>
                      <div class="data-list-item-cont">{{ data.source }}</div>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">电子邮箱：</div>
                      <div class="data-list-item-cont">{{ data.mail ? data.mail : '&#45;&#45;' }}</div>
                      <a class="data-list-item-edit" href="javascript:void(0);"
                        v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('mail', data)">修改</a>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">注册IP：</div>
                      <div class="data-list-item-cont">{{ data.registerIp }}</div>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">推荐人：</div>
                      <div class="data-list-item-cont">{{ data.referrerName }}</div>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title">注册时间：</div>
                      <div class="data-list-item-cont">{{ data.registerTime }}</div>
                    </div>
                  </div>
                  <div class="data-list">
                    <div class="data-list-item">
                      <div class="data-list-item-title" style="width: 106px">最后登录IP：</div>
                      <div class="data-list-item-cont">{{ data.lastLoginIp }}</div>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title" style="width: 106px">注册设备型号：</div>
                      <div class="data-list-item-cont">{{ data.registerMpModel ? data.registerMpModel : '--' }}</div>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title" style="width: 106px">最后登录时间：</div>
                      <div class="data-list-item-cont">{{ data.lastLogin }}</div>
                    </div>
                    <div class="data-list-item">
                      <div class="data-list-item-title" style="width: 106px">注册设备码：</div>
                      <div class="data-list-item-cont">{{ data.registerMdCode ? data.registerMdCode : '--' }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </n-card>
          </n-gi>
          <n-gi style="background: #FFF" span="3 600:1">
            <n-card :bordered="false" :header-style="{
              padding: '25px 20px 10px',
            }" :content-style="{
              padding: '0 20px 25px',
            }">
              <template #header>
                <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">账户信息</text>
              </template>
              <div class="member-info">
                <div class="info">
                  <div class="data-list-item1">
                    <div class="data-list-item-title1">会员积分：</div>
                    <div class="data-list-item-cont1">{{ data.balance }}</div>
                    <a class="data-list-item-edit1" href="javascript:void(0);"
                      v-if="hasPermission(['/pmsMember/balanceEdit'])" @click="handleEditBalance(data.balance)">修改</a>
                  </div>
                  <div class="data-list-item1">
                    <div class="data-list-item-title1">会员成长值：</div>
                    <div class="data-list-item-cont1">{{ data.exp }}</div>
                    <a class="data-list-item-edit1" href="javascript:void(0);"
                      v-if="hasPermission(['/pmsMember/expEdit'])" @click="handleEditExp(data.exp)">修改</a>
                  </div>
                </div>
              </div>
            </n-card>
          </n-gi>
        </n-grid>
      </div>
      <n-grid :cols="1">
        <n-gi>
          <n-card :bordered="false" :header-style="{
            padding: '25px 20px 20px',
          }" :content-style="{
            padding: '0 20px 25px',
          }">
            <template #header>
              <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">会员明细</text>
            </template>
            <div class="statusTab hs">
              <div :class="tabValue == 'BAL' ? 'active' : ''"><span @click="handleUpdateValue('BAL')">积分</span></div>
              <div :class="tabValue == 'EXP' ? 'active' : ''"><span @click="handleUpdateValue('EXP')">成长值</span></div>
              <div :class="tabValue == 'HOTEL' ? 'active' : ''"><span @click="handleUpdateValue('HOTEL')">住宿订单</span>
              </div>
              <div :class="tabValue == 'FOOD' ? 'active' : ''"><span @click="handleUpdateValue('FOOD')">餐厅订单</span>
              </div>
              <div :class="tabValue == 'SPA' ? 'active' : ''"><span @click="handleUpdateValue('SPA')">按摩订单</span></div>
              <div :class="tabValue == 'CAR' ? 'active' : ''"><span @click="handleUpdateValue('CAR')">接送机订单</span></div>
              <div :class="tabValue == 'CABINET' ? 'active' : ''"><span
                  @click="handleUpdateValue('CABINET')">储物柜订单</span>
              </div>
              <div :class="tabValue == 'TRAVEL' ? 'active' : ''"><span @click="handleUpdateValue('TRAVEL')">一日游订单</span></div>
              <div :class="tabValue == 'COUPON' ? 'active' : ''"><span @click="handleUpdateValue('COUPON')">优惠券</span>
              </div>
              <div :class="tabValue == 'TH_COUPON' ? 'active' : ''"><span
                  @click="handleUpdateValue('TH_COUPON')">礼品券</span>
              </div>
              <div :class="tabValue == 'JUNIOR' ? 'active' : ''"><span @click="handleUpdateValue('JUNIOR')">受邀下级</span>
              </div>
            </div>
            <template v-if="tabValue == 'BAL'">
              <BasicForm ref="searchBalanceFormRef" @register="registerBalance" @submit="reloadBalanceTable"
                @reset="reloadBalanceTable" @keyup.enter="reloadBalanceTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionBalanceRef" :columns="balanceColumns" :request="loadBalanceDataTable"
                :scroll-x="scrollBalanceX" :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'EXP'">
              <BasicForm ref="searchExpFormRef" @register="registerExp" @submit="reloadExpTable" @reset="reloadExpTable"
                @keyup.enter="reloadExpTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionExpRef" :columns="expColumns" :request="loadExpDataTable" :scroll-x="scrollExpX"
                :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'HOTEL'">
              <BasicForm ref="searchHotelOrderFormRef" @register="registerHotelOrder" @submit="reloadHotelOrderTable"
                @reset="reloadHotelOrderTable" @keyup.enter="reloadHotelOrderTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionHotelOrderRef" :columns="hotelOrderColumns" :actionColumn="hotelOrderActionColumn"
                :request="loadHotelOrderDataTable" :scroll-x="scrollHotelOrderX" :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'FOOD'">
              <BasicForm ref="searchFoodOrderFormRef" @register="registerFoodOrder" @submit="reloadFoodOrderTable"
                @reset="reloadFoodOrderTable" @keyup.enter="reloadFoodOrderTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionFoodOrderRef" :columns="foodOrderColumns" :actionColumn="foodOrderActionColumn"
                :request="loadFoodOrderDataTable" :scroll-x="scrollFoodOrderX" :resizeHeightOffset="-10000"
                @update:sorter="handleUpdateSorter">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'SPA'">
              <BasicForm ref="searchSpaOrderFormRef" @register="registerSpaOrder" @submit="reloadSpaOrderTable"
                @reset="reloadSpaOrderTable" @keyup.enter="reloadSpaOrderTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionSpaOrderRef" :columns="spaOrderColumns" :actionColumn="spaOrderActionColumn"
                :request="loadSpaOrderDataTable" :scroll-x="scrollSpaOrderX" :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'CAR'">
              <BasicForm ref="searchCarOrderFormRef" @register="registerCarOrder" @submit="reloadCarOrderTable"
                @reset="reloadCarOrderTable" @keyup.enter="reloadCarOrderTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionCarOrderRef" :columns="carOrderColumns" :actionColumn="carOrderActionColumn"
                :request="loadCarOrderDataTable" :scroll-x="scrollCarOrderX" :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'CABINET'">
              <BasicForm ref="searchCabinetOrderFormRef" @register="registerCabinetOrder"
                @submit="reloadCabinetOrderTable" @reset="reloadCabinetOrderTable"
                @keyup.enter="reloadCabinetOrderTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionCabinetOrderRef" :columns="cabinetOrderColumns"
                :actionColumn="cabinetOrderActionColumn" :request="loadCabinetOrderDataTable"
                :scroll-x="scrollCabinetOrderX" :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'TRAVEL'">
              <BasicForm ref="searchTravelOrderFormRef" @register="registerTravelOrder"
                         @submit="reloadTravelOrderTable" @reset="reloadTravelOrderTable"
                         @keyup.enter="reloadTravelOrderTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionTravelOrderRef" :columns="travelOrderColumns"
                          :actionColumn="travelOrderActionColumn" :request="loadTravelOrderDataTable"
                          :scroll-x="scrollTravelOrderX" :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'COUPON'">
              <BasicTable ref="actionCouponRef" :columns="couponColumns" :request="loadCouponDataTable"
                :scroll-x="scrollCouponX" :resizeHeightOffset="-10000" :actionColumn="actionCouponColumn">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'TH_COUPON'">
              <BasicForm ref="searchThCouponFormRef" @register="registerThCoupon" @submit="reloadThCouponTable"
                @reset="reloadThCouponTable" @keyup.enter="reloadThCouponTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionThCouponRef" :columns="thCouponColumns" :request="loadThCouponDataTable"
                :scroll-x="scrollThCouponX" :resizeHeightOffset="-10000" :actionColumn="thCouponActionColumn">
              </BasicTable>
            </template>
            <template v-else-if="tabValue == 'JUNIOR'">
              <BasicForm ref="searchJuniorFormRef" @register="registerJunior" @submit="reloadJuniorTable"
                @reset="reloadJuniorTable" @keyup.enter="reloadJuniorTable">
                <template #statusSlot="{ model, field }">
                  <n-input v-model:value="model[field]" />
                </template>
              </BasicForm>
              <BasicTable ref="actionJuniorRef" :columns="juniorColumns" :request="loadJuniorDataTable"
                :scroll-x="scrollJuniorX" :resizeHeightOffset="-10000">
              </BasicTable>
            </template>
          </n-card>
        </n-gi>
      </n-grid>
    </n-spin>

    <HotelOrderView ref="hotelOrderViewRef" />
    <FoodOrderView ref="foodOrderViewRef" />
    <CarOrderView ref="carOrderViewRef" />
    <SpaOrderView ref="spaOrderViewRef" />
    <CabinetOrderView ref="cabinetOrderViewRef" />
    <TravelOrderView ref="travelOrderViewRef" />
    <EditBase ref="editBaseRef" @reloadInfo="getInfo" />
    <EditBalance ref="editBalanceRef" @reloadInfo="reloadBalanceInfo" />
    <EditExp ref="editExpRef" @reloadInfo="reloadExpInfo" />
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, computed, reactive, h } from 'vue';
import { useRouter } from 'vue-router';
import { NButton, useMessage, useDialog } from 'naive-ui';
import HotelOrderView from '@/views/pmsAppReservation/view.vue';
import FoodOrderView from '@/views/foodOrder/view.vue';
import CarOrderView from '@/views/carOrder/view.vue';
import SpaOrderView from '@/views/spaOrder/view.vue';
import CabinetOrderView from '@/views/cabinetOrder/view.vue';
import TravelOrderView from '@/views/travelOrder/view.vue';
import { View } from '@/api/pmsMember';
import { getConfig } from "@/api/sys/config";
import { BasicForm, useForm } from "@/components/Form";
import { BasicTable, TableAction } from "@/components/Table";
import { balanceSchemas, loadBalanceOptions } from './view_balance_model';
import { expColumns, expSchemas, loadExpOptions } from './view_exp_model';
import { hotelOrderColumns, hotelOrderSchemas, loadOptions } from './view_hotel_order_model';
import { foodOrderColumns, foodOrderSchemas, loadOptions as loadFoodOptions } from './view_food_order_model';
import { spaOrderColumns, spaOrderSchemas, loadOptions as loadSpaOptions } from './view_spa_order_model';
import { carOrderColumns, carOrderSchemas, loadOptions as loadCarOptions } from './view_car_order_model';
import { cabinetOrderColumns, cabinetOrderSchemas, loadOptions as loadCabinetOptions } from './view_cabinet_order_model';
import { travelOrderColumns, travelOrderSchemas, loadOptions as loadTravelOptions } from './view_travel_order_model';
import { couponColumns, loadCouponOptions } from './view_coupon_model';
import { thCouponColumns, thCouponSchemas } from './view_th_coupon_model';
import { List as BalanceList } from "@/api/pmsBalanceChange";
import { List as ExpList } from "@/api/pmsExpChange";
import { List as HotelOrderList } from "@/api/pmsAppReservation";
import { List, List as FoodOrderList } from '@/api/foodOrder';
import { List as SpaOrderList } from '@/api/spaOrder';
import { List as CarOrderList } from '@/api/carOrder';
import { List as CabinetOrderList } from '@/api/cabinetOrder';
import { List as TravelOrderList } from '@/api/travelOrder';
import { adaTableScrollX } from "@/utils/hotgo";
import EditBase from '@/views/pmsMember/edit_base.vue';
import EditBalance from '@/views/pmsMember/edit_balance.vue';
import EditExp from '@/views/pmsMember/edit_exp.vue';
import { List as CouponList, Recycle as CouponRecycle } from "@/api/pmsCoupon";
import { List as MemberList } from "@/api/pmsMember";
import { List as ThCouponList, Recycle as ThCouponRecycle } from "@/api/thMemberCoupon";
import { juniorColumns, juniorSchemas, loadJuniorOptions } from './view_junior_model';
import defaultImg from "@/assets/images/mrtx.png";
import { usePermission } from "@/hooks/web/usePermission";
import { useSorter } from "@/hooks/common";

const dialog = useDialog();
const tabValue = ref('BAL')
const message = useMessage();
const { hasPermission } = usePermission();
const router = useRouter();
const params = router.currentRoute.value.params;
const loading = ref(false);
const recommendModel = ref('');
const data = ref({});
const editBaseRef = ref();
const editBalanceRef = ref();
const editExpRef = ref();
const actionCouponRef = ref();

const hotelOrderViewRef = ref();
const foodOrderViewRef = ref();
const carOrderViewRef = ref();
const spaOrderViewRef = ref();
const cabinetOrderViewRef = ref();
const travelOrderViewRef = ref();

const actionBalanceRef = ref();
const searchBalanceFormRef = ref<any>({});
const [registerBalance, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: balanceSchemas,
});
const scrollBalanceX = computed(() => {
  return adaTableScrollX(balanceColumns, 0);
});

const actionExpRef = ref();
const searchExpFormRef = ref<any>({});
const [registerExp, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: expSchemas,
});
const scrollExpX = computed(() => {
  return adaTableScrollX(expColumns, 0);
});
const scrollCouponX = computed(() => {
  return adaTableScrollX(couponColumns, actionCouponColumn.width);
});

const actionJuniorRef = ref();
const searchJuniorFormRef = ref<any>({});
const [registerJunior, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: juniorSchemas,
});
const scrollJuniorX = computed(() => {
  return adaTableScrollX(juniorColumns, 0);
});

const actionHotelOrderRef = ref();
const searchHotelOrderFormRef = ref<any>({});
const [registerHotelOrder, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: hotelOrderSchemas,
});
const scrollHotelOrderX = computed(() => {
  return adaTableScrollX(hotelOrderColumns, hotelOrderActionColumn.width);
});

const hotelOrderActionColumn = reactive({
  width: 160,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleHotelOrderView.bind(null, record),
        },
      ],
    });
  },
});


// 餐厅订单
const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadFoodOrderTable);
const actionFoodOrderRef = ref();
const searchFoodOrderFormRef = ref<any>({});
const [registerFoodOrder, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: foodOrderSchemas,
});
const scrollFoodOrderX = computed(() => {
  return adaTableScrollX(foodOrderColumns, foodOrderActionColumn.width);
});

const foodOrderActionColumn = reactive({
  width: 160,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleFoodOrderView.bind(null, record),
        },
      ],
    });
  },
});

// 按摩订单
const actionSpaOrderRef = ref();
const searchSpaOrderFormRef = ref<any>({});
const [registerSpaOrder, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: spaOrderSchemas,
});
const scrollSpaOrderX = computed(() => {
  return adaTableScrollX(spaOrderColumns, spaOrderActionColumn.width);
});

const spaOrderActionColumn = reactive({
  width: 160,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleSpaOrderView.bind(null, record),
        },
      ],
    });
  },
});

// 接送机订单
const actionCarOrderRef = ref();
const searchCarOrderFormRef = ref<any>({});
const [registerCarOrder, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: carOrderSchemas,
});
const scrollCarOrderX = computed(() => {
  return adaTableScrollX(carOrderColumns, carOrderActionColumn.width);
});

const carOrderActionColumn = reactive({
  width: 160,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleCarOrderView.bind(null, record),
        },
      ],
    });
  },
});

// 储物柜订单
const actionCabinetOrderRef = ref();
const searchCabinetOrderFormRef = ref<any>({});
const [registerCabinetOrder, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 90,
  schemas: cabinetOrderSchemas,
});
const scrollCabinetOrderX = computed(() => {
  return adaTableScrollX(cabinetOrderColumns, cabinetOrderActionColumn.width);
});

const cabinetOrderActionColumn = reactive({
  width: 160,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleCabinetOrderView.bind(null, record),
        },
      ],
    });
  },
});

// 一日游订单
const actionTravelOrderRef = ref();
const searchTravelOrderFormRef = ref<any>({});
const [registerTravelOrder, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 90,
  schemas: travelOrderSchemas,
});
const scrollTravelOrderX = computed(() => {
  return adaTableScrollX(travelOrderColumns, travelOrderActionColumn.width);
});

const travelOrderActionColumn = reactive({
  width: 160,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '详情',
          onClick: handleTravelOrderView.bind(null, record),
        },
      ],
    });
  },
});


// 优惠券
const actionCouponColumn = reactive({
  width: 90,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '回收',
          onClick: handleCouponRecycle.bind(null, record),
          ifShow: () => {
            return record.state === 1;
          },
          auth: ['/pmsCouponType/delete'],
        },
      ],

    });
  },
});

// 礼品券订单
const actionThCouponRef = ref();
const searchThCouponFormRef = ref<any>({});
const [registerThCoupon, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 90,
  schemas: thCouponSchemas,
});
const scrollThCouponX = computed(() => {
  return adaTableScrollX(thCouponColumns, thCouponActionColumn.width);
});

const thCouponActionColumn = reactive({
  width: 160,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '回收',
          onClick: handleThCouponRecycle.bind(null, record),
          ifShow: () => {
            return record.state === 1 || record.state === 2;
          },
          auth: ['/thMemberCoupon/recycle'],
        },
      ],
    });
  },
});

// 回收优惠券
function handleCouponRecycle(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '回收后用户无法再使用，确认回收？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      CouponRecycle({ id: record.id }).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          reloadCouponTable();
        });
      });
    },
  });
}

// 回收礼品券
function handleThCouponRecycle(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '回收后用户无法再使用，确认回收？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      ThCouponRecycle({ id: record.id }).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          reloadThCouponTable();
        });
      });
    },
  });
}

// 重新加载表格数据
function reloadCouponTable() {
  actionCouponRef.value?.reload();
}


const getInfo = () => {
  loading.value = true;
  View(params)
    .then((res) => {
      if ((res.referrer <= 0 && res.lastReferrer <= 0) || (!res.referrerDetail && !res.LastReferrerDetail)) {
        res.referrerName = '无';
      } else {
        if (recommendModel.value == 'FIRST') {
          if (!res.referrerDetail) {
            res.referrerName = '无';
          } else {
            if (res.referrerDetail.rebateMode == 'CHANNEL') {
              res.referrerName = '渠道(' + res.referrerDetail.memberNo + ')'
            } else if (res.referrerDetail.rebateMode == 'STAFF') {
              res.referrerName = '员工(' + res.referrerDetail.memberNo + ')'
            } else {
              res.referrerName = '会员(' + res.referrerDetail.memberNo + ')'
            }
          }
        } else {
          if (!res.LastReferrerDetail) {
            res.referrerName = '无';
          } else {
            if (res.LastReferrerDetail.rebateMode == 'CHANNEL') {
              res.referrerName = '渠道(' + res.LastReferrerDetail.memberNo + ')'
            } else if (res.LastReferrerDetail.rebateMode == 'STAFF') {
              res.referrerName = '员工(' + res.LastReferrerDetail.memberNo + ')'
            } else {
              res.referrerName = '会员(' + res.LastReferrerDetail.memberNo + ')'
            }
          }
        }
      }

      res.levelName = res.memberLevel.levelName
      res.groupName = res.memberGroup ? res.memberGroup.memberGroup : '--'
      data.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
};

const balanceColumns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 100,
  },
  {
    title: '方向',
    key: 'changePrice',
    align: 'left',
    width: 80,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.changePrice > 0 ? 'green' : 'red'
          },
        },
        {
          default: () => row.changePrice > 0 ? '发放' : '消耗'
        }
      )
    },
  },
  {
    title: '发生场景',
    key: 'scene',
    align: 'left',
    width: 80,
    render(row) {
      if (row.scene == 'SYSTEM') {
        return '系统'
      } else if (row.scene == 'HOTEL') {
        return '酒店'
      } else if (row.scene == 'FOOD') {
        return '餐饮'
      } else if (row.scene == 'SPA') {
        return '按摩'
      } else if (row.scene == 'CAR') {
        return '出行'
      } else if (row.scene == 'CABINET') {
        return '储物柜'
      } else if (row.scene == 'TRAVRL') {
        return '一日游'
      } else {
        return '未知'
      }
    },
  },
  {
    title: '积分变化',
    key: 'changePrice',
    align: 'left',
    width: 200,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.changePrice > 0 ? 'green' : 'red'
          },
        },
        {
          default: () => row.changePrice > 0 ? '+' + row.changePrice : row.changePrice
        }
      )
    },
  },
  {
    title: '操作人',
    key: 'des',
    align: 'left',
    width: 100,
    render(row) {
      if (row.operatorId > 0) {
        return row.adminMemberUsername
      } else {
        return '系统操作'
      }
    }
  },
  {
    title: '备注',
    key: 'des',
    align: 'left',
    width: 200,
    ellipsis: false,
    render(row) {
      if (row.des) {
        return row.des
      } else {
        return '--'
      }
    }
  },
  {
    title: '原因',
    key: 'reason',
    align: 'left',
    width: 200,
    ellipsis: false,
    render(row) {
      if (row.reason) {
        return row.reason
      } else {
        return '--'
      }
    }
  },
  {
    title: '发生时间',
    key: 'createdAt',
    align: 'left',
    width: 200,
  },
  {
    title: '关联订单号',
    key: 'orderSn',
    align: 'left',
    width: 240,
    render(row) {
      if (row.orderSn) {
        if (row.orderSn.charAt(0) == 'H') {
          return h(
            NButton,
            {
              strong: true,
              size: 'small',
              text: true,
              type: 'primary',
              onClick: () => handleHotelOrderView({ orderSn: row.orderSn, outOrderSn: '' }),
            },
            { default: () => row.orderSn }
          )
        } else if (row.orderSn.charAt(0) == 'F') {
          return h(
            NButton,
            {
              strong: true,
              size: 'small',
              text: true,
              type: 'primary',
              onClick: () => handleFoodOrderView({ orderSn: row.orderSn }),
            },
            { default: () => row.orderSn }
          )
        } else if (row.orderSn.charAt(0) == 'S') {
          return h(
            NButton,
            {
              strong: true,
              size: 'small',
              text: true,
              type: 'primary',
              onClick: () => handleSpaOrderView({ orderSn: row.orderSn }),
            },
            { default: () => row.orderSn }
          )
        } else if (row.orderSn.charAt(0) == 'C') {
          return h(
            NButton,
            {
              strong: true,
              size: 'small',
              text: true,
              type: 'primary',
              onClick: () => handleCarOrderView({ orderSn: row.orderSn }),
            },
            { default: () => row.orderSn }
          )
        } else if (row.orderSn.charAt(0) == 'B') {
          return h(
            NButton,
            {
              strong: true,
              size: 'small',
              text: true,
              type: 'primary',
              onClick: () => handleCabinetOrderView({ orderSn: row.orderSn }),
            },
            { default: () => row.orderSn }
          )
        }

      } else {
        return ''
      }
    }
  },
];

function avatarErr(e) {
  e.target.src = defaultImg
}

function handleUpdateValue(e) {
  tabValue.value = e;
}

// 加载表格数据
const loadBalanceDataTable = async (res) => {
  res.memberId = params.id
  return await BalanceList({ ...searchBalanceFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadBalanceTable() {
  actionBalanceRef.value?.reload();
}

// 加载表格数据
const loadExpDataTable = async (res) => {
  res.memberId = params.id
  return await ExpList({ ...searchExpFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadExpTable() {
  actionExpRef.value?.reload();
}

// 加载表格数据
const loadHotelOrderDataTable = async (res) => {
  res.member_id = params.id
  return await HotelOrderList({ ...searchHotelOrderFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadHotelOrderTable() {
  actionHotelOrderRef.value?.reload();
}

// 加载表格数据
const loadFoodOrderDataTable = async (res) => {
  res.memberId = params.id

  if (sortStatesRef.value.length > 0) {
    var bookSort = "";
    var createSort = "";
    sortStatesRef.value.forEach(element => {
      if (element.columnKey == 'bookDateTime') {
        bookSort = element.order
      } else if (element.columnKey == 'createdAt') {
        createSort = element.order
      }

    });
    return await FoodOrderList({ ...searchFoodOrderFormRef.value?.formModel, ...res, bookSort: bookSort, createSort: createSort });
  } else {
    return await FoodOrderList({ ...searchFoodOrderFormRef.value?.formModel, ...res });
  }


};

// 重新加载表格数据
function reloadFoodOrderTable() {
  actionFoodOrderRef.value?.reload();
}

// 加载表格数据
const loadSpaOrderDataTable = async (res) => {
  res.memberId = params.id
  return await SpaOrderList({ ...searchSpaOrderFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadSpaOrderTable() {
  actionSpaOrderRef.value?.reload();
}

// 加载表格数据
const loadCarOrderDataTable = async (res) => {
  res.memberId = params.id
  return await CarOrderList({ ...searchCarOrderFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadCarOrderTable() {
  actionCarOrderRef.value?.reload();
}

// 储物柜~~~
// 加载表格数据
const loadCabinetOrderDataTable = async (res) => {
  res.memberId = params.id
  return await CabinetOrderList({ ...searchCabinetOrderFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadCabinetOrderTable() {
  actionCabinetOrderRef.value?.reload();
}

// 一日游~~~
// 加载表格数据
const loadTravelOrderDataTable = async (res) => {
  res.memberId = params.id
  return await TravelOrderList({ ...searchTravelOrderFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadTravelOrderTable() {
  actionTravelOrderRef.value?.reload();
}
// 一日游~~~end~~~

// 加载表格数据
const loadThCouponDataTable = async (res) => {
  res.memberId = params.id
  return await ThCouponList({ ...searchThCouponFormRef.value?.formModel, ...res });
};

// 重新加载表格数据
function reloadThCouponTable() {
  actionThCouponRef.value?.reload();
}

// 重新加载表格数据
function reloadJuniorTable() {
  actionJuniorRef.value?.reload();
}

// 加载表格数据
const loadCouponDataTable = async (res) => {
  res.memberId = params.id
  return await CouponList({ ...res });
};

// 加载表格数据
const loadJuniorDataTable = async (res) => {
  res.id = params.id
  res.isJunior = 1
  return await MemberList({ ...searchJuniorFormRef.value?.formModel, ...res });
};

// // 重新加载表格数据
// function reloadCouponTable() {
//   actionCouponRef.value?.reload();
// }

// 查看详情
function handleHotelOrderView(record: Recordable) {
  hotelOrderViewRef.value.openModal(record);
}

// 查看详情
function handleFoodOrderView(record: Recordable) {
  foodOrderViewRef.value.openModal(record);
}

// 查看详情
function handleCarOrderView(record: Recordable) {
  carOrderViewRef.value.openModal(record);
}

// 查看详情
function handleCabinetOrderView(record: Recordable) {
  cabinetOrderViewRef.value.openModal(record);
}

// 查看详情
function handleTravelOrderView(record: Recordable) {
  travelOrderViewRef.value.openModal(record);
}

// 查看详情
function handleSpaOrderView(record: Recordable) {
  spaOrderViewRef.value.openModal(record);
}

// 编辑数据
function handleEditBase(type, value) {
  editBaseRef.value.openModal(data.value.id, type, value);
}
// 编辑数据
function handleEditBalance(value) {
  editBalanceRef.value.openModal(data.value.id, value);
}
// 编辑数据
function handleEditExp(value) {
  editExpRef.value.openModal(data.value.id, value);
}

function reloadBalanceInfo() {
  getInfo()
  reloadBalanceTable()
}

function reloadExpInfo() {
  getInfo()
  reloadExpTable()
}

onMounted(() => {
  if (!params.id) {
    message.error('会员ID不正确，请检查！');
    return;
  }
  getConfig({ group: 'yyconfig' })
    .then((res) => {
      recommendModel.value = res.list.recommendModel;
    })
  loadBalanceOptions();
  loadExpOptions();
  loadFoodOptions();
  loadSpaOptions();
  loadCarOptions();
  loadCabinetOptions();
  loadTravelOptions();
  loadJuniorOptions();
  loadOptions();
  loadCouponOptions();
  getInfo();
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}

.member-info {
  display: flex;

  .headimg {
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
  }

  .info {
    flex: 1;
    width: 0;
    padding-top: 10px;

    .data-list {
      width: 100%;
      border-bottom: 1px solid #EEEEEE;
      padding-bottom: 8px;
      margin-bottom: 20px;
      display: flex;
      flex-wrap: wrap;

      &:last-child {
        border-bottom: none;
        margin-bottom: 0;
      }

      .data-list-item {
        width: 50%;
        box-sizing: border-box;
        display: flex;
        align-items: center;
        margin-bottom: 12px;

        .data-list-item-title {
          width: 78px;
          font-weight: 400;
          font-size: 14px;
          color: #4E5969;
          line-height: 22px;
        }

        .data-list-item-cont {
          font-weight: 400;
          font-size: 14px;
          color: #1D2129;
          line-height: 22px;
        }

        .data-list-item-edit {
          font-weight: 400;
          font-size: 14px;
          color: #1577FF;
          line-height: 22px;
          margin-left: 14px;
        }
      }

      @media(max-width: 600px) {
        .data-list-item {
          width: 100%;
        }
      }
    }

    .data-list-item1 {
      display: flex;
      align-items: center;
      margin-bottom: 12px;

      .data-list-item-title1 {
        width: 92px;
        font-weight: 400;
        font-size: 14px;
        color: #4E5969;
        line-height: 22px;
      }

      .data-list-item-cont1 {
        font-weight: 400;
        font-size: 14px;
        color: #1D2129;
        line-height: 22px;
      }

      .data-list-item-edit1 {
        font-weight: 400;
        font-size: 14px;
        color: #1577FF;
        line-height: 22px;
        margin-left: 14px;
      }
    }
  }

}

.statusTab {
  display: flex;
  margin-bottom: 30px;
  flex-wrap: nowrap;
  overflow-x: auto;

  div {
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    white-space: nowrap;

    span {
      cursor: pointer;
    }

    &.active {
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
