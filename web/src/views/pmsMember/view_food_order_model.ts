import {h, ref} from 'vue';
import {NEllipsis, NIcon, NTag, NTooltip} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";
import {isNullObject} from "@/utils/is";
import {QuestionCircleOutlined} from "@vicons/antd";

// 表格搜索表单
export const foodOrderSchemas = ref<FormSchema[]>([
  {
    field: 'orderType',
    component: 'NSelect',
    label: '订单类型',
    defaultValue: '',
    componentProps: {
      placeholder: '请选择订单类型',
      options: [
        {
          label: 'CRS定金模式',
          value: 'CRS'
        },
        {
          label: 'CRS全款模式',
          value: 'CRSALL'
        },
        {
          label: 'TORETA订单',
          value: 'TORETA'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
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
  /*{
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
  },*/
  {
    field: 'bookingStatus',
    component: 'NSelect',
    label: '预约状态',
    defaultValue: '',
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
    field: 'createdAt',
    component: 'NDatePicker',
    label: '下单时间',
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
export const foodOrderColumns = [
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
          )
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
  /*{
    title: '订单号',
    key: 'orderSn',
    align: 'left',
    width: 200,
  },
  {
    title: '下单时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },*/
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
          type: getOptionTag(optionsstatus.value.food_order_status, row.actualOrderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(optionsstatus.value.food_order_status, row.actualOrderStatus),
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
                type: getOptionTag(optionsstatus.value.booking_status, row.bookingStatus),
                bordered: false,
              },
              {
                default: () => getOptionLabel(optionsstatus.value.booking_status, row.bookingStatus),
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
          type: getOptionTag(optionsstatus.value.booking_status, row.bookingStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(optionsstatus.value.booking_status, row.bookingStatus),
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
  /*{
    title: '套餐信息',
    key: 'id',
    align: 'left',
    width: 400,
    ellipsis: false,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.goodsDetail.goodsName,
            }
          ),
          h(
            NEllipsis,
            {
              expandTrigger: 'click',
              lineClamp: 2,
              tooltip: false,
            },
            {
              default: () => row.goodsDetail.goodsContent,
            }
          )
        ]
      )
    }
  },
  {
    title: '定金',
    key: 'depositAmount',
    align: 'left',
    width: 110,
    render(row){
      return h(
        'span',
        {
          style: {
            color: 'red',
            fontWeight: 'bold'
          }
        },
        {
          default: () => row.depositAmount + ' JPY'
        }
      )
    }
  },
  {
    title: '账款',
    key: 'orderAmount',
    align: 'left',
    width: -1,
    render(row){
      return h(
        'span',
        {
          style: {
            color: 'red',
            fontWeight: 'bold'
          }
        },
        {
          default: () => row.orderAmount + ' JPY'
        }
      )
    }
  },*/
];

// 字典数据选项
export const optionsstatus = ref({
  food_order_status: [] as Option[],
  booking_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['food_order_status','booking_status'],
  }).then((res) => {
    optionsstatus.value = res;
  });
}

