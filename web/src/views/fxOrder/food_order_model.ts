import { h, ref } from 'vue';
import { NTag} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import {defRangeShortcuts, defShortcuts} from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '系统单号',
    componentProps: {
      placeholder: '请输入系统订单号',
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
      type: 'date',
      clearable: true,
      shortcuts: defShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'checkoutDate',
    component: 'NDatePicker',
    label: '退房日期',
    componentProps: {
      type: 'date',
      clearable: true,
      shortcuts: defShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
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
    key: 'orderSn',
    align: 'left',
    width: 300,
    render(row){
      let guestName,phoneArea,phone,propertyName = ''
      if(row.guestProfileDetail){
        guestName = row.guestProfileDetail.fullName
        phoneArea = row.guestProfileDetail.areaNo
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
                  default: () => "  " + phoneArea + ' ' + phone,
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
        ]
      );
    }
  },
  {
    title: '推荐人',
    key: 'checkInDate',
    align: 'left',
    width: -1,
    render(row) {
      if(!row.referrerDetail){
        return '无'
      }else{
        return row.referrerDetail.fullName
      }
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
    title: '预定状态',
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
    title: '订单金额',
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
          default: () => row.refundAmount + "JPY",
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


