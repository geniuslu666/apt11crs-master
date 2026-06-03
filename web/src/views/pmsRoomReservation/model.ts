import {h, ref} from 'vue';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { defRangeShortcuts } from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Option} from '@/utils/hotgo';
import {NTag} from "naive-ui";
import {isNullObject} from "@/utils/is";


// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'source',
    component: 'NSelect',
    label: '订单来源',
    defaultValue: 'APP',
    componentProps: {
      placeholder: '请选择方向',
      options: [
        {
          labelField: 'APP',
          valueField: 'APP',
        },
        {
          labelField: 'AIRHOST',
          valueField: 'AIRHOST',
        }
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
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
    field: 'outOrderSn',
    component: 'NInput',
    label: 'Airhost订单号',
    componentProps: {
      placeholder: '请输入airhost订单号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'propertyName',
    component: 'NInput',
    label: '物业名称',
    componentProps: {
      placeholder: '请输入物业名称',
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
  {
    field: 'orderStatus',
    component: 'NSelect',
    label: '订单状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择订单状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
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
    field: 'checkinStatus',
    component: 'NSelect',
    label: '入住状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择入住状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '订单信息',
    key: 'orderSn',
    align: 'left',
    width: 260,
    resizable: true,
    render(row){
      var guestName = ''
      if(row.guestProfileDetail){
        if(row.guestProfileDetail.fullName){
          guestName = row.guestProfileDetail.fullName + ' # '
        }
      }
      return h(
        'div',
        [
          h(
            'div',
            {
              style: {

              },
            },
            {
              default: () => row.propertyDetail.name,
            }
          ),
          h(
            'div',
            [
              h(
                'span',
                {
                  style: {
                    fontSize: '14px',
                    fontWeight: 'bold',
                    color: '#409eff'
                  }
                },
                {
                  default: () => guestName + row.orderSn
                }
              ),
            ]
          ),
          h(
            'div',
            {
              style: {

              },
            },
            {
              default: () => '成人数：' + row.adultCount + ' 儿童数：' + row.childCount,
            }
          )
        ]
      )
    }
  },
  {
    title: 'airhost订单号',
    key: 'outOrderSn',
    align: 'left',
    width: 200,
  },
  {
    title: '订单来源',
    key: 'sourceName',
    align: 'left',
    width: 120,
  },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: 100,
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
    title: '预约状态',
    key: 'status',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.status, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.status, row.status),
        }
      );
    },
  },
  {
    title: '房间信息',
    key: 'roomInfo',
    align: 'left',
    width: 250,
    ellipsis: false,
    render(row){
      var roomTypeName = '';
      var roomNo = '';
      if(row.roomTypeDetail){
        if(row.roomTypeDetail.name){
          roomTypeName = row.roomTypeDetail.name
        }
      }
      if(row.roomUnitDetail){
        if(row.roomUnitDetail.roomNo){
          roomNo = row.roomUnitDetail.roomNo
        }
      }
      return h(
        'div',
        [
          h(
            'div',
            {
              style: {
                display: "flex",
                alignItems: "center"
              }
            },
            [
              h(
                'span',
                {
                  style: {
                    fontWeight: 'bold'
                  }
                },
                {
                  default: () => '#' + roomNo,
                }
              ),
              h(
                'span',
                {
                  style: {
                    marginLeft: '5px'
                  }
                },
                {
                  default: () => roomTypeName
                }
              )
            ]
          )
        ]
      )
    }
  },
  {
    title: '入住',
    key: 'checkinDate',
    align: 'left',
    width: 100,
  },
  {
    title: '退房',
    key: 'checkoutDate',
    align: 'left',
    width: 100,
  },
  {
    title: '入住状态',
    key: 'checkinStatus',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.checkinStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.checkin_status, row.checkinStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.checkin_status, row.checkinStatus),
        }
      );
    },
  },
  {
    title: '账款',
    key: 'bookingFee',
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
          default: () => row.bookingFee + ' JPY'
        }
      )
    }
  },
];

// 字典数据选项
export const options = ref({
  checkin_status: [] as Option[],
  status: [] as Option[],
  order_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['checkin_status','status','order_status'],
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
      }
    }
  });
}
