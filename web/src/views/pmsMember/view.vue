<template>
  <div class="member-admin-page member-detail-page p-4 space-y-4">
    <!-- Loading overlay -->
    <div v-if="loading" class="flex items-center justify-center py-20">
      <svg class="animate-spin w-8 h-8 text-primary" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
      </svg>
    </div>

    <template v-if="!loading">
      <!-- Basic Info + Account Info -->
      <div class="member-detail-grid grid grid-cols-1 gap-4">
        <!-- Basic Info -->
        <div class="member-card bg-white rounded-lg border border-border p-5">
          <h2 class="member-section-title">基本信息</h2>
          <div class="flex gap-4 md:gap-8">
            <!-- Avatar -->
            <div class="flex-shrink-0 w-12 h-12 md:w-24 md:h-24 rounded-full overflow-hidden bg-muted">
              <img v-if="data.avatar" :src="data.avatar" loading="lazy" class="w-full h-full object-cover"
                @error="($event.target as HTMLImageElement).src = defaultImg" />
              <img v-else :src="defaultImg" class="w-full h-full object-cover" />
            </div>
            <!-- Info lists -->
            <div class="flex-1 min-w-0">
              <!-- Row 1 -->
              <div class="member-info-grid grid grid-cols-1 sm:grid-cols-2 gap-x-8 gap-y-2 pb-3 mb-3 border-b border-border">
                <InfoItem label="会员号" :value="data.memberNo" />
                <InfoItem label="会员姓">
                  {{ data.lastName }}
                  <button v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('firstName', data)" class="member-inline-action">修改</button>
                </InfoItem>
                <InfoItem label="会员分组">
                  {{ data.groupName }}
                  <button v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('group', data)" class="member-inline-action">修改</button>
                </InfoItem>
                <InfoItem label="会员名">
                  {{ data.firstName }}
                  <button v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('lastName', data)" class="member-inline-action">修改</button>
                </InfoItem>
                <InfoItem label="会员等级">
                  {{ data.levelName }}
                  <button v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('level', data)" class="member-inline-action">修改</button>
                </InfoItem>
                <InfoItem label="会员全名" :value="data.fullName" />
              </div>
              <!-- Row 2 -->
              <div class="member-info-grid grid grid-cols-1 sm:grid-cols-2 gap-x-8 gap-y-2 pb-3 mb-3 border-b border-border">
                <InfoItem label="手机号码">
                  {{ data.phoneArea ? data.phoneArea + ' ' : '' }}{{ data.phone || '--' }}
                  <button v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('phone', data)" class="member-inline-action">修改</button>
                </InfoItem>
                <InfoItem label="注册来源" :value="data.source" />
                <InfoItem label="电子邮箱">
                  {{ data.mail || '--' }}
                  <button v-if="hasPermission(['/pmsMember/baseEdit'])" @click="handleEditBase('mail', data)" class="member-inline-action">修改</button>
                </InfoItem>
                <InfoItem label="注册IP" :value="data.registerIp" />
                <InfoItem label="推荐人" :value="data.referrerName" />
                <InfoItem label="注册时间" :value="data.registerTime" />
              </div>
              <!-- Row 3 -->
              <div class="member-info-grid grid grid-cols-1 sm:grid-cols-2 gap-x-8 gap-y-2">
                <InfoItem label="最后登录IP" :value="data.lastLoginIp" />
                <InfoItem label="注册设备型号" :value="data.registerMpModel || '--'" />
                <InfoItem label="最后登录时间" :value="data.lastLogin" />
                <InfoItem label="注册设备码" :value="data.registerMdCode || '--'" />
              </div>
            </div>
          </div>
        </div>

        <!-- Account Info -->
        <div class="member-card bg-white rounded-lg border border-border p-5 min-w-[220px]">
          <h2 class="member-section-title">账户信息</h2>
          <div class="space-y-2.5">
            <div class="member-account-row">
              <span>会员积分</span>
              <strong>{{ data.balance }}</strong>
              <button v-if="hasPermission(['/pmsMember/balanceEdit'])" @click="handleEditBalance(data.balance)" class="member-inline-action">修改</button>
            </div>
            <div class="member-account-row">
              <span>会员成长值</span>
              <strong>{{ data.exp }}</strong>
              <button v-if="hasPermission(['/pmsMember/expEdit'])" @click="handleEditExp(data.exp)" class="member-inline-action">修改</button>
            </div>
          </div>
        </div>
      </div>

      <!-- Detail Tabs -->
      <div class="member-card bg-white rounded-lg border border-border p-5">
        <h2 class="member-section-title">会员明细</h2>

        <!-- Tab Bar -->
        <div class="flex items-center gap-1 flex-wrap mb-4">
          <button v-for="tab in tabs" :key="tab.key" type="button"
            @click="tabValue = tab.key"
            :class="['px-3 py-1.5 text-sm rounded-full transition-colors whitespace-nowrap', tabValue === tab.key ? 'bg-muted text-primary font-medium' : 'text-muted-foreground hover:text-foreground hover:bg-muted/50']"
          >{{ tab.label }}</button>
        </div>

        <!-- BAL -->
        <template v-if="tabValue === 'BAL'">
          <ProTable
            ref="actionBalanceRef"
            :columns="balanceCols"
            :request="loadBalanceDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <span v-if="row.orderSn">
                <button @click="openOrderView(row.orderSn)" class="member-table-action">{{ row.orderSn }}</button>
              </span>
            </template>
          </ProTable>
        </template>

        <!-- EXP -->
        <template v-else-if="tabValue === 'EXP'">
          <ProTable
            ref="actionExpRef"
            :columns="expCols"
            :request="loadExpDataTable"
            :row-key="(row) => row.id"
          />
        </template>

        <!-- HOTEL -->
        <template v-else-if="tabValue === 'HOTEL'">
          <ProTable
            ref="actionHotelOrderRef"
            :columns="hotelOrderCols"
            :request="loadHotelOrderDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button @click="handleHotelOrderView(row)" class="member-table-action">详情</button>
            </template>
          </ProTable>
        </template>

        <!-- FOOD -->
        <template v-else-if="tabValue === 'FOOD'">
          <ProTable
            ref="actionFoodOrderRef"
            :columns="foodOrderCols"
            :request="loadFoodOrderDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button @click="handleFoodOrderView(row)" class="member-table-action">详情</button>
            </template>
          </ProTable>
        </template>

        <!-- SPA -->
        <template v-else-if="tabValue === 'SPA'">
          <ProTable
            ref="actionSpaOrderRef"
            :columns="spaOrderCols"
            :request="loadSpaOrderDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button @click="handleSpaOrderView(row)" class="member-table-action">详情</button>
            </template>
          </ProTable>
        </template>

        <!-- CAR -->
        <template v-else-if="tabValue === 'CAR'">
          <ProTable
            ref="actionCarOrderRef"
            :columns="carOrderCols"
            :request="loadCarOrderDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button @click="handleCarOrderView(row)" class="member-table-action">详情</button>
            </template>
          </ProTable>
        </template>

        <!-- CABINET -->
        <template v-else-if="tabValue === 'CABINET'">
          <ProTable
            ref="actionCabinetOrderRef"
            :columns="cabinetOrderCols"
            :request="loadCabinetOrderDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button @click="handleCabinetOrderView(row)" class="member-table-action">详情</button>
            </template>
          </ProTable>
        </template>

        <!-- TRAVEL -->
        <template v-else-if="tabValue === 'TRAVEL'">
          <ProTable
            ref="actionTravelOrderRef"
            :columns="travelOrderCols"
            :request="loadTravelOrderDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button @click="handleTravelOrderView(row)" class="member-table-action">详情</button>
            </template>
          </ProTable>
        </template>

        <!-- COUPON -->
        <template v-else-if="tabValue === 'COUPON'">
          <ProTable
            ref="actionCouponRef"
            :columns="couponCols"
            :request="loadCouponDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button
                v-if="row.state === 1 && hasPermission(['/pmsCouponType/delete'])"
                @click="handleCouponRecycle(row)"
                class="member-table-action member-table-action-danger"
              >回收</button>
            </template>
          </ProTable>
        </template>

        <!-- TH_COUPON -->
        <template v-else-if="tabValue === 'TH_COUPON'">
          <ProTable
            ref="actionThCouponRef"
            :columns="thCouponCols"
            :request="loadThCouponDataTable"
            :row-key="(row) => row.id"
          >
            <template #action="{ row }">
              <button
                v-if="(row.state === 1 || row.state === 2) && hasPermission(['/thMemberCoupon/recycle'])"
                @click="handleThCouponRecycle(row)"
                class="member-table-action member-table-action-danger"
              >回收</button>
            </template>
          </ProTable>
        </template>

        <!-- JUNIOR -->
        <template v-else-if="tabValue === 'JUNIOR'">
          <ProTable
            ref="actionJuniorRef"
            :columns="juniorCols"
            :request="loadJuniorDataTable"
            :row-key="(row) => row.id"
          />
        </template>
      </div>
    </template>

    <!-- Confirm Dialog -->
    <UiDialog :open="confirmDialog.show" :title="confirmDialog.title" max-width="400px" @update:open="confirmDialog.show = false">
      <p class="text-sm text-foreground">{{ confirmDialog.content }}</p>
      <template #footer>
        <UiButton variant="outline" @click="confirmDialog.show = false">取消</UiButton>
        <UiButton variant="destructive" @click="confirmDialog.onConfirm?.()">确定</UiButton>
      </template>
    </UiDialog>

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
import { onMounted, ref, h, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useMessage } from 'naive-ui';
import { UiButton, UiDialog, ProTable } from '@/components/ui';
import HotelOrderView from '@/views/pmsAppReservation/view.vue';
import FoodOrderView from '@/views/foodOrder/view.vue';
import CarOrderView from '@/views/carOrder/view.vue';
import SpaOrderView from '@/views/spaOrder/view.vue';
import CabinetOrderView from '@/views/cabinetOrder/view.vue';
import TravelOrderView from '@/views/travelOrder/view.vue';
import { View } from '@/api/pmsMember';
import { getConfig } from "@/api/sys/config";
import { List as BalanceList } from "@/api/pmsBalanceChange";
import { List as ExpList } from "@/api/pmsExpChange";
import { List as HotelOrderList } from "@/api/pmsAppReservation";
import { List as FoodOrderList } from '@/api/foodOrder';
import { List as SpaOrderList } from '@/api/spaOrder';
import { List as CarOrderList } from '@/api/carOrder';
import { List as CabinetOrderList } from '@/api/cabinetOrder';
import { List as TravelOrderList } from '@/api/travelOrder';
import { List as CouponList, Recycle as CouponRecycle } from "@/api/pmsCoupon";
import { List as MemberList } from "@/api/pmsMember";
import { List as ThCouponList, Recycle as ThCouponRecycle } from "@/api/thMemberCoupon";
import EditBase from '@/views/pmsMember/edit_base.vue';
import EditBalance from '@/views/pmsMember/edit_balance.vue';
import EditExp from '@/views/pmsMember/edit_exp.vue';
import { usePermission } from "@/hooks/web/usePermission";
import defaultImg from "@/assets/images/mrtx.png";

// Inline InfoItem component
const InfoItem = {
  props: ['label', 'value'],
  setup(props: any, { slots }: any) {
    return () => h('div', { class: 'member-info-item' }, [
      h('span', { class: 'member-info-label' }, props.label),
      slots.default ? h('span', { class: 'member-info-value' }, slots.default()) : h('span', { class: 'member-info-value' }, props.value ?? '--'),
    ]);
  },
};

const { hasPermission } = usePermission();
const router = useRouter();
const params = router.currentRoute.value.params;
const message = useMessage();
const loading = ref(false);
const data = ref<any>({});
const tabValue = ref('BAL');
const recommendModel = ref('');

const editBaseRef = ref();
const editBalanceRef = ref();
const editExpRef = ref();

const hotelOrderViewRef = ref();
const foodOrderViewRef = ref();
const carOrderViewRef = ref();
const spaOrderViewRef = ref();
const cabinetOrderViewRef = ref();
const travelOrderViewRef = ref();

const actionBalanceRef = ref();
const actionExpRef = ref();
const actionHotelOrderRef = ref();
const actionFoodOrderRef = ref();
const actionSpaOrderRef = ref();
const actionCarOrderRef = ref();
const actionCabinetOrderRef = ref();
const actionTravelOrderRef = ref();
const actionCouponRef = ref();
const actionThCouponRef = ref();
const actionJuniorRef = ref();

const confirmDialog = reactive({
  show: false,
  title: '',
  content: '',
  onConfirm: null as (() => void) | null,
});

function showConfirm(title: string, content: string, onConfirm: () => void) {
  confirmDialog.title = title;
  confirmDialog.content = content;
  confirmDialog.onConfirm = () => {
    onConfirm();
    confirmDialog.show = false;
  };
  confirmDialog.show = true;
}

const tabs = [
  { key: 'BAL', label: '积分' },
  { key: 'EXP', label: '成长值' },
  { key: 'HOTEL', label: '住宿订单' },
  { key: 'FOOD', label: '餐厅订单' },
  { key: 'SPA', label: '按摩订单' },
  { key: 'CAR', label: '接送机订单' },
  { key: 'CABINET', label: '储物柜订单' },
  { key: 'TRAVEL', label: '一日游订单' },
  { key: 'COUPON', label: '优惠券' },
  { key: 'TH_COUPON', label: '礼品券' },
  { key: 'JUNIOR', label: '受邀下级' },
];

// Balance columns
const sceneLabel = (scene: string) => {
  const map: Record<string, string> = { SYSTEM: '系统', HOTEL: '酒店', FOOD: '餐饮', SPA: '按摩', CAR: '出行', CABINET: '储物柜', TRAVRL: '一日游' };
  return map[scene] || '未知';
};

const balanceCols = [
  { key: 'id', title: 'ID', width: 80 },
  {
    key: 'changePrice_dir',
    title: '方向',
    width: 70,
    render: (row: any) => h('span', { class: row.changePrice > 0 ? 'member-text-success' : 'member-text-danger' }, row.changePrice > 0 ? '发放' : '消耗'),
  },
  {
    key: 'scene',
    title: '发生场景',
    width: 80,
    render: (row: any) => sceneLabel(row.scene),
  },
  {
    key: 'changePrice',
    title: '积分变化',
    width: 120,
    render: (row: any) => h('span', { class: row.changePrice > 0 ? 'member-text-success' : 'member-text-danger' }, row.changePrice > 0 ? `+${row.changePrice}` : String(row.changePrice)),
  },
  {
    key: 'operator',
    title: '操作人',
    width: 100,
    render: (row: any) => row.operatorId > 0 ? row.adminMemberUsername : '系统操作',
  },
  {
    key: 'des',
    title: '备注',
    width: 160,
    render: (row: any) => row.des || '--',
  },
  {
    key: 'reason',
    title: '原因',
    width: 160,
    render: (row: any) => row.reason || '--',
  },
  { key: 'createdAt', title: '发生时间', width: 170 },
  {
    key: 'orderSn',
    title: '关联订单号',
    width: 200,
    render: (row: any) => {
      if (!row.orderSn) return '';
      return h('button', {
        class: 'member-table-action',
        onClick: () => openOrderView(row.orderSn),
      }, row.orderSn);
    },
  },
];

// Exp columns
const expCols = [
  { key: 'id', title: 'ID', width: 80 },
  {
    key: 'exp_dir',
    title: '方向',
    width: 70,
    render: (row: any) => h('span', { class: row.exp > 0 ? 'member-text-success' : 'member-text-danger' }, row.exp > 0 ? '发放' : '消耗'),
  },
  {
    key: 'scene',
    title: '发生场景',
    width: 80,
    render: (row: any) => row.scene === 'SYSTEM' ? '系统' : row.scene === 'HOTEL' ? '酒店' : '未知',
  },
  {
    key: 'exp_change',
    title: '成长值变化',
    width: 120,
    render: (row: any) => h('span', { class: row.exp > 0 ? 'member-text-success' : 'member-text-danger' }, row.exp > 0 ? `+${row.exp}` : String(row.exp)),
  },
  {
    key: 'operator',
    title: '操作人',
    width: 100,
    render: (row: any) => row.operatorId > 0 ? row.adminMemberUsername : '系统操作',
  },
  { key: 'des', title: '原因', width: 160 },
  { key: 'createdAt', title: '发生时间', width: 170 },
  { key: 'orderSn', title: '关联订单号', width: 200 },
];

// Hotel order columns
const hotelOrderCols = [
  { key: 'orderSn', title: '内部订单号', width: 180 },
  { key: 'outOrderSn', title: '外部订单号', width: 180 },
  { key: 'checkInDate', title: '入住日期', width: 120 },
  { key: 'checkOutDate', title: '退房日期', width: 120 },
  { key: 'totalAmount', title: '金额', width: 100 },
  { key: 'status', title: '状态', width: 80 },
  { key: 'createdAt', title: '创建时间', width: 160 },
];

// Food order columns
const foodOrderCols = [
  { key: 'orderSn', title: '订单号', width: 180 },
  { key: 'bookDateTime', title: '预约时间', width: 140 },
  { key: 'totalAmount', title: '金额', width: 100 },
  { key: 'status', title: '状态', width: 80 },
  { key: 'createdAt', title: '创建时间', width: 160 },
];

// Spa order columns
const spaOrderCols = [
  { key: 'orderSn', title: '订单号', width: 180 },
  { key: 'bookDateTime', title: '预约时间', width: 140 },
  { key: 'totalAmount', title: '金额', width: 100 },
  { key: 'status', title: '状态', width: 80 },
  { key: 'createdAt', title: '创建时间', width: 160 },
];

// Car order columns
const carOrderCols = [
  { key: 'orderSn', title: '订单号', width: 180 },
  { key: 'pickupDateTime', title: '接机时间', width: 140 },
  { key: 'totalAmount', title: '金额', width: 100 },
  { key: 'status', title: '状态', width: 80 },
  { key: 'createdAt', title: '创建时间', width: 160 },
];

// Cabinet order columns
const cabinetOrderCols = [
  { key: 'orderSn', title: '订单号', width: 180 },
  { key: 'startDate', title: '开始日期', width: 120 },
  { key: 'endDate', title: '结束日期', width: 120 },
  { key: 'totalAmount', title: '金额', width: 100 },
  { key: 'status', title: '状态', width: 80 },
  { key: 'createdAt', title: '创建时间', width: 160 },
];

// Travel order columns
const travelOrderCols = [
  { key: 'orderSn', title: '订单号', width: 180 },
  { key: 'travelDate', title: '出游日期', width: 120 },
  { key: 'totalAmount', title: '金额', width: 100 },
  { key: 'status', title: '状态', width: 80 },
  { key: 'createdAt', title: '创建时间', width: 160 },
];

// Coupon columns
const couponCols = [
  {
    key: 'couponName',
    title: '优惠券',
    width: 200,
    render: (row: any) => row.pmsCouponType?.couponName || '--',
  },
  {
    key: 'type',
    title: '类型',
    width: 80,
    render: (row: any) => row.type === 'reward' ? '满减券' : '折扣券',
  },
  {
    key: 'coupon_amount',
    title: '金额/折扣',
    width: 110,
    render: (row: any) => row.type === 'reward' ? `${row.money}JPY` : `${row.discount}折`,
  },
  { key: 'atLeast', title: '满额使用', width: 100 },
  {
    key: 'state',
    title: '状态',
    width: 80,
    render: (row: any) => {
      const map: Record<number, string> = { 1: '已领取', 2: '已使用', 3: '已过期', 5: '已回收' };
      return map[row.state] || '已关闭';
    },
  },
  { key: 'fetchTime', title: '领取时间', width: 160 },
  { key: 'useTime', title: '使用时间', width: 160 },
];

// TH coupon columns
const thCouponCols = [
  { key: 'couponName', title: '礼品券名称', width: 180 },
  { key: 'couponCode', title: '券码', width: 160 },
  {
    key: 'state',
    title: '状态',
    width: 80,
    render: (row: any) => {
      const map: Record<number, string> = { 1: '未使用', 2: '已使用', 3: '已过期', 5: '已回收' };
      return map[row.state] || '--';
    },
  },
  { key: 'expireTime', title: '过期时间', width: 160 },
  { key: 'createdAt', title: '领取时间', width: 160 },
];

// Junior columns
const juniorCols = [
  { key: 'memberNo', title: '会员号', width: 150 },
  { key: 'firstName', title: '名', width: 100 },
  { key: 'lastName', title: '姓', width: 100 },
  { key: 'phone', title: '手机号', width: 140 },
  { key: 'createdAt', title: '注册时间', width: 160 },
];

// Data loading functions
const loadBalanceDataTable = async (res: any) => {
  res.memberId = params.id;
  return await BalanceList({ ...res });
};

const loadExpDataTable = async (res: any) => {
  res.memberId = params.id;
  return await ExpList({ ...res });
};

const loadHotelOrderDataTable = async (res: any) => {
  res.member_id = params.id;
  return await HotelOrderList({ ...res });
};

const loadFoodOrderDataTable = async (res: any) => {
  res.memberId = params.id;
  return await FoodOrderList({ ...res });
};

const loadSpaOrderDataTable = async (res: any) => {
  res.memberId = params.id;
  return await SpaOrderList({ ...res });
};

const loadCarOrderDataTable = async (res: any) => {
  res.memberId = params.id;
  return await CarOrderList({ ...res });
};

const loadCabinetOrderDataTable = async (res: any) => {
  res.memberId = params.id;
  return await CabinetOrderList({ ...res });
};

const loadTravelOrderDataTable = async (res: any) => {
  res.memberId = params.id;
  return await TravelOrderList({ ...res });
};

const loadCouponDataTable = async (res: any) => {
  res.memberId = params.id;
  return await CouponList({ ...res });
};

const loadThCouponDataTable = async (res: any) => {
  res.memberId = params.id;
  return await ThCouponList({ ...res });
};

const loadJuniorDataTable = async (res: any) => {
  res.id = params.id;
  res.isJunior = 1;
  return await MemberList({ ...res });
};

// Order detail views
function openOrderView(orderSn: string) {
  if (!orderSn) return;
  const first = orderSn.charAt(0);
  if (first === 'H') handleHotelOrderView({ orderSn, outOrderSn: '' });
  else if (first === 'F') handleFoodOrderView({ orderSn });
  else if (first === 'S') handleSpaOrderView({ orderSn });
  else if (first === 'C') handleCarOrderView({ orderSn });
  else if (first === 'B') handleCabinetOrderView({ orderSn });
}

function handleHotelOrderView(record: any) { hotelOrderViewRef.value?.openModal(record); }
function handleFoodOrderView(record: any) { foodOrderViewRef.value?.openModal(record); }
function handleCarOrderView(record: any) { carOrderViewRef.value?.openModal(record); }
function handleCabinetOrderView(record: any) { cabinetOrderViewRef.value?.openModal(record); }
function handleTravelOrderView(record: any) { travelOrderViewRef.value?.openModal(record); }
function handleSpaOrderView(record: any) { spaOrderViewRef.value?.openModal(record); }

function handleEditBase(type: string, value: any) {
  editBaseRef.value?.openModal(data.value.id, type, value);
}
function handleEditBalance(value: any) {
  editBalanceRef.value?.openModal(data.value.id, value);
}
function handleEditExp(value: any) {
  editExpRef.value?.openModal(data.value.id, value);
}

function reloadBalanceInfo() {
  getInfo();
  actionBalanceRef.value?.reload();
}

function reloadExpInfo() {
  getInfo();
  actionExpRef.value?.reload();
}

function reloadCouponTable() {
  actionCouponRef.value?.reload();
}

function reloadThCouponTable() {
  actionThCouponRef.value?.reload();
}

function handleCouponRecycle(record: any) {
  showConfirm('警告', '回收后用户无法再使用，确认回收？', () => {
    CouponRecycle({ id: record.id }).then(() => {
      message.success('操作成功');
      reloadCouponTable();
    });
  });
}

function handleThCouponRecycle(record: any) {
  showConfirm('警告', '回收后用户无法再使用，确认回收？', () => {
    ThCouponRecycle({ id: record.id }).then(() => {
      message.success('操作成功');
      reloadThCouponTable();
    });
  });
}

const getInfo = () => {
  loading.value = true;
  View(params)
    .then((res) => {
      if ((res.referrer <= 0 && res.lastReferrer <= 0) || (!res.referrerDetail && !res.LastReferrerDetail)) {
        res.referrerName = '无';
      } else {
        if (recommendModel.value === 'FIRST') {
          if (!res.referrerDetail) {
            res.referrerName = '无';
          } else {
            const mode = res.referrerDetail.rebateMode;
            const no = res.referrerDetail.memberNo;
            res.referrerName = mode === 'CHANNEL' ? `渠道(${no})` : mode === 'STAFF' ? `员工(${no})` : `会员(${no})`;
          }
        } else {
          if (!res.LastReferrerDetail) {
            res.referrerName = '无';
          } else {
            const mode = res.LastReferrerDetail.rebateMode;
            const no = res.LastReferrerDetail.memberNo;
            res.referrerName = mode === 'CHANNEL' ? `渠道(${no})` : mode === 'STAFF' ? `员工(${no})` : `会员(${no})`;
          }
        }
      }
      res.levelName = res.memberLevel?.levelName || '--';
      res.groupName = res.memberGroup ? res.memberGroup.memberGroup : '--';
      data.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
};

onMounted(() => {
  if (!params.id) {
    message.error('会员ID不正确，请检查！');
    return;
  }
  getConfig({ group: 'yyconfig' }).then((res) => {
    recommendModel.value = res.list.recommendModel;
  });
  getInfo();
});
</script>

<style lang="less" scoped>
.member-detail-page {
  --member-primary: #38aeea;
  --member-primary-strong: #128fc8;
  --member-primary-soft: #eaf7ff;
  --member-border: #dce7f2;
  --member-surface-soft: #f7fbff;
  --member-text: #152033;
  --member-muted: #6b7c93;
  --member-danger: #ef4444;
  --member-success: #16a34a;
}

.member-detail-grid {
  grid-template-columns: minmax(0, 1fr) 260px;

  @media (max-width: 1180px) {
    grid-template-columns: 1fr;
  }
}

.member-card {
  border-color: var(--member-border);
  color: var(--member-text);
}

.member-section-title {
  margin-bottom: 14px;
  color: var(--member-text);
  font-size: 15px;
  font-weight: 600;
  line-height: 22px;
}

.member-info-grid {
  border-color: var(--member-border);
}

:deep(.member-info-item) {
  display: grid;
  grid-template-columns: 88px minmax(0, 1fr);
  align-items: start;
  gap: 8px;
  min-height: 28px;
  font-size: 13px;
  line-height: 20px;
}

:deep(.member-info-label) {
  color: var(--member-muted);
  white-space: nowrap;
}

:deep(.member-info-label::after) {
  content: ":";
}

:deep(.member-info-value) {
  display: flex;
  min-width: 0;
  flex-wrap: wrap;
  align-items: center;
  gap: 6px;
  color: var(--member-text);
  overflow-wrap: anywhere;
}

.member-account-row {
  display: grid;
  grid-template-columns: 88px minmax(0, 1fr) auto;
  align-items: center;
  gap: 8px;
  min-height: 30px;
  font-size: 13px;
  line-height: 20px;

  span {
    color: var(--member-muted);
  }

  strong {
    color: var(--member-text);
    font-weight: 600;
  }
}

.member-inline-action,
.member-table-action {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-height: 24px;
  border: 0;
  border-radius: 5px;
  color: var(--member-primary-strong);
  background: transparent;
  font-size: 12px;
  line-height: 18px;
  cursor: pointer;
  transition: color 0.15s ease, background 0.15s ease;

  &:hover {
    color: #0879ad;
    background: var(--member-primary-soft);
  }
}

.member-table-action {
  min-height: 26px;
  padding: 0 7px;
  font-size: 13px;
}

.member-table-action-danger {
  color: var(--member-danger);

  &:hover {
    color: #dc2626;
    background: #fff1f2;
  }
}

:deep(.member-text-success) {
  color: var(--member-success);
  font-weight: 500;
}

:deep(.member-text-danger) {
  color: var(--member-danger);
  font-weight: 500;
}

:deep(.bg-muted) {
  background: var(--member-primary-soft);
}

:deep(.text-primary) {
  color: var(--member-primary-strong);
}
</style>
