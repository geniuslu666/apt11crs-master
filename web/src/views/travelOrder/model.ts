import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { Option } from '@/utils/hotgo';

export class State {
  public id = 0;
  public orderSn = '';
  public productId = 0;
  public memberId = 0;
  public bookingName = '';
  public phoneArea = '';
  public bookingMobile = '';
  public bookingNum = 1;
  public bookingEmail = '';
  public bookDate = '';
  public orderAmount = 0;
  public orderStatus = '';
  public verifyStaffId = 0;
  public verifyStaffName = '';
  public verifyTime = '';
  public payStatus = 0;
  public payTime = '';
  public cancelTime = '';
  public cancelFee = 0;
  public refundStatus = '';
  public refundAmount = 0;
  public refundTime = '';
  public createdAt = '';
  public updatedAt = '';
  public productInfo = {
    'id': 0,
    'title': '',
    'meetingPlace': '',
    'meetingTime': '',
  };
  public skuInfo = {
    'id': 0,
    'name': '',
  };
  public memberDeleted = false;

  constructor(state?: Partial<State>) {
    if (state) Object.assign(this, state);
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) return cloneDeep(state);
    return new State(state);
  }
  return new State();
}

// const ORDER_STATUS_OPTIONS = [
//   { label: '全部', value: '' },
//   { label: '待支付', value: 'WAIT_PAY' },
//   { label: '待核销', value: 'WAIT_VERIFY' },
//   { label: '已完成', value: 'DONE' },
//   { label: '已取消', value: 'CANCEL' },
//   { label: '已退款', value: 'REFUND' },
// ];

export const ORDER_STATUS_MAP = {
  WAIT_PAY: { label: '待支付', type: 'warning' },
  WAIT_VERIFY: { label: '待核销', type: 'info' },
  DONE: { label: '已完成', type: 'success' },
  CANCEL: { label: '已取消', type: 'default' },
  REFUND: { label: '已退款', type: 'error' },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '预约单号',
    componentProps: { placeholder: '请输入预约单号' },
  },
  {
    field: 'productName',
    component: 'NInput',
    label: '产品名称',
    componentProps: { placeholder: '请输入产品名称' },
  },
  {
    field: 'skuName',
    component: 'NInput',
    label: '车型名称',
    componentProps: { placeholder: '请输入车型名称' },
  },
  {
    field: 'memberSearch',
    component: 'NInput',
    label: '客户信息',
    componentProps: {
      placeholder: '请输入预订人姓名/手机号/用户ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookDate',
    label: '预定日期',
    slot: 'bookDateSlot',
  },
  {
    field: 'verifyTime',
    label: '核销时间',
    slot: 'verifyTimeSlot',
  },
  {
    field: 'createdAt',
    label: '下单时间',
    slot: 'createdAtSlot',
  },
  // {
  //   field: 'orderStatus',
  //   component: 'NSelect',
  //   label: '订单状态',
  //   componentProps: { placeholder: '请选择状态', options: ORDER_STATUS_OPTIONS },
  // },
]);

// 字典数据选项
export const options = ref({
  travel_order_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['travel_order_status'],
  }).then((res) => {
    options.value = res;
  });
}
