import { h, ref } from 'vue';
import { NTag, SelectRenderLabel, NTooltip, NIcon} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
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
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '订单号',
    key: 'name',
    align: 'left',
    width: 180,
    render(row){
      return row.orderDetail.orderSn
    }
  },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: 120,
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
          type: getOptionTag(options.value.car_log_order_status, row.orderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.car_log_order_status, row.orderStatus),
        }
      );
    },
  },
  {
    title: '操作',
    key: 'actionWay',
    align: 'left',
    width: 150,
    render(row) {
      if (isNullObject(row.actionWay)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.car_log_action_way, row.actionWay),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.car_log_action_way, row.actionWay),
        }
      );
    },
  },
  {
    title: '备注',
    key: 'remark',
    align: 'left',
    width: 200,
    ellipsis: false,
  },
  {
    title: '操作角色',
    key: 'operateType',
    align: 'left',
    width: 100,
  },
  {
    title: '操作人',
    key: 'operateId',
    align: 'left',
    width: -1,
    render(row){
      return  row.operateName
    }
  },
  {
    title: '操作时间',
    key: 'createdAt',
    align: 'left',
    width: 160,
  },
];

export const renderMemberLabel: SelectRenderLabel = (option) => {
  return option.memberNo + ' | ' + option.fullName + ' | ' + option.phone
};

// 字典数据选项
export const options = ref({
  car_log_order_status: [] as Option[],
  car_log_action_way: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: [ 'car_log_order_status', 'car_log_action_way'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'orderStatus':
          item.componentProps.options = options.value.car_log_order_status;
          break;
      }
    }
  });
}


