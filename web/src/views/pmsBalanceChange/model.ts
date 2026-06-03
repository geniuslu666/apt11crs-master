import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { usePermission } from '@/hooks/web/usePermission';
const { hasPermission } = usePermission();
const $message = window['$message'];
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';


// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'memberKey',
    component: 'NInput',
    label: '会员信息',
    componentProps: {
      placeholder: '请输入会员名/手机号/邮箱',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'orderNo',
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
    field: 'createdAt',
    label: '发生时间',
    slot: 'createdAtSlot',
  },
]);

// 字典数据选项
export const options = ref({
  language: [] as Option[],
  category: []
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['language'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'language':
          item.componentProps.options = options.value.language;
          break;
      }
    }
  });
}
