import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { NTag } from 'naive-ui';
import { isNullObject } from '@/utils/is';

export class State {
  public id = 0;
  public productId = 0;
  public nameLanguage = null;
  public price = 0;
  public dailyCapacity = 1;
  public meetingPlace = '';
  public meetingTime = null;
  public ggLat = '';
  public ggLng = '';
  public contactMobile = '';
  public status = 1;
  public sort = 0;
  public createdAt = '';
  public updatedAt = '';
  public allowCancel = 1; // 是否允许取消
  public freeCancelHours = 0; // 几小时前免费取消
  public cancelFeePercent = 0; // 取消手续费（百分比）

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

export const rules = {
  price: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入价格',
  },
  dailyCapacity: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入每日接待人数',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '车型名称',
    componentProps: {
      placeholder: '请输入车型名称',
    },
  },
]);

export const columns = [
  {
    title: '车型名称',
    key: 'name',
    align: 'left',
    width: 200,
  },
  {
    title: '价格（JPY）',
    key: 'price',
    align: 'left',
    width: 120,
    render(row) {
      return row.price + ' JPY';
    },
  },
  {
    title: '每日接待人数',
    key: 'dailyCapacity',
    align: 'left',
    width: 130,
  },
  {
    title: '排序',
    key: 'sort',
    align: 'left',
    width: 80,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.status)) return '';
      return h(
        NTag,
        { type: row.status === 1 ? 'success' : 'error', bordered: false },
        { default: () => (row.status === 1 ? '启用' : '禁用') }
      );
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
];
