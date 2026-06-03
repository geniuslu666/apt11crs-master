import {h, ref} from 'vue';
import {NTag, NTooltip} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";
import {isNullObject} from "@/utils/is";
import {defRangeShortcuts} from "@/utils/dateUtil";

// 表格搜索表单
export const carOrderSchemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '订单编号',
    componentProps: {
      placeholder: '请输入订单编号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'serviceType',
    component: 'NSelect',
    label: '类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择类型',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'orderType',
    component: 'NSelect',
    label: '渠道',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择渠道',
      options: [
        {
          label: '住一自营',
          value: 'CRS'
        },
        {
          label: 'INNN',
          value: 'INNN'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookStartTime',
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
export const carOrderColumns = [
  {
    title: '订单信息',
    key: 'orderSn',
    align: 'left',
    width: 180,
    render(row){
      let htmlStr = ''
      if(row.abnormalStatus == 2){
        htmlStr = h(
          NTag,
          {
            style: {
              marginLeft: '6px',
            },
            type: row.abnormalStatus == 2 ? 'error' : 'warning',
            bordered: false,
            size: 'small'
          },
          {
            default: () => row.abnormalStatus == 2 ? '异常待处理' : '异常已处理',
          }
        )
      }else if(row.abnormalStatus == 3){
        htmlStr = h(
          NTooltip,
          null,
          {
            trigger:()=>
              h(
                NTag,
                {
                  style: {
                    marginLeft: '6px',
                  },
                  type: row.abnormalStatus == 2 ? 'error' : 'warning',
                  bordered: false,
                  size: 'small'
                },
                {
                  default: () => row.abnormalStatus == 2 ? '异常待处理' : '异常已处理',
                }
              ),
            default: () => row.abnormalReason,
          },
        )
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            {
              style:{
                display: 'flex',
                alignItem: 'center'
              }
            },
            [
              h(
                'div',
                null,
                {
                  default: () => row.orderType == 'INNN' ? 'INNN' : '住一自营'
                }
              ),
              htmlStr
            ]
          ),
          h(
            'div',
            null,
            {
              default: () => row.orderSn
            }
          ),
          h(
            'div',
            {
              style:{
                color: 'red'
              }
            },
            {
              default: () => row.orderAmount + 'JPY'
            }
          ),
        ]
      )
    }
  },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: 90,
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
          type: getOptionTag(optionsstatus.value.spa_order_status, row.orderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(optionsstatus.value.spa_order_status, row.orderStatus),
        }
      );
    },
  },
  {
    title: '预约类型',
    key: 'serviceType',
    align: 'left',
    width: 100,
    render(row){
      if(row.serviceType == 'PICKUP'){
        return '接机'
      }else{
        if(row.serviceType == 'DELIVERY'){
          return '送机'
        }else{
          return '包车'
        }
      }
    }
  },
  {
    title: '预约信息',
    key: 'orderSn',
    align: 'left',
    width: 150,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.bookingName + '/' + row.pmsMemberMemberNo
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.phoneArea+row.bookingMobile,
            }
          )
        ]
      )
    }
  },
  {
    title: '用车时间',
    key: 'bookStartTime',
    align: 'left',
    width: 180,
    render(row){
      let datetimeArr = row.bookStartTime.split(' ')
      let dateArr = datetimeArr[0].split('-')
      let timeArr = datetimeArr[1].split(':')
      return parseInt(dateArr[0]) + '-' + parseInt(dateArr[1]) + '-' + parseInt(dateArr[2]) + ' '+parseInt(timeArr[0]) + ':' + timeArr[1]
    }
  },
  {
    title: '线路',
    key: 'address',
    align: 'left',
    width: 220,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => '出发：'+row.startServiceAddress.name
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '到达：'+row.endServiceAddress.name
            }
          )
        ]
      )
    }
  },
  {
    title: '司机',
    key: 'driver',
    align: 'left',
    width: 100,
    render(row){
      if(row.orderType == 'INNN'){
        return '--'
      }else{
        if(row.dispatchStatus == 'WAIT'){
          return '待指派'
        }else{
          return row.driverDetail.nickname
        }
      }
    }
  },
  {
    title: '支付状态',
    key: 'payStatus',
    align: 'left',
    width: 90,
    render(row) {
      if (isNullObject(row.payStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(optionsstatus.value.spa_order_pay_status, row.payStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(optionsstatus.value.spa_order_pay_status, row.payStatus),
        }
      );
    },
  },
  {
    title: '下单时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
];

// 字典数据选项
export const optionsstatus = ref({
  spa_order_status: [] as Option[],
  spa_order_pay_status: [] as Option[],
  car_service_type: [] as Option[],
  driver_work_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['spa_order_status','spa_order_pay_status','car_service_type','driver_work_status'],
  }).then((res) => {
    optionsstatus.value = res;
    for (const item of carOrderSchemas.value) {
      switch (item.field) {
        case 'serviceType':
          item.componentProps.options = optionsstatus.value.car_service_type;
          break;
      }
    }
  });
}

