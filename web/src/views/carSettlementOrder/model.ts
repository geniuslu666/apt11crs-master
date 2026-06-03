import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { usePermission } from '@/hooks/web/usePermission';
const { hasPermission } = usePermission();
const $message = window['$message'];
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {All} from "@/api/carDriver";


// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'driverId',
    component: 'NSelect',
    label: '司机',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择司机',
      options: [],
      labelField: 'name',
      valueField: 'id',
      filterable: true,
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
  {
    field: 'orderSn',
    component: 'NInput',
    label: '结算单号',
    componentProps: {
      placeholder: '请输入结算单号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'verifyStatus',
    component: 'NSelect',
    label: '核账状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '待核账',
          valueField: 'WAIT_VERIFY',
        },
        {
          labelField: '已核账',
          valueField: 'VERIFIED',
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
    field: 'verifyTime',
    label: '核账时间',
    slot: 'verifyTimeSlot',
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

// 加载司机
export function loadDriverOptions() {
  All({}).then((res) => {
    console.log(res.list)
    for (const item of schemas.value) {
      switch (item.field) {
        case 'driverId':
          item.componentProps.options = res.list;
          break;
      }
    }
  })
  .finally(() => {

  });
}
