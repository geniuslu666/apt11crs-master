import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public name = ''; // 名称
  public key = null; // 配置项
  public value = ''; // 配置信息
  public language = null; // 语言
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at
  public deletedAt = ''; // deleted_at

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  // {
  //   field: 'id',
  //   component: 'NInputNumber',
  //   label: 'id',
  //   componentProps: {
  //     placeholder: '请输入id',
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'name',
    component: 'NInput',
    label: '名称',
    componentProps: {
      placeholder: '请输入名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'key',
    component: 'NSelect',
    label: '配置项',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择配置项',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建日期',
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
    title: 'ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '配置项',
    key: 'key',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.key)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.app_config_key, row.key),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.app_config_key, row.key),
        }
      );
    },
  },
  // {
  //   title: '配置信息',
  //   key: 'value',
  //   align: 'left',
  //   width: -1,
  // },
  {
    title: '语言',
    key: 'language',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.language)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.language, row.language),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.language, row.language),
        }
      );
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  // {
  //   title: '更新时间',
  //   key: 'updatedAt',
  //   align: 'left',
  //   width: -1,
  // },
];

// 字典数据选项
export const options = ref({
  language: [] as Option[],
  app_config_key: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['language','app_config_key'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'language':
          item.componentProps.options = options.value.language;
          break;
        case 'app_config_key':
          item.componentProps.options = options.value.language;
          break;
        case 'key':
          item.componentProps.options = options.value.app_config_key;
          break;
      }
    }
  });
}
