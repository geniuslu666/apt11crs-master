import { h, ref } from 'vue';
import { NTag, SelectRenderLabel} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '订单编号',
    componentProps: {
      placeholder: '订单编号',
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
    title: '操作',
    key: 'actionWay',
    align: 'left',
    width: 170,
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
          type: getOptionTag(options.value.hotel_log_action_way, row.actionWay),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.hotel_log_action_way, row.actionWay),
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
    render(row){
      if(row.actionWay == "CHANGE_GUEST"){
        return "入住人信息变更"
      }else if(row.actionWay == "CHANGE_PEOPLE_NUM"){
        return "入住人数变更"
      }else if(row.actionWay == "CHANGE_DATE"){
        return "入住日期变更"
      }else if(row.actionWay == "ORDER_CONFIRMED"){
        return "订单已确认"
      }else if(row.actionWay == "ORDER_CHECKED_IN"){
        return "客人已入住"
      }else if(row.actionWay == "ORDER_CHECKED_OUT"){
        return "客人已退房"
      }else if(row.actionWay == "ORDER_BLOCKED"){
        return "订单已阻止"
      }else if(row.actionWay == "ORDER_USER_CANCELLED"){
        return "订单已取消"
      }else if(row.actionWay == "ORDER_PENDING"){
        return "订单待确认"
      }else if(row.actionWay == "ORDER_OVERLAPPED"){
        return "订单重叠"
      }else if(row.actionWay == "ORDER_CANCELLED"){
        return "订单已取消"
      }
      return  row.remark
    }
  },
  {
    title: '操作角色',
    key: 'operateType',
    align: 'left',
    width: -1,
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
  hotel_log_action_way: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: [ 'hotel_log_action_way'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        // case 'order_status':
        //   item.componentProps.options = options.value.car_log_order_status;
        //   break;
      }
    }
  });
}


