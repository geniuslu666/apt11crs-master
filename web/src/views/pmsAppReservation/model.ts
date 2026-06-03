import {h, ref} from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defShortcuts, defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // 主键
  public outOrderSn = ''; // 房间预订ID
  public memberId = 0; // 用户ID
  public puid = ''; // 物业ID
  public orderSn = ''; // 系统订单号
  public ticketId = ''; // 票号
  public roomType = ''; // 房型信息
  public roomUnit = ''; // 房间单元
  public guestProfile = ''; // 预订人信息
  public ratePlanId = ''; // 费率ID
  public mainGuest = ''; // 主要客人信息，参考GuestProfile
  public checkinDate = ''; // 入住日期
  public checkoutDate = ''; // 退房日期
  public checkinTime = ''; // 入住时间
  public checkoutTime = ''; // 退房时间
  public status = 1; // 预订状态
  public checkinStatus = null; // 入住状态
  public orderStatus = null; // 支付状态
  public refundStatus = null; // 退款状态
  public adultCount = 0; // 成人数量
  public childCount = 0; // 儿童数量
  public bookingFee = null; // 预订费
  public channelFee = null; // 渠道费
  public cleaningFee = null; // 清洁费
  public cancellationFee = null; // 取消费
  public charges = ''; // 费用详情
  public createdAt = ''; // 下单时间
  public updatedAt = ''; // 更新时间
  public viewtype = 'App入住订单'; // 更新时间
  public memberDeleted = false;

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
  orderStatus: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入支付状态',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'snKeyword',
    component: 'NInput',
    label: '系统|外部单号',
    componentProps: {
      placeholder: '请输入系统单号|外部单号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'orderStatus',
    component: 'NSelect',
    label: '预订状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择预订状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'refund_status',
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
    field: 'checkinDate',
    component: 'NDatePicker',
    label: '入住日期',
    componentProps: {
      type: 'daterange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  // {
  //   field: 'checkinDate',
  //   component: 'NDatePicker',
  //   label: '入住日期',
  //   componentProps: {
  //     type: 'date',
  //     clearable: true,
  //     shortcuts: defShortcuts(),
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'checkoutDate',
    component: 'NDatePicker',
    label: '退房日期',
    componentProps: {
      type: 'daterange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  // {
  //   field: 'checkoutDate',
  //   component: 'NDatePicker',
  //   label: '退房日期',
  //   componentProps: {
  //     type: 'date',
  //     clearable: true,
  //     shortcuts: defShortcuts(),
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'createdAt',
    label: '下单时间',
    slot: 'createdAtSlot',
  },
  {
    field: 'MemberNo',
    component: 'NInput',
    label: '会员号',
    componentProps: {
      placeholder: '请输入会员号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  // {
  //   title: '物业封面',
  //   key: 'propertyDetail.cover',
  //   align: 'left',
  //   width: 100,
  //   render(row){
  //     return h(NImage, {
  //       width: 50,
  //       height: 70,
  //       src: row.propertyDetail.cover,
  //       style: {
  //         width: '50px',
  //         height: '70px',
  //         'max-width': '100%',
  //         'max-height': '100%',
  //       },
  //     });
  //   }
  // },
  {
    title: '订单信息',
    key: 'orderSn',
    align: 'left',
    width: 300,
    render(row){
      var guestName = ''
      var phone = ''
      var propertyName = ''
      if(row.guestProfileDetail){
        guestName = row.guestProfileDetail.fullName
        phone = row.guestProfileDetail.phone
      }
      if(row.propertyDetail){
        propertyName = row.propertyDetail.name
      }
      return h(
        "div",
          [
            h(
              'p',
              {
                style: {
                  color: '#409eff',
                  margin:'0px'
                },
              },
              {
                default: () => row.orderSn,
              }
            ),
            h(
              'div',
              [
                h(
                  "span",
                  {
                    "style":{
                      margin:'0px',
                      color:'#409eff'
                    }
                  },
                  {
                    default: () => guestName,
                  }
                ),
                h(
                  "span",
                  {
                    "style":{
                      margin:'0px',
                    }
                  },
                  {
                    default: () => "   tel: "+phone,
                  }
                )
              ]
            ),
            h(
              'p',
              {
                "style":{
                  margin:'0px'
                }
              },
              {
                default: () => propertyName,
              }
            ),
            row.memberDeleted ? h(
              'p',
              {
                style: {
                  margin: '0px',
                  color: 'red'
                }
              },
              {
                default: () => '会员已注销',
              }
            ) : null,
          ]
      );
    }
  },
  {
    title: '入住日期',
    key: 'checkInDate',
    align: 'left',
    width: -1
  },
  {
    title: '退房日期',
    key: 'checkOutDate',
    align: 'left',
    width: -1
  },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.orderStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.order_status, row.orderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.order_status, row.orderStatus),
        }
      );
    },
  },
  {
    title: '预订费',
    key: 'orderAmount',
    align: 'left',
    width: -1,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: '#409eff',
          },
        },
        {
          default: () => row.orderAmount + "JPY",
        }
      )
    }
  },
  {
    title: '退款状态',
    key: 'refund_status',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.refundStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.refund_status, row.refundStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.refund_status, row.refundStatus),
        }
      );
    },
  },
  {
    title: '退款金额(含积分)',
    key: 'refundAmount',
    align: 'left',
    width: -1,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: '#409eff',
          },
        },
        {
          default: () => row.refundAmount + " JPY",
        }
      )
    }
  },
  {
    title: '下单时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
];

// 字典数据选项
export const options = ref({
  checkin_status: [] as Option[],
  status: [] as Option[],
  order_status: [] as Option[],
  refund_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['checkin_status', 'status', 'order_status','refund_status'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'checkinStatus':
          item.componentProps.options = options.value.checkin_status;
          break;
        case 'status':
          item.componentProps.options = options.value.status;
          break;
        case 'orderStatus':
          item.componentProps.options = options.value.order_status;
          break;
        case 'refund_status':
          item.componentProps.options = options.value.refund_status;
          break;
      }
    }
  });
}
