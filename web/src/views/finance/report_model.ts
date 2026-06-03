import { ref } from 'vue';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import {Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";


// 表格搜索表单
export const schemas = ref<FormSchema[]>([

  {
    field: 'payChannel',
    component: 'NSelect',
    label: '支付平台',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择支付平台',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'payType',
    component: 'NSelect',
    label: '交易方式',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择交易方式',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'timeRange',
    label: '发生时间',
    slot: 'timeRangeSlot',
  },
]);

// 字典数据选项
export const options = ref({
  pay_channel: [] as Option[],
  pay_type: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['pay_channel', 'pay_type'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'payChannel':
          item.componentProps.options = options.value.pay_channel;
          break;
        case 'payType':
          item.componentProps.options = options.value.pay_type;
          break;
      }
    }
  });
}
