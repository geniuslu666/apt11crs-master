import {h, ref} from 'vue';
import {NTag, NTooltip} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";
import {isNullObject} from "@/utils/is";
import {defRangeShortcuts} from "@/utils/dateUtil";

// 表格搜索表单
export const travelOrderSchemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '预约单号',
    componentProps: { placeholder: '请输入预约单号' },
  },
  {
    field: 'productName',
    component: 'NInput',
    label: '产品名称',
    componentProps: { placeholder: '请输入产品名称' },
  },
  {
    field: 'skuName',
    component: 'NInput',
    label: '车型名称',
    componentProps: { placeholder: '请输入车型名称' },
  },
  {
    field: 'memberSearch',
    component: 'NInput',
    label: '客户信息',
    componentProps: {
      placeholder: '请输入预订人姓名/手机号/用户ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookDate',
    component: 'NDatePicker',
    label: '预定日期',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'verifyTime',
    component: 'NDatePicker',
    label: '核销时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
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
export const travelOrderColumns = [
  { title: '预约单号', key: 'orderSn', align: 'left', width: 180 },
  {
    title: '产品',
    key: 'productName',
    align: 'left',
    width: 200,
    render(row){
      return h(
        'div',
        {
          class:'flex-item',
          style: {
            // paddingLeft: '8px',
          }
        },
        [
          h(
            'div',
            {
              style: {
                fontWeight: '400',
                fontSize: '14px',
                // color: '#1664FF',
                lineHeight: '20px'
              }
            },
            {
              default: () => row.productInfo.title,
            }
          ),
          h(
            'div',
            {
              style: {
                fontWeight: '400',
                fontSize: '14px',
                color: 'grey', //#3D3D3D
                lineHeight: '20px'
              }
            },
            {
              default: () => row.skuInfo.name,
            }
          ),
        ]
      )
    }
  },
  {
    title: '预约信息',
    key: 'orderSn',
    align: 'left',
    width: 150,
    render(row){
      let memberHtmlStr = ''
      if(row.memberDeleted){
        memberHtmlStr = h(
          'div',
          {
            style:{
              color: 'red'
            }
          },
          {
            default: () => '用户已注销'
          }
        )
      }
      let htmlStr = h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.bookDate ? row.bookDate.slice(0, 10) : '',
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.bookingName + ' / ' + row.pmsMemberMemberNo
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.phoneArea+row.bookingMobile,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.bookingNum + '人',
            }
          ),
          memberHtmlStr
        ]
      )

      return h(
        NTooltip,
        null,
        {
          trigger:()=>
            htmlStr,
          default: () => htmlStr,
        },
      )
    }
  },
  {
    title: '订单金额',
    key: 'orderAmount',
    align: 'left',
    width: 120,
    render(row) { return row.orderAmount + ' JPY'; },
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
          type: getOptionTag(options.value.travel_order_status, row.orderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.travel_order_status, row.orderStatus),
        }
      );
    },
  },
  {
    title: '核销信息',
    key: 'verify',
    align: 'left',
    width: 200,
    render(row){
      if (row.orderStatus == 'DONE'){
        let htmlStr = h(
          'div',
          null,
          [
            h(
              'div',
              null,
              {
                default: () => row.verifyStaffInfo.name,
              }
            ),
            h(
              'div',
              null,
              {
                default: () => row.verifyTime
              }
            ),
          ]
        )

        return h(
          NTooltip,
          null,
          {
            trigger:()=>
              htmlStr,
            default: () => htmlStr,
          },
        )
      }else{
        return '--';
      }

    }
  },
  { title: '下单时间', key: 'createdAt', align: 'left', width: 160 },
];

// 字典数据选项
export const options = ref({
  travel_order_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['travel_order_status'],
  }).then((res) => {
    options.value = res;
    for (const item of travelOrderSchemas.value) {
      switch (item.field) {
        case 'orderStatus':
                  item.componentProps.options = options.value.travel_order_status;
                  break;
      }
    }
  });
}

