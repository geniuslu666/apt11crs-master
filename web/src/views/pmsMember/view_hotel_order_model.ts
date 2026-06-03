import {h, ref} from 'vue';
import { NTag } from 'naive-ui';
import { FormSchema } from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";

// 表格搜索表单
export const hotelOrderSchemas = ref<FormSchema[]>([
  {
    field: 'keywordOrderSn',
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
export const hotelOrderColumns = [
  {
    title: '内部订单号',
    key: 'orderSn',
    align: 'left',
    width: 200,
  },
  {
    title: '外部订单号',
    key: 'outOrderSn',
    align: 'left',
    width: 200,
  },
  {
    title: '订单信息',
    key: 'id',
    align: 'left',
    width: 250,
    render(row){
      return h(
        'div',
        [
          h(
            'div',
            {},
            {
              default: () => row.propertyDetail.name
            }
          ),
          h(
            'div',
            {},
            {
              default: () => row.checkInDate + ' -> ' + row.checkOutDate
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
    width: 130,
    render(row){
      return h(
        NTag,
        {
          size: "small",
          class: "min-left-space",
          type: getOptionTag(optionsstatus.value.order_status, row.orderStatus)
        },
        {
          default: () => getOptionLabel(optionsstatus.value.order_status, row.orderStatus),
        }
      )
    }
  },
  {
    title: '订单金额',
    key: 'orderAmount',
    align: 'left',
    width: 150,
    render(row) {
      return h(
        'div',
        [
          row.orderAmount,
          h(
            'span',
            {
              class: "c999",
              style: {
                marginLeft: "5px"
              }
            },
            {
              default: () => "JPY"
            }
          )
        ]
      )
    },
  },
  {
    title: '下单时间',
    key: 'createdAt',
    align: 'left',
    width: 200,
  },
];

export const optionsstatus = ref({
  order_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['order_status'],
  }).then((res) => {
    optionsstatus.value = res;
  });
}

