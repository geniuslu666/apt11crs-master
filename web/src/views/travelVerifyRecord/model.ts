import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {
  public id = 0;
  public orderId = 0;
  public orderSn = '';
  public productId = 0;
  public productName = '';
  public memberId = 0;
  public bookingName = '';
  public bookingMobile = '';
  public bookDate = '';
  public verifyStaffId = 0;
  public verifyStaffName = '';
  public verifyTime = '';
  public createdAt = '';

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
    label: '活动名称',
    componentProps: { placeholder: '请输入活动名称' },
  },
  {
    field: 'verifyTime',
    component: 'NDatePicker',
    label: '核销时间',
    componentProps: { type: 'datetimerange', clearable: true },
  },
  {
    field: 'bookDate',
    component: 'NDatePicker',
    label: '预约日期',
    componentProps: { type: 'daterange', clearable: true },
  },
]);
