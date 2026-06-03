import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defShortcuts, defRangeShortcuts } from '@/utils/dateUtil';
import {NTag} from "naive-ui";
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {options} from "@/views/pmsTransaction/model";
import {Dicts} from "@/api/dict/dict";

export class State {
  public id = 0; // 主键
  public orderSn = ''; // 订单号
  public transactionSn = ''; // 支付流水号
  public refundChannel = ''; // 支付渠道
  public refundType = 'BAL'; // '支付方式   BAL 余额'
  public refundSn = ''; // 退款流水号
  public transNo = ''; // 退款交易号
  public cancelOrderSn = ''; // 取消订单号
  public refundAmount = null; // 退款金额
  public refundTime = ''; // 退款时间
  public refundStatus = 'WAIT'; // 退款状态
  public cancelId = ''; // 取消政策ID
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间

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

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '订单号',
    componentProps: {
      placeholder: '请输入订单号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'transactionSn',
    component: 'NInput',
    label: '支付流水号',
    componentProps: {
      placeholder: '请输入支付流水号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'refundChannel',
    component: 'NSelect',
    label: '支付平台',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择支付平台',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'refundType',
    component: 'NSelect',
    label: '交易方式',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择交易方式',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'refundStatus',
    component: 'NSelect',
    label: '退款状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择退款状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'refundTime',
    label: '退款时间',
    slot: 'refundTimeSlot',
  },
  {
    field: 'createdAt',
    label: '创建时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 90,
  },
  {
    title: '订单号',
    key: 'orderSn',
    align: 'left',
    width: 180,
    resizable: true,
  },
  {
    title: '支付流水号',
    key: 'transactionSn',
    align: 'left',
    width: 190,
  },

  {
    title: '支付平台',
    key: 'refundChannel',
    align: 'left',
    width: 150,
    render(record){
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.pay_channel, record.refundChannel),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.pay_channel, record.refundChannel),
        }
      );
    }
  },
  {
    title: '支付方式',
    key: 'refundType',
    align: 'left',
    width: 115,
    render(record){
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.pay_type, record.refundType),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.pay_type, record.refundType),
        }
      );
    }
  },
  {
    title: '退款流水号',
    key: 'refundSn',
    align: 'left',
    width: 170,
  },
  {
    title: '退款交易号',
    key: 'transNo',
    align: 'left',
    width: 170,
  },
  {
    title: '取消订单号',
    key: 'cancelOrderSn',
    align: 'left',
    width: 170,
  },
  {
    title: '退款金额',
    key: 'refundAmount',
    align: 'left',
    width: 100,
  },
  {
    title: '退款时间',
    key: 'refundTime',
    align: 'left',
    width: 160,
  },
  {
    title: '退款状态',
    key: 'refundStatus',
    align: 'left',
    width: 100,
    render(record) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.refund_status, record.refundStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.refund_status, record.refundStatus),
        }
      );
    }
  },
  {
    title: '创建/更新时间',
    key: 'phone',
    align: 'left',
    width: 200,
    render(row) {
      return h(
        'div',
        [
          h(
            'div',
            [
              h(
                'span',
                {
                  class: 'c999 mr-1'
                },
                {
                  default: () => '创建:',
                }
              ),
              h(
                'span',
                {},
                {
                  default: () => row.createdAt,
                }
              ),
            ]
          ),
          h(
            'div',
            [
              h(
                'span',
                {
                  class: 'c999 mr-1'
                },
                {
                  default: () => '更新:',
                }
              ),
              row.updatedAt ? row.updatedAt : '--',
            ]
          )
        ]
      )
    },
  },
];

// 字典数据选项
export const options = ref({
  pay_channel: [] as Option[],
  pay_type: [] as Option[],
  refund_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['pay_channel', 'pay_type', 'refund_status'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'refundChannel':
          item.componentProps.options = options.value.pay_channel;
          break;
        case 'refundType':
          item.componentProps.options = options.value.pay_type;
          break;
        case 'refundStatus':
          item.componentProps.options = options.value.refund_status;
          break;
      }
    }
  });
}

