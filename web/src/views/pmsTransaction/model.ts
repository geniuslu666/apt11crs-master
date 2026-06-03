import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defShortcuts, defRangeShortcuts } from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";
import {NTag} from "naive-ui";

export class State {
  public id = 0; // 主键
  public orderSn = ''; // 订单号
  public transactionSn = ''; // 支付流水号
  public paymentRequestId = ''; // 第三方支付流水号
  public payChannel = 'SYSTEM'; // SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台
  public payType = 'BAL'; // 支付方式   BAL 余额
  public amount = null; // 总金额
  public payParams = ''; // 支付参数
  public priceCurrency = ''; // 币种
  public payAmount = null; // 支付金额
  public payStatus = 'WAIT'; // 支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付
  public payTime = ''; // 支付时间
  public expiredTime = ''; // 过期时间
  public refundAmount = null; // 退款金额
  public refundStatus = 'WAIT'; // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
  public scenePayRate = null; // 场景积分抵扣比例
  public level = 0; // 等级
  public exchangeRate = null; // 积分汇率
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
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
    field: 'payChannel',
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
    field: 'payType',
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
    field: 'payStatus',
    component: 'NSelect',
    label: '支付状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择支付状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'payTime',
    label: '支付时间',
    slot: 'payTimeSlot',
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
    width: 60,
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
    title: '第三方支付流水号',
    key: 'paymentRequestId',
    align: 'left',
    width: 180,
  },
  {
    title: '支付平台',
    key: 'payChannel',
    align: 'left',
    width: 150,
    render(record){
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.pay_channel, record.payChannel),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.pay_channel, record.payChannel),
        }
      );
    }
  },
  {
    title: '支付方式',
    key: 'payType',
    align: 'left',
    width: 115,
    render(record){
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.pay_type, record.payType),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.pay_type, record.payType),
        }
      );
    }
  },
  {
    title: '总金额',
    key: 'amount',
    align: 'left',
    width: 100,
  },
  {
    title: '币种',
    key: 'priceCurrency',
    align: 'left',
    width: 60,
  },
  {
    title: '支付金额',
    key: 'payAmount',
    align: 'left',
    width: 100,
  },
  {
    title: '支付状态',
    key: 'payStatus',
    align: 'left',
    width: 80,
    render(record) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.pay_status, record.payStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.pay_status, record.payStatus),
        }
      );
    }
  },
  {
    title: '支付时间',
    key: 'payTime',
    align: 'left',
    width: 170,
  },
  {
    title: '过期时间',
    key: 'expiredTime',
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
    title: '退款状态',
    key: 'refundStatus',
    align: 'left',
    width: 100,
    render(record) {
      if(record.refundStatus == 'WAIT'){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'default',
            bordered: false,
          },
          {
            default: () => '未退款',
          }
        );
      }
      if(record.refundStatus == 'PART'){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'warning',
            bordered: false,
          },
          {
            default: () => '部分退款',
          }
        );
      }
      if(record.refundStatus == 'DONE'){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'error',
            bordered: false,
          },
          {
            default: () => '全额退款',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '场景积分抵扣比例',
    key: 'scenePayRate',
    align: 'left',
    width: 120,
  },
  {
    title: '积分汇率',
    key: 'exchangeRate',
    align: 'left',
    width: 100,
  },
  {
    title: '创建/更新时间',
    key: 'phone',
    align: 'left',
    width: 220,
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
  pay_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['pay_channel', 'pay_type', 'pay_status'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'payChannel':
          item.componentProps.options = options.value.pay_channel;
          break;
        case 'payType':
          item.componentProps.options = options.value.pay_type;
          break;
        case 'payStatus':
          item.componentProps.options = options.value.pay_status;
          break;
      }
    }
  });
}
