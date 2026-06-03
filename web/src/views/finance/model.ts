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
    field: 'searchTime',
    label: '发生时间',
    slot: 'searchTimeSlot',
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
