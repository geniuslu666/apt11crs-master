import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { usePermission } from '@/hooks/web/usePermission';
const { hasPermission } = usePermission();
const $message = window['$message'];
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {isNullObject} from "@/utils/is";
import {NTag} from "naive-ui";

export class State {
  public id = 0; // id
  public sceneName = ''; // 场景名称
  public isGetOpen = 2; // 是否允许积分获取
  public isPayOpen = 2; // 是否允许积分抵扣
  public isLimitMoney = 1; // 是否限制金额
  public limitMoney = null; // 限制金额
  public getRate = null; // 获得积分的抵扣比例
  public isPayRate = 1; // 积分抵现上限（开关）
  public payRate = null; // 能够使用的总金额比例积分
  public isOpenReward = 2;
  public rewardType = '';
  public rewardCouponTypeIds = '';
  public rewardThCouponIds = '';
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 修改时间
  public deletedAt = ''; // 删除时间

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

// 表单验证规则
export const rules = {
  limitMoney: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'number',
    validator: validate.amount,
  },
  getRate: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'number',
    validator: validate.percentage,
  },
  payRate: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'number',
    validator: validate.percentage,
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'sceneName',
    component: 'NInput',
    label: '场景名称',
    componentProps: {
      placeholder: '请输入场景名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '场景名称',
    key: 'sceneName',
    align: 'left',
    width: -1,
  },
  {
    title: '是否允许积分获取',
    key: 'isGetOpen',
    align: 'left',
    width: -1,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.isGetOpen == 'Y' ? 'green' : 'red'
          },
        },
        {
          default: () => row.isGetOpen == 'Y' ? '允许' : '不允许'
        }
      )
    },
  },
  {
    title: '是否允许积分抵扣',
    key: 'isPayOpen',
    align: 'left',
    width: -1,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.isPayOpen == 'Y' ? 'green' : 'red'
          },
        },
        {
          default: () => row.isPayOpen == 'Y' ? '允许' : '不允许'
        }
      )
    },
  },
  {
    title: '限制金额',
    key: 'limitMoney',
    align: 'left',
    width: -1,
  },
  {
    title: '积分抵扣比例',
    key: 'getRate',
    align: 'left',
    width: -1,
  },
  {
    title: '能使用的总金额比例积分',
    key: 'payRate',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 165,
  },
];

// 字典数据选项
export const options = ref({
  language: [] as Option[],
  category: []
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['language'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'language':
          item.componentProps.options = options.value.language;
          break;
      }
    }
  });
}
