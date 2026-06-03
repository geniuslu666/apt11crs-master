import {h, ref} from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {isNullObject} from "@/utils/is";
import {NEllipsis, NIcon, NTag, NTooltip} from "naive-ui";
import {QuestionCircleOutlined} from "@vicons/antd";

export class State {
  public id = 0; // id
  public orderType = '';
  public orderSn = ''; // 订单编号
  public outOrderSn = ''; // 三方订单号
  public memberId = 0; // 用户ID
  public memberDetail = {
    fullName: '',
    memberNo: ''
  };
  public restaurantId = 0; // 餐厅ID
  public restaurantDetail = {
    name: ''
  };
  public goodsDetail = {
    goodsName: ''
  };
  public orderAmount = 0; // 订单金额
  public name = ''; // 预定人姓名
  public mobile = ''; // 预定人手机
  public email = ''; // 预定人邮箱
  public bookDate = ''; // 预定日期
  public bookTime = ''; // 预定时间
  public oldBookDate = '';// 旧预定日期
  public oldBookTime = '';// 旧预定时间
  public bookingName = '';
  public phoneArea = '';
  public bookingMobile = '';
  public bookingEmail = '';
  public goodsNum = '';
  public bookingCount = '';
  public oldBookingCount = '';  // 旧预定人数
  public seatId = 0; // 座位ID
  public payModel = 3; // 1、余额支付 2、组合支付 3、纯外部支付
  public payTime = ''; // 支付时间
  public actualOrderStatus = '';// 订单状态
  public orderStatus = 'WAIT_PAY'; // 订单付款状态
  public bookingStatus = 'WAIT_CONFIRM'; // 订单预定状态
  public rebateRate = 0; // 分佣比例
  public rebateStatus = 'WAIT'; // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
  public rebateAmount = 0; // 分佣结算金额
  public rebateTime = ''; // 分佣结算时间
  public memberMessage = ''; // 购买人留言信息
  public memberMessageJa = '';
  public adultCount = 0; // 成人数量
  public childCount = 0; // 儿童数量
  public infantCount = 0; // 婴儿数量
  public refundStatus = 'WAIT'; // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
  public refundTime = ''; // 退款时间
  public refundAmount = 0; // 已退款金额
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public verifyStatus = ''; // 核销状态
  public verifyTime = ''; // 核销时间
  public adminCancelNum = 0; // 后台取消次数
  public toretaReservationStatus = 0;

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
    label: '预约单号',
    componentProps: {
      placeholder: '请输入预约单号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'restaurantName',
    component: 'NInput',
    label: '餐厅名称',
    componentProps: {
      placeholder: '请输入餐厅名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  // {
  //   field: 'orderType',
  //   component: 'NSelect',
  //   label: '渠道',
  //   defaultValue: null,
  //   componentProps: {
  //     placeholder: '请选择渠道',
  //     options: [
  //       {
  //         label: '住一自营',
  //         value: 'CRS'
  //       },
  //       {
  //         label: 'TORETA',
  //         value: 'TORETA'
  //       },
  //     ],
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'depositPayStatus',
    component: 'NSelect',
    label: '定金状态',
    defaultValue: 'HAVE_PAID',
    componentProps: {
      placeholder: '请选择定金状态',
      options: [
        {
          label: '已支付',
          value: 'HAVE_PAID'
        },
        {
          label: '待支付',
          value: 'WAIT_PAY'
        },
        {
          label: '已取消',
          value: 'CANCEL'
        },
        {
          label: '已退款',
          value: 'REFUND'
        },
        {
          label: '全部',
          value: 'ALL'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'remainPayStatus',
    component: 'NSelect',
    label: '尾款状态',
    componentProps: {
      placeholder: '请选择尾款状态',
      options: [
        {
          label: '已支付',
          value: 'HAVE_PAID'
        },
        {
          label: '待支付',
          value: 'WAIT_PAY'
        },
        {
          label: '已取消',
          value: 'CANCEL'
        },
        {
          label: '已退款',
          value: 'REFUND'
        },
        {
          label: '全部',
          value: 'ALL'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookingStatus',
    component: 'NSelect',
    label: '预约状态',
    defaultValue: 'WAIT_CONFIRM',
    componentProps: {
      placeholder: '请选择预约状态',
      options: [
        {
          label: '待确认',
          value: 'WAIT_CONFIRM'
        },
        {
          label: '已确认',
          value: 'CONFIRMED'
        },
        {
          label: '已取消',
          value: 'CANCEL'
        },
        {
          label: '全部',
          value: 'ALL'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'goodsName',
    component: 'NInput',
    label: '套餐名称',
    componentProps: {
      placeholder: '请输入套餐名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookDateTime',
    label: '用餐时间',
    slot: 'bookDateTimeSlot',
  },
  {
    field: 'createdAt',
    label: '下单时间',
    slot: 'createdAtSlot',
  },
]);

// 表格搜索表单
export const toretaSchemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '预约单号',
    componentProps: {
      placeholder: '请输入预约单号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'restaurantName',
    component: 'NInput',
    label: '餐厅名称',
    componentProps: {
      placeholder: '请输入餐厅名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'toretaPayStatus',
    component: 'NSelect',
    label: '支付状态',
    defaultValue: 'HAVE_PAID',
    componentProps: {
      placeholder: '请选择支付状态',
      options: [
        {
          label: '已支付',
          value: 'HAVE_PAID'
        },
        {
          label: '待支付',
          value: 'WAIT_PAY'
        },
        {
          label: '已取消',
          value: 'CANCEL'
        },
        {
          label: '全部',
          value: 'ALL'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'goodsName',
    component: 'NInput',
    label: '套餐名称',
    componentProps: {
      placeholder: '请输入套餐名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookDateTime',
    label: '用餐时间',
    slot: 'bookDateTimeSlot',
  },
  {
    field: 'createdAt',
    label: '下单时间',
    slot: 'createdAtSlot',
  },
]);

// 表格搜索表单
export const crsAllSchemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '预约单号',
    componentProps: {
      placeholder: '请输入预约单号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'restaurantName',
    component: 'NInput',
    label: '餐厅名称',
    componentProps: {
      placeholder: '请输入餐厅名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'toretaPayStatus',
    component: 'NSelect',
    label: '支付状态',
    defaultValue: 'HAVE_PAID',
    componentProps: {
      placeholder: '请选择支付状态',
      options: [
        {
          label: '已支付',
          value: 'HAVE_PAID'
        },
        {
          label: '待支付',
          value: 'WAIT_PAY'
        },
        {
          label: '已取消',
          value: 'CANCEL'
        },
        {
          label: '全部',
          value: 'ALL'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookingStatus',
    component: 'NSelect',
    label: '预约状态',
    defaultValue: 'WAIT_CONFIRM',
    componentProps: {
      placeholder: '请选择预约状态',
      options: [
        {
          label: '待确认',
          value: 'WAIT_CONFIRM'
        },
        {
          label: '已确认',
          value: 'CONFIRMED'
        },
        {
          label: '已取消',
          value: 'CANCEL'
        },
        {
          label: '全部',
          value: 'ALL'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'goodsName',
    component: 'NInput',
    label: '套餐名称',
    componentProps: {
      placeholder: '请输入套餐名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookDateTime',
    label: '用餐时间',
    slot: 'bookDateTimeSlot',
  },
  {
    field: 'createdAt',
    label: '下单时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: '订单信息',
    key: 'id',
    align: 'left',
    width: 250,
    render(row){
      var orderType;
      if(row.orderType == "CRS"){
        orderType = "定金模式";
      }else if(row.orderType == "CRSALL"){
        orderType = "全款模式";
      }else{
        orderType = "TORETA";
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            [
              [
                h(
                  'span',
                  {
                    style: {
                      marginRight: '6px',
                    },
                  },
                  {
                    default: () => row.restaurantDetail.name,
                  }
                ),
                h(
                  NTag,
                  {
                    style: {
                      marginRight: '6px',
                    },
                    type: row.orderType == "CRS" ? 'success' : 'warning',
                    bordered: false,
                  },
                  {
                    default: () => orderType,
                  }
                )
              ]
            ]
          ),
          h(
            'div',
            null,
            {
              default: () => '预订人：' + row.bookingName,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '预订人电话：' + row.phoneArea + row.bookingMobile,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '预订人邮箱：' + row.bookingEmail,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '人数：' + (row.bookingCount),
            }
          ),
          row.memberDeleted ? h(
            'div',
            {
              style: {
                color: 'red'
              }
            },
            {
              default: () => '会员已注销',
            }
          ) : null
        ]
      )
    }
  },
  {
    title: '用餐时间',
    key: 'bookDateTime',
    align: 'left',
    sorter: true, // 单列排序
    width: 180,
    render(row){
      return row.bookDate + ' ' + row.bookTime
    }
  },
  // {
  //   title: '下单时间',
  //   key: 'createdAt',
  //   align: 'left',
  //   width: 180,
  // },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: 110,
    render(row) {
      if (isNullObject(row.actualOrderStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.food_order_status, row.actualOrderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.food_order_status, row.actualOrderStatus),
        }
      );
    },
  },
  {
    title: '预约状态',
    key: 'bookingStatus',
    align: 'left',
    width: 120,
    render(row) {
      if (isNullObject(row.bookingStatus)) {
        return ``;
      }

      if(row.bookingStatus == "CANCEL"){
        return h(
          'div',
          {
            style: {
              display: 'flex',
              alignItems: 'center'
            },
          },
          [
            h(
              NTag,
              {
                style: {
                  marginRight: '6px',
                },
                type: getOptionTag(options.value.booking_status, row.bookingStatus),
                bordered: false,
              },
              {
                default: () => getOptionLabel(options.value.booking_status, row.bookingStatus),
              }
            ),
            h(
              NTooltip,
              null,
              {
                trigger:()=>
                  h(
                    NIcon,
                    {
                      size: 20,
                      style: {
                        marginLeft: '5px',
                      },
                    },
                    {
                      default: () => h(QuestionCircleOutlined),
                    }
                  ),
                default: () => row.confirmRefuseReason ? row.confirmRefuseReason : (row.depositCancelReason ? row.depositCancelReason : (row.remainCancelReason ? row.remainCancelReason : '--')),
              },
            )
          ],
        );
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.booking_status, row.bookingStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.booking_status, row.bookingStatus),
        }
      )
    },
  },
  {
    title: '核销状态',
    key: 'verifyStatus',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.verifyStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.verifyStatus == "VERIFIED" ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => row.verifyStatus == "VERIFIED" ? '已核销' : '待核销',
        }
      );
    },
  },
  {
    title: '订单号/下单时间',
    key: 'createdAt',
    sorter: true, // 单列排序
    align: 'left',
    width: 200,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.orderSn,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.createdAt,
            }
          ),
        ]
      )
    }
  },
  // {
  //   title: '套餐信息',
  //   key: 'id',
  //   align: 'left',
  //   width: 400,
  //   ellipsis: false,
  //   render(row){
  //     return h(
  //       'div',
  //       null,
  //       [
  //         h(
  //           'div',
  //           null,
  //           {
  //             default: () => row.goodsDetail.goodsName,
  //           }
  //         ),
  //         h(
  //           NEllipsis,
  //           {
  //             expandTrigger: 'click',
  //             lineClamp: 2,
  //             tooltip: false,
  //           },
  //           {
  //             default: () => row.goodsDetail.goodsContent,
  //           }
  //         )
  //       ]
  //     )
  //   }
  // },
  // {
  //   title: '定金',
  //   key: 'depositAmount',
  //   align: 'left',
  //   width: 110,
  //   render(row){
  //     if(row.orderType == "TORETA"){
  //       return "--";
  //     }
  //     return h(
  //       'span',
  //       {
  //         style: {
  //           color: 'red',
  //           fontWeight: 'bold'
  //         }
  //       },
  //       {
  //         default: () => row.depositAmount + ' JPY'
  //       }
  //     )
  //   }
  // },
  // {
  //   title: '账款',
  //   key: 'orderAmount',
  //   align: 'left',
  //   width: -1,
  //   render(row){
  //     return h(
  //       'span',
  //       {
  //         style: {
  //           color: 'red',
  //           fontWeight: 'bold'
  //         }
  //       },
  //       {
  //         default: () => row.orderAmount + ' JPY'
  //       }
  //     )
  //   }
  // },
];

// 表格列
export const toretaColumns = [
  {
    title: '订单信息',
    key: 'id',
    align: 'left',
    width: 250,
    render(row){
      var orderType;
      if(row.orderType == "CRS"){
        orderType = "定金模式";
      }else if(row.orderType == "CRSALL"){
        orderType = "全款模式";
      }else{
        orderType = "TORETA";
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            [
              [
                h(
                  'span',
                  {
                    style: {
                      marginRight: '6px',
                    },
                  },
                  {
                    default: () => row.restaurantDetail.name,
                  }
                ),
                h(
                  NTag,
                  {
                    style: {
                      marginRight: '6px',
                    },
                    type: row.orderType == "CRS" ? 'success' : 'warning',
                    bordered: false,
                  },
                  {
                    default: () => orderType,
                  }
                )
              ]
            ]
          ),
          h(
            'div',
            null,
            {
              default: () => '预订人：' + row.bookingName,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '预订人电话：' + row.phoneArea +  row.bookingMobile,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '预订人邮箱：' + row.bookingEmail,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '人数：' + (row.bookingCount),
            }
          ),
          row.memberDeleted ? h(
            'div',
            {
              style: {
                color: 'red'
              }
            },
            {
              default: () => '会员已注销',
            }
          ) : null
        ]
      )
    }
  },
  {
    title: '用餐时间',
    key: 'bookDateTime',
    align: 'left',
    sorter: true, // 单列排序
    width: 180,
    render(row){
      return row.bookDate + ' ' + row.bookTime
    }
  },
  // {
  //   title: '下单时间',
  //   key: 'createdAt',
  //   align: 'left',
  //   width: 180,
  // },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: 110,
    render(row) {
      if (isNullObject(row.actualOrderStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.food_order_status, row.actualOrderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.food_order_status, row.actualOrderStatus),
        }
      );
    },
  },
  {
    title: '预约状态',
    key: 'bookingStatus',
    align: 'left',
    width: 120,
    render(row) {
      if (isNullObject(row.bookingStatus)) {
        return ``;
      }

      if(row.bookingStatus == "CANCEL"){
        return h(
          'div',
          {
            style: {
              display: 'flex',
              alignItems: 'center'
            },
          },
          [
            h(
              NTag,
              {
                style: {
                  marginRight: '6px',
                },
                type: getOptionTag(options.value.booking_status, row.bookingStatus),
                bordered: false,
              },
              {
                default: () => getOptionLabel(options.value.booking_status, row.bookingStatus),
              }
            ),
            h(
              NTooltip,
              null,
              {
                trigger:()=>
                  h(
                    NIcon,
                    {
                      size: 20,
                      style: {
                        marginLeft: '5px',
                      },
                    },
                    {
                      default: () => h(QuestionCircleOutlined),
                    }
                  ),
                default: () => row.confirmRefuseReason ? row.confirmRefuseReason : (row.depositCancelReason ? row.depositCancelReason : (row.remainCancelReason ? row.remainCancelReason : '--')),
              },
            )
          ],
        );
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.booking_status, row.bookingStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.booking_status, row.bookingStatus),
        }
      )
    },
  },
  {
    title: '核销状态',
    key: 'verifyStatus',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.verifyStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.verifyStatus == "VERIFIED" ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => row.verifyStatus == "VERIFIED" ? '已核销' : '待核销',
        }
      );
    },
  },
  {
    title: '订单号/下单时间',
    key: 'createdAt',
    align: 'left',
    sorter: true, // 单列排序
    width: 200,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.orderSn,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.createdAt,
            }
          ),
        ]
      )
    }
  },
  // {
  //   title: '套餐信息',
  //   key: 'id',
  //   align: 'left',
  //   width: 400,
  //   ellipsis: false,
  //   render(row){
  //     return h(
  //       'div',
  //       null,
  //       [
  //         h(
  //           'div',
  //           null,
  //           {
  //             default: () => row.goodsDetail.goodsName,
  //           }
  //         ),
  //         h(
  //           NEllipsis,
  //           {
  //             expandTrigger: 'click',
  //             lineClamp: 2,
  //             tooltip: false,
  //           },
  //           {
  //             default: () => row.goodsDetail.goodsContent,
  //           }
  //         )
  //       ]
  //     )
  //   }
  // },
  // {
  //   title: '账款',
  //   key: 'orderAmount',
  //   align: 'left',
  //   width: -1,
  //   render(row){
  //     return h(
  //       'span',
  //       {
  //         style: {
  //           color: 'red',
  //           fontWeight: 'bold'
  //         }
  //       },
  //       {
  //         default: () => row.orderAmount + ' JPY'
  //       }
  //     )
  //   }
  // },
];

// 字典数据选项
export const options = ref({
  food_order_status: [] as Option[],
  booking_status: [] as Option[],
  toreta_order_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['food_order_status','booking_status','toreta_order_status'],
  }).then((res) => {
    options.value = res;
  });
}


