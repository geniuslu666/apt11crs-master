import {ref} from 'vue';
import { FormSchema } from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {GetMemberOption} from "@/api/org/user";

// 表格搜索表单
export const balanceSchemas = ref<FormSchema[]>([

  {
    field: 'direction',
    component: 'NSelect',
    label: '方向',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择方向',
      options: [
        {
          labelField: '发放',
          valueField: '1',
        },
        {
          labelField: '消耗',
          valueField: '2',
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
    field: 'operatorId',
    component: 'NSelect',
    label: '操作人',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择操作人',
      options: [],
      labelField: 'username',
      valueField: 'value',
      filterable: true,
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
  {
    field: 'des',
    component: 'NInput',
    label: '备注',
    componentProps: {
      placeholder: '请输入备注',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'reason',
    component: 'NInput',
    label: '原因',
    componentProps: {
      placeholder: '请输入原因',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '发生时间',
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

export const operatorList = ref([]);

// 加载字典数据选项
export function loadBalanceOptions() {
  GetMemberOption().then((res) => {
    operatorList.value = res;
    for (const item of balanceSchemas.value) {
      switch (item.field) {
        case 'operatorId':
          item.componentProps.options = operatorList.value;
          break;
      }
    }
  });
}

